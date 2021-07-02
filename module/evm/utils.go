package evm

import (
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// GetTxHash calculates the tx hash
func (ec evmClient) GetTxHash(signedTx *ethcore.Transaction) (txHash ethcmn.Hash, err error) {
	v, r, s := signedTx.RawSignatureValues()
	tx := evmtypes.MsgEthereumTx{
		Data: evmtypes.TxData{
			AccountNonce: signedTx.Nonce(),
			Price:        signedTx.GasPrice(),
			GasLimit:     signedTx.Gas(),
			Recipient:    signedTx.To(),
			Amount:       signedTx.Value(),
			Payload:      signedTx.Data(),
			V:            v,
			R:            r,
			S:            s,
		},
	}

	txBytes, err := authcli.GetTxEncoder(ec.GetCodec())(tx)
	if err != nil {
		return
	}

	txHash = ethcmn.BytesToHash(tmhash.Sum(txBytes))
	return
}
