          ; initialise valeur de début
          ; add r0 15 r1
          scall 0
          ; compteur pour stockage mémoire
          add r3, r0, r0
L_LOOP:
          ; on affiche r1
          scall 1
          ; r1 <= 1 -> fin du programme
          sle r1, 1, r2
          branz r2, L_END
          ; on teste la parité r2 = r1 & 0x0001
          and r1, 1, r2
          branz r2, L_IMPAIR

          ; si r1 est pair r1 /= 2
L_PAIR:   div r1, 2, r1
          jmp L_LOOP, r0

          ; sinon r1 = r1*3 + 1
L_IMPAIR: mul r1, 3, r1
          add r1, 1, r1
          jmp L_LOOP, r0
L_END:
          stop