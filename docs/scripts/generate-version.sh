jq ".[1]";
SHA=$(git rev-parse HEAD)
VERSION="Built from: <a target="_blank" href=\"https:\/\/github.com\/kubernetes-sigs\/container-object-storage-interface\/tree\/${SHA}\"><code>${SHA}<\/code><\/a>"
TAG=$(git tag --contains "$SHA" | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$')
if [ -n "$TAG" ]; then
  VERSION="Version: <a target="_blank" href=\"https:\/\/github.com\/kubernetes-sigs\/container-object-storage-interface\/tree\/${TAG}\"><code>${TAG}<\/code><\/a> ${VERSION}"
fi
sed "s/VERSION-PLACEHOLDER/${VERSION}/" theme/index-template.hbs > theme/index.hbs
