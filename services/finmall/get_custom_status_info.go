package finmall

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

// GetCustomStatusInfo invokes the finmall.GetCustomStatusInfo API synchronously
// api document: https://help.aliyun.com/api/finmall/getcustomstatusinfo.html
func (client *Client) GetCustomStatusInfo(request *GetCustomStatusInfoRequest) (response *GetCustomStatusInfoResponse, err error) {
	response = CreateGetCustomStatusInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetCustomStatusInfoWithChan invokes the finmall.GetCustomStatusInfo API asynchronously
// api document: https://help.aliyun.com/api/finmall/getcustomstatusinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCustomStatusInfoWithChan(request *GetCustomStatusInfoRequest) (<-chan *GetCustomStatusInfoResponse, <-chan error) {
	responseChan := make(chan *GetCustomStatusInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCustomStatusInfo(request)
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

// GetCustomStatusInfoWithCallback invokes the finmall.GetCustomStatusInfo API asynchronously
// api document: https://help.aliyun.com/api/finmall/getcustomstatusinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCustomStatusInfoWithCallback(request *GetCustomStatusInfoRequest, callback func(response *GetCustomStatusInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCustomStatusInfoResponse
		var err error
		defer close(result)
		response, err = client.GetCustomStatusInfo(request)
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

// GetCustomStatusInfoRequest is the request struct for api GetCustomStatusInfo
type GetCustomStatusInfoRequest struct {
	*requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

// GetCustomStatusInfoResponse is the response struct for api GetCustomStatusInfo
type GetCustomStatusInfoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetCustomStatusInfoRequest creates a request to invoke GetCustomStatusInfo API
func CreateGetCustomStatusInfoRequest() (request *GetCustomStatusInfoRequest) {
	request = &GetCustomStatusInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetCustomStatusInfo", "finmall", "openAPI")
	return
}

// CreateGetCustomStatusInfoResponse creates a response to parse from GetCustomStatusInfo response
func CreateGetCustomStatusInfoResponse() (response *GetCustomStatusInfoResponse) {
	response = &GetCustomStatusInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}