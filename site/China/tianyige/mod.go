package tianyige

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

// APP_ID & APP_KEY form https://gj.tianyige.com.cn/js/2.f75e590e.chunk.js
const APP_ID = "4f65a2a8247f400c8c29474bf707d680"
const APP_KEY = "G3HT5CX8FTG5GWGUUJX8B5SWJTXS1KRC"

func encrypt(pt, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	pt, err = padder.Pad(pt) // padd last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return base64.StdEncoding.EncodeToString(ct)
}

func getToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//pt := []byte(strconv.Itoa(r.Intn(900000)+100000) + strconv.FormatInt(time.Now().UnixMilli(), 10))
	pt := []byte(fmt.Sprintf("%.6d%d", r.Int31()%10000, time.Now().UnixMilli()))
	// Key size for AES is either: 16 bytes (128 bits), 24 bytes (192 bits) or 32 bytes (256 bits)
	key := []byte(APP_KEY)
	return encrypt(pt, key)
}
