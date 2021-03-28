package sign

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
)

type testSm2SignData struct {
	d    string
	x    string
	y    string
	in   string
	sign string
}

var testSignData = []testSm2SignData{
	{
		d:    "5DD701828C424B84C5D56770ECF7C4FE882E654CAC53C7CC89A66B1709068B9D",
		x:    "FF6712D3A7FC0D1B9E01FF471A87EA87525E47C7775039D19304E554DEFE0913",
		y:    "F632025F692776D4C13470ECA36AC85D560E794E1BCCF53D82C015988E0EB956",
		in:   "0102030405060708010203040506070801020304050607080102030405060708",
		sign: "30450220213C6CD6EBD6A4D5C2D0AB38E29D441836D1457A8118D34864C247D727831962022100D9248480342AC8513CCDF0F89A2250DC8F6EB4F2471E144E9A812E0AF497F801",
	},
}

func TestVerify(t *testing.T) {
	for _, data := range testSignData {

		xBytes, _ := hex.DecodeString(data.x)
		yBytes, _ := hex.DecodeString(data.y)
		pub.X = new(big.Int).SetBytes(xBytes)
		pub.Y = new(big.Int).SetBytes(yBytes)
		inBytes, _ := hex.DecodeString(data.in)
		sign, _ := hex.DecodeString(data.sign)
		result := Verify(pub, nil, inBytes, sign)
		if !result {
			t.Error("verify failed")
			return
		}
	}
}

// func TestSign(t *testing.T) {
// 	pubKeyString := []string{
// 		"FF6712D3A7FC0D1B9E01FF471A87EA87525E47C7775039D19304E554DEFE0913",
// 		"F632025F692776D4C13470ECA36AC85D560E794E1BCCF53D82C015988E0EB956",
// 	}

// 	msg, _ := hex.DecodeString("0102030405060708010203040506070801020304050607080102030405060708")

// 	fmt.Println(msg)
// 	x, _ := new(big.Int).SetString(pubKeyString[0], 16)
// 	y, _ := new(big.Int).SetString(pubKeyString[1], 16)
// 	pub := &sm2.PublicKey{
// 		Curve: sm2.P256Sm2(),
// 		X:     x,
// 		Y:     y,
// 	}

// 	sign, _ := hex.DecodeString("30450220213C6CD6EBD6A4D5C2D0AB38E29D441836D1457A8118D34864C247D727831962022100D9248480342AC8513CCDF0F89A2250DC8F6EB4F2471E144E9A812E0AF497F801")
// 	ok := pub.Verify(msg, sign)
// 	if ok != true {
// 		fmt.Printf("Verify error\n")
// 	} else {
// 		fmt.Printf("Verify ok\n")
// 	}
// }
