package sign

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

type sm2SignData struct {
	d    string
	x    string
	y    string
	in   string
	sign string
	pub  *sm2.PublicKey
	priv *sm2.PrivateKey
}

// func initKey(data Sm2SignData) sm2SignData {
// 	data.pub = &sm2.PublicKey{
// 		Curve: sm2.P256Sm2(),
// 		X:     data.x,
// 		Y:     data.y,
// 	}

// 	data.priv = &sm2.PrivateKey{
// 		PublicKey: *pub,
// 		D:         d,
// 	}

// 	return data
// }

func Newsm2SginData(priv_key string, pub_x string, pub_y string, filename string) *sm2SignData {
	return &sm2SignData{
		d:  priv_key,
		x:  pub_x,
		y:  pub_y,
		in: filename,
	}
}

func (data *sm2SignData) InitKey() {
	xBytes, _ := hex.DecodeString(data.x)
	yBytes, _ := hex.DecodeString(data.y)
	dBytes, _ := hex.DecodeString(data.d)

	data.pub = &sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     new(big.Int).SetBytes(xBytes),
		Y:     new(big.Int).SetBytes(yBytes),
	}

	data.priv = &sm2.PrivateKey{
		PublicKey: *data.pub,
		D:         new(big.Int).SetBytes(dBytes),
	}
}

// func Sign(data sm2SignData) sm2SignData {

// 	msg, _ := ioutil.ReadFile(data.in)

// 	sign, err := data.priv.Sign(rand.Reader, msg, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	signString := hex.EncodeToString(sign)
// 	data.sign = signString

// 	return data
// }

func (data *sm2SignData) SignFile() {

	msg, _ := ioutil.ReadFile(data.in)

	sign, err := data.priv.Sign(rand.Reader, msg, nil)
	if err != nil {
		log.Println(err)
	}

	signString := hex.EncodeToString(sign)
	data.sign = signString

}

func (data sm2SignData) VerifyFile() bool {
	msg, _ := ioutil.ReadFile(data.in)

	sign, _ := hex.DecodeString(data.sign)
	ok := data.pub.Verify(msg, sign)

	return ok
}
