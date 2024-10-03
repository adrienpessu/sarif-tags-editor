package main

import (
	"flag"
	"fmt"
	sarif "github.com/adrienpessu/sarif-tags-editors/libs/sarif"
	"io"
	"os"
	"strings"
)

func main() {

	// GitHub configuration
	sarifFileSource := flag.String("sarif-file-source", "", "SARIF file source")
	sarifFileDestination := flag.String("sarif-file-destination", "", "SARIF file destination")
	filtersTags := flag.String("filters-tags", "", "Filters tags")
	addTags := flag.String("add-tags", "", "add tags")
	removeTags := flag.String("remove-tags", "", "remove tags")

	flag.Parse()

	if sarifFileSource != nil && *sarifFileSource != "" && sarifFileDestination != nil && *sarifFileDestination != "" && filtersTags != nil && *filtersTags != "" && ((addTags != nil && *addTags != "") || (removeTags != nil && *removeTags != "")) {
		// read file in sarifPath and store it in the variable sarif
		file, err := os.Open(*sarifFileSource)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		sarifByte, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		sarifOutput := string(sarifByte)

		sarifFinal := sarif.SearchAndAssignOrReplaceTags(sarifOutput, strings.Split(*filtersTags, ","), strings.Split(*addTags, ","), strings.Split(*removeTags, ","))
		// write the sarif to the file sarifFileDestination
		err = os.WriteFile(*sarifFileDestination, []byte(sarifFinal), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		fmt.Println("You need to provide a SARIF file source, a SARIF file destination, filters tags and assigned tags")
		return
	}

}
