package tenantoauthstructs

import (
	"time"
	"uuid"
)

// This is just used to display all apps
// that we support
type OAuthApps struct {
	AppId   uint   `json:"AppId"`
	AppName string `json:"AppName"`
	AppLink string `json:"AppLink"`
}

// All settings
type OAuthSettings struct {
	SettingId          uint   `json:"SettingId"`
	SettingName        string `json:"SettingName"`
	SettingDescription string `json:"SettingDescription"`
	LangId             uint   `json:"LanguageId"`
}

// This holds all token data
type TenantOAuthToken struct {
	TokenId         uint
	TokenVal        string
	RefreshTokenVal string
	LastChanged     time.Time
	AppId           uint
	TenantId        uuid.UUID
}

// This holds all information for the specific tenant
type TenantOAuthApps struct {
	TenantAppId   uint          `json:"TenantAppId"`
	SyncInterval  time.Duration `json:"Interval"`
	SyncActivated time.Time     `json:"SyncActivated"`

	// TODO: move this to another struct
	// SyncSuccessfull bool          `json:"SyncSuccess"`
	// SyncErrorMsg    string        `json:"SyncErrorMsg,omitempty"`
	UserId   uint      `json:"UserId"`
	AppId    uint      `json:"AppId"`
	TenantId uuid.UUID `json:"TenantId"`
}

// This holds all settings of that tenant
type TenantSyncSettings struct {
	TenantSettingId uint      `json:"TenantSettingId"`
	LastChanged     time.Time `json:"LastChanged"`
	SettingId       uint      `json:"SettingId"`
	AppId           uint      `json:"AppId"`
	UserId          uint      `json:"UserId"`
	TenantId        uuid.UUID `json:"TenantId"`
}
