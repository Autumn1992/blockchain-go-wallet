package trc20

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func DeriveAddressFromMnemonic(mnemonic string, coinType, account, index uint32) (string, error) {
	// Step 1: 使用助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")

	// Step 2: 创建主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", err
	}

	// Step 3: 按 BIP44 路径派生子密钥
	// 路径: m/44'/<coinType>'/<account>'/0/<index>
	path := fmt.Sprintf("44'/%d'/%d'/0/%d", coinType, account, index)
	key := masterKey

	for _, part := range strings.Split(path, "/") {
		hardened := strings.HasSuffix(part, "'")
		segment, err := strconv.ParseUint(strings.TrimSuffix(part, "'"), 10, 32)
		if err != nil {
			return "", err
		}
		if hardened {
			segment += uint64(bip32.FirstHardenedChild)
		}
		key, err = key.NewChildKey(uint32(segment))
		if err != nil {
			return "", err
		}
	}
	// Step 4: 从派生密钥生成 Tron 地址
	pubKey := key.PublicKey().Key
	address := TronAddressFromPublicKey(pubKey)
	return address, nil
}

func TronAddressFromPublicKey(pubKey []byte) string {
	// Keccak256 哈希 + Tron 地址转换逻辑
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey[1:])        // 跳过首字节（04，表示未压缩的公钥）
	address := hash.Sum(nil)[12:] // 获取最后20字节
	return "41" + hex.EncodeToString(address)
}
