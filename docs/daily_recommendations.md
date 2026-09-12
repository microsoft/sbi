# Daily Recommended Images by Language

_Generated: 2026-09-12T02:17:44Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 108.0 MB | 2026-09-04 | `sha256:0aca4c2e1c24` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:0aca4c2e1c2435a74be5254599b6436773483c649acc764aa5bf2898d91025f8` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 112.0 MB | 2026-09-04 | `sha256:f1981f8923ee` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:f1981f8923ee8ecd967f002df6100b556c2ed8f7ab29df932d4625f206badb6f` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.31 | - | 0 | 0 | 0 | 126.0 MB | 2026-09-04 | `sha256:8f0aabeb4e7d` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:8f0aabeb4e7d3accdb42a69db6344962c5844dcfcdb50009a4b5b1927df5f146` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 132.0 MB | 2026-09-04 | `sha256:04acc008da2b` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:04acc008da2bb706efc2158ac1ab4f3f7f3015cfb0ad3d602eb0f996025e8e59` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 139.0 MB | 2026-09-04 | `sha256:9fd8a9a09ba5` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:9fd8a9a09ba5c0c5cb6d95c8200447de9bc46751ff77b8caa141725a29a1178c` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 0 | 70 | 961.0 MB | 2026-09-04 | `sha256:40195d80f8c2` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:40195d80f8c217195a16f65c1be7e8acad646317e9bda2801d91df681c60b3f4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 0 | 80 | 899.0 MB | 2026-09-04 | `sha256:49e105cb9e20` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:49e105cb9e208a4562c904605afb28ea2377d6a83110c8c29fd745d446bc81dd` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 52 | 239 | 193.0 MB | 2026-09-04 | `sha256:9cfa8aaf5c98` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:9cfa8aaf5c98a4cedffec74d450dd2d4510ba72ee9e663efe2658500b0321524` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 52 | 239 | 218.0 MB | 2026-09-04 | `sha256:9a464e9a7e8c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:9a464e9a7e8c6144631020975f703c89034fe386417cb740620df69c2c6cfe24` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 101 | 449 | 867.0 MB | 2026-09-04 | `sha256:5ef85cc12cb2` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:5ef85cc12cb25be6ec319a7392d1e9efd53c3bc8abb971c53d8058a473f09053` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 0 | 12 | 193.0 MB | 2026-09-11 | `sha256:104cc434e0e6` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:104cc434e0e6385522d26e575705f9c6b08020838226b1e1856865bb00bb0f30` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 0 | 12 | 198.0 MB | 2026-09-11 | `sha256:292a2258079b` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:292a2258079b972164633918da438944e8561302884ced44862e4b2a936974ef` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 0 | 12 | 203.0 MB | 2026-09-11 | `sha256:8a153b5889d7` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:8a153b5889d796b6450295b383596b13308c24c230515f8a7770ce1b94e0c460` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 0 | 12 | 217.0 MB | 2026-09-11 | `sha256:ce687d884ca7` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:ce687d884ca748a70af86b901782afbd54214c4b930148c42d85e246c9881c1f` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 0 | 12 | 223.0 MB | 2026-09-11 | `sha256:cdbdb91fbd1d` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:cdbdb91fbd1d95e8561bd126858ebceef0a0f00dd3fe90d72d8b40f0b25602b4` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 0 | 12 | 230.0 MB | 2026-09-11 | `sha256:6a94333d3751` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:6a94333d37514e385650a3c81a55e5350b67253dbe136e9cf17e499c35606a8c` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 0 | 60 | 917.0 MB | 2026-09-11 | `sha256:2fa828c68761` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:2fa828c68761b1b8c23d7662dc134421b9d3b59fe1425fdbc80804e390cdb24d` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 0 | 70 | 855.0 MB | 2026-09-11 | `sha256:d07d5770b2b7` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:d07d5770b2b74717bfe882250e981a4e7a756e0240fc4aefcc0ce93c1e8a9c36` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 11 | 81 | 854.0 MB | 2026-09-11 | `sha256:0aa5a155ba07` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:0aa5a155ba075ad0effc3f60b3af6a696c2a498703519748e09e7eba8c5586f7` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-09-11 | `sha256:6fd09ebf3b22` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6fd09ebf3b2227694000a97256c4f6c07a447519c62e52bbc0e368669be51219` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-09-11 | `sha256:12af5c6399bf` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:12af5c6399bf635af1f199d2cc635bc5c1bb42e0679a5b80933b226acba2305f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-09-11 | `sha256:8bbae400ba9d` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:8bbae400ba9d805ec6ea1f3c707f192f1fec35d62bb1dff71c16f701b5a22e1e` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-09-11 | `sha256:4ae132d3c8b0` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:4ae132d3c8b0135b3061cc90e70421fec0e32fd50cf8217cb4985751a18de9e4` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-09-11 | `sha256:050ccd0e3bd4` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:050ccd0e3bd4f379a49b97486ac7bc038b0506fb970fcc16bfaad29bfedf42be` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 522.0 MB | 2026-09-11 | `sha256:f3fad3885d92` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:f3fad3885d921907149959505057a3e85885dd479863a6e833c7d4a424d396fa` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 94 | 453.0 MB | 2026-09-11 | `sha256:add62053f97c` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:add62053f97cba2be01bf918f46bdd2a3f0688fdb838b42207550ebc0019ac85` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 94 | 480.0 MB | 2026-09-11 | `sha256:4acacc9b5af3` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:4acacc9b5af3d9b21d87815571d739c7cb7ae04b1d81dfc1371510c36378a0c6` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 94 | 525.0 MB | 2026-09-11 | `sha256:d6ace73af767` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:d6ace73af767fca4d48fb91acfffb6da4c6fdd79d4d2d2c39dbc64d7cd67403f` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 0 | 3 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | :24-nonroot | 0 | 4 | 9 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | :24 | 0 | 4 | 9 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 13 | 42 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 8 | 24 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 8 | 24 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 12 | 42 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 12 | 42 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 43 | 117 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 83.6 MB | 2026-09-11 | `sha256:6f7e96128ee1` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:6f7e96128ee1c66f167d8512178c0f6b43eca6dcdd9cd544a591f27ee5d75a4b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 83.6 MB | 2026-09-11 | `sha256:aa58cf2e8d81` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:aa58cf2e8d8190536744d947624014a4d2de744d1bb5f99c8aed97e710172fac` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-09-11 | `sha256:6a4fcacee291` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:6a4fcacee291dffbb7a95881a8549cda0ff0999628975eebeec5b7ce6acb10fa` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 70 | 842.0 MB | 2026-09-09 | `sha256:9f6b5c480db6` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:9f6b5c480db643ca536db638fc0346b5a8424fd738c916db5db6443e68ad130e` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 6 | 84 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-11 | `sha256:f9aa2435862c` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:f9aa2435862cbb05a8d7fbeb9c2d024bf7548f3838f69a96cb214c1bade88ffa` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-11 | `sha256:4377af4aa7a8` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:4377af4aa7a810b7d59f691eae5066895a71aa3eee4cfb4eba527bbebff16479` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.4 MB | 2026-09-11 | `sha256:34a22db497ff` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:34a22db497ff34a0f35ca5fc54bd38711d04238a2c1b2f65d35dc9d45dd82584` |
