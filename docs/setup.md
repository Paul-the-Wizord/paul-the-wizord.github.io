# Setup

One-time bootstrap to get the provider registry running on GitHub Pages.

This repo (`paul-the-wizord.github.io`) is a GitHub **user-pages** repo, so
its contents are served at the host root `https://paul-the-wizord.github.io/`.
That root is exactly where the Terraform registry protocol's service discovery
document must live — hence the user-pages slot is required.

## 1. Free up the `paul-the-wizord.github.io` host root

GitHub Pages serves a **user site** (at host root `<user>.github.io`) only
from a repo named exactly `<user>.github.io` — one per account. Any prior
content repo occupying that slot must be renamed first:

```bash
gh repo rename paul-the-wizord-old -R Paul-the-Wizord/paul-the-wizord.github.io
```

Then create a fresh empty repo to host both the provider source and the
registry:

```bash
gh repo create Paul-the-Wizord/paul-the-wizord.github.io --public
```

> Leave it empty (no README, no .gitignore). The first push fills it.

## 2. Generate a GPG keypair

The registry protocol requires provider checksums to be GPG-signed.

```bash
# Generate (RSA 4096, no expiry for a PoC, name it "Terraform Provider Hello")
gpg --full-generate-key

# List keys and grab the KEY_ID (the long hex after 'rsa4096/')
gpg --list-secret-keys --keyid-format=long

# Export the PUBLIC key and commit it to this repo
gpg --armor --export <KEY_ID> > gpg/public-key.asc
git add gpg/public-key.asc
git commit -m "Add GPG public key for provider signing"
git push

# Export the PRIVATE key (for the repo secret — never commit this file)
gpg --armor --export-secret-keys <KEY_ID> > /tmp/private-key.asc
```

## 3. Add repository secrets

In this repo: **Settings → Secrets and variables → Actions → New repository secret**

| Secret name       | Value                              |
|-------------------|------------------------------------|
| `GPG_PRIVATE_KEY` | contents of `/tmp/private-key.asc`  |
| `GPG_PASSPHRASE`   | the passphrase you set on the GPG key |

No `PAGES_PAT` is needed — the publish workflow commits to the same repo
using the default `GITHUB_TOKEN`.

## 4. Verify

Push a commit to `main`. Watch the three workflows run in sequence:

1. **Build and Release** — creates a `v0.0.1` tag + release.
2. **Publish Registry** — commits registry JSON to this repo (`[skip ci]`).
3. **Test** — runs `terraform apply`/change/`destroy` against the live registry.

After workflow 2 completes, confirm the discovery document is served:

```bash
curl https://paul-the-wizord.github.io/.well-known/terraform.json
# {"providers.v1": "/v1/providers/"}
```
