package v1

import (
	"net/http"

	"github.com/atompi/autom/pkg/handler"
	etcdutil "github.com/atompi/autom/pkg/util/etcd"
	"github.com/gin-gonic/gin"
)

func ListOriginsHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(opts.APIServer.Etcd)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "cannot create etcd client"})
		return
	}
	defer etcdClient.Close()

	res, err := ListOrigins(etcdClient, opts.APIServer.Etcd.Prefix+"dns/", opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "get key value failed"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": res})
}

func ListRecordsHandler(c *handler.Context) {
	opts := c.Options
	etcdClient, err := etcdutil.New(opts.APIServer.Etcd)
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
	origin, ok := j["origin"]
	if !ok {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, no origin"})
		return
	}

	res, err := ListRecords(etcdClient, opts.APIServer.Etcd.Prefix+"dns/"+origin+"/", opts.APIServer.Etcd.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "get key value failed"})
		return
	}
	c.GinContext.JSON(http.StatusOK, gin.H{"response": res})
}
