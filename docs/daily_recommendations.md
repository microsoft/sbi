# Daily Recommended Images by Language

_Generated: 2026-07-13T03:11:46Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 107.0 MB | 2026-07-08 | `sha256:a9bc87ac6946` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:a9bc87ac6946ddc6cd6d72d1659595cc13565f9e6fcc357234ca7497e94e5242` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 112.0 MB | 2026-07-08 | `sha256:df6139122b81` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:df6139122b8161ba1b0a088998f0db33960dbb963965bd7b343e2b4ac0c3dfed` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.28 | - | 0 | 0 | 0 | 126.0 MB | 2026-07-08 | `sha256:cc04104c7b59` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:cc04104c7b59c439d19916da88808bb1da1337c528e8cd9ca7e29c378f7dcb4d` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.17 | - | 0 | 0 | 0 | 132.0 MB | 2026-07-08 | `sha256:27f7f3f9f2fd` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:27f7f3f9f2fd47cda0ff48ef13702c6c94225f0b31faba3058f3933cf8e939a8` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.9 | - | 0 | 0 | 0 | 139.0 MB | 2026-07-08 | `sha256:eea5b5657cbe` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:eea5b5657cbe7f71b18633a5dadb638d431a86349d014ea85a0f8cee0fdc59a2` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.301 | - | 0 | 0 | 0 | 918.0 MB | 2026-07-08 | `sha256:935ca8833439` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:935ca883343905e1a3453ca8eca13a06a16c570cf6d515451162ab1381fd370d` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.315 | - | 0 | 0 | 10 | 880.0 MB | 2026-07-08 | `sha256:d8735845ae45` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:d8735845ae45873714c341c5e3e7e8021623a3cbbec36ec9937d31c23b5c3288` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.28 | - | 3 | 16 | 159 | 193.0 MB | 2026-06-25 | `sha256:4cbc2530d585` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:4cbc2530d5855154199432020f232205aecd8be34be87bea866f3b235f853bcb` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.28 | - | 3 | 16 | 159 | 218.0 MB | 2026-06-25 | `sha256:e78fda31142e` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:e78fda31142e28746a6908e288e0d40346793f691ca99d8d150bcbe95c0ef035` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.422 | - | 9 | 65 | 363 | 850.0 MB | 2026-06-25 | `sha256:63ebabdcde24` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:63ebabdcde24cc8134304a6f50719cf1e22bb4e9f7148ac4ef967c7680187356` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.28 | - | 0 | 0 | 35 | 193.0 MB | 2026-06-25 | `sha256:1b809194e628` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:1b809194e6288a772512b6ed6d6f8780e476b3f52ee0591eefd7c3a88add7aaf` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.17 | - | 0 | 0 | 35 | 198.0 MB | 2026-06-25 | `sha256:0f7c41465ace` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:0f7c41465ace8ae345a18e167000bf7ae16e0ed87c398e1bc73cb24ac3c1296d` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.9 | - | 0 | 0 | 35 | 203.0 MB | 2026-06-25 | `sha256:6a40d375e9c8` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:6a40d375e9c8432fcf4adebae05d7e0a276e9a90dd01174df6709a090771bebc` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.28 | - | 0 | 0 | 35 | 217.0 MB | 2026-06-25 | `sha256:77526aac4254` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:77526aac4254c2b26764c2ab09700f6bdbcad2e2cb748e0d57bc01ebdced6515` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.17 | - | 0 | 0 | 35 | 223.0 MB | 2026-06-25 | `sha256:165ce688e1fc` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:165ce688e1fc73f2a624dfc8142531a7c594e155a7287b7ea12a43dc49f6601a` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.9 | - | 0 | 0 | 35 | 230.0 MB | 2026-06-25 | `sha256:7644f992230d` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:7644f992230d35cf230017189d4038c0ae0f7388b13f4f7ae1900a155bafb597` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.301 | - | 0 | 0 | 63 | 883.0 MB | 2026-06-25 | `sha256:ea8bde36c11b` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ea8bde36c11b6e7eec2656d0e59101d4462f6bd630730f2c8201ed0572b295d5` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.315 | - | 0 | 0 | 73 | 844.0 MB | 2026-06-25 | `sha256:b128a1e799e9` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:b128a1e799e98bceb26cc913b40d56da6e13cc7b0f665a5a81a0a4ded3f1f64d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.422 | - | 0 | 10 | 83 | 844.0 MB | 2026-06-25 | `sha256:08f003ba31fc` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:08f003ba31fca1aeed32d5a076f62dc3b592faa0a811b072aa245321dd09f0ab` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.12 | - | 0 | 0 | 0 | 809.0 MB | 2026-07-10 | `sha256:e3f9d4108b4e` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:e3f9d4108b4e5880e68378c5757cc103846a1f8dd9742b725a9d8a7f11b761e8` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.5 | - | 0 | 0 | 0 | 841.0 MB | 2026-07-10 | `sha256:1c77c1cbb5de` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:1c77c1cbb5de52db3f119fe2efe7a938e734c08196bbe3ad94b3bdadbab926f9` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | - | 0 | 0 | 0 | 477.0 MB | 2026-07-10 | `sha256:a452871f500e` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:a452871f500ec5ab153b71066c728955cacbcfbafcbc1ef4fd09de603dc8decf` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | - | 0 | 0 | 0 | 522.0 MB | 2026-07-10 | `sha256:1eaab2199b83` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:1eaab2199b830bbc28810b1b1b73171908ca54be95595d00a8fd182f849a26b0` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | - | 0 | 1 | 2 | 323.0 MB | 2026-07-10 | `sha256:ddaf950179ca` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:ddaf950179ca261c9f297579dad53a01017f8203700e9b7bad67759330d85586` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | - | 0 | 1 | 2 | 326.0 MB | 2026-07-10 | `sha256:fd755889c55d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:fd755889c55d6df28ec6f3429254fcba3ae1c5c036fb21e24b6f26980d2f0f00` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | - | 0 | 1 | 2 | 354.0 MB | 2026-07-10 | `sha256:45fcc92f3c56` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:45fcc92f3c5623d951cd0a84922a258d5135689c09ea4b9efe5ba852ee5c8b5a` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | - | 0 | 1 | 2 | 399.0 MB | 2026-07-10 | `sha256:4074fb32410d` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:4074fb32410d1468cd45356e4c61df53aaaf37170ec8829e6046a874b6361e54` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | - | 0 | 1 | 87 | 432.0 MB | 2026-07-10 | `sha256:9ddf45a14ac5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:9ddf45a14ac58263b6e2b791fbe0e75d460857d15e889caf3580ebe7709b02b9` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | - | 0 | 1 | 87 | 459.0 MB | 2026-07-08 | `sha256:6b5b1b54dc38` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:6b5b1b54dc38341381f612f1c7ed7ff340f89248495c879ca5ae6ddfcb787125` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | - | 0 | 1 | 87 | 503.0 MB | 2026-07-10 | `sha256:77432a8d689f` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:77432a8d689f7923d1ef863c31a610d503e3e4eaebe06e93bc1b13fc01561aaa` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.7 MB | 2026-07-06 | `sha256:13e3cc900da5` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:13e3cc900da5a3c7f878f2089f1414e9c8d6641e20b831d999aa0c3806b33f24` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-07-06 | `sha256:6053e35b8f94` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:6053e35b8f942fd0444a87934b7233445194748af29c688404651f3df20d53d9` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-07-06 | `sha256:b63ee4b9d88f` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:b63ee4b9d88fe6435fa9add7cb64c6c8cb5653a923ff4e84ece23186233f2da0` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | :24 | 0 | 3 | 9 | 196.0 MB | 2026-07-06 | `sha256:f72657da4619` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:f72657da461922ab2aa7b2589ac395bc4858302b2aa395410f0c7041a818f87d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | :24-nonroot | 0 | 5 | 19 | 156.0 MB | 2026-07-06 | `sha256:cbacebb4c60a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:cbacebb4c60abd5af3f992c26c3d5afb3c44616534c1a1aa8f2c932b129013ab` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | :24 | 0 | 5 | 19 | 156.0 MB | 2026-07-06 | `sha256:51b8cb44ddc0` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:51b8cb44ddc0e3e362f4dcccc9782700d05089d0d37e80ba8186f8041dfd1352` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 0 | 26 | 78 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 0 | 26 | 78 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 6 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 20 | 131 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 23 | 117 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 2 | 34 | 82 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 2 | 34 | 82 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | :20 | 3 | 50 | 187 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-07-06 | `sha256:576d9769c014` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:576d9769c0146cbf0cf7946bacf536c5758464c29eadfa03ef5090ae708e641f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-07-06 | `sha256:3dcd23ead303` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:3dcd23ead3035173915ddf670ea28d48cdcb6208025c8f31ce66e185053e8444` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.4 MB | 2026-07-06 | `sha256:0cdd0c6a200f` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:0cdd0c6a200fc2b5d6da711c34228126034bd428650b43dfb7e378214e6f2d32` |
