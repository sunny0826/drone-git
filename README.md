# drone-git

[![Build Status](https://travis-ci.org/sunny0826/drone-git.svg?branch=master)](https://travis-ci.org/sunny0826/drone-git)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunny0826/drone-git)](https://goreportcard.com/report/github.com/sunny0826/drone-git)
![GitHub](https://img.shields.io/github/license/sunny0826/drone-git.svg)
![GitHub release](https://img.shields.io/github/release/sunny0826/drone-git)

drone plugin of git



```yaml
- name: 拉取配置
  image: guoxudongdocker/drone-git
  settings:
    token:
      from_secret: git_token
    config_path: https://github.com/sunny0826/configure.git
    out: configure
    check_enable: true
```