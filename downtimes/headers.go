package downtimes

import (
	// "encoding/json"
	"bytes"
	"fmt"
	// "io"
	// "io/util"
	// "log"
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

// const GET_HOSTS str = "https://api.ddog-gov.com/api/v1/hosts"
// func newRequestStatusCheck(req *http.Request, err error) (*http.Response, error) {
//     if err != nil {
//         return nil, err
//     }
//
//     if req.StatusCode != http.StatusOK {
//         return nil, fmt.Errorf("Error: %v, ReturnCode: %d", err, req.StatusCode)
//     }
//
//     return &http.Response{StatusCode: req.StatusCode}, nil
// }

func newRequestStatusCheck(req *http.Request, err error) (*http.Request, error) {
	switch {
	case err != nil:
		return nil, fmt.Errorf("Request Error: %v", err)
	default:
		return req, nil
	}
}

func responseStatusCheck(resp *http.Response, err error) (http.Response, error) {
	switch {
	case err != nil:
		return http.Response{}, err
	case resp.StatusCode != http.StatusOK:
		return http.Response{}, fmt.Errorf("Error: %v, ReturnCode: %d", err, resp.StatusCode)
	default:
		return *resp, nil
	}
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

func Get(api_key string, app_key string, url string) (string, int, error) {
	body := bytes.NewBufferString("")
	req, err := http.NewRequest("GET", url, body)
	newRequestStatusCheck(req, err)
	switch {
	case err != nil:
		fmt.Println(err)
	default:
		client := &http.Client{}
		resp, err := client.Do(req)
		handleResponse(resp, err)
	}
	return "", 0, nil
}

// return fmt.Println(body)
// resp, err := http.Get(url)
// if err != nil {
// 	return "", 0, err
// }
// defer resp.Body.Close()
//
// _, err = handleResponse(resp, err)
// if err != nil {
// 	return "", 0, err
// }

// Process the response here
// return "", resp.StatusCode, nil
// }

// func Post(api_key string, app_key string, url string) (resp string) {
// 	body := bytes.NewBufferString("")
// 	resp, err := http.NewRequest("POST", url, body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()
//
// 	body, err := util.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return fmt.Println(body)
// }

// client := http.Client{
//
// }
