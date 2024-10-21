package main

import "fmt"

type Stack []rune

func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) Pop() rune {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Top() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func getBool(ch rune) bool {
	return ch == 't'
}

// Обработка действия для оператора
func (s *Stack) action(logical rune, operands []bool) bool {
	cur := operands[0]
	if logical == '!' {
		return !cur
	}
	for _, val := range operands[1:] {
		if logical == '|' {
			cur = cur || val
		} else {
			cur = cur && val
		}
	}
	return cur
}

func parseBoolExpr(expression string) bool {
	opStack := Stack{} // Стек для операторов
	for _, ch := range expression {
		if ch == 't' || ch == 'f' || ch == '!' || ch == '&' || ch == '|' || ch == '(' {
			opStack.Push(ch) // Добавляем оператор или операнд в стек
		} else if ch == ')' {
			var operands []bool
			for opStack[len(opStack)-1] != '(' { // Извлекаем все операнды до открывающей скобки
				val := opStack.Pop()
				if val == 't' || val == 'f' {
					operands = append([]bool{getBool(val)}, operands...)
				}
			}
			opStack.Pop() // Удаляем '(' из стека

			logical := opStack.Pop() // Извлекаем оператор
			result := opStack.action(logical, operands)

			// Сохраняем результат операции
			if result {
				opStack.Push('t')
			} else {
				opStack.Push('f')
			}
		}
	}
	return getBool(opStack.Pop())
}

func main() {
	fmt.Println(parseBoolExpr("!(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))))")) // true
}
