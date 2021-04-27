package ladybug

// ladybug
// Node의 array이나 parameter 추가된 것(kind)이 있음.
type NodeList struct {
	Items []Node `json:"items"`
	Kind  string `json:"kind"`
}
