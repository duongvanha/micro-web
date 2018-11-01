# app cinema microservices

## docker compose for development env
```bash
 copy file .env.example -> .env
 export NODE_ENV=development
 ./install.sh
 
 docker-compose up -d
 
```
## build proto
```bash
    protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto
```
## run kubernetes

```bash
#   set pg password
    kubectl create secret generic pgpassword --from-literal PGPASSWORD=123456
    kubectl apply -f kubernetes 
```

## make data 
```bash
#   psql --username xmvmpwalsagthi --dbname test -f /usr/src/data/db
    dc run postgres --rm -c "psql --username xmvmpwalsagthi --dbname test -f /usr/src/data/db"

```
