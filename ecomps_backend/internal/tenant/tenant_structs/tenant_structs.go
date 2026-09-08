package tenantstructs

import (
	"time"
)

type Tenant struct {
	TenantId          uint      `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantCreation    time.Time `json:"-"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
	TenantPwId        uint      `json:"-"`
}

// Checks if a user is the admin of the tenant
func (t *Tenant) IsUserAdmin(userId uint) bool {

	if t.TenantAdminUserId == userId {
		return true
	}
	return false
}
