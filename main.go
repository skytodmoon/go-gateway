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

	serial "github.com/tarm/goserial"

	"gopkg.in/ini.v1"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var TOPIC = make(map[string]string)

// set timestamp, clientid, subscribe topic and publish topic
var timeStamp string = "1528018257135"
var clientId string = "192.168.****"
var subTopic string = "/" + productKey + "/" + deviceName + "/user/get"
var pubTopic string = "/" + productKey + "/" + deviceName + "/user/update"

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

func creat_default_config() {
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

	// 查找并打开配置文件
	_, err := os.Open("config.ini")
	if err != nil {
		fmt.Println("Can't find config.ini, will create a default one")
		creat_default_config()
	}

	// 打开配置文件
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	clientId := cfg.Section(("MQTT")).Key("server_addr").String()
	productKey := cfg.Section("MQTT").Key("product_key").String()
	deviceName := cfg.Section("MQTT").Key("device_name").String()
	deviceSecret := cfg.Section("MQTT").Key("device_secret").String()
	timeStamp := cfg.Section("MQTT").Key("mqtt_timestamp").String()

	// 设置登录代理服务器地址
	var raw_broker bytes.Buffer
	raw_broker.WriteString("tls://")
	raw_broker.WriteString(productKey)
	raw_broker.WriteString(".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883")
	opts := MQTT.NewClientOptions().AddBroker(raw_broker.String())

	// 计算出加密后的登录信息
	auth := calculate_sign(clientId, productKey, deviceName, deviceSecret, timeStamp)
	opts.SetClientID(auth.mqttClientId)
	opts.SetUsername(auth.username)
	opts.SetPassword(auth.password)
	opts.SetKeepAlive(60 * 2 * time.Second)
	//	opts.SetDefaultPublishHandler(f)

	serial_port := cfg.Section("COM").Key("rs485_port").String()
	serial_baud := cfg.Section("COM").Key("rs485_baud").MustInt()
	fmt.Println("RS485 info:", serial_port, "baudrate is :", serial_baud)

	c := &serial.Config{Name: serial_port, Baud: serial_baud}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	serial_cmd1 := cfg.Section("CMD").Key("cmd").String()
	fmt.Println("the sting will send", serial_cmd1)
	s.Write(([]byte(serial_cmd1)))

	serial_cmd2 := cfg.Section("CMD").Key("cmd2").String()
	if serial_cmd2 == "" {
		fmt.Println("nothing to be show in cmd2")
	} else {
		fmt.Println("there is sth inside")
	}

	// //获取配置文件中的配置项
	// id, err := cfg.String("COM", "COMID")
	// //设置串口编号
	// c := &serial.Config{Name: id, Baud: 115200}
	// //打开串口
	// s, err := serial.OpenPort(c)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// command, err := cfg.String("COM", "COMMAND")
	// // 写入货柜串口命令go
	// log.Printf("货柜打开指令 %s", command)
	// n, err := s.Write([]byte(command))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// buf := make([]byte, 128)
	// n, err = s.Read(buf)
	// log.Printf("读取窗口信息 %s", buf[:n])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
}
