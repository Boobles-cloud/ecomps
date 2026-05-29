package database

type DBUpdateStruct[T any] struct {
	StatementName    string
	StatementContext string
	StatementType    int
	StatementStruct  T
}

// TODO implement meee
func (d *DBUpdateStruct[T]) Execute() bool {
	return false
}
