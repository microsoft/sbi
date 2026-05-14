# Daily Recommended Images by Language

_Generated: 2026-05-14T03:15:43Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 107.0 MB | 2026-05-13 | `sha256:83db6768c5b0` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:83db6768c5b041ea1c6a09645efa047c239bc3e71d06c7be4f27aded7647a5ec` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 112.0 MB | 2026-05-13 | `sha256:f533693eef1a` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:f533693eef1a7d816c8d31f574aa0b5ccc350a29ab460621d43c745e18974968` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.27 | 0 | 0 | 0 | 126.0 MB | 2026-05-13 | `sha256:fea200ebae29` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:fea200ebae291e45c540bd31c09ba92f0ce662f1326ad9720eeaf7db5f6224a9` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 132.0 MB | 2026-05-13 | `sha256:6c0097ff5181` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:6c0097ff51813189e494b301a3d208954c4e761e0d979738ad999dceaedbf48f` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 139.0 MB | 2026-05-13 | `sha256:ad10865ffdf5` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:ad10865ffdf5706ef7e8cdb3ccb25fe6f766e643b17d7d2b8ee0e1cc00c99efb` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.300 | 0 | 0 | 0 | 918.0 MB | 2026-05-13 | `sha256:28e3d18f3e17` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:28e3d18f3e172d554d7002ee1af97c0a24b31b26ae3b5dd8d2e40a9a6c5760b1` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.314 | 0 | 0 | 10 | 879.0 MB | 2026-05-13 | `sha256:403607a1996f` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:403607a1996f8c4009cbf536793f5b061ff12736028a6f95511155cf221b4ec1` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.27 | 3 | 12 | 148 | 193.0 MB | 2026-05-12 | `sha256:a9cb56dff4de` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a9cb56dff4de2b15faa1597ed12fbff1679c06b6e29a5678e08c0669216117d5` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.27 | 3 | 12 | 148 | 218.0 MB | 2026-05-12 | `sha256:93154e00cb22` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:93154e00cb227f2fad30724455ecd1e77f46afc9a0273590af652ac220664e54` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.421 | 4 | 28 | 279 | 849.0 MB | 2026-05-12 | `sha256:ebecce75fb42` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ebecce75fb42c4c14db108f88054d20a93b0b9e6dfbacad56c8f744c342cf9ef` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.27 | 0 | 0 | 22 | 193.0 MB | 2026-05-12 | `sha256:459f18113cb1` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:459f18113cb174c916f490ed4c848eed7cfadf8285efd04296356b6f3ccb571e` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.16 | 0 | 0 | 22 | 198.0 MB | 2026-05-12 | `sha256:0f5074ba51a1` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:0f5074ba51a1f1b725ff4e567bce6ac9818b1c0a32a21662c755e160bb980c2f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.8 | 0 | 0 | 22 | 203.0 MB | 2026-05-12 | `sha256:dcc1b4539569` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:dcc1b45395697ed27239d121eb8f3d5f2e2fd195257d1b8119cb3e9eb85ad44f` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.27 | 0 | 0 | 22 | 217.0 MB | 2026-05-12 | `sha256:cdd262e8545b` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:cdd262e8545b4313dad7e375ec0ce9703bc973d6384e275f7c7d42be9be93de7` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.16 | 0 | 0 | 22 | 223.0 MB | 2026-05-12 | `sha256:bbb1d16fe64e` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:bbb1d16fe64e2be82e1860573a38fa763ee82ecdf35b8ab72a599b635b674f31` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.8 | 0 | 0 | 22 | 230.0 MB | 2026-05-12 | `sha256:9b5222b0ff8e` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:9b5222b0ff8e9eb991a7c1a64b25f0f771d21ccc05dfa1c834f5668ffd9cd73f` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.300 | 0 | 0 | 26 | 875.0 MB | 2026-05-12 | `sha256:dc8430e6024d` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:dc8430e6024d454edadad1e160e1973be3cabbb7125998ef190d9e5c6adf7dbb` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.314 | 0 | 0 | 36 | 837.0 MB | 2026-05-12 | `sha256:4bd4c1855e91` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:4bd4c1855e9184db451b1de1bef106c6eb67400ba12c892616b8e74c52c22a1a` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.421 | 0 | 10 | 46 | 836.0 MB | 2026-05-12 | `sha256:16411e23fa66` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:16411e23fa6602381d54c1926a44d934bbe0c7fb680319c89bcd8f0e360a12e6` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.10 | 0 | 0 | 0 | 809.0 MB | 2026-05-13 | `sha256:cd1e15d2c132` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:cd1e15d2c132cfb79219818b8856645286f83eb2924a77def3a6edb4f26c9320` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.3 | 0 | 0 | 0 | 841.0 MB | 2026-05-13 | `sha256:2329aeebbce4` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:2329aeebbce4d25c81d006067e1da74a43552600a7b0e25475ace1b797b3ba42` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | 0 | 0 | 0 | 321.0 MB | 2026-05-13 | `sha256:def8837eae36` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:def8837eae36dd0752889e215e9e4bd5010cc52034653d6edb47798b8fa3415e` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | 0 | 0 | 0 | 324.0 MB | 2026-05-13 | `sha256:225f11c6c22e` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:225f11c6c22e98c453034dc6268e76a9d0971c8647a0b4d9b2061301edf91a94` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | 0 | 0 | 0 | 351.0 MB | 2026-05-13 | `sha256:5a3e5b92a865` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:5a3e5b92a865d5f7a2fb3805b206cc475a93f09ce4eaf5e2d70eecca1fcab063` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | 0 | 0 | 0 | 397.0 MB | 2026-05-13 | `sha256:3381ad31b33e` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:3381ad31b33ec69e44021bb1ac101fe26c68113d7112873767c3a897f9d91b76` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | 0 | 0 | 0 | 475.0 MB | 2026-05-13 | `sha256:1b3eec6395c6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:1b3eec6395c65c09a50c8182e172f6fa5a2627faad082bb5da442f2e26d7b390` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | 0 | 0 | 0 | 520.0 MB | 2026-05-13 | `sha256:17887dbf92c7` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:17887dbf92c7987a77ebe82990e818c72e1051edba99b4e6ef3d4e3192e357ad` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | 0 | 0 | 85 | 427.0 MB | 2026-05-13 | `sha256:a7cb810db877` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:a7cb810db877ec9a2ab30ea66a662fc5befc76e7e9cbd536a642342c8127e1de` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | 0 | 0 | 85 | 455.0 MB | 2026-05-13 | `sha256:814f63f317b1` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:814f63f317b16dcd1cc5eb7421930d81fececb2e3c5007d53876c3e0af601fba` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | 0 | 0 | 85 | 499.0 MB | 2026-05-13 | `sha256:95ee9145ee53` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:95ee9145ee53fb910ae9675c42ffcefa8e9c1c27849c8cabf8cbc7960816397c` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-10 | `sha256:46e2825b3f3a` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:46e2825b3f3ab19e1a8b769ff493bd25a5d7f95af64568796023269d5f93c1db` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-10 | `sha256:46e2825b3f3a` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:46e2825b3f3ab19e1a8b769ff493bd25a5d7f95af64568796023269d5f93c1db` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-10 | `sha256:0b9aa27e25bd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:0b9aa27e25bd76196fc2d9ddbe1eca4546b6144c6cc2479ebbf5c12879b1a784` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-10 | `sha256:e5f3c66a886e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:e5f3c66a886eb53c522848120b3eb5f9c5a53f0048fda95b80c7c3973a832294` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-10 | `sha256:0b9aa27e25bd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:0b9aa27e25bd76196fc2d9ddbe1eca4546b6144c6cc2479ebbf5c12879b1a784` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-10 | `sha256:e5f3c66a886e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:e5f3c66a886eb53c522848120b3eb5f9c5a53f0048fda95b80c7c3973a832294` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 8 | 50 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 10 | 50 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-10 | `sha256:4a478f3ce1d4` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:4a478f3ce1d464ec57c85e77f39801e4a97786520d586b9a5fa05df95ad5d95f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-10 | `sha256:08bc2b702e6a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:08bc2b702e6aa2dcc856cdbe239817c18cada79fdad7d835d4e3f59c5679eb76` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-10 | `sha256:4a478f3ce1d4` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:4a478f3ce1d464ec57c85e77f39801e4a97786520d586b9a5fa05df95ad5d95f` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-10 | `sha256:08bc2b702e6a` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:08bc2b702e6aa2dcc856cdbe239817c18cada79fdad7d835d4e3f59c5679eb76` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-10 | `sha256:388ce0129704` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:388ce0129704b2702d6575e10b8abf79fc78404fc5cc9903db7d039a5602351c` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-10 | `sha256:388ce0129704` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:388ce0129704b2702d6575e10b8abf79fc78404fc5cc9903db7d039a5602351c` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 475.0 MB | 2026-05-13 | `sha256:1b3eec6395c6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:1b3eec6395c65c09a50c8182e172f6fa5a2627faad082bb5da442f2e26d7b390` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 520.0 MB | 2026-05-13 | `sha256:17887dbf92c7` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:17887dbf92c7987a77ebe82990e818c72e1051edba99b4e6ef3d4e3192e357ad` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | 2026-05-13 | `sha256:cd1e15d2c132` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:cd1e15d2c132cfb79219818b8856645286f83eb2924a77def3a6edb4f26c9320` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | 2026-05-13 | `sha256:2329aeebbce4` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:2329aeebbce4d25c81d006067e1da74a43552600a7b0e25475ace1b797b3ba42` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | 2026-05-10 | `sha256:31fa4e799021` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:31fa4e799021201e2d7b25b389037c9b8f750c102308c47954be8eb41dfa4a1c` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | 2026-05-10 | `sha256:9b9c3c4ad09e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:9b9c3c4ad09e26ef062bfb422e3af437ce7185d01d82406120750860ba3cdc0b` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | 2026-05-10 | `sha256:923029007961` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:9230290079612a2a8d30a0aa8b8f097c3194e787d09987df0c3e6b8f6271ee27` |
