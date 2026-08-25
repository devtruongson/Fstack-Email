package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

// V6_3_0 performs the DB migrations for v6.3.0.
func V6_3_0(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf, lo *log.Logger) error {
	if _, err := db.Exec(`
		INSERT INTO settings (key, value) VALUES
			('app.daily_send_limit', '0'),
			('app.message_send_delay', '"1s"')
			ON CONFLICT DO NOTHING;
	`); err != nil {
		return err
	}

	return nil
}
