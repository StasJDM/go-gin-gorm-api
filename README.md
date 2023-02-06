# Readme

## Run migrations

### Pre-install

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

```

### Up

```bash
export $(cat .env | xargs) && ./migrate -database ${POSTGRES_URL} -path db/migrations up

```

### Down

```bash
export $(cat .env | xargs) && ./migrate -database ${POSTGRES_URL} -path db/migrations down

```
