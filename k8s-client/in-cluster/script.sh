#!/bin/sh
echo "Building the application..."
GOOS=linux go build -o app .

eval $(minikube docker-env)

docker rmi sirajudheenam/in-cluster:v0.0.1
docker rmi in-cluster

echo "Building the Docker image..."
docker build -t sirajudheenam/in-cluster:v0.0.1 .


echo "Pushing the Docker image..."
docker push sirajudheenam/in-cluster:v0.0.1

echo "Running the application..."
kubectl delete pod demo --ignore-not-found
kubectl run --rm -i demo --image=sirajudheenam/in-cluster:v0.0.1
# kubectl delete deployment demo
