package i18n

import (
	"sync"
)

//go:generate go run i18n_generator.go
//go:generate go fmt ./

var lang = "en_US"
var translations = map[string]map[string]string{}
var mutex sync.Mutex

func I(key string) string {
	mutex.Lock()
	defer mutex.Unlock()
	if translations[lang] == nil {
		return key
	}
	if translations[lang][key] == "" {
		return key
	}
	return translations[lang][key]
}

func SetLang(l string) {
	mutex.Lock()
	defer mutex.Unlock()
	lang = l
}
