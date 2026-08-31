# Daily Recommended Images by Language

_Generated: 2026-08-31T02:15:51Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 108.0 MB | 2026-08-13 | `sha256:56f96c0dfa21` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:56f96c0dfa21cf170ff402ff75de39efe89c4fee9b92c0a8fb07ea751b621d45` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 112.0 MB | 2026-08-13 | `sha256:3d7b8851fa22` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:3d7b8851fa229f31031fb9c6787b5d38263ea9a6bf88f97bf17f32f446c9b9b0` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.30 | - | 0 | 0 | 0 | 126.0 MB | 2026-08-13 | `sha256:8957983e8929` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:8957983e8929735a20b978ea82bb3074d79c2a1410a10a70175b6c4d9fdd003a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 132.0 MB | 2026-08-13 | `sha256:c412debe88b0` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:c412debe88b0360c3fb477c0f49894295176f924bfcf869c0bce4aa085e501e2` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 139.0 MB | 2026-08-13 | `sha256:c2fbd1565145` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:c2fbd15651452871430ce3fbd0083c128b376dbac786797be1c7dd4091d20a23` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.317 | - | 0 | 0 | 10 | 881.0 MB | 2026-08-26 | `sha256:004142e58aeb` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:004142e58aeb2f5b1ddd5b551dd1e887e632597767e269dc699dcc094de4af0a` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.400 | - | 0 | 5 | 5 | 931.0 MB | 2026-08-26 | `sha256:0a694e0479f1` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:0a694e0479f1a604bc57dcdf95181a861f0c1b1ecd5b4ee8c1c2662b30c82150` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 4 | 26 | 197 | 193.0 MB | 2026-08-25 | `sha256:9d94ecf60a21` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:9d94ecf60a21c6e7a784cf0761fbd4a8391646617a0ff2f39621443d580cc2c3` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 4 | 26 | 197 | 218.0 MB | 2026-08-25 | `sha256:787c228ea854` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:787c228ea85457bec43c8b084e6ac360b26ea43b5c2fcbe861f721f2e8670dd3` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 13 | 77 | 407 | 850.0 MB | 2026-08-25 | `sha256:bb32ba3ba3ea` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:bb32ba3ba3ea36e38572d9d8db76fa15f7cbf722f3f886e06bca6d528bd4fba8` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 39 | 193.0 MB | 2026-08-17 | `sha256:f0de11ccd19f` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f0de11ccd19f2dd4b119cb0ec70985737b6e2f3fb96f791b2eadaead624bb428` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 39 | 198.0 MB | 2026-08-17 | `sha256:51ff44dbc1c0` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:51ff44dbc1c0898e7afb8d57e6b5759cc22735114cc2747191f2aca92e37b075` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 39 | 203.0 MB | 2026-08-17 | `sha256:a365ce6a50b0` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a365ce6a50b09176855d085c69da3fc1204a48432e36087e9a208f6e5860e235` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 39 | 217.0 MB | 2026-08-17 | `sha256:d72743fde6e2` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d72743fde6e26a853eaa5dd58a6e30cfbedf17186fe2aed6ff283b82efee9249` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 39 | 223.0 MB | 2026-08-17 | `sha256:9ee47ed695f4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:9ee47ed695f4839b0ae8fc28800982c2277fb6eeabd8ff7d7b0d1846f61fb25c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 39 | 230.0 MB | 2026-08-17 | `sha256:a4556ed033fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a4556ed033fa96f984bb7a8d348851cb2d36b1281dd2420070045f664fbb5f94` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 86 | 837.0 MB | 2026-08-17 | `sha256:278ce7b39d84` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:278ce7b39d8418e5e204eb7ef2882947af6baf8ff81396bbcdaacf5819ad2f79` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 81 | 887.0 MB | 2026-08-17 | `sha256:e1ffd2a92ae8` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1ffd2a92ae84c1291bc1b6887501f8af98e6331e7af6d4c8d37168c5e87a64c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 97 | 836.0 MB | 2026-08-17 | `sha256:2ae6f287fa86` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2ae6f287fa860c15f121474cf864b86765beb87507bbc3f48661a4f6f1ffc2b5` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.7 | - | 0 | 0 | 0 | 842.0 MB | 2026-08-26 | `sha256:8f638b09830f` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:8f638b09830f92f4005c56756cbafdd56d6460f41d2c1fd12b9bc02a5c85434c` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 6 | 14 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-08-28 | `sha256:70f9f19bee08` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:70f9f19bee082edf49b74c467abc08594f94acca77d264fc3f7da7c73fd0b0e4` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-08-28 | `sha256:1cfa74ef4921` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:1cfa74ef492105b0d2cd26adf1727a7b7ca4088c6b59fcdb820de8c6ab9fb80a` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-08-28 | `sha256:78a083b4c0c5` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:78a083b4c0c5ca8ff1b132de4f51daea17c191841dfeeb23606478b656771b74` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-08-28 | `sha256:386d7413d34e` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:386d7413d34ef9d1b1381eba9a03067022501136d662bc6a0970fc1208e8e1cf` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-28 | `sha256:50bff21357fe` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:50bff21357fecc8df7266541d43c59a3a62cf877c941ccab68950f6b80645c5e` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 522.0 MB | 2026-08-28 | `sha256:710408600d75` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:710408600d7504f43f95084dbb5e90deb11c95a236a80c45d169dd65b3d49d27` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 77 | 445.0 MB | 2026-08-28 | `sha256:0ec1ec6819e7` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:0ec1ec6819e70802ea87a154b3690ec0455e190a6f9cc86c15ffc9f55db6de41` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 77 | 473.0 MB | 2026-08-28 | `sha256:d4cb55376795` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:d4cb55376795c0facef7ff9e8cfd60dabdf792c7f8c7a26bb22c9a3e34d9b06e` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 77 | 517.0 MB | 2026-08-28 | `sha256:54387e1a7ac0` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:54387e1a7ac06e55f331973930bc6ebab4c0d6b3cf5e362d17190a11a1c12237` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 0 | 0 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 13 | 39 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 8 | 20 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 8 | 20 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 12 | 38 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 12 | 38 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 43 | 113 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 43 | 113 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 43 | 182 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 52 | 183 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 83.5 MB | 2026-08-25 | `sha256:d921452dba64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:d921452dba64944bf959f22450bb3740f5b2fff4a59faa64bd6b8eaf4c57b5b8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 83.5 MB | 2026-08-25 | `sha256:aff6cec76c03` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:aff6cec76c03b70afc2c062ebe7fba1d8a0c4ec6a2fa1c8cd1982db6a3037423` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-08-25 | `sha256:0b729c82c0dd` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:0b729c82c0ddc0769248e287d7414f0cc4e42ae4aa5b786aa99883c247e42bdb` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.0 MB | 2026-08-25 | `sha256:daa1142fc6b4` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:daa1142fc6b44e27c8112ec6b4c2d579ddb9bc6b3747504e666010a45a51faa4` |
