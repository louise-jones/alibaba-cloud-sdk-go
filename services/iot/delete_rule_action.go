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

// DeleteRuleAction invokes the iot.DeleteRuleAction API synchronously
// api document: https://help.aliyun.com/api/iot/deleteruleaction.html
func (client *Client) DeleteRuleAction(request *DeleteRuleActionRequest) (response *DeleteRuleActionResponse, err error) {
	response = CreateDeleteRuleActionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRuleActionWithChan invokes the iot.DeleteRuleAction API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteruleaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRuleActionWithChan(request *DeleteRuleActionRequest) (<-chan *DeleteRuleActionResponse, <-chan error) {
	responseChan := make(chan *DeleteRuleActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRuleAction(request)
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

// DeleteRuleActionWithCallback invokes the iot.DeleteRuleAction API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteruleaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRuleActionWithCallback(request *DeleteRuleActionRequest, callback func(response *DeleteRuleActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRuleActionResponse
		var err error
		defer close(result)
		response, err = client.DeleteRuleAction(request)
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

// DeleteRuleActionRequest is the request struct for api DeleteRuleAction
type DeleteRuleActionRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ActionId      requests.Integer `position:"Query" name:"ActionId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// DeleteRuleActionResponse is the response struct for api DeleteRuleAction
type DeleteRuleActionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteRuleActionRequest creates a request to invoke DeleteRuleAction API
func CreateDeleteRuleActionRequest() (request *DeleteRuleActionRequest) {
	request = &DeleteRuleActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteRuleAction", "Iot", "openAPI")
	return
}

// CreateDeleteRuleActionResponse creates a response to parse from DeleteRuleAction response
func CreateDeleteRuleActionResponse() (response *DeleteRuleActionResponse) {
	response = &DeleteRuleActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
