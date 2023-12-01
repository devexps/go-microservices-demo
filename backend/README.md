# Go Backend

## API documentation

### Swagger UI

- [Admin Swagger UI](http://localhost:8080/docs/)

### openapi.yaml

- [Admin openapi.yaml](http://localhost:8080/docs/openapi.yaml)

## Buf.Build uses

Use [buf.build](https://buf.build/) for engineering construction of Protobuf API.

Execute the command in the `backend` root directory:

### Update buf.lock

```bash
buf mod update
```

### Generate GO code

```bash
buf generate
```

### Generate OpenAPI v3 documentation

```bash
buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml
```

## Make build

Please execute under `app/{service name}/service`:

### Initialize development environment

```bash
make init
```

### Generate go code for API

```bash
make api
```

### Generate OpenAPI v3 documentation for the API

```bash
make openapi
```

### Generate ent code

```bash
make ent
```

### Generate wire code

```bash
make wire
```

### Build program

```bash
make build
```

### Run program

```bash
make run
```

### Build Docker image

```bash
make docker
```
