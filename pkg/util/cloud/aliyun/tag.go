package aliyun

import (
	"fmt"

	tag20180828 "github.com/alibabacloud-go/tag-20180828/v2/client"
	_util "github.com/alibabacloud-go/tea-utils/v2/service"
	csvutil "github.com/atompi/autom/pkg/util/csv"

	"github.com/alibabacloud-go/tea/tea"
)

func SetTag(client *tag20180828.Client, data *[][]string) (err error) {
	srcData, err := csvutil.DataToMap(data)
	if err != nil {
		err = fmt.Errorf("convert data to map failed: %v", err)
		return
	}

	for _, row := range *srcData {
		arn, _err := GenARN(&row)
		if _err != nil {
			err = fmt.Errorf("generate arn failed: %v", err)
			return
		}
		regionId := row["RegionId"]
		if row["Tags"] == "" || row["Tags"] == "{}" || regionId == "" {
			err = fmt.Errorf("regionId or tag is empty")
			return
		}
		tagResourcesRequest := &tag20180828.TagResourcesRequest{
			RegionId:    tea.String(regionId),
			Tags:        tea.String(row["Tags"]),
			ResourceARN: []*string{tea.String(arn)},
		}
		runtime := &_util.RuntimeOptions{}

		_, _err = client.TagResourcesWithOptions(tagResourcesRequest, runtime)
		if _err != nil {
			err = fmt.Errorf("tag resources failed: %v", _err)
			return
		}
	}

	return
}
