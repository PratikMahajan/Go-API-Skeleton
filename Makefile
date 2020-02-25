.PHONY: container container-run

CONTAINER_BIN ?= podman

APP ?= go-app-twitter
TAG_VERSION ?= staging-latest
CONTAINER_TAG ?= pratikmahajan/${APP}:${TAG_VERSION}

NAMESPACE ?= dmprj

container-build:
	${CONTAINER_BIN} build -t ${CONTAINER_TAG} .

container-push:
	${CONTAINER_BIN} push ${CONTAINER_TAG}

container: container-build container-push

container-debug:
	${CONTAINER_BIN} run -it --rm --net host --entrypoint /bin/sh ${CONTAINER_TAG}

container-run:
	${CONTAINER_BIN} run -it --rm -e APP_ACCESSTOKEN=${APP_ACCESSTOKEN} -e APP_ACCESSTOKENSECRET=${APP_ACCESSTOKENSECRET} \
	 -e APP_APIKEY=${APP_APIKEY} -e APP_APISECRETKEY=${APP_APISECRETKEY} --net host ${CONTAINER_TAG}

deploy-app:
	./deploy/deploy.sh -n ${NAMESPACE} -t "apply"

delete-app:
	./deploy/deploy.sh -n ${NAMESPACE} -t "delete"