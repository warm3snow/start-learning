package tassl

type GmsslAddr struct {
	addr string
}

func (sslAddr *GmsslAddr) Network() string {
	return "tcp"
}
func (sslAddr *GmsslAddr) String() string {
	return sslAddr.addr
}
