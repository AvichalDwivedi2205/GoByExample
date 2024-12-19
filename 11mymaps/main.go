package main

import "fmt"

func main(){
	fmt.Println("Maps In Golang")

	//We can also take it as hash table in some languages

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	//loops are interesting in GoLang

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
    // We can ignore anything with _ in golang
}