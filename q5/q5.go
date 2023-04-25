package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	if strings.Contains(s, "a") {
		return strings.ReplaceAll(s, "a", "")
	}
	if strings.Contains(s, "e") {
		return strings.ReplaceAll(s, "e", "")
	}
	if strings.Contains(s, "i") {
		return strings.ReplaceAll(s, "i", "")
	}
	if strings.Contains(s, "o") {
		return strings.ReplaceAll(s, "o", "")
	}
	if strings.Contains(s, "u") {
		return strings.ReplaceAll(s, "u", "")
	}
	if strings.Contains(s, "A") {
		return strings.ReplaceAll(s, "A", "")
	}
	if strings.Contains(s, "E") {
		return strings.ReplaceAll(s, "E", "")
	}
	if strings.Contains(s, "I") {
		return strings.ReplaceAll(s, "I", "")
	}
	if strings.Contains(s, "O") {
		return strings.ReplaceAll(s, "O", "")
	}
	if strings.Contains(s, "U") {
		return strings.ReplaceAll(s, "U", "")
	}
	
	return ""
}
