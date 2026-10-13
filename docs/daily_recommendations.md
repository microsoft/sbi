# Daily Recommended Images by Language

_Generated: 2026-10-13T02:29:53Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-01 | `sha256:584f42d16916` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:584f42d169167d6a2e9d78344be439f56504baa376b697348406d42dba8094be` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 112.0 MB | 2026-10-01 | `sha256:339937e5e3b3` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:339937e5e3b33d4a909d303d06ccde422f6a4675ef55e828bab803d15a2af374` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.31 | - | 0 | 0 | 0 | 126.0 MB | 2026-10-01 | `sha256:6d7c5efaf40e` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:6d7c5efaf40e11b600cf48ecc3c9e216ddf8ecc905d9755faaf0dbd9c0cf3935` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 132.0 MB | 2026-10-01 | `sha256:e675f8ee67cf` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:e675f8ee67cf6bd133d3180f9570cc714c471998398329569fa5c0a433945793` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 139.0 MB | 2026-10-01 | `sha256:0ab59d8cda13` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:0ab59d8cda137f6aa13ec40ba233aa524b346b2aea6c000f010f6b550de2d593` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 0 | 0 | 962.0 MB | 2026-10-06 | `sha256:7db078926447` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:7db078926447807991fce6040e9f11df8a8a114c6e2d912e7012dade5d283ee1` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-06 | `sha256:07634be0fc64` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:07634be0fc64bf7060907a8cfb0169ff8af84c9a0468a3c8825c28d0dbdebe1e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 1 | 85 | 412 | 875.0 MB | 2026-10-06 | `sha256:ec9c0a0dc5f6` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ec9c0a0dc5f60adc2065762050638dd5ed5facfc53716d4580a42d86027c8e80` |
| 2 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 56 | 255 | 193.0 MB | 2026-10-06 | `sha256:47c36b770db8` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:47c36b770db8f712ceb03e9220d799b068c769b8bc89c61d62cd41600633d9c2` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 56 | 255 | 218.0 MB | 2026-10-06 | `sha256:a3cd573ac05c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:a3cd573ac05cf88ca496e3309cb7f7c44aa16664352d1ee41eda93154f50cdbf` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 0 | 14 | 193.0 MB | 2026-10-02 | `sha256:06a688d0db2f` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:06a688d0db2fa297d420059db33cbc2932b355a6048d92cfdca9da56f386ff0f` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 0 | 14 | 198.0 MB | 2026-10-02 | `sha256:0a1d535179bf` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:0a1d535179bf7d46bb6ba4eab543780a28dd69b903bd54e8bd94ee3c14ac0f20` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 0 | 14 | 203.0 MB | 2026-10-02 | `sha256:b89586dc1778` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:b89586dc17781f25531909993658aa8161205ae38b8cec8847df4a8221a403d5` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 0 | 14 | 217.0 MB | 2026-10-02 | `sha256:c291421b6f20` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:c291421b6f209369e11899c9cc2425ccf08566f2d40b527388c45df036d86b7a` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 0 | 14 | 223.0 MB | 2026-10-02 | `sha256:f6ca5cf8e277` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:f6ca5cf8e277d5fd590ebeb0df890f1039ad8da0a953c94cb4ee8d3f2cae5dff` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 0 | 14 | 230.0 MB | 2026-10-02 | `sha256:222759b391a1` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:222759b391a1aaf241166672c8f99b2d4ada452e7b5319f3c6e8f265a37b5ad4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 0 | 18 | 917.0 MB | 2026-10-02 | `sha256:e70cdb7f80b0` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e70cdb7f80b0348f5cb85f19a8f670fca061f033d57eed12fa003d58b0e06317` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 0 | 28 | 855.0 MB | 2026-10-02 | `sha256:64c94a5a8fac` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:64c94a5a8fac55f965e0fb74b0e6a46b73b7323257c492170befbe5537dfafef` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 11 | 39 | 854.0 MB | 2026-10-02 | `sha256:536b706b6113` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:536b706b6113a3839c1712ebcf832dc33d275fd605d35a931db956792c10d26c` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.9 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-12 | `sha256:b56d1f01c961` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:b56d1f01c961316d412963dc6bc6b7b02a9b84a6e6415d01413172f353a9d415` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 54 | 298 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 324.0 MB | 2026-10-12 | `sha256:3d0d87da3362` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:3d0d87da3362c37bd443e1e53ffe62e81250d89d02a2e263fbbbbd398927816b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-12 | `sha256:f28592f3715c` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:f28592f3715cd6e90c51199ec0bf491cd6e7149ced440305d21d8c203b3b2d1c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-10-12 | `sha256:d8ad6889c401` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:d8ad6889c4017c960b5779b31ab5fba2917674e5dfcac9cc2d3b8dc01fc2b810` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 400.0 MB | 2026-10-12 | `sha256:65789a5d313e` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:65789a5d313e58295d9b9b566f27a6470c2ecfd6674cfb97f26ac553d28acf3b` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 479.0 MB | 2026-10-12 | `sha256:f7e42a4c13de` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:f7e42a4c13de87129f3279072b7344a54028d05e3c10676042d585694a7f8688` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-12 | `sha256:9c3b6e495e05` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:9c3b6e495e05d71f9db65e54d2946c92773baf71b550e26f8fd2ac25d2235831` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 77 | 436.0 MB | 2026-10-12 | `sha256:94efd1775b47` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:94efd1775b47cbb07e4c73a37a36521f254ce7949c28d51c9301330a74912404` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 77 | 463.0 MB | 2026-10-12 | `sha256:29c5ce0b6622` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:29c5ce0b66220c00e6f1c5f0cea8984e22224502a61ed6c988697efdf1e01678` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 77 | 508.0 MB | 2026-10-12 | `sha256:debcc24458e0` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:debcc24458e06b133be4d5edf9fb29320936ba25629bef1f1cb2940e20deaf64` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-10-10 | `sha256:b4f6866da097` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:b4f6866da0977e1bbf9ad4c5cdf265cb841bca03dcc2ea55cfa6a71ab655860b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 8 | 21 | 158.0 MB | 2026-10-10 | `sha256:b6eb40885d72` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:b6eb40885d725facba8541cdbfe29a20c4bb8be1eff5a2a542a8d28494e97de2` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 8 | 21 | 158.0 MB | 2026-10-10 | `sha256:7df99cb1dab0` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:7df99cb1dab062ae3f12e0804e7cb8b0aa013c238dd1dd43320ab08a6bb5027d` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 8 | 35 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 10 | 54 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 14 | 54 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 14 | 54 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 23 | 93 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 20 | 83 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 20 | 83 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-10 | `sha256:0395d881de0e` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:0395d881de0e233e193733d135982ecd603f5f84ddbd8101774c25e973f470c4` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-10 | `sha256:a51284a102e3` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:a51284a102e3f6b295fd94d29bf9f21f664878b40904d588694cffd9c4423d6c` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-10 | `sha256:313afe87b11a` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:313afe87b11a0651a0dc64651b49922947a8eadd2a0305ef0b833baa4059dc3c` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-10 | `sha256:431ae1dc2924` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:431ae1dc2924c41ae54c97c25676990f0c61d6295f42f5d07783c6976f58f1ef` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-10-10 | `sha256:7292f2bcf52c` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:7292f2bcf52c3916758f1a45aba6e4dcb93f466027d7786900ea96c2d23c48e2` |
