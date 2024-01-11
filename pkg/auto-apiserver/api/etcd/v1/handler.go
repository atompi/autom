package v1

import (
	"net/http"

	etcdutil "github.com/atompi/autom/pkg/auto-apiserver/util/etcd"
	"github.com/atompi/autom/pkg/auto-apiserver/util/handler"
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
	defer etcdClient.Close()

	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	resp, err := etcdutil.GetMemberList(etcdClient, opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "etcd cluster not health"})
		return
	}

	m := resp.Members
	c.GinContext.JSON(http.StatusOK, gin.H{"response": m})
}
