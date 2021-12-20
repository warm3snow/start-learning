package main

import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

/*
#cgo CFLAGS: -g -O2 -I/usr/local/include
#cgo darwin LDFLAGS: ${SRCDIR}/../tassl/darwin/libssl.a ${SRCDIR}/../tassl/darwin/libcrypto.a -ldl
#cgo linux LDFLAGS: ${SRCDIR}/../tassl/linux/libssl.a ${SRCDIR}/../tassl/linux/libcrypto.a -ldl

#include "stdio.h"
#include "openssl/ssl.h"
#include "openssl/pkcs12.h"
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>
#include <fcntl.h>

static SSL_CTX* gmssl_ctx = NULL;
static SSL_SESSION *reuse_session = NULL;

static int load_pkcs12(BIO *in, EVP_PKEY **pkey, X509 **cert, STACK_OF(X509) **ca, char *password)
{
	char tpass[PEM_BUFSIZE] = {0,};
	int ret = 0;
	PKCS12 *p12;
	if (password) {
		snprintf(tpass, sizeof(tpass), "%s", password);
	}
	p12 = d2i_PKCS12_bio(in, NULL);
	if (p12 == NULL) {
		printf("Error loading PKCS12 file \n");
		goto die;
	}
	if (!PKCS12_verify_mac(p12, tpass, strlen(tpass))) {
		printf("Mac verify error (wrong password?) in PKCS12 file \n");
		goto die;
	}
	ret = PKCS12_parse(p12, tpass, pkey, cert, ca);
	die:
		PKCS12_free(p12);
	return ret;
}

void gmssl_init(int gmssl) {
	const SSL_METHOD *method;
	SSLeay_add_ssl_algorithms();
	SSL_load_error_strings();
	if (gmssl) {
		method = CNTLS_client_method();
	} else {
		method = TLS_client_method();
	}
	gmssl_ctx = SSL_CTX_new(method);
	if (!gmssl_ctx) {
		printf("gm ctx init failed\n");
	}
}

void gmssl_set_cert2(char *cert_S_file, char *cert_E_file, char *ca_file, char *password) {
	if (!gmssl_ctx) {
		printf("no gm ctx\n");
		return;
	}

	if (!cert_S_file || !cert_E_file || !ca_file) {
		return;
	}
	SSL_CTX_load_verify_locations(gmssl_ctx, ca_file, NULL);
	X509 *x = NULL;
	EVP_PKEY *pkey = NULL;
	BIO *cert;
	int ret;
	cert = BIO_new_file(cert_S_file, "rb");
	ret = load_pkcs12(cert, &pkey, &x, NULL, password);
	printf("load file %s %p ret %d\n", cert_S_file, x, ret);
	BIO_free(cert);
	if (x) {
		SSL_CTX_use_certificate(gmssl_ctx, x);
	}
	if (pkey) {
		SSL_CTX_use_PrivateKey(gmssl_ctx, pkey);
	}

	cert = BIO_new_file(cert_E_file, "rb");
	ret = load_pkcs12(cert, &pkey, &x, NULL, password);
	printf("load file %s %p ret %d\n", cert_E_file, x, ret);
	BIO_free(cert);
	if (x) {
		SSL_CTX_use_enc_certificate(gmssl_ctx, x);
	}
	if (pkey) {
		SSL_CTX_use_enc_PrivateKey(gmssl_ctx, pkey);
	}

	if (reuse_session) {
		SSL_SESSION_free(reuse_session);
		reuse_session = NULL;
	}
	return;
}

void * gmssl_ssl_connect(int fd) {
	int ret;
	SSL *ssl;
	ssl = SSL_new(gmssl_ctx);
	if (ssl == NULL) {
		printf("ssl new failed\n");
		close(fd);
		return 0;
	}
	//SSL_set_cipher_list(ssl, "SM2DHE-WITH-SMS4-SM3");
	SSL_set_fd(ssl, fd);
	if (reuse_session) {
		SSL_set_session(ssl, reuse_session);
	}
	ret = SSL_connect(ssl);
	if (ret == -1) {
		printf("ssl connect failed ret %d %s\n", ret, SSL_state_string_long(ssl));
		printf(stderr);
		if (reuse_session) {
			SSL_SESSION_free(reuse_session);
			reuse_session = NULL;
		}
		SSL_shutdown(ssl);
		close(fd);
		SSL_free(ssl);
		return 0;
	}
	if (!reuse_session) {
		reuse_session = SSL_get1_session(ssl);
	}
	return (void *)ssl;
}

int gmssl_read(void *fd, void *buf, int num) {
	SSL *ssl = (SSL *)fd;
	return SSL_read(ssl, buf, num);
}

int gmssl_write(void *fd, void *buf, int num) {
	SSL *ssl = (SSL *)fd;
	return SSL_write(ssl, (const void *)buf, num);
}

void gmssl_close(void *ssl_fd, int socket_fd) {
	SSL *ssl = (SSL *)ssl_fd;
	SSL_shutdown(ssl);
	close(socket_fd);
	SSL_free(ssl);
	return;
}

void gmssl_destroy_ctx(void) {
	SSL_CTX_free(gmssl_ctx);
	gmssl_ctx = NULL;
}


void go_gmssl_set_cert(void *cert_file, int cert_file_len, void *cert_file_enc, int cert_file_len_enc, void *ca_file, int ca_file_len, void *password, int password_len) {
        char cert_file_str[1024] = {0,};
        char cert_file_str_enc[1024] = {0,};
        char ca_file_str[1024] = {0,};
        char password_str[1024] = {0,};
        memcpy(cert_file_str, cert_file, cert_file_len >= 1024 ? 1023 : cert_file_len);
        memcpy(cert_file_str_enc, cert_file_enc, cert_file_len_enc >= 1024 ? 1023 : cert_file_len_enc);
        memcpy(ca_file_str, ca_file, ca_file_len >= 1024 ? 1023 : ca_file_len);
        memcpy(password_str, password, password_len >= 1024 ? 1023 : password_len);
        gmssl_set_cert2(cert_file_str, cert_file_str_enc, ca_file_str, password_str);
}

int gmssl_socket_connect(void *server, int port) {
	int client_socket;
	struct sockaddr_in addr_server;
	client_socket = socket(AF_INET, SOCK_STREAM, 0);

	memset(&addr_server, 0, sizeof(addr_server));
	addr_server.sin_family = AF_INET;
	addr_server.sin_addr.s_addr = inet_addr((char *)server);
	addr_server.sin_port = htons(port);

	int ret;
	ret = connect(client_socket, (struct sockaddr*) &addr_server, sizeof(addr_server));
	if (ret == -1) {
			printf("connect failed\n");
			close(client_socket);
			return -1;
	}

	return client_socket;
}

void gmssl_socket_set_nonblock(int client_socket) {
	int flags = fcntl(client_socket, F_GETFL, 0);
	fcntl(client_socket, F_SETFL, flags|O_NONBLOCK);
}
*/
import "C"

