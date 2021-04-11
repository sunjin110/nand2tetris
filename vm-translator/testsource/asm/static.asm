// push static 3

@Xxx.3
D=M

@SP
A=M
M=D
@SP
M=M+1

// pop static 3
@SP
M=M-1
A=M
D=M

@Xxx.3
M=D

====
@Static_%s_%d
D=M
@SP
A=M
M=D
@SP
M=M+1

===
@SP
M=M-1
A=M
D=M
@Static_%s_%d
M=D