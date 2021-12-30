package openssl

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/spacemonkeygo/openssl"
)

var (
	serverCrt = "../testdata/servercert.pem"
	serverKey = "../testdata/serverkey.pem"

	clientCrt = "../testdata/usercert.pem"
	clientKey = "../testdata/userkey.pem"

	caCrt = "../testdata/cacert.pem"
)

func TestSSL(t *testing.T) {
	server()
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})

	err := openssl.ListenAndServeTLS(":8081", serverCrt, serverKey, nil)
	log.Fatal(err)
}
