# Daily Recommended Images by Language

_Generated: 2026-06-25T03:12:17Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.11 | - | 0 | 0 | 0 | 820.0 MB | 2026-06-24 | `sha256:18812bd6f88c` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:18812bd6f88c6a3134ddc71a2a1387cc256c2168b5f81172d5179f1a7aa1ed67` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.4 | - | 0 | 0 | 0 | 853.0 MB | 2026-06-24 | `sha256:8dbce1c5f7fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:8dbce1c5f7fde5702cdbaa7a352c31477327ca2cec236c8e6c1cae3acb7f936a` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 0 | 488.0 MB | 2026-06-24 | `sha256:9e269bdc9e92` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:9e269bdc9e92f40cc6858cb0547d83c4a25e5b814311c02f5803e45a6d08bcd8` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 0 | 533.0 MB | 2026-06-24 | `sha256:9a7088ba6562` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:9a7088ba65624c26deb3091596bc0c7148da71c105ee0cb6ba4f3696f5193b5c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 2 | 0 | 4 | 323.0 MB | 2026-06-24 | `sha256:c77932bf7ad1` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:c77932bf7ad1cec669dd050607f7712345d2c96172874695cf2f1748b18aaf8b` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 2 | 0 | 4 | 326.0 MB | 2026-06-24 | `sha256:1bb99931dceb` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:1bb99931dceb172323b214b76f69b2954b7008171a9976a317412e8446d06c3a` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 2 | 0 | 4 | 354.0 MB | 2026-06-24 | `sha256:34e488081a89` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:34e488081a897fa2e2fa22615e96b5a36e79bcf06dbc06718a97bd7456e2a8b4` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 2 | 0 | 4 | 399.0 MB | 2026-06-24 | `sha256:f1cc79f8871a` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:f1cc79f8871adc49d198e93c2bbcf6aa227c0607e17c5d78b01a6aa38436f3a0` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 0 | 87 | 440.0 MB | 2026-06-24 | `sha256:de4ffcc8f089` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:de4ffcc8f089d4a3ef439f03bf288ff16a85fe9871877c6e2986ca8f639a5312` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 0 | 87 | 467.0 MB | 2026-06-24 | `sha256:77ce4caf0cb8` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:77ce4caf0cb8db87fd1f23ff7930f43eda93a00a97146b214d97c74d9aa8ad30` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 0 | 87 | 512.0 MB | 2026-06-24 | `sha256:d96e382dfeba` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:d96e382dfeba5cbd07cef4783a1edcef2d2cb45c908a388630fe1a8dee40c951` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.7 MB | 2026-06-19 | `sha256:b9c3f3717641` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b9c3f3717641372bef57dc9cbdcc191087587154cb9b01c3dae5b0c480a77c6d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-06-19 | `sha256:d188df941951` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:d188df941951658d215aa37c9e52d01b94821bd7e356f7bb8d7603db392a8b06` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 2 | 2 | 12 | 139.0 MB | 2026-06-19 | `sha256:84be0b597731` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:84be0b597731efe4134de278e1321e5bc0fa30eee24c3775b8ee92457be3603d` |

## Dotnet

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 2 | 0 | 4 | 107.0 MB | 2026-06-09 | `sha256:e7ecea6e4a4d` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:e7ecea6e4a4d5a5a3ccdda460543a6f76e543b0951e1a9fde4bb925aaaf2eccb` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 2 | 0 | 4 | 112.0 MB | 2026-06-09 | `sha256:30a8c40c4809` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:30a8c40c480930bdb3e9f01870f5ab742679982134807c780b4671af5a5c4959` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.28 | - | 2 | 0 | 4 | 126.0 MB | 2026-06-09 | `sha256:75121c43d3d2` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:75121c43d3d2bc7e8637b6de690b50f75a43f4a64cbe772da61802e64b7760c7` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 2 | 0 | 4 | 132.0 MB | 2026-06-09 | `sha256:b37b64067912` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:b37b640679120d7c7b071796fb65a67d203510b1d9d8cc763bca86d48e47a8f1` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 2 | 0 | 4 | 139.0 MB | 2026-06-09 | `sha256:2641266fe8db` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:2641266fe8dbd11ce14dc362a0e31aff5a97b44a3afa03497df5c08fba078ace` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.301 | - | 2 | 108 | 189 | 918.0 MB | 2026-06-09 | `sha256:b9ba3498cd0e` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:b9ba3498cd0ea65e665bbc921f91215309d007a5c2e407a79ba61aa627975b5e` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.315 | - | 2 | 108 | 199 | 879.0 MB | 2026-06-09 | `sha256:920ff2f97cc9` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:920ff2f97cc94d5e281f8cf0d6a4f36ce20981788ee183ec985a0d4d034d0528` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.28 | - | 3 | 6 | 144 | 193.0 MB | 2026-06-12 | `sha256:d73109ac3176` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d73109ac31761185b1b97af576c78182af9189495c35fa451ed82994c9af23bf` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.28 | - | 3 | 6 | 144 | 218.0 MB | 2026-06-12 | `sha256:93b366e510c6` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:93b366e510c6cd01cee608447014f7d349cb7ff8809fd0f554aa3772e8587b7e` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.422 | - | 9 | 37 | 343 | 850.0 MB | 2026-06-12 | `sha256:d80fdd84f7e1` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:d80fdd84f7e18eea12f8e45c52914f1353395009c95c41197178ea19944e6d48` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.28 | - | 0 | 0 | 43 | 199.0 MB | 2026-06-09 | `sha256:7b2a0bdb49b0` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7b2a0bdb49b0830a57c5be63f65d82e796e3acabe87c9b95a27a8c69c64704fb` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.17 | - | 0 | 0 | 43 | 205.0 MB | 2026-06-09 | `sha256:12fced73002a` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:12fced73002a68c58115becc6fd01f83e7378e854a489f01639b9c1eec5fda8c` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.9 | - | 0 | 0 | 43 | 209.0 MB | 2026-06-09 | `sha256:58318ab0733b` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:58318ab0733b63d3ac0d7609c46f2718244e623a176f45991ee01fad46fbf880` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.28 | - | 0 | 0 | 43 | 223.0 MB | 2026-06-09 | `sha256:bd44aa584869` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:bd44aa584869d2204094661820c793f34b5945a431ebbfae17d2eb8abd2a5fd7` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.17 | - | 0 | 0 | 43 | 229.0 MB | 2026-06-09 | `sha256:cca5379cc0ce` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:cca5379cc0cee235b3d825f4747be2879bc85517005afb38470c6d743bedda85` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.9 | - | 0 | 0 | 43 | 236.0 MB | 2026-06-09 | `sha256:ddcf70ad1ab9` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:ddcf70ad1ab963a4fcd41fbd722a6b660e404e87567cfbd46fd2809c21b02088` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.301 | - | 0 | 0 | 53 | 882.0 MB | 2026-06-09 | `sha256:548d93f8a18a` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:548d93f8a18a1acbe6cc127bc4f47281430d34a9e35c18afa80a8d6741c2adc3` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.315 | - | 0 | 0 | 63 | 843.0 MB | 2026-06-09 | `sha256:bdf453c44d37` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:bdf453c44d3729380bab181de1a9178d9920ad4f84273382d683ed2003e4e890` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.422 | - | 0 | 10 | 73 | 842.0 MB | 2026-06-09 | `sha256:d1e5dd229d0d` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:d1e5dd229d0d1d936df1273de6ff4b7114a39be127f0b1dfb0aa61fd922d3b7e` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | :24-nonroot | 0 | 5 | 10 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | :24 | 0 | 5 | 10 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 3 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | :24 | 2 | 2 | 12 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 11 | 71 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 2 | 25 | 60 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 2 | 25 | 60 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 3 | 33 | 97 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | :20-nonroot | 3 | 37 | 77 | 106.0 MB | 2026-03-04 | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | :20 | 3 | 37 | 77 | 106.0 MB | 2026-03-04 | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-06-19 | `sha256:83c9e52e0a0e` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:83c9e52e0a0ef97d9d87b8b81da6119f748dbeca4641ae2cb8b11552e2c8f35d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 2 | 0 | 4 | 34.1 MB | 2026-06-19 | `sha256:f8f5a9bb739a` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:f8f5a9bb739ad1ec347853144c9ed4ca2260e587082277bc6066fcd5cc9973e8` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 2 | 2 | 12 | 75.3 MB | 2026-06-19 | `sha256:1c56f09437df` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:1c56f09437dfc2910faad39abaaed336265d246cf183e2adb362d4cb3b881ab6` |
