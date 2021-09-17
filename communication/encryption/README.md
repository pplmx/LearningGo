# Self-signed Certificate

## create root key and crt

```shell
# Here, rootCA.key is the same as rootKey.pem. Only the file extensions are different.
# rootCA.crt <==> rootCrt.pem. The reason is the same as the above.
openssl req -x509 -nodes -sha256 -days 10240 -newkey rsa:4096 -keyout rootCA.key -out rootCA.crt \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics, Inc./OU=Software Dept/CN=localhost"
openssl req -x509 -nodes -sha256 \
	-newkey rsa:4096 \
	-days 10240 \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics, Inc./OU=Software Dept/CN=localhost" \
	-keyout root.key.pem \
	-out root.crt.pem
```

## Self-signed Certificate by Owned CA

```shell
# OPENSSL_CONF="/etc/ssl/openssl.cnf"
OPENSSL_CONF="/System/Library/OpenSSL/openssl.cnf"
openssl req -new -nodes \
    -newkey rsa:4096 \
    -subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics, Inc./OU=Software Dept/CN=localhost" \
	-reqexts SAN \
    -config <(cat "${OPENSSL_CONF}" \
        <(printf "\n[SAN]\nsubjectAltName=DNS:localhost")) \
    -keyout localhost.key.pem \
    -out localhost.csr

openssl x509 -req -sha256 -CAcreateserial -days 365 \
	-CA root.crt.pem \
	-CAkey root.key.pem \
	-extfile <(printf "subjectAltName=DNS:localhost") \
	-in localhost.csr \
	-out localhost.crt.pem
```
