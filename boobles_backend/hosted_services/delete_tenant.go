package hostedservices

import (
	"context"
	"time"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Checks if a tenant needs to be deleted -> then deletes it and all its data
// Runs every 12h
// TODO: Add all other tenant data to this
func DeleteTenants(ctx context.Context, dh *database.DbHandler) {

	ticker := time.NewTicker(12 * time.Hour)

	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			tenantDeletions, ok := database.QueryMany[tenantstructs.TenantDeletionStruct](ctx, dh, "SelectAllTenantDeletion", []any{})

			if !ok || len(tenantDeletions) == 0 {
				logging.Log(logging.Error, "[Hosted Service | DeleteTenants] Failed to get all tenant deletion")
				ticker.Reset(12 * time.Hour)
			}

			for i := range tenantDeletions {

				if tenantDeletions[i].IsDeletionToday() {
					dh.ExecuteSQLStatement("DeleteTenantById", []any{tenantDeletions[i].TenantId})
					dh.ExecuteSQLStatement("DeleteOrderByTenantId", []any{tenantDeletions[i].TenantId})
					dh.ExecuteSQLStatement("DeleteCustomerByTenantId", []any{tenantDeletions[i].TenantId})
				}
			}

			ticker.Reset(12 * time.Hour)
		}
	}
}
