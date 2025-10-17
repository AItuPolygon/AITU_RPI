package app

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"go.bug.st/serial"
)

type ModbusRegData struct {
	Addr, Reg, Value uint8
}

type ModBusReg struct {
	Addr, Reg uint8
}

// TODO: move to app
var modbusMu sync.Mutex

func crc(data []byte) uint16 {
	var crc16 uint16 = 0xffff
	l := len(data)
	for i := 0; i < l; i++ {
		crc16 ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if crc16&0x0001 > 0 {
				crc16 = (crc16 >> 1) ^ 0xA001
			} else {
				crc16 >>= 1
			}
		}
	}
	return crc16
}

func (app *App) InitModbus() (err error) {
	app.Port = app.Config.Modbus
	return
}
func (app *App) sendPacket(addr uint8, fn uint8, reg uint8, value uint8) (rAddr uint8, packet []byte, err error) {
	modbusMu.Lock()
	defer modbusMu.Unlock()

	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open(app.Port, mode)
	if err != nil {
		return
	}
	defer port.Close()

	packet = make([]byte, 0, 16)
	packet = append(packet, addr, fn, reg)
	if fn == 1 {
		packet = append(packet, value)
	}

	packet_crc := crc(packet)
	log.Printf("Data for CRC: %v, Calculated CRC: %04x\n", packet, packet_crc)
	packet = append(packet, byte(packet_crc>>8), byte(packet_crc&0xff))
	log.Printf("MODBUS REQ: %s\n", hex.EncodeToString(packet))
	_, err = port.Write(packet)
	if err != nil {
		return
	}
	time.Sleep(time.Millisecond * 50)

	response := make([]byte, 16)
	port.SetReadTimeout(time.Second * 1)
	l, err := port.Read(response)
	if err != nil {
		return
	}
	log.Printf("MODBUS RES: %s\n", hex.EncodeToString(response[:l]))
	if l < 6 {
		err = fmt.Errorf("modbus: answer packet too small: %d", l)
		return
	}
	packet_crc = crc(response[:l-2])
	if byte(packet_crc&0xff) != response[l-1] || byte(packet_crc>>8) != response[l-2] {
		err = errors.New("modbus: CRC check failed")
		return
	}
	rAddr = response[0]
	packet = response[1 : l-2]
	return
}

func (app *App) setRegister(data *ModbusRegData) (err error) {
	if app.Port != "" {
		_, _, err = app.sendPacket(data.Addr, 1, data.Reg, data.Value)
	} else {
		return errors.New("modbus: not ready")
	}
	log.Println("Register is set. data:", data)
	return
}

func (app *App) getRegisterValue(data *ModBusReg) (uint8, error) {
	_, response, err := app.sendPacket(data.Addr, 2, data.Reg, 0)
	if err != nil {
		log.Println("Modbus error")
		return 0, err
	}

	if len(response) != 3 || response[0] != 2 || response[1] != data.Reg {
		return 0, errors.New("error: modbus unexpected response")
	}

	log.Println("binary response: ", strconv.FormatInt(int64(response[2]), 2))
	log.Println("response:", response)
	return response[2], nil
}
