package elasticsearch

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

// Result is a nested struct in elasticsearch response
type Result struct {
	Config                    map[string]interface{}     `json:"config" xml:"config"`
	KibanaPort                int                        `json:"kibanaPort" xml:"kibanaPort"`
	VpcInstanceId             string                     `json:"vpcInstanceId" xml:"vpcInstanceId"`
	UserName                  string                     `json:"userName" xml:"userName"`
	Endpoints                 string                     `json:"endpoints" xml:"endpoints"`
	PublicDomain              string                     `json:"publicDomain" xml:"publicDomain"`
	HaveKibana                bool                       `json:"haveKibana" xml:"haveKibana"`
	KibanaDomain              string                     `json:"kibanaDomain" xml:"kibanaDomain"`
	Domain                    string                     `json:"domain" xml:"domain"`
	RegionId                  string                     `json:"regionId" xml:"regionId"`
	Port                      int                        `json:"port" xml:"port"`
	Env                       string                     `json:"env" xml:"env"`
	EnableKibanaPublicNetwork bool                       `json:"enableKibanaPublicNetwork" xml:"enableKibanaPublicNetwork"`
	WarmNode                  bool                       `json:"warmNode" xml:"warmNode"`
	AdvancedDedicateMaster    bool                       `json:"advancedDedicateMaster" xml:"advancedDedicateMaster"`
	EnablePublic              bool                       `json:"enablePublic" xml:"enablePublic"`
	EsVersion                 string                     `json:"esVersion" xml:"esVersion"`
	Id                        string                     `json:"id" xml:"id"`
	PublicPort                int                        `json:"publicPort" xml:"publicPort"`
	DedicateMaster            bool                       `json:"dedicateMaster" xml:"dedicateMaster"`
	Enable                    bool                       `json:"enable" xml:"enable"`
	QuartzRegex               string                     `json:"quartzRegex" xml:"quartzRegex"`
	CreateUrl                 string                     `json:"createUrl" xml:"createUrl"`
	PipelineManagementType    string                     `json:"pipelineManagementType" xml:"pipelineManagementType"`
	Protocol                  string                     `json:"protocol" xml:"protocol"`
	PaymentType               string                     `json:"paymentType" xml:"paymentType"`
	UpdatedAt                 string                     `json:"updatedAt" xml:"updatedAt"`
	NodeAmount                int                        `json:"nodeAmount" xml:"nodeAmount"`
	InstanceId                string                     `json:"instanceId" xml:"instanceId"`
	Status                    string                     `json:"status" xml:"status"`
	ZoneCount                 int                        `json:"zoneCount" xml:"zoneCount"`
	Description               string                     `json:"description" xml:"description"`
	EsConfig                  map[string]interface{}     `json:"esConfig" xml:"esConfig"`
	Version                   string                     `json:"version" xml:"version"`
	HaveClientNode            bool                       `json:"haveClientNode" xml:"haveClientNode"`
	CreatedAt                 string                     `json:"createdAt" xml:"createdAt"`
	EsIPWhitelist             []string                   `json:"esIPWhitelist" xml:"esIPWhitelist"`
	EsIPBlacklist             []string                   `json:"esIPBlacklist" xml:"esIPBlacklist"`
	PrivateNetworkIpWhiteList []string                   `json:"privateNetworkIpWhiteList" xml:"privateNetworkIpWhiteList"`
	MasterSpec                []string                   `json:"masterSpec" xml:"masterSpec"`
	PublicIpWhitelist         []string                   `json:"publicIpWhitelist" xml:"publicIpWhitelist"`
	ClientNodeSpec            []string                   `json:"clientNodeSpec" xml:"clientNodeSpec"`
	EsVersions                []string                   `json:"esVersions" xml:"esVersions"`
	KibanaIPWhitelist         []string                   `json:"kibanaIPWhitelist" xml:"kibanaIPWhitelist"`
	PipelineIds               []string                   `json:"pipelineIds" xml:"pipelineIds"`
	Zones                     []string                   `json:"zones" xml:"zones"`
	Node                      Node                       `json:"node" xml:"node"`
	NodeSpec                  NodeSpec                   `json:"nodeSpec" xml:"nodeSpec"`
	AdvancedSetting           AdvancedSetting            `json:"advancedSetting" xml:"advancedSetting"`
	JvmConfine                JvmConfine                 `json:"jvmConfine" xml:"jvmConfine"`
	KibanaConfiguration       KibanaConfiguration        `json:"kibanaConfiguration" xml:"kibanaConfiguration"`
	WarmNodeConfiguration     WarmNodeConfiguration      `json:"warmNodeConfiguration" xml:"warmNodeConfiguration"`
	ClientNodeConfiguration   ClientNodeConfiguration    `json:"clientNodeConfiguration" xml:"clientNodeConfiguration"`
	WarmNodeProperties        WarmNodeProperties         `json:"warmNodeProperties" xml:"warmNodeProperties"`
	NetworkConfig             NetworkConfig              `json:"networkConfig" xml:"networkConfig"`
	MasterConfiguration       MasterConfiguration        `json:"masterConfiguration" xml:"masterConfiguration"`
	KibanaNodeProperties      KibanaNodeProperties       `json:"kibanaNodeProperties" xml:"kibanaNodeProperties"`
	ClientNodeAmountRange     ClientNodeAmountRange      `json:"clientNodeAmountRange" xml:"clientNodeAmountRange"`
	SupportVersions           []CategoryEntity           `json:"supportVersions" xml:"supportVersions"`
	ZoneInfos                 []ZoneInfo                 `json:"zoneInfos" xml:"zoneInfos"`
	MasterDiskList            []Disk                     `json:"masterDiskList" xml:"masterDiskList"`
	EsVersionsLatestList      []EsVersionsLatestListItem `json:"esVersionsLatestList" xml:"esVersionsLatestList"`
	EndpointList              []Endpoint                 `json:"endpointList" xml:"endpointList"`
	NodeSpecList              []NodeSpecListItem         `json:"nodeSpecList" xml:"nodeSpecList"`
	ClientNodeDiskList        []Disk                     `json:"clientNodeDiskList" xml:"clientNodeDiskList"`
	DataDiskList              []DataDiskListItem         `json:"dataDiskList" xml:"dataDiskList"`
	SynonymsDicts             []SynonymsDicts            `json:"synonymsDicts" xml:"synonymsDicts"`
	AliwsDicts                []Dict                     `json:"aliwsDicts" xml:"aliwsDicts"`
	DictList                  []DictList                 `json:"dictList" xml:"dictList"`
}