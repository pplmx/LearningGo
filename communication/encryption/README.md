# Self-signed Certificate

## create root key and crt

```shell
# Here, rootCA.key is the same as rootKey.pem. Only the file extensions are different.
# rootCA.crt <==> rootCrt.pem. The reason is the same as the above.
openssl req -x509 -nodes -sha256 -days 10240 -newkey rsa:4096 -keyout rootCA.key -out rootCA.crt \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost"
openssl req -x509 -nodes -sha256 \
	-newkey rsa:4096 \
	-days 10240 \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost" \
	-keyout root.key.pem \
	-out root.crt.pem
```

## Solution1. Self-signed Certificate by Owned CA

```shell
# OPENSSL_CONF="/etc/ssl/openssl.cnf"
OPENSSL_CONF="/System/Library/OpenSSL/openssl.cnf"
openssl req -new -nodes \
    -newkey rsa:4096 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost" \
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


## Solution2. Self-signed Certificate by Owned CA

This solution is much better.

```shell
MY_CONFIG="
[ req ]
default_bits=4096
distinguished_name=req_distinguished_name
x509_extensions=v3_ca
req_extensions=v3_req

[ req_distinguished_name ]
countryName                     = Country Name (2 letter code)
countryName_default             = CN
countryName_min                 = 2
countryName_max                 = 2

stateOrProvinceName             = State or Province Name (full name)
stateOrProvinceName_default     = Shanghai

localityName                    = Locality Name (eg, city)
localityName_default            = Shanghai

0.organizationName              = Organization Name (eg, company)
0.organizationName_default      = MyOrg, Inc.

organizationalUnitName          = Organizational Unit Name (eg, section)
organizationalUnitName_default  = Software Dept.

commonName                      = Common Name (eg, YOUR name)
commonName_default              = localhost
commonName_max                  = 64

emailAddress                    = Email Address
emailAddress_max                = 64

[ v3_req ]
subjectAltName=@alt_names
basicConstraints=CA:true

[ v3_ca ]
subjectAltName=@alt_names
basicConstraints=CA:true

[ alt_names ]
IP.1=127.0.0.1
IP.2=::1
DNS.1=localhost
"

openssl req -new -nodes \
    -newkey rsa:4096 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost" \
    -config <(echo "${MY_CONFIG}") \
    -keyout localhost.key.pem \
    -out localhost.csr

openssl x509 -req -sha256 -CAcreateserial -days 365 \
	-CA root.crt.pem \
	-CAkey root.key.pem \
	-extensions v3_ca \
	-extfile <(echo "${MY_CONFIG}") \
	-in localhost.csr \
	-out localhost.crt.pem

# show message
openssl rsa -in localhost.key.pem -noout -text
openssl req -in localhost.csr -noout -text
openssl x509 -in localhost.crt.pem -noout -text
```
