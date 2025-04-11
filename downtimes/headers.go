package downtimes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io/util"
	"log"
	"net/http"
)

// func statusCheck(resp http.Response, err error) (http.Response, error) {
// 	if err != nil {
// 		return http.Response{}, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return http.Response{}, fmt.Errorf("Error: %s, ReturnCode: %d", err, resp.StatusCode)
// 	}
// 	return resp, nil
// }

const GET_HOSTS str = "https://api.ddog-gov.com/api/v1/hosts"

func statusCheck(resp http.Response, err error) (http.Response, error) {
	switch {
	case err != nil:
		return http.Response{}, err
	case resp.StatusCode != http.StatusOK:
		return http.Response{}, fmt.Errorf("Error: %v, ReturnCode: %d", err, resp.StatusCode)
	default:
		return resp, nil
	}
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return fmt.Println(string(body))
}



func handleResponse(r *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("Error: %s, %v", err, r)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: Status code is not OK")
	}
	return r, nil
}

func get(api_key string, app_key string, url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	_, err = handleResponse(resp, err)
	if err != nil {
		return "", 0, err
	}

	// Process the response here
	return "", resp.StatusCode, nil
}





func post(api_key str, app_key str, url str) (resp str) {
	resp, err := http.Post(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Println(string(body))
}


client := http.Client{

}
