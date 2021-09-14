# Self-signed Certificate

## create ca.key and ca.crt

### one command

```shell
# one command: create ca.key and ca.crt(key.pem and cert.pem)
openssl req -x509 -sha256 -days 10240 -newkey rsa:4096 -keyout rootCA.key -out rootCA.crt \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics/OU=Software Dept/CN=www.example.com"
openssl req -x509 \
	-sha256 \
	-newkey rsa:4096 \
	-days 10240 \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics/OU=Software Dept/CN=www.example.com" \
	-keyout key.pem \
	-out cert.pem
```

### multiple command

```shell
# create root key === MUST NOT BE EXPOSED
openssl genrsa -aes256 -out ca.key 4096

# create crt with csr
openssl req -new -key ca.key -out ca.csr \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics/OU=Software Dept/CN=www.example.com"
openssl x509 -req -days 365 -signkey ca.key -in ca.csr -out ca.crt

# create crt without csr
openssl req -x509 -new -key ca.key -sha256 -days 365 -out ca.crt \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics/OU=Software Dept/CN=www.example.com"

```

## Self-signed Certificate by Owned CA

```shell
# 2. create server certificate
openssl genrsa -aes256 -out server.key 4096

# create crt with csr
openssl req -new -sha256 -key server.key -out server.csr \
	-subj "/C=CN/ST=Shanghai/L=Shanghai/O=Max Optics/OU=Software Dept/CN=www.example.com"
openssl x509 -req -sha256 -days 365 -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -in server.csr -out server.crt
```
