# MySQLVars

`mysqlvars` is a simple CLI app to dump the contentes of `SHOW VARIABLES` for MySQL.

It writes to `STDOUT` which can be piped to a file.

## Installation

`go install github.com/grokfiy/mysqlvars`

## Usage

`mysqlvars -h <host> -P <port> -u <username> -p <password>`