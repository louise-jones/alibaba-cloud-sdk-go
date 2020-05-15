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

// QueryDevicePropertyData invokes the iot.QueryDevicePropertyData API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertydata.html
func (client *Client) QueryDevicePropertyData(request *QueryDevicePropertyDataRequest) (response *QueryDevicePropertyDataResponse, err error) {
	response = CreateQueryDevicePropertyDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDevicePropertyDataWithChan invokes the iot.QueryDevicePropertyData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicePropertyDataWithChan(request *QueryDevicePropertyDataRequest) (<-chan *QueryDevicePropertyDataResponse, <-chan error) {
	responseChan := make(chan *QueryDevicePropertyDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDevicePropertyData(request)
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

// QueryDevicePropertyDataWithCallback invokes the iot.QueryDevicePropertyData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicePropertyDataWithCallback(request *QueryDevicePropertyDataRequest, callback func(response *QueryDevicePropertyDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDevicePropertyDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDevicePropertyData(request)
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

// QueryDevicePropertyDataRequest is the request struct for api QueryDevicePropertyData
type QueryDevicePropertyDataRequest struct {
	*requests.RpcRequest
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Identifier    string           `position:"Query" name:"Identifier"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	Asc           requests.Integer `position:"Query" name:"Asc"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryDevicePropertyDataResponse is the response struct for api QueryDevicePropertyData
type QueryDevicePropertyDataResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDevicePropertyData `json:"Data" xml:"Data"`
}

// CreateQueryDevicePropertyDataRequest creates a request to invoke QueryDevicePropertyData API
func CreateQueryDevicePropertyDataRequest() (request *QueryDevicePropertyDataRequest) {
	request = &QueryDevicePropertyDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDevicePropertyData", "Iot", "openAPI")
	return
}

// CreateQueryDevicePropertyDataResponse creates a response to parse from QueryDevicePropertyData response
func CreateQueryDevicePropertyDataResponse() (response *QueryDevicePropertyDataResponse) {
	response = &QueryDevicePropertyDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
