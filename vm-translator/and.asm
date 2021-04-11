// SPの値をglobal stackのポジションのする
@20
D=A
@SP
M=D

// push constant 7 : address: 256に7を追加する
@1
D=A

@SP
A=M
M=D

// SPを+1する
@SP
M=M+1

// push constant 8 address: 267に8を追加する
@1
D=A

@SP
A=M
M=D

@SP
M=M+1

// and
@SP
A=M-1
D=M
M=0
A=A-1
M=D&M
@SP
M=M-1
