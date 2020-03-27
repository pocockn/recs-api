package persistance

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"github.com/pocockn/recs-api/config"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	// GormDB holds a database connection.
	GormDB struct {
		maxConnections int
		url            string
	}
)

// NewConnection creates a new connection for the database.
func NewConnection(config config.Config) (*GormDB, error) {
	return &GormDB{
		maxConnections: config.Database.MaxConnections,
		url:            generateURL(config),
	}, nil
}

// Connect connects to the database and passes back the connection so we can
// use it throughout the application
func (g GormDB) Connect() (*gorm.DB, error) {
	var gormDB *gorm.DB
	var err error

	for i := 0; i <= 30; i++ {
		gormDB, err = gorm.Open("mysql", g.url)
		if err == nil {
			err := gormDB.DB().Ping()
			if err == nil {
				gormDB.LogMode(true)
				break
			}
		}

		if i == 15 {
			log.Fatalf("unable to connect to %s after 30 seconds", g.url)
		}

		logrus.Infof("%d attempt at connecting to the DB \n", i)
		time.Sleep(2 * time.Second)
	}

	maxConnsPerContainer := g.maxConnections / 4
	gormDB.DB().SetMaxOpenConns(maxConnsPerContainer / 2)

	return gormDB, nil
}

func generateURL(config config.Config) string {
	templateString := "%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4"

	return fmt.Sprintf(
		templateString,
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName,
	)
}
