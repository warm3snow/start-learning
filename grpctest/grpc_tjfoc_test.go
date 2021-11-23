package grpctest

import (
	"io/ioutil"
	"testing"

	"chainmaker.org/gotest/grpctest/helloworld"
	"github.com/Hyperledger-TWGC/tjfoc-gm/gmtls"
	"github.com/Hyperledger-TWGC/tjfoc-gm/gmtls/gmcredentials"
	"github.com/Hyperledger-TWGC/tjfoc-gm/x509"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestTjfocClient(t *testing.T) {
	cert, err := gmtls.LoadX509KeyPair(userCert, userKey)
	require.NoError(t, err)

	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile(ca)
	require.NoError(t, err)

	certPool.AppendCertsFromPEM(cacert)
	creds := gmcredentials.NewTLS(&gmtls.Config{
		ServerName:   "chainmaker.org",
		Certificates: []gmtls.Certificate{cert},
		RootCAs:      certPool,
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
	})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	defer conn.Close()
	require.NoError(t, err)

	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: requestMsg})
	require.NoError(t, err)
	require.Equal(t, responseMsg, r.Message)
}
