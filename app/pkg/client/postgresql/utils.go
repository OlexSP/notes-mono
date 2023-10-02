package psql

import (
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
)

func ParsePgError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return fmt.Errorf(
			"database error. message:%s, detail:%s, where:%s, sqlstate:%s",
			pgErr.Message,
			pgErr.Detail,
			pgErr.Where,
			pgErr.SQLState(),
		)
	}

	return err
}

func PrettySQL(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
