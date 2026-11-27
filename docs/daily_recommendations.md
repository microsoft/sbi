# Daily Recommended Images by Language

_Generated: 2026-11-27T02:17:00Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine` | 10.0.15 | - | 0 | 0 | 10 | 121.0 MB | 2026-11-10 | `sha256:21c3f9798416` | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine@sha256:21c3f9798416cba4c0576a947c286231cb71ea2884896cb968253f5c9668f930` |

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.34 | - | 0 | 0 | 0 | 102.0 MB | 2026-11-10 | `sha256:c67eb11d6d5a` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:c67eb11d6d5a6275561190f3e842d5f1172212ba52e32181eb379ec8ef3000fe` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.23 | - | 0 | 0 | 0 | 108.0 MB | 2026-11-10 | `sha256:b48c851f943e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:b48c851f943e086425c2a16c79553337f678eede0634ccf5ce597b3db0ab4fa0` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 112.0 MB | 2026-11-17 | `sha256:72d582f9a282` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:72d582f9a282eeb8d7dc3e769d1830db582eef276d0f40554eb1f4db73ad5b72` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.34 | - | 0 | 0 | 0 | 127.0 MB | 2026-11-10 | `sha256:9cc006d455a5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:9cc006d455a547ce9f45370f482454e78cdd2940ea31e89fd79ca9f0a87e15b6` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.23 | - | 0 | 0 | 0 | 133.0 MB | 2026-11-10 | `sha256:55738825e188` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:55738825e1888742f8b214b32a2cbc4d94d849aa57f0bf15c31974bdab0e1cc3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.15 | - | 0 | 0 | 0 | 140.0 MB | 2026-11-17 | `sha256:74beeeaa6c91` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:74beeeaa6c913d19fde119911ba2fc0ffb183752cea28bf3b106d53de8c91ade` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.404 | - | 0 | 0 | 0 | 957.0 MB | 2026-11-22 | `sha256:cb33856da043` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:cb33856da04320a799e89723630f368f62e9a5965d2b10d79c0b2bd5a3b95eca` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.321 | - | 3 | 47 | 161 | 901.0 MB | 2026-11-10 | `sha256:c785479263b3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c785479263b326cbb43a64940eb16dc0e07f3adf67e2508500242b80d64a17f7` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.428 | - | 3 | 63 | 178 | 899.0 MB | 2026-11-10 | `sha256:d925a0008079` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:d925a0008079bdde6e19a8d73c0b91d77654b1df20a2eabb3d72ccc6fa36a56e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.34 | - | 1 | 38 | 262 | 194.0 MB | 2026-11-10 | `sha256:21b0f8ba67be` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:21b0f8ba67bef2f121781d17a8e21956aeaa2d8da8e4b5db0da1ea8f91e2b6a1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.34 | - | 1 | 38 | 262 | 218.0 MB | 2026-11-10 | `sha256:dccacae9e0e1` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:dccacae9e0e103b5d3ba5fbe442dd8d3316b990cd8f2ec9be0a0ec9db0c0e520` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.428 | - | 1 | 141 | 565 | 868.0 MB | 2026-11-10 | `sha256:1b955f179b02` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:1b955f179b02df4f0eb5fadedaa898cbe44d620a64aab52eec72c8e91d1355b5` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.15 | - | 0 | 0 | 0 | 124.0 MB | 2026-11-24 | `sha256:bcdfe6b8e923` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:bcdfe6b8e923ad3d05f01dfa805ab97c99d742823ce030856002a6db38122712` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.15 | - | 0 | 0 | 1 | 162.0 MB | 2026-11-24 | `sha256:c8dd08a50d8f` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:c8dd08a50d8ffd19bd8da72edee05a7ac15b486a5b1869cf46fdb370bf1fbc76` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 203.0 MB | 2026-11-23 | `sha256:0a0e20d12b84` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:0a0e20d12b84b081fb737fddb92d8d205edce9ac2dc62796d2aab30f03804498` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 230.0 MB | 2026-11-23 | `sha256:aac50ceef103` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:aac50ceef10391e78a899e3d2ccd00bbafa2dcb56ccf4586cbccf00edb84fc21` |
| 5 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 193.0 MB | 2026-11-10 | `sha256:7f2884767083` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7f28847670838ca16b8d8e228cfacd023fe74b279927d3cd01c640ce1b453283` |
| 6 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 198.0 MB | 2026-11-10 | `sha256:1830446e7cef` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:1830446e7cef0e0b2eeb9251ac430908fdb91a23909c5cc754cb09b20416e883` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 217.0 MB | 2026-11-10 | `sha256:fb50a3d0efad` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:fb50a3d0efad014b4fa8ee85fa498818172e38216cc73742821319f9db827bb2` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 223.0 MB | 2026-11-10 | `sha256:d3c3622111a4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:d3c3622111a4f6b380975634fb20dd18b5ecf9b230731017e5eb49ddd765db64` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.404 | - | 0 | 0 | 116 | 911.0 MB | 2026-11-23 | `sha256:8c6edfb26c85` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:8c6edfb26c85153f236d26efebd9ad349aec403583f343f201ee67838b23658d` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.321 | - | 0 | 0 | 138 | 863.0 MB | 2026-11-10 | `sha256:eeaa543fc683` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:eeaa543fc683baaae7389e83786f749add545f68cd2fecff19212f23a3bec58f` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.11 | - | 0 | 0 | 0 | 845.0 MB | 2026-11-25 | `sha256:e8c8da1d9459` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:e8c8da1d9459ae8716a50a0bf01d9c5075039fcbbe7a10c456e476b088778bf1` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 83 | 286 | 811.0 MB | 2026-11-09 | `sha256:a5d4b744e914` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:a5d4b744e9142c24601aef0d780820834a3d46d47e11a414dbef1a20f1d11fab` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.1 | - | 0 | 0 | 0 | 480.0 MB | 2026-11-25 | `sha256:11429d458e12` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:11429d458e127cb29948be456bfd975ea0088e1f81d0ebf88a882d6eb4836a80` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.1 | - | 0 | 0 | 0 | 524.0 MB | 2026-11-25 | `sha256:57ae14a2bde1` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:57ae14a2bde1ccdceaf4bc1e7d08a739fbc3326285af0c7cb358ae736376c355` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33.1 | - | 0 | 0 | 1 | 324.0 MB | 2026-11-25 | `sha256:593b07399e14` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:593b07399e1481a952f82b106e12be8b080043a156a1f8447228e0a5a3aaeb9c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.1 | - | 0 | 0 | 1 | 327.0 MB | 2026-11-25 | `sha256:f80538e39a2e` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:f80538e39a2e329bf94a7b44080da45a5f605cc17e16a893410e443cd87c14cc` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.1 | - | 0 | 0 | 1 | 355.0 MB | 2026-11-25 | `sha256:57c060a7fd77` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:57c060a7fd7760d5ebed2101b4ac4d0f3286e9cbad94bcb618d0987dce73d594` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.1 | - | 0 | 0 | 1 | 400.0 MB | 2026-11-25 | `sha256:150018c6104f` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:150018c6104f6d1586743cf0766b92e22bf0040c4dc4e9ebc01d794b7d6b7bfa` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.1 | - | 0 | 0 | 283 | 431.0 MB | 2026-11-25 | `sha256:25bcfda434d9` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:25bcfda434d99dae4adc797a0a954bd4f22611e7f383873bdb0acc0f7ae85b3b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.1 | - | 0 | 0 | 283 | 458.0 MB | 2026-11-25 | `sha256:6a3d94f5b521` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:6a3d94f5b5212601654408a4de1624640389701e83aaa659660bf8b01a48df63` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.1 | - | 0 | 0 | 283 | 502.0 MB | 2026-11-25 | `sha256:ed9d057f8c95` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:ed9d057f8c95db6934a13d933a22cfcc6712d919e7b57848d81c15e99d89a40b` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.3 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-11-22 | `sha256:a3fb051354a7` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:a3fb051354a72cdaa1dac3ee9c4f74a0a4f903bd3da4152547a7bd4151a6e4eb` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.3 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-11-22 | `sha256:0b2dc2f5df44` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:0b2dc2f5df44272c07eca4450a7afde3a6028f9e4f7c0dde880d14c3c95f553b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.3 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-11-22 | `sha256:bafdd1aa0cbf` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:bafdd1aa0cbf0e17a521cfb2e96a235f31d14ea42d5d76b7b1575db31c907348` |
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
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-11-22 | `sha256:4789bedcb024` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:4789bedcb024e888dd8f4508f1f4977a07a3eb7e22f602a6e7389f7bf27fb9e2` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-11-22 | `sha256:6a7114c0debb` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:6a7114c0debbe0b58718c7040f7cc2b02245c7a526f7f581d8b3292941a57d4f` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-11-22 | `sha256:e313bff6397a` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:e313bff6397a9b36059990af1383bcde23f84126751d8dc6d8763601f5ecc314` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-11-22 | `sha256:cb8c7822089e` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:cb8c7822089e39790ca053ed47591388393c57f17e443f57004c7dfebed21770` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-11-16 | `sha256:acf5853dc365` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:acf5853dc3651b91bdbf3aee1e5383e689452f6637bd9bf74c9537771cafa6c2` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-11-22 | `sha256:c1bfb88b8e24` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:c1bfb88b8e2488f7ce88ca7d8520dba7c2e840827df0b9454b41c3951ad9f6cb` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.7 MB | 2026-11-17 | `sha256:046c3797695e` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:046c3797695ebd2cdad0c64c9e14a5004526614c78c35017c3c0e1f3d0e3bd98` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-11-16 | `sha256:11ecd8f7cd1c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:11ecd8f7cd1cad0e4fbb565ab98768f1a2ce35cb2730c867217baac011ae8645` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-11-22 | `sha256:a333b3233076` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a333b3233076d8b47d63237760e7b36576c0fe47fbf670d5d3517d2fd6fe5030` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 0 | 14.2 MB | 2026-11-24 | `sha256:2bb1d484500b` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:2bb1d484500b032aaafb75e876ffbea135b969d6259cacfcbf18a793e5d6e6ec` |
