package connections

import (
	"fmt"

	"shared-wallet-service/infrastructure/config"

	"github.com/go-sql-driver/mysql"
)

func GetDBConfig(nameDB string, action string) mysql.Config {
	return mysql.Config{
		User:                 fmt.Sprintf("%s_%s", config.Get().UString(fmt.Sprintf("%s.user", nameDB)), action),
		Passwd:               config.Get().UString(fmt.Sprintf("%s.password", nameDB)),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Get().UString(fmt.Sprintf("%s.host", nameDB)), config.Get().UString(fmt.Sprintf("%s.port", nameDB))),
		DBName:               config.Get().UString(fmt.Sprintf("%s.schema", nameDB)),
		AllowNativePasswords: true,
	}
}
