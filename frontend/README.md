# Frontend

## Install and use

### Install dependent libraries

```bash
pnpm install
```

### Run

```bash
pnpm serve
```

### Build

```bash
pnpm build
```

### OpenAPI documentation generates DTS files

```bash
npx openapi-typescript http://localhost:8080/docs/openapi.yaml --output ./types/openapi.d.ts
```
