# Self-signed Certificate

> The deprecated, legacy behavior of treating the `CommonName` field on X.509 certificates as a host name when no Subject Alternative Names are present is now disabled by default. It can be temporarily re-enabled by adding the value `x509ignoreCN=0` to the `GODEBUG` environment variable.
>
>   Note that if the `CommonName` is an invalid host name, it's always ignored, regardless of `GODEBUG` settings. Invalid names include those with any characters other than letters, digits, hyphens and underscores, and those with empty labels or trailing dots.

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

## Self-signed Certificate by Owned CA

If you don't have the servers and clients, this section is enough for you.

If not, please reference the section **Create Server Certificate** and **Create Client Certificate**.

### create by config [Recommend]

```shell
MY_CONFIG="
[ req ]
default_bits                    = 4096
distinguished_name              = req_distinguished_name
req_extensions                  = v3_req
x509_extensions                 = v3_ca

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


```

### one command [Not Recommend]

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

## Create Server Certificate

```shell
# config file is at the following
openssl req -new -nodes \
    -newkey rsa:4096 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost" \
    -config crt_ext_server.cnf \
    -keyout server.key.pem \
    -out server.csr

openssl x509 -req -sha256 -CAcreateserial -days 365 \
	-CA root.crt.pem \
	-CAkey root.key.pem \
	-extensions v3_ca \
	-extfile crt_ext_server.cnf \
	-in server.csr \
	-out server.crt.pem
```

### cat crt_ext_server.cnf

```toml
[ req ]
default_bits                    = 4096
distinguished_name              = req_distinguished_name
req_extensions                  = v3_req
x509_extensions                 = v3_ca

[ req_distinguished_name ]

[ v3_req ]
basicConstraints                = CA:false
subjectAltName                  = @alt_names

[ v3_ca ]
basicConstraints                = CA:false
nsCertType                      = server
nsComment                       = "OpenSSL Generated Server Certificate"
keyUsage                        = critical, digitalSignature, keyEncipherment
extendedKeyUsage                = serverAuth
subjectAltName                  = @alt_names

[ alt_names ]
IP.1                            = 127.0.0.1
IP.2                            = ::1
DNS.1                           = localhost

```

## Create Client Certificate

```shell
# config file is at the following
openssl req -new -nodes \
    -newkey rsa:4096 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg, Inc./OU=Software Dept/CN=localhost" \
    -config crt_ext_client.cnf \
    -keyout client.key.pem \
    -out client.csr

openssl x509 -req -sha256 -CAcreateserial -days 365 \
	-CA root.crt.pem \
	-CAkey root.key.pem \
	-extensions v3_ca \
	-extfile crt_ext_client.cnf \
	-in client.csr \
	-out client.crt.pem
```

### cat crt_ext_client.cnf

```toml
[ req ]
default_bits                    = 4096
distinguished_name              = req_distinguished_name
req_extensions                  = v3_req
x509_extensions                 = v3_ca

[ req_distinguished_name ]

[ v3_req ]
basicConstraints                = CA:false
subjectAltName                  = @alt_names

[ v3_ca ]
basicConstraints                = CA:false
nsCertType                      = client, email
nsComment                       = "OpenSSL Generated Client Certificate"
keyUsage                        = critical, nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage                = clientAuth, emailProtection
subjectAltName                  = @alt_names

[ alt_names ]
IP.1                            = 127.0.0.1
IP.2                            = ::1
DNS.1                           = localhost

```

## Show Certificate Info

```shell
# show message
openssl rsa -in localhost.key.pem -noout -text
openssl req -in localhost.csr -noout -text
openssl x509 -in localhost.crt.pem -noout -text
```
