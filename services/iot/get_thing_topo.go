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

// GetThingTopo invokes the iot.GetThingTopo API synchronously
// api document: https://help.aliyun.com/api/iot/getthingtopo.html
func (client *Client) GetThingTopo(request *GetThingTopoRequest) (response *GetThingTopoResponse, err error) {
	response = CreateGetThingTopoResponse()
	err = client.DoAction(request, response)
	return
}

// GetThingTopoWithChan invokes the iot.GetThingTopo API asynchronously
// api document: https://help.aliyun.com/api/iot/getthingtopo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThingTopoWithChan(request *GetThingTopoRequest) (<-chan *GetThingTopoResponse, <-chan error) {
	responseChan := make(chan *GetThingTopoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThingTopo(request)
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

// GetThingTopoWithCallback invokes the iot.GetThingTopo API asynchronously
// api document: https://help.aliyun.com/api/iot/getthingtopo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThingTopoWithCallback(request *GetThingTopoRequest, callback func(response *GetThingTopoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThingTopoResponse
		var err error
		defer close(result)
		response, err = client.GetThingTopo(request)
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

// GetThingTopoRequest is the request struct for api GetThingTopo
type GetThingTopoRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	PageNo        requests.Integer `position:"Query" name:"PageNo"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// GetThingTopoResponse is the response struct for api GetThingTopo
type GetThingTopoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetThingTopoRequest creates a request to invoke GetThingTopo API
func CreateGetThingTopoRequest() (request *GetThingTopoRequest) {
	request = &GetThingTopoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetThingTopo", "Iot", "openAPI")
	return
}

// CreateGetThingTopoResponse creates a response to parse from GetThingTopo response
func CreateGetThingTopoResponse() (response *GetThingTopoResponse) {
	response = &GetThingTopoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
