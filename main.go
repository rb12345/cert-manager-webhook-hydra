package main

import (
	"encoding/json"
	"context"
	"fmt"
	"os"

	extapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/antihax/optional"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"github.com/cert-manager/webhook-example/swagger"
)

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		panic("GROUP_NAME must be specified")
	}

	// This will register our custom DNS provider with the webhook serving
	// library, making it available as an API under the provided GroupName.
	// You can register multiple DNS provider implementations with a single
	// webhook, where the Name() method will be used to disambiguate between
	// the different implementations.
	cmd.RunWebhookServer(GroupName,
		&hydraDNSProviderSolver{},
	)
}

// hydraDNSProviderSolver implements the provider-specific logic needed to
// 'present' an ACME challenge TXT record for your own DNS provider.
// To do so, it must implement the `github.com/cert-manager/cert-manager/pkg/acme/webhook.Solver`
// interface.
type hydraDNSProviderSolver struct {
	// If a Kubernetes 'clientset' is needed, you must:
	// 1. uncomment the additional `client` field in this structure below
	// 2. uncomment the "k8s.io/client-go/kubernetes" import at the top of the file
	// 3. uncomment the relevant code in the Initialize method below
	// 4. ensure your webhook's service account has the required RBAC role
	//    assigned to it for interacting with the Kubernetes APIs you need.
	client kubernetes.Clientset
}

