package handler

import (
	"net/http"

	etcdutil "github.com/atompi/autom/pkg/util/etcd"
	"github.com/gin-gonic/gin"
)

func EtcdStatusHandler(c *Context) {
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
	_, err = etcdutil.CheckClusterHealth(etcdClient)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "etcd cluster not health"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": "etcd cluster is health"})
}
