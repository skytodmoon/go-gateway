package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"main/modbus_mod"
	"main/mqtt_mod"
)

func show_logo() {
	logo_buf, err := ioutil.ReadFile("logo.log")
	if err != nil {
		fmt.Println("logo file doesn't exist....check it?")
	}
	fmt.Println(string(logo_buf))
}

func main() {
	// 展示公司logo
	show_logo()

	err := modbus_mod.Modbus_init()
	if err != nil {
		fmt.Println("Modbus 初始化失败~")
	}

	err = mqtt_mod.Mqtt_init()
	if err != nil {
		fmt.Println("MQTT 初始化失败")
	}

	go func() {

		for {
			fmt.Println("i am OK")
			time.Sleep(1 * time.Second)
		}
	}()

}
