package main

import (
	"demo03/ch4/JSON/web-json"
	"fmt"
	"log"
	"os"
)

func main() {
	result, err := web_json.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
