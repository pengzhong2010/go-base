package mysqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var Mysqldb *gorm.DB

func init() {
	var (
		env_type  string
		db_host   string
		db_port   string
		db_user   string
		db_passwd string
		db_name   string
	)
	env_type = os.Getenv("ENV_TYPE")

	db_host = os.Getenv("MYSQL_HOST")
	db_port = os.Getenv("MYSQL_PORT")
	db_user = os.Getenv("MYSQL_USER")
	db_passwd = os.Getenv("MYSQL_PASSWD")
	db_name = os.Getenv("MYSQL_NAME")
	var err error
	if env_type == "dev" {
		Mysqldb, err = gorm.Open("mysql", "root:caidaoninb@tcp(localhost:33306)/cmdb1?charset=utf8&parseTime=True&loc=Local")
	} else {
		Mysqldb, err = gorm.Open("mysql", db_user+":"+db_passwd+"@tcp("+db_host+":"+db_port+")/"+db_name+"?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	Mysqldb.DB().SetMaxIdleConns(20)
	Mysqldb.DB().SetMaxOpenConns(20)
}
