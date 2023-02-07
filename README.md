# Readme

## Run migrations

### Pre-install

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

```

### Create

```bash
/devtools/migrate create -ext sql -dir db/migrations create_user_table
```

### Up

```bash
export $(cat .env | xargs) && /devtools/migrate -database ${POSTGRES_URL} -path db/migrations up

```

### Down

```bash
export $(cat .env | xargs) && /devtools/migrate -database ${POSTGRES_URL} -path db/migrations down

```
