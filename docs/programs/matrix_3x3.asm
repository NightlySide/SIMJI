; program to compute the square of a 3x3 matrix a

    add r0,3,r10   ; N=3 in R10

; precompute the RESULT offset in memory
    mult r10,r10,r11 ; OFFSET in R11 for storing result

; start of the multiplication...
    add  r0,0,r1   ; i=0

L5: slt  r1,r10,r2   ; i < N
    braz r2, Label_end
    add  r0,0,r2   ; j=0

L4: slt  r2,r10,r8   ; j < N
    braz r8, L1    ; no?  goto L1
    add  r0,0,r4   ; yes? s=0
    add  r0,0,r5   ; k=0

L2: slt  r5,r10,r6   ; k < 2
    braz r6,L3
    ; computation for adresse [i,k]
    mult  r1,r10,r6   ; r6 as tmp
    add  r6,r5,r6  ; @[i,k] -> r6
    ; computation for adresse @[k,j]
    mult  r5,r10,r7
    add  r7,r2,r7  ; @[k,j] -> r7
    load r6,r0,r8  ; a[i,k]
    load r7,r0,r9  ; b[k,j]
    mult  r8,r9,r8  ; tmp=a[i,k]*b[k,j]
    add  r4,r8,r4  ; s+=tmp
    add  r5,1,r5   ; k++
    jmp  L2,r0
L3:
    mult r1,r10,r6
    add r6,r2,r6   ; @[i,j] for C
    add r6,r11,r6  ; @[i,j]+OFFSET_C
    store r6,0,r4  ; c[i,j]=s
    add r2,1,r2    ; j++
    jmp L4,r0

L1: add r1,1,r1    ; i++
    jmp L5,r0

Label_end:
    stop
