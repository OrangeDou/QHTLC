package blockchain

type Block struct {
	Header   *BlockHeader
	Data     *BlockData
	Metadata *BlockMetadata
}

type BlockHeader struct {
	Number       uint64
	PreviousHash []byte
	DataHash     []byte
}

type BlockData struct {
	Transactions []*Transaction
}

type Transaction struct {
	ID             string
	ValidationCode int32
	Payload        []byte
}

type BlockMetadata struct {
	Metadata []byte
}

func (b *Block) GetBlockChainData(tid int) []byte {
	return b.Data.Transactions[tid].Payload
}
