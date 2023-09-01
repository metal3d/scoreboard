package i18n

import (
	"sync"
)

//go:generate go run i18n_generator.go
//go:generate go fmt ./

const defaultLang = "en_US"

var lang = defaultLang
var translations = map[string]map[string]string{}
var mutex sync.Mutex

func I(key string) string {
	mutex.Lock()
	defer mutex.Unlock()
	if translations[lang] == nil {
		return key
	}
	if v, ok := translations[lang][key]; !ok {
		// try to find the key in the english version
		if v, ok := translations[defaultLang][key]; ok {
			return v
		}
		return key // not found, return the key
	} else {
		return v // found, return the translation
	}
}

func SetLang(l string) {
	mutex.Lock()
	defer mutex.Unlock()
	lang = l
}
