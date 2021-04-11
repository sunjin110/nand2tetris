// global stack 256から
// SPは、256にある

// push constant 1
// push constant 2
// push constant 3
// add
// add
// init 
// @256
// D=A

// @SP
// A=D

// SPの値をglobal stackのポジションのする
@20
D=A
@SP
M=D

// push constant 7 : address: 256に7を追加する
@7
D=A

@SP
A=M
M=D

// SPを+1する
@SP
M=M+1

// push constant 8 address: 267に8を追加する
@8
D=A

@SP
A=M
M=D

@SP
M=M+1

// add
@SP
A=M-1
D=M
M=0
A=A-1
M=D+M

// SPを-1にする
@SP
M=M-1
