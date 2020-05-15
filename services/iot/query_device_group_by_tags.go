package iot

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

// QueryDeviceGroupByTags invokes the iot.QueryDeviceGroupByTags API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicegroupbytags.html
func (client *Client) QueryDeviceGroupByTags(request *QueryDeviceGroupByTagsRequest) (response *QueryDeviceGroupByTagsResponse, err error) {
	response = CreateQueryDeviceGroupByTagsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceGroupByTagsWithChan invokes the iot.QueryDeviceGroupByTags API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicegroupbytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceGroupByTagsWithChan(request *QueryDeviceGroupByTagsRequest) (<-chan *QueryDeviceGroupByTagsResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceGroupByTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceGroupByTags(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryDeviceGroupByTagsWithCallback invokes the iot.QueryDeviceGroupByTags API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicegroupbytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceGroupByTagsWithCallback(request *QueryDeviceGroupByTagsRequest, callback func(response *QueryDeviceGroupByTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceGroupByTagsResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceGroupByTags(request)
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

// QueryDeviceGroupByTagsRequest is the request struct for api QueryDeviceGroupByTags
type QueryDeviceGroupByTagsRequest struct {
	*requests.RpcRequest
	IotInstanceId string                       `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer             `position:"Query" name:"PageSize"`
	Tag           *[]QueryDeviceGroupByTagsTag `position:"Query" name:"Tag"  type:"Repeated"`
	CurrentPage   requests.Integer             `position:"Query" name:"CurrentPage"`
	ApiProduct    string                       `position:"Body" name:"ApiProduct"`
	ApiRevision   string                       `position:"Body" name:"ApiRevision"`
}

// QueryDeviceGroupByTagsTag is a repeated param struct in QueryDeviceGroupByTagsRequest
type QueryDeviceGroupByTagsTag struct {
	TagValue string `name:"TagValue"`
	TagKey   string `name:"TagKey"`
}

// QueryDeviceGroupByTagsResponse is the response struct for api QueryDeviceGroupByTags
type QueryDeviceGroupByTagsResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Page         int                          `json:"Page" xml:"Page"`
	PageSize     int                          `json:"PageSize" xml:"PageSize"`
	PageCount    int                          `json:"PageCount" xml:"PageCount"`
	Total        int                          `json:"Total" xml:"Total"`
	Data         DataInQueryDeviceGroupByTags `json:"Data" xml:"Data"`
}

// CreateQueryDeviceGroupByTagsRequest creates a request to invoke QueryDeviceGroupByTags API
func CreateQueryDeviceGroupByTagsRequest() (request *QueryDeviceGroupByTagsRequest) {
	request = &QueryDeviceGroupByTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceGroupByTags", "Iot", "openAPI")
	return
}

// CreateQueryDeviceGroupByTagsResponse creates a response to parse from QueryDeviceGroupByTags response
func CreateQueryDeviceGroupByTagsResponse() (response *QueryDeviceGroupByTagsResponse) {
	response = &QueryDeviceGroupByTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