// hydraDNSProviderConfig is a structure that is used to decode into when
// solving a DNS01 challenge.
// This information is provided by cert-manager, and may be a reference to
// additional configuration that's needed to solve the challenge for this
// particular certificate or issuer.
// This typically includes references to Secret resources containing DNS
// provider credentials, in cases where a 'multi-tenant' DNS solver is being
// created.
// If you do *not* require per-issuer or per-certificate configuration to be
// provided to your webhook, you can skip decoding altogether in favour of
// using CLI flags or similar to provide configuration.
// You should not include sensitive information here. If credentials need to
// be used by your provider here, you should reference a Kubernetes Secret
// resource and fetch these credentials using a Kubernetes clientset.
type hydraDNSProviderConfig struct {
	// Change the two fields below according to the format of the configuration
	// to be decoded.
	// These fields will be set by users in the
	// `issuer.spec.acme.dns01.providers.webhook.config` field.

	//Email           string `json:"email"`
	//APIKeySecretRef v1alpha1.SecretKeySelector `json:"apiKeySecretRef"`
	HydraBasePath string `json:hydraBasePath"`
	HydraTokenSecretRef cmmeta.SecretKeySelector `json:"hydraTokenSecretRef"`
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource.
// This should be unique **within the group name**, i.e. you can have two
// solvers configured with the same Name() **so long as they do not co-exist
// within a single webhook deployment**.
// For example, `cloudflare` may be used as the name of a solver.
func (c *hydraDNSProviderSolver) Name() string {
	return "hydra-dns01-solver"
}

// Present is responsible for actually presenting the DNS record with the
// DNS provider.
// This method should tolerate being called multiple times with the same value.
// cert-manager itself will later perform a self check to ensure that the
// solver has correctly configured the DNS provider.
func (c *hydraDNSProviderSolver) Present(ch *v1alpha1.ChallengeRequest) error {
	cfg, err := loadConfig(ch.Config)
	if err != nil {
		return err
	}

	fmt.Printf("Decoded configuration %v", cfg)

	if ch.Type != "" && ch.Type != "dns-01" {
		return fmt.Errorf("Unsupported challenge type: '%s'", ch.Type)
	}

	// TODO: add code that sets a record in the DNS provider's console
	swaggerCfg := swagger.NewConfiguration()
	swaggerCfg.BasePath = cfg.HydraBasePath
	hydra := swagger.NewAPIClient(swaggerCfg)

	auth, err := c.getAuthenticationDetails(ch.ResourceNamespace, cfg.HydraTokenSecretRef.LocalObjectReference.Name)
	if err != nil {
		return err
	}
	searchString := fmt.Sprintf("%s type:TXT", ch.ResolvedFQDN)
	searchOpts := swagger.RecordsApiListRecordsOpts{Q: optional.NewString(searchString)}
	fmt.Printf("CleanUp: Querying Hydra for records: %s\n", searchString)
	records, _, err := hydra.RecordsApi.ListRecords(auth, &searchOpts)
	if err != nil {
		return err
	}
	fmt.Printf("%v", records)
	exists := false
	for _, record := range records {
		if record.Hostname == ch.ResolvedFQDN && record.Type_ == "TXT" && record.Content == ch.Key && record.Id != "" {
			exists = true
		}
	}
	if !exists {
		// Create record
		record, _, err := hydra.RecordsApi.CreateRecord(auth, ch.ResolvedFQDN, "TXT", ch.Key, "cert-manager challenge response", 300, "")
		if err != nil {
			return err
		}
		fmt.Printf("Present: Published record with ID %s\n", record.Id)
	}
	return nil
}

// CleanUp should delete the relevant TXT record from the DNS provider console.
// If multiple TXT records exist with the same record name (e.g.
// _acme-challenge.example.com) then **only** the record with the same `key`
// value provided on the ChallengeRequest should be cleaned up.
// This is in order to facilitate multiple DNS validations for the same domain
// concurrently.
func (c *hydraDNSProviderSolver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	cfg, err := loadConfig(ch.Config)
	if err != nil {
		return err
	}
	fmt.Printf("Decoded configuration %v", cfg)

	if ch.Type != "" && ch.Type != "dns-01" {
		return fmt.Errorf("Unsupported challenge type: %s", ch.Type)
	}

	swaggerCfg := swagger.NewConfiguration()
	swaggerCfg.BasePath = cfg.HydraBasePath
	hydra := swagger.NewAPIClient(swaggerCfg)

	auth, err := c.getAuthenticationDetails(ch.ResourceNamespace, cfg.HydraTokenSecretRef.LocalObjectReference.Name)
	if err != nil {
		return err
	}
	searchString := fmt.Sprintf("%s type:TXT", ch.ResolvedFQDN)
	searchOpts := swagger.RecordsApiListRecordsOpts{Q: optional.NewString(searchString)}
	fmt.Printf("CleanUp: Querying Hydra for records: %s\n", searchString)
	records, _, err := hydra.RecordsApi.ListRecords(auth, &searchOpts)
	if err != nil {
		return err
	}
	fmt.Printf("%v", records)
	for _, record := range records {
		if record.Hostname == ch.ResolvedFQDN && record.Type_ == "TXT" && record.Content == ch.Key && record.Id != "" {
			fmt.Printf("CleanUp: Deleting Hydra record ID: %s\n", record.Id)
			_, err := hydra.RecordsApi.RecordsIdDelete(auth, record.Id, &swagger.RecordsApiRecordsIdDeleteOpts{})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Initialize will be called when the webhook first starts.
// This method can be used to instantiate the webhook, i.e. initialising
// connections or warming up caches.
// Typically, the kubeClientConfig parameter is used to build a Kubernetes
// client that can be used to fetch resources from the Kubernetes API, e.g.
// Secret resources containing credentials used to authenticate with DNS
// provider accounts.
// The stopCh can be used to handle early termination of the webhook, in cases
// where a SIGTERM or similar signal is sent to the webhook process.
func (c *hydraDNSProviderSolver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	///// UNCOMMENT THE BELOW CODE TO MAKE A KUBERNETES CLIENTSET AVAILABLE TO
	///// YOUR CUSTOM DNS PROVIDER

	cl, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return err
	}

	c.client = *cl


	///// END OF CODE TO MAKE KUBERNETES CLIENTSET AVAILABLE
	return nil
}

// loadConfig is a small helper function that decodes JSON configuration into
// the typed config struct.
func loadConfig(cfgJSON *extapi.JSON) (hydraDNSProviderConfig, error) {
	cfg := hydraDNSProviderConfig{}
	// handle the 'base case' where no configuration has been provided
	if cfgJSON == nil {
		return cfg, nil
	}
	if err := json.Unmarshal(cfgJSON.Raw, &cfg); err != nil {
		return cfg, fmt.Errorf("error decoding solver config: %v", err)
	}

	return cfg, nil
}

func (c *hydraDNSProviderSolver) getAuthenticationDetails(namespace string, secretName string) (context.Context, error) {
	// TODO: Get Kubernetes secret containing Hydra API token
	ctx := context.Background()
	secret, err := c.client.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return ctx, err
	}
	if secret.Type != corev1.SecretTypeBasicAuth {
		return ctx, fmt.Errorf("Invalid secret type: %s", secret.Type)
	}
	if secret.Data[corev1.BasicAuthUsernameKey] == nil {
		return ctx, fmt.Errorf("Missing username")
	}
	if secret.Data[corev1.BasicAuthPasswordKey] == nil {
		return ctx, fmt.Errorf("Missing password")
	}
	username := string(secret.Data[corev1.BasicAuthUsernameKey])
	password := string(secret.Data[corev1.BasicAuthPasswordKey])

	auth := context.WithValue(context.Background(), swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: username,
		Password: password,
	})
	return auth, nil
}
