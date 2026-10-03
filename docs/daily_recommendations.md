# Daily Recommended Images by Language

_Generated: 2026-10-03T02:18:25Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 0 | 0 | 962.0 MB | 2026-10-01 | `sha256:da4e71e176bd` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:da4e71e176bd65433cfef1e095b2f4f0416b709d5976bd9402c3da487b11d29a` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-01 | `sha256:b64b47f07ab5` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:b64b47f07ab5718c23c8981eee2737c4e4f477c2a1e8b7689e9f426832615b95` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 55 | 262 | 193.0 MB | 2026-09-19 | `sha256:37466ea190f6` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:37466ea190f696105c1c3ae67c15e32d4e199face9a0b2ad5b9a37c464db8f30` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 55 | 262 | 218.0 MB | 2026-09-19 | `sha256:2f202e1169ec` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:2f202e1169ec507bdc07007cf68c14d0ff3a098110b17c460a60185e1f36a9d1` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 105 | 470 | 867.0 MB | 2026-09-19 | `sha256:78235e09001f` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:78235e09001f52b6592c458ac010775ebac6725422e80cd0c1650590f67b2743` |

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

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-10-02 | `sha256:dece9b994dca` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:dece9b994dcace5621e2dba8ced1f91c5de0bf4c7faed4bd7c0d0c31d54fb8ef` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-02 | `sha256:7b364f0e08e4` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:7b364f0e08e4518a848496f0bf99132af414bae6f9b234c23e171d4176c5b9f3` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-10-02 | `sha256:6cd83dc07745` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:6cd83dc077454327243b0e4fd0f35e8e2272bf48bb23961e8907312510c3e597` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-10-02 | `sha256:5610262615ea` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:5610262615ea433b5cdc08dc0172972ac15177301a8ae3aeb202646356b18f0f` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 479.0 MB | 2026-10-02 | `sha256:84cb92c92e97` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:84cb92c92e970dce6b017bdf0495d64b50bee54c6981d8bba7ae37208f4ae12f` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 523.0 MB | 2026-10-02 | `sha256:02105a323f18` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:02105a323f1837ffd915941f03b795bea08299d81f5d9f53d6974607a1ea350a` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 77 | 435.0 MB | 2026-10-02 | `sha256:4eb94422de7a` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:4eb94422de7a774c81d20c3dc7997e9ecccc49e5c02caf34d5f6fe5a360266af` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 77 | 463.0 MB | 2026-10-02 | `sha256:31c92e1f6f60` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:31c92e1f6f60be322703b14f1d06ed74c8c809c191ca8e77a202e93620052d16` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 77 | 507.0 MB | 2026-10-02 | `sha256:2427e09f4418` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:2427e09f4418fc0a322c72bc250e0b6ba34dc98c04ae49bb2fa9a9ac46cd983c` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-09-23 | `sha256:5583f79d517a` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:5583f79d517aff6422b50c8b75b328061cbd2be9ddbc3a14683783952a290b20` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 6 | 33 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 8 | 20 | 158.0 MB | 2026-09-23 | `sha256:aefd87999acb` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:aefd87999acb71f972a7c851f6b2a1cdc11a7d26f16918a180368422f0394722` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 8 | 20 | 158.0 MB | 2026-09-23 | `sha256:eeca1fcd6ace` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:eeca1fcd6acef7ee241a3f362a4ed176f04970a661642b3447580c01c859651d` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 8 | 52 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 14 | 53 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 14 | 53 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 21 | 91 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 20 | 82 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 20 | 82 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-09-23 | `sha256:a654be0e0e6a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:a654be0e0e6a9d872a963597caadc17084bb911f83e3888a9b5a603ff39e620f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-09-23 | `sha256:a95f17a6dbea` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:a95f17a6dbeac3d43ff3652864116a9ff56868ee8e044b29c15896ab6d36d5d5` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-09-23 | `sha256:b006d366ab67` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:b006d366ab67c2170e5ed5b7d34500e0291373df7faedc1d1ca2472944f6c7b2` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 4 | 845.0 MB | 2026-09-30 | `sha256:7e3429507d92` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7e3429507d92bb1e12fe45b7650c3dd3f24a4f0d1ca45622b1452834f54d8020` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 20 | 162 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-23 | `sha256:2b5cec59b51c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:2b5cec59b51cb0509157e3a1b730e19aa42912b99d7f5300b87eac7000bfea19` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-09-23 | `sha256:1324a2cf7ed3` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:1324a2cf7ed34e5f48a1022816b205782b86c7305651658e611dcd3d30756751` |
