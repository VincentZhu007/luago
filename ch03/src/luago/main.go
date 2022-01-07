/**
 * main.go
 *
 * 测试解析lua指令
 */

package main

import "fmt"
import "io/ioutil"
import "os"
import "luago/binchunk"
import "luago/vm"

/*
 * 打印函数原型信息
 */
func list(f *binchunk.Prototype) {
	printHeader(f) 	// 打印函数头部
	printCode(f)	// 打印指令表
	printDetail(f)	// 打印调试信息
	for _, p := range f.Protos {
		list(p) 	// 递归调用打印子函数信息
	}
}

func printHeader(f *binchunk.Prototype) {
	// 设置函数类型
	funcType := "main"	// Lua main函数的起始行号固定是0,其它函数非0
	if f.LineDefined > 0 { funcType = "function" }

	varargFlag := ""

	if (f.IsVararg > 0) { // 可变参数，使用+标记
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType,
		f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	
	fmt.Printf("%d%s params, %d slots, %d upvalues, %d locals, %d constants, %d functions\n", 
		f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues),
		len(f.LocVars), len(f.Constants), len(f.Protos))
}

func printCode(f *binchunk.Prototype)  {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s \t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		if i.BMode() != vm.OpArgN {
			if b > 0xFF {
				fmt.Printf(" %d", -1-b&0xFF)
			} else {
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != vm.OpArgN {
			if c > 0xFF {
				fmt.Printf(" %d", -1-c&0xFF)
			} else {
				fmt.Printf(" %d", c)
			}
		}
	case vm.IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf(" %d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf(" %d", bx)
		}
	case vm.IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)
	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}
}

func printDetail(f *binchunk.Prototype) {
	// 打印常量信息
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(k))
	}
	
	// 打印局部变量信息
	fmt.Printf("locals (%d):\n", len(f.LocVars))
	for i, l := range f.LocVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i+1,
			l.VarName, l.StartPC, l.EndPC)
	}

	// 打印upvalue
	fmt.Printf("upvalues (%d):\n", len(f.Upvalues))
	for i, u := range f.Upvalues {
		fmt.Printf("\t%d\t%d\t%d\t%d\n", i, upvalueName(f, i),
			u.Instack, u.Idx)
	}
}

// 将常量转换为字符串
func constantToString(k interface{}) string {
	switch k.(type) {
	case nil:		return "nil"
	case bool:		return fmt.Sprintf("%t", k)
	case float64:	return fmt.Sprintf("%g", k)
	case int64:		return fmt.Sprintf("%d", k)
	case string:	return fmt.Sprintf("%q", k)
	default: 		return "?"
	}
}

func upvalueName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1]) // 读取文件内容
		if err != nil { panic(err) }
		proto := binchunk.Undump(data)	// 反汇编lua字节码
		list(proto)	// 打印函数原型信息
	}
}

