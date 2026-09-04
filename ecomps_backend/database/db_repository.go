package database

import (
	"context"
	"database/sql"
	"errors"
)

// This gets used by our services
// IMPORTANT: dont forgett to implement "ToArgsFunc"
type DatabaseRepository[T any] struct {
	db         *DbHandler
	entityName string
	toArgs     ToArgsFunc[T]
}

type ToArgsFunc[T any] func(item T) []any

func NewDatabaseRepository[T any](db *DbHandler, entityName string, toArgs ToArgsFunc[T]) *DatabaseRepository[T] {
	return &DatabaseRepository[T]{
		db:         db,
		entityName: entityName,
		toArgs:     toArgs,
	}
}

func (d *DatabaseRepository[T]) GetById(ctx context.Context, id uint) (T, error) {

	var item T

	item, ok := QueryOne[T](ctx, d.db, "Select"+d.entityName+"ById", id)

	if !ok {
		return item, sql.ErrNoRows
	}

	return item, nil
}

func (d *DatabaseRepository[T]) GetAllByTenantId(ctx context.Context, tenantId uint) ([]T, error) {

	var items []T

	items, ok := QueryMany[T](ctx, d.db, "Select"+d.entityName+"ByTenantId", tenantId)

	if !ok {
		return items, sql.ErrNoRows
	}

	return items, nil
}

func (d *DatabaseRepository[T]) Create(ctx context.Context, item T) error {

	args := d.toArgs(item)

	if result := d.db.ExecuteSQLStatement(ctx, "Insert"+d.entityName, args); !result.Ok {
		return errors.New("Failed to create item")
	}

	return nil
}

func (d *DatabaseRepository[T]) Update(ctx context.Context, item T, filterName string) error {

	if ok := UpdateDatabaseEntry[T](d.db, "Update"+d.entityName, filterName, item); !ok {
		return errors.New("Failed to update item")
	}

	return nil
}

func (d *DatabaseRepository[T]) Delete(ctx context.Context, id uint) error {

	if result := d.db.ExecuteSQLStatement(ctx, "Delete"+d.entityName, []any{id}); !result.Ok {
		return errors.New("Failed to delete item")
	}

	return nil
}
