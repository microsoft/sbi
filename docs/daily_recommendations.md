# Daily Recommended Images by Language

_Generated: 2026-05-13T03:14:07Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 107.0 MB | 2026-05-12 | `sha256:248e6cb9e2a6` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:248e6cb9e2a63392f7536ff2347cdc966f4a2f437e150ba97740fb1fb67de7a9` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 112.0 MB | 2026-05-12 | `sha256:2a73d0c572d5` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:2a73d0c572d5c35dcf56e871adb90405650a0d17273c5325b99c4b542320e43e` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.27 | 0 | 0 | 0 | 126.0 MB | 2026-05-12 | `sha256:5effe1628853` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:5effe162885379dbbf940f3ae72573799b13d8538b4fe1efaa7b79e20942440a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 132.0 MB | 2026-05-12 | `sha256:5795f946005b` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:5795f946005b72a5c1f2f59494eebfc993f433f0a4621f9fc6f2f2818705c363` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 139.0 MB | 2026-05-12 | `sha256:f0d5456cfeda` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:f0d5456cfeda617f0309401be1817a98d34998447b59f46869142dd378e19cd9` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.300 | 0 | 0 | 0 | 918.0 MB | 2026-05-12 | `sha256:9295edd1d3da` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:9295edd1d3dab97acb5a526c2129b5549374f95b2c8734674ec90b04364c8260` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.314 | 0 | 0 | 10 | 879.0 MB | 2026-05-12 | `sha256:3dac2f5cea77` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:3dac2f5cea7792ee60ff1c2e4ba7137b1ddfe03080f1bb3306d578b6a407357f` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.27 | 2 | 13 | 148 | 193.0 MB | 2026-05-12 | `sha256:a9cb56dff4de` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a9cb56dff4de2b15faa1597ed12fbff1679c06b6e29a5678e08c0669216117d5` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.27 | 2 | 13 | 148 | 218.0 MB | 2026-05-12 | `sha256:93154e00cb22` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:93154e00cb227f2fad30724455ecd1e77f46afc9a0273590af652ac220664e54` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.421 | 3 | 26 | 279 | 849.0 MB | 2026-05-12 | `sha256:ebecce75fb42` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ebecce75fb42c4c14db108f88054d20a93b0b9e6dfbacad56c8f744c342cf9ef` |

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
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.10 | 0 | 0 | 0 | 809.0 MB | 2026-05-11 | `sha256:31211bf7d7fb` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:31211bf7d7fbddfe501802c7e6c91e7605a7e2d90ff903a10d6da5dd30e181e9` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.3 | 0 | 0 | 0 | 841.0 MB | 2026-05-11 | `sha256:6b48a8fb2ae3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:6b48a8fb2ae39953f6a01a4077eae21c759635babb41ca458d364610af626666` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | 0 | 0 | 0 | 321.0 MB | 2026-05-12 | `sha256:a3c7b67f4902` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:a3c7b67f4902f4325122e23456d30da597e7cf2460a93f1eb84b59cb2cbf5adb` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | 0 | 0 | 0 | 324.0 MB | 2026-05-12 | `sha256:49b60242c934` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:49b60242c934230947e52b2925458ae6e41d746e2cbc2e3c619a8c7fe119a607` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | 0 | 0 | 0 | 351.0 MB | 2026-05-12 | `sha256:f2dea1752a70` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:f2dea1752a7052149ee1baa9c7d86fe2716d50f281edcea153b53aae303a62ba` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | 0 | 0 | 0 | 397.0 MB | 2026-05-12 | `sha256:4511fc87aa53` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:4511fc87aa535636744b1edb2c72aef6584d945ff50a017e0117ba142ba65e9c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | 0 | 0 | 0 | 475.0 MB | 2026-05-12 | `sha256:fd14722fab26` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:fd14722fab26f0ec24206cec0335cf82639e9d02b94ebd810372a3a46fb5f1c1` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | 0 | 0 | 0 | 520.0 MB | 2026-05-12 | `sha256:ab68247308b1` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:ab68247308b1d12f834b59be3e923d137064d2fe6f2fd7c816a6dd0651272e86` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | 0 | 0 | 85 | 427.0 MB | 2026-05-12 | `sha256:5f6c72a5ee94` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:5f6c72a5ee94495f43929b227072da5028b02402c2980a5d5382aaf8b5f47368` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | 0 | 0 | 85 | 455.0 MB | 2026-05-12 | `sha256:e6cec562d4bf` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:e6cec562d4bf20ba4d216f44248074526b15037e8691f915204bce82a2194611` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | 0 | 0 | 85 | 499.0 MB | 2026-05-12 | `sha256:12b290110ec0` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:12b290110ec0afeb0912fe9d095d5bb31dfbffdfb8e0f10c4fdf46cbb4cb221d` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-09 | `sha256:a0d8e48ae705` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:a0d8e48ae705882480340e393cf528a4cab134b4e3e8251caf32dc323119f9ea` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-09 | `sha256:a0d8e48ae705` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:a0d8e48ae705882480340e393cf528a4cab134b4e3e8251caf32dc323119f9ea` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-09 | `sha256:2b7bc798de75` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:2b7bc798de75eed10379e13fa1f828329fb097ff900aad79200add7653c11d37` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-09 | `sha256:d9579bf850b6` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:d9579bf850b66deb11f92212bddfedd79003a575caad6d64ca059c23a77b85af` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-09 | `sha256:2b7bc798de75` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:2b7bc798de75eed10379e13fa1f828329fb097ff900aad79200add7653c11d37` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | 0 | 5 | 8 | 122.0 MB | 2026-05-09 | `sha256:d9579bf850b6` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:d9579bf850b66deb11f92212bddfedd79003a575caad6d64ca059c23a77b85af` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 8 | 50 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 10 | 50 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-09 | `sha256:a632fdff9455` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:a632fdff945567b1d8a5535142131d9da6357f69877e4e13bdebbe36a66794bb` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-09 | `sha256:8ea9673360c2` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:8ea9673360c2cc00cac6fb2205c3f4ca694c31aa892544071973d904a0a2bb9d` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-09 | `sha256:a632fdff9455` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:a632fdff945567b1d8a5535142131d9da6357f69877e4e13bdebbe36a66794bb` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-09 | `sha256:8ea9673360c2` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:8ea9673360c2cc00cac6fb2205c3f4ca694c31aa892544071973d904a0a2bb9d` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-09 | `sha256:dc4849f644bc` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:dc4849f644bca03c3cd15d3db8dc2e5a1d60ba749c87375b54c7b4fd8c1a3e7a` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-09 | `sha256:dc4849f644bc` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:dc4849f644bca03c3cd15d3db8dc2e5a1d60ba749c87375b54c7b4fd8c1a3e7a` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 475.0 MB | 2026-05-12 | `sha256:fd14722fab26` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:fd14722fab26f0ec24206cec0335cf82639e9d02b94ebd810372a3a46fb5f1c1` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 520.0 MB | 2026-05-12 | `sha256:ab68247308b1` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:ab68247308b1d12f834b59be3e923d137064d2fe6f2fd7c816a6dd0651272e86` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | 2026-05-11 | `sha256:31211bf7d7fb` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:31211bf7d7fbddfe501802c7e6c91e7605a7e2d90ff903a10d6da5dd30e181e9` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | 2026-05-11 | `sha256:6b48a8fb2ae3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:6b48a8fb2ae39953f6a01a4077eae21c759635babb41ca458d364610af626666` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | 2026-05-06 | `sha256:4c30ebfa4129` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4c30ebfa41297a79e98dc08117cc17e80245aacff9dfd0578cbc8c75c2368566` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | 2026-05-06 | `sha256:e79ae62e905e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:e79ae62e905ee0b492d0755abe37ac68001082aa638be12340540a535fc26324` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | 2026-05-06 | `sha256:82e37ddcf271` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:82e37ddcf271af1b720036697d8ab8c95f7001c3eaf7694a9ca17b35d84085de` |
