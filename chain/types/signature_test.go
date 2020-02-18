package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/specs-actors/actors/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
