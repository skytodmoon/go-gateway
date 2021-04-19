package hwconfig_json

import (
	"encoding/json"
	"log"
	"os"
)

type Com_rs232 struct {
	Com_enable   interface{} `json:"rs232_enable"`
	Com_port     interface{} `json:"rs232_port"`
	Com_baudrate interface{} `json:"rs232_baudrate"`
	Com_databits interface{} `json:"rs232_databits"`
	Com_stopbits interface{} `json:"rs232_stopbits"`
}

type Com_rs485 struct {
	Com_enable   interface{}
	Com_port     interface{}
	Com_baudrate interface{}
	Com_databits interface{}
	Com_stopbits interface{}
}

func hw_read_json() {

	json_rs232, err := os.Open("config/hwconfig.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer json_rs232.Close()

	config_rs232 := Com_rs232{}

	json_rs232_dec := json.NewDecoder(json_rs232)
	err = json_rs232_dec.Decode(config_rs232)

	err := json.Unmarshal(str, &config_rs232)

}
