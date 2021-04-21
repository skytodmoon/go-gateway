package modbus_mod

import (
	"log"
	"main/hwconfig_json"
	"time"

	"github.com/goburrow/serial"
	"github.com/tbrandon/mbserver"
)

func Modbus_server_init() error {
	rs232_config := hwconfig_json.Com_rs232{}
	err := hwconfig_json.Hw_read_json(&rs232_config)

	serv := mbserver.NewServer() // argument for constructor determines the ModbusSlave-ID

	err = serv.ListenRTU(&serial.Config{
		Address:  rs232_config.Com_port,
		BaudRate: rs232_config.Com_baudrate,
		DataBits: rs232_config.Com_databits,
		StopBits: rs232_config.Com_stopbits,
		Parity:   rs232_config.Com_parity,
		Timeout:  0,
		RS485:    serial.RS485Config{},
	})
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}

}
