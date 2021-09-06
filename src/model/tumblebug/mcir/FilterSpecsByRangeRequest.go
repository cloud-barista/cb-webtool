package mcir

//
type FilterSpecsByRangeRequest struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ConnectionName string `json:"connectionName"`
	CspSpecName    string `json:"cspSpecName"`
	OsType         string `json:"os_type"`

	CostPerHour Range `json:"cost_per_hour"`
	EbsBwMbps   Range `json:"ebs_bw_Mbps"`

	EvaluationScore01 Range  `json:"evaluationScore_01"`
	EvaluationScore02 Range  `json:"evaluationScore_02"`
	EvaluationScore03 Range  `json:"evaluationScore_03"`
	EvaluationScore04 Range  `json:"evaluationScore_04"`
	EvaluationScore05 Range  `json:"evaluationScore_05"`
	EvaluationScore06 Range  `json:"evaluationScore_06"`
	EvaluationScore07 Range  `json:"evaluationScore_07"`
	EvaluationScore08 Range  `json:"evaluationScore_08"`
	EvaluationScore09 Range  `json:"evaluationScore_09"`
	EvaluationScore10 Range  `json:"evaluationScore_10"`
	EvaluationStatus  string `json:"evaluationStatus"`

	GpuModel  string `json:"gpu_model"`
	GpuP2p    string `json:"gpu_p2p"`
	GpumemGiB Range  `json:"gpumem_GiB"`

	MaxNumStorage      Range `json:"max_num_storage"`
	MaxTotalStorageTiB Range `json:"max_total_storage_TiB"`
	MemGiB             Range `json:"mem_GiB"`

	NetBwGbps  Range `json:"net_bw_Gbps"`
	NumCore    Range `json:"num_core"`
	NumGpu     Range `json:"num_gpu"`
	NumStorage Range `json:"num_storage"`
	NumVCPU    Range `json:"num_vCPU"`
	StorageGiB Range `json:"storage_GiB"`
}
