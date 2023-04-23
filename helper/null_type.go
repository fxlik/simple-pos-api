package helper

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type NullInt32 struct {
	sql.NullInt32
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}
