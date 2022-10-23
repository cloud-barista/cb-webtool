package spider

type ClusterReqInfo struct {
	Name               string          `json:"Name"`
	Version            string          `json:"Version"`
	VPCName            string          `json:"VPCName"`
	SubnetNames        []string        `json:"SubnetNames"`
	SecurityGroupNames []string        `json:"SecurityGroupNames"`
	NodeGroupList      []NodeGroupInfo `json:"NodeGroupList"`
}

/*
"Name": "myk8scluser-01",
                        "Version": "1.20.6",

                         "VPCName": "vpc-01",
                        "SubnetNames": ["subnet-01"],
                        "SecurityGroupNames": ["sg-01"],

                        "NodeGroupList": [
                                {"Name": "Economy", "ImageName": "img-pi0ii46r", "VMSpecName": "S3.MEDIUM2", "KeyPairName": "keypair-01",
                                        "OnAutoScaling": "true", "DesiredNodeSize": "2", "MinNodeSize": "2", "MaxNodeSize": "2"}
                                ]
*/
