# Ninja Boilerplate

## Database

```bash
docker run -d -p 5438:5432 \
--name boilerplate-db \
-e POSTGRES_USER=boilerplate \
-e POSTGRES_PASSWORD=dev \
-e POSTGRES_DB=boilerplate \
postgres:11-alpine
```

```
docker exec -it boilerplate-db psql -U boilerplate
```

```sql
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## Load Balancer

```
cd configs
../bin/caddy -conf Caddyfile
```

## Web

```
cd web
npm install
npm start
```

## Server

```
cd server
./db-prepare.sh
go generate ./...
../bin/realize start
```
