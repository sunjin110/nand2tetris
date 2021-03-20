// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(LOOP)
    @KBD
    D=M

    // 0以上なら黒く塗りつぶす
    @ON
    D;JGT

    // 0なら白
    @OFF
    0;JMP

(ON)
    // drow valueを追加
    @R0
    M=-1

    // Draw
    @DRAW
    0;JMP

(OFF)
    @R0
    M=0

    @DRAW
    0;JMP

(DRAW)
    // 実行する必要のあるblits数
    @8191
    D=A
    @R1
    M=D

    (NEXT)
        @R1
        D=M
        @pos
        M=D
        @SCREEN
        D=A
        @pos
        M=M+D

        @R0
        D=M
        @pos
        A=M
        M=D

        @R1
        D=M-1
        M=D

        @NEXT
        D;JGE

// 描画完了したらLoopに戻す
@LOOP
0;JMP
