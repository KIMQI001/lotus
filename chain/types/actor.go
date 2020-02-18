package types

import (
	"fmt"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = fmt.Errorf("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
