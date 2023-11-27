package peda

import (
	"fmt"
	"testing"

	"github.com/whatsauth/watoken"
)

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println("privateKey" + privateKey)
	fmt.Println("publicKey" + publicKey)
}

func TestEncode(t *testing.T) {
	privateKey := "8c92e028bf9dc2ad8d7244a08b611220845a7f43fb7986a6af805cbda811b96c02a51c4853c18c4a4bbedd29661618e7bec9041d4063caec8cc84c601c20281b"
	userid := "mubarok"

	tokenstring, err := watoken.Encode(userid, privateKey)
	fmt.Println("error : ", err)
	fmt.Println("token : ", tokenstring)
}

func TestDecode(t *testing.T) {
	publicKey := "02a51c4853c18c4a4bbedd29661618e7bec9041d4063caec8cc84c601c20281b"

	tokenstring := "v4.public.eyJleHAiOiIyMDIzLTExLTIzVDExOjA5OjQ0KzA3OjAwIiwiaWF0IjoiMjAyMy0xMS0yM1QwOTowOTo0NCswNzowMCIsImlkIjoiaWJyb2hpbSIsIm5iZiI6IjIwMjMtMTEtMjNUMDk6MDk6NDQrMDc6MDAifRrPRTXnMvVDjYN_Eb27_GYyovqHZCwI8ds5Rk7RxM2OyqiujCUzsTfZy1PlaAl7kv7wkQk9ST0oFJ3WD2Ih-Qg"
	body := watoken.DecodeGetId(publicKey, tokenstring)
	fmt.Println("isi : ", body)
}

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnectionTest("mongodb+srv://fitrahaliakbar:fitrah354@cluster0.vb2aato.mongodb.net/", "GIS")
	datagedung := GeoIntersects(mconn, 106.68136271186847, -6.193543362832116)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnectionTest("mongodb+srv://fitrahaliakbar:fitrah354@cluster0.vb2aato.mongodb.net/", "GIS")
	coordinates := [][][]float64{
		{
			{106.68136271186847, -6.193543362832116},
			{106.68132861056944, -6.19370318772549},
			{106.68150155287583, -6.193732246791583},
			{106.68149911706826, -6.19357000031799},
			{106.68136271186847, -6.193543362832116},
		},
	}
	datagedung := GeoWithin(mconn, coordinates)
	fmt.Println(datagedung)
}
