---
title: Build YaraHunter
---

# Build YaraHunter

YaraHunter is a self-contained docker-based tool. Clone the [YaraHunter repository](https://github.com/Sam12121/YaraHunter), then build:

```bash
docker build --rm=true --tag=toaeio/toae_malware_scanner_ce:2.0.0 -f Dockerfile .
```

Alternatively, you can pull the official toae image at `toaeio/toae_malware_scanner_ce:2.0.0`.

```bash
docker pull toaeio/toae_malware_scanner_ce:2.0.0
```
