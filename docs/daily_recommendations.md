# Daily Recommended Images by Language

_Generated: 2026-10-17T02:20:15Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.21 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-13 | `sha256:72b00129426f` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:72b00129426fcd37d04eadcd3184e41a6a9ff19b679b0748386367648fc4f84f` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.13 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-13 | `sha256:e21e0262d286` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:e21e0262d286b0d23e7ae60ab8da09110b4ded4a72e9faa2712f634d827694ef` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.32 | - | 0 | 0 | 0 | 126.0 MB | 2026-10-13 | `sha256:42d59c90812c` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:42d59c90812ce53255a3102225c4b5a51fbd2471b32e34289fdd5ea10155da4e` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.21 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-13 | `sha256:c8fddc889464` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:c8fddc889464be4889e8afbac395983c63a7c9d673475205ca5310e5eb1c777d` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.13 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-13 | `sha256:42f3714eaa91` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:42f3714eaa91aad7843633f943221f3a36f7894b7d93f512177e9757861c024f` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.319 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-16 | `sha256:18f5f0f1c8a8` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:18f5f0f1c8a81689880383c95816da71d74f4a66757ba85ed9e79cb804fdbb53` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.402 | - | 0 | 4 | 5 | 963.0 MB | 2026-10-16 | `sha256:5798415b0662` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:5798415b0662aa6443c87d5dfcddb4e7a189431310365442a0246ade7707893d` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.426 | - | 1 | 96 | 454 | 875.0 MB | 2026-10-13 | `sha256:f7e28a0a3f6f` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:f7e28a0a3f6f9235929ec19274fd0ddf8534e9159f99fac602ac73304f1ce99f` |
| 2 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.32 | - | 4 | 62 | 255 | 194.0 MB | 2026-10-13 | `sha256:60d11c2be381` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:60d11c2be381e7f90b8bd984472dd2432933858aba55087d5936980cf2071d5a` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.32 | - | 4 | 62 | 255 | 218.0 MB | 2026-10-13 | `sha256:160138077567` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:1601380775673021a23ea5985669f1daf03719faaa684f839da1752fb48cfc7f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.32 | - | 0 | 0 | 26 | 193.0 MB | 2026-10-13 | `sha256:32c87009a88e` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:32c87009a88e78b6cbac570084dd1b144988454273221c79fd4ccc245e479b13` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.21 | - | 0 | 0 | 26 | 198.0 MB | 2026-10-13 | `sha256:51d9c14e5df3` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:51d9c14e5df3669a21f336b2411d8fe2fa94e03a2bc3f2baac241d9d51e15246` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.13 | - | 0 | 0 | 26 | 203.0 MB | 2026-10-13 | `sha256:939555726cae` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:939555726cae4c916d212323ac3c0792996dcac2614f00e82cc40fcf059e1af3` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.32 | - | 0 | 0 | 26 | 217.0 MB | 2026-10-13 | `sha256:6f9add0af1b2` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:6f9add0af1b2e16f9fe690c016f0c4fe847a064abb94b49bbc77cb5345d00d38` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.21 | - | 0 | 0 | 26 | 223.0 MB | 2026-10-13 | `sha256:7da398155e0d` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:7da398155e0d6439cf0058bc1afa937817a03e320714b876a5f58a7d64f79d50` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.13 | - | 0 | 0 | 26 | 230.0 MB | 2026-10-13 | `sha256:d5633778c754` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:d5633778c754e89a386882eddf8dbee6942447a8a4b8267c19cf1e1158d0b0ac` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.319 | - | 0 | 0 | 69 | 855.0 MB | 2026-10-13 | `sha256:95f436a35242` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:95f436a35242b0c595ed3c9379ba1a358be9aa52cf33749a780f0393c6c851d6` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.402 | - | 0 | 4 | 64 | 918.0 MB | 2026-10-13 | `sha256:0d99ebf1ccad` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:0d99ebf1ccad15037b5c98c9fe8979d6b8dbac23812feae2277a0b277b5c5ffa` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.426 | - | 0 | 16 | 86 | 854.0 MB | 2026-10-13 | `sha256:35f39dc26de2` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:35f39dc26de234f9bd24d482b8c7a9391d263dfbc1e62f24bb59de1168dd8d21` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.9 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-14 | `sha256:39876189c354` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:39876189c354904163b60a5d24e5bcce52e7285a43586cbdeb1bdc243e210d46` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 62 | 326 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 324.0 MB | 2026-10-16 | `sha256:461a519fc064` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:461a519fc064afb0e1d378819a6bdb4ebfaf2fbe6b79f14fdaa90f01b0db0c9f` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-16 | `sha256:40b25af15c93` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:40b25af15c939bd4b30478f72db91b0f9333ad47fb40507485bcff3facee6c9d` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-10-16 | `sha256:502828e07c69` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:502828e07c6925b2de332f9c9c7dda8c7d939ae39d7eb147cee5115ce6a208bf` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 400.0 MB | 2026-10-16 | `sha256:e9adec1bed64` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:e9adec1bed64e18ddecf94a20bdc4ef1e5fe107a4f9dfeb068cc9da7b89c7f5d` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 479.0 MB | 2026-10-16 | `sha256:e284d4837e3b` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:e284d4837e3b46d435f74fc05d83480ba0c42fcf0ecbcf68a4540aa864524f68` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-16 | `sha256:7d71f8106021` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:7d71f81060212f6ef5b206acbd189e86957649c2b89522c98ce7c7af59a4a5f0` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 90 | 437.0 MB | 2026-10-16 | `sha256:802c01106979` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:802c011069791c6d922d42410e433fcd33b9e5dcdd679fce506b04a4147bfa6b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 90 | 464.0 MB | 2026-10-16 | `sha256:1deedd1c6e2e` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:1deedd1c6e2ea1cdc885325d6b96be087aef43b1db66be8d84c060c33956665a` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 90 | 509.0 MB | 2026-10-16 | `sha256:80f3656b4581` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:80f3656b458111134f99fe2102ec47a741c4a468f39e9043a489a82d5a0ffc16` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 8 | 21 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 8 | 21 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 17 | 65 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 19 | 84 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 21 | 80 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 21 | 80 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 32 | 123 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 27 | 109 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 27 | 109 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-16 | `sha256:b266b38f7faa` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b266b38f7faad1d2baf06901d709bf53db1052f6a0b35043d39216ae9df204a7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-16 | `sha256:21f2c8ce7b64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:21f2c8ce7b64860535298ae3907ec56affb9667ce6f1c793aa11ecfc8fbe07d1` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-16 | `sha256:98a774c86033` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:98a774c86033b17a976da188757e57b2ab976379f1233fe21f02bbdf550458c0` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-10 | `sha256:431ae1dc2924` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:431ae1dc2924c41ae54c97c25676990f0c61d6295f42f5d07783c6976f58f1ef` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-10-15 | `sha256:af39b0587a0f` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:af39b0587a0fd132e574faabf9aaf56833d69908b9ff227131d519ee219759e7` |
