package fsutil

import (
	"fmt"
	"os"
	"testing"
)

func TestReadPersistData(t *testing.T) {
	// Test the ReadData and PersistDataToFile functions
	dFile := "testdata.json"
	defer func() { _ = os.Remove(dFile) }()

	// Sample data structure for testing
	type TestData struct {
		Name  string
		Value int
	}

	// Create sample data
	dataToPersist := TestData{
		Name:  "Test",
		Value: 42,
	}

	// Persist the data to a file
	err := PersistDataToFile(dFile, dataToPersist)
	if err != nil {
		t.Fatalf("PersistDataToFile error: %v", err)
	}

	// Read the data back from the file
	var loadedData TestData
	err = ReadData(dFile, &loadedData)
	if err != nil {
		t.Fatalf("ReadData error: %v", err)
	}

	// Compare the loaded data with the original data
	if loadedData != dataToPersist {
		t.Errorf("Loaded data (%v) is not equal to the original data (%v)", loadedData, dataToPersist)
	}
}

func TestGetPwdmgDir(t *testing.T) {
	// Test the GetPwdmgDir function
	dir, err := GetPwdmgDir()
	if err != nil {
		t.Fatalf("GetPwdmgDir error: %v", err)
	}

	// Check if the directory exists
	if !FileExists(dir) {
		t.Errorf("GetPwdmgDir did not create the directory %s", dir)
	}
}

func TestFileExists(t *testing.T) {
	// Test the FileExists function
	nonExistingFile := "nonexistent.json"

	existingFile, err := os.CreateTemp("./", "testdata")
	if err != nil {
		return
	}
	defer func() { _ = os.Remove(existingFile.Name()) }()

	fmt.Print(
		existingFile.Name())

	if !FileExists(existingFile.Name()) {
		t.Errorf("FileExists returned false for an existing file: %s", existingFile.Name())
	}

	if FileExists(nonExistingFile) {
		t.Errorf("FileExists returned true for a non-existing file: %s", nonExistingFile)
	}

}
