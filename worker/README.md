config env (db using postgres)

### RUN
    go run main.go


### RUN TEST
    go test ./...

### BUILD
    export GOOS=linux
    go build

### BUILD image
    cd ..
    docker build -t blademaster996/micro-go micro-go/

### Make docker swarm
    docker-machine create \
      --driver virtualbox \
      --virtualbox-cpu-count 2 \
      --virtualbox-memory 2048 \
      --virtualbox-disk-size 20000 \
      swarm-manager-1

    docker $(docker-machine config swarm-manager-1) swarm init --advertise-addr $(docker-machine ip swarm-manager-1)

    docker network create --driver overlay my_network

    docker service create --name=movieservice --replicas=1 --network=my_network -p=3000:3000 blademaster996/micro-go