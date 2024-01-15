package v1

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// GetMemberList returns members list of etcd cluster
func GetMemberList(c *clientv3.Client, timeout int) (*clientv3.MemberListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	resp, err := c.Cluster.MemberList(ctx)
	cancel()
	return resp, err
}
