package main

import "fmt"

func main() {

	fighters := make(map[string]map[string][]string)
	fighters["Fedor"] = map[string][]string{
		"books":     {"Bibliya", "Uroki memologii", "Kak ne bit' dushnim"},
		"newsPaper": {"bloomberg", "playboy", "kommersant"},
	}
	fighters["Habib"] = map[string][]string{
		"books":     {"Karan", "Uroki geografiya 9 klass", "Kak bit' dushnim"},
		"newsPaper": {"aldgazira", "playboy", "svobodniy Chtototamrstan"},
	}
	fighters["Neznaika"] = map[string][]string{
		"books":     {},
		"newsPaper": {},
	}
	with := 0
	for k, _ := range fighters {
		summaIzdanii := len(fighters[k]["books"]) + len(fighters[k]["newsPaper"])
		if summaIzdanii != 0 {
			with++
		}
		fmt.Println("Общее количество изданий у", k, "=", summaIzdanii)
	}
	fmt.Println("Количество учеников без изданий на руках =", with)
}
