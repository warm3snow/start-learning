#!/bin/bash

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=root-cert/CN=ca.wx-org5.chainmaker.org"
openssl ecparam -out ca.key -name prime256v1 -genkey
openssl req -key ca.key -new -out ca.csr -subj $SUBJ
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt -days 3650

rm -rf ca.csr

# 生成节点证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=common/CN=common1.sign.wx-org5.chainmaker.org"
openssl ecparam -out common1.sign.key -name prime256v1 -genkey
openssl req -key common1.sign.key -new -out common1.sign.csr -subj $SUBJ
openssl x509 -req -in common1.sign.csr -CA ca.crt -CAkey ca.key -out common1.sign.crt -days 3650

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=common/CN=common1.tls.wx-org5.chainmaker.org"
openssl ecparam -out common1.tls.key -name prime256v1 -genkey
openssl req -key common1.tls.key -new -out common1.tls.csr -subj $SUBJ
openssl x509 -req -in common1.tls.csr -CA ca.crt -CAkey ca.key -out common1.tls.crt -days 3650

rm -rf common1.sign.csr common1.tls.csr


# 生成管理员证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=admin/CN=admin1.sign.wx-org5.chainmaker.org"
openssl ecparam -out admin1.sign.key -name prime256v1 -genkey
openssl req -key admin1.sign.key -new -out admin1.sign.csr -subj $SUBJ
openssl x509 -req -in admin1.sign.csr -CA ca.crt -CAkey ca.key -out admin1.sign.crt -days 3650

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=admin/CN=admin1.tls.wx-org5.chainmaker.org"
openssl ecparam -out admin1.tls.key -name prime256v1 -genkey
openssl req -key admin1.tls.key -new -out admin1.tls.csr -subj $SUBJ
openssl x509 -req -in admin1.tls.csr -CA ca.crt -CAkey ca.key -out admin1.tls.crt -days 3650

rm -rf admin1.sign.csr admin1.tls.csr

# 生成普通用户证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=client/CN=client1.sign.wx-org5.chainmaker.org"
openssl ecparam -out client1.sign.key -name prime256v1 -genkey
openssl req -key client1.sign.key -new -out client1.sign.csr -subj $SUBJ
openssl x509 -req -in client1.sign.csr -CA ca.crt -CAkey ca.key -out client1.sign.crt -days 3650

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=client/CN=client1.tls.wx-org5.chainmaker.org"
openssl ecparam -out client1.tls.key -name prime256v1 -genkey
openssl req -key client1.tls.key -new -out client1.tls.csr -subj $SUBJ
openssl x509 -req -in client1.tls.csr -CA ca.crt -CAkey ca.key -out client1.tls.crt -days 3650

rm -rf client1.sign.csr client1.tls.csr