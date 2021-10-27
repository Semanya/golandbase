package main

import "fmt"

func main() {
	var fighters map[string]map[string][]string = map[string]map[string][]string{
		"Fedor": {
			"books":     {"Bibliya", "Uroki memologii", "Kak ne bit' dushnim"},
			"newsPaper": {"bloomberg", "playboy", "kommersant"},
		},
		"Habib": {
			"books":     {"Karan", "Uroki geografiya 9 klass", "Kak bit' dushnim"},
			"newsPaper": {"aldgazira", "playboy", "svobodniy Chtototamrstan"},
		},
		"Neznaika": {
			"books":     {},
			"newsPaper": {},
		},
	}

	keys := make([]string, 0, len(fighters))
	for k := range fighters {
		keys = append(keys, k)
	}
	without := 0
	for _, v := range keys {
		izdaniya := fighters[v]
		summaIzdanii := len(izdaniya["books"]) + len(izdaniya["newsPaper"])
		if summaIzdanii == 0 {
			without = without + 1
		}
		fmt.Println("Общее количество изданий у", v, "=", summaIzdanii)
	}
	fmt.Println("Количество учеников без изданий на руках =", without)
}
