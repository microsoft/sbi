# Daily Recommended Images by Language

_Generated: 2026-10-25T02:17:30Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.21 | - | 0 | 0 | 0 | 108.0 MB | 2026-10-22 | `sha256:308fc618c272` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:308fc618c272877f7e448beaadcd34ed8b054680abda231144200811e0961ae2` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.13 | - | 0 | 0 | 0 | 113.0 MB | 2026-10-22 | `sha256:c334e6863649` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:c334e686364939560906b4a300b829c806397195b77c2320cb6cb081d265b8c6` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.32 | - | 0 | 0 | 0 | 126.0 MB | 2026-10-22 | `sha256:b3dd2a8b1098` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:b3dd2a8b1098226dd6c92c04b4a59f52f413bf73b4e4daa198f43c8b5ae8e13a` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.21 | - | 0 | 0 | 0 | 133.0 MB | 2026-10-22 | `sha256:57ba034fd207` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:57ba034fd207357fb9326e834bef0d7a9cbd8ead1d7917ec9524efee0718a7b1` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.13 | - | 0 | 0 | 0 | 140.0 MB | 2026-10-22 | `sha256:fa26a8a53d81` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:fa26a8a53d81c3a2aab53fe5d92d88e2569175d5718f56224b0543a008be1e69` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.319 | - | 0 | 0 | 10 | 900.0 MB | 2026-10-22 | `sha256:a57657f48b51` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:a57657f48b513031b8474e63c616c91076014685b6a245fb82bea926db58703e` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.402 | - | 0 | 4 | 5 | 963.0 MB | 2026-10-22 | `sha256:09b8d78b079f` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:09b8d78b079ff2237dbb56aa3ee70db2747ce23da2a6b5b095a7e16b8bf20d29` |

### Debian

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.32 | - | 1 | 58 | 243 | 194.0 MB | 2026-10-21 | `sha256:1b2f288e3a87` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:1b2f288e3a87cd251322f304e176327b6c00ec3bc1847f38b69a58cb7251e994` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.32 | - | 1 | 58 | 243 | 218.0 MB | 2026-10-21 | `sha256:74a1d7c07211` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:74a1d7c07211a7e8a2e35ff3c70ed8e4e1d6451a1ff7cf7bdf526f9eafa23023` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.426 | - | 1 | 98 | 459 | 868.0 MB | 2026-10-21 | `sha256:738f72aff955` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:738f72aff955d5c5690e06b6fc8ecd1b90ead564099a257ac528ff720b8da85f` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.32 | - | 0 | 0 | 27 | 193.0 MB | 2026-10-21 | `sha256:d1a464e7b68a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:d1a464e7b68a37abf1739f6d97d3e164a0ed908e8b3bb04a9861b9d940acb876` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.21 | - | 0 | 0 | 27 | 198.0 MB | 2026-10-21 | `sha256:cf7601452fc4` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:cf7601452fc41575ba15538c0f12e5df31d9c0098820a73c222f23bfa736a004` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.13 | - | 0 | 0 | 27 | 203.0 MB | 2026-10-21 | `sha256:0cf09aaccf6a` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:0cf09aaccf6a9f68900aed4c73fb28fe25bb59ec85de4dea08ba02240a8635f3` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.32 | - | 0 | 0 | 27 | 217.0 MB | 2026-10-21 | `sha256:9449b45658df` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:9449b45658dfb30286dad89aa55ec3640eb0cea2edb7f6842c9d734920ed9b2b` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.21 | - | 0 | 0 | 27 | 223.0 MB | 2026-10-21 | `sha256:f154a9699db4` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:f154a9699db4bc90db46159a3214c8e59dc981524bdc6bc2a9adcc83614cfa82` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.13 | - | 0 | 0 | 27 | 230.0 MB | 2026-10-21 | `sha256:8af87eba5e7c` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8af87eba5e7cdafb1514d7388be287365da97c132509e84dfe252ffd552398ca` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.319 | - | 0 | 0 | 43 | 855.0 MB | 2026-10-21 | `sha256:0e0342bb18ef` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:0e0342bb18ef8c8eb7c5a189463d4840cf595706aa9d0494349a6a64965c5428` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.402 | - | 0 | 4 | 38 | 918.0 MB | 2026-10-21 | `sha256:747cbade5bfd` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:747cbade5bfd918a175ac32f8f1d931edec2c7d0d5397a3de72eaef1ac1f66d2` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.426 | - | 0 | 16 | 60 | 854.0 MB | 2026-10-21 | `sha256:ada801965b46` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:ada801965b46b26af97d6d366300ec06ca793c581259664c226253d9da55d254` |

