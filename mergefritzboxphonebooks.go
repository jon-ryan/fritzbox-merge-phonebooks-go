package mergefritzboxphonebooks

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// define structs to read xml file

// Phonebooks holds all phonebooks
type Phonebooks struct {
	XMLName    xml.Name    `xml:"phonebooks"`
	Phonebooks []Phonebook `xml:"phonebook"`
}

// Phonebook holds contact entries
type Phonebook struct {
	XMLName  xml.Name  `xml:"phonebook"`
	Name     string    `xml:"name,attr"`
	Contacts []Contact `xml:"contact"`
}

// Contact holds the contact information
type Contact struct {
	XMLName   xml.Name  `xml:"contact"`
	Category  int32     `xml:"category"`
	Person    Person    `xml:"person"`
	Telephony Telephony `xml:"telephony"`
	Services  Services  `xml:"services"`
	ModTime   string    `xml:"mod_time,omitempty"`
	UniqueID  int32     `xml:"uniqueid"`
}

// Person holds the person information
type Person struct {
	XMLName  xml.Name `xml:"person"`
	RealName string   `xml:"realName"`
}

// Telephony holds the numbers associated with the person
type Telephony struct {
	XMLName xml.Name `xml:"telephony"`
	Nid     int32    `xml:"nid,attr"`
	Numbers []Number `xml:"number"`
}

// Number holds all information about a specific number
type Number struct {
	XMLName   xml.Name `xml:"number"`
	Type      Type     `xml:"type,attr"`
	Quickdail string   `xml:"quickdial,attr"`
	Vanity    string   `xml:"vanity,attr"`
	Priority  Priority `xml:"prio,attr"`
	ID        int32    `xml:"id,attr"`
	Number    string   `xml:",chardata"`
}

// Services hols information about email addresses
type Services struct {
	XMLName xml.Name `xml:"services"`
	Nid     int32    `xml:"nid,attr,omitempty"`
	Emails  []Email  `xml:"email"`
}

// Email holds information about a persons email address
type Email struct {
	XMLName    xml.Name `xml:"email"`
	Classifier string   `xml:"classifier,attr"`
	ID         int32    `xml:"id,attr"`
	Address    string   `xml:",chardata"`
}

// Type enum
type Type string

const (
	home   Type = "home"
	mobile      = "mobile"
	work        = "work"
)

// Priority enum
type Priority int32

const (
	normal Priority = 0
	high            = 1
)

// clean the phonenumber from '.', ',', '-' or ' '
func cleanPhonenumbers(book *Phonebooks) {
	for i, phonebook := range book.Phonebooks {
		for j, contact := range phonebook.Contacts {
			for k := range contact.Telephony.Numbers {
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, ".", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, ",", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, " ", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "-", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "(", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, ")", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "/", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "\\", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "!", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "?", "", -1)
				book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number = strings.Replace(book.Phonebooks[i].Contacts[j].Telephony.Numbers[k].Number, "$", "", -1)
			}
		}
	}
}

// mergePhonebooks requires two Phonebooks and merges them into one discarding duplicate entries
// result will be in book1
func mergePhonebooks(book1 *Phonebooks, book2 *Phonebooks) {
	// temp counter
	tempCounter := 0

	// temp storage for non-duplicate contacts
	var contacts []Contact

	// output
	fmt.Println("\n----------")

	// check phonenumbers per contact to find out duplicates
	// outer loop iterates over book 2
	// duplicates will be found in book 1
	// if a contact is not a duplicate it will be added from book 1
	// duplicates will be appended to book1.Phonebooks[0]
	for b2i := range book2.Phonebooks {
		for b1i := range book1.Phonebooks {

		checkContactInBook2:
			for b2j := range book2.Phonebooks[b2i].Contacts {
				for b1j := range book1.Phonebooks[b1i].Contacts {

					// check if a number of Contact[j] is part of a Contact in book1
					for b2k := range book2.Phonebooks[b2i].Contacts[b2j].Telephony.Numbers {
						for b1k := range book1.Phonebooks[b1i].Contacts[b1j].Telephony.Numbers {
							// check for duplicate numbers per contact
							if book1.Phonebooks[b1i].Contacts[b1j].Telephony.Numbers[b1k].Number == book2.Phonebooks[b2i].Contacts[b2j].Telephony.Numbers[b2k].Number {
								//fmt.Println("Duplicate:", book1.Phonebooks[b1i].Contacts[b1j].Person.RealName, "and", book2.Phonebooks[b2i].Contacts[b2j].Person.RealName)
								tempCounter++
								continue checkContactInBook2
							}
						}
					}
				}
				// it's not a duplicate -> append
				contacts = append(contacts, book2.Phonebooks[b2i].Contacts[b2j])
			}
		}
		// output
		fmt.Println("Count of non-duplicates:", len(contacts))
		// append contacts
		book1.Phonebooks[0].Contacts = append(book1.Phonebooks[0].Contacts, contacts...)
		contacts = nil
	}

	fmt.Println("Duplicates found:", tempCounter)
	fmt.Println("----------")
}

