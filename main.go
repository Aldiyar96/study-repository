package main

import (
	"Study/feature"
	feature_postgres "Study/feature-postgres"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Aldiyar")
	feature.Feature()
	feature_postgres.CheckConnection()
}
