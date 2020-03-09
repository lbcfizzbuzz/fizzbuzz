package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/samyy321/fizzbuzz/config"
	ds "github.com/samyy321/fizzbuzz/datastore"
	serv "github.com/samyy321/fizzbuzz/server"
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
	config := conf.Configuration{}
	if err := config.Read(args.ConfigPath); err != nil {
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

	// Launch the server
	s := serv.Server{Db: datastore}
	s.Run()
}
