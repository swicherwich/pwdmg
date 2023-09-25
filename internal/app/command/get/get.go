package get

import (
	"errors"
	"fmt"
	"pwd-manager/internal/app/domain"
	"pwd-manager/internal/pkg/fsutil"
	"pwd-manager/internal/pkg/secutil"
)

func PwdByLogin(dName, acc string) (string, error) {
	dHash := secutil.HashStr(dName)
	accHash := secutil.HashStr(acc)
	pwdmgDir, _ := fsutil.GetPwdmgDir()
	dFile := fmt.Sprintf("%s/%s%s", pwdmgDir, dHash, ".json")

	var d domain.Domain

	err := fsutil.ReadData(dFile, &d)
	if err != nil {
		fmt.Print("Error reading file:", err)
		return "", err
	}

	for _, account := range d.Accounts {
		if accHash == account.Login {
			pwdD, err := secutil.DecodeBase64(account.Password)
			if err != nil {
				return "", err
			}
			return pwdD, nil
		}
	}

	return "", errors.New(fmt.Sprintf("no such login: %s", acc))
}
