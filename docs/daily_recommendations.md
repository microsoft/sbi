# Daily Recommended Images by Language

_Generated: 2026-07-23T03:10:53Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.18 | - | 0 | 0 | 0 | 107.0 MB | 2026-07-14 | `sha256:8f9f04927791` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:8f9f049277916be27b72704649a5da6a82b0a3527c61f275658dfeea4ab81311` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.10 | - | 0 | 0 | 0 | 112.0 MB | 2026-07-14 | `sha256:c715e19ee81a` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:c715e19ee81a6482bfb667f205a01a23ec5c9c1b6834fbc242d0f311da478f10` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.29 | - | 0 | 0 | 0 | 126.0 MB | 2026-07-14 | `sha256:7d4e87003e81` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:7d4e87003e810d81c016f3bb2a4e540bf0f34ad7a385104cf3b141d7e1233663` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.18 | - | 0 | 0 | 0 | 132.0 MB | 2026-07-14 | `sha256:7b12beadd29e` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:7b12beadd29e5c3a699f04cd509750c3ce00cd7c6310baecc5834aa18602fc37` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.10 | - | 0 | 0 | 0 | 139.0 MB | 2026-07-14 | `sha256:814f7b303e77` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:814f7b303e771fe8cbdf29233cb053728eaff8fe9efae95c81220dba44e81237` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.302 | - | 35 | 40 | 110 | 919.0 MB | 2026-07-14 | `sha256:675817acabaa` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:675817acabaae7140642e228aa4c795dfa30c960c4daa3776b2e749305ffd3bb` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.316 | - | 35 | 40 | 120 | 881.0 MB | 2026-07-14 | `sha256:ff06e00ec2cc` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ff06e00ec2cc3b87b7f383dda323c92f52ac6e284e8c04ae4930f04a035d44e0` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.29 | - | 5 | 17 | 162 | 193.0 MB | 2026-07-14 | `sha256:f7519e953be7` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:f7519e953be71b6a8bc8987530b8c453ab13cfd3db2dff5c9ae7babc0e1740ae` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.29 | - | 5 | 17 | 162 | 218.0 MB | 2026-07-14 | `sha256:9822a6201c64` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:9822a6201c64c9f38b8a190832a7ae21cd30150d509b3ea170fd83b4f4e6166a` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.423 | - | 17 | 72 | 364 | 850.0 MB | 2026-07-14 | `sha256:89ce6291bde9` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:89ce6291bde9acdf59594e79fb8277c6d84c46e4b1f5bf126a4f18766e4bd597` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.29 | - | 0 | 0 | 35 | 193.0 MB | 2026-07-14 | `sha256:0f038d362c97` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:0f038d362c97c1d34ce2481891e7dadfcbb67a014fed7e010ba2ef5bbb9c5820` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.18 | - | 0 | 0 | 35 | 198.0 MB | 2026-07-14 | `sha256:156e9fd13513` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:156e9fd1351359ac2b8cd3e05676de78bfb1a8937f9b221b50f9ccf6984c7093` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.10 | - | 0 | 0 | 35 | 203.0 MB | 2026-07-14 | `sha256:ed5d539b2784` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:ed5d539b27842d656a06a5984dbcb5114d3e885fbada612a49a5a7c3c3a44e1c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.29 | - | 0 | 0 | 35 | 217.0 MB | 2026-07-14 | `sha256:15f696229ddd` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:15f696229dddb7e68aae3fa9551f1af9093d7cdda1803834fdf722013383014c` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.18 | - | 0 | 0 | 35 | 223.0 MB | 2026-07-14 | `sha256:5ad129827db5` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:5ad129827db5bf710906a43db07f4f1e2209bd2a8e3403ffea1ffe6961a82dc1` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.10 | - | 0 | 0 | 35 | 230.0 MB | 2026-07-14 | `sha256:1fa23fc4872d` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:1fa23fc4872d95fd71c2833ebe65d7e84a43b2d51a31d119516852f13d9505a7` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.302 | - | 0 | 5 | 59 | 883.0 MB | 2026-07-14 | `sha256:ed034a8bf0b2` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ed034a8bf0b24ded0cbbac07e17825d8e9ebfe21e308191d0f7421eaf5ad4664` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.316 | - | 0 | 5 | 69 | 845.0 MB | 2026-07-14 | `sha256:72b2c1fba104` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:72b2c1fba104eed0765e76c66256dd57b8b00c5e7c7fd16ad3eb254ad18db3fc` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.423 | - | 0 | 16 | 80 | 844.0 MB | 2026-07-14 | `sha256:283164eecee1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:283164eecee1fc80a590410d3c56207af455bd51b8bf3cbf2ded3592b9014b29` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 0 | 0 | 0 | 323.0 MB | 2026-07-22 | `sha256:6a39278d840f` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6a39278d840fb711cb804da516884305f7208cca0860b71b2577c1c7d12a2e95` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 0 | 0 | 0 | 326.0 MB | 2026-07-22 | `sha256:33fb4d4baf0c` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:33fb4d4baf0cc71ee1adfeb72c87b0a5dffcc1675ea3d105a42bf44cf8616d7f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 0 | 0 | 0 | 354.0 MB | 2026-07-22 | `sha256:e3cbf3c93b95` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:e3cbf3c93b95541811232c8824123130646ca34b847c66d33dc064aea01e9f90` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 0 | 0 | 0 | 399.0 MB | 2026-07-22 | `sha256:80acccf56744` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:80acccf567442efec792ffc22934c1974d8ec39efa8fe9d38a8fd8c49046c7a7` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 0 | 477.0 MB | 2026-07-22 | `sha256:25b757fbed33` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:25b757fbed33c8b3af4c49dc0f8cb514ca9dac31df8bbe3144fac798873a2b04` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 0 | 522.0 MB | 2026-07-22 | `sha256:0b2a4c71f2c8` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0b2a4c71f2c8821e07c831857c070963e5411e59a248591500c997be908c220c` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 1 | 101 | 432.0 MB | 2026-07-10 | `sha256:9ddf45a14ac5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:9ddf45a14ac58263b6e2b791fbe0e75d460857d15e889caf3580ebe7709b02b9` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 1 | 101 | 459.0 MB | 2026-07-08 | `sha256:6b5b1b54dc38` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:6b5b1b54dc38341381f612f1c7ed7ff340f89248495c879ca5ae6ddfcb787125` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 1 | 101 | 503.0 MB | 2026-07-10 | `sha256:77432a8d689f` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:77432a8d689f7923d1ef863c31a610d503e3e4eaebe06e93bc1b13fc01561aaa` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | :24 | 0 | 0 | 0 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | :24-nonroot | 1 | 4 | 15 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | :24 | 1 | 4 | 15 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 28 | 83 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 28 | 83 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 6 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 28 | 147 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 31 | 133 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 3 | 36 | 87 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 3 | 36 | 87 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 3 | 58 | 203 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.7 MB | 2026-07-22 | `sha256:3e8b0c32cf30` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:3e8b0c32cf3093429b02c5ccc62dd08820ba94f77206c369c54e563c7646eb1c` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-07-22 | `sha256:37cf91ef0121` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:37cf91ef0121b0d03ab72894a7ee02c67ddc7b71cda987b2a800ed085bc8f90c` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-07-22 | `sha256:ba723369765c` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:ba723369765c5532f606ccde2877f7f76028350d9edaa19022abbd0912cff7ff` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.12 | - | 35 | 0 | 35 | 809.0 MB | 2026-07-22 | `sha256:039726bb73b2` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:039726bb73b24b07b075fcf934a868962674dcfa2e7da36185563521615e34f2` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.5 | - | 35 | 0 | 35 | 841.0 MB | 2026-07-22 | `sha256:f0363c90b1a4` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:f0363c90b1a42a06fb72eced44dd0330003ae1d9a855c74f209363a9ba9a684d` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-07-06 | `sha256:576d9769c014` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:576d9769c0146cbf0cf7946bacf536c5758464c29eadfa03ef5090ae708e641f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-07-22 | `sha256:178f25fadf46` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:178f25fadf466549d31e234b3091bf815161159f2f2bc98720bbf39f7368aff4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.4 MB | 2026-07-22 | `sha256:a30e18dd24a8` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a30e18dd24a8080ee0b72d0f998a688e99380678a407bdd7c3a0ac7417b15eb3` |
