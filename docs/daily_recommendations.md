# Daily Recommended Images by Language

_Generated: 2026-07-06T03:11:37Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **37 configured sources** across 8 groups (see [repositories.json](../config/repositories.json)):

**Base / minimal images (no runtime):**

- `mcr.microsoft.com/azurelinux/base/core:3.0`
- `mcr.microsoft.com/azurelinux/distroless/base:3.0`
- `mcr.microsoft.com/azurelinux/distroless/minimal:3.0`

**Azure Linux Python and Node.js images:**

- `azurelinux/base/python`
- `azurelinux/base/nodejs`

**Azure Linux distroless runtime images:**

- `azurelinux/distroless/python`
- `azurelinux/distroless/nodejs`

**.NET Azure Linux images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`

**.NET Ubuntu (Noble) images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-noble`
- `mcr.microsoft.com/dotnet/aspnet:9.0-noble`
- `mcr.microsoft.com/dotnet/aspnet:8.0-noble`
- `mcr.microsoft.com/dotnet/runtime:10.0-noble`
- `mcr.microsoft.com/dotnet/runtime:9.0-noble`
- `mcr.microsoft.com/dotnet/runtime:8.0-noble`
- `mcr.microsoft.com/dotnet/sdk:10.0-noble`
- `mcr.microsoft.com/dotnet/sdk:9.0-noble`
- `mcr.microsoft.com/dotnet/sdk:8.0-noble`

**.NET Debian images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

**OpenJDK images:**

- `mcr.microsoft.com/openjdk/jdk:25-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:25-distroless`
- `mcr.microsoft.com/openjdk/jdk:25-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:21-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:21-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:17-distroless`
- `mcr.microsoft.com/openjdk/jdk:17-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:11-distroless`

