
```bash
go run main.go
curl "http://localhost:8080/user?id=42"
# {"id":"42","name":"User 42"}


docker build -t rest-server .
docker run --rm -p 8080:8080 rest-server

curl "http://localhost:8080/user?id=1"

# create a deployment
kubectl apply -f rest-deployment.yaml
```
