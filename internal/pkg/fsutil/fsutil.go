package fsutil

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

const sysDir = ".pwdmg"

func ReadData(dFile string, d any) error {
	c, err := os.ReadFile(dFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return err
	}

	err = json.Unmarshal(c, &d)
	if err != nil {
		fmt.Println("Error deserializing from json:", err)
		return err
	}
	return nil
}

func PersistDataToFile(dFile string, d any) error {
	file, err := os.Create(dFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	jData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to json:", err)
		return err
	}

	_, err = file.Write(jData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	_, err = fmt.Fprintf(file, "")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

func GetPwdmgDir() (string, error) {
	dirUser, err := user.Current()
	if err != nil {
		fmt.Println("Error obtaining home dir:", err)
		return "", err
	}

	pwdmgDir := fmt.Sprintf("%s/%s", dirUser.HomeDir, sysDir)
	err = os.MkdirAll(pwdmgDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating pwdmg:", err)
		return "", err
	}

	return pwdmgDir, nil
}

func FileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
