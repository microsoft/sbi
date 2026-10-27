# Daily Recommended Images by Language

_Generated: 2026-10-27T02:25:47Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-24 | `sha256:c9af2430171d` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:c9af2430171d7ca2252372f8656e134aa62cf599c227a0773ca13239cd30219f` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-24 | `sha256:20246a3c6190` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:20246a3c6190e9ffac89625c300978163c32b1682d2b629556c32368a9046885` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 126.0 MB | 2026-10-24 | `sha256:e27d09562205` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e27d0956220534c06c533469fe325fdf6a3c0b49e20d94e26b877fa3d1a8f88a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-24 | `sha256:844c5d3a1c0d` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:844c5d3a1c0dc772f568ce0306b1551ead353427494bb967075710bb4ec65dc1` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-24 | `sha256:45908708ba21` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:45908708ba217e57515ab3ba03bfb1eeb5de75cf22a9e196f809a18fdfa13a01` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.403 | - | 0 | 0 | 0 | 963.0 MB | 2026-10-24 | `sha256:f0414bcafb27` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:f0414bcafb2751af8dcc7f72f6055eb456ca9931074b2eb837c6c9d166360fb5` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.320 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-24 | `sha256:41fe79a756a3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:41fe79a756a3ca3aa1cb79ce4f84e399b21ea871e38795fc7afcac02bf71f431` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.33 | - | 1 | 58 | 242 | 195.0 MB | 2026-10-24 | `sha256:ac29abc7855e` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:ac29abc7855ef3dd4697be50ac8baa99151d893d5c4545f48b44579ce2d8f439` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.33 | - | 1 | 58 | 242 | 219.0 MB | 2026-10-24 | `sha256:63a1e42746a2` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:63a1e42746a20795273c219f1577a8826644f12f67d829e818d8ccec92915493` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.427 | - | 1 | 95 | 453 | 869.0 MB | 2026-10-24 | `sha256:ff39a765f046` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ff39a765f04619c8d146796fc09b38499ae5984f4a2a4c3a0a57d0022855daee` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.33 | - | 0 | 0 | 26 | 193.0 MB | 2026-10-26 | `sha256:2ba6f26ec6dc` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:2ba6f26ec6dc479aa7f8016e1b0bd80797c52104872ba97766d0b160bb97d8f6` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.22 | - | 0 | 0 | 26 | 198.0 MB | 2026-10-26 | `sha256:3562dcf10791` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:3562dcf10791d9b1a20097500da37dc5390a1a6c59f56e40806f552cdd7c1666` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.14 | - | 0 | 0 | 26 | 203.0 MB | 2026-10-26 | `sha256:a3f9f81954b9` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a3f9f81954b9c810af2ed18849716c8dc74a4707bd0907fdbd69be81554c0f30` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.33 | - | 0 | 0 | 26 | 217.0 MB | 2026-10-26 | `sha256:d0a0fa004862` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d0a0fa004862a9cca83208800f706dcc2523c4db29f09d501395f80969e03070` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.22 | - | 0 | 0 | 26 | 223.0 MB | 2026-10-26 | `sha256:b7459513d45f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:b7459513d45fa48ac9a1ca2a191c720aa15b8a6a1b961a1ab2c6965816afe7d3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.14 | - | 0 | 0 | 26 | 230.0 MB | 2026-10-26 | `sha256:de0ee358cdd6` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:de0ee358cdd672f7017a003bed21e9ac9ed48a27c922a59b18175b6e82fa7fee` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.403 | - | 0 | 0 | 33 | 918.0 MB | 2026-10-26 | `sha256:de4fe90a9884` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:de4fe90a98845d75910b9769662465c84de1a934665c431279d1382b2bddc02b` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.320 | - | 0 | 0 | 43 | 855.0 MB | 2026-10-26 | `sha256:9a905addac4d` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:9a905addac4d4d3bfdd28185030f874251614559b3a038200849619710b3136b` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.427 | - | 0 | 16 | 60 | 854.0 MB | 2026-10-26 | `sha256:42eb9e70b190` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:42eb9e70b190d0cca893d93724e67fb89e7570e4309c4b1c20b7f6016a90202d` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-26 | `sha256:42d595b0d9b5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:42d595b0d9b51f8c200e12129708764fb7531eff9b6b87e832b9be68190a502a` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 69 | 339 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-10-26 | `sha256:c20b32d4e946` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:c20b32d4e946eceaf3cbfe9a29ee75cc08f8597c795b4c1b183e5eded3f0d27b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-26 | `sha256:d3215d9ab2fc` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:d3215d9ab2fc5f2ee7e3a77e83c9750f18b70a735a4514f5b2d3c5861bce4905` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 1 | 1 | 324.0 MB | 2026-10-26 | `sha256:60e50e38ea19` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:60e50e38ea197c28f7f1df97fdfe7078f1c3d4e76851483161e0e176b127f476` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 1 | 1 | 327.0 MB | 2026-10-26 | `sha256:c1f579232d45` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:c1f579232d45eb76a3a53a935a149cacf07b9521cf1a31f8ac87a793e251f666` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 1 | 1 | 355.0 MB | 2026-10-26 | `sha256:95d38c6bda2b` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:95d38c6bda2ba9dd69b564094f81b3e5911a0f30d2cd45d43d094297772f091d` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 1 | 1 | 400.0 MB | 2026-10-26 | `sha256:9141ae68beff` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:9141ae68beff7a1d603c0f8d6a0d1bc0b16702ae9821a82f095c18a0775c324f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 1 | 91 | 432.0 MB | 2026-10-26 | `sha256:1e3d6f51cdae` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:1e3d6f51cdaee953344f406433150ede04acb0be84455e5bacb1c2c0c3380bf2` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 1 | 91 | 459.0 MB | 2026-10-26 | `sha256:028abee4fba0` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:028abee4fba04b2100db738436c65582b358b63811e9140afb9d52cd9463df25` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 1 | 91 | 503.0 MB | 2026-10-26 | `sha256:955b399669af` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:955b399669af0064f678077401c798f0f8b4285682f4f24e3c33fe208c0252b7` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.1 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-10-22 | `sha256:4d63e9b5fedc` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:4d63e9b5fedc1aab072f67e30d59586eee3eeda0fae160338398be7d80997878` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | - | 0 | 3 | 9 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.1 | :24-nonroot | 0 | 5 | 15 | 159.0 MB | 2026-10-22 | `sha256:c61fc93be756` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:c61fc93be756a3018184886a353439cb2aee68962a572db6c65a09108027b27b` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.1 | :24 | 0 | 5 | 15 | 159.0 MB | 2026-10-22 | `sha256:b722a77359d4` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:b722a77359d4722f865341e63d36690ce6eb8356d171d8da7f7b8ca962af7ccb` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 0 | 12 | 33 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 0 | 12 | 33 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 20 | 80 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 22 | 99 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 25 | 92 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 25 | 92 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-26 | `sha256:53bcb6cfbebe` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:53bcb6cfbebe225e3de8d46755e70ec77fecf061c1a9fd38a2e2256563ebdbb7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-26 | `sha256:d10dcddeb70d` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:d10dcddeb70ddad7a82276334477aba8e8ee2b10126593e5a0144942d9c9404f` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-26 | `sha256:4c138f8a3c43` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:4c138f8a3c4399006d928c3942a928a19f63390ac1dac91e9a3ebb667dbb0e84` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-10-22 | `sha256:c91e907c537a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:c91e907c537ad861cb90bc18aece90479b0754b090afd4eb8722dd15477dc9e8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-22 | `sha256:6b85a86416a2` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:6b85a86416a29cdcfd404e9fbd9c5ef49cc0612882cecc702535dcac39c9446c` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-10-22 | `sha256:bd88359666e3` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:bd88359666e3f20296e7fe639a480dfd3c3c8eba299b21f26e4560353e21a656` |
