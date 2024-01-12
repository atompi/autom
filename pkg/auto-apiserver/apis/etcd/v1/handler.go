package v1

import (
	"net/http"

	"github.com/atompi/autom/pkg/auto-apiserver/handler"
	etcdutil "github.com/atompi/autom/pkg/auto-apiserver/util/etcd"
	"github.com/gin-gonic/gin"
)

func ListMembersHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(
		opts.Etcd.Endpoints,
		opts.Etcd.Tls.Ca,
		opts.Etcd.Tls.Cert,
		opts.Etcd.Tls.Key,
		opts.Etcd.DialTimeout,
	)
	defer etcdClient.Close()

	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	resp, err := etcdutil.GetMemberList(etcdClient, opts.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "etcd cluster not health"})
		return
	}

	m := resp.Members
	c.GinContext.JSON(http.StatusOK, gin.H{"response": m})
}
