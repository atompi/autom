package v1

import (
	"net/http"

	"github.com/atompi/autom/pkg/handler"
	etcdutil "github.com/atompi/autom/pkg/util/etcd"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func ListMembersHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(
		opts.APIServer.Etcd.Endpoints,
		opts.APIServer.Etcd.Tls.Ca,
		opts.APIServer.Etcd.Tls.Cert,
		opts.APIServer.Etcd.Tls.Key,
		opts.APIServer.Etcd.DialTimeout,
	)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	defer etcdClient.Close()

	resp, err := GetMemberList(etcdClient, opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot list etcd members"})
		return
	}

	m := resp.Members
	c.GinContext.JSON(http.StatusOK, gin.H{"response": m})
}

func GetHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(
		opts.APIServer.Etcd.Endpoints,
		opts.APIServer.Etcd.Tls.Ca,
		opts.APIServer.Etcd.Tls.Cert,
		opts.APIServer.Etcd.Tls.Key,
		opts.APIServer.Etcd.DialTimeout,
	)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	defer etcdClient.Close()

	j := make(map[string]string)
	err = c.GinContext.BindJSON(&j)
	if err != nil {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, body must be in json format"})
		return
	}
	key, ok := j["key"]
	if !ok {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, no key"})
		return
	}
	prefix := j["prefix"]
	var resp *clientv3.GetResponse
	if prefix == "true" {
		resp, err = GetWithPrefix(etcdClient, key, opts.APIServer.Etcd.DialTimeout)
	} else {
		resp, err = Get(etcdClient, key, opts.APIServer.Etcd.DialTimeout)
	}
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "get key value failed"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": resp})
}

func PutHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(
		opts.APIServer.Etcd.Endpoints,
		opts.APIServer.Etcd.Tls.Ca,
		opts.APIServer.Etcd.Tls.Cert,
		opts.APIServer.Etcd.Tls.Key,
		opts.APIServer.Etcd.DialTimeout,
	)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	defer etcdClient.Close()

	j := make(map[string]string)
	err = c.GinContext.BindJSON(&j)
	if err != nil {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, body must be in json format"})
		return
	}
	key, ok := j["key"]
	if !ok {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, no key"})
		return
	}
	value, ok := j["value"]
	if !ok {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, no value"})
		return
	}
	resp, err := Put(etcdClient, key, value, opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "put key value failed"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": resp})
}
