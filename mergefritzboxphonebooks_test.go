package mergefritzboxphonebooks

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSimpleAppend(t *testing.T) {
	// open the test files
	testXML1, err := os.OpenFile("testAppend1.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testAppend1.xml")
		t.Error(err)
	}
	defer testXML1.Close()

	testXML2, err := os.OpenFile("testAppend2.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testAppend2.xml")
		t.Error(err)
	}
	defer testXML2.Close()

	// call merging function
	MergeFritzBoxPhoneBooks(testXML1, testXML2, false, "resultTestAppend", "testAppend1")

	// open result file and validation file
	resultXML, err := os.OpenFile("resultTestAppend.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open resultTestAppend.xml")
		t.Error(err)
	}
	defer resultXML.Close()

	validationXML, err := os.OpenFile("validateTestAppend.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open validateTestAppend.xml")
		t.Error(err)
	}
	defer validationXML.Close()

	// read content of result file and validation file
	result, err := ioutil.ReadAll(resultXML)
	if err != nil {
		t.Log("Could not read resultTestAppend.xml")
		t.Error(err)
	}

	validation, err := ioutil.ReadAll(validationXML)
	if err != nil {
		t.Log("Could not read validateTestAppend.xml")
		t.Error(err)
	}

	// convert to string
	resultString := string(result)
	validationString := string(validation)

	// compare
	if resultString != validationString {
		t.Error("Test Append: Strings do not match")
	}
}

func TestRemoveSpecialCharacters(t *testing.T) {
	// open the test files
	testXML1, err := os.OpenFile("testRemoveSpecialCharacters1.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testRemoveSpecialCharacters1.xml")
		t.Error(err)
	}
	defer testXML1.Close()

	testXML2, err := os.OpenFile("testRemoveSpecialCharacters2.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testRemoveSpecialCharacters2.xml")
		t.Error(err)
	}
	defer testXML2.Close()

	// call merging function
	MergeFritzBoxPhoneBooks(testXML1, testXML2, true, "resultTestRemoveSpecialCharacters", "testRemoveSpecialCharacters1")

	// open result file and validation file
	resultXML, err := os.OpenFile("resultTestRemoveSpecialCharacters.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open resultTestRemoveSpecialCharacters.xml")
		t.Error(err)
	}
	defer resultXML.Close()

	validationXML, err := os.OpenFile("validateRemoveSpecialCharacters.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open validateRemoveSpecialCharacters.xml")
		t.Error(err)
	}
	defer validationXML.Close()

	// read content of result file and validation file
	result, err := ioutil.ReadAll(resultXML)
	if err != nil {
		t.Log("Could not read resultTestRemoveSpecialCharacters.xml")
		t.Error(err)
	}

	validation, err := ioutil.ReadAll(validationXML)
	if err != nil {
		t.Log("Could not read validateRemoveSpecialCharacters.xml")
		t.Error(err)
	}

	// convert to string
	resultString := string(result)
	validationString := string(validation)

	// compare
	if resultString != validationString {
		t.Error("Test Remove Special: Strings do not match")
	}
}

func TestSimpleDuplicate(t *testing.T) {
	// open the test files
	testXML1, err := os.OpenFile("testSimpleDuplicates1.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testSimpleDuplicates1.xml")
		t.Error(err)
	}
	defer testXML1.Close()

	testXML2, err := os.OpenFile("testSimpleDuplicates2.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testSimpleDuplicates2.xml")
		t.Error(err)
	}
	defer testXML2.Close()

	// call merging function
	MergeFritzBoxPhoneBooks(testXML1, testXML2, true, "resultSimpleDuplicates", "testDuplicate1")

	// open result file and validation file
	resultXML, err := os.OpenFile("resultSimpleDuplicates.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open resultSimpleDuplicates.xml")
		t.Error(err)
	}
	defer resultXML.Close()

	validationXML, err := os.OpenFile("validateSimpleDuplicate.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open validateSimpleDuplicate.xml")
		t.Error(err)
	}
	defer validationXML.Close()

	// read content of result file and validation file
	result, err := ioutil.ReadAll(resultXML)
	if err != nil {
		t.Log("Could not read resultSimpleDuplicates.xml")
		t.Error(err)
	}

	validation, err := ioutil.ReadAll(validationXML)
	if err != nil {
		t.Log("Could not read validateSimpleDuplicate.xml")
		t.Error(err)
	}

	// convert to string
	resultString := string(result)
	validationString := string(validation)

	// compare
	if resultString != validationString {
		t.Error("Test Simple Duplicate: Strings do not match")
	}
}

func TestComplexDuplicate(t *testing.T) {
	// open the test files
	testXML1, err := os.OpenFile("testComplexDuplicates1.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testComplexDuplicates1.xml")
		t.Error(err)
	}
	defer testXML1.Close()

	testXML2, err := os.OpenFile("testComplexDuplicates2.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open testComplexDuplicates2.xml")
		t.Error(err)
	}
	defer testXML2.Close()

	// call merging function
	MergeFritzBoxPhoneBooks(testXML1, testXML2, true, "resultComplexDuplicates", "testDuplicate1")

	// open result file and validation file
	resultXML, err := os.OpenFile("resultComplexDuplicates.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open resultComplexDuplicates.xml")
		t.Error(err)
	}
	defer resultXML.Close()

	validationXML, err := os.OpenFile("validateComplexDuplicate.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Log("Could not open validateComplexDuplicate.xml")
		t.Error(err)
	}
	defer validationXML.Close()

	// read content of result file and validation file
	result, err := ioutil.ReadAll(resultXML)
	if err != nil {
		t.Log("Could not read resultComplexDuplicates.xml")
		t.Error(err)
	}

	validation, err := ioutil.ReadAll(validationXML)
	if err != nil {
		t.Log("Could not read validateComplexDuplicate.xml")
		t.Error(err)
	}

	// convert to string
	resultString := string(result)
	validationString := string(validation)

	// compare
	if resultString != validationString {
		t.Error("Test Complex Duplicate: Strings do not match")
	}
}
