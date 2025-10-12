#!/bin/bash
# This script generates a self-signed CA certificate and a server certificate signed by this CA.
# 1. Generate a private key for the CA
openssl genrsa -out ca.key 4096

# 2. Create a self-signed certificate for the CA (valid for 10 years)
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt \
  -subj "/C=DE/ST=Berlin/L=Berlin/O=LocalCA/OU=Dev/CN=MyLocalCA"

# 1. Generate server private key
openssl genrsa -out server.key 2048

# 2. Create a certificate signing request (CSR)
openssl req -new -key server.key -out server.csr \
  -subj "/C=DE/ST=Berlin/L=Berlin/O=Server/OU=Dev/CN=localhost"

# 3. Create a configuration file for the extensions
cat > server.ext <<EOF
subjectAltName = DNS:localhost,IP:127.0.0.1
EOF

# 4. Sign the server certificate with the CA (valid for 1 year)
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.ext

# You now have:
# - server.crt — the certificate you’ll serve to clients
# - server.key — private key for encryption

# Step 3. Create the client certificate

# 1. Generate client private key
openssl genrsa -out client.key 2048

# 2. Create a CSR for client
openssl req -new -key client.key -out client.csr \
  -subj "/C=DE/ST=Berlin/L=Berlin/O=Client/OU=Dev/CN=client"

# 3. Sign it with your CA
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out client.crt -days 365 -sha256

# You now have:
# - client.crt — the client certificate
# - client.key — private key for client authentication

# File	Role
# ca.crt	Trusted CA certificate
# ca.key	CA private key (keep safe)
# server.crt	Server public certificate
# server.key	Server private key
# client.crt	Client public certificate
# client.key	Client private key