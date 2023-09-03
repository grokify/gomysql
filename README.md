# MySQLVars

`mysqlvars` is a simple CLI app to dump the contents of `SHOW VARIABLES` for MySQL.

It writes to `STDOUT` which can be piped to a file.

## CLI Installation and Usage

```
% go install github.com/grokify/mysqlvars@latest
% mysqlvars -h <host> -P <port> -u <username> -p <password>
```

## Script Installation and Usage

```
% git clone https://github.com/grokify/mysqlvars
% cd mysqlvars
% go run main.go -h <host> -P <port> -u <username> -p <password>
```