# app

## docker compose for development env
```bash
 copy file .env.example -> .env
 export NODE_ENV=development
 
 docker-compose up -d
 
```
# run kubernetes

```bash
#   set pg password
    kubectl create secret generic pgpassword --from-literal PGPASSWORD=123456
    kubectl apply -f kubernetes 
```

