# Daily Recommended Images by Language

_Generated: 2026-08-08T02:48:56Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.18 | - | 0 | 0 | 0 | 107.0 MB | 2026-07-23 | `sha256:6c2313d68fa1` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:6c2313d68fa19680d31c0a89b95245fbafba33af8438dc4504ba16f5927c7701` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.10 | - | 0 | 0 | 0 | 112.0 MB | 2026-07-23 | `sha256:2bf17be77419` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:2bf17be77419916c2354b20a474c314adcdc85dcbdda42829c577d45bede75aa` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.29 | - | 0 | 0 | 0 | 126.0 MB | 2026-07-23 | `sha256:bdebaa5c36e8` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:bdebaa5c36e8e307d57b7ce3b5c7dd81c0313033578dd79e1d99e61f07520f51` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.18 | - | 0 | 0 | 0 | 132.0 MB | 2026-07-23 | `sha256:9763271662f4` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:9763271662f44b5c694ba050998eda81e93003e752836c89164e5364f64b8915` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.10 | - | 0 | 0 | 0 | 139.0 MB | 2026-07-23 | `sha256:d61165fe0f81` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:d61165fe0f812193589d83d666dec228519f9e0c4faebcaa4ed8775f22de2b42` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.302 | - | 0 | 6 | 7 | 919.0 MB | 2026-07-30 | `sha256:dbcd91573287` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:dbcd91573287583286212aaa14ae39ccf04f80818b33e2770dbbd5d948d23230` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.316 | - | 0 | 6 | 17 | 881.0 MB | 2026-07-30 | `sha256:ff1b47803c37` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ff1b47803c37c1678474d91275afaf920b1d82f2f7dcb4e3c0c208ee391c42fb` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.29 | - | 5 | 17 | 167 | 193.0 MB | 2026-08-05 | `sha256:794b603ad8cf` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:794b603ad8cf882df3a37dcb124eccdcd688c367864cd66e93e921cda9caba76` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.29 | - | 5 | 17 | 167 | 218.0 MB | 2026-08-05 | `sha256:fd7596eaea7a` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:fd7596eaea7ad453fe7ac16724a3c9ae36edcda894ba13743d6a5c83d6a3b36d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.423 | - | 17 | 70 | 373 | 850.0 MB | 2026-08-05 | `sha256:e2f26f26169f` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:e2f26f26169fd10d6f1b426e01c97397717b32e9d5ab4ee4a7d5497ed9403007` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.29 | - | 0 | 0 | 15 | 199.0 MB | 2026-08-01 | `sha256:b0642831def7` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:b0642831def7f75d1686be15c3232fadd51ea479b12c9b140e4b7a03f9511474` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.18 | - | 0 | 0 | 15 | 205.0 MB | 2026-08-01 | `sha256:1f73eeff2942` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:1f73eeff294238d87142a79554795b4b7f93195042075a098846a8459429daf2` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.10 | - | 0 | 0 | 15 | 209.0 MB | 2026-08-01 | `sha256:68d35011fe04` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:68d35011fe04a39cca38208d392ed48f2df15653633dca16dbc4582d07342b9f` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.29 | - | 0 | 0 | 15 | 223.0 MB | 2026-08-01 | `sha256:91ccba025be7` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:91ccba025be7d9ecb231cb4f1dfe463007cd0b87c9f76c7ab5730c9ae16bcabb` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.18 | - | 0 | 0 | 15 | 229.0 MB | 2026-08-01 | `sha256:60a0648fab58` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:60a0648fab58fa6e374bf091a612d70241d82f58fd454605585ab6ee14527028` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.10 | - | 0 | 0 | 15 | 236.0 MB | 2026-08-01 | `sha256:f1126d438ccc` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:f1126d438ccc359f51cc6d4701a8deae513856cf10f5fe645d29ea6403dcac6b` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.302 | - | 0 | 5 | 24 | 882.0 MB | 2026-08-01 | `sha256:72dd743782f2` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:72dd743782f2ae7e5476fd64f6a460045e3998dc862218b80e6944cba79a01b0` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.316 | - | 0 | 5 | 34 | 844.0 MB | 2026-08-01 | `sha256:a87faeffaa78` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:a87faeffaa7890ea1ae97e72e327208f01ec810401ec5f5d917d5ea630aaed5c` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.423 | - | 0 | 16 | 45 | 842.0 MB | 2026-08-01 | `sha256:72b30253425d` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:72b30253425d2707ea1dda364477136003586a9bdab63a988a84d1710f940d35` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.12 | - | 0 | 0 | 0 | 810.0 MB | 2026-08-05 | `sha256:143f9e785ddb` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:143f9e785ddb8f1091a420a96de35158faf076d577dd0f4c715d97ed15cd0e3d` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.5 | - | 0 | 0 | 0 | 843.0 MB | 2026-08-05 | `sha256:7eaa7ec1b6c1` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7eaa7ec1b6c116d1b914d4699ff7726189e0dd78ff29801af48b559a5922a3d6` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32 | - | 0 | 0 | 0 | 323.0 MB | 2026-08-07 | `sha256:33400c9efb80` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:33400c9efb80b49341a31696002395e2a93be6aba0fa342c95e174e69ee14632` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20 | - | 0 | 0 | 0 | 326.0 MB | 2026-08-07 | `sha256:59df8d7666c2` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:59df8d7666c2d1b91ac5b32e3afd27d15016b5909a80796e107ef2fe8925193c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12 | - | 0 | 0 | 0 | 354.0 MB | 2026-08-07 | `sha256:4c68f970b59b` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:4c68f970b59bd0adef7a4c44e100ed5709dbee40d985ac6c225c7f7dfad47de3` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4 | - | 0 | 0 | 0 | 399.0 MB | 2026-08-07 | `sha256:ecc55d0450fb` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:ecc55d0450fb4b45b02d242f51fcea558259449bc99c052a9b9493a27c1e048f` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12 | - | 0 | 0 | 0 | 478.0 MB | 2026-08-07 | `sha256:cd0f9db7c6d9` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:cd0f9db7c6d987bf8164f4703dcd664b035ddde67889334d3913a3945838d758` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4 | - | 0 | 0 | 0 | 523.0 MB | 2026-08-07 | `sha256:38f561fc8858` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:38f561fc8858d9bc33d1853c3304598600bdb52f800b32987cd471137275cd10` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20 | - | 0 | 0 | 75 | 429.0 MB | 2026-08-07 | `sha256:bdcad55ae91f` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:bdcad55ae91f28637a9e7de87c121cb0e363b0d6d79542e6014828c2b331c22c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12 | - | 0 | 0 | 75 | 457.0 MB | 2026-08-07 | `sha256:570bc01ebda1` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:570bc01ebda1e7de6b2f362b56fb37a8fc53199c4bc567fa2f1d0e8d3051108a` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4 | - | 0 | 0 | 75 | 501.0 MB | 2026-08-07 | `sha256:59a1113820e3` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:59a1113820e3e0c82a7e04e96c2f91715c1a31040097e3375da46f40550b0696` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.7 MB | 2026-07-31 | `sha256:81accc5dfe9a` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:81accc5dfe9a015fd026bda4856d3073cdc0928e54ceb8ce1df85c306d5a5ca7` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-07-31 | `sha256:8b486e043e64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:8b486e043e646e69a9bfc25a404909f1036b8af386a9e1d410f90936d923fb23` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 1 | 2 | 139.0 MB | 2026-07-31 | `sha256:66097f09fcca` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:66097f09fcca9f87dfa6481938c1fa4a8b81ea5d6319f1129c610367f1436c06` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | :24 | 0 | 1 | 2 | 197.0 MB | 2026-07-31 | `sha256:6523505968bf` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:6523505968bf37b0f97975cedfd202a3dbaee44ed653d17890d36deead3828cc` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 8 | 21 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | :24-nonroot | 1 | 6 | 19 | 157.0 MB | 2026-07-31 | `sha256:ad7995197668` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ad7995197668d617eb71bb198d90d1d95049f6f43028b27c0a694cddb430c073` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | :24 | 1 | 6 | 19 | 157.0 MB | 2026-07-31 | `sha256:8652f9953ce3` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:8652f9953ce393d2018727371948afda3fc085aaceb07b2ade8f8b1c112e1b1c` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 10 | 37 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 38 | 110 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 38 | 162 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 44 | 163 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-07-06 | `sha256:576d9769c014` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:576d9769c0146cbf0cf7946bacf536c5758464c29eadfa03ef5090ae708e641f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-07-22 | `sha256:178f25fadf46` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:178f25fadf466549d31e234b3091bf815161159f2f2bc98720bbf39f7368aff4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 1 | 2 | 75.4 MB | 2026-07-29 | `sha256:4ecd6b297db8` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:4ecd6b297db85c54ec2df07145a28536c3655a3e98e54eb2364189bc4e6eac23` |
