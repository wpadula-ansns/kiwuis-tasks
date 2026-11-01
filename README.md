# kiw-lab-tasks

Laboratory tasks for the containerization and virtualization course.
The `main` branch holds only shared information. Each laboratory lives on its own branch, so switch to it before you start.

## Branches

| Branch | Laboratory |
| --- | --- |
| `lab-01` | L01: first containerised application (Go web server) |

## Requirements

| Tool | Version | Used for |
| --- | --- | --- |
| `git` | 2.23+ (`git switch`) | cloning and switching branches |
| Docker with Compose v2 | Docker Desktop with WSL 2 integration, or Docker Engine 24+ | building and running containers |
| `mise` (optional) | latest | installs the Go version pinned in `mise.toml` (Go 1.25.11) |
| `make` (optional) | GNU Make 4+ | shortcuts from `Makefile` |

Go is needed only to run the app outside a container: `mise install`, then `make run`.
The container build uses the `golang:1.25-alpine` image, so Docker alone is enough.

## Getting started

```bash
git clone https://github.com/storm-legacy/kiwuis-labs.git
cd kiwuis-labs
git switch lab-01
cp .env.example .env   # set NUMER_ALBUMU
make up                # or: docker compose up -d --build
```

Open http://localhost:8080.

Switching with uncommitted changes:

```bash
git stash push -u   # save changes, including new files
git switch lab-02
git switch lab-01
git stash pop       # restore saved changes
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT, see [LICENSE](LICENSE).
