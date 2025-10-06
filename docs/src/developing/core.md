# Developing "core" COSI

With “core” COSI we refer to the common set of API and controllers that are required to run any COSI driver.

Before your first contribution, you should follow [the Kubernetes Contributor Guide](https://www.kubernetes.dev/docs/guide/#contributor-guide).

To further understand the COSI architecture, please refer to [KEP-1979: Object Storage
Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1979-object-storage-support).

Before contributing a Pull Request, ensure a [GitHub issue](https://github.com/kubernetes-sigs/container-object-storage-interface/issues) exists corresponding to the change.

## Local code development

For new contributors, use `make help`, and use **Core** targets as needed.
These targets ensure changes build successfully, pass basic checks, and are ready for end-to-end tests run in COSI's automated CI.

Other more advanced targets are available and also described in `make help` output.

Some specific workflows are documented below.

### Building and deploying COSI controller changes locally

```sh
export CONTROLLER_TAG="$MY_REPO"/cosi-controller:latest # replace MY_REPO with desired dev repo
make -j prebuild
make build.controller
docker push "$CONTROLLER_TAG"
make deploy
```
