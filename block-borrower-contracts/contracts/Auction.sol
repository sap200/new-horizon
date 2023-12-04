// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;
pragma abicoder v2;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721Metadata.sol";


contract Auction {
     enum AuctionState {NOT_STARTED, RUNNING, FINALIZED, CANCELLED, ENDED}

        struct AuctionItem {
            uint256 loanId;
            address pledger;
            uint256 tokenId;
            uint256 auctionId;
            uint256 auctionStartTime;
            uint256 repaymentAmount;
            uint256 bankToWhichItisPledged;
            address nftContractAddress;
            address bankAddress;
            address bankAdmin;
            AuctionState auctionState;
            uint256 highestBindingBid;
            address payable highestBidder;
            string symbol;
        }

        // Maps loan ID to Auction
        mapping(uint256 => AuctionItem) public auction; 
        // Maps auctionId to all bidder's list
        mapping(uint256 => address[]) public bidderList; 
        // Maps auction ID and bidder address to bid amount
        mapping(uint256 => mapping(address => uint256)) public bidsForAuction; 
        // Maps auction ID and address to blacklist status
        mapping(uint256 => mapping(address => bool)) public blackListed; 
        // Maps loan ID to an array of auction IDs
        mapping(uint256 => uint256[]) public loanIdToAllAuctionIdMappings;

        // State variables
        uint256 public auctionCounter = 1;
        uint256 public daysAfterWhichAuctionShouldEnd = 20 minutes;
        uint256 public minimumAmountToPlaceABid = 1 ether;

        function startAuction(uint256 bankId, uint256 loanId, address pledger, uint256 tokenId, uint256 repaymentAmount, address nftContractAddress, address bankAddress, address bankAdmin) public returns (uint256) {
            require(loanId != 0, "Invalid loan Id");
            require(bankId != 0, "Invalid bank Id");
            require(IERC721(nftContractAddress).ownerOf(tokenId) == msg.sender, "You do not own the NFT");
            require(auction[loanId].auctionState != AuctionState.RUNNING, "Already an auction is in progress !!");

            AuctionItem memory auctionItem;
            auctionItem.loanId = loanId;
            auctionItem.pledger = pledger;
            auctionItem.tokenId = tokenId;
            auctionItem.auctionId = auctionCounter;
            auctionCounter++;
            auctionItem.auctionStartTime = block.timestamp;
            auctionItem.repaymentAmount = repaymentAmount;
            auctionItem.bankToWhichItisPledged = bankId;
            auctionItem.nftContractAddress = nftContractAddress;
            auctionItem.bankAddress = bankAddress;
            auctionItem.bankAdmin = bankAdmin;
            auctionItem.auctionState = AuctionState.RUNNING;
            auctionItem.highestBindingBid = repaymentAmount;
            auctionItem.highestBidder = payable(bankAddress);
            auctionItem.symbol = IERC721Metadata(nftContractAddress).symbol();

            auction[auctionItem.loanId] = auctionItem;
            loanIdToAllAuctionIdMappings[loanId].push(auctionItem.auctionId);

            return auctionItem.auctionId;
        }

        function blackListUsersFromAuction(uint256 loanId, uint256 auctionId, address[] memory list) public {
            require(loanId != 0, "Loan Id invalid");
            require(auction[loanId].auctionState == AuctionState.RUNNING, "Auction is not running");
            require(msg.sender == auction[loanId].bankAddress, "Contract can blacklist");
            require(auction[loanId].auctionId == auctionId);

            // set all of them blacklisted
            // bank lenders cannot bid to increase the price
            for(uint256 i = 0; i < list.length; i++) {
                blackListed[auctionId][list[i]] = true;
            }  
        }

        function finalizeAuctionBeforeTime(uint256 loanId) public returns (uint256, address) {
            require(auction[loanId].auctionState == AuctionState.RUNNING, "Auction is not running");
            require(msg.sender == auction[loanId].bankAddress, "Only contract can finalize the auction");
            require(block.timestamp < (auction[loanId].auctionStartTime + daysAfterWhichAuctionShouldEnd), "Auction has ended ! Cannot finalize");
            auction[loanId].auctionState = AuctionState.FINALIZED;

            return takeAuctionDecision(loanId);
        }

        function cancelAuctionBeforeTime(uint256 loanId) public returns (uint256, address) {
            require(auction[loanId].auctionState == AuctionState.RUNNING, "Auction is not running");
            require(msg.sender == auction[loanId].bankAddress, "Only contract can cancel the auction");
            require(block.timestamp < (auction[loanId].auctionStartTime + daysAfterWhichAuctionShouldEnd), "Auction has ended ! Cannot cancel");
            auction[loanId].auctionState = AuctionState.CANCELLED;

            return (0, auction[loanId].bankAddress);
        }

        function checkIfAuctionHasExpired(uint256 loanId) public returns (uint256, address) {
            require(auction[loanId].auctionState == AuctionState.RUNNING, "Auction is in other states");
            require(msg.sender == auction[loanId].bankAddress, "Only contract can check if auction has ended");
            if(block.timestamp > (auction[loanId].auctionStartTime + daysAfterWhichAuctionShouldEnd) ) {
                auction[loanId].auctionState = AuctionState.ENDED;
                return takeAuctionDecision(loanId);
            }

            return (0, address(0));
        }


        function takeAuctionDecision(uint256 loanId) private returns (uint256, address) {
            if(auction[loanId].highestBidder == auction[loanId].bankAddress) {
                // nft belongs to bank, don't do anything instead return 0 and banker's address
                return (0, auction[loanId].bankAddress);

            } else {
                // transfer the NFT to highest bidder
                IERC721(auction[loanId].nftContractAddress).transferFrom(auction[loanId].bankAddress, auction[loanId].highestBidder, auction[loanId].tokenId);
                // reduce his bids so he cannot withdraw any money
                bidsForAuction[auction[loanId].auctionId][auction[loanId].highestBidder] = 0;
                // send the money to the contract
                payable(auction[loanId].bankAddress).transfer(auction[loanId].highestBindingBid);

                // return amount sent and the address where nft is sent.
                return (auction[loanId].highestBindingBid, auction[loanId].highestBidder);

            }
        }

        function withdrawMoney(uint256 loanId) public payable {
            uint256 totalMoney = 0;
            for(uint256 i = 0; i < loanIdToAllAuctionIdMappings[loanId].length; i++) {
                uint256 currentAuctionId = loanIdToAllAuctionIdMappings[loanId][i];
                if (auction[loanId].auctionState == AuctionState.RUNNING && auction[loanId].auctionId == currentAuctionId) {
                    continue;
                }

                totalMoney += bidsForAuction[currentAuctionId][msg.sender];
                // make his balance for considered auctions 0
                bidsForAuction[currentAuctionId][msg.sender] = 0;
            }

            require(totalMoney >= 1 ether, "Total claimable balance should be >= 1 ether");

            // transfer total money to user
            payable(msg.sender).transfer(totalMoney);
        }

        function getAllAuctionIdsForLoanId(uint256 loanId) public view returns (uint256[] memory) {
            return loanIdToAllAuctionIdMappings[loanId];
        }

        function getAllBiddersList(uint256 auctionId) public view returns (address[] memory) {
            return bidderList[auctionId];
        }

        function getUsersBidValueForAuctionId(uint256 auctionId) public view returns (uint256) {
            return bidsForAuction[auctionId][msg.sender];
        }

        function placeBidForAuction(uint256 loanId) public payable {
            require(auction[loanId].auctionState == AuctionState.RUNNING, "Auction State is not running");
            uint256 cAuctionId = auction[loanId].auctionId;
            require(blackListed[cAuctionId][msg.sender] == false, "you are blacklisted");
            require(block.timestamp < auction[loanId].auctionStartTime + daysAfterWhichAuctionShouldEnd, "Auction ended, please press end auction button !!");
            uint256 currentBidValue = bidsForAuction[cAuctionId][msg.sender] + msg.value;
            require(currentBidValue > (auction[cAuctionId].highestBindingBid + minimumAmountToPlaceABid), "bid should be greater than highest binding bid");
            // set highest binding bid and bidder
            auction[loanId].highestBindingBid = currentBidValue;
            auction[loanId].highestBidder = msg.sender;

            // increase his bids
            bidsForAuction[cAuctionId][msg.sender] = currentBidValue;

            // append him to bidders list if already not existing
            bool alreadyBidded = bidderAlreadyExistsInBiddersList(cAuctionId, msg.sender);
            if(!alreadyBidded) {
                bidderList[cAuctionId].push(msg.sender);
            }
        }

        function bidderAlreadyExistsInBiddersList(uint256 auctionId, address bidder) private view returns (bool) {
            for(uint256 i = 0; i < bidderList[auctionId].length; i++) {
                if(bidderList[auctionId][i] == bidder) {
                    return true;
                }
            }

            return false;
        }
}