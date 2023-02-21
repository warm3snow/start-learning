package wss

//
//import (
//	"context"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net"
//	"net/http"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	cmtls "chainmaker.org/chainmaker/common/v2/crypto/tls"
//
//	cmhttp "chainmaker.org/chainmaker/common/v2/crypto/tls/http"
//	"github.com/gorilla/websocket"
//)
//
//var (
//	path      = "/Users/hxy/go/src/chainmaker.org/gotest/chainmakertest/wss/"
//	caCert    = path + "cacert.pem"
//	serverCrt = path + "servercert.pem"
//	serverKey = path + "serverkey.pem"
//	clientCrt = path + "usercert.pem"
//	clientKey = path + "userkey.pem"
//)
//
//type msg struct {
//	Num int
//}
//
//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	content, err := ioutil.ReadFile("index.html")
//	if err != nil {
//		fmt.Println("Could not open file.", err)
//	}
//	fmt.Fprintf(w, "%s", content)
//}
//
//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Header.Get("Origin") != "http://"+r.Host {
//		http.Error(w, "Origin not allowed", 403)
//		return
//	}
//	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
//	if err != nil {
//		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
//	}
//
//	go echo(conn)
//}
//
//func echo(conn *websocket.Conn) {
//	for {
//		m := msg{}
//
//		err := conn.ReadJSON(&m)
//		if err != nil {
//			fmt.Println("Error reading json.", err)
//		}
//
//		fmt.Printf("Got message: %#v\n", m)
//
//		if err = conn.WriteJSON(m); err != nil {
//			fmt.Println(err)
//		}
//	}
//}
//
//func TestServer(t *testing.T) {
//	http.HandleFunc("/ws", wsHandler)
//	http.HandleFunc("/", rootHandler)
//
//	panic(cmhttp.ListenAndServeTLS(":8080", serverCrt, serverKey, caCert, http.DefaultServeMux))
//}
//
//func TestClient(t *testing.T) {
//	tlsConfig, err := cmhttp.GetConfig(clientCrt, clientKey, caCert, false)
//	client, res, err := NewClient(tlsConfig)
//	resp, err := client.Get(url)
//	assert.NoError(t, err)
//
//	buf, err := ioutil.ReadAll(resp.Body)
//	assert.NoError(t, err)
//	assert.Equal(t, msg, buf)
//	log.Println("receive from server: " + string(buf))
//}
//
//func NewClient(tlsConfig cmtls.Config) (websocket.Dialer, error) {
//	d := websocket.Dialer{
//		ReadBufferSize:  10,
//		WriteBufferSize: 10,
//
//		NetDialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
//			dialer := &net.Dialer{}
//			conn, err := cmtls.DialWithDialer(dialer, network, addr, tlsConfig)
//			if err != nil {
//				return nil, err
//			}
//
//			return conn, nil
//		},
//	}
//	return d, err
//}
