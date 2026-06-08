package crypto

import tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"

// Creates the master key for the given tenant
func CreateMasterKey(t tenantstructs.Tenant) (string, bool)