// MergeFritzBoxPhoneBooks takes two pointers to two XML exports of FritzBox phonebooks
// If removeSpecialCharacters is true the numbers will be clean from the characters spaces, commas, dashes, full stops, brackets slash, backslash, !, ? or $
// The result will be written to a new XML file. The name is specified by 'outputFilename'
// The name of the resulting phonebook is specified at 'phonebookName'
func MergeFritzBoxPhoneBooks(book1, book2 *os.File, removeSpecialCharacters bool, outputFilename, phonebookName string) {
	// read content of book 1
	byteContentBook1, errRead1 := ioutil.ReadAll(book1)
	if errRead1 != nil {
		fmt.Println("Error reading book 1:")
		fmt.Println(errRead1)
	}

	// read content of book 2
	byteContentBook2, errRead2 := ioutil.ReadAll(book2)
	if errRead2 != nil {
		fmt.Println("Error reading book 2:")
		fmt.Println(errRead2)
	}

	// initialize phonebooks1 object
	var phonebooks1 Phonebooks
	var phonebooks2 Phonebooks

	// unmarshal book1 xml file
	unmarshalError := xml.Unmarshal(byteContentBook1, &phonebooks1)
	if unmarshalError != nil {
		log.Fatal(unmarshalError)
	}

	// unmarshal book2 xml file
	unmarshalError = xml.Unmarshal(byteContentBook2, &phonebooks2)
	if unmarshalError != nil {
		log.Fatal(unmarshalError)
	}

	if removeSpecialCharacters {
		// clean contact numbers from '.', '-', ' ', '(' and ')'
		cleanPhonenumbers(&phonebooks1)
		cleanPhonenumbers(&phonebooks2)
	}

	// print stats
	fmt.Println("--- Shape of phonebooks ---")
	fmt.Println("Book 1:")
	fmt.Println("Phonebooks:", len(phonebooks1.Phonebooks))
	for i := 0; i < len(phonebooks1.Phonebooks); i++ {
		fmt.Println("Contacts in phonebook", i+1, ":", len(phonebooks1.Phonebooks[i].Contacts))
	}
	fmt.Println("---")
	fmt.Println("Book 2:")
	fmt.Println("Phonebooks:", len(phonebooks2.Phonebooks))
	for i := 0; i < len(phonebooks2.Phonebooks); i++ {
		fmt.Println("Contacts in phonebook", i+1, ":", len(phonebooks2.Phonebooks[i].Contacts))
	}

	// merge the two books
	// result in phonebook1
	mergePhonebooks(&phonebooks1, &phonebooks2)

	// rename phonebok
	phonebooks1.Phonebooks[0].Name = phonebookName

	fmt.Println("\n--- Shape of Result Book ---")
	fmt.Println("Name of Phonebook:", phonebooks1.Phonebooks[0].Name)
	fmt.Println("Phonebooks:", len(phonebooks1.Phonebooks))
	for i := 0; i < len(phonebooks1.Phonebooks); i++ {
		fmt.Println("Contacts in phonebook", i+1, ":", len(phonebooks1.Phonebooks[i].Contacts))
	}
	fmt.Println("------------")

	// wirte to new file
	content, marshalError := xml.MarshalIndent(phonebooks1, "", "    ")
	if marshalError != nil {
		fmt.Println(marshalError)
	}

	// open target file
	f, fileError := os.OpenFile(outputFilename+".xml", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if fileError != nil {
		fmt.Println(fileError)
	}

	_, headerWriteErr := f.WriteString(xml.Header)
	if headerWriteErr != nil {
		fmt.Println(headerWriteErr)
	}
	_, bodyWriteErr := f.Write(content)
	if bodyWriteErr != nil {
		fmt.Println(bodyWriteErr)
	}
	// done
	fmt.Println("Done")
}
