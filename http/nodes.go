package main

type NodeRequest struct {
	Kind string `json:"kind"` // Node kind (eg linux)
	Size string `json:"size"` // size
	// optional 	Default value : XS, The following parameters refer to the SGCLOUD fields :
	// * XS
	// * S
	// * M
	// * L
	// * L32 (L With 32 gb memory)
	// * XL32
	// * XL32 (XL With 32 gb memory)
	// * XL64 (XL With 64 gb memory).
	// You can check couples vCPU/vRAM and reversibility on top of this page
	// Value must be one of: XS, S, M, L, L32, XL, XL32, XL64.
	Disksize         int    `json:"disksize"`
	Leasedays        int    `json:"leasedays"`
	EnvironmentType  string `json:"environment_type"`
	CentrifyZone     string `json:"centrify_zone"`
	Description      string `json:"description"`
	Usergroup        string `json:"usergroup"`
	ServiceAccount   string `json:"service_account"`
	AppTrigram       string `json:"app_trigram"`
	Region           string `json:"region"`
	AvailabilityZone string `json:"availability_zone"`
	Subnet           string `json:"subnet"`
	Notifymail       string `json:"notifymail"`
}
