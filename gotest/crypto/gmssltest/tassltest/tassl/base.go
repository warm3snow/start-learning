package tassl

/*
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
	return ret
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
		//ERR_print_errors_fp(stderr);
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
		gmssl_init(1);

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
