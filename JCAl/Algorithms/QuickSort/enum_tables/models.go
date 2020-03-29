package enumtables

import (
	"database/sql"
	"github.com/ryu/hxs/config"
	"time"
)

type Enum struct {
	ID        sql.NullInt64
	UpdatedAt *time.Time
	Type      sql.NullString
	Name      sql.NullString
	Symbol    sql.NullString
	Attrs     sql.NullString
	Ordinal   sql.NullInt64
	DeletedAt *time.Time
}

func GetEnum(id sql.NullInt64) (Enum, error) {
	row := config.DB.QueryRow(
		"SELECT "+
			"* "+
			"FROM "+
			"enum_tables "+
			"WHERE "+
			"id = $1",
		id)

	e := Enum{}
	err := row.Scan(
		&e.ID,
		&e.UpdatedAt,
		&e.Type,
		&e.Name,
		&e.Symbol,
		&e.Ordinal,
		&e.Attrs,
		&e.DeletedAt,
	)

	if err != nil {
		return e, nil
	}

	return e, nil
}
