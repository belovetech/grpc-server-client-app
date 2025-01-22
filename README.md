# GRPC Implementation

This grpc mini project demonstrates the use of GRPC as communication protocol between server(go) and client(python).

## Set up server

    - cd into server directory `cd server`
    - run go mod tidy to install go dependencies
    - start server

## Set up client

    - cd into client directory `cd client`
    - setup python environment by running `make env`
    - activate the environment by running `make activate`
    - run `pip install -r requirements.txt` to install python dependencies
    - start client

## Start server

```bash
    make run

    or

    go run server.go
```

## Invoke Client

```bash
    make run

    or

    python3 client.py
```
