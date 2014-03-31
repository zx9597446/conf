//load json config file to struct,
//create file if not exist
package conf

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

//load config file to struct
func Load(filename string, result interface{}) error {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		cf, err := os.Create(filename)
		defer cf.Close()
		if err != nil {
			log.Println(err)
			return err
		}
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Println(err)
			return err
		}
		cf.Write(data)
		msg := "config file not exist, created"
		log.Println(msg)
		return errors.New(msg)
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return err
}
