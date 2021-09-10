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
