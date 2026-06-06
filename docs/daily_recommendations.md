# Daily Recommended Images by Language

_Generated: 2026-06-06T03:13:42Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.16 | - | 0 | 0 | 0 | 107.0 MB | 2026-06-05 | `sha256:b5cfd3688a5f` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:b5cfd3688a5f40042b765e6e4127da9ee61d820f98f429ce56efc5f4764cba09` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.8 | - | 0 | 0 | 0 | 112.0 MB | 2026-06-05 | `sha256:3b63c797f9cb` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:3b63c797f9cbb1d46f547f09695b53f0f14d30e25563c6086d3e91c0729b12bb` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.27 | - | 0 | 0 | 0 | 126.0 MB | 2026-06-05 | `sha256:4ab637e6ea0e` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:4ab637e6ea0e888a5f4bc970467a98257c878c80380f4914d0e6766e5f44b9ea` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.16 | - | 0 | 0 | 0 | 132.0 MB | 2026-06-05 | `sha256:9013290cb672` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:9013290cb6727d0daa1c4a445fc25d49cca64a45189225403bd65bcba44c8e1a` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.8 | - | 0 | 0 | 0 | 139.0 MB | 2026-06-05 | `sha256:42643ee4e0d4` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:42643ee4e0d488a4177fd834085939221ce54d3ecaafccdb429cdc9495b03494` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.300 | - | 0 | 0 | 0 | 918.0 MB | 2026-06-05 | `sha256:b254b994500b` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:b254b994500b174ffcb1522aefb993961996f360fc8fcf53a579cb4261fe82dc` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.314 | - | 0 | 0 | 10 | 879.0 MB | 2026-06-05 | `sha256:c7629fbdf4a0` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c7629fbdf4a0002d32b5261bd05e6bb83d7383f6bcf97b467cba56e710ade8f3` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.27 | - | 5 | 9 | 133 | 193.0 MB | 2026-05-20 | `sha256:6810933f14d0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:6810933f14d0b8a0eb193ab60e3a6429f8d263a7af891fa9ebd5f0eb6703febb` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.27 | - | 5 | 9 | 133 | 218.0 MB | 2026-05-20 | `sha256:19be23fe71e8` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:19be23fe71e885186495d8da1f2d417e553daaa99e16e5148c5ca3ee2f062512` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.421 | - | 11 | 39 | 291 | 849.0 MB | 2026-05-20 | `sha256:fc69dc5e0c97` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:fc69dc5e0c9789adaac5c8efce71ead4d016a51318667c4f26ce93574b1b9403` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.27 | - | 0 | 0 | 40 | 193.0 MB | 2026-05-20 | `sha256:7bb91a5cf371` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7bb91a5cf371059c9ebbb92142db6e6f7024a75da4cc306b181de0b57f3058e3` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.16 | - | 0 | 0 | 40 | 198.0 MB | 2026-05-20 | `sha256:b22e7cc84d54` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:b22e7cc84d549540db21a778a28012ba1527689ae24a7a06851d176328e24ed7` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.8 | - | 0 | 0 | 40 | 203.0 MB | 2026-05-20 | `sha256:d399699ebc8a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:d399699ebc8a27ab34665707ae6dc8f77ae478bd319444841a33a2b3840c5c9a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.27 | - | 0 | 0 | 40 | 217.0 MB | 2026-05-20 | `sha256:052a40facc3d` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:052a40facc3d32caa9920381df76b20628382409f021bf5d9f38dd67a6f936fa` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.16 | - | 0 | 0 | 40 | 223.0 MB | 2026-05-20 | `sha256:53683436482e` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:53683436482e444f31c9db6c34921f408a84e7ebcef41e2d3862c7c10051076a` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.8 | - | 0 | 0 | 40 | 230.0 MB | 2026-05-20 | `sha256:8c0b6857eab7` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8c0b6857eab7b2aa57884c839bf4678414606bd7d17370f18a842ac5cf414711` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.300 | - | 0 | 0 | 44 | 875.0 MB | 2026-05-20 | `sha256:c07906393326` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:c0790639332692a0d56cdd81ed581cfd24d040d9839764c138994866df89a3b6` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.314 | - | 0 | 0 | 54 | 837.0 MB | 2026-05-20 | `sha256:fda2105a0319` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:fda2105a0319b5ae5d7d51bd540ab0f7fc548ee6ea48cbc95c2b63b62ecad3a3` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.421 | - | 0 | 10 | 64 | 836.0 MB | 2026-05-20 | `sha256:6b0b7f73dc7c` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:6b0b7f73dc7cce85fe9eaf7cfcfd1dc109accc5b3782c8cba006fbe036da424e` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.11 | - | 0 | 0 | 0 | 809.0 MB | 2026-06-03 | `sha256:91664b8bf244` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:91664b8bf244b7abb63bed58abedbd64907ebad4dd15c3486d879b3a23303fa9` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.4 | - | 0 | 0 | 0 | 841.0 MB | 2026-06-03 | `sha256:5f95bf70f4c4` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:5f95bf70f4c437de4a6ba1f72ef1532f69fca392c131fff2172993cabbea359c` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 0 | 0 | 0 | 323.0 MB | 2026-06-05 | `sha256:7fc7401b1f2e` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:7fc7401b1f2e8ae23212160e217eb21c06c814d6afe840ba41ce4632775d959c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 0 | 0 | 0 | 326.0 MB | 2026-06-05 | `sha256:f11f36fdae9c` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:f11f36fdae9cc874acef40c818a4bb5078b80cecab8da728531ec0dba5f506e3` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 0 | 0 | 0 | 354.0 MB | 2026-06-05 | `sha256:9a435dd72d96` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:9a435dd72d967fb1d60ad798c5723c12b91d021a8e31536aa17744837251f962` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 0 | 0 | 0 | 399.0 MB | 2026-06-05 | `sha256:f277b436f45e` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:f277b436f45ea66c86e7e3ac33ed0fcd9807f08425588f05e2a83deab66de6aa` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 0 | 477.0 MB | 2026-06-05 | `sha256:0424dd358a2c` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:0424dd358a2c0579085ed1a2c74a0ffa532ffa870cc8247bd3d68bbd77ce2346` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 0 | 522.0 MB | 2026-06-05 | `sha256:e7505afb11f5` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e7505afb11f5d6c9544943a2990a1d36e7fecf41933cc12eec0e7fd88a314f3c` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 0 | 88 | 433.0 MB | 2026-06-05 | `sha256:3b68bb925af4` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:3b68bb925af4efb4da7e0788c8e01cf86e3258f1f5986783fe441f3595c91329` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 0 | 88 | 460.0 MB | 2026-06-05 | `sha256:ae2cc3730ca0` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:ae2cc3730ca05ac6edc4cab655c41cbb73fc4c5f239e921977a5f0e0f61a49eb` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 0 | 88 | 505.0 MB | 2026-06-05 | `sha256:f690875c6872` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:f690875c6872dc7f8c40eb1eab3aac0aabf84d69c1a9010d9752137f4e703264` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | :24 | 0 | 0 | 0 | 193.0 MB | 2026-06-05 | `sha256:68425482d43e` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:68425482d43e84cd0887f56c9c2b0451750065035e8d82b23b9bf02184fa2c47` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | :24-nonroot | 0 | 5 | 9 | 153.0 MB | 2026-06-05 | `sha256:e4ae7b8125f5` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:e4ae7b8125f5963a0881df021c189e5cf154169e172aef4d29a964619954c586` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | :24 | 0 | 5 | 9 | 153.0 MB | 2026-06-05 | `sha256:b4a890da3286` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:b4a890da32868cc682ceca42e78b5b9a5121e0b46e743ccd0c4b619a058d09ce` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 0 | 9 | 51 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 0 | 25 | 55 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 0 | 25 | 55 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 1 | 11 | 51 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | :20-nonroot | 1 | 17 | 46 | 106.0 MB | 2026-03-04 | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | :20 | 1 | 17 | 46 | 106.0 MB | 2026-03-04 | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-06-05 | `sha256:7eef130986e3` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:7eef130986e3ab36c9bd5069b7260276343beab88b39927fdb523c5daf60012d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | - | 0 | 0 | 0 | 83.7 MB | 2026-06-05 | `sha256:d83c8a8f7356` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:d83c8a8f73567d7f3888262685693b8be4bcc1f5616692bed33e304da1e6e17a` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-06-05 | `sha256:086a668fc26d` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:086a668fc26db53e11452f977f303d6da11a57e6408add617f18eb9568eb0308` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | - | 0 | 2 | 2 | 83.7 MB | 2026-05-17 | `sha256:c0279d3b8bdd` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:c0279d3b8bdddbe189585e9b084b234eed6285ea60dc2c5dec42ba6c8f3cf10b` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-06-02 | `sha256:15f9b83a828e` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:15f9b83a828eb6ae3a9057ff90ce52eae221997f14dd0221501c761682a08b3d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-06-02 | `sha256:91162180f572` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:91162180f5723e79e83dd65050de6ea4ae38cc4d4d132f287690cfc59b2c1d6a` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.3 MB | 2026-06-02 | `sha256:cd38424f36dd` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:cd38424f36dd2db09d0ccd4b0d6fa2203b3dcf7282fb4c06ea1663b23dcaf7df` |
