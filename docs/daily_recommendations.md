# Daily Recommended Images by Language

_Generated: 2026-08-19T02:35:47Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.317 | - | 0 | 0 | 10 | 881.0 MB | 2026-08-13 | `sha256:c9c558d6730a` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c9c558d6730a709144f578823ac0a3934eb7715bc8e18c5c1e5072bf42b3e0de` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.400 | - | 0 | 5 | 5 | 931.0 MB | 2026-08-13 | `sha256:d1348cda1563` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:d1348cda1563f5dc1a1d419023a2067389a3dbe7a1e0686913d9839f85053f00` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 5 | 19 | 179 | 193.0 MB | 2026-08-10 | `sha256:a27922704a14` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a27922704a14d6d1ce103697aa3efa885902db33651b41a84ad1594fdf5be66e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 5 | 19 | 179 | 218.0 MB | 2026-08-10 | `sha256:b0beb9cc1dee` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b0beb9cc1dee1c1b0749796110d4734292071b814207ad0d4f40611f7db04f7b` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 17 | 66 | 385 | 850.0 MB | 2026-08-10 | `sha256:306301580fca` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:306301580fcaa5b445180e759db59309979002d1000669cb4cf58a567d0014bc` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 15 | 193.0 MB | 2026-08-17 | `sha256:f0de11ccd19f` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f0de11ccd19f2dd4b119cb0ec70985737b6e2f3fb96f791b2eadaead624bb428` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 15 | 198.0 MB | 2026-08-17 | `sha256:51ff44dbc1c0` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:51ff44dbc1c0898e7afb8d57e6b5759cc22735114cc2747191f2aca92e37b075` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 15 | 203.0 MB | 2026-08-17 | `sha256:a365ce6a50b0` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a365ce6a50b09176855d085c69da3fc1204a48432e36087e9a208f6e5860e235` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 15 | 217.0 MB | 2026-08-17 | `sha256:d72743fde6e2` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d72743fde6e26a853eaa5dd58a6e30cfbedf17186fe2aed6ff283b82efee9249` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 15 | 223.0 MB | 2026-08-17 | `sha256:9ee47ed695f4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:9ee47ed695f4839b0ae8fc28800982c2277fb6eeabd8ff7d7b0d1846f61fb25c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 15 | 230.0 MB | 2026-08-17 | `sha256:a4556ed033fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a4556ed033fa96f984bb7a8d348851cb2d36b1281dd2420070045f664fbb5f94` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 29 | 837.0 MB | 2026-08-17 | `sha256:278ce7b39d84` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:278ce7b39d8418e5e204eb7ef2882947af6baf8ff81396bbcdaacf5819ad2f79` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 24 | 887.0 MB | 2026-08-17 | `sha256:e1ffd2a92ae8` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1ffd2a92ae84c1291bc1b6887501f8af98e6331e7af6d4c8d37168c5e87a64c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 40 | 836.0 MB | 2026-08-17 | `sha256:2ae6f287fa86` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2ae6f287fa860c15f121474cf864b86765beb87507bbc3f48661a4f6f1ffc2b5` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.13 | - | 0 | 0 | 0 | 809.0 MB | 2026-08-17 | `sha256:58269007190d` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:58269007190dca4778302320b520f44b73395d469b89b846e107d252d44bd671` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.6 | - | 0 | 0 | 0 | 842.0 MB | 2026-08-17 | `sha256:77e1ecb8e71d` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:77e1ecb8e71d4bf2bcd36ead98693ebae5663ab9fe0d32216889e8da4a2334be` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-18 | `sha256:067fad82d3a3` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:067fad82d3a35827531e9e40deea1ee6daace7738c9fb60f30ed6fb2b3211819` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4 | - | 0 | 0 | 0 | 522.0 MB | 2026-08-18 | `sha256:e148b26e65fe` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e148b26e65fe4fd0f0d848df23b9edbdb0989ed8bc6cf95e7e6a3ff5c19d63d1` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32 | - | 0 | 0 | 1 | 323.0 MB | 2026-08-18 | `sha256:8717ad06b1df` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:8717ad06b1df4cc73277d81652df5466033bd72f6a4ff8c7da6d37488465e58c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20 | - | 0 | 0 | 1 | 326.0 MB | 2026-08-18 | `sha256:3768707c7573` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:3768707c7573b05b59542e999bfae019876d1fdaa1b98d436d1d2634996b4e59` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12 | - | 0 | 0 | 1 | 354.0 MB | 2026-08-18 | `sha256:34042d53df9e` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:34042d53df9e1f06b7aba89db5e7626abfa88e74cdd599fd5ad102aef3f918de` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4 | - | 0 | 0 | 1 | 399.0 MB | 2026-08-18 | `sha256:2096b49e1d19` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:2096b49e1d19bd827acac26add73e9c67f09b65cb09543593ed5b072461e42fd` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20 | - | 0 | 0 | 75 | 429.0 MB | 2026-08-18 | `sha256:aafb0051121e` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:aafb0051121e8df5c49e9878e78a596b0b2e6f562996d4f08d09a1496e3e5e31` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12 | - | 0 | 0 | 75 | 457.0 MB | 2026-08-18 | `sha256:9cfbddd99f7c` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:9cfbddd99f7cf6fd2efc67d6e05c7f90e9f9561eb5f6046e6d52cfab7b6af124` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4 | - | 0 | 0 | 75 | 501.0 MB | 2026-08-18 | `sha256:092ed93dfee6` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:092ed93dfee667454c43bbbe8c44ff413de8068e15af7d14d9ce276d0f867149` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 0 | 0 | 197.0 MB | 2026-08-11 | `sha256:e4578b763a16` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:e4578b763a1698192f53e8de82175f0fa315946acac6d24e49ea064b1387215b` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 8 | 24 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 6 | 19 | 157.0 MB | 2026-08-11 | `sha256:58202c56c1d2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:58202c56c1d28e3ad3962143c9d2ec55c190b42cc19d1918d5a57b205fe7af6f` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 6 | 19 | 157.0 MB | 2026-08-11 | `sha256:bb5ac511e1f9` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:bb5ac511e1f92dcc192a3414ccd6941b1044111db7d7ee6ee4b2bddf76d90257` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 38 | 165 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 44 | 166 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.8 MB | 2026-08-11 | `sha256:6f12a7ae467f` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:6f12a7ae467f01f7f4f36893fb484813bf7a0cce555da9b20cc78938b6b27164` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.8 MB | 2026-08-11 | `sha256:8f0a5a7be5bb` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:8f0a5a7be5bb57f6605ffde72e5aff701d0c75d34a6797da3bfcf64140d63b8d` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-08-11 | `sha256:722b6224c23b` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:722b6224c23b3f21f5268e2073f80c0f396bc626e3193b6dbf66e40d89478f03` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.5 MB | 2026-08-11 | `sha256:8bb51342bd5e` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:8bb51342bd5eba915990ab608f91060d502bb7891a2d3d909e0419b932533029` |
