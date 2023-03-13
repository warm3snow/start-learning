package main

import "C"
import (
	"bytes"
	"fmt"
	"github.com/warm3snow/start-learning/gotest/crypto/gmssltest/tassltest/tassl"
	"io/ioutil"
	"net/http"
)

func main() {
	// CA证书
	caFile := "./ca.crt"
	// 客户端证书
	certFile := "./cs.p12"
	// 客户端加密证书
	certFile2 := "./ce.p12"
	// 证书加密密码
	XgsCertPassword := "123456"
	// 访问地址
	url := "http://10.75.2.245:443"

	tassl.SetCert(caFile, certFile, certFile2, XgsCertPassword)
	client := tassl.CreateHttp()
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
