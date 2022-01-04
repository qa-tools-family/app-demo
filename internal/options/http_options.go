package options

import (
	"fmt"
	"github.com/spf13/pflag"
)

type HttpOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	// BindPort is ignored when Listener is set, will serve HTTPS even with 0.
	BindPort int `json:"bind-port"    mapstructure:"bind-port"`
	// Mode
	Mode string `json:"mode"    mapstructure:"mode"`
	// Middlewares
	Middlewares []string `json:"middlewares"    mapstructure:"middlewares"`
}

func (s *HttpOptions) Validate() []error {
	var errors []error
	if s.BindPort < 80 || s.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--bind-port %v must be between 80 and 65535",
				s.BindPort,
			),
		)
	}
	return errors
}

func (s *HttpOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "http.bind-address", s.BindAddress, ""+
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&s.BindPort, "http.bind-port", s.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated grpc access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")

	fs.StringVar(&s.Mode, "http.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")
}

func (s *HttpOptions) Complete() {
	if s.BindPort == 0 {
		s.BindPort = 8081
	}
	if s.BindAddress == "" {
		s.BindAddress = "0.0.0.0"
	}
}

func NewHttpOptions() *HttpOptions {
	return &HttpOptions{
		Mode: "release",
		Middlewares: []string{},
	}
}
