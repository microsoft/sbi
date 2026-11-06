# Daily Recommended Images by Language

_Generated: 2026-11-06T02:20:46Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **46 configured sources** across 12 groups (see [repositories.json](../config/repositories.json)):

**Base / minimal images (no runtime):**

- `mcr.microsoft.com/azurelinux/base/core:3.0`
- `mcr.microsoft.com/azurelinux/distroless/base:3.0`
- `mcr.microsoft.com/azurelinux/distroless/minimal:3.0`
- `mcr.microsoft.com/azurelinux/busybox:1.36`

**.NET self-contained / Native AOT deployment bases:**

- `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled`
- `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless`

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
- `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0`

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

**.NET Ubuntu Chiseled images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled`
- `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra`

**.NET Alpine images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-alpine`

**.NET Debian images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

**Rust build images:**

- `mcr.microsoft.com/azurelinux/base/rust:1`

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

### Alpine

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine` | 10.0.14 | - | 0 | 0 | 0 | 121.0 MB | 2026-11-05 | `sha256:486297ae006e` | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine@sha256:486297ae006e95ee2ee3112a2c03e3d26e41fbbc64b7a29e28e7a18cc36b12bb` |

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 102.0 MB | 2026-10-29 | `sha256:68c4eda2e145` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:68c4eda2e1459e4284a5ea593cb131f264d624ac0bbd8b31671caaeeee13b707` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-29 | `sha256:369fd8d3246f` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:369fd8d3246f76695dbd5a3ed521982d099a8c97d64eba40f6ad9de704c8d7b4` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-29 | `sha256:9bf51dbfda85` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:9bf51dbfda85c9289bca3606d2c79c11a09457dbf0f2ed4f23a9e2b9b19dde4b` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 127.0 MB | 2026-10-29 | `sha256:e18779f67a33` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e18779f67a33317667fdb5e375fc5f9f9b79ed43710f01f920ad123e3ec68986` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-29 | `sha256:1c77a718863a` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:1c77a718863a5ac26cb17d125fecef8bba1f3f2bb25e19062f6bce8619a6458c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-29 | `sha256:79cb8e7d5437` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:79cb8e7d543740ccdae265ad2e98e4939946f8f43b91efbcef31bdb94b234cf1` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.403 | - | 0 | 0 | 41 | 951.0 MB | 2026-10-29 | `sha256:ffd76bb46785` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:ffd76bb46785445d2d029809009d35b522619b97df88c0abaa05d549db697b43` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.320 | - | 0 | 0 | 51 | 900.0 MB | 2026-10-29 | `sha256:17e7ceadf4df` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:17e7ceadf4dfab25b6d5fbbcd516320c3950b11365f5e17d808186e3364ee082` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.427 | - | 0 | 16 | 68 | 899.0 MB | 2026-10-29 | `sha256:224ae6ffdac4` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:224ae6ffdac4888bb88bbc8a9d0d482b9fabf6a764f51b4aa6e771a6985d051e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.33 | - | 1 | 33 | 217 | 194.0 MB | 2026-11-05 | `sha256:09f71a9459e2` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:09f71a9459e260fe4bab0528f0bd9bfbfed7c8330bed3aa5b665281760205bec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.33 | - | 1 | 33 | 217 | 218.0 MB | 2026-11-05 | `sha256:b14a41431d25` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b14a41431d252c640aae69b52320e512050fc53bafcb5012b97ccb9ca173c2d5` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.427 | - | 1 | 73 | 431 | 868.0 MB | 2026-11-05 | `sha256:341aa6e8bbac` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:341aa6e8bbaceb1194732328ca27c15bfa9c531687db83a66af3c0c1833f5534` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.14 | - | 0 | 0 | 2 | 124.0 MB | 2026-11-05 | `sha256:f2557fb81b1c` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:f2557fb81b1cfcfdfcde9d3bf247ed60b2cd38f00a88153f61e34737fab8ade7` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.14 | - | 0 | 0 | 3 | 162.0 MB | 2026-11-05 | `sha256:a9298ccc39c6` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:a9298ccc39c6258ae71228479e2f3f8d2ceb6639a9927262a4b0ff7d319a0e7e` |
| 3 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.33 | - | 0 | 2 | 28 | 193.0 MB | 2026-11-05 | `sha256:4c1e0075bc54` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4c1e0075bc541cb181872f11b24bd1ed4aa4618517011b1663269482c156d67a` |
| 4 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.22 | - | 0 | 2 | 28 | 198.0 MB | 2026-11-05 | `sha256:ed8d3f58736b` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:ed8d3f58736ba1ccd954aa3d13e582830fb799848e9b739c4d9b8fd0f028f688` |
| 5 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.14 | - | 0 | 2 | 28 | 203.0 MB | 2026-11-05 | `sha256:4f1f30cf0bb4` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:4f1f30cf0bb4ab0c8b2daa84613051683456f8df4a9bf7c812e20d74962992ac` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.33 | - | 0 | 2 | 28 | 217.0 MB | 2026-11-05 | `sha256:188854d2ed80` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:188854d2ed80c2768395bddde1af897ebe562509fbbbae8a3c1b504e42c65eae` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.22 | - | 0 | 2 | 28 | 223.0 MB | 2026-11-05 | `sha256:12e7bdfe76c8` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:12e7bdfe76c8210a5e47e365ff90330474e4a3504e122479dc4ded3fdf52de43` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.14 | - | 0 | 2 | 28 | 230.0 MB | 2026-11-05 | `sha256:1a48394b6f29` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:1a48394b6f29d3a32c79d9edecd4abda0ecbe8776d567dbfc4590ed6979c740c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.403 | - | 0 | 2 | 37 | 906.0 MB | 2026-11-05 | `sha256:212fe4cb7462` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:212fe4cb74627d5adf9817a4088906ae3edff1bd853bb27565d4b46731efe61e` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.320 | - | 0 | 2 | 47 | 855.0 MB | 2026-11-05 | `sha256:61c49333124a` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:61c49333124a241d3c14dee86467e8978aab73d4e569dcb1586dddeb9d844bd9` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33 | - | 0 | 0 | 0 | 324.0 MB | 2026-11-05 | `sha256:8b3c785bdd4b` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:8b3c785bdd4b281f6fcef7e379927f58dac03f4ae75d43f58c906968b159b28b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 0 | 0 | 327.0 MB | 2026-11-04 | `sha256:ff980531a09d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:ff980531a09d5d704abc30ec001b674a3ec2cf6876c2f3a925a81a617c365461` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 0 | 0 | 355.0 MB | 2026-11-04 | `sha256:4314d33822cf` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:4314d33822cff14d4b4954593c3a4b8b290974bd8e29c962a3bb0186c46c066a` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 0 | 0 | 400.0 MB | 2026-11-04 | `sha256:e9b08176aa05` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:e9b08176aa054bee301aebd3d039a7034621780360078e08a15b4b6fc8468ae6` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-11-04 | `sha256:a556447b524b` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:a556447b524b528dc14928229abf52f555bd04627b1330bb29d6b863583c9d9d` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-11-04 | `sha256:d80e32fb93c1` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:d80e32fb93c16b34f8d2af115c4a08664f0d1c66b47e1a72797cf7237ffef126` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 2 | 85 | 431.0 MB | 2026-11-04 | `sha256:e4a8c72b11ea` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:e4a8c72b11eae72b5b95bf0481207afe61f034601643312ab73e8a62f7b90906` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 2 | 85 | 458.0 MB | 2026-11-04 | `sha256:f509c2eacc1d` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:f509c2eacc1d2c7126ff6cc87cdda4bba003381ccdf7a014f218e05b6beed81c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 2 | 85 | 503.0 MB | 2026-11-04 | `sha256:9e7c3ba2316c` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:9e7c3ba2316c7714c4c6e87a1da978f98e1cc5fb23d1a41ac1841426203b09d9` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.1 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-10-28 | `sha256:affe8e33adb5` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:affe8e33adb5aa9039d6bf37156fd07d0eb8bb3ace44bd178835484dc2911b5f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.1 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-10-28 | `sha256:546ae02c9a35` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:546ae02c9a3537e0a322fe9e594da019350273b82b7e56929bdb32b5482e33cd` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.1 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-10-28 | `sha256:4997e0b8d0dd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:4997e0b8d0dd2058d7fd7fc3023dde06913f1d164e043ed147cb03620ef531af` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 3 | 61 | 267 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | - | 4 | 12 | 37 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 4 | 22 | 65 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 4 | 22 | 65 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 4 | 29 | 108 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 4 | 31 | 127 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 4 | 35 | 124 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:9c735f30f9d2` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:9c735f30f9d2344b4af3f8a2f715587ca1cd1afeab8c3c9836ac69d8326637b7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:10bcfa3c563f` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:10bcfa3c563fd05f9b11495551bd36bfe919b4a2f32b68388e1c393bbeea64ac` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-28 | `sha256:71a1f5a1821f` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:71a1f5a1821f4c56e45aff52e792ed6c0500631e201cfc140c4c9d22886b29f9` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 35 | 843.0 MB | 2026-11-04 | `sha256:129cab4502dd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:129cab4502dd2554942c3d8ee8910bf32d7f9017371a3f72b23fe55decfb7dae` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 6 | 69 | 383 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 38 | 1.74 GB | 2026-10-28 | `sha256:d554fbb204d8` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:d554fbb204d8330fd3a64234a07e4f0a6b51c0c2ac2ef48b000ea74a25e9cf8d` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.9 MB | 2026-10-28 | `sha256:ca4f1b9b5d4b` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:ca4f1b9b5d4bb21445fbf5cf7288063ea477a9fa88affa4b13757245e3214914` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-10-28 | `sha256:19ffa1dd7adf` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:19ffa1dd7adfb8644954eff77ce03a1e6617e710a8137b2fccfcb434db774314` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.8 MB | 2026-10-29 | `sha256:d523c75e720b` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:d523c75e720b64058b529dbbb1b6ec4023b3e54f3b83583fef4de4ed73925922` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-28 | `sha256:e95286512259` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:e952865122598043cfb54f0257f016e28e02c21904e172b47dcb2148e095693a` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-10-28 | `sha256:88fc2b4a4258` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:88fc2b4a425801b3b4d8fe5342e24fb6134669fb718ec970340bdd36dee50c35` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 2 | 14.2 MB | 2026-11-05 | `sha256:e9346023014b` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:e9346023014b406b9bc8be969b3aa4d39efa3a43b6abcc495f01be566acb10e9` |
