package tenantstructs

import "time"

type TenantDeletionStruct struct {
	TenantDeletionId uint      `json:"-"`
	IssuedFrom       string    `json:"IssuedFrom"`
	IssuedOn         time.Time `json:"-"`
	WhenToComplete   time.Time `json:"-"`
	Deleted          bool      `json:"-"`
	TenantId         uint      `json:"TenantId"`
}

// Checks if a tenant has its deletion date today
func (t *TenantDeletionStruct) DeletionToday() bool {
	y1, m1, d1 := t.WhenToComplete.Date()
	y2, m2, d2 := time.Now().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
