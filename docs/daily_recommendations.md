# Daily Recommended Images by Language

_Generated: 2026-09-18T02:19:46Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 108.0 MB | 2026-09-12 | `sha256:e5d3046590bf` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:e5d3046590bf050a78776f350fb3eff963742719e8a3bf50ff167a73981bbd91` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 112.0 MB | 2026-09-12 | `sha256:cb5995f1878d` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:cb5995f1878da8b87c7496ebae6e4a0c3f5d3262ca48977f330a6b334af3f783` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.31 | - | 0 | 0 | 0 | 126.0 MB | 2026-09-12 | `sha256:1e1aa734db68` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:1e1aa734db68ba3c21b22a5acada92adde68f28674ee687df9b6c29ae08f7a96` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.20 | - | 0 | 0 | 0 | 132.0 MB | 2026-09-12 | `sha256:dd3762bcb3b6` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:dd3762bcb3b6c70a1bb970f3a29ac171610af8508dce23147c6d1448804bfe48` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.12 | - | 0 | 0 | 0 | 139.0 MB | 2026-09-12 | `sha256:42e74b1e732f` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:42e74b1e732f4b17c6a292c189f2110d3501b6caa96bb59edda698eb059ecc9b` |
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 0 | 0 | 962.0 MB | 2026-09-12 | `sha256:01bb0a1008dd` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:01bb0a1008dd2b753ff059c9c18ec7e728b3d98d0024fa6e8c84905cdb86ae4e` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 0 | 10 | 900.0 MB | 2026-09-12 | `sha256:20455bfb45bc` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:20455bfb45bca8a6a6ae7f6df54af1c67f96a1a685e810fca2f09c8e7967c52c` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 55 | 242 | 193.0 MB | 2026-09-04 | `sha256:9cfa8aaf5c98` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:9cfa8aaf5c98a4cedffec74d450dd2d4510ba72ee9e663efe2658500b0321524` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 55 | 242 | 218.0 MB | 2026-09-04 | `sha256:9a464e9a7e8c` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:9a464e9a7e8c6144631020975f703c89034fe386417cb740620df69c2c6cfe24` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 105 | 452 | 867.0 MB | 2026-09-04 | `sha256:5ef85cc12cb2` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:5ef85cc12cb25be6ec319a7392d1e9efd53c3bc8abb971c53d8058a473f09053` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 0 | 12 | 193.0 MB | 2026-09-11 | `sha256:104cc434e0e6` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:104cc434e0e6385522d26e575705f9c6b08020838226b1e1856865bb00bb0f30` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 0 | 12 | 198.0 MB | 2026-09-11 | `sha256:292a2258079b` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:292a2258079b972164633918da438944e8561302884ced44862e4b2a936974ef` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 0 | 12 | 203.0 MB | 2026-09-11 | `sha256:8a153b5889d7` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:8a153b5889d796b6450295b383596b13308c24c230515f8a7770ce1b94e0c460` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 0 | 12 | 217.0 MB | 2026-09-11 | `sha256:ce687d884ca7` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:ce687d884ca748a70af86b901782afbd54214c4b930148c42d85e246c9881c1f` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 0 | 12 | 223.0 MB | 2026-09-11 | `sha256:cdbdb91fbd1d` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:cdbdb91fbd1d95e8561bd126858ebceef0a0f00dd3fe90d72d8b40f0b25602b4` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 0 | 12 | 230.0 MB | 2026-09-11 | `sha256:6a94333d3751` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:6a94333d37514e385650a3c81a55e5350b67253dbe136e9cf17e499c35606a8c` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 0 | 60 | 917.0 MB | 2026-09-11 | `sha256:2fa828c68761` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:2fa828c68761b1b8c23d7662dc134421b9d3b59fe1425fdbc80804e390cdb24d` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 0 | 70 | 855.0 MB | 2026-09-11 | `sha256:d07d5770b2b7` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:d07d5770b2b74717bfe882250e981a4e7a756e0240fc4aefcc0ce93c1e8a9c36` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 11 | 81 | 854.0 MB | 2026-09-11 | `sha256:0aa5a155ba07` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:0aa5a155ba075ad0effc3f60b3af6a696c2a498703519748e09e7eba8c5586f7` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 0 | 843.0 MB | 2026-09-16 | `sha256:13a693fd5c7f` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:13a693fd5c7f7501ff91f3b941911de66dfabe1b36583ab2f711a907c1400731` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 8 | 115 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-09-16 | `sha256:b87c51265858` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:b87c51265858740dfecea1aa5d431ef21bca48940007a66fc13a76c9dded30f1` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-09-16 | `sha256:6e5bc91d7512` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:6e5bc91d75127e8d48b3ea8731a2e7e6a7022cef2472bea1114461d843c8b732` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-09-16 | `sha256:fe5aa2db35c8` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:fe5aa2db35c8ec6a75702b084873bbfcbe25133d6190b50f4634046da08a882d` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-09-16 | `sha256:43ead3eeea8e` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:43ead3eeea8e75602ca811ff753ea1d14638959a6ef9eb8d58bd0afc1501586c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 478.0 MB | 2026-09-16 | `sha256:ce0f927cd0ed` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:ce0f927cd0ed790f97d1df82cbce8d887bb5b672a0633b84455fc183a5f9ca5e` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 523.0 MB | 2026-09-16 | `sha256:cc2d6b167e58` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:cc2d6b167e58443746dafb354d46d2dc01b4f2082529488c926d42912af3fe23` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 94 | 454.0 MB | 2026-09-16 | `sha256:3639b2d3e3cf` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:3639b2d3e3cf44ac06ea468dbccc8cfb278e269426b6ed9b9c39160bed7d0b77` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 94 | 482.0 MB | 2026-09-16 | `sha256:fa7221968d66` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:fa7221968d66318fc30725712070f2db7256d458fca1e79357ea325cacf6c26d` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 94 | 526.0 MB | 2026-09-16 | `sha256:0d750c32f0fa` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:0d750c32f0fa61e47f70f63dbfe48b9c6fc47d011f57655f4025534be14f3b3a` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 0 | 0 | 199.0 MB | 2026-09-16 | `sha256:9d24caf762ea` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:9d24caf762eaea25ab6742648553d19ee2b07897d545bd5a9884e6b313546df5` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 0 | 3 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 3 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 2 | 22 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 4 | 9 | 158.0 MB | 2026-09-16 | `sha256:4321c8923aad` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:4321c8923aad0e89345e9dade530f9f3f951e4d962333c1d51a0febb951744af` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 4 | 9 | 158.0 MB | 2026-09-16 | `sha256:550a59eaa4cb` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:550a59eaa4cbdbf06933741197173fcfddd73c5492b6b18a4ff8e3d1682425a5` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 4 | 12 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 4 | 12 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 15 | 61 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 10 | 41 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 10 | 41 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 0 | 0 | 83.6 MB | 2026-09-11 | `sha256:6f7e96128ee1` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:6f7e96128ee1c66f167d8512178c0f6b43eca6dcdd9cd544a591f27ee5d75a4b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 83.6 MB | 2026-09-11 | `sha256:aa58cf2e8d81` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:aa58cf2e8d8190536744d947624014a4d2de744d1bb5f99c8aed97e710172fac` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-09-11 | `sha256:6a4fcacee291` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:6a4fcacee291dffbb7a95881a8549cda0ff0999628975eebeec5b7ce6acb10fa` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-11 | `sha256:f9aa2435862c` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:f9aa2435862cbb05a8d7fbeb9c2d024bf7548f3838f69a96cb214c1bade88ffa` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-11 | `sha256:4377af4aa7a8` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:4377af4aa7a810b7d59f691eae5066895a71aa3eee4cfb4eba527bbebff16479` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.4 MB | 2026-09-11 | `sha256:34a22db497ff` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:34a22db497ff34a0f35ca5fc54bd38711d04238a2c1b2f65d35dc9d45dd82584` |
