package gossip

import (
	"fmt"
	"testing"
)

func Test_whoAmI(t *testing.T) {
	gotAddresses := whoAmI()
	fmt.Println(gotAddresses)
	t.Fail()

}
