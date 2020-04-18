# FritzBox - Merge phonebooks
This programm aims to merge two phonebooks managed by a FritzBox. The phonebooks have to be exported as an XML file. They will then be merged which will drop duplicate entries. Duplicates are determined by phonenumber. Also, if a phonenumber contains spaces, commas, dashes, full stops, brackets slash, backslash, !, ? or $ and the 'removeSpecialCharacters' flag is true it will remove them.

# Signature
The avaiable function has the signature
- MergeFritzBoxPhoneBooks(book1, book2 *os.File, removeSpecialCharacters bool, outputFilename, phonebookName string)
where
- _book1_ and _book2_ are the file descriptors to XML Files in questsion
- _removeSpecialCharacters_ sets if the special characters should be removed
- _outputFilename_ specifies the resulting XML-File name
- _phonebookName_ sets the name of the resulting phonebook
