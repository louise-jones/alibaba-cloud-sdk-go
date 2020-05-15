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

// DeleteProduct invokes the iot.DeleteProduct API synchronously
// api document: https://help.aliyun.com/api/iot/deleteproduct.html
func (client *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
	response = CreateDeleteProductResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProductWithChan invokes the iot.DeleteProduct API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProductWithChan(request *DeleteProductRequest) (<-chan *DeleteProductResponse, <-chan error) {
	responseChan := make(chan *DeleteProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProduct(request)
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

// DeleteProductWithCallback invokes the iot.DeleteProduct API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProductWithCallback(request *DeleteProductRequest, callback func(response *DeleteProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProductResponse
		var err error
		defer close(result)
		response, err = client.DeleteProduct(request)
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

// DeleteProductRequest is the request struct for api DeleteProduct
type DeleteProductRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteProductResponse is the response struct for api DeleteProduct
type DeleteProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteProductRequest creates a request to invoke DeleteProduct API
func CreateDeleteProductRequest() (request *DeleteProductRequest) {
	request = &DeleteProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteProduct", "Iot", "openAPI")
	return
}

// CreateDeleteProductResponse creates a response to parse from DeleteProduct response
func CreateDeleteProductResponse() (response *DeleteProductResponse) {
	response = &DeleteProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
