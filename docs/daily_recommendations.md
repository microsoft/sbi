# Daily Recommended Images by Language

_Generated: 2026-08-16T02:40:29Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 108.0 MB | 2026-08-13 | `sha256:56f96c0dfa21` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:56f96c0dfa21cf170ff402ff75de39efe89c4fee9b92c0a8fb07ea751b621d45` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 112.0 MB | 2026-08-13 | `sha256:3d7b8851fa22` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:3d7b8851fa229f31031fb9c6787b5d38263ea9a6bf88f97bf17f32f446c9b9b0` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.30 | - | 0 | 0 | 0 | 126.0 MB | 2026-08-13 | `sha256:8957983e8929` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:8957983e8929735a20b978ea82bb3074d79c2a1410a10a70175b6c4d9fdd003a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.19 | - | 0 | 0 | 0 | 132.0 MB | 2026-08-13 | `sha256:c412debe88b0` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:c412debe88b0360c3fb477c0f49894295176f924bfcf869c0bce4aa085e501e2` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.11 | - | 0 | 0 | 0 | 139.0 MB | 2026-08-13 | `sha256:c2fbd1565145` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:c2fbd15651452871430ce3fbd0083c128b376dbac786797be1c7dd4091d20a23` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.317 | - | 0 | 0 | 10 | 881.0 MB | 2026-08-13 | `sha256:c9c558d6730a` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:c9c558d6730a709144f578823ac0a3934eb7715bc8e18c5c1e5072bf42b3e0de` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.400 | - | 0 | 5 | 5 | 931.0 MB | 2026-08-13 | `sha256:d1348cda1563` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:d1348cda1563f5dc1a1d419023a2067389a3dbe7a1e0686913d9839f85053f00` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.30 | - | 5 | 17 | 173 | 193.0 MB | 2026-08-10 | `sha256:a27922704a14` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:a27922704a14d6d1ce103697aa3efa885902db33651b41a84ad1594fdf5be66e` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.30 | - | 5 | 17 | 173 | 218.0 MB | 2026-08-10 | `sha256:b0beb9cc1dee` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:b0beb9cc1dee1c1b0749796110d4734292071b814207ad0d4f40611f7db04f7b` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.424 | - | 17 | 64 | 379 | 850.0 MB | 2026-08-10 | `sha256:306301580fca` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:306301580fcaa5b445180e759db59309979002d1000669cb4cf58a567d0014bc` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.30 | - | 0 | 0 | 19 | 199.0 MB | 2026-08-10 | `sha256:6a585d286cb7` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:6a585d286cb758eeeba86166bb28beaf1da3de35ed0f47f5850ff9ee750d0f43` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.19 | - | 0 | 0 | 19 | 205.0 MB | 2026-08-10 | `sha256:8e8a639a9806` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:8e8a639a98063440299b397e36e0520d24457483badd4f34c133bcb5ece6e162` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.11 | - | 0 | 0 | 19 | 209.0 MB | 2026-08-10 | `sha256:acad02eb5c4f` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:acad02eb5c4fbf57d15296f9c08d56cd4036e915bdae5b4dd48a06523d452617` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.30 | - | 0 | 0 | 19 | 223.0 MB | 2026-08-10 | `sha256:2b27185ab3f3` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:2b27185ab3f3fd7ee77d9c9515f333758ff34b46f2ab507b36535e8693ae3e5c` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.19 | - | 0 | 0 | 19 | 229.0 MB | 2026-08-10 | `sha256:546d3a96069b` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:546d3a96069bd7442ec3e9d26b245dfca58fc0824a3dc277312db0fd99a412fc` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.11 | - | 0 | 0 | 19 | 236.0 MB | 2026-08-10 | `sha256:207cc5149677` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:207cc51496778557731c81ff670333d8ade4a4fec22768fd1be8e78474a84ecf` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.317 | - | 0 | 0 | 33 | 844.0 MB | 2026-08-10 | `sha256:840c88158a49` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:840c88158a49d65ab59451369988a0db6b420ab3ea7948a5e261c56a4ab3404d` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.400 | - | 0 | 5 | 28 | 894.0 MB | 2026-08-10 | `sha256:e1fc6e423f54` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e1fc6e423f543119c406d24e2e687d67c569f18f04a37a8b0005d80ad0dcee80` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.424 | - | 0 | 11 | 44 | 842.0 MB | 2026-08-10 | `sha256:f7b81c81fde9` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:f7b81c81fde99517f5ac7ff48188500782f9bff210bf3e8648a6ff75f651c43f` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.13 | - | 0 | 0 | 0 | 809.0 MB | 2026-08-14 | `sha256:7b9b44f1f025` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:7b9b44f1f02557a111c6cfa8c7df1a7ea8bcd7ecf9424a49129f0d9cf4956547` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.6 | - | 0 | 0 | 0 | 842.0 MB | 2026-08-14 | `sha256:0d009d1cb486` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:0d009d1cb486a1a23c35dcfee3fbd82e22adf9f7ef0cf6bb35020679621d378f` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-14 | `sha256:c10eb5aeaafb` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:c10eb5aeaafb2ed9f856073cedeb4c43f2577259221f499e257dd1ef75abafbd` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4 | - | 0 | 0 | 0 | 522.0 MB | 2026-08-14 | `sha256:afbfadded27b` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:afbfadded27b5cc34be616a73bb6b740589b8a5ea6e509e53aa385d9ac641f80` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32 | - | 0 | 8 | 8 | 323.0 MB | 2026-08-14 | `sha256:81d4073f3637` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:81d4073f36376a24427233886d5505e08bb6da4aa117952122400ad9d97405fc` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20 | - | 0 | 8 | 8 | 326.0 MB | 2026-08-14 | `sha256:4f5db737903b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:4f5db737903b584341db699496fe9bbe6bfeee645c865ca4bc5ce8f785307d20` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12 | - | 0 | 8 | 8 | 354.0 MB | 2026-08-14 | `sha256:154f86e8262c` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:154f86e8262c2378acbc6720eb050910d87356f757dad048295f563f543e799b` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4 | - | 0 | 8 | 8 | 399.0 MB | 2026-08-14 | `sha256:b9e55f98b462` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:b9e55f98b46221fd40eb22ade9d0918764f9f765b3234bee92ee9bdea3ba5938` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20 | - | 0 | 8 | 83 | 430.0 MB | 2026-08-14 | `sha256:5edf663c3ca6` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:5edf663c3ca6a33d61ebf9d35a53fd6475eb1f61d7863bc13bc2a71a49442924` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12 | - | 0 | 8 | 83 | 458.0 MB | 2026-08-14 | `sha256:b94a615a6484` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:b94a615a64842966dbea0fa57c5244c90bca57b7a3f36b5060973e777adbce59` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4 | - | 0 | 8 | 83 | 502.0 MB | 2026-08-14 | `sha256:57207a5b2793` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:57207a5b2793adc7d60ee3f3a502d03c2d4fbec701535d002b7b6cc3d605eede` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 0 | 0 | 197.0 MB | 2026-08-11 | `sha256:e4578b763a16` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:e4578b763a1698192f53e8de82175f0fa315946acac6d24e49ea064b1387215b` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 8 | 24 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 6 | 19 | 157.0 MB | 2026-08-11 | `sha256:58202c56c1d2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:58202c56c1d28e3ad3962143c9d2ec55c190b42cc19d1918d5a57b205fe7af6f` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 6 | 19 | 157.0 MB | 2026-08-11 | `sha256:bb5ac511e1f9` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:bb5ac511e1f92dcc192a3414ccd6941b1044111db7d7ee6ee4b2bddf76d90257` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 38 | 165 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 44 | 166 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.8 MB | 2026-08-11 | `sha256:6f12a7ae467f` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:6f12a7ae467f01f7f4f36893fb484813bf7a0cce555da9b20cc78938b6b27164` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.8 MB | 2026-08-11 | `sha256:8f0a5a7be5bb` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:8f0a5a7be5bb57f6605ffde72e5aff701d0c75d34a6797da3bfcf64140d63b8d` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-08-11 | `sha256:722b6224c23b` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:722b6224c23b3f21f5268e2073f80c0f396bc626e3193b6dbf66e40d89478f03` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.5 MB | 2026-08-11 | `sha256:8bb51342bd5e` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:8bb51342bd5eba915990ab608f91060d502bb7891a2d3d909e0419b932533029` |
