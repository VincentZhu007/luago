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
