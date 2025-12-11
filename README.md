<p align="center">
  <img src="https://raw.githubusercontent.com/cert-manager/cert-manager/d53c0b9270f8cd90d908460d69502694e1838f5f/logo/logo-small.png" height="256" width="256" alt="cert-manager project logo" />
</p>

# Hydra IPAM ACME webhook implementation

This repository contains a basic `cert-manager` webhook implementation that
allows users to respond to ACME DNS01 challenges using the Oxford Hydra IPAM
system.

## Hosting details

The webhook is deployed as a Kubernetes API service so access can be restricted
via Kubernetes RBAC controls.  This prevents arbitrary people with access to the
webhook from completing ACME challenge validations and obtaining certificates.

# Running the test suite

Before deploying  any modifications, you **must** run the DNS01 provider
conformance testing suite as configured in [main_test.go](https://github.com/rb12345/cert-manager-webhook-hydra/blob/master/main_test.go).

You can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.ox.ac.uk. TEST_DNS_SERVER=pigeon.dns.ox.ac.uk:53 make test
```

after first creating a Kubernetes Secret manifest file at
`testdata/hydra-dns01-solver/hydratoken.yaml` with credentials for the Hydra
sandpit environment.

# Using the new provider

First, create a Hydra token with the following mask:

* `hostname: %.subdomain.ox.ac.uk.`
* `ip: [blank]`
* `content: [blank]`
* `target: [blank]`

where `subdomain.ox.ac.uk` is the domain for which challenges should be issued.

Next, install the webhook using Helm:

    cd deploy/cert-manager-webhook-hydra/
    helm install cert-manager-webhook-hydra .

Once the webhook is installed, create a new Basic auth secret containing the
Hydra token:

    apiVersion: v1
    kind: Secret
    metadata:
      name: hydra-api-token
    type: kubernetes.io/basic-auth
    stringData:
      username: "$TOKEN_USERNAME"
      password: "$TOKEN_PASSWORD"

then use the following `Issuer` configuration to make use of the webhook:

    apiVersion: cert-manager.io/v1
    kind: Issuer
    metadata:
      name: example-issuer
    spec:
      acme:
        email: first.last@it.ox.ac.uk
        server: https://acme-staging-v02.api.letsencrypt.org/directory
        privateKeySecretRef:
          name: example-issuer-account-key
        solvers:
        - dns01:
            webhook:
              groupName: "acme.ox.ac.uk"
              solverName: "hydra-dns01-solver"
              config:
                # For the sandpit, use:
                # hydraBasePath: "https://www.networks.it.ox.ac.uk/api/sandpit/ipam/"
                hydraBasePath: "https://www.networks.it.ox.ac.uk/api/ipam/"
                hydraTokenSecretRef:
                  name: "hydra-api-token"

You should then be able to request a new certificate via:

    apiVersion: cert-manager.io/v1
    kind: Certificate
    metadata:
      name: testname-cert
    spec:
      dnsNames:
        - testname.subdomain.ox.ac.uk
      secretName: testname-cert-tls
      issuerRef:
        name: example-issuer

and watch your new certificate appear within 5-10 minutes after DNS has been
updated.
