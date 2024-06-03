# TLS with Go

### Generate Certs:

```bash
# Generate server key
openssl genrsa -out server.key 2048

# Generate server certificate
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365 -subj "/CN=localhost"

# Generate client key
openssl genrsa -out client.key 2048

# Generate client certificate signing request
openssl req -new -key client.key -out client.csr -subj "/CN=client"

# Generate client certificate
openssl x509 -req -in client.csr -CA server.crt -CAkey server.key -CAcreateserial -out client.crt -days 365 -sha256
```

### Building and Running the Containers

#### To build and run the server:

```bash
docker build -t tls-server -f Dockerfile.server .
docker run -p 8443:8443 tls-server

```

#### To build and run the client

```bash
docker build -t tls-client -f Dockerfile.client .
docker run tls-client
```
