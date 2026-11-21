package tenantstructs

import (
	"time"
	"uuid"
)

type Tenant struct {
	TenantId          uuid.UUID `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantCreation    time.Time `json:"-"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
}

// Checks if a user is the admin of the tenant
func (t *Tenant) IsUserAdmin(userId uint) bool {

	if t.TenantAdminUserId == userId {
		return true
	}
	return false
}
