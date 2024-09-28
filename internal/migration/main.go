package migration

import (
	"fmt"
	"funnymovies/config"
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
	migrationutil "funnymovies/util/migration"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Run executes the migration
func Run() (respErr error) {
	fmt.Println("Start migration function...")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := dbutil.New(cfg.DbDsn, true)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				respErr = fmt.Errorf("%s", x)
			case error:
				respErr = x
			default:
				respErr = fmt.Errorf("unknown error: %+v", x)
			}
		}
	}()

	fmt.Println("db connected: " + db.Name())

	initSQL := "CREATE TABLE IF NOT EXISTS migrations (id VARCHAR(255) PRIMARY KEY)"
	if err := db.Exec(initSQL).Error; err != nil {
		return err
	}

	migrationutil.Run(db, []*gormigrate.Migration{
		{
			ID: "202428090942",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					&model.User{},
					&model.Link{},
				); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					"user",
					"link",
				)
			},
		},
	})
	return nil
}
