#!/usr/bin/env bash
set -o errexit
set -o nounset

# with nounset, these will fail if necessary vars are missing
echo "GIT_TAG: ${GIT_TAG}"
echo "PULL_BASE_REF: ${PULL_BASE_REF}"
echo "PLATFORM: ${PLATFORM}"

# debug the rest of the script in case of image/CI build issues
set -o xtrace

REPO="gcr.io/k8s-staging-sig-storage"

CONTROLLER_IMAGE="${REPO}/objectstorage-controller"
SIDECAR_IMAGE="${REPO}/objectstorage-sidecar"

# args to 'make build'
export DOCKER="/buildx-entrypoint" # available in gcr.io/k8s-testimages/gcb-docker-gcloud image
export PLATFORM
export SIDECAR_TAG="${SIDECAR_IMAGE}:${GIT_TAG}"
export CONTROLLER_TAG="${CONTROLLER_IMAGE}:${GIT_TAG}"

ADDITIONAL_BUILD_ARGS="--push"
ADDITIONAL_CONTROLLER_TAGS=()
ADDITIONAL_SIDECAR_TAGS=()

# PULL_BASE_REF is 'main' for non-tagged commits on the main branch
if [[ "${PULL_BASE_REF}" == main ]]; then
  echo " ! ! ! this is a main branch build ! ! !"
  # 'main' tag follows the main branch head
  ADDITIONAL_CONTROLLER_TAGS+=("${CONTROLLER_IMAGE}:main")
  ADDITIONAL_SIDECAR_TAGS+=("${SIDECAR_IMAGE}:main")
  # 'latest' tag follows 'main' for easy use by developers
  ADDITIONAL_CONTROLLER_TAGS+=("${CONTROLLER_IMAGE}:latest")
  ADDITIONAL_SIDECAR_TAGS+=("${SIDECAR_IMAGE}:latest")
fi

# PULL_BASE_REF is 'release-*' for non-tagged commits on release branches
if [[ "${PULL_BASE_REF}" == release-* ]]; then
  echo " ! ! ! this is a ${PULL_BASE_REF} release branch build ! ! !"
  # 'release-*' tags that follow each release branch head
  ADDITIONAL_CONTROLLER_TAGS+=("${CONTROLLER_IMAGE}:${PULL_BASE_REF}")
  ADDITIONAL_SIDECAR_TAGS+=("${SIDECAR_IMAGE}:${PULL_BASE_REF}")
fi

# PULL_BASE_REF is 'controller/TAG' for a tagged controller release
if [[ "${PULL_BASE_REF}" == controller/* ]]; then
  echo " ! ! ! this is a tagged controller release ! ! !"
  TAG="${PULL_BASE_REF#controller/*}"
  ADDITIONAL_CONTROLLER_TAGS+=("${CONTROLLER_IMAGE}:${TAG}")
fi

# PULL_BASE_REF is 'sidecar/TAG' for a tagged sidecar release
if [[ "${PULL_BASE_REF}" == sidecar/* ]]; then
  echo " ! ! ! this is a tagged sidecar release ! ! !"
  TAG="${PULL_BASE_REF#sidecar/*}"
  ADDITIONAL_SIDECAR_TAGS+=("${SIDECAR_IMAGE}:${TAG}")
fi

# PULL_BASE_REF is 'v0.y.z*' for tagged alpha releases where controller and sidecar are released simultaneously
# hand wave over complex matching logic by just looking for 'v0.' prefix
if [[ "${PULL_BASE_REF}" == 'v0.'* ]]; then
  echo " ! ! ! this is a tagged controller + sidecar release ! ! !"
  TAG="${PULL_BASE_REF}"
  ADDITIONAL_CONTROLLER_TAGS+=("${CONTROLLER_IMAGE}:${TAG}")
  ADDITIONAL_SIDECAR_TAGS+=("${SIDECAR_IMAGE}:${TAG}")
fi

# else, PULL_BASE_REF is something that doesn't release image(s) to staging, like:
#  - a random branch name (e.g., feature-xyz)
#  - a version tag for a subdir with no image associated (e.g., client/v0.2.0, proto/v0.2.0)

# 'gcloud container images add-tag' within the cloudbuild infrastructure doesn't preserve the date
# of the underlying image when adding a new tag, resulting in tags dated Dec 31, 1969 (the epoch).
# To ensure the right date on all built image tags, do the build with '--tag' args for all tags.

BUILD_ARGS="${ADDITIONAL_BUILD_ARGS}"
for tag in "${ADDITIONAL_CONTROLLER_TAGS[@]}"; do
  BUILD_ARGS="${BUILD_ARGS} --tag=${tag}"
done
export BUILD_ARGS
make build.controller

BUILD_ARGS="${ADDITIONAL_BUILD_ARGS}"
for tag in "${ADDITIONAL_SIDECAR_TAGS[@]}"; do
  BUILD_ARGS="${BUILD_ARGS} --tag=${tag}"
done
export BUILD_ARGS
make build.sidecar
