package main

import (
	"os"

	"github.com/pjpillai/MyProjects/go_projects/03_database/01_cloudantdb/cloudant"
	couchdb "github.com/timjacobi/go-couchdb"
)

var testClient *cloudant.Client
var testDB *cloudant.DB
var err error

func main() {

	url := ""
	username := ""
	password := ""

	if testClient, err = cloudant.NewClient(url, username, password); err != nil {
		os.Exit(1)
	}

	dbName := "primary_ansible_pj_ecmconfig"
	// Get the test DB object
	testDB = testClient.DB(dbName)

	println(testDB.Name)

	type alldocsResult struct {
		TotalRows int `json:"total_rows"`
		Offset    int
		Rows      []map[string]interface{}
	}

	var result alldocsResult
	err := testDB.AllDocs(&result, couchdb.Options{
		"offset": 5,
		"limit":  100,
	})

	if err != nil {
		panic(err)
	}

	for k, v := range result.Rows {
		println(k, v)
	}

	println(result.Rows)

}
