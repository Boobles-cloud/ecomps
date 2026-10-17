package permissionstructs

type Permission struct {
	PermissionId          uint   `json:"PermissionId"`
	PermissionName        string `json:"PermissionName"`
	PermissionDescription string `json:"PermissionDescription"`
	UserId                uint   `json:"UserId"`
}
