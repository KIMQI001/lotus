package retrievaladapter

import (
	"context"
	"github.com/filecoin-project/lotus/lib/sharedutils"

	"github.com/filecoin-project/go-fil-components/retrievalmarket"
	retrievaladdress "github.com/filecoin-project/go-fil-components/shared/address"
	retrievaltoken "github.com/filecoin-project/go-fil-components/shared/tokenamount"
	retrievaltypes "github.com/filecoin-project/go-fil-components/shared/types"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/storage/sectorblocks"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
)

type retrievalProviderNode struct {
	sectorBlocks *sectorblocks.SectorBlocks
	full         api.FullNode
}

// NewRetrievalProviderNode returns a new node adapter for a retrieval provider that talks to the
// Lotus Node
func NewRetrievalProviderNode(sectorBlocks *sectorblocks.SectorBlocks, full api.FullNode) retrievalmarket.RetrievalProviderNode {
	return &retrievalProviderNode{sectorBlocks, full}
}

func (rpn *retrievalProviderNode) GetPieceSize(pieceCid []byte) (uint64, error) {
	asCid, err := cid.Cast(pieceCid)
	if err != nil {
		return 0, err
	}
	return rpn.sectorBlocks.GetSize(asCid)
}

func (rpn *retrievalProviderNode) SealedBlockstore(approveUnseal func() error) blockstore.Blockstore {
	return rpn.sectorBlocks.SealedBlockstore(approveUnseal)
}

func (rpn *retrievalProviderNode) SavePaymentVoucher(ctx context.Context, paymentChannel retrievaladdress.Address, voucher *retrievaltypes.SignedVoucher, proof []byte, expectedAmount retrievaltoken.TokenAmount) (retrievaltoken.TokenAmount, error) {
	localVoucher, err := sharedutils.FromSharedSignedVoucher(voucher)
	if err != nil {
		return retrievaltoken.FromInt(0), err
	}
	added, err := rpn.full.PaychVoucherAdd(ctx, sharedutils.FromSharedAddress(paymentChannel), localVoucher, proof, sharedutils.FromSharedTokenAmount(expectedAmount))
	return sharedutils.ToSharedTokenAmount(added), err
}