---
name: Bug Report
about: Use THIS AND ONLY THIS template for reporting bugs or issues
title: "Title"
labels: "kind/bug"
---
# Bug Report

<!--
You MUST use this template while reporting a bug.
Provide as much info as possible so that we may understand the issue and assist.
Enclose all log/console output in triple-backticks.
Issues that do not fill in this template or follow its guidelines will be closed.
Thanks!

If the matter is security related, please disclose it privately via https://kubernetes.io/security/
-->

**What happened**:

**What you expected to happen**:

**How to reproduce this bug (as minimally and precisely as possible)**:

**Anything else relevant for this bug report?**:

**Resources and logs to submit**:

Copy all relevant COSI resources here in yaml format:

```yaml
# BucketClass
# BucketAccessClass
# BucketClaim
# Bucket
# BucketAccess
```

```
# Copy COSI controller pod logs here
```

```
# Copy COSI sidecar logs here for the relevant driver
```

<!--
To get resources in yaml format, use `kubectl -n <namespace> get -o yaml`
To get logs, use `kubectl -n <namespace> logs <pod name>`

When pasting logs, always surround them with triple-backticks.
Read the GitHub documentation if you need help:
https://help.github.com/en/articles/creating-and-highlighting-code-blocks
-->

**Environment**:

- Kubernetes version (use `kubectl version`), please list client and server:
- Sidecar version (provide the release tag or commit hash):
- Provisoner name and version (provide the release tag or commit hash):
- Cloud provider or hardware configuration:
- OS (e.g: `cat /etc/os-release`):
- Kernel (e.g. `uname -a`):
- Install tools:
- Network plugin and version (if this is a network-related bug):
- Others:
