# Daily Recommended Images by Language

_Generated: 2026-08-12T03:01:22Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 107.0 MB | 2026-08-10 | `sha256:dec201d3be75` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:dec201d3be75681fc79661f0981ec47623e4ce7d1bff5c3d4d29120502205b15` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 112.0 MB | 2026-08-10 | `sha256:83fbb26e97c9` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:83fbb26e97c98799a3a059f2ca05381a64c9b3799dd16205320fa079145294ea` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.30 | - | 0 | 0 | 0 | 126.0 MB | 2026-08-10 | `sha256:3c0d5f8d1f5a` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:3c0d5f8d1f5aed7e3c24ddc0bc47e5a11dde5560041b70c3f93d8e0fba048990` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 132.0 MB | 2026-08-10 | `sha256:f01ebb6d2e3a` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:f01ebb6d2e3acce39ff057bd2408ca5eae9fa536bd7d92e1ffb25eb032ef3d81` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 139.0 MB | 2026-08-10 | `sha256:83986d537223` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:83986d537223c9b31e22f5b9e2b25d99fe7b039877f376721b08bf616c73c32c` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.317 | - | 0 | 1 | 12 | 881.0 MB | 2026-08-10 | `sha256:4550fda34470` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:4550fda344705dcaf19116031eff5125fff3419c34879a211ceb1c125ddabb4b` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.400 | - | 0 | 6 | 7 | 931.0 MB | 2026-08-10 | `sha256:74fdeba262ce` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:74fdeba262ceb9d2daecfdf4a20ec67e9701cd4fd4bb32f82e0ebeb2ff068902` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 5 | 17 | 172 | 193.0 MB | 2026-08-10 | `sha256:a27922704a14` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a27922704a14d6d1ce103697aa3efa885902db33651b41a84ad1594fdf5be66e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 5 | 17 | 172 | 218.0 MB | 2026-08-10 | `sha256:b0beb9cc1dee` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b0beb9cc1dee1c1b0749796110d4734292071b814207ad0d4f40611f7db04f7b` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 17 | 64 | 375 | 850.0 MB | 2026-08-10 | `sha256:306301580fca` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:306301580fcaa5b445180e759db59309979002d1000669cb4cf58a567d0014bc` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 19 | 199.0 MB | 2026-08-10 | `sha256:6a585d286cb7` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:6a585d286cb758eeeba86166bb28beaf1da3de35ed0f47f5850ff9ee750d0f43` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 19 | 205.0 MB | 2026-08-10 | `sha256:8e8a639a9806` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:8e8a639a98063440299b397e36e0520d24457483badd4f34c133bcb5ece6e162` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 19 | 209.0 MB | 2026-08-10 | `sha256:acad02eb5c4f` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:acad02eb5c4fbf57d15296f9c08d56cd4036e915bdae5b4dd48a06523d452617` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 19 | 223.0 MB | 2026-08-10 | `sha256:2b27185ab3f3` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:2b27185ab3f3fd7ee77d9c9515f333758ff34b46f2ab507b36535e8693ae3e5c` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 19 | 229.0 MB | 2026-08-10 | `sha256:546d3a96069b` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:546d3a96069bd7442ec3e9d26b245dfca58fc0824a3dc277312db0fd99a412fc` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 19 | 236.0 MB | 2026-08-10 | `sha256:207cc5149677` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:207cc51496778557731c81ff670333d8ade4a4fec22768fd1be8e78474a84ecf` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 33 | 844.0 MB | 2026-08-10 | `sha256:840c88158a49` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:840c88158a49d65ab59451369988a0db6b420ab3ea7948a5e261c56a4ab3404d` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 28 | 894.0 MB | 2026-08-10 | `sha256:e1fc6e423f54` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1fc6e423f543119c406d24e2e687d67c569f18f04a37a8b0005d80ad0dcee80` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 44 | 842.0 MB | 2026-08-10 | `sha256:f7b81c81fde9` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:f7b81c81fde99517f5ac7ff48188500782f9bff210bf3e8648a6ff75f651c43f` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.12 | - | 0 | 0 | 0 | 810.0 MB | 2026-08-10 | `sha256:86f70d8ddc3f` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:86f70d8ddc3fe9e0eeb1be30c7a4f73040261350b267c77505e641fbb60db239` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.5 | - | 0 | 0 | 0 | 843.0 MB | 2026-08-10 | `sha256:d0ad9fcd7803` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:d0ad9fcd7803caf36f5411356d34a44989f8688ac5bebef76d27fb28235d2815` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32 | - | 0 | 0 | 0 | 323.0 MB | 2026-08-10 | `sha256:e8ed4416fc65` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:e8ed4416fc65ff9dfcfe601552e220694b57a09896d92077f78df480741575eb` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20 | - | 0 | 0 | 0 | 326.0 MB | 2026-08-10 | `sha256:1e8f3bf3e78e` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:1e8f3bf3e78eb468f81104306f6b907ba73b2f509bdddd73975764e3012191ce` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12 | - | 0 | 0 | 0 | 354.0 MB | 2026-08-10 | `sha256:abf1de62e423` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:abf1de62e4230baef7771133637ae03760e230ee7ac47c2685e19dac43182abb` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4 | - | 0 | 0 | 0 | 399.0 MB | 2026-08-10 | `sha256:e8864d0c0e7a` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:e8864d0c0e7ab7996cdb8dcfc5b07924362c5ddfeb5b599bbc96fb476ccd662b` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-10 | `sha256:9c7ab5898718` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:9c7ab5898718fc157cb4826f90c1ed483f3f66aed0708c8773aeb0374e8b312c` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4 | - | 0 | 0 | 0 | 523.0 MB | 2026-08-10 | `sha256:f7fd5134d4fe` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:f7fd5134d4fee8c5d1ab8af1e39ee797f5ad3c45ecd9362f436ed8889a91087d` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20 | - | 0 | 0 | 79 | 429.0 MB | 2026-08-10 | `sha256:de403cfd76e3` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:de403cfd76e3075b24dbbb14aee004a48965d172ecde36d944b076ebb0e6238b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12 | - | 0 | 0 | 79 | 457.0 MB | 2026-08-10 | `sha256:a5b01acc7eec` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:a5b01acc7eeca9a569d3fb3a4bc23fb89daf580a97fb6479505bc0742485d84c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4 | - | 0 | 0 | 79 | 501.0 MB | 2026-08-10 | `sha256:0813705cc3ce` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:0813705cc3ce323d342975941a94fd1dc3f27104b86e58ef2fcf1e456ce470ba` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 6 | 24 | 83.7 MB | 2026-07-31 | `sha256:81accc5dfe9a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:81accc5dfe9a015fd026bda4856d3073cdc0928e54ceb8ce1df85c306d5a5ca7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 6 | 24 | 83.7 MB | 2026-07-31 | `sha256:8b486e043e64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:8b486e043e646e69a9bfc25a404909f1036b8af386a9e1d410f90936d923fb23` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 7 | 26 | 139.0 MB | 2026-07-31 | `sha256:66097f09fcca` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:66097f09fcca9f87dfa6481938c1fa4a8b81ea5d6319f1129c610367f1436c06` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 1 | 2 | 197.0 MB | 2026-07-31 | `sha256:6523505968bf` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:6523505968bf37b0f97975cedfd202a3dbaee44ed653d17890d36deead3828cc` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 8 | 21 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 6 | 19 | 157.0 MB | 2026-07-31 | `sha256:ad7995197668` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ad7995197668d617eb71bb198d90d1d95049f6f43028b27c0a694cddb430c073` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 6 | 19 | 157.0 MB | 2026-07-31 | `sha256:8652f9953ce3` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:8652f9953ce393d2018727371948afda3fc085aaceb07b2ade8f8b1c112e1b1c` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 38 | 162 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 44 | 163 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-07-06 | `sha256:576d9769c014` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:576d9769c0146cbf0cf7946bacf536c5758464c29eadfa03ef5090ae708e641f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-07-22 | `sha256:178f25fadf46` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:178f25fadf466549d31e234b3091bf815161159f2f2bc98720bbf39f7368aff4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 1 | 2 | 75.4 MB | 2026-07-29 | `sha256:4ecd6b297db8` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:4ecd6b297db85c54ec2df07145a28536c3655a3e98e54eb2364189bc4e6eac23` |
