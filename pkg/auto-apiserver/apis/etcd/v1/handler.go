package v1

import (
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
