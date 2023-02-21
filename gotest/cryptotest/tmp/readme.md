Step-1: Generate private key
First we would need a private key to generate the rootCA certificate:
[root@controller certs_x509]# openssl genrsa  -out cakey.pem 4096


Step-2: Create openssl configuration file
Next we will create one custom openssl configuration file required to generate the Certificate Signing request and add X.509 extensions to our RootCA certificate:

[root@controller certs_x509]# cat openssl.cnf
[ req ]
distinguished_name = req_distinguished_name
policy             = policy_match
x509_extensions     = v3_ca

# For the CA policy
[ policy_match ]
countryName             = optional
stateOrProvinceName     = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = supplied
emailAddress            = optional

[ req_distinguished_name ]
countryName                     = Country Name (2 letter code)
countryName_default             = IN
countryName_min                 = 2
countryName_max                 = 2
stateOrProvinceName             = State or Province Name (full name) ## Print this message
stateOrProvinceName_default     = KARNATAKA ## This is the default value
localityName                    = Locality Name (eg, city) ## Print this message
localityName_default            = BANGALORE ## This is the default value
0.organizationName              = Organization Name (eg, company) ## Print this message
0.organizationName_default      = GoLinuxCloud ## This is the default value
organizationalUnitName          = Organizational Unit Name (eg, section) ## Print this message
organizationalUnitName_default  = Admin ## This is the default value
commonName                      = Common Name (eg, your name or your server hostname) ## Print this message
commonName_max                  = 64
emailAddress                    = Email Address ## Print this message
emailAddress_max                = 64

[ v3_ca ]
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical,CA:true
nsComment = "OpenSSL Generated Certificate"


Here I have added

policy which must be applied to the RootCA certificate which is used while signing any certificate
distinguised_name which will be used to write the CSR
v3_ca field is used to define the X.509 extensions which will be added to the RootCA certificate


Step-3: Generate RootCA certificate
Let us go ahead and create our RootCA certificate:

[root@controller certs_x509]# openssl req -new -x509 -days 3650 -config openssl.cnf  -key cakey.pem -out cacert.pem


Step-4: Verify X.509 Extensions inside RootCA certificate
Our rootca certificate has successfully been created. let us verify the content of the certificate to make sure that our extensions were properly added:

[root@controller certs_x509]# openssl x509 -text -noout -in cacert.pem | grep -A10 "X509v3 extensions"
X509v3 extensions:
X509v3 Subject Key Identifier:
D2:84:32:48:45:86:23:E4:8F:02:22:BC:4D:E8:37:39:EF:FD:AF:7C
X509v3 Authority Key Identifier:
keyid:D2:84:32:48:45:86:23:E4:8F:02:22:BC:4D:E8:37:39:EF:FD:AF:7C

            X509v3 Basic Constraints: critical
                CA:TRUE
            Netscape Comment:
                OpenSSL Generated Certificate
    Signature Algorithm: sha256WithRSAEncryption


Scenario-2: Add X.509 extensions to Certificate Signing Request (CSR)
In this section I will share the steps required to add X.509 extensions to a certificate Signing request which can be used to sign any server or client certificate later.

