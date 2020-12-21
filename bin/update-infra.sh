# /usr/bin/env bash

set -ex

BASEPATH=${BASEPATH:-"infra"}
WITHPUSH=${WITHPUSH:-"no"}

if [ -z "$IMAGE_NAME" ]; then
  echo "image name cannot be empty. please use IMAGE_NAME."
  exit 1
fi

cd $BASEPATH

# replace the image name in the deployment with the image we just built
replaced=$(cat kubernetes/deployment.yml | yq w - 'spec.template.spec.containers[0].image' $IMAGE_NAME)
echo "$replaced" > kubernetes/deployment.yml

if [ "$WITHPUSH" == "yes" ]; then
  echo "with push enabled, committing changes"
  git add kubernetes/deployment.yml
  git commit -m "chore(deploy): automated image update"

  git push
fi