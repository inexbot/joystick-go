package socket

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"net"
	"time"
)

var socket net.Conn

// StartSocket 开始连接伺服
func StartSocket() {
	var err interface{}
	socket, err = net.Dial("tcp", "192.168.0.63:9000")
	for err != nil {
		println(err)
		time.Sleep(time.Second * 2)
		StartSocket()
		return
	}
}

// CloseSocket 关闭socket
func CloseSocket() {
	socket.Close()
}

// SendSocketMessage 发送socket消息
func SendSocketMessage(data map[string]interface{}, command int) {
	tmpToSend := make([]byte, 0)
	tmpToSend = append(tmpToSend, 0x4E)
	tmpToSend = append(tmpToSend, 0x66)
	dataBytes, _ := json.Marshal(data)
	dataLen := len(dataBytes)
	_dd := int(dataLen / 256)
	_dd1 := int(dataLen % 256)
	tmpToSend = append(tmpToSend, byte(_dd))
	tmpToSend = append(tmpToSend, byte(_dd1))
	_co1 := int(command / 256)
	_co2 := int(command % 256)
	tmpToSend = append(tmpToSend, byte(_co1))
	tmpToSend = append(tmpToSend, byte(_co2))
	tmpToSend = append(tmpToSend, dataBytes...)
	_Crc := crc32.ChecksumIEEE(tmpToSend[2:])
	dataCrc := bytes.NewBuffer([]byte{})
	binary.Write(dataCrc, binary.BigEndian, _Crc)
	_dataCrc := dataCrc.Bytes()
	tmpToSend = append(tmpToSend, _dataCrc...)
	socket.Write(tmpToSend)
	fmt.Println("send", data)
	return
}
