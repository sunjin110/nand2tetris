@20
D=A
@SP
M=D

@11
D=A
@SP
A=M
M=D
@SP
M=M+1

@11
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
D=M-D // x-y
M=-1 // 先に-1を入れる
@LT
D;JLT
@SP
A=M-1
A=A-1
M=0
(LT)
@SP
M=M-1


===
@SP
A=M-1
D=M
M=0
A=A-1
D=M-D
M=-1
@LT_%s_%d
D;JLT
@SP
A=M-1
A=A-1
M=0
(LT_%s_%d)
@SP
M=M-1
