package ecs

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

func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

type CreateInstanceRequest struct {
	*requests.RpcRequest
	SpotPriceLimit              requests.Float            `position:"Query" name:"SpotPriceLimit"`
	SystemDiskCategory          string                    `position:"Query" name:"SystemDisk.Category"`
	SpotStrategy                string                    `position:"Query" name:"SpotStrategy"`
	ZoneId                      string                    `position:"Query" name:"ZoneId"`
	ClusterId                   string                    `position:"Query" name:"ClusterId"`
	InnerIpAddress              string                    `position:"Query" name:"InnerIpAddress"`
	Password                    string                    `position:"Query" name:"Password"`
	PeriodUnit                  string                    `position:"Query" name:"PeriodUnit"`
	InstanceType                string                    `position:"Query" name:"InstanceType"`
	Tag3Key                     string                    `position:"Query" name:"Tag.3.Key"`
	Tag5Value                   string                    `position:"Query" name:"Tag.5.Value"`
	Description                 string                    `position:"Query" name:"Description"`
	Tag1Key                     string                    `position:"Query" name:"Tag.1.Key"`
	Tag1Value                   string                    `position:"Query" name:"Tag.1.Value"`
	ResourceOwnerId             requests.Integer          `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                string                    `position:"Query" name:"OwnerAccount"`
	AutoRenewPeriod             requests.Integer          `position:"Query" name:"AutoRenewPeriod"`
	VSwitchId                   string                    `position:"Query" name:"VSwitchId"`
	ClientToken                 string                    `position:"Query" name:"ClientToken"`
	InternetChargeType          string                    `position:"Query" name:"InternetChargeType"`
	SecurityGroupId             string                    `position:"Query" name:"SecurityGroupId"`
	HostName                    string                    `position:"Query" name:"HostName"`
	KeyPairName                 string                    `position:"Query" name:"KeyPairName"`
	OwnerId                     requests.Integer          `position:"Query" name:"OwnerId"`
	SystemDiskDiskName          string                    `position:"Query" name:"SystemDisk.DiskName"`
	Period                      requests.Integer          `position:"Query" name:"Period"`
	Tag5Key                     string                    `position:"Query" name:"Tag.5.Key"`
	IoOptimized                 string                    `position:"Query" name:"IoOptimized"`
	UseAdditionalService        requests.Boolean          `position:"Query" name:"UseAdditionalService"`
	InstanceName                string                    `position:"Query" name:"InstanceName"`
	Tag2Key                     string                    `position:"Query" name:"Tag.2.Key"`
	Tag4Key                     string                    `position:"Query" name:"Tag.4.Key"`
	SecurityEnhancementStrategy string                    `position:"Query" name:"SecurityEnhancementStrategy"`
	UserData                    string                    `position:"Query" name:"UserData"`
	ResourceOwnerAccount        string                    `position:"Query" name:"ResourceOwnerAccount"`
	NodeControllerId            string                    `position:"Query" name:"NodeControllerId"`
	AutoRenew                   requests.Boolean          `position:"Query" name:"AutoRenew"`
	Tag4Value                   string                    `position:"Query" name:"Tag.4.Value"`
	SystemDiskSize              requests.Integer          `position:"Query" name:"SystemDisk.Size"`
	DataDisk                    *[]CreateInstanceDataDisk `position:"Query" name:"DataDisk"  type:"Repeated"`
	InternetMaxBandwidthOut     requests.Integer          `position:"Query" name:"InternetMaxBandwidthOut"`
	PrivateIpAddress            string                    `position:"Query" name:"PrivateIpAddress"`
	VlanId                      string                    `position:"Query" name:"VlanId"`
	ResourceGroupId             string                    `position:"Query" name:"ResourceGroupId"`
	ImageId                     string                    `position:"Query" name:"ImageId"`
	HpcClusterId                string                    `position:"Query" name:"HpcClusterId"`
	RamRoleName                 string                    `position:"Query" name:"RamRoleName"`
	InternetMaxBandwidthIn      requests.Integer          `position:"Query" name:"InternetMaxBandwidthIn"`
	DeploymentSetId             string                    `position:"Query" name:"DeploymentSetId"`
	SystemDiskDescription       string                    `position:"Query" name:"SystemDisk.Description"`
	InstanceChargeType          string                    `position:"Query" name:"InstanceChargeType"`
	Tag3Value                   string                    `position:"Query" name:"Tag.3.Value"`
	Tag2Value                   string                    `position:"Query" name:"Tag.2.Value"`
}

type CreateInstanceDataDisk struct {
	Size               string `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
	Encrypted          string `name:"Encrypted"`
}

type CreateInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateInstance", "", "")
	return
}

func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}