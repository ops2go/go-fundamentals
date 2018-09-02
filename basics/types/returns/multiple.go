package returns

import (
	"fmt"
)

func FullName(fname, lname string) (string, string) {
	return fmt.Sprint(fname), fmt.Sprint(lname)

}
