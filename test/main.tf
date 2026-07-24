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
