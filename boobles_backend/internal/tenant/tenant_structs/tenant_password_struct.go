package tenantstructs

type TenantPwStruct struct {
	TenantPwId  uint
	TenantPwVal string
}

func (t *TenantPwStruct) CreateTenantPwInDatabase() bool {
	// TODO
	return false
}
