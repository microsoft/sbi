# Daily Recommended Images by Language

_Generated: 2026-07-31T03:13:16Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.302 | - | 0 | 5 | 5 | 919.0 MB | 2026-07-30 | `sha256:dbcd91573287` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:dbcd91573287583286212aaa14ae39ccf04f80818b33e2770dbbd5d948d23230` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.316 | - | 0 | 5 | 15 | 881.0 MB | 2026-07-30 | `sha256:ff1b47803c37` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ff1b47803c37c1678474d91275afaf920b1d82f2f7dcb4e3c0c208ee391c42fb` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.29 | - | 5 | 17 | 164 | 193.0 MB | 2026-07-24 | `sha256:4d0117b33547` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:4d0117b335477ad6c35cfa70c75424e72bbf05645a5d77c1f68a74fdd057676a` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.29 | - | 5 | 17 | 164 | 218.0 MB | 2026-07-24 | `sha256:8f6a307ae32f` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:8f6a307ae32fb393a4b4bcde0a81f0ce3f0a715de0e5575df71f3030448f2dde` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.423 | - | 17 | 71 | 370 | 850.0 MB | 2026-07-24 | `sha256:3c0edbfe1549` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:3c0edbfe1549dd93fb789dc96299a40df865ad7bffefcaf38e8c05940686d641` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.29 | - | 0 | 0 | 43 | 193.0 MB | 2026-07-14 | `sha256:0f038d362c97` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:0f038d362c97c1d34ce2481891e7dadfcbb67a014fed7e010ba2ef5bbb9c5820` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.18 | - | 0 | 0 | 43 | 198.0 MB | 2026-07-14 | `sha256:156e9fd13513` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:156e9fd1351359ac2b8cd3e05676de78bfb1a8937f9b221b50f9ccf6984c7093` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.10 | - | 0 | 0 | 43 | 203.0 MB | 2026-07-14 | `sha256:ed5d539b2784` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:ed5d539b27842d656a06a5984dbcb5114d3e885fbada612a49a5a7c3c3a44e1c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.29 | - | 0 | 0 | 43 | 217.0 MB | 2026-07-14 | `sha256:15f696229ddd` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:15f696229dddb7e68aae3fa9551f1af9093d7cdda1803834fdf722013383014c` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.18 | - | 0 | 0 | 43 | 223.0 MB | 2026-07-14 | `sha256:5ad129827db5` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:5ad129827db5bf710906a43db07f4f1e2209bd2a8e3403ffea1ffe6961a82dc1` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.10 | - | 0 | 0 | 43 | 230.0 MB | 2026-07-14 | `sha256:1fa23fc4872d` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:1fa23fc4872d95fd71c2833ebe65d7e84a43b2d51a31d119516852f13d9505a7` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.302 | - | 0 | 5 | 67 | 883.0 MB | 2026-07-14 | `sha256:ed034a8bf0b2` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ed034a8bf0b24ded0cbbac07e17825d8e9ebfe21e308191d0f7421eaf5ad4664` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.316 | - | 0 | 5 | 77 | 845.0 MB | 2026-07-14 | `sha256:72b2c1fba104` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:72b2c1fba104eed0765e76c66256dd57b8b00c5e7c7fd16ad3eb254ad18db3fc` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.423 | - | 0 | 16 | 88 | 844.0 MB | 2026-07-14 | `sha256:283164eecee1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:283164eecee1fc80a590410d3c56207af455bd51b8bf3cbf2ded3592b9014b29` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32 | - | 0 | 0 | 0 | 323.0 MB | 2026-07-29 | `sha256:6bb30a8150c9` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6bb30a8150c93731eb50fa180c45151bd042d831cb1a5e1aa4109bf81dbba71e` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20 | - | 0 | 0 | 0 | 326.0 MB | 2026-07-27 | `sha256:fa7a8b3dbf76` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:fa7a8b3dbf76f07334cd9a466164e4c86e53c5382854965f5e3f568054dff06d` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12 | - | 0 | 0 | 0 | 354.0 MB | 2026-07-29 | `sha256:f34554b34f97` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:f34554b34f970743644b9bb44c08f2f1670c2843c0f8126d704b426fe26269e5` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4 | - | 0 | 0 | 0 | 399.0 MB | 2026-07-29 | `sha256:990139175a7d` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:990139175a7d498319e50ccb3d0bada3730a55d2e11bb22076e13f9f4b6ce59a` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12 | - | 0 | 4 | 4 | 477.0 MB | 2026-07-29 | `sha256:1e76ccf6882f` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:1e76ccf6882fd6d7df4c537a1bb6f519d5fa67fba4c3fa8b8f59fbcb585b283a` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4 | - | 0 | 4 | 4 | 522.0 MB | 2026-07-29 | `sha256:d51e08c026e5` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:d51e08c026e5e08e17770d1c9f07ee43dd279d48941247a38a69729fa5b60594` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20 | - | 0 | 0 | 75 | 450.0 MB | 2026-07-29 | `sha256:904a1d8dde44` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:904a1d8dde443bc1718ecfe56a453c0f880b44b848003dbed6830965fb9847ca` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12 | - | 0 | 0 | 75 | 477.0 MB | 2026-07-29 | `sha256:cb85a4d7d9c4` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:cb85a4d7d9c4a886742ad4ee79e9022401fff1721e427ed1a28c960626c5ce16` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4 | - | 0 | 0 | 75 | 522.0 MB | 2026-07-29 | `sha256:f654b767640f` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:f654b767640fef9da648e45b65243407159f8b166b5e72f40c4a28209a52d329` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.0 | :24 | 0 | 0 | 0 | 197.0 MB | 2026-07-29 | `sha256:0e75641e2e38` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:0e75641e2e38595bd04caaeb2050738dc90ca2ecebf5522705ba07ac3da7a526` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 4 | 4 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.0 | :24-nonroot | 1 | 4 | 12 | 157.0 MB | 2026-07-29 | `sha256:265a12a4c363` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:265a12a4c3631cd23ea7013f0601a88e2b4412cf9eb61cd5b4b8772a947a8355` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.0 | :24 | 1 | 4 | 12 | 157.0 MB | 2026-07-29 | `sha256:5a2fa3b1ccc3` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:5a2fa3b1ccc3c60bc21c06ca606f37058df63ef5fe6787bd61d64d1da3b2df6e` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 5 | 17 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 5 | 17 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 31 | 91 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14` | 24.14.1 | - | 1 | 31 | 91 | 153.0 MB | 2026-06-19 | `sha256:9f931b8beae1` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14@sha256:9f931b8beae125bd4863ead3308cb700e44154af264067c9804abb66965a5b26` |
| 9 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | - | 2 | 35 | 158 | 163.0 MB | 2026-04-01 | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 10 | `mcr.microsoft.com/azurelinux/base/nodejs:24.14` | 24.14.1 | - | 2 | 38 | 144 | 193.0 MB | 2026-06-19 | `sha256:2cb9bed9f0d2` | `mcr.microsoft.com/azurelinux/base/nodejs:24.14@sha256:2cb9bed9f0d2aba3d711b09da1ca62dd11ef594e0ae9b87352bb7eea34f3297c` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | :3-nonroot | 0 | 0 | 0 | 83.7 MB | 2026-07-29 | `sha256:eb1c2baadb95` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:eb1c2baadb95379102db6129406ab00deda74c2656e2cbfa28dd40c78d34c79d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 83.7 MB | 2026-07-29 | `sha256:6d35905ed696` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:6d35905ed696c91bb33dff68139fd47a5f2c43e23634a968af778e54f48bebb7` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-07-29 | `sha256:8b00ae816515` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8b00ae81651563f82c466eb7dbdb3df660962b0278c6ece65c14aa141d6612fb` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.12 | - | 0 | 4 | 4 | 809.0 MB | 2026-07-29 | `sha256:e20e3f786105` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:e20e3f78610593a499e108966bd73dba88689f68de01934af42dfccd64cbee81` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.5 | - | 0 | 4 | 4 | 841.0 MB | 2026-07-29 | `sha256:e86735636d61` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:e86735636d6188f60df5ecdb8a2fa4e55bfed6cf71ed95272004b3a53db17c25` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-07-06 | `sha256:576d9769c014` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:576d9769c0146cbf0cf7946bacf536c5758464c29eadfa03ef5090ae708e641f` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.1 MB | 2026-07-22 | `sha256:178f25fadf46` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:178f25fadf466549d31e234b3091bf815161159f2f2bc98720bbf39f7368aff4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 75.4 MB | 2026-07-29 | `sha256:4ecd6b297db8` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:4ecd6b297db85c54ec2df07145a28536c3655a3e98e54eb2364189bc4e6eac23` |
