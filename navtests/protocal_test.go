package navtests

import (
	"testing"
	"github.com/NAVCoin/navd/wire"
)

func TestProtocolsSettings(t *testing.T) {

	var NavProtocolVersion uint32
	NavProtocolVersion = 70020

	 if wire.ProtocolVersion != NavProtocolVersion {
	 	t.Errorf("Incorrect protocal version expected %d for %d", NavProtocolVersion, wire.ProtocolVersion)
	 }


}