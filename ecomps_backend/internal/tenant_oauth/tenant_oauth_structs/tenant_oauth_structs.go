package tenantoauthstructs

import "time"

// This is just used to display all apps
// that we support
type OAuthApps struct {
	AppId   uint   `json:"AppId"`
	AppName string `json:"AppName"`
	AppLink string `json:"AppLink"`
}

// This holds all token data
type TenantOAuthToken struct {
	TokenId         uint
	TokenVal        string
	RefreshTokenVal string
	LastChanged     time.Time
	AppId           uint
	TenantId        uint
}

// This holds all information for the specific tenant
type TenantOAuthApps struct {
	TenantAppId     uint          `json:"TenantAppId"`
	SyncInterval    time.Duration `json:"Interval"`
	SyncActivated   time.Time     `json:"SyncActivated"`
	SyncSuccessfull bool          `json:"SyncSuccess"`
	SyncErrorMsg    string        `json:"SyncErrorMsg,omitempty"`
	UserId          uint          `json:"UserId"`
	AppId           uint          `json:"AppId"`
	TenantId        uint          `json:"TenantId"`
}

// This holds all settings of that tenant
type TenantSyncSettings struct {
	TenantSettingId uint      `json:"TenantSettingId"`
	LastChanged     time.Time `json:"LastChanged"`
	SettingId       uint      `json:"SettingId"`
	UserId          uint      `json:"UserId"`
	TenantId        uint      `json:"TenantId"`
}

// All settings
type OAuthSettings struct {
	SettingId          uint   `json:"SettingId"`
	SettingName        string `json:"SettingName"`
	SettingDescription string `json:"SettingDescription"`
	LangId             uint   `json:"LanguageId"`
}
