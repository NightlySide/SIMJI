package assembler

// OpCodes représente tous les codes associés aux instructions
var OpCodes = map[string]int{
	"stop": 	0,		"add": 		1, 
	"sub": 		2, 		"mult": 	3, 
	"div": 		4,		"and": 		5, 
	"or": 		6,		"xor": 		7, 
	"shl": 		8, 		"shr": 		9, 
	"slt": 		10, 	"sle": 		11, 
	"seq": 		12, 	"load": 	13,
	"store": 	14,		"jmp": 		15, 
	"braz": 	16, 	"branz":	17, 
	"scall": 	18,
}
