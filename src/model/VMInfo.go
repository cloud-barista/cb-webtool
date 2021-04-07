package model

type VMInfo struct {
	ID  string `json:"id"`
	Name  string `json:"name"`
	ConnectionName  string `json:"connectionName"`

	
	SpecId  string `json:"specId"`
	ImageID  string `json:"imageId"`

	VNetId  string `json:"vNetId"`
	SubnetId  string `json:"subnetId"`
	SecurityGroupIDs  []string `json:"securityGroupIds"`
	
	SshKeyId  string `json:"sshKeyId"`
	
	VmUserAccount  string `json:"vmUserAccount"`
	VmUserPassword  string `json:"vmUserPassword"`

	Description  string `json:"description"`
	Label  string `json:"label"`

	Location LocationInfo `json:"location"`
	RegionZone NameSystemIdInfo `json:"region"`

	publicIP  string `json:"publicIP"`
    publicDNS  string `json:"publicDNS"`
	privateIP  string `json:"privateIP"`
	privateDNS  string `json:"privateDNS"`

	
	
	VmBootDisk  string `json:"vmBootDisk"`
	VmBlockDisk  string `json:"vmBlockDisk"`
	
	Status  string `json:"status"`
	TargetStatus  string `json:"targetStatus"`
	TargetAction  string `json:"targetAction"`
	
	MonAgentStatus  string `json:"monAgentStatus"` // "monAgentStatus": "[installed, notInstalled, failed]",
	
	CspViewVmDetail	CspViewVmDetailInfo `json:"cspViewVmDetail"`
}


type CspViewVmDetailInfo struct {
	Name	string `json:"name"`
	ImageName  string `json:"imageName"`
	Vpcname  string `json:"vpcname"`
	SubnetName  string `json:"subnetName"`
	SecurityGroupNames  []string `json:"securityGroupNames"`

	KeyPairName  string `json:"keyPairName"`

	VmspecName  string `json:"vmspecName"`

	VmuserId  string `json:"vmuserId"`
	VmuserPasswd  string `json:"vmuserPasswd"`	

	ConnectionName  string `json:"connectionName"`
	IID			NameSystemIdInfo `json:"iid"`
	ImageIID	NameSystemIdInfo `json:"imageIId"`
	VpcIID  NameSystemIdInfo `json:"vpcIID"`
	SubnetIID  NameSystemIdInfo `json:"subnetIID"`
	SecurityGroupIIDs []NameSystemIdInfo `json:"securityGroupIIds"`
	keyPairIID	NameSystemIdInfo `json:"keyPairIId"`

	StartTime  string `json:"startTime"`
	RegionZone NameSystemIdInfo `json:"region"`

	NetworkInterface	string `json:"networkInterface"`

	publicIP	string `json:"publicIP"`
	publicDNS	string `json:"publicDNS"`
	privateIP	string `json:"privateIP"`
	privateDNS	string `json:"privateDNS"`
	
	VmbootDisk  string `json:"vmbootDisk"`
	VmblockDisk  string `json:"vmblockDisk"`
	
	KeyValueInfos []KeyValueInfoList `json:"keyValueList"`
	
}

type LocationInfo struct{
	BriefAddr  string `json:"briefAddr"`
	CloudType  string `json:"cloudType"`
	Latitude  string `json:"latitude"`
	Longitude  string `json:"longitude"`
	NativeRegion  string `json:"nativeRegion"`
}
// 
type NameSystemIdInfo struct{
	NameID  string `json:"nameId"`
	SystemID  string `json:"systemId"`	
}
// type IIDInfo struct {
// 	NameID  string `json:"nameId"`
// 	SystemID  string `json:"systemId"`	
// }
// type ImageIIDInfo struct {
// 	NameID  string `json:"nameId"`
// 	SystemID  string `json:"systemId"`
// }
// type keyPairIIDInfo struct {
// 	NameID  string `json:"nameId"`
// 	SystemID  string `json:"systemId"`
// }

// // 원래는 RegionInfo가 맞으나, Resouce에서 RegionInfo를 사용하므로 RegionZoneInfo로 명명
// type RegionZoneInfo struct {
// 	Region  string `json:"region"`
// 	Zone  string `json:"zone"`
// }



// type SecurityGroupIIdsInfo struct {
// 	NameID  string `json:"nameId"`
// 	SystemID  string `json:"systemId"`
// }