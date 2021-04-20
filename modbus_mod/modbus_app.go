package modbus_mod

import (
	"fmt"
	"log"
	"main/hwconfig_json"
)

// modbus 初始化
func Modbus_init() error {
	mb_config := hwconfig_json.Modbus_config{}

	err := hwconfig_json.Hw_read_json(&mb_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(mb_config)

	return nil
}
