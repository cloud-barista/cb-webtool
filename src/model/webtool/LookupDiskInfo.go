package webtool

type LookupDiskInfo struct {
	// create disk list
	Provider string `json:"provider"`

	RootDiskType []string `json:"rootdisktype"`

	//
	DataDiskType []string `json:"datadisktype"`

	// disk size range by diskType
	DiskSize []string `json:"disksize"`
}
