package helper

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gadhittana01/todolist/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(c *config.GlobalConfig) {
	path := "config/todolist-http.yaml"
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	replaceWithENVValue(c)

}

func replaceWithENVValue(c *config.GlobalConfig) {
	// DB
	DBHost := os.Getenv("MYSQL_HOST")
	if DBHost != "" {
		c.DB.Host = DBHost
	}

	DBPort := os.Getenv("MYSQL_PORT")
	if DBPort != "" {
		port, err := strconv.Atoi(DBPort)
		if err != nil {
			log.Fatalf("Error Parse : %v", err)
		}
		c.DB.Port = int32(port)
	}

	DBUser := os.Getenv("MYSQL_USER")
	if DBUser != "" {
		c.DB.User = DBUser
	}

	DBPassword := os.Getenv("MYSQL_PASSWORD")
	if DBPassword != "" {
		c.DB.Password = DBPassword
	}

	DBName := os.Getenv("MYSQL_DBNAME")
	if DBName != "" {
		c.DB.Name = DBName
	}
}
