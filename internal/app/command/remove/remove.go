package remove

import (
	"errors"
	"fmt"
	"github.com/swicherwich/pwdmg/internal/app/domain"
	"github.com/swicherwich/pwdmg/internal/pkg/fsutil"
	"github.com/swicherwich/pwdmg/internal/pkg/secutil"
	"time"
)

func RemoveAccount(d, acc string) error {
	pwdmgDir, _ := fsutil.GetPwdmgDir()

	dHash := secutil.HashStr(d)
	accHash := secutil.HashStr(acc)
	dFile := fmt.Sprintf("%s/%s%s", pwdmgDir, dHash, ".json")

	if fsutil.FileExists(dFile) {
		var d domain.Domain

		if err := fsutil.ReadData(dFile, &d); err != nil {
			return err
		}

		for i := range d.Accounts {
			account := &d.Accounts[i]
			if account.Login == accHash {
				d.Metadata = domain.Metadata{
					Entries:      d.Metadata.Entries - 1,
					LastModified: time.Now().Format(time.DateTime),
				}

				d.Accounts = append(d.Accounts[:i], d.Accounts[i+1:]...)

				if err := fsutil.PersistDataToFile(dFile, d); err != nil {
					return err
				}

				return nil
			}
		}
	}

	return errors.New("no such login")
}
