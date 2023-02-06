# Readme

## Run migrations

Up

```bash
export $(cat .env | xargs) && ./devtools/migrate -database ${POSTGRES_URL} -path db/migrations up

```

Down

```bash
export $(cat .env | xargs) && ./devtools/migrate -database ${POSTGRES_URL} -path db/migrations down

```
