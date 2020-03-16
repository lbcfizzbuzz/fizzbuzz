package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	cfg "github.com/lbcfizzbuzz/fizzbuzz/config"
	ds "github.com/lbcfizzbuzz/fizzbuzz/datastore"
	srv "github.com/lbcfizzbuzz/fizzbuzz/server"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	// Read the command line arguments
	var args Args
	args.Init()

	err := args.Check()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the configuration file
	config := cfg.Configuration{}
	if err := config.Read(args.ConfigPath); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting ...")
		return
	}

	err = config.Validate()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting ...")
		return
	}

	// Get the datastore
	datastore, err := ds.GetDatastore(config)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting ...")
		return
	}
	err = datastore.Init()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting ...")
		return
	}

	// Get the logger
	logger := log.New()
	if config.LogInJSON {
		logger.SetFormatter(&log.JSONFormatter{})
	} else {
		logger.SetFormatter(&log.TextFormatter{})
	}
	if config.LogFilePath != "" {
		f, err := os.OpenFile(config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Exiting ...")
			return
		}
		defer f.Close()
		fmt.Println("Logs will be in file " + config.LogFilePath)
		logger.SetOutput(f)
	} else {
		fmt.Println("Stdout will be used to logs")
		logger.SetOutput(os.Stdout)
	}

	// Launch the server
	s := srv.Server{Db: datastore, Config: &config, Logger: logger}
	fmt.Println("Launching server listening on port " + strconv.Itoa(config.Port))
	s.Run()
}
