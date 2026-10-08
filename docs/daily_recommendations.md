# Daily Recommended Images by Language

_Generated: 2026-10-08T02:27:44Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 1 | 82 | 396 | 875.0 MB | 2026-10-06 | `sha256:ec9c0a0dc5f6` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ec9c0a0dc5f60adc2065762050638dd5ed5facfc53716d4580a42d86027c8e80` |
| 2 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 54 | 249 | 193.0 MB | 2026-10-06 | `sha256:47c36b770db8` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:47c36b770db8f712ceb03e9220d799b068c769b8bc89c61d62cd41600633d9c2` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 54 | 249 | 218.0 MB | 2026-10-06 | `sha256:a3cd573ac05c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:a3cd573ac05cf88ca496e3309cb7f7c44aa16664352d1ee41eda93154f50cdbf` |

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
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-07 | `sha256:6ff8558f06dd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:6ff8558f06ddcb8673c73629f1500ac7238cc00d8c714b1884baa821e212ad6e` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 20 | 162 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-10-07 | `sha256:838328420277` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:838328420277c5fb439ca7e3cf9f694124a733d3f0b3d120edd8f2b2b88f5cc4` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-07 | `sha256:8074a0fe16b1` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:8074a0fe16b12da6ab804381cc86ebd201c403335d075740ec68f612bc0c0f94` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-10-07 | `sha256:7fc30c462f62` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:7fc30c462f6227640b522b7323d8fcfd147e96dee9a2dda7379e010397521eae` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-10-07 | `sha256:6fb36103f5d1` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:6fb36103f5d1904f46486622e77886aa0389d57467f78ddff72b1fd3aeaa3fb3` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 479.0 MB | 2026-10-07 | `sha256:deacce81f1a2` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:deacce81f1a2173690cf674b2377c6a23cdbab04a6eddde50501e968d1918f5c` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 523.0 MB | 2026-10-07 | `sha256:904ce51ad4aa` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:904ce51ad4aa632ecc232a4715201019ff040a79932f59baa94bb1294b98d586` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 78 | 435.0 MB | 2026-10-07 | `sha256:594087a43a0a` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:594087a43a0a795184ce759b3f2f489a3f4b14302eaf7471d3abb6b4cf388d49` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 78 | 463.0 MB | 2026-10-07 | `sha256:fbfd4fa69a3f` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:fbfd4fa69a3fb1396baf319a59750dc29a78b358342e6afabf818d6622736e27` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 78 | 507.0 MB | 2026-10-07 | `sha256:f181a5a6dbee` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:f181a5a6dbee409127fc054e2a9fe443594d69465a88a0c4a8cc6cf46184b444` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-10-06 | `sha256:c30e39a396fb` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:c30e39a396fb0ad2c7bf2fe6b4e2fe1598a9b79f09831c07d763fe09927e8779` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 6 | 33 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 8 | 21 | 158.0 MB | 2026-10-06 | `sha256:be12f1235ae6` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:be12f1235ae675a4246ac2a276ac20aac8b573e70914f726c17b96615a573ced` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 8 | 21 | 158.0 MB | 2026-10-06 | `sha256:0375789e568e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:0375789e568e198d670b2c14e99493adfa07068a9f9615c9a0664b9df4dd8fc1` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 8 | 52 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 14 | 54 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 14 | 54 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 21 | 91 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 20 | 83 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 20 | 83 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-06 | `sha256:bb3696f7552a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:bb3696f7552a7f9a6550558343dff5a4d45c58fc38e1eb93c7493d685a68db17` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-06 | `sha256:88eb93599227` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:88eb935992271c4ab82773efbb34290e7448c9df778c9505d0e5b184ec25d07e` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-06 | `sha256:d1a693353b7d` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:d1a693353b7db0383e66a61416a049f617fdfb89435dcaac543749fb4f14950a` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-23 | `sha256:2b5cec59b51c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:2b5cec59b51cb0509157e3a1b730e19aa42912b99d7f5300b87eac7000bfea19` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-10-05 | `sha256:bfd3e44899fe` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:bfd3e44899fe7c17f6fda42a6ef2a322f2178c2dafb88e69fd87675cdcac39ec` |
