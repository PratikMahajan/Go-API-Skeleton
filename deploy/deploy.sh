#!/usr/bin/env bash

set -e

# Turn colors in this script off by setting the NO_COLOR variable in your
# environment to any value:

NO_COLOR=${NO_COLOR:-""}
if [ -z "$NO_COLOR" ]; then
  header=$'\e[1;33m'
  reset=$'\e[0m'
else
  header=''
  reset=''
fi

while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -n|--namespace)
    namespace="$2"
    shift # past argument
    shift # past value
    ;;
    -t|--type)
    install="$2"
    shift
    shift
    ;;
esac
done


function header_text {
  echo "$header$*$reset"
}


echo "Using namespace/Project $namespace"
oc project $namespace

if [ "$install" == "apply" ]; then
    header_text "Starting deployment on OpenShift!"
    for filename in deploy/templates/*.yaml; do
        header_text "Applying configuration $filename on $namespace"
        kubectl -n $namespace apply -f $filename
    done
fi

if [ "$install" == "delete" ]; then
    header_text "Deleting deployment from OpenShift!"
    for filename in deploy/templates/*.yaml; do
        header_text "Deleting configuration $filename on $namespace"
        kubectl -n $namespace delete -f $filename
    done
fi

header_text "All Configurations applied on $namespace"
