# Daily Recommended Images by Language

_Generated: 2026-08-24T02:40:59Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 4 | 28 | 183 | 193.0 MB | 2026-08-10 | `sha256:a27922704a14` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a27922704a14d6d1ce103697aa3efa885902db33651b41a84ad1594fdf5be66e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 4 | 28 | 183 | 218.0 MB | 2026-08-10 | `sha256:b0beb9cc1dee` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b0beb9cc1dee1c1b0749796110d4734292071b814207ad0d4f40611f7db04f7b` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 13 | 78 | 392 | 850.0 MB | 2026-08-10 | `sha256:306301580fca` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:306301580fcaa5b445180e759db59309979002d1000669cb4cf58a567d0014bc` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 15 | 193.0 MB | 2026-08-17 | `sha256:f0de11ccd19f` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f0de11ccd19f2dd4b119cb0ec70985737b6e2f3fb96f791b2eadaead624bb428` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 15 | 198.0 MB | 2026-08-17 | `sha256:51ff44dbc1c0` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:51ff44dbc1c0898e7afb8d57e6b5759cc22735114cc2747191f2aca92e37b075` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 15 | 203.0 MB | 2026-08-17 | `sha256:a365ce6a50b0` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a365ce6a50b09176855d085c69da3fc1204a48432e36087e9a208f6e5860e235` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 15 | 217.0 MB | 2026-08-17 | `sha256:d72743fde6e2` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d72743fde6e26a853eaa5dd58a6e30cfbedf17186fe2aed6ff283b82efee9249` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 15 | 223.0 MB | 2026-08-17 | `sha256:9ee47ed695f4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:9ee47ed695f4839b0ae8fc28800982c2277fb6eeabd8ff7d7b0d1846f61fb25c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 15 | 230.0 MB | 2026-08-17 | `sha256:a4556ed033fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a4556ed033fa96f984bb7a8d348851cb2d36b1281dd2420070045f664fbb5f94` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 32 | 837.0 MB | 2026-08-17 | `sha256:278ce7b39d84` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:278ce7b39d8418e5e204eb7ef2882947af6baf8ff81396bbcdaacf5819ad2f79` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 27 | 887.0 MB | 2026-08-17 | `sha256:e1ffd2a92ae8` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1ffd2a92ae84c1291bc1b6887501f8af98e6331e7af6d4c8d37168c5e87a64c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 43 | 836.0 MB | 2026-08-17 | `sha256:2ae6f287fa86` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2ae6f287fa860c15f121474cf864b86765beb87507bbc3f48661a4f6f1ffc2b5` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-08-21 | `sha256:6d5ddc4fca35` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6d5ddc4fca350e33644301a8c098aa89ac16545b848e61bbda13cc40f6b948e3` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-08-21 | `sha256:935e45e386f2` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:935e45e386f2f2b824e7e5462cd3ddfde3867809d8f26078f0b26de32da07939` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-08-21 | `sha256:f82670c3905b` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:f82670c3905b27d25e4a95ca764e285318eec7899273bf98a98a77df051016ea` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-08-21 | `sha256:f466a0056188` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:f466a0056188b8180c77fbb55ac877c491d304a44bba2d5a5aae50425993289d` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-21 | `sha256:9c1e17003564` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:9c1e170035641c96e8b40d95f1ae61bcf851f5a7e21554f7540ed6aeeea31090` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 522.0 MB | 2026-08-21 | `sha256:709fa8bdbe19` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:709fa8bdbe1923ab3a11d24f668a01bc8020ef922058c678cccc8b3cc50d42d1` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 74 | 430.0 MB | 2026-08-21 | `sha256:ec004c1de2a5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:ec004c1de2a5e31627b82fbc7a62f26743ac5072a70fbde5bb57c92f9896084b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 74 | 457.0 MB | 2026-08-21 | `sha256:459569f82fa9` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:459569f82fa95456ec0362623fb73908d614124bda50be32ba0bb3804c4a29a6` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 74 | 502.0 MB | 2026-08-21 | `sha256:7fae9cdc81cf` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:7fae9cdc81cf2521990d615b471682e22c562ea54c91c057ce338d485dcfe4d1` |

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

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 0 | 4 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.7 | - | 0 | 0 | 4 | 842.0 MB | 2026-08-20 | `sha256:3960d75b1ddd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:3960d75b1ddd28b4e5cffaa13f3ec0c95df95a3be84f67994ace5dc18110f563` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.0 MB | 2026-08-20 | `sha256:76325dbdd581` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:76325dbdd581f039ac2edffca3ff40aef88951da7aa24cf0c07b40992e6d6bd2` |
