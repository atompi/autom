package aliyun

import (
	"fmt"
	"strings"

	resourcecenter "github.com/alibabacloud-go/resourcecenter-20221201/client"

	"github.com/alibabacloud-go/tea/tea"
)

func FilterResourceCenter(client *resourcecenter.Client, fKey string, mType string, fValues []string) (data *[][]string, err error) {
	nextToken := tea.String("")
	maxResults := tea.Int32(int32(100))
	sort := &resourcecenter.SearchResourcesRequestSortCriterion{
		Key:   tea.String("CreateTime"),
		Order: tea.String("ASC"),
	}

	head := []string{
		"AccountId",
		"ResourceId",
		"RegionId",
		"ZoneId",
		"ResourceName",
		"ResourceType",
		"CreateTime",
		"ResourceGroupId",
		"Tags",
		"IpAddresses",
	}
	res := [][]string{head}

	filter := &resourcecenter.SearchResourcesRequestFilter{
		Key:       tea.String(fKey),
		MatchType: tea.String(mType),
		Value:     tea.StringSlice(fValues),
	}

	for {
		request := &resourcecenter.SearchResourcesRequest{
			MaxResults:    maxResults,
			NextToken:     nextToken,
			SortCriterion: sort,
			Filter:        []*resourcecenter.SearchResourcesRequestFilter{filter},
		}

		response, _err := client.SearchResources(request)
		err = _err
		if err != nil {
			err = fmt.Errorf("request SearchResources failed: %v", err)
			return
		}

		body := response.Body
		for _, resource := range body.Resources {
			ss := []string{}
			for _, tag := range resource.Tags {
				s := "\"" + tea.StringValue(tag.Key) + "\":\"" + tea.StringValue(tag.Value) + "\""
				ss = append(ss, s)
			}
			tags := "{" + strings.Join(ss, ",") + "}"
			ss = []string{}
			for _, ip := range resource.IpAddresses {
				ss = append(ss, tea.StringValue(ip))
			}
			ips := strings.Join(ss, ",")
			res = append(
				res,
				[]string{
					tea.StringValue(resource.AccountId),
					tea.StringValue(resource.ResourceId),
					tea.StringValue(resource.RegionId),
					tea.StringValue(resource.ZoneId),
					tea.StringValue(resource.ResourceName),
					tea.StringValue(resource.ResourceType),
					tea.StringValue(resource.CreateTime),
					tea.StringValue(resource.ResourceGroupId),
					tags,
					ips,
				},
			)
		}

		nextToken = body.NextToken
		if nextToken == nil || tea.StringValue(nextToken) == "" {
			break
		}
	}

	data = &res

	return
}
