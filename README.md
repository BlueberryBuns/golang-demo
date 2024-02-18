# Golang demo 18.02.2024

This is demo of go for internal side project. It' purpose is to show how to work with the language itself.

### Useful resources

- [Official Go website](https://go.dev/)
- [Go Standard library docs](https://pkg.go.dev/std)
- [Golang project structure - unofficial projects layout GitHub](https://github.com/golang-standards/project-layout)
- [GrafanaLabs: How I write HTTP services in Go after 13 years by Mat Ryer](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
- [Awesome Go - Repo of useful packages](https://github.com/avelino/awesome-go)
- [JetBrains tutorial for all of JB fans - not for broke a** devs](https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/)
- [Middleware Patterns in Go](https://drstearns.github.io/tutorials/gomiddleware/)
- [Making and using Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)

## Prerequisites

- Docker installed - [How to install docker](https://docs.docker.com/engine/install/) - We are intentionally skipping docker compose for now!
    - [Optional Prerequisite] [How to install WSL2 on Windows](https://learn.microsoft.com/en-us/windows/wsl/install)
- Golang image downloaded [version 1.22.0](https://hub.docker.com/_/golang)
- Alpine Linux [image](https://hub.docker.com/_/alpine) downloaded
- Golang installed on your computer [version 1.22.0](https://go.dev/dl/) - required as utility
- Editor installed VSCode + golang plugin (recommended) / Vim

## Presentation plan

1. Why Go - advantages / disadvantages
2. Comparison Go vs Python (the most basic structures and use cases + glimpse of composition-based OOP)
3. Docker example and running your code
4. Writing HTTP server (Why standard library and not a framework)
5. Testing in go
6. Deploying your package