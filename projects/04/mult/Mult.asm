// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

// Start R2 at 0.

// R0を2
// R1を3と仮定する

//  R2に0を追加
@R2
M=0

// R0のデータをDに追加する
@R0
D=M

// Dが0以上であればSTEPにJMPする
@STEP
D;JGT

@END
0;JMP

(STEP)
    // R2の値をDに展開
    @R2
    D=M

    // R1
    // 0+3
    // D=3
    @R1
    D=D+M

    // R2に3を追加
    @R2
    M=D

    // R0を-1する 
    // R0=1
    @R0
    D=M-1
    M=D

    // R0が0以上の場合はSTEPに戻る
    @STEP
    D;JGT

(END)
    @END
    0;JMP
