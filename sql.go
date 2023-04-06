package optional

import (
	"database/sql"
	"database/sql/driver"
)

type scannerValuer interface {
	sql.Scanner
	driver.Valuer
}

type SQL[T scannerValuer] struct {
	Optional[T]
}

func NoneSQL[T scannerValuer]() SQL[T] {
	return SQL[T]{Optional: None[T]()}
}

func SomeSQL[T scannerValuer](v T) SQL[T] {
	return SQL[T]{Optional: Some(v)}
}

func FromPtrSQL[T scannerValuer](ptr *T) SQL[T] {
	return SQL[T]{Optional: FromPtr(ptr)}
}

func (o *SQL[T]) Scan(value any) (err error) {
	if value == nil {
		*o = NoneSQL[T]()
		return
	}

	if err := o.value.Scan(value); err != nil {
		*o = NoneSQL[T]()
		return err
	}

	o.valid = true
	return nil
}

func (o SQL[T]) Value() (driver.Value, error) {
	if !o.Is() {
		return nil, nil
	}
	return o.value.Value()
}
