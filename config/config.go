package config

type Config struct {
	Users []User `json:"users"`
}

type User struct {
	Account         string `json:"Account"`
	DeptCode        string `json:"DeptCode"`
	CompanyName     string `json:"CompanyName"`
	DeptName        string `json:"DeptName"`
	UserName        string `json:"UserName"`
	Sex             string `json:"Sex"`
	Age             string `json:"Age"`
	Gwmc            string `json:"Gwmc"`
	Phone           string `json:"Phone"`
	WorkPlace       string `json:"WorkPlace"`
	NewStatus       string `json:"NewStatus"`
	UpTime          string `json:"UpTime"`
	Province        string `json:"Province"`
	City            string `json:"City"`
	Region          string `json:"Region"`
	WzStatus        string `json:"WzStatus"`
	NowStatus       string `json:"NowStatus"`
	IsLeaveCompany  string `json:"IsLeaveCompany"`
	NowNear         string `json:"NowNear"`
	WzSProvince     string `json:"Wz_sProvince"`
	WzSCity         string `json:"Wz_sCity"`
	WzSRegion       string `json:"Wz_sRegion"`
	WzEProvince     string `json:"Wz_eProvince"`
	WzECity         string `json:"Wz_eCity"`
	WzERegion       string `json:"Wz_eRegion"`
	WzStime         string `json:"Wz_stime"`
	WzEtime         string `json:"Wz_etime"`
	WzPtime         string `json:"Wz_ptime"`
	WzTravel        string `json:"Wz_travel"`
	WzReson         string `json:"Wz_reson"`
	Status          string `json:"Status"`
	CommunityRecord string `json:"CommunityRecord"`
	GLstime         string `json:"GLstime"`
	GLetime         string `json:"GLetime"`
	GLstreet        string `json:"GLstreet"`
	GLcommunity     string `json:"GLcommunity"`
	GLaddress       string `json:"GLaddress"`
	GLremark        string `json:"GLremark"`
	GLremark2       string `json:"GLremark2"`
	FamilyStatus    string `json:"FamilyStatus"`
	IsToCompany     string `json:"IsToCompany"`
	WzRemark        string `json:"Wz_remark"`
	Temperature     string `json:"Temperature"`
	Province2       string `json:"Province2"`
	City2           string `json:"City2"`
	Region2         string `json:"Region2"`
}
