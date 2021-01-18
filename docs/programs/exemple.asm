add r0 1 r3
L_loop:
    sle r2 1 r4
    seq r4 r0 r4
    braz r4 L_END
    mult r3 r2 r3
    sub r2 1 r2
    jmp L_loop r0
L_end:
stop