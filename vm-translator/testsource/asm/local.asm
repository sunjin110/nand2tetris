// SPの値をglobal stackのポジションのする
@10 // 256
D=A
@SP
M=D

// Localのポジションを設定する
@20 // 3000
D=A
@LCL
M=D

// とりあえずこれを再現する
// push constant 10
// pop local 0

// push constant 10
@3
D=A
@SP
A=M
M=D
@SP
M=M+1




// pop local 5
@5
D=A
@LCL // 一時的にポジションを変更する
M=D+M
@SP
M=M-1
A=M
D=M
// ポジションに追加する
@LCL
A=M
M=D
// LCLのポジションをもとに戻す
@5
D=A
@LCL 
M=M-D

// push local 5
@5
D=A
@LCL // 一時的にポジションを変更する
M=D+M
A=M
D=M // そのポジションにあるデータを取得する

@SP
A=M // そのポジションに移動
M=D // そのポジションに取り出したLCLのデータを追加

@SP // カウント+1
M=M+1

// LCLのポジションを戻す
@5
D=A

@LCL
M=M-D


===
@5
D=A
@LCL
M=D+M
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@LCL
M=M-D
