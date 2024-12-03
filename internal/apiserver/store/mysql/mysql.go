package mysql

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/internal/model/options"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type datastore struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Posts() store.PostStore {
	return newPosts(ds)
}

func (ds *datastore) Likes() store.LikeStore {
	return newLikes(ds)
}

func (ds *datastore) Tokens() store.TokenStore {
	return newTokens(ds)
}

func (ds *datastore) Problems() store.ProbStore {
	return newProblems(ds)
}

func (ds *datastore) Close() error {
	_, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return ds.rdb.Close()
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

		rdb := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", "127.0.0.1", 6379), // Redis服务地址
			Password:     "",                                      // 密码
			DB:           0,                                       // 默认数据库
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
		})

		mysqlFactory = &datastore{
			db:  db,
			rdb: rdb,
		}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

func flushRedis(rdb *redis.Client) error {
	var ctx context.Context
	err := rdb.FlushDB(ctx).Err()
	if err != nil {
		logrus.Println("Failed to flush database:", err)
		return err
	}

	return nil
}

func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&model.User{}); err != nil {
		return errors.Wrap(err, "drop user table failed")
	}

	if err := db.Migrator().DropTable(&model.Post{}); err != nil {
		return errors.Wrap(err, "drop post table failed")
	}

	if err := db.Migrator().DropTable(&model.Like{}); err != nil {
		return errors.Wrap(err, "drop like table failed")
	}

	if err := db.Migrator().DropTable(&model.Problem{}); err != nil {
		return errors.Wrap(err, "drop prob table failed")
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

	if err := db.AutoMigrate(&model.Like{}); err != nil {
		return errors.Wrap(err, "migrate like table failed")
	}

	if err := db.AutoMigrate(&model.Problem{}); err != nil {
		return errors.Wrap(err, "migrate prob table failed")
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
