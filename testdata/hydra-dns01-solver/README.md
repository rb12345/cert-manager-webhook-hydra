# Solver testdata directory

You will need to create a file `hydratoken.yaml` with the following contents:

    apiVersion: v1
    kind: Secret
    metadata:
      name: hydra-api-token
    type: kubernetes.io/basic-auth
    stringData:
      username: "$HYDRA_TOKEN_USERNAME"
      password: "$HYDRA_TOKEN_PASSWORD"

replacing `$HYDRA_TOKEN_USERNAME` and `$HYDRA_TOKEN_PASSWORD` with the actual
values for your token.
