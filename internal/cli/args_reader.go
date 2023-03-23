package cli

import (
	"flag"
	"fmt"
	"net"
)

const (
	defaultAddress = "0.0.0.0"
	defaultPort    = 8090
	minPort        = 0
	maxPort        = 65535
)

type ArgsReader struct {
}

func NewArgsReader() *ArgsReader {
	return &ArgsReader{}
}

func (ar *ArgsReader) ReadArgs() (string, int, error) {
	var (
		address string
		port    int
	)

	flag.StringVar(&address, "addr", defaultAddress, "Provide server address")
	flag.IntVar(&port, "p", defaultPort, "Provide server port")
	flag.Parse()

	err := ar.validateAddress(address)
	if err != nil {
		return address, port, err
	}

	err = ar.validatePort(port)
	if err != nil {
		return address, port, err
	}

	return address, port, nil
}

func (*ArgsReader) validateAddress(address string) error {
	if net.ParseIP(address) == nil {
		return ArgsReaderError{
			errMessage: "Server address value is not valid",
		}
	}

	return nil
}

func (*ArgsReader) validatePort(port int) error {
	if port < minPort || port > maxPort {
		return ArgsReaderError{
			errMessage: fmt.Sprintf("Server port value is not in the range %d - %d: ", minPort, maxPort),
		}
	}

	return nil
}

type ArgsReaderError struct {
	errMessage string
}

func (re ArgsReaderError) Error() string {
	return re.errMessage
}
