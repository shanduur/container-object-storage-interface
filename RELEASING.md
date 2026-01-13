# Releasing

This repository ships multiple Go modules:

- `sigs.k8s.io/container-object-storage-interface` (root)
- `sigs.k8s.io/container-object-storage-interface/client` (client/)
- `sigs.k8s.io/container-object-storage-interface/proto` (proto/)

Each module requires its own semantic version tag in Git.

## Release candidate (`-rc.X`) flow

1. Start from a clean `main` (or the targeted release) commit.
2. Choose the next version (for example `v0.2.0-rc.1`).
3. Create annotated tags for the modules that changed (usually all three):
   ```bash
   git tag -a v0.2.0-rc.1 -m "root module v0.2.0-rc.1"
   git tag -a client/v0.2.0-rc.1 -m "client module v0.2.0-rc.1"
   git tag -a proto/v0.2.0-rc.1 -m "proto module v0.2.0-rc.1"
   git push origin v0.2.0-rc.1 client/v0.2.0-rc.1 proto/v0.2.0-rc.1
   ```
4. Use the RC tags for downstream testing and validation.

## Promotion to a final release

1. When the RC is validated, reuse the same commit. Do **not** cherry-pick changes after the RC; cut a new RC if needed.
2. Create annotated tags without the `-rc.X` suffix, matching each module:
   ```bash
   git tag -a v0.2.0 -m "root module v0.2.0"
   git tag -a client/v0.2.0 -m "client module v0.2.0"
   git tag -a proto/v0.2.0 -m "proto module v0.2.0"
   git push origin v0.2.0 client/v0.2.0 proto/v0.2.0
   ```
3. Publish release notes (if applicable) referencing the final tags.

## Tagging notes

- Tag names must follow Go module tagging rules: submodules use `path-prefix/vX.Y.Z`.
- Tag only the modules that changed; keep versions aligned across modules when possible.
- Sign tags (`git tag -s`) if required by project policy.
