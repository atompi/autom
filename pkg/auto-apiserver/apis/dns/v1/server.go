package v1

import (
	"context"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetOrigins(c *clientv3.Client, key string, timeout int) (res []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	var resp *clientv3.GetResponse
	resp, err = c.Get(
		ctx,
		key,
		clientv3.WithKeysOnly(),
		clientv3.WithPrefix(),
	)
	cancel()
	if err != nil {
		return
	}

	kvs := resp.Kvs
	for _, kv := range kvs {
		k := string(kv.Key)
		if strings.Count(k, "/") > 2 {
			continue
		}
		res = append(res, k[len(key):])
	}
	return
}
