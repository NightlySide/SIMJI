STARTING_LABEL:
; adding
add r1, r0, r2
add r1, 5,  r1
; substracting
sub r1, r0, r2
sub r1, 3,  r2
; multiplication
mult r1, r1, r2
mult r1, 5 , r2
; division
div r1, r1, r2
div r1, 5 , r2
; logical operations
and r1, 1,  r1
and r1, r2, r2
or  r1, 1,  r2
or  r1, r2, r2
xor r1, 1,  r2
xor r1, r2, r2
; shifting
shl r1, 2, r2
shr r2, 2, r2
; comparison
slt r1, 5, r2
sle r1, 6, r2
seq r1, r2, r2
; memory
load r1, r0, r2
store r2, r0, r1
; jmps
jmp ENDING_LABEL, r0
braz r2, STARTING_LABEL
branz r0, STARTING_LABEL
; system
scall 1
ENDING_LABEL:
; stop the program
stop