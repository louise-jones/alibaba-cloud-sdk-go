package rds

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

func (client *Client) DescribeOptimizeAdviceOnMissIndex(request *DescribeOptimizeAdviceOnMissIndexRequest) (response *DescribeOptimizeAdviceOnMissIndexResponse, err error) {
	response = CreateDescribeOptimizeAdviceOnMissIndexResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeOptimizeAdviceOnMissIndexWithChan(request *DescribeOptimizeAdviceOnMissIndexRequest) (<-chan *DescribeOptimizeAdviceOnMissIndexResponse, <-chan error) {
	responseChan := make(chan *DescribeOptimizeAdviceOnMissIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOptimizeAdviceOnMissIndex(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeOptimizeAdviceOnMissIndexWithCallback(request *DescribeOptimizeAdviceOnMissIndexRequest, callback func(response *DescribeOptimizeAdviceOnMissIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOptimizeAdviceOnMissIndexResponse
		var err error
		defer close(result)
		response, err = client.DescribeOptimizeAdviceOnMissIndex(request)
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

type DescribeOptimizeAdviceOnMissIndexRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeOptimizeAdviceOnMissIndexResponse struct {
	*responses.BaseResponse
	RequestId         string           `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string           `json:"DBInstanceId" xml:"DBInstanceId"`
	TotalRecordsCount requests.Integer `json:"TotalRecordsCount" xml:"TotalRecordsCount"`
	PageNumber        requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount   requests.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
	Items             struct {
		AdviceOnMissIndex []struct {
			DBName      string `json:"DBName" xml:"DBName"`
			TableName   string `json:"TableName" xml:"TableName"`
			QueryColumn string `json:"QueryColumn" xml:"QueryColumn"`
			SQLText     string `json:"SQLText" xml:"SQLText"`
		} `json:"AdviceOnMissIndex" xml:"AdviceOnMissIndex"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeOptimizeAdviceOnMissIndexRequest() (request *DescribeOptimizeAdviceOnMissIndexRequest) {
	request = &DescribeOptimizeAdviceOnMissIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissIndex", "", "")
	return
}

func CreateDescribeOptimizeAdviceOnMissIndexResponse() (response *DescribeOptimizeAdviceOnMissIndexResponse) {
	response = &DescribeOptimizeAdviceOnMissIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}