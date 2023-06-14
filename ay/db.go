package ay

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func GetDB() (err error) {
	InitializeRedis()
	var option gorm.Dialector

	switch Yaml.GetString("sql.type") {
	case "mysql":
		dsn := Yaml.GetString("sql.user") + ":" + Yaml.GetString("sql.password") + "@tcp(" + Yaml.GetString("sql.localhost") + ":" + Yaml.GetString("sql.port") + ")/" + Yaml.GetString("sql.database") + "?charset=utf8mb4&parseTime=true&loc=Local"
		option = mysql.New(mysql.Config{DSN: dsn})
	case "sqlite":
		//option = sqlite.Open(Yaml.GetString("sql.localhost"))
	default:

	}

	if Db, err = gorm.Open(option, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		return
	}

	return
}

func IgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

func IgnoreNotFoundResultOk[T any](result T, err error) (T, bool, error) {
	ok := true
	if err != nil {
		ok = false
		err = IgnoreNotFound(err)
	}
	return result, ok, err
}

func IgnoreNotFoundReturn[T any](r T, err error) (T, error) {
	err = IgnoreNotFound(err)
	return r, err
}

func IsFound(err error) (bool, error) {
	if err == nil {
		return true, nil
	}
	return false, IgnoreNotFound(err)
}

func InsertIGNORE(db *gorm.DB) *gorm.DB {
	return db.Clauses(clause.Insert{Modifier: "IGNORE"})
}

func ForUpdate(db *gorm.DB) *gorm.DB {
	return db.Clauses(clause.Locking{Strength: "UPDATE"})
}
