# Daily Recommended Images by Language

_Generated: 2026-10-22T02:19:42Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.32 | - | 1 | 58 | 242 | 194.0 MB | 2026-10-21 | `sha256:1b2f288e3a87` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:1b2f288e3a87cd251322f304e176327b6c00ec3bc1847f38b69a58cb7251e994` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.32 | - | 1 | 58 | 242 | 218.0 MB | 2026-10-21 | `sha256:74a1d7c07211` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:74a1d7c07211a7e8a2e35ff3c70ed8e4e1d6451a1ff7cf7bdf526f9eafa23023` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.426 | - | 1 | 96 | 454 | 868.0 MB | 2026-10-21 | `sha256:738f72aff955` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:738f72aff955d5c5690e06b6fc8ecd1b90ead564099a257ac528ff720b8da85f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.32 | - | 0 | 0 | 26 | 193.0 MB | 2026-10-21 | `sha256:d1a464e7b68a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:d1a464e7b68a37abf1739f6d97d3e164a0ed908e8b3bb04a9861b9d940acb876` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.21 | - | 0 | 0 | 26 | 198.0 MB | 2026-10-21 | `sha256:cf7601452fc4` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:cf7601452fc41575ba15538c0f12e5df31d9c0098820a73c222f23bfa736a004` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.13 | - | 0 | 0 | 26 | 203.0 MB | 2026-10-21 | `sha256:0cf09aaccf6a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:0cf09aaccf6a9f68900aed4c73fb28fe25bb59ec85de4dea08ba02240a8635f3` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.32 | - | 0 | 0 | 26 | 217.0 MB | 2026-10-21 | `sha256:9449b45658df` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:9449b45658dfb30286dad89aa55ec3640eb0cea2edb7f6842c9d734920ed9b2b` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.21 | - | 0 | 0 | 26 | 223.0 MB | 2026-10-21 | `sha256:f154a9699db4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:f154a9699db4bc90db46159a3214c8e59dc981524bdc6bc2a9adcc83614cfa82` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.13 | - | 0 | 0 | 26 | 230.0 MB | 2026-10-21 | `sha256:8af87eba5e7c` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8af87eba5e7cdafb1514d7388be287365da97c132509e84dfe252ffd552398ca` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.319 | - | 0 | 0 | 42 | 855.0 MB | 2026-10-21 | `sha256:0e0342bb18ef` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:0e0342bb18ef8c8eb7c5a189463d4840cf595706aa9d0494349a6a64965c5428` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.402 | - | 0 | 4 | 37 | 918.0 MB | 2026-10-21 | `sha256:747cbade5bfd` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:747cbade5bfd918a175ac32f8f1d931edec2c7d0d5397a3de72eaef1ac1f66d2` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.426 | - | 0 | 16 | 59 | 854.0 MB | 2026-10-21 | `sha256:ada801965b46` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:ada801965b46b26af97d6d366300ec06ca793c581259664c226253d9da55d254` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-21 | `sha256:d104fef62900` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:d104fef62900c2ca101f52568f43e0e6b6499b60be324fdd4d2ad5bae4bc3203` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 69 | 339 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 479.0 MB | 2026-10-21 | `sha256:410124e58fe5` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:410124e58fe5f552825fefb0c03f5b0f7581a3572141709babc93dcddfd6ff8c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-21 | `sha256:bf552454269a` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:bf552454269ad43c276e1aaad729fc1ba2bb00d90a24931faf4199719b9da699` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 1 | 1 | 324.0 MB | 2026-10-21 | `sha256:6c0ce63ab516` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6c0ce63ab51617977749e24a2369f0fadc115c4c613e53c3ef6747072e06e8a0` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 1 | 1 | 327.0 MB | 2026-10-21 | `sha256:1bd1bfe2221d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:1bd1bfe2221d7c5294fae10c8d373d8e8d7f58f97b553e7161e253ec4ac96051` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 1 | 1 | 354.0 MB | 2026-10-21 | `sha256:e42bb14a20d7` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:e42bb14a20d7fe8a30d13d603047a3b4e0fa83596ab32edc099b71325e51fa90` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 1 | 1 | 400.0 MB | 2026-10-21 | `sha256:d89a22137a73` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:d89a22137a7360478015105ddcb60b4de94304feac233e09e3c5525448911af4` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 1 | 90 | 431.0 MB | 2026-10-21 | `sha256:fadd15520f39` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:fadd15520f396beb223d4d2951fff9199b019f1b1304edfb9feca603f0f3f7ee` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 1 | 90 | 458.0 MB | 2026-10-21 | `sha256:08c68d2c69d5` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:08c68d2c69d56f1e8096c66136799d33c19bcb5a3b285be9055cd9a432a922c1` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 1 | 90 | 503.0 MB | 2026-10-21 | `sha256:aeb66b5c108a` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:aeb66b5c108aa4a0fcd5d0cacd1f71d8d973aebc693dd31cb2594db8cca215c4` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-16 | `sha256:b266b38f7faa` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b266b38f7faad1d2baf06901d709bf53db1052f6a0b35043d39216ae9df204a7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-16 | `sha256:21f2c8ce7b64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:21f2c8ce7b64860535298ae3907ec56affb9667ce6f1c793aa11ecfc8fbe07d1` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-16 | `sha256:98a774c86033` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:98a774c86033b17a976da188757e57b2ab976379f1233fe21f02bbdf550458c0` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 3 | 9 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 11 | 30 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 11 | 30 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 20 | 80 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 22 | 99 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 24 | 89 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 24 | 89 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 35 | 138 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 30 | 118 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 30 | 118 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-10 | `sha256:431ae1dc2924` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:431ae1dc2924c41ae54c97c25676990f0c61d6295f42f5d07783c6976f58f1ef` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-10-15 | `sha256:af39b0587a0f` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:af39b0587a0fd132e574faabf9aaf56833d69908b9ff227131d519ee219759e7` |
