package options

import (
	"encoding/json"
	"github.com/qa-tools-family/go-common-kit/cli/flag"
)

type Options struct {
	// 保证配置的全的
	HttpOptions *HttpOptions    `json:"http"   mapstructure:"http"`
	GRPCOptions *GRPCOptions    `json:"grpc"   mapstructure:"grpc"`
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.HttpOptions.Validate()...)
	errs = append(errs, o.GRPCOptions.Validate()...)
	return errs
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.HttpOptions.AddFlags(fss.FlagSet("http"))
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

func (o *Options) Complete() error {
	o.HttpOptions.Complete()
	o.GRPCOptions.Complete()
	return nil
}

// NewOptions 生成类型
func NewOptions() *Options {
	o := Options{
		HttpOptions: NewHttpOptions(),
		GRPCOptions: NewGRPCOptions(),
	}
	return &o
}
