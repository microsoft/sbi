# Daily Recommended Images by Language

_Generated: 2026-04-23T03:10:12Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 107.0 MB | `sha256:f57b292ed7d6` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:f57b292ed7d6f7c88d3b53addfcc02e520aae273eddb57cf0ea49f9396cb2424` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 112.0 MB | `sha256:79af73677bbc` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:79af73677bbc271ed7789d11df2f25974b65e7bb85399b6e9be06c7c37122fa7` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.26 | 0 | 0 | 0 | 126.0 MB | `sha256:742fe85d9ef9` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:742fe85d9ef99cd3d43fc11d4e97e1fb3fe5d30d02d183c7d52aa6238f05904c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 132.0 MB | `sha256:22b1bd549a20` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:22b1bd549a200249bea512b6b2fce8885a47d98c346d0d33511b2bf70653d207` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.7 | 0 | 0 | 0 | 139.0 MB | `sha256:b782ff63dec5` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:b782ff63dec54c5537b217ef1f3ab1900058f070c786bc214687c6fe53bba1e6` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.203 | 0 | 3 | 27 | 943.0 MB | `sha256:4f08dbd55f00` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:4f08dbd55f00d1f7eb232ddacb24ecdab62fd64774a25c1da5b6daeb76e76a75` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.313 | 0 | 4 | 28 | 888.0 MB | `sha256:13b64ec155e9` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:13b64ec155e9aefa220263957b6d86f1add651106fc5df3cbecacb7a74eb3b06` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.26 | 1 | 7 | 131 | 193.0 MB | `sha256:24117fb11f2f` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:24117fb11f2f27bdacc6f1ffab16e3a7cddd2417b4eae5e6bc825810a0d1746d` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.26 | 1 | 7 | 131 | 218.0 MB | `sha256:f88c77644f4c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:f88c77644f4c480a62d3b46dc74db8d5472a24e282df8b1e56195c689d35a6db` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.420 | 1 | 22 | 233 | 850.0 MB | `sha256:4b1cdaa57eed` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:4b1cdaa57eed2cecabcf29bdb9bce11e8ca1c287d39dfd2c8b534663ea94d493` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.26 | 0 | 0 | 20 | 193.0 MB | `sha256:f1994beca66a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f1994beca66a16766311f5234fbfc9a0ce6efb1fa400fa0a3cd95eb1fd6f2a71` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.15 | 0 | 0 | 20 | 198.0 MB | `sha256:08ec9b342063` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:08ec9b342063d70cf18e7ee10a3f95f22180be321af77e03a8ed2151093ef65f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.7 | 0 | 0 | 20 | 203.0 MB | `sha256:8fb7ff015fcf` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:8fb7ff015fcf0ebc6e90105bd6db06875954e6dc3d374b9dbb34c732867d13e4` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.26 | 0 | 0 | 20 | 217.0 MB | `sha256:b0811426a260` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:b0811426a2602ce0be3d1f605257c601b38c3a26e7b49283c9330de01c67cefd` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.15 | 0 | 0 | 20 | 223.0 MB | `sha256:4180d760569f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:4180d760569fdf8f639cc79d62de294356c166eeebe3e4ab19d4f012960b5be8` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.7 | 0 | 0 | 20 | 230.0 MB | `sha256:55e37c7795bf` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:55e37c7795bfaf6b9cc5d77c155811d9569f529d86e20647704bc1d7dd9741d4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.313 | 0 | 2 | 36 | 837.0 MB | `sha256:c4d10394421c` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:c4d10394421c0e7c116b6f8866b7c80fa74352ea294629cda1fe4f2acd8f8f6f` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.203 | 0 | 2 | 38 | 892.0 MB | `sha256:8a90a473da52` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:8a90a473da5205a16979de99d2fc20975e922c68304f5c79d564e666dc3982fc` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.420 | 0 | 12 | 46 | 836.0 MB | `sha256:a91a987babf7` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:a91a987babf792bfe535ad5d85e94c492f66ffc2efc25aea343432a5bd606e93` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.9 | 0 | 0 | 0 | 820.0 MB | `sha256:e248a20eddb7` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:e248a20eddb7cd542942489ac18c675c7669066660e63d61ba7476e51ffac231` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.2 | 0 | 0 | 0 | 853.0 MB | `sha256:ee3b11755cb3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:ee3b11755cb3401848dc3140bbaae8da81dce6f7c7569857b2f45f040ae8868b` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 486.0 MB | `sha256:83be8d641034` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:83be8d64103415f8867d6742eb6a6c4dbed2f707d01048a986e5609513478b74` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 530.0 MB | `sha256:5bc2481b6cfc` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5bc2481b6cfcf77d809f5644326ef33ca2abb5b41b79a69c8c72d905a843f0fb` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 10 | 320.0 MB | `sha256:7b15e181d31b` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:7b15e181d31b2352b99ee93bf1ee057d1527360f5600adac364a9d3e38a0cd45` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 10 | 324.0 MB | `sha256:4f0a05eeced0` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:4f0a05eeced0d4e682f8cc357160f48445ff5989075ad842bed9b7e91ef85f86` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 10 | 351.0 MB | `sha256:e290f1a0e746` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:e290f1a0e7466ce166b22ecf577b790853dc1492e81a6f03ab6073317f0abb69` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 10 | 396.0 MB | `sha256:bb5d66d81fc8` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:bb5d66d81fc877446ba157c632f621e51712517efd70f98f73c7b3af5e8a36e7` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 91 | 427.0 MB | `sha256:126462b802ef` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:126462b802ef798664fb0883f6368d28587b71da1991628d7df2dbbd34c1dfce` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 91 | 454.0 MB | `sha256:c0c504cacb5b` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:c0c504cacb5b4ced74d47efe02067bbb99664971fc2085c5f6dae660981f5e0d` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 91 | 499.0 MB | `sha256:bbee4563d5d9` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:bbee4563d5d9b88e75970d807061a8c63016e01c5eb93b3035365e3e1a6a3e2c` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 486.0 MB | `sha256:83be8d641034` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:83be8d64103415f8867d6742eb6a6c4dbed2f707d01048a986e5609513478b74` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 530.0 MB | `sha256:5bc2481b6cfc` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5bc2481b6cfcf77d809f5644326ef33ca2abb5b41b79a69c8c72d905a843f0fb` |
| 3 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 820.0 MB | `sha256:e248a20eddb7` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:e248a20eddb7cd542942489ac18c675c7669066660e63d61ba7476e51ffac231` |
| 4 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 853.0 MB | `sha256:ee3b11755cb3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:ee3b11755cb3401848dc3140bbaae8da81dce6f7c7569857b2f45f040ae8868b` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 1 | 17 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 9 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 1 | 24 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 10 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 1 | 24 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 5 | 41 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.13.0 | 0 | 5 | 41 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 23 | 51 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 23 | 51 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.13.0 | 0 | 23 | 51 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.13.0 | 0 | 23 | 51 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 7 | 41 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 1 | 7 | 41 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 1 | 15 | 42 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 1 | 15 | 42 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 10 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 1 | 23 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |
