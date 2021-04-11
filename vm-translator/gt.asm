@20
D=A
@SP
M=D

@9
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
D=M-D // x-y
M=-1 // 先に-1を入れる
@GT
D;JGT
@SP
A=M-1
A=A-1
M=0
(GT)
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
@GT_%s_%d
D;JGT
@SP
A=M-1
A=A-1
M=0
(GT_%s_%d)
@SP
M=M-1
