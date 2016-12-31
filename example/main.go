package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tomsteele/boom"
)

func main() {
	err := boom.BadRequest("invalid payload", nil)
	fmt.Println(err)
	data, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	err = boom.BadImplementation(errors.New("sql: no rows found in result set"))
	fmt.Println(err)
	data, err = json.Marshal(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
