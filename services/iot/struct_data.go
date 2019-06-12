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

// Data is a nested struct in iot response
type Data struct {
	DataFormat            int                              `json:"DataFormat" xml:"DataFormat"`
	RequestProtocol       string                           `json:"RequestProtocol" xml:"RequestProtocol"`
	UtcCreate             string                           `json:"UtcCreate" xml:"UtcCreate"`
	DeviceActive          int                              `json:"DeviceActive" xml:"DeviceActive"`
	RequestMethod         string                           `json:"RequestMethod" xml:"RequestMethod"`
	Nickname              string                           `json:"Nickname" xml:"Nickname"`
	CategoryName          string                           `json:"CategoryName" xml:"CategoryName"`
	PageNo                int                              `json:"PageNo" xml:"PageNo"`
	DevEui                string                           `json:"DevEui" xml:"DevEui"`
	GroupId               string                           `json:"GroupId" xml:"GroupId"`
	OnlineCount           int64                            `json:"onlineCount" xml:"onlineCount"`
	FileId                string                           `json:"FileId" xml:"FileId"`
	LastUpdateTime        int64                            `json:"LastUpdateTime" xml:"LastUpdateTime"`
	Versions              string                           `json:"Versions" xml:"Versions"`
	AliyunCommodityCode   string                           `json:"AliyunCommodityCode" xml:"AliyunCommodityCode"`
	UtcCreatedOn          string                           `json:"UtcCreatedOn" xml:"UtcCreatedOn"`
	ApplyId               int64                            `json:"ApplyId" xml:"ApplyId"`
	GmtCreate             int64                            `json:"GmtCreate" xml:"GmtCreate"`
	MessageId             string                           `json:"MessageId" xml:"MessageId"`
	DeviceName            string                           `json:"DeviceName" xml:"DeviceName"`
	Size                  string                           `json:"Size" xml:"Size"`
	Id2                   bool                             `json:"Id2" xml:"Id2"`
	Owner                 bool                             `json:"Owner" xml:"Owner"`
	NodeType              int                              `json:"NodeType" xml:"NodeType"`
	ApiSrn                string                           `json:"ApiSrn" xml:"ApiSrn"`
	Name                  string                           `json:"Name" xml:"Name"`
	ProductName           string                           `json:"ProductName" xml:"ProductName"`
	GroupName             string                           `json:"GroupName" xml:"GroupName"`
	DownloadUrl           string                           `json:"DownloadUrl" xml:"DownloadUrl"`
	CategoryKey           string                           `json:"CategoryKey" xml:"CategoryKey"`
	CreateTime            int64                            `json:"CreateTime" xml:"CreateTime"`
	PageSize              int                              `json:"PageSize" xml:"PageSize"`
	DeviceCount           int64                            `json:"deviceCount" xml:"deviceCount"`
	ActiveCount           int64                            `json:"activeCount" xml:"activeCount"`
	DateFormat            string                           `json:"DateFormat" xml:"DateFormat"`
	Description           string                           `json:"Description" xml:"Description"`
	ApiPath               string                           `json:"ApiPath" xml:"ApiPath"`
	DeviceOnline          int                              `json:"DeviceOnline" xml:"DeviceOnline"`
	Status                int                              `json:"Status" xml:"Status"`
	ProductSecret         string                           `json:"ProductSecret" xml:"ProductSecret"`
	DeviceSecret          string                           `json:"DeviceSecret" xml:"DeviceSecret"`
	Result                string                           `json:"Result" xml:"Result"`
	ProductKey            string                           `json:"ProductKey" xml:"ProductKey"`
	DisplayName           string                           `json:"DisplayName" xml:"DisplayName"`
	JoinEui               string                           `json:"JoinEui" xml:"JoinEui"`
	IotId                 string                           `json:"IotId" xml:"IotId"`
	GroupDesc             string                           `json:"GroupDesc" xml:"GroupDesc"`
	DeviceCount           int                              `json:"DeviceCount" xml:"DeviceCount"`
	ProtocolType          string                           `json:"ProtocolType" xml:"ProtocolType"`
	ProductStatus         string                           `json:"ProductStatus" xml:"ProductStatus"`
	NetType               int                              `json:"NetType" xml:"NetType"`
	FieldNameList         FieldNameList                    `json:"FieldNameList" xml:"FieldNameList"`
	InvalidDeviceNameList InvalidDeviceNameList            `json:"InvalidDeviceNameList" xml:"InvalidDeviceNameList"`
	ResultList            ResultList                       `json:"ResultList" xml:"ResultList"`
	SqlTemplateDTO        SqlTemplateDTO                   `json:"SqlTemplateDTO" xml:"SqlTemplateDTO"`
	List                  ListInQueryDeviceDesiredProperty `json:"List" xml:"List"`
}