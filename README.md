# Nothing to see here, this is just for learning.

Links
- https://www.alexedwards.net/blog/i18n-managing-translations
- https://pkg.go.dev/golang.org/x/text@v0.3.7/message

1. Generate languages including english (source lang)
   `gotext update -lang=en,de,pt-BR`
2. `locales/{lang}/out.gettext.json` files will be generated
3. `cp locales/en/out.gettext.json locales/en/messages.gettext.json` to update the english text, including plurals.
4. `cp locales/en/messages.gettext.json locales/en/out.gettext.json`
5. Push `locales/en/out.gettext.json` to crowdin
6. Ask for translations
7. Download translations
8. Run `gotext update -out catalog.go -lang=en,de,pt-BR` to update the `catalog.go` file.


# Sample response

`$ go run .`

```txt
There are 1,500 flowers in our garden. 1,500 again
Existem 1.500 flores em nosso jardim. 1.500 novamente
Você tem 0 problema
Você tem 1 problema
Você tem 2 problemas

Du hast 0 Probleme
Du hast ein Problem
Du hast 2 Probleme

You have 0 problems
You have one problem
You have 2 problems

R$ 100.00
€ 100.00
$ 100.00
zł 100.00
```
