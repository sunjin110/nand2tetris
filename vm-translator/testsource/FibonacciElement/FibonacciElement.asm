// SPを初期化

@256
D=A
@SP
M=D

// sys.initをcallする
@Sys.init
0;JMP





// call ※引数1

// return addressの場所をpushする
@RET_ADDRESS_CALL_0
D=A
@SP
M=D
@SP
M=M+1

// LCLの場所をpushする
@LCL
D=A
@SP
M=D
@SP
M=M+1

// ARGの場所をpushする
@ARG
D=A
@SP
M=D
@SP
M=M+1

// THISの場所をpushする
@THIS
D=A
@SP
M=D
@SP
M=M+1

// THATの場所をpushする
@THAT
D=A
@SP
M=D
@SP
M=M+1

// ARGを別の場所に移す
@6 // -n-5
D=A
@SP
D=M-D
@ARG
M=D

// LCLを別の場所に移動する
@SP
D=M
@LCL
M=D

// GOTO funciton
@function.name
0;JMP



(RET_ADDRESS_CALL_0)



=====================



// call ※引数1

@%s
D=A
@SP
M=D
@SP
M=M+1
@LCL
D=A
@SP
M=D
@SP
M=M+1
@ARG
D=A
@SP
M=D
@SP
M=M+1
@THIS
D=A
@SP
M=D
@SP
M=M+1
@THAT
D=A
@SP
M=D
@SP
M=M+1
@%d
D=A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
@%s
0;JMP
(%s)






============================

// 1. push return address
@%s
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@%d
D=D+A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
@%s
0;JMP
(%s)










