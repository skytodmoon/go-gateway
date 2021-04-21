package mqtt_mod

import (
	"fmt"
	"log"
	"main/hwconfig_json"
)

func Mqtt_init() error {

	// 读取MQTT参数配置
	mqtt_config := hwconfig_json.Mqtt_config{}

	err := hwconfig_json.Hw_read_json(&mqtt_config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// MQTT 版本判断
	if mqtt_config.Mqtt_version == "V3" {
		fmt.Println("MQTT 使用 V3 版本")
	} else if mqtt_config.Mqtt_version == "V5" {
		fmt.Println("MQTT 使用 V5 版本")
	} else {
		fmt.Println("MQTT 版本未知，请检查！")
		return fmt.Errorf("unknown mqtt version")
	}

	fmt.Println("MQTT 初始化成功")
	return nil
}
