/**
 * binary_chunk.go
 *
 * 解析二进制chunk文件
 */

package binchunk

const (
	LUA_SIGNATURE		= "\x1bLua"
	LUAC_VERSION		= 0x53
	LUAC_FORMAT			= 0
	LUAC_DATA			= "\x19\x93\r\n\x1a\n"
	CINT_SIZE			= 4
	CSZIET_SIZE			= 8
	INSTRUCTION_SIZE	= 4
	LUA_INTEGER_SIZE	= 8
	LUA_NUMBER_SIZE		= 8
	LUAC_INT			= 0x5678
	LUAC_NUM			= 370.5
)

const (
	TAG_NIL				= 0x00
	TAG_BOOLEAN			= 0x01
	TAG_NUMBER			= 0x03
	TAG_INTEGER			= 0x13
	TAG_SHORT_STR		= 0x04
	TAG_LONG_STR		= 0x14
)

type header struct {
	signature 			[4]byte			// 签名，Lua二进制魔数，用来识别文件格式
	version 			byte			// 版本号
	format				byte			// 格式号
	luacData			[6]byte			// 存放特定数字，用来校验文件
	cintSize			byte			// 整数宽度
	sizetSize			byte			// size_t宽度
	instructionSize		byte			// lua虚拟机指令宽度
	luaIntegerSize		byte			// lua整数宽度
	luaNumberSize		byte			// lua数值宽度
	luacInt				int64			// 存放整型数0x5678，用来判断大小端
	luacNum				float64			// 存放浮点数370.5，用来判断浮点数格式
}

type Upvalue struct {
	Instack byte
	Idx 	byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC 	uint32
}

type Prototype struct {
	Source				string			// 源文件名
	LineDefined			uint32			// 函数的起始行号
	LastLineDefined 	uint32			// 函数的结束行号
	NumParams			byte			// 固定参数个数
	IsVararg			byte			// 是否可变参数
	MaxStackSize		byte			// 寄存器个数
	Code				[]uint32		// Lua虚拟机指令表
	Constants			[]interface{}	// 函数中使用的常量表
	Upvalues			[]Upvalue		// Upvalue列表
	Protos				[]*Prototype		// 子函数列表
	LineInfo			[]uint32		// 行号表
	LocVars				[]LocVar		// 局部变量表
	UpvalueNames		[]string		// Upvalue名列表
}	

type binarychunk struct {
	header								// 头部
	sizeUpvalues 		byte			// 主函数Upvalue数量
	mainFunc			*Prototype		// 主函数原型
}


func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()			// 校验头部
	reader.readByte()				// 跳过Upvalue数量
	return reader.readProto("")		// 读取函数原型
}
