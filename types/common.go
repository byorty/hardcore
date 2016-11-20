package types

const (
	MsgpackNil            byte = 0xc0
	MsgpackFalse          byte = 0xc2
	MsgpackTrue           byte = 0xc3
	MsgpackFloat32        byte = 0xca
	MsgpackFloat64        byte = 0xcb
	MsgpackUint8          byte = 0xcc
	MsgpackUint16         byte = 0xcd
	MsgpackUint32         byte = 0xce
	MsgpackUint64         byte = 0xcf
	MsgpackInt8           byte = 0xd0
	MsgpackInt16          byte = 0xd1
	MsgpackInt32          byte = 0xd2
	MsgpackInt64          byte = 0xd3
	MsgpackStr8           byte = 0xd9
	MsgpackStr16          byte = 0xda
	MsgpackStr32          byte = 0xdb
	MsgpackBin8           byte = 0xc4
	MsgpackBin16          byte = 0xc5
	MsgpackBin32          byte = 0xc6
	MsgpackArray16        byte = 0xdc
	MsgpackArray32        byte = 0xdd
	MsgpackMap16          byte = 0xde
	MsgpackMap32          byte = 0xdf
	MsgpackPositiveFixInt byte = 0x00
	MsgpackFixMap         byte = 0x80
	MsgpackFixArray       byte = 0x90
	MsgpackFixRaw         byte = 0xa0
	MsgpackNegativeFixInt byte = 0xe0
	MsgpackMaxFixMap           = 16
	MsgpackMaxFixAarray        = 16
	MsgpackMaxFixRaw           = 32
	MsgpackMax16Bit            = 2 << (16 - 1)
)
