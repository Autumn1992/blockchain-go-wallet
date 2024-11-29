package mnemonic

import (
	"github.com/tyler-smith/go-bip39"
	"walletserver/log"
)

func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256) // 128位熵
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)

	if err != nil {
		log.Error("生成助记词失败", err)
	} else {
		if !ValidateMnemonic(mnemonic) {
			return GenerateMnemonic()
		}
	}
	return mnemonic, err
}

func ValidateMnemonic(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}
