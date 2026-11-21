package smstructs

import "time"

type SecretEntry struct {
	Key    string
	Value  string
	Expire time.Time
}
