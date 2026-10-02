# Daily Recommended Images by Language

_Generated: 2026-10-02T02:22:58Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 55 | 259 | 193.0 MB | 2026-09-19 | `sha256:37466ea190f6` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:37466ea190f696105c1c3ae67c15e32d4e199face9a0b2ad5b9a37c464db8f30` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 55 | 259 | 218.0 MB | 2026-09-19 | `sha256:2f202e1169ec` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:2f202e1169ec507bdc07007cf68c14d0ff3a098110b17c460a60185e1f36a9d1` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 105 | 466 | 867.0 MB | 2026-09-19 | `sha256:78235e09001f` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:78235e09001f52b6592c458ac010775ebac6725422e80cd0c1650590f67b2743` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 2 | 26 | 193.0 MB | 2026-09-21 | `sha256:4b8f52c09215` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4b8f52c092153540f8d3ae94181f7749f4ced9d076bdc8e7c8742d4b517077f4` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 2 | 26 | 198.0 MB | 2026-09-21 | `sha256:b016fbf79133` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:b016fbf79133cb79f57a06f7b5df756eca697fc49bd74ab539f10d994d918edf` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 2 | 26 | 203.0 MB | 2026-09-21 | `sha256:ff17a18b639a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:ff17a18b639a0327e52c7c296fa2e1abe6e03eb61d8121a8ef67cc6aa430a27e` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 2 | 26 | 217.0 MB | 2026-09-21 | `sha256:84892b9bf258` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:84892b9bf258042b0b9175509292b9158e8c90c721f1e6a1fa8256f21a652764` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 2 | 26 | 223.0 MB | 2026-09-21 | `sha256:196831e5c6a2` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:196831e5c6a26dba1c0db2f15856332938c2cc6d9192715e45d935ef679e6af9` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 2 | 26 | 230.0 MB | 2026-09-21 | `sha256:2d584d8147fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:2d584d8147faddb0d678c5748d47953e5b8e18621ed4fb7049a91381d9d7746f` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 2 | 53 | 917.0 MB | 2026-09-21 | `sha256:35d40304542c` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:35d40304542c8689331f8cab17c65926cdf48fe711e289321d71924b230a7d29` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 2 | 63 | 855.0 MB | 2026-09-21 | `sha256:0f97a4002de8` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:0f97a4002de8050867e7e55e03a0487aea4f4a933158fe1d2cc35f570fc81a0d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 13 | 74 | 854.0 MB | 2026-09-21 | `sha256:2e171ed9da38` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2e171ed9da38a01abb2005b8b2ce9df0c40e583a815e977a2c84d47778005dde` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 0 | 845.0 MB | 2026-09-30 | `sha256:7e3429507d92` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7e3429507d92bb1e12fe45b7650c3dd3f24a4f0d1ca45622b1452834f54d8020` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 20 | 158 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-09-30 | `sha256:b71fc487d268` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:b71fc487d2689839a1ccefd51b7dbdbe71e85f636593c5741f3a9b6cb011583c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-09-30 | `sha256:d11f59923b1a` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:d11f59923b1a98f4f40fc6efe4ee3a6e03d355c94630768302ce9b1f2d043b6f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-09-30 | `sha256:7e3f3e40cd49` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:7e3f3e40cd49738c5cb284ec5c1ac9306704c3410a2ee3cd30135614003cb974` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-09-30 | `sha256:3ee3e8f26b9c` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:3ee3e8f26b9c0d72108efbc996766e165579c1f8d45aa941e7aca4c7e8f7f3fb` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 481.0 MB | 2026-09-30 | `sha256:4d26725a1ccb` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4d26725a1ccbe10eed239eb8d92e66fb10993df5ea1a5ecbe75c65d1ca3c31ad` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 526.0 MB | 2026-09-30 | `sha256:6ca9ba67bf61` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:6ca9ba67bf61dc24fcf95cfae096b44876f3331cc962cfc3a54f129e8a4797bb` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 77 | 460.0 MB | 2026-09-30 | `sha256:24a94e287f3f` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:24a94e287f3f49167b442cb181277322f63b96a72f88fc02c898e61e17158cf2` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 77 | 488.0 MB | 2026-09-30 | `sha256:2f667ec86805` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:2f667ec868058bcd19387175c3c27d0eb29b28a86e88d4cdafa07db15c78f40c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 77 | 532.0 MB | 2026-09-30 | `sha256:02f7c9295fcc` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:02f7c9295fccc595728c1c65fb98d98d453ce9e3acc023cab797acf11a717ba0` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-09-23 | `sha256:5583f79d517a` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:5583f79d517aff6422b50c8b75b328061cbd2be9ddbc3a14683783952a290b20` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 6 | 33 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 7 | 19 | 158.0 MB | 2026-09-23 | `sha256:aefd87999acb` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:aefd87999acb71f972a7c851f6b2a1cdc11a7d26f16918a180368422f0394722` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 7 | 19 | 158.0 MB | 2026-09-23 | `sha256:eeca1fcd6ace` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:eeca1fcd6acef7ee241a3f362a4ed176f04970a661642b3447580c01c859651d` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 8 | 52 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 13 | 52 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 13 | 52 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 21 | 91 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 19 | 81 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 19 | 81 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-09-23 | `sha256:a654be0e0e6a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:a654be0e0e6a9d872a963597caadc17084bb911f83e3888a9b5a603ff39e620f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-09-23 | `sha256:a95f17a6dbea` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:a95f17a6dbeac3d43ff3652864116a9ff56868ee8e044b29c15896ab6d36d5d5` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-09-23 | `sha256:b006d366ab67` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:b006d366ab67c2170e5ed5b7d34500e0291373df7faedc1d1ca2472944f6c7b2` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-23 | `sha256:792ea6ed971a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:792ea6ed971a69bbca863882354d4ff197ca461e0f699d0256653f7886a62d42` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-23 | `sha256:2b5cec59b51c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:2b5cec59b51cb0509157e3a1b730e19aa42912b99d7f5300b87eac7000bfea19` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.8 MB | 2026-09-23 | `sha256:1324a2cf7ed3` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:1324a2cf7ed34e5f48a1022816b205782b86c7305651658e611dcd3d30756751` |
