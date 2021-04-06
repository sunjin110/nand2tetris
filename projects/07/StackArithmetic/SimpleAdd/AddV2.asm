// global stack 256から
// SPは、256にある

// push constant 1
// push constant 2
// push constant 3
// add
// add
// init 
// @256
// D=A

// @SP
// A=D

// SPの値をglobal stackのポジションのする
@256
D=A
@SP
M=D

// TODO 256にデータ1を追加する
@1 // 1はpush conatant 1の定数
D=A
@SP
AM=D
// TODO SPの値を1プラスする

// @2
// D=A
// @SP
// M=D

