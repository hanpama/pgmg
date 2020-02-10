package testutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// TestSnapshot tests the bytes equals the data written on fp file.
func TestSnapshot(got []byte, fp string) (err error) {
	if _, err = os.Stat(fp); os.IsNotExist(err) {
		if err = ioutil.WriteFile(fp, got, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		expected, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}
		if string(got) != string(expected) {
			if err = ioutil.WriteFile(fp+".expected", expected, os.ModePerm); err != nil {
				panic(err)
			}
			if err = ioutil.WriteFile(fp, got, os.ModePerm); err != nil {
				panic(err)
			}
			return fmt.Errorf("Unexpected inspection result: %s", string(got))
		}
	}
	return nil
}
