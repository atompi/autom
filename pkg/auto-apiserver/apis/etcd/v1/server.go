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

func Get(c *clientv3.Client, prefix bool, key string, timeout int) (res []map[string]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	var resp *clientv3.GetResponse
	if prefix {
		resp, err = c.Get(
			ctx,
			key,
			clientv3.WithPrefix(),
		)
	} else {
		resp, err = c.Get(ctx, key)
	}
	cancel()
	if err != nil {
		return
	}

	kvs := resp.Kvs
	for _, kv := range kvs {
		r := map[string]string{}
		r[string(kv.Key)] = string(kv.Value)
		res = append(res, r)
	}
	return
}

func Put(c *clientv3.Client, key string, value string, timeout int) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	resp, err := c.Put(ctx, key, value)
	cancel()
	return resp, err
}
