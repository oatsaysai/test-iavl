# TEST-IAVL (golang)

Test set key/value to Tendermint IAVL

## Prerequisites

- Go version >= 1.9.2

  - [Install Go](https://golang.org/dl/) by following [installation instructions.](https://golang.org/doc/install)
  - Set GOPATH environment variable (https://github.com/golang/go/wiki/SettingGOPATH)

## Getting started

1.  Get dependency

    ```sh
    dep ensure
    ```

2.  Run a IAVL server

    ```sh
    SERVER_PORT=8080 \
    DB_TYPE=goleveldb \
    DB_DIR_PATH=DB \
    go run server/*
    ```

**Environment variable options**

- `SERVER_PORT`: IAVL server port [Default: `8080`]
- `DB_TYPE`: Database type (same options as Tendermint's `db_backend`) [Default: `goleveldb`]
- `DB_DIR_PATH`: Directory path for database files [Default: `DB`]

3. Run a client

    ```sh
    SERVER_ADDRESS=http://127.0.0.1:8080 \
    TXPERSEC=10 \
    DURATION=60 \
    go run client/main.go
    ```

**Environment variable options**

- `SERVER_ADDRESS`: Address of IAVL server [Default: `http://127.0.0.1:8080`]
- `TXPERSEC`: Transaction per second [Default: `10`]
- `DURATION`: Duration (second) [Default: `60`]