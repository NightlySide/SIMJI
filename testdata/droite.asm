; r2 = A, r3 = B, r4 = N, r1 = valeur actuelle
; on demande A
scall 0
add r1, r0, r2
; on demande B
scall 0
add r1, r0, r3
; on demande le nombre de points
scall 0
add r1, r0, r4
; on met la valeur actuelle à B
xor r1, r1, r1
add r1, r3, r1
; boucle
BEGIN_LOOP:
    ; on modifie le compteur
    sub r4, 1, r4
    ; on affiche le résultat
    scall 1
    ; tant que N n'est pas 0
    braz r4, END_LOOP
    ; on modifie la valeur actuelle
    add r1, r2, r1
    ; on revient au debut de la boucle
    jmp BEGIN_LOOP, r0
END_LOOP:
stop