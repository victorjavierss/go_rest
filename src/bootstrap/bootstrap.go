package bootstrap

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"database/sql"

    "gopkg.in/yaml.v2"
	_ "github.com/go-sql-driver/mysql"
)

type AppConfig struct {
    Database struct {
        Name string `yaml:"name"`
        Password string `yaml:"password"`
        User string `yaml:"user"`
        Addr string `yaml:"addr"`
    }
	Port string `yaml:"port"`
}

type Dependecies struct {
	Database *sql.DB
	AppConfig AppConfig
}

func Init () Dependecies {

	appConfigFile := os.Getenv("APP_CONFIG")
	yamlFile, err := ioutil.ReadFile(appConfigFile)

	log.Printf("Loading configuration [%s]", appConfigFile)

	if err != nil {
		log.Fatalf("Unable to locate config file\n#%v", err)
	}

	var config AppConfig;
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatalf("Unable to load configuration\n %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.Database.User, config.Database.Password, config.Database.Addr, config.Database.Name)
	log.Printf("DB DSN [%s]", dsn)

	db, errDb := sql.Open("mysql", dsn)

	if errDb != nil {
		log.Fatalf("Unable to connect to DB \n %v", errDb)
	}

	errDb = db.Ping() 

	if errDb != nil {
		log.Fatalf("DB connection Error \n %#v", errDb)
	}

	deps := Dependecies{
		Database  : db,
		AppConfig : config,
	}

	return deps
}