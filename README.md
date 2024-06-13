## Temporal, Go-Gin Server, Traefik - Poc

This is a Temporal, Go-Gin Server, Traefik - Poc

## How to use

- Install temporal server from [temporal.io](https://github.com/temporalio/temporal)
- Run the below command to start the server
  ```bash
  temporal server start-dev
  ```

- Run the below command to run the worker
  ```bash
    go run cmd/worker/main.go
  ```
- Run the below command to run the server
  ```bash
    go run cmd/server/main.go
  ```

## How to check the workflow of the server
- check the workflow on `http://localhost:8233/namespaces/default/workflows`

## How to run the traefik
- Run the below command to start the traefik
  ```bash
  docker-compose up -d
  ```

## How to stop the traefik
- Run the below command to stop the traefik
  ```bash
  docker-compose down
  ```

## How to check the traefik
- Run the below command to check the traefik
  ```bash
  docker-compose ps
  ```