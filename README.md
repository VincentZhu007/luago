# luago

*用go实现的lua*

## 第1章、环境搭建

开发环境：ubuntu 20.04

### 1.1 安装go

``` bash
$ sudo apt install golang-go
```

### 1.2 安装lua

``` bash
$ sudo apt install lua5.3
```

## 第2章、二进制chunk

### 2.1 查看二进制chunk

```bash
zgd@zgd-dell:~/01-srccode/luago/ch01$ luac -o hello.out hello.lua 
zgd@zgd-dell:~/01-srccode/luago/ch01$ lua hello.out
hello, lua!
zgd@zgd-dell:~/01-srccode/luago/ch01$ luac -l hello.out

main <hello.lua:0,0> (4 instructions at 0x561287fb9c50)
0+ params, 2 slots, 1 upvalue, 0 locals, 2 constants, 0 functions
    1	[2]	GETTABUP 	0 0 -1	; _ENV "print"
    2	[2]	LOADK    	1 -2	; "hello, lua!"
    3	[2]	CALL     	0 2 1
    4	[2]	RETURN   	0 1
```

### 2.2 二进制chunk解析

#### 2.2.1 编译运行

在luago目录编译运行：

```bash
$ cd ch01
$ luac -o hello.out hello.lua       # 为ch02demo准备lua二进制chunk文件
$ cd ..
```

```bash
$ cd ch02
$ export GOPATH="${PATH}/"          # 指定ch02目录为go工作目录
$ go install lualist                # 编译lualist二进制
$ ./bin/lualist ../ch01/hello.out   # 使用demo程序解析二进制chunk
```

#### 2.2.2 输出效果

`luac -l`输出：
```bash
$ luac -l ../ch01/hello.out                           

main <hello.lua:0,0> (4 instructions at 0x7f9026406dc0)
0+ params, 2 slots, 1 upvalue, 0 locals, 2 constants, 0 functions
        1       [2]     GETTABUP        0 0 -1  ; _ENV "print"
        2       [2]     LOADK           1 -2    ; "hello, lua!"
        3       [2]     CALL            0 2 1
        4       [2]     RETURN          0 1

```

`lualist`输出：
```
$ ./bin/lualist ../ch01/hello.out

main <@hello.lua:0,0> (4 instructions)
0+ params, 2 slots, 1 upvalues, 0 locals, 2 constants, 0 functions
        1       [2]     0x00400006
        2       [2]     0x00004041
        3       [2]     0x01004024
        4       [2]     0x00800026
constants (2):
        1       "print"
        2       "hello, lua!"
locals (0):
upvalues (1):
        0       %!d(string=_ENV)        1       0
```

## 第3章、指令集

Lua虚拟机：

- 基于寄存器
- 定长指令集，单条指令占4个字节

Lua指令集编码模式：
- iABC
- iABx
- iAxBx
- iAx

编译执行：
```bash
$ cd ch03
$ export GOPATH="${PATH}/"          # 指定ch03目录为go工作目录
$ go install luago                  # 编译luago二进制
$ ./bin/luago ../ch01/hello.out

main <@hello.lua:0,0> (4 instructions)
0+ params, 2 slots, 1 upvalues, 0 locals, 2 constants, 0 functions
        1       [2]     GETTABUP        0 0 -1
        2       [2]     LOADK           1 -2
        3       [2]     CALL            0 2 1
        4       [2]     RETURN          0 1
constants (2):
        1       "print"
        2       "hello, lua!"
locals (0):
upvalues (1):
        0       %!d(string=_ENV)        1       0
```

## 第4章、Lua API


