# GoLem
This project is a dictionary based lemmatizer written in go. 

### What?
A lemmatizer is a tool that finds the base form of words.
|lang|input|output|
|----|-----|------|
|en: | aligning | align |
|sv: | sprungit | springa |
|fr: | abattaient | abattre |

It's based on the dictionaries found on [lexiconista.com](http://www.lexiconista.com/datasets/lemmatization/), which are available under the [Open Database License](https://opendatacommons.org/licenses/odbl/summary/). This project would not be feasible without them. 

### Languages
At the moment I have added English, Swedish, French, Spanish & German, but implementing another language on the list should be no more trouble than getting the dictionary for that language. Some of which are already available on lexiconista.

### Basic usage
```go
package main

import (
	"github.com/aaaton/golem"
)

func main() {
	// "en" and "english" will give an english lemmatizer
	lemmatizer, err := golem.New("english")
	if err != nil {
		panic(err)
	}
	word, err := lemmatizer.Lemma("Abducting")
	if err != nil {
		panic(err)
	}
	if word != "abduct" {
		panic("The output is not what is expected!")
	}
}

```