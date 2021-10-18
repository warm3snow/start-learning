module chainmaker.org/gotest

go 1.16

require (
	chainmaker.org/chainmaker/common/v2 v2.0.1-0.20211015124616-4d3c85fa0a79
	github.com/Hyperledger-TWGC/tjfoc-gm v1.4.0
	github.com/miekg/pkcs11 v1.0.3
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	github.com/tjfoc/gmsm v1.4.1
)

replace chainmaker.org/chainmaker/common/v2 => chainmaker.org/chainmaker/common/v2 v2.0.1-0.20211015124616-4d3c85fa0a79 // indirect
