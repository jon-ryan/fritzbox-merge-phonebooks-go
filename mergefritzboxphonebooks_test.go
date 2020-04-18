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
