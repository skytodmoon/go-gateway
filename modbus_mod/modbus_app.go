package modbus_mod

import (
	"fmt"
	"log"
	"main/hwconfig_json"
	"time"

	"github.com/goburrow/serial"
	modbus "github.com/thinkgos/gomodbus/v2"
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

	// 调试口打印 json参数
	// fmt.Println(mb_config)
	// fmt.Println(rs232_config)
	// fmt.Println(rs485_1_config)
	// fmt.Println(rs485_2_config)

	// 判断 Modbus 功能是否启用
	if mb_config.Modbus_enable == true {
		fmt.Println("Modbus 功能启用")

		// 判断 Modbus 工作模式
		if mb_config.Modbus_mode == "tcp" {
			fmt.Println("Modbus 采用 TCP 协议通信")

			// 读取 TCP 配置服务器地址及端口
			tcp_server := mb_config.Modbus_tcp_addr + ":" + mb_config.Modbus_tcp_port
			fmt.Println(tcp_server)

			// 开始启动 Modbus 的 TCP模式进行连接
			p := modbus.NewTCPClientProvider(tcp_server, modbus.WithEnableLogger())
			client := modbus.NewClient(p)
			err := client.Connect()
			if err != nil {
				fmt.Println("连接失败，", err)
				return err
			}
			defer client.Close()

			fmt.Println("连接成功，开始工作...")
			for {
				results, err := client.ReadHoldingRegisters(1, 1, 10)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Printf("ReadDiscreteInputs %#v\r\n", results)

				time.Sleep(time.Second * 2)
			}

		} else if mb_config.Modbus_mode == "rtu" {
			fmt.Println("Modbus 采用 RTU 协议通信")

			// 开始启动 Modbus 的 RTU 模式进行连接
			p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
				modbus.WithSerialConfig(serial.Config{
					Address:  rs485_1_config.Com_port,
					BaudRate: rs485_1_config.Com_baudrate,
					DataBits: rs485_1_config.Com_databits,
					StopBits: rs485_1_config.Com_stopbits,
					Parity:   rs485_1_config.Com_parity,
					Timeout:  modbus.SerialDefaultTimeout,
					RS485: serial.RS485Config{
						Enabled:            false,
						DelayRtsBeforeSend: 0,
						DelayRtsAfterSend:  0,
						RtsHighDuringSend:  false,
						RtsHighAfterSend:   false,
						RxDuringTx:         false,
					},
				}))

			client := modbus.NewClient(p)
			err := client.Connect()
			if err != nil {
				fmt.Println("连接失败，", err)
				log.Fatal(err)
			}
			defer client.Close()

			fmt.Println("连接成功，开始工作...")
			for {
				results, err := client.ReadHoldingRegisters(mb_config.Modbus_slave_id_1, 1, 1)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Printf("ReadDiscreteInputs %#v\r\n", results)

				time.Sleep(time.Second * 2)
			}
		} else if mb_config.Modbus_mode == "ascii" {
			fmt.Println("Modbus 采用 ASCII 协议通信")

			// 开始启动 Modbus 的 ASCII 模式进行连接
			p := modbus.NewASCIIClientProvider(modbus.WithEnableLogger(),
				modbus.WithSerialConfig(serial.Config{
					Address:  rs485_2_config.Com_port,
					BaudRate: rs485_2_config.Com_baudrate,
					DataBits: rs485_2_config.Com_databits,
					StopBits: rs485_2_config.Com_stopbits,
					Parity:   rs485_2_config.Com_parity,
					Timeout:  modbus.SerialDefaultTimeout,
					RS485: serial.RS485Config{
						Enabled:            false,
						DelayRtsBeforeSend: 0,
						DelayRtsAfterSend:  0,
						RtsHighDuringSend:  false,
						RtsHighAfterSend:   false,
						RxDuringTx:         false,
					},
				}))

			client := modbus.NewClient(p)
			err := client.Connect()
			if err != nil {
				fmt.Println("连接失败，", err)
				log.Fatal(err)
			}
			defer client.Close()

			fmt.Println("连接成功，开始工作...")

			for {
				results, err := client.ReadHoldingRegisters(1, 1, 10)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Printf("ReadDiscreteInputs %#v\r\n", results)

				time.Sleep(time.Second * 2)
			}
		} else {
			fmt.Println("Modbus 模式:", mb_config.Modbus_mode, "解析错误！请检查 json 配置！")
		}
	} else {
		fmt.Println("Modbus 功能未启用")
	}

	fmt.Println("Modbus 初始化成功")
	return nil
}
