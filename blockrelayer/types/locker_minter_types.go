package types

type LockerNFTDetails struct {
	TokenId                     uint   `json:"TokenId"`
	NftContractAddress          string `json:"NftContractAddress"`
	Owner                       string `json:"Owner"`
	ReceipentInDestinationChain string `json:"ReceipentInDestinationChain"`
	DestinationChainName        string `json:"DestinationChainName"`
	LockState                   uint   `json:"LockState"`
	Symbol                      string `json:"Symbol"`
	TokenURI                    string `json:"TokenURI"`
}

type MinterNFTDetails struct {
	NftId                         uint   `json:"NftId"`
	DestinationAddress            string `json:"DestinationAddress"`
	SourceAddress                 string `json:"SourceAddress"`
	ChainName                     string `json:"ChainName"`
	SourceTokenId                 uint   `json:"SourceTokenId"`
	DestinationTokenId            uint   `json:"DestinationTokenId"`
	ContractAddress               string `json:"ContractAddress"`
	ReceipentAddressInSourceChain string `json:"ReceipentAddressInSourceChain"`
	Symbol                        string `json:"Symbol"`
	TokenURI                      string `json:"TokenURI"`
	Status                        uint   `json:"Status"`
}
