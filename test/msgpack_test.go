package test

import (
	"bytes"
	"fmt"
	"github.com/byorty/hardcore/decoder"
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/test/exporters"
	"github.com/byorty/hardcore/test/importers"
	"github.com/byorty/hardcore/test/models"
	msgpack "github.com/msgpack/msgpack-go"
	"testing"
)

func TestMsgpack(t *testing.T) {
	buf := new(bytes.Buffer)

	fmt.Println("msgNil     byte = ", 0xc0)
	fmt.Println("msgFalse   byte = ", 0xc2)
	fmt.Println("msgTrue    byte = ", 0xc3)
	fmt.Println("msgFloat32 byte = ", 0xca)
	fmt.Println("msgFloat64 byte = ", 0xcb)
	fmt.Println("msgUint8   byte = ", 0xcc)
	fmt.Println("msgUint16  byte = ", 0xcd)
	fmt.Println("msgUint32  byte = ", 0xce)
	fmt.Println("msgUint64  byte = ", 0xcf)
	fmt.Println("msgInt8    byte = ", 0xd0)
	fmt.Println("msgInt16   byte = ", 0xd1)
	fmt.Println("msgInt32   byte = ", 0xd2)
	fmt.Println("msgInt64   byte = ", 0xd3)
	fmt.Println("msgStr8    byte = ", 0xd9)
	fmt.Println("msgStr16   byte = ", 0xda)
	fmt.Println("msgStr32   byte = ", 0xdb)
	fmt.Println("msgBin8    byte = ", 0xc4)
	fmt.Println("msgBin16   byte = ", 0xc5)
	fmt.Println("msgBin32   byte = ", 0xc6)
	fmt.Println("msgArray16 byte = ", 0xdc)
	fmt.Println("msgArray32 byte = ", 0xdd)
	fmt.Println("msgMap16   byte = ", 0xde)
	fmt.Println("msgMap32   byte = ", 0xdf)
	fmt.Println("msgPositiveFixInt byte = ", 0x00)
	fmt.Println("msgFixMap         byte = ", 0x80)
	fmt.Println("msgFixArray       byte = ", 0x90)
	fmt.Println("msgFixRaw         byte = ", 0xa0)
	fmt.Println("msgNegativeFixInt byte = ", 0xe0)

	fmt.Println("FIXMAP = ", msgpack.FIXMAP|byte(2))
	fmt.Println("MAP16 = ", msgpack.MAP16)
	fmt.Println("MAP32 = ", msgpack.MAP32)

	email := "user@example.com"
	//role := models.LoggedUserRole
	userMap := map[string]interface{}{
		"id":    int64(1),
		"email": email,
		//"role": role.GetId(),
	}
	msgpack.Pack(buf, userMap)
	msgBuf := buf.Bytes()
	fmt.Println("m:", msgBuf)

	user := new(models.User).
		SetId(1).
		SetEmail(email)
		//SetRole(role)
		//SetRegisterDate(now)
	encoderBuf := encoder.NewMsgpack().Encode(exporters.NewUser(user))
	fmt.Println("e:", encoderBuf)

	//r := bytes.NewReader(msgBuf)
	//v, _, err := msgpack.Unpack(r)
	//fmt.Println("unpack m:", v, err)
	r1 := bytes.NewReader(encoderBuf)
	v1, _, err1 := msgpack.Unpack(r1)
	fmt.Println("unpack e:", v1, err1)

	user1 := new(models.User)
	decoder.NewMsgpack(encoderBuf).Decode(importers.NewUser(user1))
	fmt.Println(user1)

	t.Fail()
}
