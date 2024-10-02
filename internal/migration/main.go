package migration

import (
	"fmt"
	"funnymovies/config"
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
	migrationutil "funnymovies/util/migration"
	"strings"

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
		{
			ID: "202410021101",
			Migrate: func(tx *gorm.DB) error {
				changes := []string{
					`INSERT INTO public."user" 
						(email, "password", created_at, updated_at, deleted_at)
					VALUES
						('testemail@gmail.com', '$2a$10$9f8GMasxXmawm4WMuoxbuuzBLMvHWeD5jgIMOG2QDLWNKPBodKMe2', now(), now(), NULL),
						('oldd@gmail.com', '$2a$10$SKzIgnu9ko9Nw4LguHmFweLPkorjXr53RYiZ94Lb.I62B2I4mEnT6', now(), now(), NULL),
						('divine@gmail.com', '$2a$10$abv2fwmXUC.7MbakBvMOA.uhwaQ0pYTShU55L5YAmaJuGDvi5hamm', now(), now(), NULL),
						('thunder@gmail.com', '$2a$10$y.8fnc5AdctYDB2RLM/.4.fKKARaEh0SoD7fTZP/FqCsamRnUYF66', now(), now(), NULL);`,

					`INSERT INTO public.link
						(url, user_id, created_at, updated_at, deleted_at)
					VALUES
						('https://www.youtube.com/watch?v=RcNy4NrBltg', 1, now(), now(), NULL),
						('https://www.youtube.com/watch?v=U44qKaKpAMk', 2, now(), now(), NULL),
						('https://www.youtube.com/watch?v=JgTZvDbaTtg', 3, now(), now(), NULL),
						('https://www.youtube.com/watch?v=cJD4fc5l3fM', 4, now(), now(), NULL),
						('https://www.youtube.com/watch?v=svv-IYQHKNc', 4, now(), now(), NULL),
						('https://www.youtube.com/watch?v=Phrs_u5Jofk', 3, now(), now(), NULL),
						('https://www.youtube.com/watch?v=e5Td3zrVdX4', 2, now(), now(), NULL),
						('https://www.youtube.com/watch?v=n6Pnzi6r9NU', 1, now(), now(), NULL),
						('https://www.youtube.com/watch?v=QY7qwMuxll8', 2, now(), now(), NULL),
						('https://www.youtube.com/watch?v=cCvr5QMnsG8', 3, now(), now(), NULL),
						('https://www.youtube.com/watch?v=JgTZvDbaTtg', 4, now(), now(), NULL),
						('https://www.youtube.com/watch?v=2YM4j-oP_qQ', 1, now(), now(), NULL),
						('https://www.youtube.com/watch?v=83OiPzXQvIc', 2, now(), now(), NULL),
						('https://www.youtube.com/watch?v=OZYYZi0Hoo4', 3, now(), now(), NULL),
						('https://www.youtube.com/watch?v=kPFPRs211bQ', 1, now(), now(), NULL),
						('https://www.youtube.com/watch?v=JYPIDIQSvb8', 4, now(), now(), NULL);`,
				}

				return migrationutil.ExecMultiple(tx, strings.Join(changes, " "))
			},
			Rollback: func(tx *gorm.DB) error {
				changes := []string{
					`DELETE FROM public."user" WHERE id IN (1,2,3,4)`,
					`ALTER SEQUENCE "user_id_seq" RESTART WITH 1;`,
					`DELETE FROM public.link WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16)`,
					`ALTER SEQUENCE "link_id_seq" RESTART WITH 1;`,
				}
				return migrationutil.ExecMultiple(tx, strings.Join(changes, " "))
			},
		},
	})
	return nil
}
