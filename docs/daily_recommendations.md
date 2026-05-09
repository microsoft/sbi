# Daily Recommended Images by Language

_Generated: 2026-05-09T03:11:01Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 107.0 MB | `sha256:f57b292ed7d6` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:f57b292ed7d6f7c88d3b53addfcc02e520aae273eddb57cf0ea49f9396cb2424` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 112.0 MB | `sha256:79af73677bbc` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:79af73677bbc271ed7789d11df2f25974b65e7bb85399b6e9be06c7c37122fa7` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.26 | 0 | 0 | 0 | 126.0 MB | `sha256:742fe85d9ef9` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:742fe85d9ef99cd3d43fc11d4e97e1fb3fe5d30d02d183c7d52aa6238f05904c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 132.0 MB | `sha256:22b1bd549a20` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:22b1bd549a200249bea512b6b2fce8885a47d98c346d0d33511b2bf70653d207` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 139.0 MB | `sha256:b782ff63dec5` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:b782ff63dec54c5537b217ef1f3ab1900058f070c786bc214687c6fe53bba1e6` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.203 | 0 | 6 | 31 | 943.0 MB | `sha256:4f08dbd55f00` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:4f08dbd55f00d1f7eb232ddacb24ecdab62fd64774a25c1da5b6daeb76e76a75` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.313 | 0 | 7 | 32 | 888.0 MB | `sha256:13b64ec155e9` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:13b64ec155e9aefa220263957b6d86f1add651106fc5df3cbecacb7a74eb3b06` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.26 | 2 | 13 | 148 | 193.0 MB | `sha256:513ab9d7d5cb` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:513ab9d7d5cbc75e81fa933113a688d00a178a18ffa54c5631eed92902498e87` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.26 | 2 | 13 | 148 | 218.0 MB | `sha256:1657ced1f401` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:1657ced1f4019bf174ff6ca844ad2545af44e787c135549a843f6a6d609d1c0e` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.420 | 2 | 29 | 280 | 850.0 MB | `sha256:27ee23dc62a9` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:27ee23dc62a99643cfab2dc291582b6f2611ddc2598a0ff0958ddd0972012e52` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.26 | 0 | 0 | 22 | 193.0 MB | `sha256:f1994beca66a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f1994beca66a16766311f5234fbfc9a0ce6efb1fa400fa0a3cd95eb1fd6f2a71` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.15 | 0 | 0 | 22 | 198.0 MB | `sha256:08ec9b342063` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:08ec9b342063d70cf18e7ee10a3f95f22180be321af77e03a8ed2151093ef65f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.7 | 0 | 0 | 22 | 203.0 MB | `sha256:8fb7ff015fcf` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:8fb7ff015fcf0ebc6e90105bd6db06875954e6dc3d374b9dbb34c732867d13e4` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.26 | 0 | 0 | 22 | 217.0 MB | `sha256:b0811426a260` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:b0811426a2602ce0be3d1f605257c601b38c3a26e7b49283c9330de01c67cefd` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.15 | 0 | 0 | 22 | 223.0 MB | `sha256:4180d760569f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:4180d760569fdf8f639cc79d62de294356c166eeebe3e4ab19d4f012960b5be8` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.7 | 0 | 0 | 22 | 230.0 MB | `sha256:55e37c7795bf` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:55e37c7795bfaf6b9cc5d77c155811d9569f529d86e20647704bc1d7dd9741d4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.313 | 0 | 2 | 60 | 837.0 MB | `sha256:c4d10394421c` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:c4d10394421c0e7c116b6f8866b7c80fa74352ea294629cda1fe4f2acd8f8f6f` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.203 | 0 | 2 | 62 | 892.0 MB | `sha256:8a90a473da52` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:8a90a473da5205a16979de99d2fc20975e922c68304f5c79d564e666dc3982fc` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.420 | 0 | 12 | 70 | 836.0 MB | `sha256:a91a987babf7` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:a91a987babf792bfe535ad5d85e94c492f66ffc2efc25aea343432a5bd606e93` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.10 | 0 | 0 | 0 | 822.0 MB | `sha256:1c574d476ff9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:1c574d476ff9e75e5d0ecca45476cf9a02d9d725c7be5be67260289c2ca115c1` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.3 | 0 | 0 | 0 | 854.0 MB | `sha256:29d44da27f9e` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:29d44da27f9e66f74e24673a79e86a53bfe575752014611df58e910ce135f6fb` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | 0 | 0 | 0 | 488.0 MB | `sha256:e6c8aa31d624` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:e6c8aa31d624f38218ab52fc21d94f960e6b3acd1e3646f35211020f4de822ce` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | 0 | 0 | 0 | 532.0 MB | `sha256:89939a506046` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:89939a506046deef73522b93948ae52dcaeb957453c09cbf1d3f41f27060bb90` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | 0 | 0 | 18 | 321.0 MB | `sha256:8be4d3180880` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:8be4d3180880b3f341323d6a5ac3b634d4ea1d657deca144e6e1745f99443b36` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | 0 | 0 | 18 | 324.0 MB | `sha256:da7be11f4c21` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:da7be11f4c211f9b8a2226ee385217045e4658a2a6e908e19071638f5d269c3b` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | 0 | 0 | 18 | 351.0 MB | `sha256:e5daf017dca2` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:e5daf017dca2af648d17638604acaf91c0d93e0e7c3d201cbd2e1190d5144ed0` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | 0 | 0 | 18 | 397.0 MB | `sha256:d6ed7ad5b480` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:d6ed7ad5b48080603110af68475cea04a8efff0c64fba4c0babe179598f8849b` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | 0 | 0 | 93 | 428.0 MB | `sha256:4bee6f4dece5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:4bee6f4dece598711f159232c22c186f121481ee02cd83588b232559557bcbc8` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | 0 | 0 | 93 | 455.0 MB | `sha256:b8675968c362` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:b8675968c36256ec5f4e10a1af58a39e1912dccc9ebbe99b83adb86fe58628f7` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | 0 | 0 | 93 | 499.0 MB | `sha256:481d82ceb702` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:481d82ceb7020e99cf6a325e3d5857603b34c3fe58885a7753d556e641a7f47e` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 488.0 MB | `sha256:e6c8aa31d624` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:e6c8aa31d624f38218ab52fc21d94f960e6b3acd1e3646f35211020f4de822ce` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 532.0 MB | `sha256:89939a506046` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:89939a506046deef73522b93948ae52dcaeb957453c09cbf1d3f41f27060bb90` |
| 3 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 822.0 MB | `sha256:1c574d476ff9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:1c574d476ff9e75e5d0ecca45476cf9a02d9d725c7be5be67260289c2ca115c1` |
| 4 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 854.0 MB | `sha256:29d44da27f9e` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:29d44da27f9e66f74e24673a79e86a53bfe575752014611df58e910ce135f6fb` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 9 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 4 | 28 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 10 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 4 | 28 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 8 | 45 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.13.0 | 0 | 8 | 45 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 23 | 52 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 23 | 52 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.13.0 | 0 | 23 | 52 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.13.0 | 0 | 23 | 52 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 10 | 45 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 1 | 10 | 45 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 1 | 15 | 43 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 1 | 15 | 43 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 10 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 4 | 27 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |
