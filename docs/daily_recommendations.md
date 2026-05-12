# Daily Recommended Images by Language

_Generated: 2026-05-12T03:11:49Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 107.0 MB | 2026-05-09 | `sha256:7d4c6f2f2482` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:7d4c6f2f24820fff509ff4b37026eeddf9e1b1164814f7a713b15309f6b9a5f3` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 112.0 MB | 2026-05-09 | `sha256:4adc13cf4961` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:4adc13cf496142a778d198ea6c4b0d37ef1f1f5b2919221013bd9ba482e5427e` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.26 | 0 | 0 | 0 | 126.0 MB | 2026-05-09 | `sha256:9c8e95576c07` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:9c8e95576c076c380459c2b8c93b51d6ff970b812d3791deb9c4834f4145be2d` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 132.0 MB | 2026-05-09 | `sha256:23c142d5e835` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:23c142d5e8350d379dc35ac1cb8505e6438f251b3bbf8fcb1b9c4f59274182cc` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 139.0 MB | 2026-05-09 | `sha256:4e1a78fb4a2b` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:4e1a78fb4a2bc80e74f4dfb6638f8407be2feefb7f8b90e8e5381ce7e602d6bd` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.313 | 0 | 2 | 12 | 880.0 MB | 2026-05-09 | `sha256:59ee21d8b2e1` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:59ee21d8b2e113b93d8d483ae15dea1f9a3f11db61597e65597ea8e415eb7feb` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.203 | 0 | 2 | 14 | 935.0 MB | 2026-05-09 | `sha256:28e566bb0a16` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:28e566bb0a1608a7a9772da9a642451903dab766423c3e65e282dc539d16c497` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.26 | 2 | 13 | 148 | 193.0 MB | 2026-05-08 | `sha256:513ab9d7d5cb` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:513ab9d7d5cbc75e81fa933113a688d00a178a18ffa54c5631eed92902498e87` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.26 | 2 | 13 | 148 | 218.0 MB | 2026-05-08 | `sha256:1657ced1f401` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:1657ced1f4019bf174ff6ca844ad2545af44e787c135549a843f6a6d609d1c0e` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.420 | 3 | 28 | 281 | 850.0 MB | 2026-05-08 | `sha256:27ee23dc62a9` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:27ee23dc62a99643cfab2dc291582b6f2611ddc2598a0ff0958ddd0972012e52` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.26 | 0 | 0 | 22 | 193.0 MB | 2026-04-13 | `sha256:f1994beca66a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f1994beca66a16766311f5234fbfc9a0ce6efb1fa400fa0a3cd95eb1fd6f2a71` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.15 | 0 | 0 | 22 | 198.0 MB | 2026-04-13 | `sha256:08ec9b342063` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:08ec9b342063d70cf18e7ee10a3f95f22180be321af77e03a8ed2151093ef65f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.7 | 0 | 0 | 22 | 203.0 MB | 2026-04-21 | `sha256:8fb7ff015fcf` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:8fb7ff015fcf0ebc6e90105bd6db06875954e6dc3d374b9dbb34c732867d13e4` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.26 | 0 | 0 | 22 | 217.0 MB | 2026-04-13 | `sha256:b0811426a260` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:b0811426a2602ce0be3d1f605257c601b38c3a26e7b49283c9330de01c67cefd` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.15 | 0 | 0 | 22 | 223.0 MB | 2026-04-13 | `sha256:4180d760569f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:4180d760569fdf8f639cc79d62de294356c166eeebe3e4ab19d4f012960b5be8` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.7 | 0 | 0 | 22 | 230.0 MB | 2026-04-21 | `sha256:55e37c7795bf` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:55e37c7795bfaf6b9cc5d77c155811d9569f529d86e20647704bc1d7dd9741d4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.313 | 0 | 2 | 60 | 837.0 MB | 2026-04-13 | `sha256:c4d10394421c` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:c4d10394421c0e7c116b6f8866b7c80fa74352ea294629cda1fe4f2acd8f8f6f` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.203 | 0 | 2 | 62 | 892.0 MB | 2026-04-21 | `sha256:8a90a473da52` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:8a90a473da5205a16979de99d2fc20975e922c68304f5c79d564e666dc3982fc` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.420 | 0 | 12 | 70 | 836.0 MB | 2026-04-13 | `sha256:a91a987babf7` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:a91a987babf792bfe535ad5d85e94c492f66ffc2efc25aea343432a5bd606e93` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.10 | 0 | 0 | 0 | 809.0 MB | 2026-05-11 | `sha256:31211bf7d7fb` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:31211bf7d7fbddfe501802c7e6c91e7605a7e2d90ff903a10d6da5dd30e181e9` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.3 | 0 | 0 | 0 | 841.0 MB | 2026-05-11 | `sha256:6b48a8fb2ae3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:6b48a8fb2ae39953f6a01a4077eae21c759635babb41ca458d364610af626666` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | 0 | 0 | 0 | 475.0 MB | 2026-05-11 | `sha256:a8fdafa467f0` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:a8fdafa467f0607300ec4fbb321bb2f5ee2b41baf5934ef3dc31f4f10e9a20db` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | 0 | 0 | 0 | 519.0 MB | 2026-05-11 | `sha256:5db63263b9db` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5db63263b9db377f6ae63ac0dfcc6c9fb04942df37976f68980635e995398246` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | 0 | 5 | 8 | 321.0 MB | 2026-05-11 | `sha256:661c26089a19` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:661c26089a193864654d1c83d68bad199e1f8513d2bd15ba02921a2a5fe2a651` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | 0 | 5 | 8 | 324.0 MB | 2026-05-11 | `sha256:333506545a6d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:333506545a6d35ad06e3ca6d1b372f4e9f87145b8979454a6fcad3b017b9f894` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | 0 | 5 | 8 | 351.0 MB | 2026-05-11 | `sha256:f6f92ef45e5b` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:f6f92ef45e5b94aa028ab76b0e9fc27027174816ddba4ce3cd0f3cbb2f506fcc` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | 0 | 5 | 8 | 397.0 MB | 2026-05-11 | `sha256:5464115fae48` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:5464115fae488d4025a523fcc8099b0b7e4f0f361f4e429eebb98c9cacced18f` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | 0 | 5 | 93 | 427.0 MB | 2026-05-11 | `sha256:3cfcafae9bdd` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:3cfcafae9bdd64f5c254680f1f47cba288449d310c9591a34131834be53bbb89` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | 0 | 5 | 93 | 455.0 MB | 2026-05-11 | `sha256:aa8fe0cfd84a` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:aa8fe0cfd84a4a9e34185feadd404f98d277249c822a4556e952e104b9db501b` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | 0 | 5 | 93 | 499.0 MB | 2026-05-11 | `sha256:6dadd73c3889` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:6dadd73c388978e9573963b40c2101bc03f7399ea0b32b3c834a6e7ebd774776` |

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
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 475.0 MB | 2026-05-11 | `sha256:a8fdafa467f0` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:a8fdafa467f0607300ec4fbb321bb2f5ee2b41baf5934ef3dc31f4f10e9a20db` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 519.0 MB | 2026-05-11 | `sha256:5db63263b9db` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5db63263b9db377f6ae63ac0dfcc6c9fb04942df37976f68980635e995398246` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | 2026-05-11 | `sha256:31211bf7d7fb` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:31211bf7d7fbddfe501802c7e6c91e7605a7e2d90ff903a10d6da5dd30e181e9` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | 2026-05-11 | `sha256:6b48a8fb2ae3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:6b48a8fb2ae39953f6a01a4077eae21c759635babb41ca458d364610af626666` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | 2026-05-06 | `sha256:4c30ebfa4129` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4c30ebfa41297a79e98dc08117cc17e80245aacff9dfd0578cbc8c75c2368566` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | 2026-05-06 | `sha256:e79ae62e905e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:e79ae62e905ee0b492d0755abe37ac68001082aa638be12340540a535fc26324` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | 2026-05-06 | `sha256:82e37ddcf271` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:82e37ddcf271af1b720036697d8ab8c95f7001c3eaf7694a9ca17b35d84085de` |
