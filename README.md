# CloudCoreo CLI

[![Build Status](https://travis-ci.org/CloudCoreo/cli.svg?branch=master)](https://travis-ci.org/CloudCoreo/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/CloudCoreo/cli)](https://goreportcard.com/report/github.com/CloudCoreo/cli)

CLI is a tool for managing CloudCoreo resources. 

Use CLI to...

- Add/remove teams, cloud accounts, Git keys and API tokens
- Intelligently manage your composites manifest files
- Manage releases of CloudCoreo composites/plans
- Create reproducible plans of your CloudCoreo applications

## Install

**DISCLAIMER: These are PRE-RELEASE binaries -- use at your own peril for now**

### OSX

Download `coreo` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_darwin_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_darwin_amd64)

```sh
$ mkdir coreo && cd coreo
$ wget -q -O coreo https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_darwin_amd64
$ chmod +x coreo
$ export PATH=$PATH:${PWD}   # Add current dir where coreo has been downloaded to
$ coreo
```

### Linux

Download `coreo` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_linux_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_linux_amd64)

```sh
$ mkdir coreo && cd coreo
$ wget -q -O coreo https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_linux_amd64
$ chmod +x coreo
$ export PATH=$PATH:${PWD}   # Add current dir where coreo has been downloaded to
$ coreo
```

### Windows

Download `coreo.exe` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_windows_amd64.exe](https://github.com/CloudCoreo/cli/releases/download/v0.0.17/coreo_windows_amd64.exe)

```
C:\Users\Username\Downloads> rename coreo_windows_amd64.exe coreo.exe
C:\Users\Username\Downloads> coreo.exe
```

### Building from source

Build instructions are as follows (see [install golang](https://docs.minio.io/docs/how-to-install-golang) for setting up a working golang environment):

```sh
$ go get -d github.com/CloudCoreo/cli
$ cd $GOPATH/src/github.com/CloudCoreo/cli/cmd
$ go build -o $GOPATH/bin/coreo
$ coreo
```

## Docs

Get started with [Coreo commands](docs/coreo/coreo.md), setup for [Coreo bash completion](docs/bash-completion.md)

## Community, discussion, contribution, and support

GitHub's Issue tracker is to be used for managing bug reports, feature requests, and general items that need to be addressed or discussed.

From a non-developer standpoint, Bug Reports or Feature Requests should be opened being as descriptive as possible.

### Code of conduct

Participation in the CloudCoreo community is governed by the [Coreo Code of Conduct](code-of-conduct.md).
