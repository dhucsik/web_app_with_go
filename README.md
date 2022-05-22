In this repo I'm learning how to use go programming language in backend.

## Installation

delete go.mod and go.sum files

```bash
go mod init github.com/dhucsik/web_app_with_go 
```
you can write what you want but then you need to change imports in go files

```bash
go get github.com/alexedwards/scs/v2
go get github.com/go-chi/chi/v5
go get github.com/justinas/nosurf 
```

## Start application
Mac OS

```bash
go run cmd/web/*.go
```
Others
```bash
go run ./cmd/web
```

# Go modules i used: 

github.com/go-chi/chi/v5

github.com/alexedwards/scs/v2

github.com/justinas/nosurf
