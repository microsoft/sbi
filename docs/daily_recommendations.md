# Daily Recommended Images by Language

_Generated: 2026-11-10T02:18:34Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine` | 10.0.14 | - | 0 | 0 | 1 | 121.0 MB | 2026-11-05 | `sha256:486297ae006e` | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine@sha256:486297ae006e95ee2ee3112a2c03e3d26e41fbbc64b7a29e28e7a18cc36b12bb` |

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 102.0 MB | 2026-11-07 | `sha256:b18510f0c023` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:b18510f0c023eb5809ab0d3f074a09990222bf1c7b54837a12a43fa8f1cd2840` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 108.0 MB | 2026-11-07 | `sha256:6fe5550db17e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:6fe5550db17e001c0a9a48671b1ee9fd74435192417e032156ebba38be242c1f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 113.0 MB | 2026-11-07 | `sha256:0306c8a5fcd8` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:0306c8a5fcd8af4ebe7c7c52765e3be99e10df01b4a35ed08dddbf869f51ac4d` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 127.0 MB | 2026-11-07 | `sha256:2274d664b6de` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:2274d664b6de6fd5f02ff2d64f2e62089d86dc30d7c804512b04d4f7c3f7a017` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 133.0 MB | 2026-11-07 | `sha256:6c81ba54b8ad` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:6c81ba54b8adc6eaf0dc1b2f77f3ca4649c62634501a482f8a1742452d820a7f` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 140.0 MB | 2026-11-07 | `sha256:603c3f61189d` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:603c3f61189d5252ceb2c169b9e6920e44c2aa0e1c47a5fce3b29e89bd880260` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.403 | - | 0 | 0 | 0 | 951.0 MB | 2026-11-07 | `sha256:51845cad473c` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:51845cad473cfffeb6e35107a42c5f2078caf096c2c9d2127d10825a8da7c9f0` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.320 | - | 0 | 0 | 10 | 900.0 MB | 2026-11-07 | `sha256:18edf9c1261f` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:18edf9c1261fa6c196bc1b6d64c5e8ee7a1644b38e8b627884a012198cb07fb6` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.427 | - | 0 | 16 | 27 | 899.0 MB | 2026-11-07 | `sha256:8b002a6c9bd2` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:8b002a6c9bd247e73264bcba8cefcf850d1470ead0fc99e091c3829dbfa7cf02` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.33 | - | 1 | 34 | 217 | 194.0 MB | 2026-11-05 | `sha256:09f71a9459e2` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:09f71a9459e260fe4bab0528f0bd9bfbfed7c8330bed3aa5b665281760205bec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.33 | - | 1 | 34 | 217 | 218.0 MB | 2026-11-05 | `sha256:b14a41431d25` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b14a41431d252c640aae69b52320e512050fc53bafcb5012b97ccb9ca173c2d5` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.427 | - | 1 | 75 | 432 | 868.0 MB | 2026-11-05 | `sha256:341aa6e8bbac` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:341aa6e8bbaceb1194732328ca27c15bfa9c531687db83a66af3c0c1833f5534` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.14 | - | 0 | 0 | 0 | 124.0 MB | 2026-11-05 | `sha256:f2557fb81b1c` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:f2557fb81b1cfcfdfcde9d3bf247ed60b2cd38f00a88153f61e34737fab8ade7` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.14 | - | 0 | 0 | 1 | 162.0 MB | 2026-11-05 | `sha256:a9298ccc39c6` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:a9298ccc39c6258ae71228479e2f3f8d2ceb6639a9927262a4b0ff7d319a0e7e` |
| 3 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.33 | - | 0 | 2 | 23 | 193.0 MB | 2026-11-05 | `sha256:4c1e0075bc54` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4c1e0075bc541cb181872f11b24bd1ed4aa4618517011b1663269482c156d67a` |
| 4 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.22 | - | 0 | 2 | 23 | 198.0 MB | 2026-11-05 | `sha256:ed8d3f58736b` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:ed8d3f58736ba1ccd954aa3d13e582830fb799848e9b739c4d9b8fd0f028f688` |
| 5 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.14 | - | 0 | 2 | 23 | 203.0 MB | 2026-11-05 | `sha256:4f1f30cf0bb4` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:4f1f30cf0bb4ab0c8b2daa84613051683456f8df4a9bf7c812e20d74962992ac` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.33 | - | 0 | 2 | 23 | 217.0 MB | 2026-11-05 | `sha256:188854d2ed80` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:188854d2ed80c2768395bddde1af897ebe562509fbbbae8a3c1b504e42c65eae` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.22 | - | 0 | 2 | 23 | 223.0 MB | 2026-11-05 | `sha256:12e7bdfe76c8` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:12e7bdfe76c8210a5e47e365ff90330474e4a3504e122479dc4ded3fdf52de43` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.14 | - | 0 | 2 | 23 | 230.0 MB | 2026-11-05 | `sha256:1a48394b6f29` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:1a48394b6f29d3a32c79d9edecd4abda0ecbe8776d567dbfc4590ed6979c740c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.403 | - | 0 | 2 | 30 | 906.0 MB | 2026-11-05 | `sha256:212fe4cb7462` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:212fe4cb74627d5adf9817a4088906ae3edff1bd853bb27565d4b46731efe61e` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.320 | - | 0 | 2 | 40 | 855.0 MB | 2026-11-05 | `sha256:61c49333124a` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:61c49333124a241d3c14dee86467e8978aab73d4e569dcb1586dddeb9d844bd9` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-11-09 | `sha256:ad2f3b8eabd6` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:ad2f3b8eabd6110d3537e270d633cbfaa2108b98cd04a41df430512a9617f7bb` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 36 | 126 | 811.0 MB | 2026-11-09 | `sha256:a5d4b744e914` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:a5d4b744e9142c24601aef0d780820834a3d46d47e11a414dbef1a20f1d11fab` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33 | - | 0 | 0 | 0 | 324.0 MB | 2026-11-09 | `sha256:8126a9311a26` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:8126a9311a263aa51c99ab0d322f107e76d431315ea9ed76fed62d087df03c67` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 0 | 0 | 327.0 MB | 2026-11-09 | `sha256:808994a71485` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:808994a71485a8f89000976856f6639b30f404a8a3add236e3345075055e7b83` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 0 | 0 | 355.0 MB | 2026-11-09 | `sha256:13540e4c3910` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:13540e4c39107e09d2c024950edf1a79fba0dd85f51ccda64a41b337b75394a5` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 0 | 0 | 400.0 MB | 2026-11-09 | `sha256:39fc33df1e6b` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:39fc33df1e6b839a426c30ff17cb39e7602ef87f58453f3844de426242a9a9c6` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-11-09 | `sha256:a275348a31a5` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:a275348a31a51aeca34356e87238c447d5860ff0c45b0fb2a8e31af376f0f704` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-11-09 | `sha256:5d7e579e3d1c` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5d7e579e3d1caaff84b04f3732f904f928e3ebe45c07cbf2f3d577b1ffac139f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 0 | 71 | 435.0 MB | 2026-11-09 | `sha256:8d54dbeb745d` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:8d54dbeb745def9040ada99c8f5042d9b5d8108d76d77c36c319ee6e25fdaeb0` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 0 | 71 | 462.0 MB | 2026-11-09 | `sha256:72f64c1ce780` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:72f64c1ce780c7436c238822d9df2908d76b080910cb8dae7db3a0962df9307c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 0 | 71 | 507.0 MB | 2026-11-09 | `sha256:d04cfe302a59` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:d04cfe302a59691ae5db01f8d6a7b87d5a210c2785c490fff60c90916fb1c8d3` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.1 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-11-05 | `sha256:700f81898b1a` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:700f81898b1a66c186a9a64a3fd20f61942e1d8335887e78ce4b7a562e0b7ff8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.1 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-11-05 | `sha256:d7e4008fd2a2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:d7e4008fd2a2b87d144f316e6310c75a83a1128f930bd3da81fdb25613ba0e55` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.1 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-11-05 | `sha256:5b9ff8def1fa` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:5b9ff8def1fa8fad6c0ba57f955c8cd5b2118fa9a9962d463a75a233333d719c` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 3 | 62 | 271 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | - | 4 | 13 | 41 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 4 | 23 | 67 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 4 | 23 | 67 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 4 | 30 | 112 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 4 | 32 | 131 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 4 | 36 | 126 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-11-05 | `sha256:4ab51be1200a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:4ab51be1200ac07000790db1330e308765fb8d75ea3159b0c4930f0b197b2fbe` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-11-05 | `sha256:e66ce0c54a5a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:e66ce0c54a5a0a8e2c679cbb8c66ef40a95bee7faed7ef87bfd52c4661073b06` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-11-05 | `sha256:379135d4dbe6` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:379135d4dbe66b12825a54d4ed7aef21aa61227533ff1b9151d70f95f9266746` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-11-05 | `sha256:6142d570852e` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:6142d570852e09fb676788262568ed2a2bf028b4eabc9b8f512206b1a9da28c3` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.9 MB | 2026-11-05 | `sha256:bc4d9a4e0d33` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:bc4d9a4e0d3317999e235cb5895dde28bf766d25533ae8e6ba311bd6d253f996` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-11-05 | `sha256:141581e735b7` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:141581e735b73dfed58f739c7fc8f5bd37546cba8de22051550eb9dfeecdcc2e` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.8 MB | 2026-11-07 | `sha256:74c4d9520192` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:74c4d9520192bd327f7c06ff4f80ad12478cd58e6bd5cc1a191bf7f2ec87d727` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-11-05 | `sha256:370fb2c55e7f` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:370fb2c55e7f0474fed019623cce6a8db65a4d70dedfca9219fc69d6c7627f89` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-11-05 | `sha256:bfb630450e9c` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:bfb630450e9c3d7428d59f25ce6df95ba5634d664d96b06017aa56f87da5ec0e` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 0 | 14.2 MB | 2026-11-05 | `sha256:e9346023014b` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:e9346023014b406b9bc8be969b3aa4d39efa3a43b6abcc495f01be566acb10e9` |
