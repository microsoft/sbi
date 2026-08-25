# Daily Recommended Images by Language

_Generated: 2026-08-25T02:36:40Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.317 | - | 0 | 0 | 10 | 881.0 MB | 2026-08-20 | `sha256:f8007eff0818` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:f8007eff0818c839a7fcace720f69e3d0437230983ff42311dbe637784186305` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.400 | - | 0 | 5 | 5 | 931.0 MB | 2026-08-20 | `sha256:148df6ae5a1a` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:148df6ae5a1a242c4d737aecea047eabd7764c05f9d7016433ce64d6bb6fe00c` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 4 | 28 | 183 | 193.0 MB | 2026-08-25 | `sha256:9d94ecf60a21` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:9d94ecf60a21c6e7a784cf0761fbd4a8391646617a0ff2f39621443d580cc2c3` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 4 | 28 | 183 | 218.0 MB | 2026-08-25 | `sha256:787c228ea854` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:787c228ea85457bec43c8b084e6ac360b26ea43b5c2fcbe861f721f2e8670dd3` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 13 | 78 | 392 | 850.0 MB | 2026-08-25 | `sha256:bb32ba3ba3ea` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:bb32ba3ba3ea36e38572d9d8db76fa15f7cbf722f3f886e06bca6d528bd4fba8` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 16 | 193.0 MB | 2026-08-17 | `sha256:f0de11ccd19f` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f0de11ccd19f2dd4b119cb0ec70985737b6e2f3fb96f791b2eadaead624bb428` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 16 | 198.0 MB | 2026-08-17 | `sha256:51ff44dbc1c0` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:51ff44dbc1c0898e7afb8d57e6b5759cc22735114cc2747191f2aca92e37b075` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 16 | 203.0 MB | 2026-08-17 | `sha256:a365ce6a50b0` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a365ce6a50b09176855d085c69da3fc1204a48432e36087e9a208f6e5860e235` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 16 | 217.0 MB | 2026-08-17 | `sha256:d72743fde6e2` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d72743fde6e26a853eaa5dd58a6e30cfbedf17186fe2aed6ff283b82efee9249` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 16 | 223.0 MB | 2026-08-17 | `sha256:9ee47ed695f4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:9ee47ed695f4839b0ae8fc28800982c2277fb6eeabd8ff7d7b0d1846f61fb25c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 16 | 230.0 MB | 2026-08-17 | `sha256:a4556ed033fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a4556ed033fa96f984bb7a8d348851cb2d36b1281dd2420070045f664fbb5f94` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 33 | 837.0 MB | 2026-08-17 | `sha256:278ce7b39d84` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:278ce7b39d8418e5e204eb7ef2882947af6baf8ff81396bbcdaacf5819ad2f79` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 28 | 887.0 MB | 2026-08-17 | `sha256:e1ffd2a92ae8` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1ffd2a92ae84c1291bc1b6887501f8af98e6331e7af6d4c8d37168c5e87a64c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 44 | 836.0 MB | 2026-08-17 | `sha256:2ae6f287fa86` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2ae6f287fa860c15f121474cf864b86765beb87507bbc3f48661a4f6f1ffc2b5` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.7 | - | 0 | 0 | 0 | 842.0 MB | 2026-08-24 | `sha256:9fdbea751306` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:9fdbea751306c03fe0250f6cc0334fd321e5a847484d0cf32357ec1ea27fdcbc` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 0 | 4 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-08-24 | `sha256:386144d0952e` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:386144d0952e9b3e1cde8eac65077d4b865fffa6ac1b99c8f08fe3eeeffe9114` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-08-24 | `sha256:9941d1670738` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:9941d16707383c392eba14073c871dc275934fde7d17877b957d045e75b4c051` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-08-24 | `sha256:67cd622af5b8` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:67cd622af5b8a17b82a9ac5cca2191eb640d1b66984fc98a09d16d82c2b0972e` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-08-24 | `sha256:761d0db32563` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:761d0db3256388108dba6d61eb8cc2ee536cacc8607d6d94d95a7b780ec306bd` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-24 | `sha256:693592c22910` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:693592c2291097a8c50a5067da53a3a6e0e59270eb1b214052edd0e62ac63b57` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 522.0 MB | 2026-08-24 | `sha256:3a8ea5809f4d` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:3a8ea5809f4d1e6074085344dfbd382ad3cac641acda5f27015b45513ce4566f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 75 | 430.0 MB | 2026-08-24 | `sha256:34344055b3d5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:34344055b3d573c32f319314fb6e03265278692e296b9b8f74c91d8ca8f3beaf` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 75 | 457.0 MB | 2026-08-24 | `sha256:7943428e84be` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:7943428e84bef36f75f2f4880eb168653017256b7bbbaaaf9f19168bbc37d6e9` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 75 | 502.0 MB | 2026-08-24 | `sha256:c2da1d3b87f0` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:c2da1d3b87f04278cc03299913a1dd13a134fb1900768313c92cac86e9fd2548` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 0 | 0 | 197.0 MB | 2026-08-20 | `sha256:0ce7fd8a4147` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:0ce7fd8a4147cd5f5bd96adc781ab661905318c8ae78a19e278fae0ad9caab18` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 9 | 31 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 7 | 19 | 157.0 MB | 2026-08-20 | `sha256:4a441147583f` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:4a441147583fcfbf3bbf3bb4de73645847a1132177f59d4899912cf868eadb94` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 7 | 19 | 157.0 MB | 2026-08-20 | `sha256:816f0ecfafda` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:816f0ecfafda49fddbdd473a1e33ca8690cd1528b3d0f0723f8336d28eedff8e` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 11 | 37 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 11 | 37 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 42 | 112 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 42 | 112 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 39 | 174 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 48 | 175 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 83.5 MB | 2026-08-20 | `sha256:db54ecb65d38` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:db54ecb65d383bc96039c11c7c10e2ecbbf76a4bc9a5cc2ce776be0b199be14d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 83.5 MB | 2026-08-20 | `sha256:6220d0e89206` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:6220d0e89206a47c11d7928f7bee15ca9c67d4626c3262fc11fe5b8b0cf04e91` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-08-20 | `sha256:df0cc451d3d3` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:df0cc451d3d37ad9d670e3bbb3ffb09ae4d4d1fb5665c196452873dbffb557b9` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.0 MB | 2026-08-20 | `sha256:76325dbdd581` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:76325dbdd581f039ac2edffca3ff40aef88951da7aa24cf0c07b40992e6d6bd2` |
