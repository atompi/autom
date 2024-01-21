package etcd

import (
	"crypto/tls"
	"time"

	"github.com/atompi/autom/pkg/options"
	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// New creates a new Etcd client
func New(opts options.EtcdOptions) (*clientv3.Client, error) {
	var err error
	var tlsConfig *tls.Config
	var tlsInfo transport.TLSInfo
	if !opts.Tls.InsecureSkipVerify {
		tlsInfo = transport.TLSInfo{
			CertFile:           opts.Tls.Cert,
			KeyFile:            opts.Tls.Key,
			TrustedCAFile:      opts.Tls.Ca,
			InsecureSkipVerify: opts.Tls.InsecureSkipVerify,
		}
	} else {
		tlsInfo = transport.TLSInfo{
			InsecureSkipVerify: true,
		}
	}
	tlsConfig, err = tlsInfo.ClientConfig()
	if err != nil {
		return nil, err
	}

	return clientv3.New(clientv3.Config{
		Endpoints:   opts.Endpoints,
		DialTimeout: time.Duration(opts.DialTimeout) * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(), // block until the underlying connection is up
		},
		TLS: tlsConfig,
	})
}
