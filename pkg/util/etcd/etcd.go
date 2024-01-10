package etcd

import (
	"context"
	"crypto/tls"
	"time"

	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// New creates a new Etcd client
func New(endpoints []string, ca, cert, key string, dialTimeout int) (*clientv3.Client, error) {
	var err error
	var tlsConfig *tls.Config
	if ca != "" || cert != "" || key != "" {
		tlsInfo := transport.TLSInfo{
			CertFile:      cert,
			KeyFile:       key,
			TrustedCAFile: ca,
		}
		tlsConfig, err = tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}
	}

	return clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Duration(dialTimeout) * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(), // block until the underlying connection is up
		},
		TLS: tlsConfig,
	})
}

// GetMemberList returns members list of etcd cluster
func GetMemberList(c *clientv3.Client, timeout int) (*clientv3.MemberListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	resp, err := c.Cluster.MemberList(ctx)
	cancel()
	return resp, err
}