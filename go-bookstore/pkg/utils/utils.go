package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func ParseBody(r *http.Request,x interface{}){
	if body,err := ioutil.ReadAll(r.Body)
	 err == nil{
		if err:= json.Unmarshal([]byte(body),x); err !=nil{
			fmt.Printf(string(err.Error()))
			return
		}
	}else{
		fmt.Println("Error while reading body")
	}
}