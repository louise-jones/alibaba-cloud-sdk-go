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

// EnableThing invokes the iot.EnableThing API synchronously
// api document: https://help.aliyun.com/api/iot/enablething.html
func (client *Client) EnableThing(request *EnableThingRequest) (response *EnableThingResponse, err error) {
	response = CreateEnableThingResponse()
	err = client.DoAction(request, response)
	return
}

// EnableThingWithChan invokes the iot.EnableThing API asynchronously
// api document: https://help.aliyun.com/api/iot/enablething.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableThingWithChan(request *EnableThingRequest) (<-chan *EnableThingResponse, <-chan error) {
	responseChan := make(chan *EnableThingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableThing(request)
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

// EnableThingWithCallback invokes the iot.EnableThing API asynchronously
// api document: https://help.aliyun.com/api/iot/enablething.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableThingWithCallback(request *EnableThingRequest, callback func(response *EnableThingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableThingResponse
		var err error
		defer close(result)
		response, err = client.EnableThing(request)
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

// EnableThingRequest is the request struct for api EnableThing
type EnableThingRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// EnableThingResponse is the response struct for api EnableThing
type EnableThingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateEnableThingRequest creates a request to invoke EnableThing API
func CreateEnableThingRequest() (request *EnableThingRequest) {
	request = &EnableThingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "EnableThing", "Iot", "openAPI")
	return
}

// CreateEnableThingResponse creates a response to parse from EnableThing response
func CreateEnableThingResponse() (response *EnableThingResponse) {
	response = &EnableThingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
