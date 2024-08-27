# Orders Challenge

## Rodando o Projeto

1. Suba o banco de dados:
   ```sh
   docker compose up -d
   ```
2. Compile e execulte cada serviço:

- REST (porta 8080):

```sh
go run ./cmd/rest/main.go
```

- gRPC (porta 50051):

```sh
go run ./cmd/grpc/main.go
```

- GraphQL (porta 4000):

```sh
go run ./cmd/graphql/main.go
```

3. Ultilize o arquivo `api.http` para testar os endpoints.

## Requisitos

- Docker
- Go
