@20
D=A
@SP
M=D

@7
D=A
@SP
A=M
M=D
@SP
M=M+1

@8
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
A=M-1
D=M
M=0
A=A-1
D=D-M // x-y
M=-1 // 先に-1を入れる
@EQ
D;JEQ
@SP
A=M-1
A=A-1
M=0
(EQ)
@SP
M=M-1



