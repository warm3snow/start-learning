package credentials

import (
	"chainmaker.org/chainmaker/common/crypto/tls/credentials"
	"chainmaker.org/gmtlstest/grpctest/helloworld"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"

	cmtls "chainmaker.org/chainmaker/common/crypto/tls"
	cmx509 "chainmaker.org/chainmaker/common/crypto/x509"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

const (
	requestMsg  = "hello, I'm client"
	responseMsg = "hi, I'm server"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received %s", req.Name)
	return &helloworld.HelloReply{Message: responseMsg}, nil
}

//tls certs
const (
	ca         = "../testdata/cacert.pem"
	serverCert = "../testdata/servercert.pem"
	serverKey  = "../testdata/serverkey.pem"
	userCert   = "../testdata/usercert.pem"
	userKey    = "../testdata/userkey.pem"
)

//grpc server
func serverRunWithTls(t *testing.T) {
	signCert, err := cmtls.LoadX509KeyPair(serverCert, serverKey)
	require.NoError(t, err)

	certPool := cmx509.NewCertPool()
	cacert, err := ioutil.ReadFile(ca)
	require.NoError(t, err)

	certPool.AppendCertsFromPEM(cacert)
	lis, err := net.Listen("tcp", port)
	require.NoError(t, err)

	creds := credentials.NewTLS(&cmtls.Config{
		ClientAuth:   cmtls.RequireAndVerifyClientCert,
		Certificates: []cmtls.Certificate{signCert},
		ClientCAs:    certPool,
		MaxVersion:   cmtls.VersionTLS12,
	})

	s := grpc.NewServer(grpc.Creds(creds))
	helloworld.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)

	require.NoError(t, err)
}

func clientRunWithTls(t *testing.T, stop chan struct{}) {
	cert, err := cmtls.LoadX509KeyPair(userCert, userKey)
	require.NoError(t, err)

	certPool := cmx509.NewCertPool()
	cacert, err := ioutil.ReadFile(ca)
	require.NoError(t, err)

	certPool.AppendCertsFromPEM(cacert)
	creds := credentials.NewTLS(&cmtls.Config{
		ServerName:   "chainmaker.org",
		Certificates: []cmtls.Certificate{cert},
		RootCAs:      certPool,
		ClientAuth:   cmtls.RequireAndVerifyClientCert,
	})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	defer conn.Close()
	require.NoError(t, err)

	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: requestMsg})
	require.NoError(t, err)
	require.Equal(t, responseMsg, r.Message)

	stop <- struct{}{}
}

//grpc server
func serverRunNoTls(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	require.NoError(t, err)
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)

	require.NoError(t, err)
}

func Test_GrpcTlsWith2WayAuth(t *testing.T) {
	stop := make(chan struct{}, 1)
	go serverRunWithTls(t)
	time.Sleep(time.Second * 3) //wait for server start
	go clientRunWithTls(t, stop)
	<-stop
}

func Test_StartGrpcServer(t *testing.T) {
	serverRunWithTls(t)
}

func Test_StartGrpcServerNoTls(t *testing.T) {
	serverRunNoTls(t)
}
