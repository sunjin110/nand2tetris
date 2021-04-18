
// function
(F)

// local 0を0で初期化
@0
D=A
@LCL // 一時的にポジションを辺こする
M=D+M
A=M
M=0 // そのポジションを0で初期化する

// LCLのポジションを戻す
@0
D=A
@LCL
M=M-D


====
@%d
D=A
@LCL
M=D+M
A=M
M=0
@%d
D=A
@LCL
M=M-D



=====

// FRAME = R13
// RET = R14

// return
@LCL
D=M
@FRAME // 一時的に保存
M=D

// リターンアドレスを取得する
@5
D=A
@FRAME
A=M-D
D=M
@RET
M=D

// 戻り値を戻ったときの、SPの先頭に来るようにする
@SP
A=M-1
D=M
@ARG
A=M
M=D


@ARG // SPの場所を返す
D=M+1
@SP
M=D

// THATの場所を戻す

@1
D=A
@FRAME
A=M-D
D=M
@THAT
M=D

// THISの場所を戻す
@2
D=A
@FRAME
A=M-D
D=M
@THIS
M=D

// ARGの場所を戻す
@3
D=A
@FRAME
A=M-D
D=M
@ARG
M=D

// LCLの場所を戻す
@4
D=A
@FRAME
A=M-D
D=M
@LCL
M=D

// 呼び出し元に戻る
@RET
A=M
0;JMP


========


@LCL
D=M
@R13
M=D
@5
D=A
@R13
A=M-D
D=M
@R14
M=D
@ARG
D=M+1
@SP
M=D
@1
D=A
@R13
A=M-D
D=M
@THAT
M=D
@2
D=A
@R13
A=M-D
D=M
@THIS
M=D
@3
D=A
@R13
A=M-D
D=M
@ARG
M=D
@4
D=A
@R13
A=M-D
D=M
@LCL
M=D
@R14
0;JMP



===


@LCL
D=M
@R13
M=D
@R13
D=M
@5
A=D-A
D=M
@R14
M=D
@SP
A=M-1
D=M
@ARG
A=M
M=D
@ARG
D=M+1
@SP
M=D
@R13
A=M-1
D=M
@THAT
M=D
@R13
D=M
@2
A=D-A
D=M
@THIS
M=D
@R13
D=M
@3
A=D-A
D=M
@ARG
M=D
@R13
D=M
@4
A=D-A
D=M
@LCL
M=D
@R14
A=M
0;JMP


===

@LCL
D=M
@R13
M=D
@5
D=A
@R13
A=M-D
D=M
@R14
M=D
@SP
A=M-1
D=M
@ARG
A=M
M=D
@ARG
D=M+1
@SP
M=D
@1
D=A
@R13
A=M-D
D=M
@THAT
M=D
@2
D=A
@R13
A=M-D
D=M
@THIS
M=D
@3
D=A
@R13
A=M-D
D=M
@ARG
M=D
@4
D=A
@R13
A=M-D
D=M
@LCL
M=D
@R14
A=M
0;JMP
