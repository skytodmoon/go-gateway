package modbus_mod

import (
	"fmt"
	"log"
	"main/hwconfig_json"
)

// modbus 初始化
func Modbus_init() error {
	// 读取 MODBUS 配置信息
	mb_config := hwconfig_json.Modbus_config{}
	rs232_config := hwconfig_json.Com_rs232{}
	rs485_1_config := hwconfig_json.Com_rs485_1{}
	rs485_2_config := hwconfig_json.Com_rs485_2{}

	// 解析 modbus 配置结构体
	err := hwconfig_json.Hw_read_json(&mb_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = hwconfig_json.Hw_read_json(&rs232_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = hwconfig_json.Hw_read_json(&rs485_1_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = hwconfig_json.Hw_read_json(&rs485_2_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(mb_config)
	fmt.Println(rs232_config)
	fmt.Println(rs485_1_config)
	fmt.Println(rs485_2_config)

	return nil
}
