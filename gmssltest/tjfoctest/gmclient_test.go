package gmcredentials

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/warm3snow/gmsm/gmtls"
	"github.com/warm3snow/gmsm/x509"

	cmtls "chainmaker.org/chainmaker/common/v2/crypto/tls"
)

var (
	caCert = "../tassl_demo_certs/CA.crt"
	csCert = "../tassl_demo_certs/CS.crt"
	csKey  = "../tassl_demo_certs/CS.key"

	ceCert = "../tassl_demo_certs/CE.crt"
	ceKey  = "../tassl_demo_certs/CE.key"
)

// gmGCMClientRun GCM模式测试
func TestGmClient(t *testing.T) {

	// 信任的根证书
	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile(caCert)
	assert.NoError(t, err)

	certPool.AppendCertsFromPEM(cacert)
	cs, err := gmtls.LoadX509KeyPair(csCert, csKey)
	assert.NoError(t, err)
	ce, err := gmtls.LoadX509KeyPair(ceCert, ceKey)
	assert.NoError(t, err)

	config := &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		RootCAs:      certPool,
		Certificates: []gmtls.Certificate{cs, ce},
		CipherSuites: []uint16{gmtls.GMTLS_SM2_WITH_SM4_SM3},
		//CipherSuites: []uint16{gmtls.GMTLS_ECDHE_SM2_WITH_SM4_SM3},
		ServerName: "server sign (SM2)",
		//InsecureSkipVerify: true,
	}

	conn, err := gmtls.Dial("tcp", "localhost:44330", config)
	assert.NoError(t, err)
	defer conn.Close()

	req := []byte("GET / HTTP/1.1\r\n" +
		"Host: localhost\r\n" +
		"Connection: close\r\n\r\n")
	_, _ = conn.Write(req)
	//buff := make([]byte, 1024)
	//for {
	//	n, _ := conn.Read(buff)
	//	if n <= 0 {
	//		break
	//	} else {
	//		fmt.Printf("%s", buff[0:n])
	//	}
	//}
	//fmt.Println(">> SM2_SM4_GCM_SM3 suite [PASS]")
	//end <- true
}

func TestLoadX509KeyPair(t *testing.T) {
	cs, err := cmtls.LoadX509KeyPair(csCert, csKey)
	assert.NoError(t, err)
	assert.NotNil(t, cs)

	csx, err := gmtls.LoadX509KeyPair(csCert, csKey)
	assert.NoError(t, err)
	assert.NotNil(t, csx)
}
