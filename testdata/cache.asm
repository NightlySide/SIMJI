; store 5 dans le cache
add r0, 5, r5
store r0, 10, r5
; load la valeur dans r1
load r0, 10, r1
; devrait afficher 5
scall 1
stop