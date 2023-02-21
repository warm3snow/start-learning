package tassl

import "C"
import (
	"bytes"
	"errors"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

//#cgo CFLAGS: -g -O2 -I/usr/local/include
//#cgo darwin LDFLAGS: ${SRCDIR}/../tassl/darwin/libssl.a ${SRCDIR}/../tassl/darwin/libcrypto.a -ldl
//#cgo linux LDFLAGS: ${SRCDIR}/../tassl/linux/libssl.a ${SRCDIR}/../tassl/linux/libcrypto.a -ldl

var Mutex sync.Mutex

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
