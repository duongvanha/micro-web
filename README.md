# app

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
# run kubernetes

```bash
#   set pg password
    kubectl create secret generic pgpassword --from-literal PGPASSWORD=123456
    kubectl apply -f kubernetes 
```

#make data 
```bash
#    pg_restore -U xmvmpwalsagthi -d test -1 /usr/src/data/d396662e-8721-44de-b579-0ea90f650032
    dc run postgres --rm -c "pg_restore -U xmvmpwalsagthi -d test -1 /usr/src/data/d396662e-8721-44de-b579-0ea90f650032"

```