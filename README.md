# Terraform Provider Hello

A PoC Terraform provider published via the [Provider Registry Protocol][registry]
on GitHub Pages. One resource, `hello_world`, prints messages on create/update/delete.

[registry]: https://developer.hashicorp.com/terraform/internals/provider-registry-protocol

| Lifecycle | Output            |
| --------- | ----------------- |
| Create    | `hello, <name>!`  |
| Update    | `<name> changed!` |
| Delete    | `bye, <name>!`    |

## Using the provider

```hcl
terraform {
  required_providers {
    hello = {
      source  = "paul-the-wizord.github.io/example/hello"
      version = "~> 0.0"
    }
  }
}

provider "hello" {}

resource "hello_world" "demo" {
  name = "world"
}
```

```bash
terraform init      # discovers the provider via .well-known/terraform.json
terraform apply     # prints: hello, world!
terraform destroy   # prints: bye, world!
```

Messages appear as Terraform warnings (no `TF_LOG` needed).

## Architecture

This repo is a GitHub **user-pages** repo (`paul-the-wizord.github.io`), so its
contents are served at the host root — exactly where the registry protocol's
discovery document (`.well-known/terraform.json`) must live. It holds both the
provider source and the served registry metadata.

Provider binaries live as **GitHub Release assets**; the Pages site serves only
lightweight metadata JSON pointing at them.

```
paul-the-wizord.github.io
├── main.go / internal/provider/         # provider source
├── .github/workflows/                   # build → publish → test pipeline
├── test/                                # end-to-end terraform test
├── gpg/public-key.asc                   # signing public key
├── .nojekyll                            # disables Jekyll (serves dotfiles)
├── .well-known/terraform.json           # {"providers.v1": "/v1/providers/"}
└── v1/providers/example/hello/          # registry metadata (extensionless)
    ├── versions                         # list available versions
    └── <ver>/download/linux/amd64       # download package metadata
```

Registry files are **extensionless** because the protocol requests paths like
`/versions` and `/download/linux/amd64` — GitHub Pages serves these as static
files when `.nojekyll` is present.

## CI pipeline

| # | Workflow | Trigger | What it does |
|---|----------|---------|-------------|
| 1 | `1-build-and-release` | push to `main` | bump patch tag, build linux/amd64, GPG-sign, release |
| 2 | `2-publish-registry` | `workflow_run` (#1) | regenerate registry metadata, commit to this repo |
| 3 | `3-test` | `workflow_run` (#2) | pin latest version, `terraform init/apply/change/destroy` |

Chained via `workflow_run` because releases created with `GITHUB_TOKEN` don't
emit `release` events. Workflow 2's commit uses `[skip ci]`.

## Setup

See [`docs/setup.md`](docs/setup.md). Requires a GPG keypair (`gpg/public-key.asc`
committed; private key + passphrase as repo secrets `GPG_PRIVATE_KEY` and
`GPG_PASSPHRASE`).

## Local development

```hcl
# ~/.terraformrc
provider_installation {
  dev_overrides {
    "paul-the-wizord.github.io/example/hello" = "/path/to/repo/bin"
  }
  direct {}
}
```

```bash
make build VERSION=0.0.1
cd test && terraform apply
```
