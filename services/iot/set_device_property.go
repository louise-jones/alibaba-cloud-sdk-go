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

// SetDeviceProperty invokes the iot.SetDeviceProperty API synchronously
// api document: https://help.aliyun.com/api/iot/setdeviceproperty.html
func (client *Client) SetDeviceProperty(request *SetDevicePropertyRequest) (response *SetDevicePropertyResponse, err error) {
	response = CreateSetDevicePropertyResponse()
	err = client.DoAction(request, response)
	return
}

// SetDevicePropertyWithChan invokes the iot.SetDeviceProperty API asynchronously
// api document: https://help.aliyun.com/api/iot/setdeviceproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDevicePropertyWithChan(request *SetDevicePropertyRequest) (<-chan *SetDevicePropertyResponse, <-chan error) {
	responseChan := make(chan *SetDevicePropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDeviceProperty(request)
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

// SetDevicePropertyWithCallback invokes the iot.SetDeviceProperty API asynchronously
// api document: https://help.aliyun.com/api/iot/setdeviceproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDevicePropertyWithCallback(request *SetDevicePropertyRequest, callback func(response *SetDevicePropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDevicePropertyResponse
		var err error
		defer close(result)
		response, err = client.SetDeviceProperty(request)
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

// SetDevicePropertyRequest is the request struct for api SetDeviceProperty
type SetDevicePropertyRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	Items         string `position:"Query" name:"Items"`
}

// SetDevicePropertyResponse is the response struct for api SetDeviceProperty
type SetDevicePropertyResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateSetDevicePropertyRequest creates a request to invoke SetDeviceProperty API
func CreateSetDevicePropertyRequest() (request *SetDevicePropertyRequest) {
	request = &SetDevicePropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SetDeviceProperty", "Iot", "openAPI")
	return
}

// CreateSetDevicePropertyResponse creates a response to parse from SetDeviceProperty response
func CreateSetDevicePropertyResponse() (response *SetDevicePropertyResponse) {
	response = &SetDevicePropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
