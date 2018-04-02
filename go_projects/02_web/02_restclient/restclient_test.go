package restclient

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestAdder2(t *testing.T) {
	println("TestAdder2")
	resp, err := http.Get("http://quotes.rest/qod.json?category=inspire")
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		t.Fatal("service returned status code " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	//contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(contents))

}

func TestAdder1(t *testing.T) {
	println("TestAdder1")
	resp, err := http.Get("http://quotes.rest/qod.json?category=inspire")
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		t.Fatal("service returned status code " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	//contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(contents))

}
