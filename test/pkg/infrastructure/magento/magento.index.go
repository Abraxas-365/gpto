package magento

import (
	"database/sql"
	"github/Abraxas-365/akeneo-connector/pkg/core/ports"
)

type magnetoRepo struct {
	db       *sql.DB
	url      string
	user     string
	password string
}

func MagentoFactory(db *sql.DB, url string, user string, password string) ports.Magento {
	return &magnetoRepo{
		db:       db,
		url:      url,
		user:     user,
		password: password,
	}
}
