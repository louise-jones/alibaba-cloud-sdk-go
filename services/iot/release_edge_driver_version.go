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

// ReleaseEdgeDriverVersion invokes the iot.ReleaseEdgeDriverVersion API synchronously
// api document: https://help.aliyun.com/api/iot/releaseedgedriverversion.html
func (client *Client) ReleaseEdgeDriverVersion(request *ReleaseEdgeDriverVersionRequest) (response *ReleaseEdgeDriverVersionResponse, err error) {
	response = CreateReleaseEdgeDriverVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseEdgeDriverVersionWithChan invokes the iot.ReleaseEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/releaseedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseEdgeDriverVersionWithChan(request *ReleaseEdgeDriverVersionRequest) (<-chan *ReleaseEdgeDriverVersionResponse, <-chan error) {
	responseChan := make(chan *ReleaseEdgeDriverVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseEdgeDriverVersion(request)
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

// ReleaseEdgeDriverVersionWithCallback invokes the iot.ReleaseEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/releaseedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseEdgeDriverVersionWithCallback(request *ReleaseEdgeDriverVersionRequest, callback func(response *ReleaseEdgeDriverVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseEdgeDriverVersionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseEdgeDriverVersion(request)
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

// ReleaseEdgeDriverVersionRequest is the request struct for api ReleaseEdgeDriverVersion
type ReleaseEdgeDriverVersionRequest struct {
	*requests.RpcRequest
	DriverId      string `position:"Query" name:"DriverId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DriverVersion string `position:"Query" name:"DriverVersion"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// ReleaseEdgeDriverVersionResponse is the response struct for api ReleaseEdgeDriverVersion
type ReleaseEdgeDriverVersionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateReleaseEdgeDriverVersionRequest creates a request to invoke ReleaseEdgeDriverVersion API
func CreateReleaseEdgeDriverVersionRequest() (request *ReleaseEdgeDriverVersionRequest) {
	request = &ReleaseEdgeDriverVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ReleaseEdgeDriverVersion", "Iot", "openAPI")
	return
}

// CreateReleaseEdgeDriverVersionResponse creates a response to parse from ReleaseEdgeDriverVersion response
func CreateReleaseEdgeDriverVersionResponse() (response *ReleaseEdgeDriverVersionResponse) {
	response = &ReleaseEdgeDriverVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
