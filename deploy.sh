#!/usr/bin/env bash

# $SHA is string unique to force restart 'kubectl deployment'
docker build -t blademaster996/micro-web-client:latest -t blademaster996/micro-web-client:$SHA -f ./client/docker/Dockerfile.production ./client
docker build -t blademaster996/micro-web-server:latest -t blademaster996/micro-web-server:$SHA -f ./server/Dockerfile.production ./server
docker build -t blademaster996/micro-web-worker:latest -t blademaster996/micro-web-worker:$SHA -f ./worker/Dockerfile.production ./worker

docker push blademaster996/micro-web-client:latest
docker push blademaster996/micro-web-server:latest
docker push blademaster996/micro-web-worker:latest

docker push blademaster996/micro-web-client:$SHA
docker push blademaster996/micro-web-server:$SHA
docker push blademaster996/micro-web-worker:$SHA

kubectl apply -f kubernetes

kubectl set image deployments/client-deployment client=blademaster996/micro-web-client:$SHA
kubectl set image deployments/server-deployment server=blademaster996/micro-web-server:$SHA
kubectl set image deployments/worker-deployment worker=blademaster996/micro-web-worker:$SHA
