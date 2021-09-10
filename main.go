package main

//go:generate gotext update -out catalog.go -lang=en,de,pt-BR

import (
	"fmt"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	total := 1500
	p := message.NewPrinter(language.English)
	_, _ = p.Printf("There are %v flowers in our garden. %[1]v again\n", total)

	p = message.NewPrinter(language.BrazilianPortuguese)
	_, _ = p.Printf("There are %v flowers in our garden. %[1]v again\n", total)

	PrintProblems(language.BrazilianPortuguese, 0)
	PrintProblems(language.BrazilianPortuguese, 1)
	PrintProblems(language.BrazilianPortuguese, 2)
	fmt.Println()

	PrintProblems(language.German, 0)
	PrintProblems(language.German, 1)
	PrintProblems(language.German, 2)
	fmt.Println()

	PrintProblems(language.English, 0)
	PrintProblems(language.English, 1)
	PrintProblems(language.English, 2)

	fmt.Println()
	PrintMoney(language.BrazilianPortuguese)
	PrintMoney(language.German)
	PrintMoney(language.AmericanEnglish)
	PrintMoney(language.Polish)
}

func PrintProblems(lang language.Tag, problems int) {
	p := message.NewPrinter(lang)
	_, _ = p.Printf("You have %d problems\n", problems)
}

func PrintMoney(lang language.Tag) {
	p := message.NewPrinter(lang)
	unit, _ := currency.FromTag(lang)
	amount := unit.Amount(100.0)
	formatter := currency.Symbol.Default(unit)
	_, _ = p.Println(formatter(amount))
}
