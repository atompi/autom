package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	tag20180828 "github.com/alibabacloud-go/tag-20180828/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateTag20180828Client(accessKeyId, accessKeySecret, regionId, endpoint *string) (*tag20180828.Client, error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	config.RegionId = regionId
	config.Endpoint = tea.String(*endpoint)

	return tag20180828.NewClient(config)
}
