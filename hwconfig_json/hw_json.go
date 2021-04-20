package hwconfig_json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type South_com struct {
	com_rs232   Com_rs232
	com_rs485_1 Com_rs485_1
	com_rs485_2 Com_rs485_2
}

type Com_rs232 struct {
	Com_enable   interface{} `json:"rs232_enable"`
	Com_port     interface{} `json:"rs232_port"`
	Com_baudrate interface{} `json:"rs232_baudrate"`
	Com_databits interface{} `json:"rs232_databits"`
	Com_stopbits interface{} `json:"rs232_stopbits"`
}

type Com_rs485_1 struct {
	Com_enable   interface{} `json:"rs485_1_enable"`
	Com_port     interface{} `json:"rs485_1_port"`
	Com_baudrate interface{} `json:"rs485_1_baudrate"`
	Com_databits interface{} `json:"rs485_1_databits"`
	Com_stopbits interface{} `json:"rs485_1_stopbits"`
}

type Com_rs485_2 struct {
	Com_enable   interface{} `json:"rs485_2_enable"`
	Com_port     interface{} `json:"rs485_2_port"`
	Com_baudrate interface{} `json:"rs485_2_baudrate"`
	Com_databits interface{} `json:"rs485_2_databits"`
	Com_stopbits interface{} `json:"rs485_2_stopbits"`
}

type Modbus_config struct {
	Modbus_mode      interface{} `json:"modbus_mode"`
	Modbus_device_id interface{} `json:"modbus_device_id"`
}

// 解析 hwconfig.json 文件
func Hw_read_json() {

	file_json, err := os.ReadFile("config/hwconfig.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 读取并解析 RS232 配置
	rs232_json := Com_rs232{}
	rs485_1_json := Com_rs485_1{}
	rs485_2_json := Com_rs485_2{}
	modbus_json := Modbus_config{}

	err = json.Unmarshal(file_json, &rs232_json)
	if err != nil {
		log.Fatal(err)
	} else {
		err = json.Unmarshal(file_json, &rs485_1_json)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		err = json.Unmarshal(file_json, &rs485_2_json)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		err = json.Unmarshal(file_json, &modbus_json)
	}

	fmt.Println(rs232_json)
	fmt.Println(rs485_1_json)
	fmt.Println(rs485_2_json)
	fmt.Println(modbus_json)
}
