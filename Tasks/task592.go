package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack []rune

type Fraction struct {
	numerator   int
	denominator int
}

func (s *Stack) add(value rune) {
	*s = append(*s, value)
}

func (s *Stack) pop() rune {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) popLeft() rune {
	item := (*s)[0]
	*s = (*s)[1:]
	return item
}

func (s *Stack) getLast() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func isOperator(value rune) bool {
	switch value {
	case '+':
		return true
	case '-':
		return true
	default:
		return false
	}
}

func stringToFraction(stack *Stack, fractions *[]Fraction) {
	numerator := strings.Builder{}
	denominator := strings.Builder{}
	isNumerator := true

	for !stack.isEmpty() {
		char := stack.popLeft()
		if char == '/' {
			isNumerator = false
			continue
		}
		if isNumerator {
			numerator.WriteRune(char)
		} else {
			denominator.WriteRune(char)
		}
	}
	intNumerator, _ := strconv.Atoi(numerator.String())
	intDenominator, _ := strconv.Atoi(denominator.String())
	fraction := Fraction{numerator: intNumerator, denominator: intDenominator}
	*fractions = append(*fractions, fraction)
}

func fractionsToAction(fractions []Fraction, operators Stack) string {
	resFraction := Fraction{numerator: fractions[0].numerator, denominator: fractions[0].denominator}
	for _, fraction := range fractions[1:] {
		action := operators.popLeft()
		if resFraction.denominator != fraction.denominator {
			generalDenominator := resFraction.denominator * fraction.denominator
			resFraction.numerator *= generalDenominator / resFraction.denominator
			resFraction.denominator = generalDenominator
			fraction.numerator *= generalDenominator / fraction.denominator
			fraction.denominator = generalDenominator
		}
		resFraction = fractionsActions(resFraction, fraction, action)
	}
	resFraction = dividerFraction(resFraction)
	numeratorIsGood(&resFraction)
	return fmt.Sprintf("%d/%d", resFraction.numerator, resFraction.denominator)
}

func dividerFraction(fraction Fraction) Fraction {
	number := int(math.Min(math.Abs(float64(fraction.numerator)), math.Abs(float64(fraction.denominator))))
	for {
		flag := false
		for i := 2; i <= number; i++ {
			if fraction.numerator%i == 0 && fraction.denominator%i == 0 {
				fraction.numerator /= i
				fraction.denominator /= i
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return fraction
}

func numeratorIsGood(fraction *Fraction) {
	if (*fraction).numerator == 0 {
		(*fraction).denominator = 1
	}
}

func fractionsActions(fraction1, fraction2 Fraction, action rune) Fraction {
	switch action {
	case '+':
		fraction1.numerator += fraction2.numerator
	case '-':
		fraction1.numerator -= fraction2.numerator
	}
	return fraction1
}

func fractionAddition(expression string) string {
	stack := Stack{}
	operators := Stack{}
	fractions := make([]Fraction, 0)
	for _, char := range expression {
		if isOperator(char) && !stack.isEmpty() {
			operators.add(char)
			stringToFraction(&stack, &fractions)
			continue
		}
		stack.add(char)
	}
	stringToFraction(&stack, &fractions)
	return fractionsToAction(fractions, operators)
}

func main() {
	fmt.Println(fractionAddition("-1/4-4/5-1/4"))
}
