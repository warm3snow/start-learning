package ac

type ResoucePolicy []ResoucePolicyElement

type ResoucePolicyElement struct {
	Policy       Policy `json:"policy"`
	ResourceName string `json:"resource_name"`
}

type Policy struct {
	RoleList []RoleList `json:"role_list"`
	Rule     string     `json:"rule"`
	OrgList  []string   `json:"org_list"`
}

type RoleList string

const (
	Admin     RoleList = "ADMIN"
	Client    RoleList = "CLIENT"
	Common    RoleList = "COMMON"
	Consensus RoleList = "CONSENSUS"
	Light     RoleList = "LIGHT"
)
