package userstructs

import "time"

type UserDeletion struct {
	UserId                   uint `json:"UserId"`
	UserWantsTenantDeletion  bool `json:"UserWantsTenantDeletion"`
	UserWnatsTenantTransfare bool `json:"UserWantsTenantTransfare"`
	NewTenantAdminUser       uint `json:"NewTenantAdminUser"`
}

// NOTE: this struct is only used when a user needs a deletion after a specific date
type UserDeletionDatabase struct {
	DeletionId     uint
	IssuedOn       time.Time
	WhenToComplete time.Time
	UserId         uint
}

func (t *UserDeletionDatabase) IsDeletionToday() bool {
	y1, m1, d1 := t.WhenToComplete.Date()
	y2, m2, d2 := time.Now().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
