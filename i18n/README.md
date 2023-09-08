# I18n

Internationalization is generated from yaml files. The procedure is simple:

- If the file doens't exists yet, create an empty file named `xx_XX.yaml` in this folder containing `{}` where `xx_XX` is the languages code (e.g. `de_DE`): `echo "{}" > i18n/xx_XX.yaml`
- call the generator: `go generate ./...`
- the file is now filled with keys and translation. All the translation value are **in english** at this time.
- change **the values** (not the keys), we mean that you must now translate english to the target langage.
- call the generator one more time, to apply your changes: `go genetare ./...`
- now the `xx_XX.go` file contains the translations. Also, the `loader.go` file should reference the new langage


In short:

```bash
# replace xx_XX to the lang code:
app_lang="xx_XX"

# make the file if it does not exists
[ -f i18n/${app_lang}.yaml ] || echo "{}" > i18n/${app_lang}.yaml

unset app_lang

# generate keys
go genetare ./...

# => open the i18n/xx_XX.yaml file, translate the values, then
# regenerate
go generate ./...
```


## What does the generator?

The generator searches for YAML file in the `i18n` folder. It loads them in memory. Then, it searches in the entire go files something like `I(...)`.

If the key is not found in the yaml, it is appended. Existing keys are not removed, not modified. The newest keys contain the english translation if it exists. If the english translation doesn't exist, so the key is set as value (`the.word: the.word`).

Then, it rewrites the YAML file to include the newest keys.

Then, it writes all `xx_XX.go` files to convert the YAML content to a `map`.

Finally, it writes the `loader.go` file to reference all langages and translation.
