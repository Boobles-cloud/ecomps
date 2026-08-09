package userstructs

type UserDeletion struct {
	UserId                   uint `json:"UserId"`
	UserWantsTenantDeletion  bool `json:"UserWantsTenantDeletion"`
	UserWnatsTenantTransfare bool `json:"UserWantsTenantTransfare"`
	NewTenantAdminUser       uint `json:"NewTenantAdminUser"`
}
