# Daily Recommended Images by Language

_Generated: 2026-11-22T02:18:08Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.404 | - | 0 | 0 | 0 | 956.0 MB | 2026-11-20 | `sha256:796e932783f7` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:796e932783f7bf7c2d1f236cd92475d9df9b3cb3365615b493ebc0a0157c4017` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.321 | - | 2 | 43 | 135 | 901.0 MB | 2026-11-10 | `sha256:c785479263b3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c785479263b326cbb43a64940eb16dc0e07f3adf67e2508500242b80d64a17f7` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.428 | - | 2 | 59 | 152 | 899.0 MB | 2026-11-10 | `sha256:d925a0008079` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:d925a0008079bdde6e19a8d73c0b91d77654b1df20a2eabb3d72ccc6fa36a56e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.34 | - | 1 | 37 | 251 | 194.0 MB | 2026-11-10 | `sha256:21b0f8ba67be` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:21b0f8ba67bef2f121781d17a8e21956aeaa2d8da8e4b5db0da1ea8f91e2b6a1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.34 | - | 1 | 37 | 251 | 218.0 MB | 2026-11-10 | `sha256:dccacae9e0e1` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:dccacae9e0e103b5d3ba5fbe442dd8d3316b990cd8f2ec9be0a0ec9db0c0e520` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.428 | - | 1 | 140 | 552 | 868.0 MB | 2026-11-10 | `sha256:1b955f179b02` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:1b955f179b02df4f0eb5fadedaa898cbe44d620a64aab52eec72c8e91d1355b5` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.15 | - | 0 | 0 | 6 | 124.0 MB | 2026-11-10 | `sha256:7c625e44e5d5` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:7c625e44e5d588b4d7b2a913e94ca4bd5098ce47442da404829e4c7936a764e1` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.15 | - | 0 | 0 | 7 | 162.0 MB | 2026-11-10 | `sha256:cae5b962bfb7` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:cae5b962bfb786606def03d250c6a645cc98839095c89e806e4bd8978a68115b` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 219.0 MB | 2026-11-18 | `sha256:64a538092439` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:64a538092439382d055c87c0f49092efb8bca883c736c54ae523da7fd8e02d18` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.15 | - | 0 | 0 | 24 | 246.0 MB | 2026-11-18 | `sha256:12b1135f5f8b` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:12b1135f5f8b3c0a214f2a680ff2033dff4f684f445d6c697d9524ceeb303586` |
| 5 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 193.0 MB | 2026-11-10 | `sha256:7f2884767083` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:7f28847670838ca16b8d8e228cfacd023fe74b279927d3cd01c640ce1b453283` |
| 6 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 198.0 MB | 2026-11-10 | `sha256:1830446e7cef` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:1830446e7cef0e0b2eeb9251ac430908fdb91a23909c5cc754cb09b20416e883` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.34 | - | 0 | 0 | 37 | 217.0 MB | 2026-11-10 | `sha256:fb50a3d0efad` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:fb50a3d0efad014b4fa8ee85fa498818172e38216cc73742821319f9db827bb2` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.23 | - | 0 | 0 | 37 | 223.0 MB | 2026-11-10 | `sha256:d3c3622111a4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:d3c3622111a4f6b380975634fb20dd18b5ecf9b230731017e5eb49ddd765db64` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.404 | - | 0 | 0 | 46 | 926.0 MB | 2026-11-18 | `sha256:bef136a37196` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:bef136a371963aaa46ec97882845ef6f54dec15728bb4b128b1c2a400ffbc9b4` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.321 | - | 0 | 0 | 68 | 863.0 MB | 2026-11-10 | `sha256:eeaa543fc683` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:eeaa543fc683baaae7389e83786f749add545f68cd2fecff19212f23a3bec58f` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.3 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-11-20 | `sha256:ff81cb003a04` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:ff81cb003a046a411377a2736fb04623e2581e0606c8aedf0952c26118fc76e4` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.3 | :24-nonroot | 0 | 6 | 19 | 159.0 MB | 2026-11-20 | `sha256:d390b2602413` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:d390b2602413e19bd6b2c71f950b450fc49b7accf7134ab6cdaf58321c7787cc` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.3 | :24 | 0 | 6 | 19 | 159.0 MB | 2026-11-20 | `sha256:ec7dd5917dc7` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:ec7dd5917dc7a14947d8576781e3d3532c7b4a828cd3b3e7f0e53b89a1279b5c` |
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
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-11-20 | `sha256:0be7ee124b9f` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:0be7ee124b9f8e19e59770522f571489a224959721cbea4ddfc986062a4a9919` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-11-20 | `sha256:5d05f59de428` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:5d05f59de428636eb5cb4ce2bb40a73ffdb93bda6f3ce9c0367b4da37f252a3f` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-11-20 | `sha256:98a966e03aa9` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:98a966e03aa9bc24929274f93eb5e1d9457073c0afa10a919393e96177b05d4e` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-11-20 | `sha256:4d350bb424e7` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:4d350bb424e73b1c9032d1f24099403fc63e426bfece1a4b3c4317305a547c83` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.33.1 | - | 0 | 0 | 1 | 324.0 MB | 2026-11-20 | `sha256:1fe1f1cd5c28` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:1fe1f1cd5c2811db725c236bb462d4258b556cf151aab44cc0e13708ff2fff67` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.1 | - | 0 | 0 | 1 | 327.0 MB | 2026-11-20 | `sha256:137b883e675d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:137b883e675d29e32c602ad56dfb60c9acae64ed418eb12773b69f97dea3a747` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.1 | - | 0 | 0 | 1 | 355.0 MB | 2026-11-20 | `sha256:8368af3ae68c` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:8368af3ae68c29c6b96d94d6fbafd2a302f4cd115a3abb0de6a8c87d47518b8f` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.1 | - | 0 | 0 | 1 | 400.0 MB | 2026-11-20 | `sha256:e84386716537` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:e843867165371f415eb7bf14399eb97bad703577c788a0b389a26b2d26938080` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.1 | - | 0 | 5 | 12 | 483.0 MB | 2026-11-20 | `sha256:def4104404b6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:def4104404b6d83f22a3ac1ae63687e5cd3c73065bbc37772ad5c6e8b5eb6d1c` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.1 | - | 0 | 5 | 12 | 528.0 MB | 2026-11-20 | `sha256:ebf940d9bb65` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:ebf940d9bb65fc1d29fec1c880678ab1b23355a5defb247a1f98fbf9a1f2a4ee` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.1 | - | 0 | 0 | 104 | 431.0 MB | 2026-11-20 | `sha256:b6306560b865` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:b6306560b8658f9e2dc5a14c4e23d7e8e34f199ec8db4f39cf6e66284424b35f` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.1 | - | 0 | 0 | 104 | 458.0 MB | 2026-11-20 | `sha256:e6afbe9013b4` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:e6afbe9013b484e43c51b087a605b77a5bcec69ce5cda35f32478de91fc36986` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.1 | - | 0 | 0 | 104 | 502.0 MB | 2026-11-20 | `sha256:d616626511f7` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:d616626511f758ce7b0af62201ef3ff9f3363f55727eb9e1cf41a0cd13e33069` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.11 | - | 0 | 41 | 89 | 843.0 MB | 2026-11-18 | `sha256:4b01bf2cc3d9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:4b01bf2cc3d952f3f34d46a8b11fbfbe48cd4f06e2f1629b26a749ee4500e68b` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 2 | 79 | 260 | 811.0 MB | 2026-11-09 | `sha256:a5d4b744e914` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:a5d4b744e9142c24601aef0d780820834a3d46d47e11a414dbef1a20f1d11fab` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-11-16 | `sha256:acf5853dc365` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:acf5853dc3651b91bdbf3aee1e5383e689452f6637bd9bf74c9537771cafa6c2` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-11-20 | `sha256:11710db17093` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:11710db17093145cc8348f7652b6f1670d610e8cdfb9b0a9179d374cad0f95dd` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.7 MB | 2026-11-17 | `sha256:046c3797695e` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:046c3797695ebd2cdad0c64c9e14a5004526614c78c35017c3c0e1f3d0e3bd98` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-11-16 | `sha256:11ecd8f7cd1c` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:11ecd8f7cd1cad0e4fbb565ab98768f1a2ce35cb2730c867217baac011ae8645` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-11-20 | `sha256:a3546fb21809` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a3546fb21809426d10585a238fcc3c702f5fb120738c2d528c03c4f311b800e7` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 6 | 14.2 MB | 2026-11-10 | `sha256:40eba5db6959` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:40eba5db69594d517c3a61e96bdacc2e3d6f8c99a44b52d55edee844d33518d0` |
