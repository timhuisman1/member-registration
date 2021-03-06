package dbRepo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
)

type DBRepo struct {
	db *sqlx.DB
}

func InitDBRepo(db *sqlx.DB) DBRepo {
	return DBRepo{
		db: db,
	}
}

// Logging
func (repo DBRepo) logDebug(domain string, message string) {
	logger.Debug(fmt.Sprintf("repositories/db/%s", domain), message)
}
func (repo DBRepo) logInfo(domain string, message string) {
	logger.Debug(fmt.Sprintf("repositories/db/%s", domain), message)
}
func (repo DBRepo) logWarning(domain string, err error) {
	logger.Warn(fmt.Sprintf("repositories/db/%s", domain), err)
}
func (repo DBRepo) logError(domain string, err error) {
	logger.Warn(fmt.Sprintf("repositories/db/%s", domain), err)
}
