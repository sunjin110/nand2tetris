

// pop temp 3
@SP
M=M-1
A=M
D=M

@8 // 5 + 3
M=D


// push temp 3
@8 // 5 + 3
D=M

@SP
A=M
M=D

@SP
M=M+1

===
@%d
D=M
@SP
A=M
M=D
@SP
M=M+1