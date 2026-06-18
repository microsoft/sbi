# Daily Recommended Images by Language

_Generated: 2026-06-18T03:14:13Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 107.0 MB | 2026-06-09 | `sha256:e7ecea6e4a4d` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:e7ecea6e4a4d5a5a3ccdda460543a6f76e543b0951e1a9fde4bb925aaaf2eccb` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 112.0 MB | 2026-06-09 | `sha256:30a8c40c4809` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:30a8c40c480930bdb3e9f01870f5ab742679982134807c780b4671af5a5c4959` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.28 | - | 0 | 0 | 0 | 126.0 MB | 2026-06-09 | `sha256:75121c43d3d2` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:75121c43d3d2bc7e8637b6de690b50f75a43f4a64cbe772da61802e64b7760c7` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 132.0 MB | 2026-06-09 | `sha256:b37b64067912` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:b37b640679120d7c7b071796fb65a67d203510b1d9d8cc763bca86d48e47a8f1` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 139.0 MB | 2026-06-09 | `sha256:2641266fe8db` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:2641266fe8dbd11ce14dc362a0e31aff5a97b44a3afa03497df5c08fba078ace` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.301 | - | 0 | 108 | 143 | 918.0 MB | 2026-06-09 | `sha256:b9ba3498cd0e` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:b9ba3498cd0ea65e665bbc921f91215309d007a5c2e407a79ba61aa627975b5e` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.315 | - | 0 | 108 | 153 | 879.0 MB | 2026-06-09 | `sha256:920ff2f97cc9` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:920ff2f97cc94d5e281f8cf0d6a4f36ce20981788ee183ec985a0d4d034d0528` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.28 | - | 3 | 6 | 143 | 193.0 MB | 2026-06-12 | `sha256:d73109ac3176` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d73109ac31761185b1b97af576c78182af9189495c35fa451ed82994c9af23bf` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.28 | - | 3 | 6 | 143 | 218.0 MB | 2026-06-12 | `sha256:93b366e510c6` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:93b366e510c6cd01cee608447014f7d349cb7ff8809fd0f554aa3772e8587b7e` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.422 | - | 9 | 36 | 300 | 850.0 MB | 2026-06-12 | `sha256:d80fdd84f7e1` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:d80fdd84f7e18eea12f8e45c52914f1353395009c95c41197178ea19944e6d48` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.28 | - | 0 | 0 | 41 | 199.0 MB | 2026-06-09 | `sha256:7b2a0bdb49b0` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7b2a0bdb49b0830a57c5be63f65d82e796e3acabe87c9b95a27a8c69c64704fb` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.17 | - | 0 | 0 | 41 | 205.0 MB | 2026-06-09 | `sha256:12fced73002a` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:12fced73002a68c58115becc6fd01f83e7378e854a489f01639b9c1eec5fda8c` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.9 | - | 0 | 0 | 41 | 209.0 MB | 2026-06-09 | `sha256:58318ab0733b` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:58318ab0733b63d3ac0d7609c46f2718244e623a176f45991ee01fad46fbf880` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.28 | - | 0 | 0 | 41 | 223.0 MB | 2026-06-09 | `sha256:bd44aa584869` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:bd44aa584869d2204094661820c793f34b5945a431ebbfae17d2eb8abd2a5fd7` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.17 | - | 0 | 0 | 41 | 229.0 MB | 2026-06-09 | `sha256:cca5379cc0ce` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:cca5379cc0cee235b3d825f4747be2879bc85517005afb38470c6d743bedda85` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.9 | - | 0 | 0 | 41 | 236.0 MB | 2026-06-09 | `sha256:ddcf70ad1ab9` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:ddcf70ad1ab963a4fcd41fbd722a6b660e404e87567cfbd46fd2809c21b02088` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.301 | - | 0 | 0 | 45 | 882.0 MB | 2026-06-09 | `sha256:548d93f8a18a` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:548d93f8a18a1acbe6cc127bc4f47281430d34a9e35c18afa80a8d6741c2adc3` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.315 | - | 0 | 0 | 55 | 843.0 MB | 2026-06-09 | `sha256:bdf453c44d37` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:bdf453c44d3729380bab181de1a9178d9920ad4f84273382d683ed2003e4e890` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.422 | - | 0 | 10 | 65 | 842.0 MB | 2026-06-09 | `sha256:d1e5dd229d0d` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:d1e5dd229d0d1d936df1273de6ff4b7114a39be127f0b1dfb0aa61fd922d3b7e` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.11 | - | 0 | 0 | 0 | 812.0 MB | 2026-06-17 | `sha256:7b1514de0a13` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:7b1514de0a139e2bf9b78d797370cb5d163f689803aeac7c25df5e9e3706a97d` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.4 | - | 0 | 0 | 0 | 844.0 MB | 2026-06-17 | `sha256:9d27dd674cd2` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:9d27dd674cd2b53f01f9462d6da5da929b04d4ffdb1476ebca632f2fd608ba9c` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 0 | 0 | 0 | 323.0 MB | 2026-06-17 | `sha256:d4d401550e73` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:d4d401550e73569b9ed1ddeb8d9c32f21ec81205fa9cf1c5f8fdf5df0b998e32` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 0 | 0 | 0 | 326.0 MB | 2026-06-17 | `sha256:6447d6f60184` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:6447d6f601841ed821f925852abfc62ca7bfe83a1835c501ae18c8f31f92b997` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 0 | 0 | 0 | 354.0 MB | 2026-06-17 | `sha256:3656e32ffd03` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:3656e32ffd03e0ad9a33239963811aa4475c06eb9d840862facd9857a61b9f9c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 0 | 0 | 0 | 399.0 MB | 2026-06-17 | `sha256:153561ca57ce` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:153561ca57ceb3e1ea7e95ecefae84cc5437c541f093aa10fc4f0cc7a1b55dba` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 0 | 480.0 MB | 2026-06-17 | `sha256:c922a59db86e` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:c922a59db86e48700539f6db8e78333a4c8fe7f9ab34027a61bea597a189e684` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 0 | 523.0 MB | 2026-06-15 | `sha256:bd6af70fa8f2` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:bd6af70fa8f26d06434fb47154d9221536ed3c1c6a7b2c62f2bed7b58c87cd98` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 0 | 85 | 440.0 MB | 2026-06-17 | `sha256:6545b2200afe` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:6545b2200afe7d1e55ad85d5bd38f8751a76d4fcf0ed6e7ec18dad20cd6cc21a` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 0 | 85 | 467.0 MB | 2026-06-17 | `sha256:8397a0b98e28` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:8397a0b98e28326b8fc876ff5865108b3cdccc9752e0a67c9c2f1828c9027914` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 0 | 85 | 512.0 MB | 2026-06-17 | `sha256:43c1f1c31038` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:43c1f1c31038b1aeebb17817c9804bc93f03d158ba5289fa0cd42d84fabc64f3` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 2 | 2 | 83.7 MB | 2026-06-05 | `sha256:d83c8a8f7356` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:d83c8a8f73567d7f3888262685693b8be4bcc1f5616692bed33e304da1e6e17a` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 2 | 2 | 83.7 MB | 2026-06-05 | `sha256:7eef130986e3` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:7eef130986e3ab36c9bd5069b7260276343beab88b39927fdb523c5daf60012d` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 2 | 2 | 139.0 MB | 2026-06-05 | `sha256:086a668fc26d` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:086a668fc26db53e11452f977f303d6da11a57e6408add617f18eb9568eb0308` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | :24 | 0 | 2 | 2 | 193.0 MB | 2026-06-05 | `sha256:68425482d43e` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:68425482d43e84cd0887f56c9c2b0451750065035e8d82b23b9bf02184fa2c47` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | :24-nonroot | 0 | 5 | 10 | 153.0 MB | 2026-06-05 | `sha256:e4ae7b8125f5` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:e4ae7b8125f5963a0881df021c189e5cf154169e172aef4d29a964619954c586` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | :24 | 0 | 5 | 10 | 153.0 MB | 2026-06-05 | `sha256:b4a890da3286` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:b4a890da32868cc682ceca42e78b5b9a5121e0b46e743ccd0c4b619a058d09ce` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 0 | 11 | 61 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 0 | 25 | 56 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 0 | 25 | 56 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 1 | 33 | 87 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | :20-nonroot | 1 | 37 | 73 | 106.0 MB | 2026-03-04 | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | :20 | 1 | 37 | 73 | 106.0 MB | 2026-03-04 | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-06-02 | `sha256:15f9b83a828e` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:15f9b83a828eb6ae3a9057ff90ce52eae221997f14dd0221501c761682a08b3d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-06-02 | `sha256:91162180f572` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:91162180f5723e79e83dd65050de6ea4ae38cc4d4d132f287690cfc59b2c1d6a` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 2 | 2 | 75.3 MB | 2026-06-02 | `sha256:cd38424f36dd` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:cd38424f36dd2db09d0ccd4b0d6fa2203b3dcf7282fb4c06ea1663b23dcaf7df` |
