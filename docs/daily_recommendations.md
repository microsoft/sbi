# Daily Recommended Images by Language

_Generated: 2026-10-31T02:18:57Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine` | 10.0.14 | - | 0 | 0 | 0 | 121.0 MB | 2026-10-24 | `sha256:254fe0174959` | `mcr.microsoft.com/dotnet/aspnet:10.0-alpine@sha256:254fe017495996d3abcc487ffacdcce2b73affb380d3af87af12c8dff41fe0b0` |

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 102.0 MB | 2026-10-29 | `sha256:68c4eda2e145` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:68c4eda2e1459e4284a5ea593cb131f264d624ac0bbd8b31671caaeeee13b707` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-29 | `sha256:369fd8d3246f` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:369fd8d3246f76695dbd5a3ed521982d099a8c97d64eba40f6ad9de704c8d7b4` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-29 | `sha256:9bf51dbfda85` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:9bf51dbfda85c9289bca3606d2c79c11a09457dbf0f2ed4f23a9e2b9b19dde4b` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 127.0 MB | 2026-10-29 | `sha256:e18779f67a33` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e18779f67a33317667fdb5e375fc5f9f9b79ed43710f01f920ad123e3ec68986` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-29 | `sha256:1c77a718863a` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:1c77a718863a5ac26cb17d125fecef8bba1f3f2bb25e19062f6bce8619a6458c` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-29 | `sha256:79cb8e7d5437` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:79cb8e7d543740ccdae265ad2e98e4939946f8f43b91efbcef31bdb94b234cf1` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.403 | - | 0 | 0 | 0 | 951.0 MB | 2026-10-29 | `sha256:ffd76bb46785` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:ffd76bb46785445d2d029809009d35b522619b97df88c0abaa05d549db697b43` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.320 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-29 | `sha256:17e7ceadf4df` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:17e7ceadf4dfab25b6d5fbbcd516320c3950b11365f5e17d808186e3364ee082` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.427 | - | 0 | 16 | 27 | 899.0 MB | 2026-10-29 | `sha256:224ae6ffdac4` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:224ae6ffdac4888bb88bbc8a9d0d482b9fabf6a764f51b4aa6e771a6985d051e` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.33 | - | 1 | 56 | 231 | 194.0 MB | 2026-10-28 | `sha256:67e16b7d5682` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:67e16b7d568208713e05db7f7bcc81827679a636222160bf57bb148acd23622e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.33 | - | 1 | 56 | 231 | 218.0 MB | 2026-10-28 | `sha256:ea721098d6ea` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:ea721098d6ea2894ff9bd8387821a1659d4765d6c014b8d648bb1da16f016f5a` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.427 | - | 1 | 93 | 444 | 868.0 MB | 2026-10-28 | `sha256:6ba9644da07b` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:6ba9644da07bb5b560a402acec9ca38711a0762969f1b61da261df5c9f6f85da` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled` | 10.0.14 | - | 0 | 0 | 2 | 124.0 MB | 2026-10-24 | `sha256:c89a6a61f494` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled@sha256:c89a6a61f494f20a5ae302990e814627a8d91b332fd4a8e7c1707c2d4da9555d` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra` | 10.0.14 | - | 0 | 0 | 3 | 162.0 MB | 2026-10-24 | `sha256:9b755ea799ae` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra@sha256:9b755ea799ae649141a073e30d8b62d8095a5465af2b7694d17866e8c05131d8` |
| 3 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.33 | - | 0 | 0 | 26 | 193.0 MB | 2026-10-26 | `sha256:2ba6f26ec6dc` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:2ba6f26ec6dc479aa7f8016e1b0bd80797c52104872ba97766d0b160bb97d8f6` |
| 4 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.22 | - | 0 | 0 | 26 | 198.0 MB | 2026-10-26 | `sha256:3562dcf10791` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:3562dcf10791d9b1a20097500da37dc5390a1a6c59f56e40806f552cdd7c1666` |
| 5 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.14 | - | 0 | 0 | 26 | 203.0 MB | 2026-10-26 | `sha256:a3f9f81954b9` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:a3f9f81954b9c810af2ed18849716c8dc74a4707bd0907fdbd69be81554c0f30` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.33 | - | 0 | 0 | 26 | 217.0 MB | 2026-10-26 | `sha256:d0a0fa004862` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:d0a0fa004862a9cca83208800f706dcc2523c4db29f09d501395f80969e03070` |
| 7 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.22 | - | 0 | 0 | 26 | 223.0 MB | 2026-10-26 | `sha256:b7459513d45f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:b7459513d45fa48ac9a1ca2a191c720aa15b8a6a1b961a1ab2c6965816afe7d3` |
| 8 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.14 | - | 0 | 0 | 26 | 230.0 MB | 2026-10-26 | `sha256:de0ee358cdd6` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:de0ee358cdd672f7017a003bed21e9ac9ed48a27c922a59b18175b6e82fa7fee` |
| 9 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.403 | - | 0 | 1 | 33 | 918.0 MB | 2026-10-26 | `sha256:de4fe90a9884` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:de4fe90a98845d75910b9769662465c84de1a934665c431279d1382b2bddc02b` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.320 | - | 0 | 1 | 43 | 855.0 MB | 2026-10-26 | `sha256:9a905addac4d` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:9a905addac4d4d3bfdd28185030f874251614559b3a038200849619710b3136b` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-28 | `sha256:da5a59d8b746` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:da5a59d8b7462abfdb98bdf47ec8bfb677d9f1cf8afac71dce9a91cbcecd003b` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 6 | 69 | 342 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 324.0 MB | 2026-10-30 | `sha256:e51ee84d0f14` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:e51ee84d0f1464c396426a8ac5f52769d615f234fac596036f220e6320196cbb` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-30 | `sha256:cb00e538f98f` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:cb00e538f98fbf293750ae48d3f31620b72c4d2764adbc1a7558e2f878ce6f70` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 0 | 0 | 355.0 MB | 2026-10-30 | `sha256:914c517ed746` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:914c517ed7468cbfb9537164e3371a9f99af09d3ba11defff566faf1c4ff747c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 0 | 0 | 400.0 MB | 2026-10-30 | `sha256:5f334cdb0bd1` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:5f334cdb0bd16e7a31ad69d4280031d8b95d2cf2d52399990d64237312cda2b5` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-10-30 | `sha256:cb4b231c08da` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:cb4b231c08da6eee123162eb24e7a0b60477628c87c6da26c96b8c13552e701b` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-30 | `sha256:61002c0b9ccd` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:61002c0b9ccdd84295342d305e4cf6f956d205cf1cd2eabc2f9e376aff1e70f6` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 0 | 89 | 430.0 MB | 2026-10-30 | `sha256:19714fba16f6` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:19714fba16f608f6234d2de483fbab6d27e7b4d0a5ef286e6afd660d764a3433` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 0 | 89 | 458.0 MB | 2026-10-30 | `sha256:71ca7d9ec1d1` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:71ca7d9ec1d116f5e421b3749a617184e54e9cdcfc72301760495d07fbb54a76` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 0 | 89 | 502.0 MB | 2026-10-30 | `sha256:c9b2c1de3aa9` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:c9b2c1de3aa92dd19378f03cb0cc3d99d1a21ae022391939f8c843f4e2e93c57` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.22` | 24.22.1 | :24 | 0 | 0 | 0 | 200.0 MB | 2026-10-28 | `sha256:affe8e33adb5` | `mcr.microsoft.com/azurelinux/base/nodejs:24.22@sha256:affe8e33adb5aa9039d6bf37156fd07d0eb8bb3ace44bd178835484dc2911b5f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot` | 24.22.1 | :24-nonroot | 0 | 6 | 17 | 159.0 MB | 2026-10-28 | `sha256:546ae02c9a35` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22-nonroot@sha256:546ae02c9a3537e0a322fe9e594da019350273b82b7e56929bdb32b5482e33cd` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22` | 24.22.1 | :24 | 0 | 6 | 17 | 159.0 MB | 2026-10-28 | `sha256:4997e0b8d0dd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.22@sha256:4997e0b8d0dd2058d7fd7fc3023dde06913f1d164e043ed147cb03620ef531af` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 3 | 61 | 267 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | - | 4 | 12 | 37 | 199.0 MB | 2026-10-16 | `sha256:74d685815e58` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:74d685815e58a5bce05088452adfccd3091368511236fbff1c7ad328b4a21b63` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | - | 4 | 22 | 63 | 158.0 MB | 2026-10-16 | `sha256:c1f9ccc08d03` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:c1f9ccc08d03674c7644bd395e82584e0d16ce65bbb15a26eae4ea02239eec57` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | - | 4 | 22 | 63 | 158.0 MB | 2026-10-16 | `sha256:802db036c3a8` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:802db036c3a84ddd010267727c85b582b61a378fe538368de6c611506b7475ec` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 4 | 29 | 108 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 4 | 31 | 127 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 4 | 35 | 122 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.16 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:9c735f30f9d2` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:9c735f30f9d2344b4af3f8a2f715587ca1cd1afeab8c3c9836ac69d8326637b7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:10bcfa3c563f` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:10bcfa3c563fd05f9b11495551bd36bfe919b4a2f32b68388e1c393bbeea64ac` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.16 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-28 | `sha256:71a1f5a1821f` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:71a1f5a1821f4c56e45aff52e792ed6c0500631e201cfc140c4c9d22886b29f9` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-10-28 | `sha256:d554fbb204d8` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:d554fbb204d8330fd3a64234a07e4f0a6b51c0c2ac2ef48b000ea74a25e9cf8d` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.9 MB | 2026-10-28 | `sha256:ca4f1b9b5d4b` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:ca4f1b9b5d4bb21445fbf5cf7288063ea477a9fa88affa4b13757245e3214914` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-10-28 | `sha256:19ffa1dd7adf` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:19ffa1dd7adfb8644954eff77ce03a1e6617e710a8137b2fccfcb434db774314` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.8 MB | 2026-10-29 | `sha256:d523c75e720b` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:d523c75e720b64058b529dbbb1b6ec4023b3e54f3b83583fef4de4ed73925922` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-28 | `sha256:e95286512259` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:e952865122598043cfb54f0257f016e28e02c21904e172b47dcb2148e095693a` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-10-28 | `sha256:88fc2b4a4258` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:88fc2b4a425801b3b4d8fe5342e24fb6134669fb718ec970340bdd36dee50c35` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 2 | 14.2 MB | 2026-10-24 | `sha256:4309626281bf` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:4309626281bfe550c0f6bb4242c435cd5cdabed1d05a5bc5843f659336f53296` |
