package etcd

import (
	"encoding/json"
	"fmt"
)

// etcdHealth encodes data returned from etcd /healthz handler.
type etcdHealth struct {
	// Note this has to be public so the json library can modify it.
	Health string `json:"health"`
}

// EtcdHealthCheck decodes data returned from etcd /healthz handler.
func EtcdHealthCheck(data []byte) error {
	obj := etcdHealth{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	if obj.Health != "true" {
		return fmt.Errorf("unhealthy status: %s", obj.Health)
	}
	return nil
}
