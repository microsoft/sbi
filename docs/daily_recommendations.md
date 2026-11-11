# Daily Recommended Images by Language

_Generated: 2026-11-11T02:17:22Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **46 configured sources** across 12 groups (see [repositories.json](../config/repositories.json)):

**Base / minimal images (no runtime):**

- `mcr.microsoft.com/azurelinux/base/core:3.0`
- `mcr.microsoft.com/azurelinux/distroless/base:3.0`
- `mcr.microsoft.com/azurelinux/distroless/minimal:3.0`
- `mcr.microsoft.com/azurelinux/busybox:1.36`

**.NET self-contained / Native AOT deployment bases:**

- `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled`
- `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless`

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
- `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0`

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

**.NET Ubuntu Chiseled images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled`
- `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra`

**.NET Alpine images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-alpine`

**.NET Debian images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

**Rust build images:**

- `mcr.microsoft.com/azurelinux/base/rust:1`

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

### Alpine

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine` | 10.0.15 | - | 0 | 0 | 0 | 121.0 MB | 2026-11-10 | `sha256:21c3f9798416` | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine@sha256:21c3f9798416cba4c0576a947c286231cb71ea2884896cb968253f5c9668f930` |

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.34 | - | 0 | 0 | 0 | 102.0 MB | 2026-11-10 | `sha256:c67eb11d6d5a` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:c67eb11d6d5a6275561190f3e842d5f1172212ba52e32181eb379ec8ef3000fe` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.23 | - | 0 | 0 | 0 | 108.0 MB | 2026-11-10 | `sha256:b48c851f943e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:b48c851f943e086425c2a16c79553337f678eede0634ccf5ce597b3db0ab4fa0` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 113.0 MB | 2026-11-10 | `sha256:7454e37035d8` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:7454e37035d87dc47c511a465837e74ec83bcb3c4e5cbfda2d8d22437dca40fa` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.34 | - | 0 | 0 | 0 | 127.0 MB | 2026-11-10 | `sha256:9cc006d455a5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:9cc006d455a547ce9f45370f482454e78cdd2940ea31e89fd79ca9f0a87e15b6` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.23 | - | 0 | 0 | 0 | 133.0 MB | 2026-11-10 | `sha256:55738825e188` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:55738825e1888742f8b214b32a2cbc4d94d849aa57f0bf15c31974bdab0e1cc3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 140.0 MB | 2026-11-10 | `sha256:cc62624ed65e` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:cc62624ed65e094785e3a5efc2a85da2635204b383fd619f2c50df59fdb14260` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.404 | - | 0 | 0 | 0 | 956.0 MB | 2026-11-10 | `sha256:559c47af49c8` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:559c47af49c81eee4057de212d8b16f0c1a5df8b0d0d38ae6499093ea4e6f5a1` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.321 | - | 0 | 0 | 10 | 901.0 MB | 2026-11-10 | `sha256:c785479263b3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c785479263b326cbb43a64940eb16dc0e07f3adf67e2508500242b80d64a17f7` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.428 | - | 0 | 16 | 27 | 899.0 MB | 2026-11-10 | `sha256:d925a0008079` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:d925a0008079bdde6e19a8d73c0b91d77654b1df20a2eabb3d72ccc6fa36a56e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.34 | - | 1 | 34 | 227 | 194.0 MB | 2026-11-10 | `sha256:21b0f8ba67be` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:21b0f8ba67bef2f121781d17a8e21956aeaa2d8da8e4b5db0da1ea8f91e2b6a1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.34 | - | 1 | 34 | 227 | 218.0 MB | 2026-11-10 | `sha256:dccacae9e0e1` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:dccacae9e0e103b5d3ba5fbe442dd8d3316b990cd8f2ec9be0a0ec9db0c0e520` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.428 | - | 1 | 131 | 510 | 868.0 MB | 2026-11-10 | `sha256:1b955f179b02` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:1b955f179b02df4f0eb5fadedaa898cbe44d620a64aab52eec72c8e91d1355b5` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.15 | - | 0 | 0 | 0 | 124.0 MB | 2026-11-10 | `sha256:7c625e44e5d5` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:7c625e44e5d588b4d7b2a913e94ca4bd5098ce47442da404829e4c7936a764e1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.15 | - | 0 | 0 | 1 | 162.0 MB | 2026-11-10 | `sha256:cae5b962bfb7` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:cae5b962bfb786606def03d250c6a645cc98839095c89e806e4bd8978a68115b` |
| 3 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.34 | - | 0 | 0 | 21 | 193.0 MB | 2026-11-10 | `sha256:7f2884767083` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7f28847670838ca16b8d8e228cfacd023fe74b279927d3cd01c640ce1b453283` |
| 4 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.23 | - | 0 | 0 | 21 | 198.0 MB | 2026-11-10 | `sha256:1830446e7cef` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:1830446e7cef0e0b2eeb9251ac430908fdb91a23909c5cc754cb09b20416e883` |
| 5 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.15 | - | 0 | 0 | 21 | 203.0 MB | 2026-11-10 | `sha256:80a51954ac70` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:80a51954ac70925e5783cb1ed35129a4243dcb1cddef20066b67346e17db284d` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.34 | - | 0 | 0 | 21 | 217.0 MB | 2026-11-10 | `sha256:fb50a3d0efad` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:fb50a3d0efad014b4fa8ee85fa498818172e38216cc73742821319f9db827bb2` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.23 | - | 0 | 0 | 21 | 223.0 MB | 2026-11-10 | `sha256:d3c3622111a4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:d3c3622111a4f6b380975634fb20dd18b5ecf9b230731017e5eb49ddd765db64` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.15 | - | 0 | 0 | 21 | 230.0 MB | 2026-11-10 | `sha256:bfd912c83295` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:bfd912c83295ce66d4b706ebb8541b9f14259b0e76fbb88b2751f58ce2399d2e` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.404 | - | 0 | 0 | 32 | 919.0 MB | 2026-11-10 | `sha256:fe0fd4e93bd2` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:fe0fd4e93bd2fbb47c883b7a6369e5c8d3db8bf01cdc5981f6ee35fbdf3ff4fb` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.321 | - | 0 | 0 | 42 | 863.0 MB | 2026-11-10 | `sha256:eeaa543fc683` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:eeaa543fc683baaae7389e83786f749add545f68cd2fecff19212f23a3bec58f` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-11-09 | `sha256:ad2f3b8eabd6` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:ad2f3b8eabd6110d3537e270d633cbfaa2108b98cd04a41df430512a9617f7bb` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 36 | 126 | 811.0 MB | 2026-11-09 | `sha256:a5d4b744e914` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:a5d4b744e9142c24601aef0d780820834a3d46d47e11a414dbef1a20f1d11fab` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33 | - | 0 | 0 | 0 | 324.0 MB | 2026-11-10 | `sha256:d2a8407d6976` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:d2a8407d6976948bfe357e1bc48539486ac533f0f24c4a2f146d6a6fc5c45aec` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 0 | 0 | 327.0 MB | 2026-11-10 | `sha256:c8376bcc8e9b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:c8376bcc8e9b886586309e99acfa4fafaf2a98d0c16cb14545be5f817a6a460f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 0 | 0 | 355.0 MB | 2026-11-10 | `sha256:86b9793a9c95` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:86b9793a9c95f3b6136422d680696def825e721812143e245436b5958a00fcdc` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 0 | 0 | 400.0 MB | 2026-11-10 | `sha256:956924449100` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:95692444910009dd13ee4fcd23fb6aea153aa5c916e0f12e1f43e10022ea0f77` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-11-10 | `sha256:e48471c5a5ca` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:e48471c5a5caec1f2a8436894fb4ee9fa9aebc3bba4868d397c3348d2c8b5302` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-11-10 | `sha256:e6c83426b583` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e6c83426b5839974adc1544c6ef46fc57e2b237e22c41791e9fb1cae430a3d02` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 0 | 72 | 443.0 MB | 2026-11-10 | `sha256:cb1f402e57a5` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:cb1f402e57a50f04b92999c029c455afe0f7fbe6d8f8550251808dd9d51f9060` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 0 | 72 | 470.0 MB | 2026-11-10 | `sha256:b349e3835860` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:b349e3835860a751aa0420473752d014184bfc53dc7ed161ca89cb6518e65029` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 0 | 72 | 514.0 MB | 2026-11-10 | `sha256:f9c5358b03dc` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:f9c5358b03dcb27f539a9df18b1da26353dda8bf8b4b5de44a871f67a404bf7c` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.1 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-11-05 | `sha256:700f81898b1a` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:700f81898b1a66c186a9a64a3fd20f61942e1d8335887e78ce4b7a562e0b7ff8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.1 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-11-05 | `sha256:d7e4008fd2a2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:d7e4008fd2a2b87d144f316e6310c75a83a1128f930bd3da81fdb25613ba0e55` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.1 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-11-05 | `sha256:5b9ff8def1fa` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:5b9ff8def1fa8fad6c0ba57f955c8cd5b2118fa9a9962d463a75a233333d719c` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 3 | 62 | 271 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | - | 4 | 13 | 41 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 4 | 23 | 67 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 4 | 23 | 67 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 4 | 30 | 112 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 4 | 32 | 131 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 4 | 36 | 126 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-11-05 | `sha256:4ab51be1200a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:4ab51be1200ac07000790db1330e308765fb8d75ea3159b0c4930f0b197b2fbe` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-11-05 | `sha256:e66ce0c54a5a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:e66ce0c54a5a0a8e2c679cbb8c66ef40a95bee7faed7ef87bfd52c4661073b06` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-11-05 | `sha256:379135d4dbe6` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:379135d4dbe66b12825a54d4ed7aef21aa61227533ff1b9151d70f95f9266746` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-11-05 | `sha256:6142d570852e` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:6142d570852e09fb676788262568ed2a2bf028b4eabc9b8f512206b1a9da28c3` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.9 MB | 2026-11-05 | `sha256:bc4d9a4e0d33` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:bc4d9a4e0d3317999e235cb5895dde28bf766d25533ae8e6ba311bd6d253f996` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-11-05 | `sha256:141581e735b7` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:141581e735b73dfed58f739c7fc8f5bd37546cba8de22051550eb9dfeecdcc2e` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.8 MB | 2026-11-10 | `sha256:39ff56ed5da2` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:39ff56ed5da2e66b5369fac94519e1c25bda7b69b10f4053de5f048c07cd7b81` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-11-05 | `sha256:370fb2c55e7f` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:370fb2c55e7f0474fed019623cce6a8db65a4d70dedfca9219fc69d6c7627f89` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-11-05 | `sha256:bfb630450e9c` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:bfb630450e9c3d7428d59f25ce6df95ba5634d664d96b06017aa56f87da5ec0e` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 0 | 14.2 MB | 2026-11-10 | `sha256:40eba5db6959` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:40eba5db69594d517c3a61e96bdacc2e3d6f8c99a44b52d55edee844d33518d0` |
