package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	"github.com/alibabacloud-go/tea/tea"
)

func CreateClientConfig(accessKeyId, accessKeySecret, regionId, endpoint *string) *openapi.Config {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	config.RegionId = regionId
	config.Endpoint = tea.String(*endpoint)
	return config
}
