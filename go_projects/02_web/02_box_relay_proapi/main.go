package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	authurl := "https://prov-api-aws-relay-prod.ibmbxrelay.com/gafapp/ido73g/svcacct"
	tenantQueryURL := "https://prov-api-aws-relay-prod.ibmbxrelay.com/ProvisionAPI/api/tenant/query/getTenantsByRange?startIndex=0&endIndex=4&sortProperty=enterpriseId&ascending=true"
	println("Connecting to " + authurl)

	req1, err := http.NewRequest("GET", authurl, nil)
	req1.SetBasicAuth("x", "x")
	cli := &http.Client{}

	resp1, err := cli.Do(req1)

	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()

	println("Successfully authenticated against box-relay provisioning-api")
	println("======================================")
	println("Printing HTTP Headers....")
	fmt.Println(resp1.Header)
	println("======================================")
	println("Running Tenant Query: " + tenantQueryURL)

	req2, err := http.NewRequest("GET", tenantQueryURL, nil)
	req2.Header = req1.Header
	resp2, err := cli.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()
	content2, err := ioutil.ReadAll(resp2.Body)
	println(string(content2))
}
