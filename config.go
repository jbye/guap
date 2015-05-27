package main

import (
  "io/ioutil"
  "os"
  "time"

  logger "github.com/Sirupsen/logrus"
  "gopkg.in/yaml.v2"
)

// -
var Config struct {
  Intervals struct {
    Heroes  time.Duration
    Leagues time.Duration
  }
  Rethink struct {
    Host     string
    Database string
  }
  Steam struct {
    Key string
  }
}

func loadConfig() {
  yamlPath := "config.yaml"
  if _, err := os.Stat(yamlPath); err != nil {
    logger.Error("Config path is not valid")
    panic(err)
  }
  ymlData, err := ioutil.ReadFile(yamlPath)
  if err != nil {
    logger.Error("Unable to parse config", err)
    panic(err)
  }
  err = yaml.Unmarshal([]byte(ymlData), &Config)
  if err != nil {
    logger.Error("Unable to unmarshal config", err)
    panic(err)
  }

  // ENV overrides
  /*
     servicePort := os.Getenv("PORT")
     if servicePort != "" {
       app.Config.Service.Host = ":" + servicePort
     }

     databaseURL := os.Getenv("DATABASE_URL")
     var connectionString string
     if databaseURL != "" {
       connectionString, _ := pq.ParseURL(databaseURL)
       connectionString += " sslmode=require"
     } else {
       connectionString = "postgres://" + app.Config.Database.User + ":" + app.Config.Database.Password +
         "@" + app.Config.Database.Host + ":5432/" + app.Config.Database.Name +
         "?sslmode=disable"
     }
     app.Config.Database.ConnectionString = connectionString
  */
}
