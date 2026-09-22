# Daily Recommended Images by Language

_Generated: 2026-09-22T02:17:47Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.401 | - | 0 | 8 | 33 | 962.0 MB | 2026-09-12 | `sha256:01bb0a1008dd` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:01bb0a1008dd2b753ff059c9c18ec7e728b3d98d0024fa6e8c84905cdb86ae4e` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.318 | - | 0 | 8 | 43 | 900.0 MB | 2026-09-12 | `sha256:20455bfb45bc` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:20455bfb45bca8a6a6ae7f6df54af1c67f96a1a685e810fca2f09c8e7967c52c` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.31 | - | 4 | 52 | 235 | 193.0 MB | 2026-09-19 | `sha256:37466ea190f6` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:37466ea190f696105c1c3ae67c15e32d4e199face9a0b2ad5b9a37c464db8f30` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.31 | - | 4 | 52 | 235 | 218.0 MB | 2026-09-19 | `sha256:2f202e1169ec` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:2f202e1169ec507bdc07007cf68c14d0ff3a098110b17c460a60185e1f36a9d1` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.425 | - | 13 | 100 | 443 | 867.0 MB | 2026-09-19 | `sha256:78235e09001f` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:78235e09001f52b6592c458ac010775ebac6725422e80cd0c1650590f67b2743` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.31 | - | 0 | 0 | 10 | 193.0 MB | 2026-09-21 | `sha256:4b8f52c09215` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4b8f52c092153540f8d3ae94181f7749f4ced9d076bdc8e7c8742d4b517077f4` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.20 | - | 0 | 0 | 10 | 198.0 MB | 2026-09-21 | `sha256:b016fbf79133` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:b016fbf79133cb79f57a06f7b5df756eca697fc49bd74ab539f10d994d918edf` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.12 | - | 0 | 0 | 10 | 203.0 MB | 2026-09-21 | `sha256:ff17a18b639a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:ff17a18b639a0327e52c7c296fa2e1abe6e03eb61d8121a8ef67cc6aa430a27e` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.31 | - | 0 | 0 | 10 | 217.0 MB | 2026-09-21 | `sha256:84892b9bf258` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:84892b9bf258042b0b9175509292b9158e8c90c721f1e6a1fa8256f21a652764` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.20 | - | 0 | 0 | 10 | 223.0 MB | 2026-09-21 | `sha256:196831e5c6a2` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:196831e5c6a26dba1c0db2f15856332938c2cc6d9192715e45d935ef679e6af9` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.12 | - | 0 | 0 | 10 | 230.0 MB | 2026-09-21 | `sha256:2d584d8147fa` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:2d584d8147faddb0d678c5748d47953e5b8e18621ed4fb7049a91381d9d7746f` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.401 | - | 0 | 0 | 52 | 917.0 MB | 2026-09-21 | `sha256:35d40304542c` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:35d40304542c8689331f8cab17c65926cdf48fe711e289321d71924b230a7d29` |
| 8 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.318 | - | 0 | 0 | 62 | 855.0 MB | 2026-09-21 | `sha256:0f97a4002de8` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:0f97a4002de8050867e7e55e03a0487aea4f4a933158fe1d2cc35f570fc81a0d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.425 | - | 0 | 11 | 73 | 854.0 MB | 2026-09-21 | `sha256:2e171ed9da38` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:2e171ed9da38a01abb2005b8b2ce9df0c40e583a815e977a2c84d47778005dde` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.8 | - | 0 | 0 | 0 | 845.0 MB | 2026-09-21 | `sha256:429ea96b4261` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:429ea96b4261cdd5f49e9c1b82d4cf702fbfd09a62ade8ca6419521cd5ce3a94` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 0 | 20 | 158 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 0 | 0 | 323.0 MB | 2026-09-21 | `sha256:0a20478c03a3` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:0a20478c03a3a278be275a7d58b5647a448a658bf8f3231c0d61d823f187c812` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.20.1 | - | 0 | 0 | 0 | 327.0 MB | 2026-09-21 | `sha256:4fa6898e66f6` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:4fa6898e66f6dfe116b8efaac7377b672bd0a3f15be2b078ca504094e672c4a4` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.12.1 | - | 0 | 0 | 0 | 354.0 MB | 2026-09-21 | `sha256:338c150e1217` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:338c150e121763608677e70062ef3a5e67f62ee7325d9199ffd1fe786acdf32d` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.4.1 | - | 0 | 0 | 0 | 399.0 MB | 2026-09-21 | `sha256:3eadbf5975d4` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:3eadbf5975d47d6131d19f59708272daf3a61ec764fb1da55437a2e0a63aab6c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.12.1 | - | 0 | 0 | 0 | 481.0 MB | 2026-09-21 | `sha256:0aa59b319917` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:0aa59b3199178b65ac08169a3ab106516a0ff3a44ad8bb1569ec660ca6844906` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.4.1 | - | 0 | 0 | 0 | 526.0 MB | 2026-09-21 | `sha256:f2ddde481b6c` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:f2ddde481b6c8a7ce2f9396c82897c9a2ef451f1d35cdc2d212819b1393240c6` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.20.1 | - | 0 | 0 | 94 | 454.0 MB | 2026-09-21 | `sha256:f0d62a58aba4` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:f0d62a58aba4aa7f4e8d26ed131e674ad25f5eeda1adc92e73a5bfc37192a40b` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.12.1 | - | 0 | 0 | 94 | 482.0 MB | 2026-09-21 | `sha256:69e7c7cc0b53` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:69e7c7cc0b5365e40718d70759f77b7c4b16e86ddaacfdf499d1c4807ba592d5` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.4.1 | - | 0 | 0 | 94 | 526.0 MB | 2026-09-21 | `sha256:0f29a83f438f` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:0f29a83f438f5759c2b3c3c77b7de9032eaea7d0a83c2f4c4011c3d04ca3c51d` |

## Python

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.14 | :3-nonroot | 0 | 8 | 33 | 83.6 MB | 2026-09-11 | `sha256:6f7e96128ee1` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:6f7e96128ee1c66f167d8512178c0f6b43eca6dcdd9cd544a591f27ee5d75a4b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.14 | :3 | 0 | 8 | 33 | 83.6 MB | 2026-09-11 | `sha256:aa58cf2e8d81` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:aa58cf2e8d8190536744d947624014a4d2de744d1bb5f99c8aed97e710172fac` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.14 | :3 | 0 | 8 | 33 | 140.0 MB | 2026-09-11 | `sha256:6a4fcacee291` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:6a4fcacee291dffbb7a95881a8549cda0ff0999628975eebeec5b7ce6acb10fa` |

## Node

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.21` | 24.21.0 | :24 | 0 | 6 | 30 | 199.0 MB | 2026-09-16 | `sha256:9d24caf762ea` | `mcr.microsoft.com/azurelinux/base/nodejs:24.21@sha256:9d24caf762eaea25ab6742648553d19ee2b07897d545bd5a9884e6b313546df5` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24.20` | 24.20.0 | - | 0 | 6 | 33 | 199.0 MB | 2026-09-11 | `sha256:65efb6be4323` | `mcr.microsoft.com/azurelinux/base/nodejs:24.20@sha256:65efb6be4323d2da4f73531a37d84fff0e451cddc82e4dabc4b55d78ac85261b` |
| 3 | `mcr.microsoft.com/azurelinux/base/nodejs:24.18` | 24.18.1 | - | 0 | 8 | 52 | 197.0 MB | 2026-08-25 | `sha256:adfee798b577` | `mcr.microsoft.com/azurelinux/base/nodejs:24.18@sha256:adfee798b577f7f2d037dbb0c96d13fb78082aa1cff2598f1cd0417bf1b9e7ad` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot` | 24.21.0 | :24-nonroot | 0 | 10 | 39 | 158.0 MB | 2026-09-16 | `sha256:4321c8923aad` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21-nonroot@sha256:4321c8923aad0e89345e9dade530f9f3f951e4d962333c1d51a0febb951744af` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21` | 24.21.0 | :24 | 0 | 10 | 39 | 158.0 MB | 2026-09-16 | `sha256:550a59eaa4cb` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.21@sha256:550a59eaa4cbdbf06933741197173fcfddd73c5492b6b18a4ff8e3d1682425a5` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot` | 24.20.0 | - | 0 | 10 | 42 | 158.0 MB | 2026-09-11 | `sha256:6538b9fc8550` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20-nonroot@sha256:6538b9fc85501e2de5f68037b15eb7d08a7fe6220c1667861c79874238a5f9db` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20` | 24.20.0 | - | 0 | 10 | 42 | 158.0 MB | 2026-09-11 | `sha256:ca19b8b60f6a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.20@sha256:ca19b8b60f6af9efeb5d884a824aa2925b45723919884a45c164615d6330ac6a` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:24.17` | 24.17.0 | - | 0 | 21 | 91 | 196.0 MB | 2026-07-22 | `sha256:3d90ac240f72` | `mcr.microsoft.com/azurelinux/base/nodejs:24.17@sha256:3d90ac240f72fd1304281072a55b3e8d95eb8cca9ac88c375ec03bf3933f395b` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot` | 24.18.1 | - | 1 | 16 | 71 | 157.0 MB | 2026-08-25 | `sha256:ef2fa2bfcdcd` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18-nonroot@sha256:ef2fa2bfcdcd1d77255c424a928b5ff7fa324685cdea351105621b48c692dc15` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18` | 24.18.1 | - | 1 | 16 | 71 | 157.0 MB | 2026-08-25 | `sha256:b74f98dd614e` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.18@sha256:b74f98dd614ea4ced30ebea9ce267ae80083ed99da090806ba9285294e4e9721` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-09-11 | `sha256:f9aa2435862c` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:f9aa2435862cbb05a8d7fbeb9c2d024bf7548f3838f69a96cb214c1bade88ffa` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-09-11 | `sha256:4377af4aa7a8` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:4377af4aa7a810b7d59f691eae5066895a71aa3eee4cfb4eba527bbebff16479` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 6 | 30 | 76.4 MB | 2026-09-11 | `sha256:34a22db497ff` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:34a22db497ff34a0f35ca5fc54bd38711d04238a2c1b2f65d35dc9d45dd82584` |
