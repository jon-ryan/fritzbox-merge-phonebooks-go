package main

import (
	"flag"
	"fmt"
	"log"
	mergefritzboxphonebooks "mergeFritzBoxPhoneBooks"
	"os"
)

func main() {

	// get the xml files
	phonebook1 := flag.String("phonebook1", "", "Set the filename of the first phonebook")
	phonebook2 := flag.String("phonebook2", "", "Set the filename of the second phonebook")

	flag.Parse()

	// check for empty parameters
	if *phonebook1 == "" || *phonebook2 == "" {
		fmt.Println("Please specifiy two phonebooks")
		return
	}

	// read the xml files
	xmlPhonebook1, errPhonebook1 := os.OpenFile(*phonebook1, os.O_RDONLY, 0755)
	if errPhonebook1 != nil {
		fmt.Println("Error opening book 1:")
		log.Fatal(errPhonebook1)
	}
	defer xmlPhonebook1.Close()

	xmlPhonebook2, errPhonebook2 := os.OpenFile(*phonebook2, os.O_RDONLY, 0755)
	if errPhonebook2 != nil {
		fmt.Println("Error opening book 2:")
		log.Fatal(errPhonebook2)
	}
	defer xmlPhonebook2.Close()

	// merge
	mergefritzboxphonebooks.MergeFritzBoxPhoneBooks(xmlPhonebook1, xmlPhonebook2)

}
