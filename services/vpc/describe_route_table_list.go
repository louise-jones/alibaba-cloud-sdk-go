package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeRouteTableList(request *DescribeRouteTableListRequest) (response *DescribeRouteTableListResponse, err error) {
	response = CreateDescribeRouteTableListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRouteTableListWithChan(request *DescribeRouteTableListRequest) (<-chan *DescribeRouteTableListResponse, <-chan error) {
	responseChan := make(chan *DescribeRouteTableListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouteTableList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeRouteTableListWithCallback(request *DescribeRouteTableListRequest, callback func(response *DescribeRouteTableListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouteTableListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouteTableList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeRouteTableListRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string           `position:"Query" name:"RouterId"`
	RouterType           string           `position:"Query" name:"RouterType"`
	RouteTableName       string           `position:"Query" name:"RouteTableName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeRouteTableListResponse struct {
	*responses.BaseResponse
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	Code            string           `json:"Code" xml:"Code"`
	Message         string           `json:"Message" xml:"Message"`
	Success         requests.Boolean `json:"Success" xml:"Success"`
	PageSize        requests.Integer `json:"PageSize" xml:"PageSize"`
	PageNumber      requests.Integer `json:"PageNumber" xml:"PageNumber"`
	TotalCount      requests.Integer `json:"TotalCount" xml:"TotalCount"`
	RouterTableList struct {
		RouterTableListType []struct {
			VpcId          string `json:"VpcId" xml:"VpcId"`
			RouterType     string `json:"RouterType" xml:"RouterType"`
			RouterId       string `json:"RouterId" xml:"RouterId"`
			RouteTableId   string `json:"RouteTableId" xml:"RouteTableId"`
			RouteTableName string `json:"RouteTableName" xml:"RouteTableName"`
			RouteTableType string `json:"RouteTableType" xml:"RouteTableType"`
			Description    string `json:"Description" xml:"Description"`
			CreationTime   string `json:"CreationTime" xml:"CreationTime"`
		} `json:"RouterTableListType" xml:"RouterTableListType"`
	} `json:"RouterTableList" xml:"RouterTableList"`
}

func CreateDescribeRouteTableListRequest() (request *DescribeRouteTableListRequest) {
	request = &DescribeRouteTableListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteTableList", "", "")
	return
}

func CreateDescribeRouteTableListResponse() (response *DescribeRouteTableListResponse) {
	response = &DescribeRouteTableListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}