package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"main/modbus_mod"
)

var TOPIC = make(map[string]string)

// set timestamp, clientid, subscribe topic and publish topic
// var timeStamp string = "1528018257135"
// var clientId string = "192.168.****"
// var subTopic string = "/" + productKey + "/" + deviceName + "/user/get"
// var pubTopic string = "/" + productKey + "/" + deviceName + "/user/update"

// 设置MQTT登录信息验证结构体
type AuthInfo struct {
	password, username, mqttClientId string
}

// 构造MQTT登录加密序列
func calculate_sign(clientId, productKey, deviceName, deviceSecret, timeStamp string) AuthInfo {
	var raw_passwd bytes.Buffer
	raw_passwd.WriteString("clientId" + clientId)
	raw_passwd.WriteString("deviceName")
	raw_passwd.WriteString(deviceName)
	raw_passwd.WriteString("productKey")
	raw_passwd.WriteString(productKey)
	raw_passwd.WriteString("timestamp")
	raw_passwd.WriteString(timeStamp)
	fmt.Println(raw_passwd.String())

	// 使用SHA1的加密
	mac := hmac.New(sha1.New, []byte(deviceSecret))
	mac.Write([]byte(raw_passwd.String()))
	password := fmt.Sprintf("%02x", mac.Sum(nil))
	fmt.Println(password)
	username := deviceName + "&" + productKey

	var MQTTClientId bytes.Buffer
	MQTTClientId.WriteString(clientId)
	// hmac, use sha1; securemode=2 means TLS connection
	MQTTClientId.WriteString("|securemode=2,_v=paho-go-1.0.0,signmethod=hmacsha1,timestamp=")
	MQTTClientId.WriteString(timeStamp)
	MQTTClientId.WriteString("|")

	auth := AuthInfo{password: password, username: username, mqttClientId: MQTTClientId.String()}
	return auth
}

func show_logo() {
	logo_buf, err := ioutil.ReadFile("logo.log")
	if err != nil {
		fmt.Println("logo file doesn't exist....check it?")
	}
	fmt.Println(string(logo_buf))
}

// 加载默认配置文件
func load_default_config() {
	file, err := os.Create("config/config.ini")

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadFile("config/config_default.ini")

	file.WriteString(string(buf))
	fmt.Println(string(buf))
	file.Close()
}

func main() {
	// 展示公司logo
	show_logo()

	err := modbus_mod.Modbus_init()
	if err != nil {
		fmt.Println("mb init error~")
	}

	go func() {

		for {
			fmt.Println("i am OK")
			time.Sleep(1 * time.Second)
		}
	}()

}
