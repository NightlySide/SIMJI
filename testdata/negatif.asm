; on init la valeur de début
add r0, 0, r1
; on affiche le 0
scall 1
; on ajoute un nombre positif
add r1, 20, r1
; on affiche le res
scall 1
; on ajoute un nb négatif
add r1, -5, r1
; on affiche
scall 1
; on reset r1
sub r1, r1, r1
; doit afficher 0
add r0, -10, r1
; doit afficher -10
scall 1
stop