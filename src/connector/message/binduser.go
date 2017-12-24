package message

import (
	"falcon/src/connector/connection"
	"falcon/src/connector/protocol"
	"bytes"
)

type BindUserMessage struct {
	ByteBufMessage
	UserId string
	Tags   string
	Data   string
}

//创建消息
func NewBindUserMessage(conn *connection.Conn) *BindUserMessage {

	msg := &BindUserMessage{}

	packet := protocol.NewPacket(protocol.CMD_BIND)
	packet.SessionId = msg.GenSessionId()

	msg.packet = packet
	msg.conn = conn
	msg.BaseMessage.Child = msg
	
	return msg
}

//解码消息
func (me *BindUserMessage) Decode(body []byte) {
	//输出消息看内容
	buf := new(bytes.Buffer)
	buf.Write(body)

	me.UserId = me.DecodeString(buf)
	me.Data = me.DecodeString(buf)
	me.Tags = me.DecodeString(buf)
}

//编码消息
func (me *BindUserMessage) Encode() []byte {

	buf := new(bytes.Buffer)

	me.EncodeString(buf, me.UserId)
	me.EncodeString(buf, me.Data)
	me.EncodeString(buf, me.Tags)

	return buf.Bytes()
}
