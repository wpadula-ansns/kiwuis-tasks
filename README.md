# kiw-lab-tasks

Laboratory tasks for the containerization and virtualization course.
The `main` branch holds only shared information. Each laboratory lives on its own branch, so switch to it before you start.

## Branches

| Branch | Laboratory |
| --- | --- |
| `lab-01` | L01: first containerised application (Go web server) |

## Getting started

```bash
git clone https://github.com/storm-legacy/kiwuis-labs.git
cd kiwuis-labs
git switch lab-01
```

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

Code is licensed under MIT (see LICENSE-CODE), content under CC BY 4.0 (see LICENSE-CONTENT).
