// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;
pragma abicoder v2;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721Metadata.sol";



contract Locker {

    enum LockState {LOCK_INITIATED, LOCKED, UNLOCKED, WITHDRAWN}

    address public owner;

    struct LockedNFTDetails {
        uint256 tokenId;
        address nftContractAddress;
        address owner;
        address receipentInDestinationChain;
        string destinationChainName;
        LockState lockState;
        string symbol;
        string tokenURI;
    }

    // mapping from contract Address and tokenId to LockedNFT Details
    mapping (address=>mapping(uint256=>LockedNFTDetails)) public tokenIdToNFTMap;
    mapping(address=>address[]) private userAddressToContractAddressMapping;
    mapping(address=>mapping(address=>uint256[])) userAndContractAddressToTokensMapping;

    // events
    event LockedNFT(uint256 indexed _timestamp, LockedNFTDetails lockedNFTDetails, string status);
    event UnlockedNFT(uint256 indexed _timestamp, LockedNFTDetails lockedNFTDetails, string status);
    event WithdrawnNFT(uint256 indexed _timestamp, LockedNFTDetails lockedNFTDetails, string status);

    // methods
    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function sendNFT(uint256 tokenId, address nftContractAddress, address receipentAddress, string memory destinationChainName) public {
        require(nftContractAddress != address(0), "Invalid contract address");
        require(receipentAddress != address(0), "Receipent address is 0");
        require(tokenId != 0, "Token Id is 0");
        require(IERC721(nftContractAddress).ownerOf(tokenId) == msg.sender, "sender not owner");
        require(tokenId != 0, "token id is 0");
        LockedNFTDetails memory lockedNFTDetails;
        lockedNFTDetails.tokenId = tokenId;
        lockedNFTDetails.nftContractAddress = nftContractAddress;
        lockedNFTDetails.owner = msg.sender;
        lockedNFTDetails.receipentInDestinationChain = receipentAddress;
        lockedNFTDetails.destinationChainName = destinationChainName;
        lockedNFTDetails.lockState = LockState.LOCK_INITIATED;
        lockedNFTDetails.symbol = IERC721Metadata(nftContractAddress).symbol();
        lockedNFTDetails.tokenURI = IERC721Metadata(nftContractAddress).tokenURI(tokenId);

        tokenIdToNFTMap[nftContractAddress][tokenId] = lockedNFTDetails;

        // transfer the ownership
        // need approval from user before transferring.
        IERC721(nftContractAddress).transferFrom(msg.sender, address(this), tokenId);

        bool contractExists = checkIfNFTContractExists(msg.sender, nftContractAddress);
        if(!contractExists) {
            userAddressToContractAddressMapping[msg.sender].push(nftContractAddress);
        }

        bool tokenExists = checkIfTokenExists(msg.sender, nftContractAddress, tokenId);
        if(!tokenExists) {
            userAndContractAddressToTokensMapping[msg.sender][nftContractAddress].push(tokenId);
        }

        emit LockedNFT(block.timestamp, lockedNFTDetails, "LOCK_INITIATED");
    }

    function checkIfTokenExists(address userAddress, address nftContractAddress, uint256 tokenId) private view returns (bool) {
        for(uint256 i = 0; i < userAndContractAddressToTokensMapping[userAddress][nftContractAddress].length; i++) {
            if(userAndContractAddressToTokensMapping[userAddress][nftContractAddress][i] == tokenId) {
                return true;
            }
        }

        return false;
    }

    function checkIfNFTContractExists(address userAddress, address nftContractAddress) private view returns (bool) {
        for(uint256 i = 0; i < userAddressToContractAddressMapping[userAddress].length; i++) {
            if(userAddressToContractAddressMapping[userAddress][i] == nftContractAddress) {
                return true;
            }
        }

        return false;
    }

    // finalize locking
    function finalizeLocking(address nftContractAddress, uint256 tokenId) public onlyOwner {
        require(tokenId != 0, "Token Id is 0");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].tokenId != 0, "Non existing");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].lockState == LockState.LOCK_INITIATED, "NFT is not in lock initiate state");


        tokenIdToNFTMap[nftContractAddress][tokenId].lockState = LockState.LOCKED;

        emit LockedNFT(block.timestamp, tokenIdToNFTMap[nftContractAddress][tokenId], "LOCKED");        
    }

    // finalize unlocking
    function finalizeUnlocking(address nftContractAddress, address receipentAddress, uint256 tokenId) public onlyOwner {
        require(tokenId != 0, "Token Id is 0");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].tokenId != 0, "Non existing");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].lockState != LockState.WITHDRAWN, "NFT is withdrawn");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].lockState != LockState.UNLOCKED, "NFT is already unlocked");


        tokenIdToNFTMap[nftContractAddress][tokenId].owner = receipentAddress;
        tokenIdToNFTMap[nftContractAddress][tokenId].lockState = LockState.UNLOCKED;

        emit UnlockedNFT(block.timestamp, tokenIdToNFTMap[nftContractAddress][tokenId], "UNLOCKED");
    }

    // fallback unlocking
    function fallbackUnlocking(address nftContractAddress, address receipentAddress, uint256 tokenId) public onlyOwner {
        require(tokenId != 0, "Token Id is 0");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].tokenId != 0, "Non existing");
        tokenIdToNFTMap[nftContractAddress][tokenId].owner = receipentAddress;
        tokenIdToNFTMap[nftContractAddress][tokenId].lockState = LockState.UNLOCKED;

        emit UnlockedNFT(block.timestamp, tokenIdToNFTMap[nftContractAddress][tokenId], "FALLBACK_UNLOCKING");
    }

    // withdraw NFT
    function withdrawNFT(address nftContractAddress, uint256 tokenId) public {
        require(tokenId != 0, "Token Id is 0");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].tokenId != 0, "Non existing");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].owner == msg.sender, "Non meant for sender");
        require(tokenIdToNFTMap[nftContractAddress][tokenId].lockState == LockState.UNLOCKED, "NFT is not unlocked");
        IERC721(nftContractAddress).transferFrom(address(this), msg.sender, tokenId);
        tokenIdToNFTMap[nftContractAddress][tokenId].lockState = LockState.WITHDRAWN;
        emit WithdrawnNFT(block.timestamp, tokenIdToNFTMap[nftContractAddress][tokenId], "WITHDRAWN");
    }

    function returnAllMyContracts() public view returns (address[] memory) {
        return userAddressToContractAddressMapping[msg.sender];
    }

    function returnAllMyTokens(address nftContractAddress) public view returns (uint256[] memory) {
        return userAndContractAddressToTokensMapping[msg.sender][nftContractAddress];
    }

}