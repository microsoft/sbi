# Daily Recommended Images by Language

_Generated: 2026-11-19T02:17:41Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 112.0 MB | 2026-11-17 | `sha256:72d582f9a282` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:72d582f9a282eeb8d7dc3e769d1830db582eef276d0f40554eb1f4db73ad5b72` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.34 | - | 0 | 0 | 0 | 127.0 MB | 2026-11-10 | `sha256:9cc006d455a5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:9cc006d455a547ce9f45370f482454e78cdd2940ea31e89fd79ca9f0a87e15b6` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.23 | - | 0 | 0 | 0 | 133.0 MB | 2026-11-10 | `sha256:55738825e188` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:55738825e1888742f8b214b32a2cbc4d94d849aa57f0bf15c31974bdab0e1cc3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 140.0 MB | 2026-11-17 | `sha256:74beeeaa6c91` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:74beeeaa6c913d19fde119911ba2fc0ffb183752cea28bf3b106d53de8c91ade` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.404 | - | 0 | 0 | 0 | 956.0 MB | 2026-11-17 | `sha256:4c4b2493a05e` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:4c4b2493a05ea496ce1b1b57f951ae0630aa4946bea13029bb2414131a107fd9` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.321 | - | 2 | 2 | 44 | 901.0 MB | 2026-11-10 | `sha256:c785479263b3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c785479263b326cbb43a64940eb16dc0e07f3adf67e2508500242b80d64a17f7` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.428 | - | 2 | 18 | 61 | 899.0 MB | 2026-11-10 | `sha256:d925a0008079` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:d925a0008079bdde6e19a8d73c0b91d77654b1df20a2eabb3d72ccc6fa36a56e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.34 | - | 1 | 37 | 252 | 194.0 MB | 2026-11-10 | `sha256:21b0f8ba67be` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:21b0f8ba67bef2f121781d17a8e21956aeaa2d8da8e4b5db0da1ea8f91e2b6a1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.34 | - | 1 | 37 | 252 | 218.0 MB | 2026-11-10 | `sha256:dccacae9e0e1` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:dccacae9e0e103b5d3ba5fbe442dd8d3316b990cd8f2ec9be0a0ec9db0c0e520` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.428 | - | 1 | 140 | 551 | 868.0 MB | 2026-11-10 | `sha256:1b955f179b02` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:1b955f179b02df4f0eb5fadedaa898cbe44d620a64aab52eec72c8e91d1355b5` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.15 | - | 0 | 0 | 6 | 124.0 MB | 2026-11-10 | `sha256:7c625e44e5d5` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:7c625e44e5d588b4d7b2a913e94ca4bd5098ce47442da404829e4c7936a764e1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.15 | - | 0 | 0 | 7 | 162.0 MB | 2026-11-10 | `sha256:cae5b962bfb7` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:cae5b962bfb786606def03d250c6a645cc98839095c89e806e4bd8978a68115b` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 219.0 MB | 2026-11-18 | `sha256:64a538092439` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:64a538092439382d055c87c0f49092efb8bca883c736c54ae523da7fd8e02d18` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 246.0 MB | 2026-11-18 | `sha256:12b1135f5f8b` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:12b1135f5f8b3c0a214f2a680ff2033dff4f684f445d6c697d9524ceeb303586` |
| 5 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.404 | - | 0 | 0 | 36 | 926.0 MB | 2026-11-18 | `sha256:bef136a37196` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:bef136a371963aaa46ec97882845ef6f54dec15728bb4b128b1c2a400ffbc9b4` |
| 6 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 193.0 MB | 2026-11-10 | `sha256:7f2884767083` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7f28847670838ca16b8d8e228cfacd023fe74b279927d3cd01c640ce1b453283` |
| 7 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 198.0 MB | 2026-11-10 | `sha256:1830446e7cef` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:1830446e7cef0e0b2eeb9251ac430908fdb91a23909c5cc754cb09b20416e883` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 217.0 MB | 2026-11-10 | `sha256:fb50a3d0efad` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:fb50a3d0efad014b4fa8ee85fa498818172e38216cc73742821319f9db827bb2` |
| 9 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 223.0 MB | 2026-11-10 | `sha256:d3c3622111a4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:d3c3622111a4f6b380975634fb20dd18b5ecf9b230731017e5eb49ddd765db64` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.321 | - | 0 | 0 | 58 | 863.0 MB | 2026-11-10 | `sha256:eeaa543fc683` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:eeaa543fc683baaae7389e83786f749add545f68cd2fecff19212f23a3bec58f` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.11 | - | 0 | 0 | 0 | 843.0 MB | 2026-11-18 | `sha256:4b01bf2cc3d9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:4b01bf2cc3d952f3f34d46a8b11fbfbe48cd4f06e2f1629b26a749ee4500e68b` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 2 | 47 | 169 | 811.0 MB | 2026-11-09 | `sha256:a5d4b744e914` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:a5d4b744e9142c24601aef0d780820834a3d46d47e11a414dbef1a20f1d11fab` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.1 | - | 0 | 0 | 0 | 480.0 MB | 2026-11-18 | `sha256:721ba2b35796` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:721ba2b3579652a2011d151efafad75eeddadd8361651ed95eec348222d15aea` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.1 | - | 0 | 0 | 0 | 524.0 MB | 2026-11-18 | `sha256:69d43fa472eb` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:69d43fa472ebd23228870d8edf85dd9f3167cb526ee7a0422745fc59e9e2c28b` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33.1 | - | 0 | 1 | 1 | 324.0 MB | 2026-11-18 | `sha256:99164be48235` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:99164be48235ef2cdf728a0a5e77c374cb1f41a00dda266b2f4cd862d7600bf1` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.1 | - | 0 | 1 | 1 | 327.0 MB | 2026-11-18 | `sha256:bfc1a464a310` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:bfc1a464a310305d1b2528619c87c1948561878d8a2bc2ad4adfffabe20521eb` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.1 | - | 0 | 1 | 1 | 355.0 MB | 2026-11-18 | `sha256:d3de6d821f9d` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:d3de6d821f9da36988e7c7c558a234dc006e7843245eea3e7e81798a65da1ee9` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.1 | - | 0 | 1 | 1 | 400.0 MB | 2026-11-18 | `sha256:f2cfd8c250de` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:f2cfd8c250de10b5062417de7721a2f6b01b48386adccb6848c5c3d86a34e96a` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.1 | - | 0 | 1 | 89 | 446.0 MB | 2026-11-18 | `sha256:3f6843252b9c` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:3f6843252b9c884668eb906aa76c146926ac15395c2523b3dde92994d0780e5c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.1 | - | 0 | 1 | 89 | 474.0 MB | 2026-11-18 | `sha256:a17f2b4e1ddf` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:a17f2b4e1ddf0789b6982f40dfffbe61dbf80f893ba23313dafe4b0e2fa3d4ef` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.1 | - | 0 | 1 | 89 | 518.0 MB | 2026-11-18 | `sha256:b72dc403d4e0` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:b72dc403d4e0ee5965f00979838c41ff07dbe98e622bafcb46b2c7464644777e` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.3 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-11-16 | `sha256:1c87a74ac10b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:1c87a74ac10b0ad7711bcc862fa6a2b13e6a4db96d3a026dbb7fcf8486d57ba3` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.3 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-11-16 | `sha256:64653c45a2fe` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:64653c45a2fe9899f744e00716b140a3ea7d2202deb3a9d6307ddf62732e3e3c` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.3 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-11-16 | `sha256:6b65f0c0fa7a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:6b65f0c0fa7ad56456e33ef53e6f5377c4aeb3aba39e42b8431d7a71959fa3dd` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 4 | 26 | 85 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 4 | 26 | 85 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 4 | 38 | 144 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 4 | 38 | 144 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | - | 4 | 62 | 189 | 123.0 MB | 2026-04-01 | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | - | 4 | 62 | 189 | 123.0 MB | 2026-04-01 | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 5 | 44 | 173 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-11-16 | `sha256:8961b59dde86` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:8961b59dde86fd38fe8efca25f50134b2772dca9389458f636f8cad070c3922c` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-11-16 | `sha256:a5318b273539` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:a5318b273539611641edb51055a71cc3cff22a393512ed8dc8c50ed4a6bfcb07` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-11-16 | `sha256:63357b11d5bc` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:63357b11d5bcd4ac06b39566fdcf33677b894fc11f0b2482ce69c920194aeb1e` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-11-16 | `sha256:69c54d82520a` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:69c54d82520acc1e7b1e0a821459154e4351f26a57d0e6b04ebe904c315e59fc` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-11-16 | `sha256:acf5853dc365` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:acf5853dc3651b91bdbf3aee1e5383e689452f6637bd9bf74c9537771cafa6c2` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-11-16 | `sha256:0e08811897dc` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:0e08811897dc1f97f761b9d256365ef33bc61ade0737bf28e52a3769bdf12df4` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.7 MB | 2026-11-17 | `sha256:046c3797695e` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:046c3797695ebd2cdad0c64c9e14a5004526614c78c35017c3c0e1f3d0e3bd98` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-11-16 | `sha256:11ecd8f7cd1c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:11ecd8f7cd1cad0e4fbb565ab98768f1a2ce35cb2730c867217baac011ae8645` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-11-16 | `sha256:3d16d8799f71` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:3d16d8799f71a78f5bddb3b7cda4a57df1f40ecfa58bac89933010596aa86880` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 6 | 14.2 MB | 2026-11-10 | `sha256:40eba5db6959` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:40eba5db69594d517c3a61e96bdacc2e3d6f8c99a44b52d55edee844d33518d0` |
