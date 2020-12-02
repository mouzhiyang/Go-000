package work

import (
	"database/sql"

	"github.com/pkg/errors"
)

func dao() ([]byte, error) {
	return nil, errors.Wrap(sql.ErrNoRows, "NotFound")
}
