// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;
pragma abicoder v2;

import "https://raw.githubusercontent.com/OpenZeppelin/openzeppelin-contracts/v3.4.0/contracts/token/ERC721/ERC721.sol";
import "https://raw.githubusercontent.com/OpenZeppelin/openzeppelin-contracts/v3.4.0/contracts/utils/Counters.sol";

contract Minter is ERC721 {

    using  Counters for Counters.Counter;
    Counters.Counter public _tokenIds;
    address public owner;

    string private nftName = "IBCCNFT";
    string private nftIdentifier = "IBCNFT";

    enum Status {MINTED, BURN_INITIATED, BURNED, REVERT_BURN}

    struct NFTDetails {
        uint256 nftId;
        address destinationAddress;
        address sourceAddress;
        string chainName;
        uint256 sourceTokenId;
        address nftContractAddress;
        address receipentAddressInSourceChain;
        string symbol;
        Status status;
        string tokenURI;
    }

    mapping(address => mapping(uint256 => NFTDetails)) public tokenIdToNFTDetailsMapping;

     // Custom Events based on your ABI
    event BurnInitiated(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, NFTDetails nftDetails, string eventType);
    event Burned(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, NFTDetails nftDetails, string eventType);
    event Minted(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, NFTDetails nftDetails, string eventType);
    event RevertedBurn(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, NFTDetails nftDetails, string eventType);




    constructor() ERC721(nftName, nftIdentifier) {
        owner = msg.sender;
    }

    function mintNFT(address _receipentAddress, string memory tokenURI) private returns (uint256) {
        // increment counter
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();
        _mint(_receipentAddress, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Sender not owner");
        _;
    }

      function startMintNFT(address sourceAddress, address destinationAddress, string memory chainName, string memory tokenURI, string memory symbol, address nftContractAddress, uint256 sTokenId) public onlyOwner returns (NFTDetails memory) {

        uint256 tokenId = mintNFT(destinationAddress, tokenURI);

        NFTDetails memory nftDetails;
        nftDetails.nftId = tokenId;
        nftDetails.destinationAddress = destinationAddress;
        nftDetails.sourceAddress = sourceAddress;
        nftDetails.chainName = chainName;
        nftDetails.sourceTokenId = sTokenId;
        nftDetails.nftContractAddress = nftContractAddress;
        nftDetails.receipentAddressInSourceChain = address(0);
        nftDetails.symbol = symbol;
        nftDetails.status = Status.MINTED;
        nftDetails.tokenURI = tokenURI;

        tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId] = nftDetails;

        emit Minted(block.timestamp, nftContractAddress, sTokenId, tokenId, nftDetails, "MINTED");
        return nftDetails;
      }

      function sendBackNFT(address nftContractAddress, uint256 sTokenId, address receiverAddress) public {
            require(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId != 0, "Token non existing");
            require(sTokenId != 0, "source token is 0");
            require(receiverAddress != address(0), "Address is 0");
            require(nftContractAddress != address(0), "Invalid nft contract address");

            // take the possession
            _transfer(msg.sender, address(this), tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId);
            tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].receipentAddressInSourceChain = receiverAddress;
            tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].status = Status.BURN_INITIATED;

            emit BurnInitiated(block.timestamp, nftContractAddress, sTokenId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId], "BURN_INITIATED");
      }

    function finalizeBurn(address nftContractAddress, uint256 sTokenId) public onlyOwner {
            require(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId != 0, "Token non existing");
            require(sTokenId != 0, "source token is 0");
            require(nftContractAddress != address(0), "Invalid nft contract address");
            require(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].status == Status.BURN_INITIATED, "Burn not initiated");
            _burn(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId);
            tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].status = Status.BURNED;

            emit Burned(block.timestamp, nftContractAddress, sTokenId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId], "BURNED");
    }

    function revertBurn(address nftContractAddress, uint256 sTokenId) public onlyOwner {
            require(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId != 0, "Token non existing");
            require(sTokenId != 0, "source token is 0");
            require(nftContractAddress != address(0), "Invalid nft contract address");
            require(tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].status == Status.BURN_INITIATED, "Burn not initiated");
            _transfer(address(this), tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].destinationAddress, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId);
            tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].status = Status.REVERT_BURN;
            tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].receipentAddressInSourceChain = address(0);

            emit RevertedBurn(block.timestamp, nftContractAddress, sTokenId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId].nftId, tokenIdToNFTDetailsMapping[nftContractAddress][sTokenId], "REVERTED_BURN");
    }


}