---
title: "SIMJI - Simulateur de Jeu d'Instructions"
subtitle: "Un projet en Golang pour l'ENSTA Bretagne"
author: [Alexandre Froehlich]
date: "2021-03-02"
subject: "SIMJI"
keywords: [SIMJI, asm, alexandre, froehlich]
lang: "fr"
titlepage: true
toc-own-page: true
header-left: "\\theauthor"
header-right: "\\thetitle"
---

# Introduction

Nous sommes entourés de machines de Turing. Des automates qui agissent et réagissent à la moindre requête qu'on leur soumet. Il serait intéressant de revenir au base et de comprendre comment, à _bas niveau_ un ordinateur est capable d'interpréter une série de commandes qu'un programmeur lui donne : le programme.

## Motivation

SIMJI est un simulateur de jeu d'instruction développé dans le cadre de mes études à l'[ENSTA Bretagne](https://www.ensta-bretagne.fr/fr). L'enseignant nous laissant le choix du langage j'ai décidé d'utiliser ce projet pour sortir de ma zone de confort et me tourner vers un langage plus "_moderne_", j'ai donc choisi le [Golang](https://golang.org/).

![Le logo de Golang (prononcé "Go") est une marmotte bleue](https://miro.medium.com/max/3152/1*Ifpd_HtDiK9u6h68SZgNuA.png){ width=50% }

## Objectifs

Les objectifs principaux de ce projet sont les suivants :

-   découvrir l'architecture des processeurs
-   découvrir la compilation des langages

A cette fin nous avons réalisé un simulateur de jeu d'instruction se décomposant en plusieurs sous-programmes :

-   **un programme d'assemblage** permettant de prendre un programme rédigé en mini-MIPS et sortant des instructions machine
-   **un simulateur d'instructions** prenant ces instructions machines et exécutant le stack

# Architecture et langage

## Les registres

Le CPU (Compute Processing Unit) est une puce de silicone contenant, aujourd'hui, des milliards de transistors faisant passer ou non un signal électrique. Afin de traiter des calculs plus complexes que ceux effectués par les portes logiques que représentent ces transistors, il est intéressant de stocker les résultats obtenus.

C'est ici qu'interviennent les registres. Les registres sont des espaces mémoires d'accès extrêmement rapide en comparaison des disques durs classiques. Ils sont en nombre limités mais permettent de stocker temporairement des valeurs provenant du processeur.

## Les instructions

Pour que le programmeur puisse donner des ordres à la machine il doit écrire une série d'instructions. Nous allons pour ce projet nous baser sur le jeu d'instructions suivant, similaire au jeu [MIPS](https://fr.wikipedia.org/wiki/Architecture_MIPS) :

![Jeu d'instructions "mini-MIPS"](../jeu_d-instructions.png){ width=70% }

Pour se simplifier le travail, on va supposer que l'architecture du processeur que nous allons simuler est de 32 bits. Ainsi chacune des instructions sera elle même encodée sur 32 bits. On va ensuite pouvoir réserver chacun de ces bits pour une fonctionnalité donné :

-   4 bits pour l'opcode : le type d'instruction
-   5 bits pour chaque adresse de registre
-   1 bit pour indiquer si la valeur qui suit est immédiate ou bien un registre
-   16 bits pour la valeur immédiate ou bien l'adresse de registre

![Structure des instructions du jeu mini-MIPS sur 32 bits](../trame.drawio.png){ width=90% }

Cependant l'expression en binaire de ce code est bien trop lourde. On lui préférera l'écriture hexadécimale qui permet de ne pas perdre en informations tout en rendant l'écriture plus lisible.

Par exemple si nous souhaitons ajouter $5$ au registre `r1` l'instruction correspondante en assembleur est :

```asm
add r1, 5, r1
```

Qui se traduit alors en instruction machine par : `0110 1000 0110 0000 0000 0000 1010 0001` ou bien en hexadécimal `0x686000a1`.

On peut tout de suite mieux comprendre l'intérêt d'une telle écriture.

## La mémoire et le cache

J'ai parlé précédemment des registres. Cependant les registres étant en nombre limité, il faut trouver un autre moyen pour stocker l'information.

Pour cela le jeu d'instructions mini-MIPS nous fournit des instructions `load` et `store` pour manipuler ces valeurs dans la mémoire. Cependant leur accès reste bien plus lent, en témoigne le tableau suivant donnant le nombre de cycles nécessaires pour accéder à l'information en fonction de l'endroit où elle est stockée :

| Type              | Taille des données   | Nombre de cycles |
| ----------------- | -------------------- | ---------------: |
| Registre          | 4 octets = 32 bits   |                0 |
| Cache             | 64 octets            |                1 |
| Mémoire virtuelle | pages de 4kB         |              100 |
| Disque dur        | secteurs de disque   |          100 000 |
| Stockage réseau   | morceaux de fichiers |       10 000 000 |

Pour cela on implémente un _cache_ qui va enregistrer une donnée dans un espace mémoire plus rapide si elle est utilisée plus fréquemment afin d'améliorer les performances du système.

Dans ce projet on ne va s'intéresser qu'à l'implémentation d'un cache à correspondance directe. Dans ce cas chaque `set` du cache ne contient qu'une ligne. On peut alors résumer le fonctionnement du cache par le schéma suivant :

![Cycle de vie du cache avec des "hits" et des "misses". Ce cache fonctionne en mode "write through" c'est a dire en écrivant à la fois dans le cache et dans la mémoire.](../cache.drawio.png){ width=80% }

# L'assembleur

> Un assembleur est un programme d'ordinateur qui traduit un programme écrit en langage assembleur — essentiellement, une représentation mnémonique du langage machine — en code objet.  
> -- Wikipedia

L'assembleur est un outil qui prend en entrée des instructions rédigées en assembleur et qui les traduit en code compréhensible par la machine.

Dans ce projet nous devrons être capable d'interpréter les instructions du jeu mini-MIPS. De plus le langage doit pouvoir supporter les labels qui seront des références dans le programme pour la gestion de tout ce qui est "saut" dans le code.

Pour se faire il faut procéder de la manière suivante :

1. Récupérer le texte correspondant au programme
2. Séparer le texte en ensemble de lignes
3. Parcourir une première fois les lignes et récupérer les adresses des labels
4. Parcourir une seconde fois les lignes :
    1. Remplacer les références de labels par leur adresse
    2. Traduire les arguments par des codes
5. "Assembler" les codes des arguments en une instruction hexadécimale

Nous allons voir chacune des étapes précédentes et expliquer les quelques spécificités ou diffcultés encontrées.

## Manipulation de fichier

En golang, la manipulation des fichiers se réalise à l'aide d'une bibliothèque standard "outil". Une autre spécificité du langage est la gestion des erreurs. Ces derniers se manipulent comme des variables : on tente une action et on récupère la présence d'erreur ou non dans une variable. Si cette variable est non nulle, il y a eu une erreur et on peut agir en conséquences.

```go
// OpenFile ouvre un fichier et retourne ses lignes
// sous la forme d'un tableau
func OpenFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	// si il y a eu une erreur durant l'ouverture
	if err != nil { panic() }

	// on sépare les lignes obtenues dans le fichier
	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}
```

## Adresses des labels

La prochaines étape est de parcourir l'ensemble des lignes du programme pour trouver les labels et stocker leur adresse dans un dictionnaire. Pour cela golang possède un type standard : la map. On peut ainsi définit le type des clés ainsi que le type des valeurs : `var dictionnaire map[string]int`.

Pour détecter un label on doit d'abord être capable de détecter qu'une ligne n'est pas un commentaire ou bien une ligne vide :

```go
// EstVideOuCommentaire retourne vrai si la ligne est
// un commentaire ou bien est vide
func EstVideOuCommentaire(ligne string) bool {
	// si la ligne est vide on retourne vrai
	if ligne == "" { return true }

	// on sépare la ligne en arguments
	args := strings.Split(ligne, " ")

	// on vérifie si la ligne entière est un commentaire
	return args[0][0] == ";"
}
```

Ensuite une ligne définissant un label commence soit par le label, ou bien le label compose la ligne. Il suffit alors de vérifier que le premier mot finisse par `:` pour s'assurer de la définition d'un label :

```go
// ExtraitLabel retourne le label si il existe dans une ligne
func ExtraitLabel(ligne string) string {
	// on sépare la ligne en arguments
	args := strings.Split(ligne, " ")

	// on vérifie si le premier mot est un label
	if args[0][len(args[0])-1] != ":" { return "" }

	// on extrait le label
	return args[0][:len(args[0])-1]
}
```

Il ne reste alors plus qu'à parcourir le programme pour définir un dictionnaire des adresses des labels :

```go
// AdressesLabels construit un dictionnaire contenant
// les adresses des labels
func AdressesLabels(lignes []string) map[string]int {
	var adresses map[string]int
	var pc int

	for _, ligne := range lignes {
		// si la ligne n'est pas viable
		if !EstVideOuCommentaire(ligne) { continue }

		label := ExtraitLabel(ligne)
		// si il y a un label on enregistre son adresse
		if label != "" { adresses[label] = pc }
		pc++
	}

	return adresses
}
```

![SIMJI affiche cette étape intermédiaire en utilisant le flag `--debug` lors de son exécution.](../analyse_labels.png){ width=60% }

## Traduction des instructions en codes

Cette étape est la plus importante, en terme d'efforts, dans l'assembleur. Il va falloir itérer sur chacune des lignes et traduire chaque argument en code qui sera ensuite traduit en hexadécimal.

On peut commencer en concevant une fonction permettant de dire si un argument est un registre et si oui quel est son numéro de registre. Pour cela golang nous permet de retourner plusieurs valeurs qui ne sont pas du même type, un peu comme python. Cela se déclare dans la signature de la fonction :

```go
// EstUnRegistre permet de parser un argument et de dire si c'est un registre
func EstUnRegistre(argument string, labels map[string]int) (int, bool) {
	// on retire le "r" du registre si il est présent
	if argument[0] == 'r' {
		// on essaie de parser l'argument
		value, err := strconv.Atoi(argument[1:])
		// si il y a une erreur c'est un mauvais registre
		if err != nil {
			panic("Error while parsing register: ", argument)
		}
		return value, true
	}

	// si il s'agit d'un label on retourne son adresse
	if value, ok := labels[argument]; ok {
		return value, false
	}

	// on essaie de parser l'argument
	value, err := strconv.Atoi(argument)
	// si il a une erreur c'est qu'on ne sait pas quelle est cette valeur
	if err != nil {
		panic("Error while parsing immediate: ", argument)
	}

	// sinon on a réussi à parser la valeur
	return value, false
}
```

On est maintenant capable de comprendre chaque argument séparément en traduisant les adresses des labels et en séparant les registres des valeurs immédiates. Il ne reste plus qu'à traduire les instructions.

Pour information voici les codes associés à chaque type d'instruction :

```go
var OpCodes = map[string]int{
	"stop":   0,	"add": 	  1,
	"sub": 	  2, 	"mul": 	  3,
	"div": 	  4, 	"and": 	  5,
	"or": 	  6, 	"xor": 	  7,
	"shl":    8, 	"shr": 	  9,
	"slt":   10, 	"sle":   11,
	"seq":   12, 	"load":  13,
	"store": 14, 	"jmp":   15,
	"braz":  16, 	"branz": 17,
	"scall": 18,
}
```

![SIMJI affiche cette étape intermédiaire en utilisant le flag `--debug` lors de son exécution.](../analyse_code_trad.png){ width=60% }

## Traduction des codes en hexadécimal

C'est la dernière étape de l'assemblage. Il faut maintenant récupérer les codes et les assembler dans une instruction unique sur 32 bits. Pour cela on va s'aider des décalages binaires.

En golang les décalages binaires sont identique au C. On utilise pour cela les chevrons doubles : `<<`. Enfin pour reconstruire l'instruction on peut créer une fonction qui va respecter les règles que j'ai décrit au chapitre sur les instructions :

```go
func TraductionHexa(inst []int) int {
	decInstr := instr[0] << 27

	// on fait un switch sur le nombre d'arguments
	switch len(instr) {
	case 1:
		break
	case 2:
		// scall
		decInstr += instr[1] // num
		break
	case 3:
		// braz
		decInstr += instr[1] << 22 // reg
		decInstr += instr[2]       // address
	case 4:
		// jmp
		decInstr += instr[1] << 26                      // imm
		decInstr += BinaryComplement(instr[2], 21) << 5 // o
		decInstr += instr[3]                            // r
	case 5:
		// add, load, store ...
		decInstr += instr[1] << 22                      // reg
		decInstr += instr[2] << 21                      // imm
		decInstr += BinaryComplement(instr[3], 16) << 5 // o
		decInstr += instr[4]                            // reg
		break
	}

	return decInstr
}
```

On a utilisé ici une fonction maison permettant d'écrire un nombre entier signé en nombre binaire en complément à 2. Voici son implémentation :

```go
// BinaryComplement permet de calculer le complément à 2 d'un entier
func BinaryComplement(number int, size int) int {
	// aucun travail à faire
	if number >= 0 { return number }

	return (1 << (size - 1)) - number
}
```

## Exemple

Pour le programme suivant, permettant de calculer les termes de la suite de syracuse, on a le programme en assembleur :

```asm
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
L_IMPAIR: mult r1, 3, r1
          add r1, 1, r1
          jmp L_LOOP, r0
L_END:
          stop
```

Qui donne après assemblage par notre programme :

![Assemblage du programme syracuse en utilisant la commande "assemble" de SIMJI](../assemble_syracuse.png){ width=80% }

# La machine virtuelle

## Interprétation des instructions

## Désassemblage

# Le logiciel

## La ligne de commande (CLI)

![SIMJI utilise la bibliothèque [Cobra](https://github.com/spf13/cobra) pour standardiser la CLI](../cli_help.png)

![SIMJI est capable d'indiquer ce qui se passe à l'utilisateur, notamment lorsqu'il oublie de préciser un fichier source](../cli_missing_file.png)

## L'interface utilisateur (GUI)

![En utilisant le flag `--gui` SIMJI ouvre une interface graphique permettant d'interagir avec le code](../graphical_ui.png)

# Conclusion

```

```
