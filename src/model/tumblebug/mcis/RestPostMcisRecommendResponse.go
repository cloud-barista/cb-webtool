package mcis

import (
	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
)

type RestPostMcisRecommendResponse struct {
	PlacementAlgo  string                `json:"placementAlgo"`
	PlacementParam []tbcommon.TbKeyValue `json:"placementParam"`
	VmRecommend    []TbVmRecommendInfo   `json:"vmReq"`
}
