package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetIPAdress() (string, error) { // TODO wonering if there is a cleaner way to get public ip.
	url := "https://api.ipify.org?format=text"
	fmt.Printf("Getting IP address from  ipify ...\n")
	resp, err := http.Get(url)
	if err != nil {
		return ",", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ",", err
	}
	fmt.Println(string(ip))
	return string(ip), nil
}