## Go

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.10 | - | 0 | 0 | 0 | 843.0 MB | 2026-10-21 | `sha256:d104fef62900` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:d104fef62900c2ca101f52568f43e0e6b6499b60be324fdd4d2ad5bae4bc3203` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.14 | - | 3 | 69 | 339 | 810.0 MB | 2026-08-20 | `sha256:8d9cba1312f5` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8d9cba1312f5dc497d59b5e3d67b065025f548b10e9dbb439bf9fb170048612b` |

## Java

### Azure Linux

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.13.0.2 | - | 0 | 0 | 0 | 480.0 MB | 2026-10-23 | `sha256:23f0878343f6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:23f0878343f60c9a5be5994204bd9790b6ee74e5bf728a7680341a685c634e06` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.5.0.2 | - | 0 | 0 | 0 | 524.0 MB | 2026-10-23 | `sha256:6ed12091e98f` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:6ed12091e98f50d6924b7f7690353a3b85ca6a94e617286420223086391a0a76` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.32.1 | - | 0 | 1 | 1 | 324.0 MB | 2026-10-24 | `sha256:470f6727be5f` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:470f6727be5f25a259cdfb2f0b9b7a4defba450d47fbd555bd71a2f91b907693` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.21.0.2 | - | 0 | 1 | 1 | 327.0 MB | 2026-10-23 | `sha256:1bff96c8e962` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:1bff96c8e962a88fb5926d775f5603a9d83f234c8fecf545be768c4385d31c68` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.13.0.2 | - | 0 | 1 | 1 | 355.0 MB | 2026-10-23 | `sha256:e5b59ca756ec` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:e5b59ca756ec27899df8909a98b9c73c9271b7107b41334855acd5647359fba0` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.5.0.2 | - | 0 | 1 | 1 | 400.0 MB | 2026-10-23 | `sha256:9bf874a31638` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:9bf874a3163858de23ebc7ed13d73fbc8d4dba7f166deb8d38797e69925b461b` |

### Ubuntu

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.21.0.2 | - | 0 | 1 | 90 | 432.0 MB | 2026-10-23 | `sha256:37cda9c1b3e7` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:37cda9c1b3e78f852909f3ff6c5c98ab21a0596f766b38af73baf8f6536ffd0d` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.13.0.2 | - | 0 | 1 | 90 | 459.0 MB | 2026-10-23 | `sha256:fb47807c17b9` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:fb47807c17b994a0d28e3d4a046b471c24e0e4d15fa5d3640adffa9404e0e48c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.5.0.2 | - | 0 | 1 | 90 | 503.0 MB | 2026-10-23 | `sha256:ae229b654046` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:ae229b6540460ecdaa751f9ef9ef6385095ace03d40e88321106f29e1238ca15` |

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
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.15 | :3-nonroot | 0 | 0 | 0 | 84.0 MB | 2026-10-22 | `sha256:f5b31eb11254` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:f5b31eb11254a444ba92837d788c349344cde899110f629100f8ca1efd49ade3` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 84.0 MB | 2026-10-22 | `sha256:4b61cf6b0730` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:4b61cf6b07307909fd71235dc5e16810582660c60b8269b32f7f89d6209f01c9` |
| 3 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.15 | :3 | 0 | 0 | 0 | 140.0 MB | 2026-10-22 | `sha256:9343005b72a7` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:9343005b72a7cc34847dcb8d09bb8a0c145311eff88c151c60d9c4904edaf6e9` |

## Base / No Runtime

| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |
|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | - | 0 | 0 | 0 | 3.8 MB | 2026-10-22 | `sha256:c91e907c537a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:c91e907c537ad861cb90bc18aece90479b0754b090afd4eb8722dd15477dc9e8` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | - | 0 | 0 | 0 | 34.3 MB | 2026-10-22 | `sha256:6b85a86416a2` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:6b85a86416a29cdcfd404e9fbd9c5ef49cc0612882cecc702535dcac39c9446c` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | - | 0 | 0 | 0 | 76.9 MB | 2026-10-22 | `sha256:bd88359666e3` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:bd88359666e3f20296e7fe639a480dfd3c3c8eba299b21f26e4560353e21a656` |
