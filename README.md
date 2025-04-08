# GoLem

This project is a dictionary based lemmatizer written in go. 

Since v4 all dictionaries need to be gotten individually.

```
go get github.com/aaaton/golem/v4
```


### What?

A [lemmatizer](https://en.wikipedia.org/wiki/Lemmatisation) is a tool that finds the base form of words.

| Lang    | Input      | Output  |
| ------- | ---------- | ------- |
| English | aligning   | align   |
| Swedish | sprungit   | springa |
| French  | abattaient | abattre |

It's based on the dictionaries found on [michmech/lemmatization-lists](https://github.com/michmech/lemmatization-lists), which are available under the [Open Database License](https://opendatacommons.org/licenses/odbl/summary/). This project would not be feasible without them.

### Languages

At the moment golem supports English, Swedish, French, Spanish, Italian & German, but adding another language should be no more trouble than getting the dictionary for that language. Some of which are already available on lexiconista. Please let me know if there is something you would like to see in here, or fork the project and create a pull request.

English
```
go get github.com/aaaton/golem/v4/dicts/en
```

Swedish
```
go get github.com/aaaton/golem/v4/dicts/sv
```

French
```
go get github.com/aaaton/golem/v4/dicts/fr
```

German
```
go get github.com/aaaton/golem/v4/dicts/de
```

Spanish
```
go get github.com/aaaton/golem/v4/dicts/es
```

Italian
```
go get github.com/aaaton/golem/v4/dicts/it
```

### Basic usage

```golang
package main

import (
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

func main() {
	// the language packages are available under golem/dicts
	// "en" is for english
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	word := lemmatizer.Lemma("Abducting")
	if word != "abduct" {
		panic("The output is not what is expected!")
	}
}
```

### Contributors

- axamon
- charlesgiroux
- glaslos
- ptdewey
