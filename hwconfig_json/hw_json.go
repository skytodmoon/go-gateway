package hwconfig_json

import (
	"encoding/json"
	"log"
	"os"
)

type South_com struct {
	com_rs232   Com_rs232
	com_rs485_1 Com_rs485_1
	com_rs485_2 Com_rs485_2
}

type Com_rs232 struct {
	Com_enable   bool   `json:"rs232_enable"`
	Com_port     string `json:"rs232_port"`
	Com_baudrate int    `json:"rs232_baudrate"`
	Com_databits int    `json:"rs232_databits"`
	Com_stopbits int    `json:"rs232_stopbits"`
	Com_parity   string `json:"rs232_parity"`
}

type Com_rs485_1 struct {
	Com_enable   bool   `json:"rs485_1_enable"`
	Com_port     string `json:"rs485_1_port"`
	Com_baudrate int    `json:"rs485_1_baudrate"`
	Com_databits int    `json:"rs485_1_databits"`
	Com_stopbits int    `json:"rs485_1_stopbits"`
	Com_parity   string `json:"rs485_1_parity"`
}

type Com_rs485_2 struct {
	Com_enable   bool   `json:"rs485_2_enable"`
	Com_port     string `json:"rs485_2_port"`
	Com_baudrate int    `json:"rs485_2_baudrate"`
	Com_databits int    `json:"rs485_2_databits"`
	Com_stopbits int    `json:"rs485_2_stopbits"`
	Com_parity   string `json:"rs485_2_parity"`
}

type Com_can struct {
	Com_enable   bool `json:"can_enable"`
	Com_port     int  `json:"can_port"`
	Com_baudrate int  `json:"can_baudrate"`
	Com_mode     int  `json:"can_mode"`
}

// Modbus 参数 结构体
type Modbus_config struct {
	Modbus_enable    bool        `json:"modbus_enable"`
	Modbus_mode      interface{} `json:"modbus_mode"`
	Modbus_timeout   int         `json:"modbus_timeout"`
	Modbus_device_id interface{} `json:"modbus_device_id"`
	// TCP 参数
	Modbus_tcp_addr string `json:"modbus_tcp_addr"`
	Modbus_tcp_port string `json:"modbus_tcp_port"`

	// 从机
	Modbus_slave_id_1 byte `json:"modbus_slave_id_1"`
	Modbus_slave_id_2 byte `json:"modbus_slave_id_2"`
	Modbus_slave_id_3 byte `json:"modbus_slave_id_3"`
	Modbus_slave_id_4 byte `json:"modbus_slave_id_4"`
	Modbus_slave_id_5 byte `json:"modbus_slave_id_5"`
	Modbus_slave_id_6 byte `json:"modbus_slave_id_6"`
}

// MQTT 参数 结构体
type Mqtt_config struct {
	Mqtt_version   string `json:"mqtt_version"`
	Mqtt_broker    string `json:"mqtt_broker"`
	Mqtt_port      int    `json:"mqtt_port"`
	Mqtt_use_tls   bool   `json:"mqtt_use_tls"`
	Mqtt_client_id string `json:"mqtt_clinet_id"`
	Mqtt_user_name string `json:"mqtt_user_name"`
	Mqtt_user_pass string `json:"mqtt_user_pass"`
	Mqtt_topic     string `json:"mqtt_topic"`
	Mqtt_keepalive int    `json:"mqtt_keepalive"`
}

// 解析 hwconfig.json 文件
func Hw_read_json(s interface{}) error {

	file_json, err := os.ReadFile("config/hwconfig.json")
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 解析 JSON 配置文件并放回到 s 结构体中
	err = json.Unmarshal(file_json, &s)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
