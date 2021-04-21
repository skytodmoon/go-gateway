package mqtt_mod

import (
	"fmt"
	"log"
	"main/hwconfig_json"
	"time"

	mqtt_v3 "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt_v3.MessageHandler = func(client mqtt_v3.Client, msg mqtt_v3.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt_v3.OnConnectHandler = func(client mqtt_v3.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt_v3.ConnectionLostHandler = func(client mqtt_v3.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func mqtt_sub(client mqtt_v3.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s \r\n", topic)
}

func mqtt_pub(client mqtt_v3.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

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

		// 配置 MQTT 连接参数
		opts := mqtt_v3.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqtt_config.Mqtt_broker, mqtt_config.Mqtt_port))
		opts.SetClientID(mqtt_config.Mqtt_client_id)
		opts.SetUsername(mqtt_config.Mqtt_user_name)
		opts.SetPassword(mqtt_config.Mqtt_user_pass)
		opts.SetDefaultPublishHandler(messagePubHandler)
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler

		// 判断是否用 TLS 加密传输
		if mqtt_config.Mqtt_use_tls != true {
			// 采用非 TLS 加密传输
			client := mqtt_v3.NewClient(opts)
			if token := client.Connect(); token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
			}

			// 启动订阅
			mqtt_sub(client, mqtt_config.Mqtt_topic)

		} else {
			// 采用 TLS 加密传输
		}

	} else if mqtt_config.Mqtt_version == "V5" {
		fmt.Println("MQTT 使用 V5 版本")

		// 配置 MQTT 连接参数
		fmt.Println("TBD...")

	} else {
		fmt.Println("MQTT:", mqtt_config.Mqtt_version, "版本未知，请检查！")
		return fmt.Errorf("mqtt_config.Mqtt_version")
	}

	fmt.Println("MQTT 初始化成功")
	return nil
}
