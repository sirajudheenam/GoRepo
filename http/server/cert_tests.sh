#!/bin/bash
# This script runs tests on the generated certificates by making HTTPS requests to a local server.
# It assumes that the server is running on localhost:8443 and requires client authentication.
# Make sure to run the server before executing this script.

# Exit on any error
set -e

openssl s_client -connect localhost:8443 \
  -cert certs/client.crt -key certs/client.key -CAfile certs/ca.crt
openssl s_client -connect localhost:8443 -CAfile certs/ca.crt