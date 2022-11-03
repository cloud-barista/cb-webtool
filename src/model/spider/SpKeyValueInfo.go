package spider

// Spider 에서 KeyValue 를 인자로 갖는 응답에서 사용
type SpKeyValueInfo struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type SpKeyValueList []SpKeyValueInfo
