// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;

import "https://raw.githubusercontent.com/OpenZeppelin/openzeppelin-contracts/v3.4.0/contracts/token/ERC721/ERC721.sol";
import "https://raw.githubusercontent.com/OpenZeppelin/openzeppelin-contracts/v3.4.0/contracts/utils/Counters.sol";

contract DogNFT is ERC721 {
    using  Counters for Counters.Counter;
    Counters.Counter public _tokenIds;

    string private nftName = "DOGGIES";
    string private nftIdentifier = "DOG";

    constructor() ERC721(nftName, nftIdentifier) {}

    function mintNFT(address _receipentAddress, string calldata tokenURI) public returns (uint256) {
        // increment counter
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();
        _mint(_receipentAddress, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }
}