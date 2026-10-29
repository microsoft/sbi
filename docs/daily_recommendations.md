# Daily Recommended Images by Language

_Generated: 2026-10-29T02:24:50Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 102.0 MB | 2026-10-24 | `sha256:c56bc7059dc5` | `mcr.microsoft.com/dotnet/runtime:8.0-azurelinux3.0-distroless@sha256:c56bc7059dc57a5a9b4dfac7317f71f13423d2449cbaae311cb169fe2cc2ddee` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-24 | `sha256:c9af2430171d` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:c9af2430171d7ca2252372f8656e134aa62cf599c227a0773ca13239cd30219f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-24 | `sha256:20246a3c6190` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:20246a3c6190e9ffac89625c300978163c32b1682d2b629556c32368a9046885` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.33 | - | 0 | 0 | 0 | 126.0 MB | 2026-10-24 | `sha256:e27d09562205` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e27d0956220534c06c533469fe325fdf6a3c0b49e20d94e26b877fa3d1a8f88a` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.22 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-24 | `sha256:844c5d3a1c0d` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:844c5d3a1c0dc772f568ce0306b1551ead353427494bb967075710bb4ec65dc1` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.14 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-24 | `sha256:45908708ba21` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:45908708ba217e57515ab3ba03bfb1eeb5de75cf22a9e196f809a18fdfa13a01` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.403 | - | 2 | 0 | 2 | 963.0 MB | 2026-10-24 | `sha256:f0414bcafb27` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:f0414bcafb2751af8dcc7f72f6055eb456ca9931074b2eb837c6c9d166360fb5` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.320 | - | 2 | 0 | 12 | 900.0 MB | 2026-10-24 | `sha256:41fe79a756a3` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:41fe79a756a3ca3aa1cb79ce4f84e399b21ea871e38795fc7afcac02bf71f431` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0` | 8.0.427 | - | 2 | 16 | 29 | 899.0 MB | 2026-10-24 | `sha256:3763c3691c72` | `mcr.microsoft.com/dotnet/sdk:8.0-azurelinux3.0@sha256:3763c3691c726fb90d60feaec899af309e4035705bec6b84f34a9cad6d8f8d59` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.33 | - | 1 | 56 | 230 | 194.0 MB | 2026-10-28 | `sha256:67e16b7d5682` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:67e16b7d568208713e05db7f7bcc81827679a636222160bf57bb148acd23622e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.33 | - | 1 | 56 | 230 | 218.0 MB | 2026-10-28 | `sha256:ea721098d6ea` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:ea721098d6ea2894ff9bd8387821a1659d4765d6c014b8d648bb1da16f016f5a` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.427 | - | 1 | 93 | 443 | 868.0 MB | 2026-10-28 | `sha256:6ba9644da07b` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:6ba9644da07bb5b560a402acec9ca38711a0762969f1b61da261df5c9f6f85da` |

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
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 324.0 MB | 2026-10-28 | `sha256:29c7dbaaebb6` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:29c7dbaaebb6e7cd89d7d3b4da21e9d6965ad96423891c81b9dee5b5c0141ffc` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 0 | 0 | 327.0 MB | 2026-10-28 | `sha256:e97a1b1b321a` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:e97a1b1b321a716eb16b007c124ae48c73787cf5b0fb49a546681fa1eecacee7` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 0 | 0 | 355.0 MB | 2026-10-28 | `sha256:39f167754b7b` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:39f167754b7b474756f2ec45317da4e7acaa80584b2969661516d4f86f63c78e` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 0 | 0 | 400.0 MB | 2026-10-28 | `sha256:3bc54f31c0f1` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:3bc54f31c0f10b715186a1a406e11bbd96393ba7fec0136d863e8be1fab2b331` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-10-28 | `sha256:75276b45d6ab` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:75276b45d6abf3e5345613399374d49e7b9f6c9dad67a97b3bfc31a2d8c489a9` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-28 | `sha256:963c4eaec8d9` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:963c4eaec8d9bd5d97b560c86967a1f32274420f01964f17a83293e9a32dfec6` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 0 | 89 | 430.0 MB | 2026-10-28 | `sha256:b895056a8b14` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:b895056a8b1478fc65bcfbdc618140e00ef9ef32215d269e693e88ce713d66b8` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 0 | 89 | 458.0 MB | 2026-10-28 | `sha256:083b83667c42` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:083b83667c425da82c26bb7476ae0faa51ebc8d95a80bea493176888f4a2e2c3` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 0 | 89 | 502.0 MB | 2026-10-28 | `sha256:72fa7b7c214b` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:72fa7b7c214b24dd65df71d58b82da71e166e8ae6b59cccba5359e053f872a0a` |

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
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:2271cda6e66c` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:2271cda6e66cb1a2938e4c8e8e5aad47008c9e03956c98ed9cbbd37d7092216b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-28 | `sha256:7858851aed77` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:7858851aed77194b4426d4e3d46bec5e2c118cbc74ad2e95f18e234b1f32156c` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-28 | `sha256:b1142b387666` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:b1142b387666a49adb1ca6707e39029e8ad229d04494abae23d16945008c31b6` |

## Rust

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/rust:1` | 1.98.1 | - | 0 | 0 | 0 | 1.74 GB | 2026-10-28 | `sha256:bd589f56a574` | `mcr.microsoft.com/azurelinux/base/rust:1@sha256:bd589f56a574c35f889b44af3387ad852127b611e3aff7717307b1ef2d8b6172` |

## Base / No Runtime

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.9 MB | 2026-10-28 | `sha256:ca4f1b9b5d4b` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:ca4f1b9b5d4bb21445fbf5cf7288063ea477a9fa88affa4b13757245e3214914` |
| 2 | `mcr.microsoft.com/azurelinux/busybox:1.36` | 3.0 | - | 0 | 0 | 0 | 14.1 MB | 2026-10-28 | `sha256:19ffa1dd7adf` | `mcr.microsoft.com/azurelinux/busybox:1.36@sha256:19ffa1dd7adfb8644954eff77ce03a1e6617e710a8137b2fccfcb434db774314` |
| 3 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless` | 3.0 | - | 0 | 0 | 0 | 29.7 MB | 2026-10-24 | `sha256:e87da9b71656` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-azurelinux3.0-distroless@sha256:e87da9b716561f8da6480292e450c8f5bbafec725e905df5da9e5f6a31d5e2eb` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-28 | `sha256:e95286512259` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:e952865122598043cfb54f0257f016e28e02c21904e172b47dcb2148e095693a` |
| 5 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-10-28 | `sha256:88fc2b4a4258` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:88fc2b4a425801b3b4d8fe5342e24fb6134669fb718ec970340bdd36dee50c35` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled` | 24.04 | - | 0 | 0 | 2 | 14.2 MB | 2026-10-24 | `sha256:4309626281bf` | `mcr.microsoft.com/dotnet/runtime-deps:10.0-noble-chiseled@sha256:4309626281bfe550c0f6bb4242c435cd5cdabed1d05a5bc5843f659336f53296` |
