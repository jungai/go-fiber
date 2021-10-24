package utils

import "fmt"

type ConfigDb struct {
	user     string
	password string
	hostname string
	dbname   string
}

var JungaiTest ConfigDb = ConfigDb{
	user:     "host",
	password: "host",
	hostname: "127.0.0.1:3306",
	dbname:   "jungai_test",
}

func GetDsn(c *ConfigDb) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.user, c.password, c.hostname, c.dbname)
}
