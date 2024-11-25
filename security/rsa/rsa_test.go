package rsa

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestKeyGenerate(t *testing.T) {
	convey.Convey("生成密钥对", t, func() {
		keyPair, err := GenerateKeyPair()
		convey.So(err, convey.ShouldBeNil)
		convey.So(keyPair, convey.ShouldNotBeNil)

		t.Log(string(keyPair.publicKey.ToPemMust()))
		t.Log(string(keyPair.privateKey.ToPemMust()))
	})

}
