# FritzBox - Merge phonebooks
This programm aims to merge two phonebooks managed by a FritzBox. The phonebooks have to be exported as an XML file. They will then be merged which will drop duplicate entries. Duplicates are determined by phonenumber. Also, if a phonenumber contains spaces, commas, dashes, full stops or brackets and the 'removeSpecialCharacters' flag is true it will remove them.
