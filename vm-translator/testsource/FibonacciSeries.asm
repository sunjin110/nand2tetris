@256
D=A
@SP
M=D
@300
D=A
@LCL
M=D
@400
D=A
@ARG
M=D
@1
D=A
@ARG
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@ARG
M=M-D
@SP
M=M-1
A=M
D=M
@THAT
M=D
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@THAT
M=D+M
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@0
D=A
@THAT
M=M-D
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@THAT
M=D+M
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@1
D=A
@THAT
M=M-D
@0
D=A
@ARG
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@ARG
M=M-D
@2
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
M=M-D
@SP
M=M-1
@0
D=A
@ARG
M=D+M
@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D
@0
D=A
@ARG
M=M-D
(MAIN_LOOP_START)
@0
D=A
@ARG
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@ARG
M=M-D
@SP
A=M-1
D=M
@SP
M=M-1
@COMPUTE_ELEMENT
D;JNE
@END_PROGRAM
0;JMP
(COMPUTE_ELEMENT)
@0
D=A
@THAT
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@THAT
M=M-D
@1
D=A
@THAT
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@THAT
M=M-D
@SP
A=M-1
D=M
M=0
A=A-1
M=D+M
@SP
M=M-1
@2
D=A
@THAT
M=D+M
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@2
D=A
@THAT
M=M-D
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
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
M=D+M
@SP
M=M-1
@SP
M=M-1
A=M
D=M
@THAT
M=D
@0
D=A
@ARG
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@ARG
M=M-D
@1
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
M=M-D
@SP
M=M-1
@0
D=A
@ARG
M=D+M
@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D
@0
D=A
@ARG
M=M-D
@MAIN_LOOP_START
0;JMP
(END_PROGRAM)