## Dotnet

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 107.0 MB | 2026-06-25 | `sha256:e51d733993da` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:e51d733993da9264bfeab967cf7f6a148f2bbb2ff58127c979deb2b282ea37b7` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 112.0 MB | 2026-06-25 | `sha256:fe26ec36e3b6` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:fe26ec36e3b6ab24c097e28a80112f7876fdbdf30f92a8a1b31cf4369e254a9f` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.28 | - | 0 | 0 | 0 | 126.0 MB | 2026-06-25 | `sha256:25f642a9fc39` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:25f642a9fc39ef55ce2740e695d66614cb69d6619fb55790b7d5f0dd6d83e21a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 132.0 MB | 2026-06-25 | `sha256:cf4f772d89c6` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:cf4f772d89c691e21e2d93a5d04328fa9fd025bb4a79b8fa069da8e78ab01a1c` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 139.0 MB | 2026-06-25 | `sha256:cf989ccd6e0f` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:cf989ccd6e0fe34e0efb3a6373ab297f168c2de0c1896bf4883f68b0cadafb15` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.301 | - | 0 | 38 | 80 | 927.0 MB | 2026-06-25 | `sha256:a996dc54d5ad` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:a996dc54d5ad919926fa00fce172fdd980c2312d2f7a800beade5cfc3a828e75` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.315 | - | 0 | 38 | 90 | 888.0 MB | 2026-06-25 | `sha256:bdf7cb101362` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:bdf7cb101362974a98f5db53a07b87733d4c9ebd75840c68d0ddb949d95d0f1f` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.28 | - | 3 | 8 | 158 | 193.0 MB | 2026-06-25 | `sha256:4cbc2530d585` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:4cbc2530d5855154199432020f232205aecd8be34be87bea866f3b235f853bcb` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.28 | - | 3 | 8 | 158 | 218.0 MB | 2026-06-25 | `sha256:e78fda31142e` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:e78fda31142e28746a6908e288e0d40346793f691ca99d8d150bcbe95c0ef035` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.422 | - | 9 | 40 | 358 | 850.0 MB | 2026-06-25 | `sha256:63ebabdcde24` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:63ebabdcde24cc8134304a6f50719cf1e22bb4e9f7148ac4ef967c7680187356` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.28 | - | 0 | 0 | 33 | 193.0 MB | 2026-06-25 | `sha256:1b809194e628` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:1b809194e6288a772512b6ed6d6f8780e476b3f52ee0591eefd7c3a88add7aaf` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.17 | - | 0 | 0 | 33 | 198.0 MB | 2026-06-25 | `sha256:0f7c41465ace` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:0f7c41465ace8ae345a18e167000bf7ae16e0ed87c398e1bc73cb24ac3c1296d` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.9 | - | 0 | 0 | 33 | 203.0 MB | 2026-06-25 | `sha256:6a40d375e9c8` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:6a40d375e9c8432fcf4adebae05d7e0a276e9a90dd01174df6709a090771bebc` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.28 | - | 0 | 0 | 33 | 217.0 MB | 2026-06-25 | `sha256:77526aac4254` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:77526aac4254c2b26764c2ab09700f6bdbcad2e2cb748e0d57bc01ebdced6515` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.17 | - | 0 | 0 | 33 | 223.0 MB | 2026-06-25 | `sha256:165ce688e1fc` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:165ce688e1fc73f2a624dfc8142531a7c594e155a7287b7ea12a43dc49f6601a` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.9 | - | 0 | 0 | 33 | 230.0 MB | 2026-06-25 | `sha256:7644f992230d` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:7644f992230d35cf230017189d4038c0ae0f7388b13f4f7ae1900a155bafb597` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.301 | - | 0 | 0 | 63 | 883.0 MB | 2026-06-25 | `sha256:ea8bde36c11b` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ea8bde36c11b6e7eec2656d0e59101d4462f6bd630730f2c8201ed0572b295d5` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.315 | - | 0 | 0 | 73 | 844.0 MB | 2026-06-25 | `sha256:b128a1e799e9` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:b128a1e799e98bceb26cc913b40d56da6e13cc7b0f665a5a81a0a4ded3f1f64d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.422 | - | 0 | 10 | 83 | 844.0 MB | 2026-06-25 | `sha256:08f003ba31fc` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:08f003ba31fca1aeed32d5a076f62dc3b592faa0a811b072aa245321dd09f0ab` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 1 | 492.0 MB | 2026-06-29 | `sha256:880d5ca6fe3e` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:880d5ca6fe3e120c86be3241fc4de603f9930f2560606f5423f699eb72bbcb66` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 1 | 537.0 MB | 2026-06-29 | `sha256:0a1c77cd005b` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0a1c77cd005bd101b716e5ac6bb66144b2c64963460b74d56e961294fdcc19c2` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 2 | 8 | 22 | 323.0 MB | 2026-06-29 | `sha256:a671bb00f995` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:a671bb00f995b0bc74ee0b5d44e50b6e6e78cd739748b9cdbff3f23ecda271dd` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 2 | 8 | 22 | 326.0 MB | 2026-06-29 | `sha256:e5ca11b6b85b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:e5ca11b6b85b94568462c1ba7b801ea93db8245610a4ae359931c23dd7ec91b5` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 2 | 8 | 22 | 354.0 MB | 2026-06-29 | `sha256:598b7d34e94f` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:598b7d34e94f54b6daaf44aaabd6249ff375c00c664cc64fa77c627d7f1c6b93` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 2 | 8 | 22 | 399.0 MB | 2026-06-29 | `sha256:5ebed5990a48` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:5ebed5990a48fa9f6f09a8f014854b7e7ea52c7fc6950eb3b795f81bbb375180` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 0 | 91 | 448.0 MB | 2026-06-29 | `sha256:03eae6fb6d68` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:03eae6fb6d68fba9c7591e0f77556fcd474dc046fdc9bbef71b9b2dc4662bf87` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 0 | 91 | 475.0 MB | 2026-06-29 | `sha256:6e4d851f85e7` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:6e4d851f85e74a57352868bf8509ad0df86bf3a763a306984179ca986de0b88c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 0 | 91 | 519.0 MB | 2026-06-29 | `sha256:b9b2f8b852bd` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:b9b2f8b852bdb0bf0ded73f21cc6c780fbe3a42f771bbfa4317d1391b8cb76ea` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 8 | 19 | 83.7 MB | 2026-06-19 | `sha256:b9c3f3717641` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b9c3f3717641372bef57dc9cbdcc191087587154cb9b01c3dae5b0c480a77c6d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 8 | 19 | 83.7 MB | 2026-06-19 | `sha256:d188df941951` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:d188df941951658d215aa37c9e52d01b94821bd7e356f7bb8d7603db392a8b06` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 2 | 11 | 32 | 139.0 MB | 2026-06-19 | `sha256:84be0b597731` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:84be0b597731efe4134de278e1321e5bc0fa30eee24c3775b8ee92457be3603d` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.11 | - | 0 | 0 | 35 | 825.0 MB | 2026-07-01 | `sha256:0f4515bec95c` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:0f4515bec95c6f6f1c0008551b27457bf7ca20966f63e6bd700f0d810fe61c0d` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.4 | - | 0 | 0 | 35 | 857.0 MB | 2026-07-01 | `sha256:877a4c2aed04` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:877a4c2aed044ce1b36c043c1af0c939c3129a7a854453117e5e2b9d574d4b52` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | :24-nonroot | 0 | 14 | 32 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | :24 | 0 | 14 | 32 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 3 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | :24 | 2 | 11 | 32 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 20 | 91 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 2 | 34 | 81 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 2 | 34 | 81 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 3 | 42 | 117 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | :20-nonroot | 3 | 46 | 98 | 106.0 MB | 2026-03-04 | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | :20 | 3 | 46 | 98 | 106.0 MB | 2026-03-04 | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-06-19 | `sha256:83c9e52e0a0e` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:83c9e52e0a0ef97d9d87b8b81da6119f748dbeca4641ae2cb8b11552e2c8f35d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 2 | 8 | 22 | 34.1 MB | 2026-06-19 | `sha256:f8f5a9bb739a` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:f8f5a9bb739ad1ec347853144c9ed4ca2260e587082277bc6066fcd5cc9973e8` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 2 | 11 | 32 | 75.3 MB | 2026-06-19 | `sha256:1c56f09437df` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:1c56f09437dfc2910faad39abaaed336265d246cf183e2adb362d4cb3b881ab6` |
