# GoMySQL

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

GoMySQL is a Go module that includes various helpers for MySQL, including `mysqlvars`, is a simple CLI app to dump the contents of `SHOW VARIABLES` for MySQL. It writes to `STDOUT` which can be piped to a file.

## CLI Installation and Usage

```
% go install github.com/grokify/gomysql/cmd/mysqlvars@latest
% mysqlvars -h <host> -P <port> -u <username> -p <password>
```

## Script Installation and Usage

```
% git clone https://github.com/grokify/gomysql
% cd gomysql/cmd/mysqlvars
% go run main.go -h <host> -P <port> -u <username> -p <password>
```

 [build-status-svg]: https://github.com/grokify/gomysql/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/gomysql/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gomysql
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/gomysql
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/gomysql
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/gomysql
 [loc-svg]: https://tokei.rs/b1/github/grokify/gomysql
 [repo-url]: https://github.com/grokify/gomysql
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/gomysql/blob/master/LICENSE
