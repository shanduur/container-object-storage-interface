![status](https://img.shields.io/badge/status-pre--alpha-lightblue)
![apiVersion](https://img.shields.io/badge/apiVersion-v1alpha2-lightblue)
[![docs](https://img.shields.io/badge/docs-latest-lightblue)](https://container-object-storage-interface.sigs.k8s.io/)

# Container Object Storage Interface

This repository hosts the Container Object Storage Interface (COSI) project.

> [!IMPORTANT]
> This `main` branch contains pre-alpha code and APIs for COSI `v1alpha2`.<br>
> For `v1alpha1` APIs, code, or development, use branch `release-0.2`

## Documentation

To deploy, run `kubectl apply -k .`

Documentation can be found under: https://container-object-storage-interface.sigs.k8s.io/

## References

- [Weekly meetings](https://www.kubernetes.dev/resources/calendar/): Thursdays from 13:30 to 14:00 US Eastern Time
- [Roadmap](https://github.com/orgs/kubernetes-sigs/projects/63/)

## Community, discussion, contribution, and support

You can reach the maintainers of this project at:

- [#sig-storage-cosi](https://kubernetes.slack.com/messages/sig-storage-cosi) Slack channel **(preferred)**
- [GitHub Issues](https://github.com/kubernetes-sigs/container-object-storage-interface/issues)
- [container-object-storage-interface](https://groups.google.com/g/container-object-storage-interface-wg) mailing list

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

## Developer Guide

All API definitions and behavior must follow the [`v1alpha2` KEP PR](https://github.com/kubernetes/enhancements/pull/4599).
Minor deviation from the KEP is acceptable in order to fix bugs.

`v1alpha2` is currently pre-release.
Changes may break compatibility up until `v1alpha2` is released with a semver tag.
After the first `v1alpha2` semver release (e.g., 0.3.0), all changes must be backwards compatible.

Before making a COSI contribution, please read and follow the
[core developer guide](https://container-object-storage-interface.sigs.k8s.io/developing/core.html).
