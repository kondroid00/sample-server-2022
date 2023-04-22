package ent

import (
	"database/sql"
	"github.com/kondroid00/sample-server-2022/package/errors"
)

func (tx *Tx) RollbackUnlessCommitted() error {
	err := tx.Rollback()
	if errors.Is(err, sql.ErrTxDone) {
		return nil
	}
	return err
}
