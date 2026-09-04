package userstructs

type UserStruct struct {
	UserId        uint   `json:"UserId"`
	UserName      string `json:"UserName"`
	UserPW        string `json:"UserPw"`
	UserMail      string `json:"UserMail"`
	UserTel       string `json:"UserTel"`
	UserHas2FA    bool   `json:"UserHas2Fa"`
	UserHasTenant bool   `json:"UserHasTenant"`
	TenantId      uint   `json:"TenantId"`
}
