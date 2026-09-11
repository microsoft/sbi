# Daily Recommended Images by Language

_Generated: 2026-09-11T02:16:29Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 108.0 MB | 2026-09-04 | `sha256:0aca4c2e1c24` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:0aca4c2e1c2435a74be5254599b6436773483c649acc764aa5bf2898d91025f8` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 112.0 MB | 2026-09-04 | `sha256:f1981f8923ee` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:f1981f8923ee8ecd967f002df6100b556c2ed8f7ab29df932d4625f206badb6f` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.31 | - | 0 | 0 | 0 | 126.0 MB | 2026-09-04 | `sha256:8f0aabeb4e7d` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:8f0aabeb4e7d3accdb42a69db6344962c5844dcfcdb50009a4b5b1927df5f146` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 132.0 MB | 2026-09-04 | `sha256:04acc008da2b` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:04acc008da2bb706efc2158ac1ab4f3f7f3015cfb0ad3d602eb0f996025e8e59` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 139.0 MB | 2026-09-04 | `sha256:9fd8a9a09ba5` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:9fd8a9a09ba5c0c5cb6d95c8200447de9bc46751ff77b8caa141725a29a1178c` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 0 | 70 | 961.0 MB | 2026-09-04 | `sha256:40195d80f8c2` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:40195d80f8c217195a16f65c1be7e8acad646317e9bda2801d91df681c60b3f4` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 0 | 80 | 899.0 MB | 2026-09-04 | `sha256:49e105cb9e20` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:49e105cb9e208a4562c904605afb28ea2377d6a83110c8c29fd745d446bc81dd` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 52 | 236 | 193.0 MB | 2026-09-04 | `sha256:9cfa8aaf5c98` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:9cfa8aaf5c98a4cedffec74d450dd2d4510ba72ee9e663efe2658500b0321524` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 52 | 236 | 218.0 MB | 2026-09-04 | `sha256:9a464e9a7e8c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:9a464e9a7e8c6144631020975f703c89034fe386417cb740620df69c2c6cfe24` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 101 | 446 | 867.0 MB | 2026-09-04 | `sha256:5ef85cc12cb2` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:5ef85cc12cb25be6ec319a7392d1e9efd53c3bc8abb971c53d8058a473f09053` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 0 | 11 | 193.0 MB | 2026-09-09 | `sha256:b95377a87b69` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:b95377a87b69721af927b5f7ec2957149fd69d2b73751f62c2d4f16ef52b9667` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 0 | 11 | 198.0 MB | 2026-09-09 | `sha256:3dbab8a69856` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:3dbab8a69856ebb6fb64a9a996f31dcbb7e8e5f02e825d1f8f75d0604f918268` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 0 | 11 | 203.0 MB | 2026-09-09 | `sha256:399e54a8a7e3` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:399e54a8a7e35c3aba78398b2840455d45185cba20b831b8a2b46f849f4f5001` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 0 | 11 | 217.0 MB | 2026-09-09 | `sha256:8315cdc745a9` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:8315cdc745a93bbc30b28d798c45503665fe171bd32f5ecaca59f8759909f2d8` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 0 | 11 | 223.0 MB | 2026-09-09 | `sha256:cb5d8a8a6f61` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:cb5d8a8a6f6160cbd3016d4e774236a1103f0df4d7d86749c576ea0c7805aeb8` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 0 | 11 | 230.0 MB | 2026-09-09 | `sha256:1fe86375600b` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:1fe86375600b62e6566b465da9553eef0621f13c67f40fe764cd8dbb1dee1497` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 0 | 44 | 917.0 MB | 2026-09-09 | `sha256:4ea6fe75dd36` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:4ea6fe75dd36706bb6d8c3c293d4c4315840f5d76ea28ac97def77e3ec487fa5` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 0 | 54 | 855.0 MB | 2026-09-09 | `sha256:57ff403c1235` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:57ff403c12358b5a93b464b48d06be297dbe466dcf5c71e74059f7cb5ee3c73d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 11 | 65 | 854.0 MB | 2026-09-09 | `sha256:1960d429d16a` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:1960d429d16a609e102e9d08c3fa127553d6792353ab95ebd8ec59faba7e8cbd` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-09-10 | `sha256:2eb96427f00c` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:2eb96427f00c6c74d36bae28b2487d2477feca4e88e1e05e45e01ec1048626fa` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-09-09 | `sha256:db296051bfdb` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:db296051bfdb0d017e2a6ab675dbea1a07fd98bf5a3e58987caa06a51819d465` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-09-10 | `sha256:b714f3aba01f` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:b714f3aba01f114f9ded54dcb9e53e49a062fe2eca9676f63571c8bfa0620a4c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-09-10 | `sha256:9bb039af0dc1` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:9bb039af0dc1e9b6d1b06ac1eb5bd9a85d7999033ce80b667ae129cda5111e18` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-09-10 | `sha256:177acadce661` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:177acadce6616dc9d4e7bc6eb7099767c25d59fdfe8e3407c54fca4cfe16087f` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 522.0 MB | 2026-09-10 | `sha256:1ef36daad47b` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:1ef36daad47bf774c48eb397efd81e9248e893c7d6b8682153fb248cd1e6d3b4` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 93 | 453.0 MB | 2026-09-10 | `sha256:171374eee16f` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:171374eee16f562bdb4fbc5ccfe0f597744958948f81efaff6d769daace169f7` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 93 | 480.0 MB | 2026-09-10 | `sha256:8b487a92b54b` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:8b487a92b54b478d016ce516f08e44a3322af5323153ce28be6343eb9716e590` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 93 | 525.0 MB | 2026-09-10 | `sha256:c05d7f192979` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:c05d7f1929792d6cbb63ef6783466a9faa238b6a5d661a77bd52af1cf70666b1` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | :24 | 0 | 0 | 0 | 198.0 MB | 2026-09-04 | `sha256:467d6d9ddee3` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:467d6d9ddee33b22cb8983a7317d40299a33e8dfc15f3f37e8b2013ed7084b20` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 0 | 3 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | :24-nonroot | 0 | 4 | 9 | 158.0 MB | 2026-09-04 | `sha256:9803a0a3e4a5` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:9803a0a3e4a5690e67aa1bddbbc64a55b921030efcc6ee732f761fe7ceb819d5` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | :24 | 0 | 4 | 9 | 158.0 MB | 2026-09-04 | `sha256:e7c65e70c1a2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:e7c65e70c1a259f1aca64458062001be52f0993eed6e39c3aa74273bf4812528` |
| 5 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 13 | 42 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 8 | 24 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 8 | 24 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot` | 24.17.0 | - | 1 | 12 | 42 | 156.0 MB | 2026-07-22 | `sha256:82b0b67e71ec` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17-nonroot@sha256:82b0b67e71ecc975ccae4250aab389c7ca58e2ce2531579b56833faaa9bd2b15` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17` | 24.17.0 | - | 1 | 12 | 42 | 156.0 MB | 2026-07-22 | `sha256:2ef223579c32` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.17@sha256:2ef223579c32d5f4fd065961d15c7220c3b0ea9fed50afd77061d58ab243a738` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot` | 24.14.1 | - | 1 | 43 | 117 | 153.0 MB | 2026-06-19 | `sha256:786da51aaf5c` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.14-nonroot@sha256:786da51aaf5c3056e0bf1ad4a7a1c44591bead144bc29198d7248190af5359fb` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 83.5 MB | 2026-08-25 | `sha256:d921452dba64` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:d921452dba64944bf959f22450bb3740f5b2fff4a59faa64bd6b8eaf4c57b5b8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 83.5 MB | 2026-08-25 | `sha256:aff6cec76c03` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:aff6cec76c03b70afc2c062ebe7fba1d8a0c4ec6a2fa1c8cd1982db6a3037423` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 139.0 MB | 2026-08-25 | `sha256:0b729c82c0dd` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:0b729c82c0ddc0769248e287d7414f0cc4e42ae4aa5b786aa99883c247e42bdb` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 70 | 842.0 MB | 2026-09-09 | `sha256:9f6b5c480db6` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:9f6b5c480db643ca536db638fc0346b5a8424fd738c916db5db6443e68ad130e` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 6 | 84 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.7 MB | 2026-08-11 | `sha256:4435f90009c1` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:4435f90009c17fb750e5518a3f43a24a629ac4c4f8c222b50f6adfe5e0d0bf2d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.2 MB | 2026-08-11 | `sha256:387a603a274e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:387a603a274e74568fd7a0e6d48ef68e631990e3b5149801515fe749a74b5b29` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.0 MB | 2026-08-25 | `sha256:daa1142fc6b4` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:daa1142fc6b44e27c8112ec6b4c2d579ddb9bc6b3747504e666010a45a51faa4` |
