package ccc

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

func (client *Client) ListCallDetailRecords(request *ListCallDetailRecordsRequest) (response *ListCallDetailRecordsResponse, err error) {
	response = CreateListCallDetailRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListCallDetailRecordsWithChan(request *ListCallDetailRecordsRequest) (<-chan *ListCallDetailRecordsResponse, <-chan error) {
	responseChan := make(chan *ListCallDetailRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCallDetailRecords(request)
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

func (client *Client) ListCallDetailRecordsWithCallback(request *ListCallDetailRecordsRequest, callback func(response *ListCallDetailRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCallDetailRecordsResponse
		var err error
		defer close(result)
		response, err = client.ListCallDetailRecords(request)
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

type ListCallDetailRecordsRequest struct {
	*requests.RpcRequest
	InstanceId         string           `position:"Query" name:"InstanceId"`
	PhoneNumber        string           `position:"Query" name:"PhoneNumber"`
	StartTime          requests.Integer `position:"Query" name:"StartTime"`
	StopTime           requests.Integer `position:"Query" name:"StopTime"`
	Criteria           string           `position:"Query" name:"Criteria"`
	ContactType        string           `position:"Query" name:"ContactType"`
	ContactDisposition string           `position:"Query" name:"ContactDisposition"`
	WithRecording      requests.Boolean `position:"Query" name:"WithRecording"`
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
}

type ListCallDetailRecordsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	HttpStatusCode    int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	CallDetailRecords CallDetailRecords `json:"CallDetailRecords" xml:"CallDetailRecords"`
}

func CreateListCallDetailRecordsRequest() (request *ListCallDetailRecordsRequest) {
	request = &ListCallDetailRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListCallDetailRecords", "", "")
	return
}

func CreateListCallDetailRecordsResponse() (response *ListCallDetailRecordsResponse) {
	response = &ListCallDetailRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}