package v1

import (
	"io"
	"net/http"

	"github.com/atompi/autom/pkg/handler"
	etcdutil "github.com/atompi/autom/pkg/util/etcd"
	"github.com/gin-gonic/gin"
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

	b, err := io.ReadAll(c.GinContext.Request.Body)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot read request body"})
		return
	}
	key := string(b)

	resp, err := Get(etcdClient, key, opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot list etcd members"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": resp})
}
