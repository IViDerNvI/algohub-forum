package mysql

import (
	"fmt"
	"sync"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/internal/model/options"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Posts() store.PostStore {
	return newPosts(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts options.MYSQLOptions) (store.Factory, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
			opts.Username,
			opts.Password,
			opts.Host,
			opts.Database,
			true,
			"Local")
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			logrus.Fatalf("connect to database failed: %s", err)
		}
		cleanDatabase(db)
		migrateDatabase(db)
		mysqlFactory = &datastore{db}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&model.User{}); err != nil {
		return errors.Wrap(err, "drop user table failed")
	}

	if err := db.Migrator().DropTable(&model.Post{}); err != nil {
		return errors.Wrap(err, "drop post table failed")
	}

	return nil
}

func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}

	if err := db.AutoMigrate(&model.Post{}); err != nil {
		return errors.Wrap(err, "migrate post model failed")
	}

	return nil
}

func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}
