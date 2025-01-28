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
export BUILD_ARGS="--push"
export PLATFORM
export SIDECAR_TAG="${SIDECAR_IMAGE}:${GIT_TAG}"
export CONTROLLER_TAG="${CONTROLLER_IMAGE}:${GIT_TAG}"

make build

# PULL_BASE_REF is 'main' for non-tagged commits on the main branch
if [[ "${PULL_BASE_REF}" == main ]]; then
  echo " ! ! ! this is a main branch build ! ! !"
  # 'main' tag follows the main branch head
  gcloud container images add-tag "${CONTROLLER_TAG}" "${CONTROLLER_IMAGE}:main"
  gcloud container images add-tag "${SIDECAR_TAG}" "${SIDECAR_IMAGE}:main"
  # 'latest' tag follows 'main' for easy use by developers
  gcloud container images add-tag "${CONTROLLER_TAG}" "${CONTROLLER_IMAGE}:latest"
  gcloud container images add-tag "${SIDECAR_TAG}" "${SIDECAR_IMAGE}:latest"
fi

# PULL_BASE_REF is 'release-*' for non-tagged commits on release branches
if [[ "${PULL_BASE_REF}" == release-* ]]; then
  echo " ! ! ! this is a ${PULL_BASE_REF} release branch build ! ! !"
  # 'release-*' tags that follow each release branch head
  gcloud container images add-tag "${CONTROLLER_TAG}" "${CONTROLLER_IMAGE}:${PULL_BASE_REF}"
  gcloud container images add-tag "${SIDECAR_TAG}" "${SIDECAR_IMAGE}:${PULL_BASE_REF}"
fi

# PULL_BASE_REF is 'controller/TAG' for a tagged controller release
if [[ "${PULL_BASE_REF}" == controller/* ]]; then
  echo " ! ! ! this is a tagged controller release ! ! !"
  TAG="${PULL_BASE_REF#controller/*}"
  gcloud container images add-tag "${CONTROLLER_TAG}" "${CONTROLLER_IMAGE}:${TAG}"
fi

# PULL_BASE_REF is 'sidecar/TAG' for a tagged sidecar release
if [[ "${PULL_BASE_REF}" == sidecar/* ]]; then
  echo " ! ! ! this is a tagged sidecar release ! ! !"
  TAG="${PULL_BASE_REF#sidecar/*}"
  gcloud container images add-tag "${SIDECAR_TAG}" "${SIDECAR_IMAGE}:${TAG}"
fi

# PULL_BASE_REF is 'v0.y.z*' for tagged alpha releases where controller and sidecar are released simultaneously
# hand wave over complex matching logic by just looking for 'v0.' prefix
if [[ "${PULL_BASE_REF}" == 'v0.'* ]]; then
  echo " ! ! ! this is a tagged controller + sidecar release ! ! !"
  TAG="${PULL_BASE_REF}"
  gcloud container images add-tag "${CONTROLLER_TAG}" "${CONTROLLER_IMAGE}:${TAG}"
  gcloud container images add-tag "${SIDECAR_TAG}" "${SIDECAR_IMAGE}:${TAG}"
fi

# else, PULL_BASE_REF is something that doesn't release image(s) to staging, like:
#  - a random branch name (e.g., feature-xyz)
#  - a version tag for a subdir with no image associated (e.g., client/v0.2.0, proto/v0.2.0)
