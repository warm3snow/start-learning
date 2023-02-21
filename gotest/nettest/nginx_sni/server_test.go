package nginx_sni

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type HelloStruct struct {
	msg string
}

func (h HelloStruct) sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("I'm %s", h.msg)))
}

func TestServer1(t *testing.T) {
	//DNS:chainmaker.org, DNS:localhost, DNS:common1.tls.wx-org1.chainmaker.org, IP Address:127.0.0.1
	http.HandleFunc("/", HelloStruct{msg: "Server1"}.sayHello)
	err := http.ListenAndServeTLS(":9090", "../certs/SS.crt", "../certs/SS.key", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestServer2(t *testing.T) {
	//DNS:chainmaker.org, DNS:localhost, DNS:consensus1.tls.wx-org1.chainmaker.org, IP Address:127.0.0.1
	http.HandleFunc("/", HelloStruct{msg: "Server2"}.sayHello)
	err := http.ListenAndServeTLS(":9091", "../certs/server.crt", "../certs/server.key", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
