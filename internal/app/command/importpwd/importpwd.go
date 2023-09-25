package importpwd

import (
	"encoding/csv"
	"fmt"
	"github.com/swicherwich/pwdmg/internal/app/command/save"
	"os"
)

func ImportFromChrome(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		d := record[0]
		acc := record[2]
		pwd := record[3]
		save.PersistAccount(d, acc, pwd)
	}
}
