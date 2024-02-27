package aliyun

import (
	"fmt"
	"strings"
)

func GenARN(data *map[string]string) (arn string, err error) {
	if (*data)["ResourceType"] == "" || (*data)["RegionId"] == "" || (*data)["AccountId"] == "" || (*data)["ResourceId"] == "" {
		err = fmt.Errorf("some field is empty")
		return
	}
	productCode := strings.ToLower(strings.Split((*data)["ResourceType"], "::")[1])
	resourceType := strings.ToLower(strings.Split((*data)["ResourceType"], "::")[2])
	arn = fmt.Sprintf("arn:acs:%s:%s:%s:%s/%s", productCode, (*data)["RegionId"], (*data)["AccountId"], resourceType, (*data)["ResourceId"])
	return
}
