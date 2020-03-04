# Go API Skeleton

## Setup
### Twitter API keys 
Add your Twitter API keys in `go-twitter-secret.yaml` file in `deploy/templates`.

Check `go-twitter-secret.yaml.example` file in `deploy/templates` for reference

_All the secrets should be in `base64` format in the `go-twitter-secret.yaml`_ \
_P.S. your secrets file should always be gitignored and dockerignored to avoid leaking secrets_


## RUN

### Build the container image 
`make container-build`

### Debug the container image on local machine
`make container-debug` \
_The default entrypoint while running the docker image on local machine is `/bin/sh`_

### Run the container image on local machine
`make container-run` \
_Set the secrets as environment variables_

### Upload the container image to container-registry
`make container-push`

### Build and Upload the container image to repository 
`make container`

### Deploy the application 
`make deploy-app` \

### Delete the application 
`make delete-app` 

## Misc.
* _To change the default namespace and container repository see `Makefile`_