var Mutex sync.Mutex

func main() {
	// CA证书
	caFile := "./ca.crt"
	// 客户端证书
	certFile := "./client.p12"
	// 客户端加密证书
	certFile2 := "./client_enc.p12"
	// 证书加密密码
	XgsCertPassword := "xxxxxx"
	// 访问地址
	url := "http://10.75.2.245:443"

	_, caErr := os.Stat(caFile)
	_, certErr := os.Stat(certFile)
	_, certErr2 := os.Stat(certFile2)

	fmt.Println(caErr)
	fmt.Println(certErr)
	fmt.Println(certErr2)

	SetCert(caFile, certFile, certFile2, XgsCertPassword)
	client := CreateHttp()
	requestBody := bytes.NewReader([]byte(""))
	req, _ := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 返回值
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(respBody))
}

// 构建 https Client
func CreateHttp() *http.Client {
	client := http.Client{
		Transport: &http.Transport{
			Dial:                GmDial,
			MaxIdleConns:        8,
			MaxIdleConnsPerHost: 2,
		},
	}
	return &client
}

type GmsslAddr struct {
	addr string
}

func (sslAddr *GmsslAddr) Network() string {
	return "tcp"
}
func (sslAddr *GmsslAddr) String() string {
	return sslAddr.addr
}

type GmsslConn struct {
	socketFd  C.int
	sslFd     unsafe.Pointer
	addr      string
	connMutex sync.Mutex
	connState int
	deadline  time.Time
	ref       int
}

