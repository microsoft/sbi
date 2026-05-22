# Daily Recommended Images by Language

_Generated: 2026-05-22T03:15:57Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 107.0 MB | 2026-05-20 | `sha256:489ecee4e885` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:489ecee4e885605c200248f430aeef6b096bb453e888aae753b5725df12d4d3d` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 112.0 MB | 2026-05-20 | `sha256:0c9c1039f5b9` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:0c9c1039f5b922b726c7c0509fde80e1c77cdf2c0f01294f2f0935ef13ee1a5b` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.27 | 0 | 0 | 0 | 126.0 MB | 2026-05-20 | `sha256:d50e7721162c` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:d50e7721162c8fabea3c9e50b0febb9405337e72283efc581728ecd5374bac74` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.16 | 0 | 0 | 0 | 132.0 MB | 2026-05-20 | `sha256:f304309c5331` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:f304309c533100192e68bdfcfb922afd15dd636f58f2e1b64ad272402f6f92b6` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.8 | 0 | 0 | 0 | 139.0 MB | 2026-05-20 | `sha256:d5ac88bf5784` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:d5ac88bf57841739c2dafeebce9adeb77247d40874cff9e8922aa72004957a92` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.300 | 0 | 0 | 0 | 918.0 MB | 2026-05-20 | `sha256:6b157661bfc0` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:6b157661bfc04c13c2307276b84e6e34f03e2f3991e5b2486a6bac8d2f4010d7` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.314 | 0 | 0 | 10 | 879.0 MB | 2026-05-20 | `sha256:523573a6ae86` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:523573a6ae86d7d1c2b4b0ad8ad7cbc38c8c3f62347292938d0808bdd3869c80` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.27 | 3 | 6 | 123 | 193.0 MB | 2026-05-20 | `sha256:6810933f14d0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:6810933f14d0b8a0eb193ab60e3a6429f8d263a7af891fa9ebd5f0eb6703febb` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.27 | 3 | 6 | 123 | 218.0 MB | 2026-05-20 | `sha256:19be23fe71e8` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:19be23fe71e885186495d8da1f2d417e553daaa99e16e5148c5ca3ee2f062512` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.421 | 3 | 31 | 253 | 849.0 MB | 2026-05-20 | `sha256:fc69dc5e0c97` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:fc69dc5e0c9789adaac5c8efce71ead4d016a51318667c4f26ce93574b1b9403` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.27 | 0 | 0 | 32 | 193.0 MB | 2026-05-20 | `sha256:7bb91a5cf371` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7bb91a5cf371059c9ebbb92142db6e6f7024a75da4cc306b181de0b57f3058e3` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.16 | 0 | 0 | 32 | 198.0 MB | 2026-05-20 | `sha256:b22e7cc84d54` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:b22e7cc84d549540db21a778a28012ba1527689ae24a7a06851d176328e24ed7` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.8 | 0 | 0 | 32 | 203.0 MB | 2026-05-20 | `sha256:d399699ebc8a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:d399699ebc8a27ab34665707ae6dc8f77ae478bd319444841a33a2b3840c5c9a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.27 | 0 | 0 | 32 | 217.0 MB | 2026-05-20 | `sha256:052a40facc3d` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:052a40facc3d32caa9920381df76b20628382409f021bf5d9f38dd67a6f936fa` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.16 | 0 | 0 | 32 | 223.0 MB | 2026-05-20 | `sha256:53683436482e` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:53683436482e444f31c9db6c34921f408a84e7ebcef41e2d3862c7c10051076a` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.8 | 0 | 0 | 32 | 230.0 MB | 2026-05-20 | `sha256:8c0b6857eab7` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8c0b6857eab7b2aa57884c839bf4678414606bd7d17370f18a842ac5cf414711` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.300 | 0 | 0 | 36 | 875.0 MB | 2026-05-20 | `sha256:c07906393326` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:c0790639332692a0d56cdd81ed581cfd24d040d9839764c138994866df89a3b6` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.314 | 0 | 0 | 46 | 837.0 MB | 2026-05-20 | `sha256:fda2105a0319` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:fda2105a0319b5ae5d7d51bd540ab0f7fc548ee6ea48cbc95c2b63b62ecad3a3` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.421 | 0 | 10 | 56 | 836.0 MB | 2026-05-20 | `sha256:6b0b7f73dc7c` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:6b0b7f73dc7cce85fe9eaf7cfcfd1dc109accc5b3782c8cba006fbe036da424e` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.10 | 0 | 0 | 0 | 809.0 MB | 2026-05-20 | `sha256:8937359450e9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8937359450e9c9faacaa9618579adba704d777402841322cee8fff828f4cf00c` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.3 | 0 | 0 | 0 | 841.0 MB | 2026-05-20 | `sha256:75a5c75f33e2` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:75a5c75f33e2238c6a5b93167cc4a7acd96d39e217d6e370ee7adcf5c29e2a6d` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.31 | 0 | 0 | 0 | 321.0 MB | 2026-05-20 | `sha256:71f0dfe51fc7` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:71f0dfe51fc74da5d3cb9d16bdb13aca2553181a853601dd4179f9ad64d404d2` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.19 | 0 | 0 | 0 | 324.0 MB | 2026-05-20 | `sha256:84d198301c9d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:84d198301c9d006c36c45270042d41c7ed437409bf2c8b3978b26914c58d7532` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.11 | 0 | 0 | 0 | 351.0 MB | 2026-05-20 | `sha256:236523e452f7` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:236523e452f7ec6a94b5841435e8292cabdfe60309e172b72b315d211d1f7183` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.3 | 0 | 0 | 0 | 397.0 MB | 2026-05-20 | `sha256:7dc91c61be71` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:7dc91c61be71989ac02e3555ef62a10677121c807ea63979db94c546182a2178` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.11 | 0 | 0 | 0 | 475.0 MB | 2026-05-20 | `sha256:60d7ded7c32f` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:60d7ded7c32f7d2de0f9c0ecaf9bcfa848c99928401c7473d6f433ae98044d0f` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.3 | 0 | 0 | 0 | 520.0 MB | 2026-05-20 | `sha256:52e2b60a57a4` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:52e2b60a57a4beeeaafb580d886aa8808216392c5224d6e556f22a3ea404b055` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.19 | 0 | 0 | 97 | 427.0 MB | 2026-05-20 | `sha256:db034e425aba` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:db034e425ababb507ebb7085148acddb5638678edd97d1c49b9f022bcfcff54c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.11 | 0 | 0 | 97 | 455.0 MB | 2026-05-20 | `sha256:cfdacaa6248f` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:cfdacaa6248fdf71705271071d20aa60b877dceb047771ff26c2e210f406c28f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.3 | 0 | 0 | 97 | 499.0 MB | 2026-05-20 | `sha256:9fd6bfd0ad11` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:9fd6bfd0ad11cffe2de6d2f52ef16af097d09a7d3328712f9512bdfcc32358ee` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-17 | `sha256:20f7ab20fab6` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:20f7ab20fab66f75d70753ac87a8f8c966521d811b109786174e15f2615464c2` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | 0 | 0 | 0 | 161.0 MB | 2026-05-17 | `sha256:20f7ab20fab6` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:20f7ab20fab66f75d70753ac87a8f8c966521d811b109786174e15f2615464c2` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.14.1 | 0 | 5 | 9 | 122.0 MB | 2026-05-17 | `sha256:ce9e8f2c397d` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:ce9e8f2c397d0456ee0358e93c580df23853926e97b07e505550a910f95bef09` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.14.1 | 0 | 5 | 9 | 122.0 MB | 2026-05-17 | `sha256:9bfbdc2fbf23` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:9bfbdc2fbf23c114c9ef18926930a11e53536d0519c895cd75bf6f0fc368bf7b` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | 0 | 5 | 9 | 122.0 MB | 2026-05-17 | `sha256:ce9e8f2c397d` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:ce9e8f2c397d0456ee0358e93c580df23853926e97b07e505550a910f95bef09` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | 0 | 5 | 9 | 122.0 MB | 2026-05-17 | `sha256:9bfbdc2fbf23` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9bfbdc2fbf23c114c9ef18926930a11e53536d0519c895cd75bf6f0fc368bf7b` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 8 | 50 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 23 | 53 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 10 | 50 | 146.0 MB | 2026-03-04 | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-17 | `sha256:c0279d3b8bdd` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:c0279d3b8bdddbe189585e9b084b234eed6285ea60dc2c5dec42ba6c8f3cf10b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-17 | `sha256:6e56d7b3b3b6` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:6e56d7b3b3b6d846401e42a32d969e045c13f476dbd6ad560d82649096d3f81f` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-17 | `sha256:c0279d3b8bdd` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:c0279d3b8bdddbe189585e9b084b234eed6285ea60dc2c5dec42ba6c8f3cf10b` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | 2026-05-17 | `sha256:6e56d7b3b3b6` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:6e56d7b3b3b6d846401e42a32d969e045c13f476dbd6ad560d82649096d3f81f` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-17 | `sha256:485299b016fe` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:485299b016fe5ae745ffee27f0b8a850576841205ed1d420c9a84b126198e320` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | 2026-05-17 | `sha256:485299b016fe` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:485299b016fe5ae745ffee27f0b8a850576841205ed1d420c9a84b126198e320` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 475.0 MB | 2026-05-20 | `sha256:60d7ded7c32f` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:60d7ded7c32f7d2de0f9c0ecaf9bcfa848c99928401c7473d6f433ae98044d0f` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 520.0 MB | 2026-05-20 | `sha256:52e2b60a57a4` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:52e2b60a57a4beeeaafb580d886aa8808216392c5224d6e556f22a3ea404b055` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | 2026-05-20 | `sha256:8937359450e9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8937359450e9c9faacaa9618579adba704d777402841322cee8fff828f4cf00c` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | 2026-05-20 | `sha256:75a5c75f33e2` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:75a5c75f33e2238c6a5b93167cc4a7acd96d39e217d6e370ee7adcf5c29e2a6d` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | 2026-05-17 | `sha256:0c64ab9cfc44` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:0c64ab9cfc44d4f100c0590bd59ead9afedda6cc54f14bb7465b5f9c35ddc037` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | 2026-05-17 | `sha256:f550b116f2db` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:f550b116f2db19ec755e25f180e1eb1beb9546263df4863678efd7e1318258d2` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | 2026-05-17 | `sha256:f5e224c47997` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:f5e224c47997aa4a5d3d8addfcc3866e175e7026368a71ce1be2c0eed1876f04` |
