package sign

import (
	"fmt"
	"testing"
)

var Data = sm2SignData{
	d:    "ABA58D6570E4E9C0C071861A015F4FFAE6FCFE2011354285CD24A1959E7EB0E3",
	x:    "59B2898C1598711489AD650A930A3FB96C4586E191C7085B79DC3D4FA2D94970",
	y:    "4686BA4EFC4AEAF247FAA8B662F8D0B5650A4D98EBCF77D1C795F42D268F8D78",
	in:   "abc.txt",
	sign: "30450220345C574FA146B2230970CDCF1F2B66C3D5875E455F08263DE4F36EDD51198ABC022100CF80E9EEC53922ABFDB1CE99C3B3F4A5C32AB322EEF87CD0F896191C75AFC8ED",
}

func TestSignTool(t *testing.T) {
	testSignData := Newsm2SginData(Data.d, Data.x, Data.y, Data.in)
	testSignData.InitKey()
	// testSignData.SignFile()
	testSignData.sign = Data.sign
	fmt.Println(testSignData.sign)

	ok := testSignData.VerifyFile()
	if ok != true {
		fmt.Printf("Verify error\n")
	} else {
		fmt.Printf("Verify ok\n")
	}
}
