package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &x)
	if err != nil {
		return err
	}

	return nil
}

// func ParseBody(r *http.Request, x interface{}) {

// 	fmt.Println("FOOPPPP")
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		err = json.Unmarshal([]byte(body), x)
// 		if err != nil {
// 			return
// 		}
// 	}
// 	fmt.Println("FOOPPPP")
// }
