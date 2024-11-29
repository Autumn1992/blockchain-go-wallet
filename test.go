package main

import (
	"fmt"
	"walletserver/mnemonic"
	"walletserver/wallet/chain/trc20"
)

func main() {

	var mem, _ = mnemonic.GenerateMnemonic()

	for i := 0; i < 20; i++ {
		var address1, _ = trc20.DeriveAddressFromMnemonic(mem, 195, 0, uint32(i))
		var address2, _ = trc20.DeriveAddressFromMnemonic(mem, 195, 1, uint32(i))
		fmt.Println(address1, address2)
	}
	fmt.Println("mem", mem)
}
