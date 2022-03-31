package http2

import (
	"crypto/tls"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/http2"

	"golang.org/x/crypto/acme/autocert"
)

func TestHttps(t *testing.T) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

	log.Fatalln(server.ListenAndServeTLS("", ""))
}

func TestHttp2(t *testing.T) {
	srv := http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http2"))
	})

	http2.ConfigureServer(&srv, &http2.Server{})

	log.Fatalln(srv.ListenAndServeTLS("./server.crt", "./server.key"))
}

func TestH2c(t *testing.T) {

}
