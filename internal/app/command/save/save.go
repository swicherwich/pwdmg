package save

import (
	"fmt"
	"github.com/swicherwich/pwdmg/internal/app/domain"
	"github.com/swicherwich/pwdmg/internal/pkg/fsutil"
	"github.com/swicherwich/pwdmg/internal/pkg/secutil"
	"time"
)

func PersistAccount(dom, acc, pwd string) {
	pwdmgDir, _ := fsutil.GetPwdmgDir()

	dHash := secutil.HashStr(dom)
	accHash := secutil.HashStr(acc)
	pwd64 := secutil.EncodeBase64(pwd)
	dFile := fmt.Sprintf("%s/%s%s", pwdmgDir, dHash, ".json")

	if fsutil.FileExists(dFile) {
		var d domain.Domain

		err := fsutil.ReadData(dFile, &d)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		d.Metadata = domain.Metadata{
			Entries:      d.Metadata.Entries + 1,
			LastModified: time.Now().Format(time.DateOnly),
		}

		if len(d.Accounts) > 0 {
			for i := range d.Accounts {
				account := &d.Accounts[i]
				if account.Login == accHash {
					fmt.Printf("Login %s is existing for domain %s\n", acc, dom)
					return
				}
			}
		}

		d.Accounts = append(d.Accounts, domain.Account{Login: accHash, Password: pwd64})

		err = fsutil.PersistDataToFile(dFile, d)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	} else {
		d := domain.Domain{
			Metadata: domain.Metadata{
				Entries:      1,
				LastModified: time.Now().Format(time.DateOnly),
			},
			Accounts: []domain.Account{
				{Login: accHash, Password: pwd64},
			},
		}
		err := fsutil.PersistDataToFile(dFile, d)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
