package bootstrap

import (
	"os"
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
	"github.com/bradfitz/gomemcache/memcache"
)


type AppConfig struct {
    Database struct {
        Name string `yaml:"name"`
        Password string `yaml:"password"`
        User string `yaml:"user"`
        Addr string `yaml:"addr"`
    }

	Memcached struct {
        Addr string `yaml:"addr"`
    }

	Port string `yaml:"port"`
}

type Dependecies struct {
	Database sqlbuilder.Database
	AppConfig AppConfig
	Cache *memcache.Client
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
		log.Fatalf("Unable to load configuration\n#%v", err)
	}

	// ConnectionURL implements a MySQL connection struct.
	DBsettings := mysql.ConnectionURL{
		User     : config.Database.User,
		Password : config.Database.Password,
		Host     : config.Database.Addr,
		Database : config.Database.Name,
	}
  
  	db, errDb := mysql.Open(DBsettings)

	log.Printf("Using Database [%s@%s]", config.Database.Name, config.Database.Addr)

	if errDb != nil {
		log.Fatalf("Unable to connect to DB \n#%v", errDb)
	}

	errDb = db.Ping() 

	if errDb != nil {
		log.Fatalf("DB connection error\n%#v", errDb)
	}

	log.Printf("Memcached [%+v]", config.Memcached);

	mc := memcache.New( config.Memcached.Addr )

	deps := Dependecies{
		Database  : db,
		AppConfig : config,
		Cache     : mc,
	}

	return deps
}