What is this?
=============

> **Warning**
> DO NOT use any of these certificates in production environments!

For testing LDAP locally, we need a test certificate authority and some test certificates.
This directory contains some ready-made certificates for use in testing scenarios.

If these certificates expire, below is documentation about how to generate more certificates.

Setup
-----

Use the tool [step](https://smallstep.com) to generate all required certificates (CA, server- and client certificate).

1. `step certificate create --ca root_ca.crt --ca-key root_ca_key --not-after "8760h" localhost server.crt server.key`
2. `step certificate create --ca root_ca.crt --ca-key root_ca_key --not-after "8760h" --template client.template "domain-verificator-project" client.crt client.key`
3. Remove password from `server.key`: `step crypto change-pass server.key --no-password --insecure`
4. Remove password from `client.key`: `step crypto change-pass client.key --no-password --insecure`
5. Move `server.*` to `certs`
