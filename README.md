# Cyclone

<p align="center"><img src="docs/images/logo.jpeg" width="200"></p>

[![Build Status](https://travis-ci.org/caicloud/cyclone.svg?branch=master)](https://travis-ci.org/caicloud/cyclone)
[![Go Report Card](https://goreportcard.com/badge/github.com/caicloud/cyclone?style=flat-square)](https://goreportcard.com/report/github.com/caicloud/cyclone)
[![Coverage Status](https://coveralls.io/repos/github/caicloud/cyclone/badge.svg?branch=master)](https://coveralls.io/github/caicloud/cyclone?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/caicloud/cyclone)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)

Cyclone is a powerful workflow engine and end-to-end pipeline solution implemented with native Kubernetes resources,
with no extra dependencies. It can run anywhere Kubernetes is deployed: public cloud, on-prem or hybrid cloud.

Cyclone is architectured with a low-level workflow engine that is application agnostic, offering capabilities like
workflow DAG scheduling, resource lifecycle management and most importantly, a pluggable and extensible framework for
extending the core APIs. Above which, Cyclone provides built-in support for high-level functionalities, with CI/CD
pipelines and AI DevOps being two notable examples, and it's possible expand to more use cases as well.

With Cyclone, users end up with the flexibility of workflow orchestration and the usability of complete solutions.

## Features

- DAG graph scheduling: Cyclone supports DAG workflow execution
- Parameterization: stage (unit of execution) can be parameterized to maximize configuration reuse
- External dependencies: external systems like SCM, docker registry, S3 can be easily integrated with Cyclone
- Triggers: Cyclone supports cron and webhook trigger today, with upcoming support for other types of triggers
- Controllability: workflow execution can be paused, resumed, retried or cancelled
- Multi-cluster: workflow can be executed in different clusters from where Cyclone is running
- Multi-tenant: resources manifests and workflow executions are grouped and isolated per tenant
- Garbadge Collection: automatic resource cleanup after workflow execution
- Logging: logs are persisted and indpendent from workflow lifecycle, enabling offline inspection
- Built-in Pipeline: curated DAG templates and stage runtimes for running DevOps pipelines

## Quick Start

Build and push images. `<registry>/<project>` specifies registry and project where to push your images, for example, `test.caicloudprivatetest.com/release`.

```bash
$ make push REGISTRIES=<registry>/<project>
```

Deploy to Kubernetes cluster, please make sure `kubectl` is installed and appropriately configured.

```bash
$ make deploy REGISTRIES=<registry>/<project>
```

## Community

- **Slack**: Join [Cyclone Community](https://caicloud-cyclone.slack.com/join/signup) for disscussions and posting questions.

## Roadmap

[Cyclone Roadmap](./docs/ROADMAP.md)

## Contributing

If you are interested in contributing to Cyclone, please checkout [CONTRIBUTING.md](./CONTRIBUTING.md).
We welcome any code or non-code contribution!

## Licensing

Cyclone is licensed under the Apache License, Version 2.0. See [LICENSE](./LICENSE) for the full license text.