func (sslConn *GmsslConn) Read(b []byte) (int, error) {
	sslConn.connMutex.Lock()
	if sslConn.connState == 1 {
		sslConn.connMutex.Unlock()
		return 0, errors.New("conn closed")
	}
	sslConn.ref += 1
	sslConn.connMutex.Unlock()
	ret := C.gmssl_read(sslConn.sslFd, unsafe.Pointer(&b[0]), C.int(len(b)))
	sslConn.connMutex.Lock()
	sslConn.ref -= 1
	sslConn.connMutex.Unlock()

	fixTarget := []byte("HTTP/1.1 301")
	if int(ret) >= len(fixTarget) {
		if bytes.Equal(b[0:len(fixTarget)], fixTarget) == true {
			b[len(fixTarget)-1] = byte('7')
		}
	}
	if int(ret) <= 0 {
		return 0, errors.New("ssl read err")
	} else {
		return int(ret), nil
	}
}

func (sslConn *GmsslConn) Write(b []byte) (int, error) {
	sslConn.connMutex.Lock()
	if sslConn.connState == 1 {
		sslConn.connMutex.Unlock()
		return 0, errors.New("conn closed")
	}
	sslConn.ref += 1
	sslConn.connMutex.Unlock()
	ret := int(C.gmssl_write(sslConn.sslFd, unsafe.Pointer(&b[0]), C.int(len(b))))
	sslConn.connMutex.Lock()
	sslConn.ref -= 1
	sslConn.connMutex.Unlock()
	if ret <= 0 {
		return 0, errors.New("ssl write err")
	} else {
		return ret, nil
	}
}

func (sslConn *GmsslConn) Close() error {
	C.gmssl_socket_set_nonblock(sslConn.socketFd)
	sslConn.connMutex.Lock()
	for sslConn.ref != 0 {
		sslConn.connMutex.Unlock()
		time.Sleep(time.Duration(100) * time.Millisecond)
		sslConn.connMutex.Lock()
	}
	sslConn.connState = 1
	if sslConn.sslFd != nil {
		C.gmssl_close(sslConn.sslFd, sslConn.socketFd)
		sslConn.sslFd = nil
	}
	sslConn.connMutex.Unlock()
	return nil
}

func (sslConn *GmsslConn) LocalAddr() net.Addr {
	return nil
}

func (sslConn *GmsslConn) RemoteAddr() net.Addr {
	return &GmsslAddr{sslConn.addr}
}

func (sslConn *GmsslConn) SetDeadline(t time.Time) error {
	sslConn.deadline = t
	return nil
}

func (sslConn *GmsslConn) SetReadDeadline(t time.Time) error {
	sslConn.deadline = t
	return nil
}

func (sslConn *GmsslConn) SetWriteDeadline(t time.Time) error {
	sslConn.deadline = t
	return nil
}

func GmDial(_, addr string) (net.Conn, error) {
	var conn net.Conn
	sslConn := new(GmsslConn)
	sslConn.socketFd = 0
	sslConn.sslFd = nil
	sslConn.addr = addr
	sslConn.connState = 0
	sslConn.connMutex = sync.Mutex{}
	sslConn.ref = 0

	now := time.Now()
	sslConn.deadline = now.Add(time.Hour)

	server := strings.Split(addr, ":")[0]
	port, _ := strconv.Atoi(strings.Split(addr, ":")[1])
	socketFd := C.gmssl_socket_connect(unsafe.Pointer(&[]byte(server)[0]), C.int(port))
	if C.int(socketFd) == -1 {
		return nil, errors.New("socket create failed")
	}
	sslConn.socketFd = socketFd
	Mutex.Lock()
	sslFd := C.gmssl_ssl_connect(socketFd)
	Mutex.Unlock()
	if sslFd == nil {
		return nil, errors.New("ssl create failed")
	}
	sslConn.sslFd = sslFd

	conn = sslConn
	return conn, nil
}

func SetCert(caFile string, certFile string, certFileEnc string, certPassword string) {
	C.go_gmssl_set_cert(unsafe.Pointer(&[]byte(certFile)[0]), C.int(len(certFile)),
		unsafe.Pointer(&[]byte(certFileEnc)[0]), C.int(len(certFileEnc)),
		unsafe.Pointer(&[]byte(caFile)[0]), C.int(len(caFile)),
		unsafe.Pointer(&[]byte(certPassword)[0]), C.int(len(certPassword)))
}
