package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)


type Config struct {
    Tables []Table `json:"tables"`
}

type Columns struct {
    Columns []Column `json:"columns"`
}

// User struct which contains a name
// a type and a list of social links
type Table struct {
	Name   string `json:"name"`
	Columns []Column `json:"columns"`
	Schema string `json:"schema"`
}

type Column struct {
	Name   string `json:"name"`
	Generator string `json:"generator"`
	Min int `json:"min"`
	Max int `json:"max"`
	Timezone string `json:"timezone"`
	SQLFunction string `json:"function"`
}


func read_config(fileName string) Config {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("Successfully Opened : ", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	// we initialize our Tables array
	var tables Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'tables' which we defined above
	if err := json.Unmarshal(byteValue, &tables); err != nil {
		panic(err)
	}

	tlist := []string{}

	for i := 0; i < len(tables.Tables); i++ {
		log.Debug("uuu", i)
		tlist = append(tlist, tables.Tables[i].Name)

		for j := 0; j < len(tables.Tables[i].Columns); j++ {

			fullname := fmt.Sprintf("%s.%s.%s", tables.Tables[i].Schema,
				tables.Tables[i].Name,
				tables.Tables[i].Columns[j].Name)
			log.Debug(fullname)
		}
	}

	return tables
}

func get_cols(conf Table, columName string) (Column, int) {
	err := 1
	var result Column
	for j := 0; j < len(conf.Columns); j++ {
		if columName == conf.Columns[j].Name {
			result = conf.Columns[j]
			err = 0
		}

	}

	return result, err
}
