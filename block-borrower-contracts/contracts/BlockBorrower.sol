// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;
pragma abicoder v2;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/token/ERC721/IERC721Metadata.sol";
import "./Auction.sol";


contract BlockBorrower {
    enum AuctionState {NOT_STARTED, RUNNING, FINALIZED}

    struct Lender {
        bool status;
        uint256 balance;
    }

     // NFT Struct
    struct NFT {
        uint256 loanId;
        address pledger;
        uint256 tokenId;
        bool canWithdraw;
        uint256 loanAmount;
        uint256 interestRate;
        uint256 tenure;
        uint256 repaymentAmount;
        uint256 interest;
        uint256 lateChargesIncurred;
        uint256 amountRepaid;
        uint256 bankToWhichItisPledged;
        address nftContractAddress;
        bool isActiveLoan;
        address lastKnownCurrentOwner;
        uint256 amountReturnedToUser;
        AuctionState auctionState;
        uint256 auctionId;
        uint256 amountEarnedFromAuction;
        bool canBeAuctioned;
        string nftSymbol;
        bool hasClaimed;
    }

    // RPS Struct
    struct RepaymentSchedule {
        uint256 rpsIndex;
        uint256 amount;
        uint256 dueDate;
        uint256 amountRepaid;
    }

      // BANK Struct
      struct Bank {
        address admin;
        string bankName;
        mapping(address=>Lender) lenders;
        address[] keys;
        address[] customers;
        uint256 interestRate;
        uint256 tenure;
        uint256 totalBalance;
        bool functional;
        uint256 loanPercentage;
        uint32 defaultingLimit;
        uint256 defaultingTimeLimit;
    }

     // Mappings
    mapping(uint256 => Bank) public bankList;
    mapping(address => uint256) public lenderAddressToBankMapping;
    mapping(uint256 => address) public loanIdToAddressMapping;
    mapping(uint256 => uint256) private loanIdToBankIdMapping;
    mapping(address => uint256[]) allMyLoanMapping;
    mapping(address => uint256[]) allMyBankMapping;
    mapping(uint256 => mapping(address=>NFT[])) private nftDeposits;
    mapping(uint256 => RepaymentSchedule[]) repaymentSchedules;


    // constants
    uint256 public constant creationThreshold = 45 ether;
    uint256 public constant joiningThreshold = 20 ether;
    uint256 public constant latePaymentCharge = 1 ether;
    uint256 public constant daysAfterWhichAuctionShouldEnd = 7 days;
    uint256 public constant minimumDepositAmount = 1 ether;
    uint256 private constant oneMonth = 30 days;


    uint256 private bankCounter = 1;
    uint256 private loanCounter = 1;


    Auction auctionSmartContract;

    constructor (address _auctionSmartContract) {
        auctionSmartContract = Auction(_auctionSmartContract);
    }


    // ------------------------------------------------------------------ BANK FUNCTIONALITIES ----------------------------------------------------------------------------


    function createBank(string memory bankName, uint256 _interestRate, uint256 _tenure, uint256 _loanPercentage, uint32 _defaultingLimit, uint32 _defaultingTimeLimit) public payable {
        require(_interestRate != 0, "Interest rate is 0");
        require(_tenure != 0, "Tenure is 0");
        require(_loanPercentage != 0, "Loan percentage is 0");
        require(_defaultingLimit != 0, "Defaulting Limit is 0");
        require(_defaultingTimeLimit != 0, "Defaulting time limit is 0");
        require(lenderAddressToBankMapping[msg.sender] == 0, "Address is already an admin or lender of the bank");
        require(msg.value >= creationThreshold, "Need a minimum of 450 XRP to create bank");

        Bank storage bank = bankList[bankCounter];
        bankCounter++;
        bank.admin = msg.sender;
        bank.bankName = bankName;
        bank.lenders[msg.sender] = Lender({status:true, balance:msg.value});
        bank.keys.push(msg.sender);
        bank.interestRate = _interestRate;
        bank.tenure = _tenure;
        bank.totalBalance = msg.value;
        bank.functional = true;
        bank.loanPercentage = _loanPercentage;
        bank.defaultingLimit = _defaultingLimit;
        bank.defaultingTimeLimit = _defaultingTimeLimit;

        lenderAddressToBankMapping[msg.sender] = bankCounter-1;
    }

    function joinBank(uint256 bankId) public payable {
        require(lenderAddressToBankMapping[msg.sender] == 0, "Address is already an admin or lender of the bank");
        require(msg.value >= joiningThreshold, "Need a minimum of 200 XRP to create bank");
        require(bankId != 0, "Invalid bankId");
        require(bankList[bankId].functional, "Bank is not functional");

        // add him to the bank
        bankList[bankId].lenders[msg.sender] = Lender({status:true, balance:msg.value});
        bankList[bankId].totalBalance += msg.value;
        bool lenderAlreadyExists = checkIfLenderExistsInBank(bankId, msg.sender);
        if(!lenderAlreadyExists) {
            bankList[bankId].keys.push(msg.sender);
        }

        // append in lenders mapping
        lenderAddressToBankMapping[msg.sender] = bankId;
    }

    function checkIfLenderExistsInBank(uint256 bankId, address lender) private view returns (bool) {
        for(uint256 i = 0; i < bankList[bankId].keys.length; i++) {
            if(bankList[bankId].keys[i] == lender) {
                return true;
            }
        }

        return false;
    }

    function depositToBank(uint256 bankId) public payable {
        require(bankId != 0, "Invalid bankId");
        require(bankList[bankId].functional, "Bank is not functional");
        require(lenderAddressToBankMapping[msg.sender] == bankId, "Lender not part of the bank");
        require(bankList[bankId].lenders[msg.sender].status == true, "Not an active member");
        require(msg.value >= minimumDepositAmount, "Should send minimum 1 xrp");
        bankList[bankId].totalBalance += msg.value;
        bankList[bankId].lenders[msg.sender].balance += msg.value;
    }

    function withdrawBalance(uint256 amount) public {
        // determine bank
        require(amount != 0, "Cannot withdraw 0 XRP");
        uint256 bankId = lenderAddressToBankMapping[msg.sender];
        require(bankId != 0, "No Bank found");
        require(bankList[bankId].functional, "Bank is not functional");

        // check bank balance and lenders balance
        require(bankList[bankId].lenders[msg.sender].balance >= amount, "Insufficient fund");

        // pay the user
        payable(msg.sender).transfer(amount);

        bankList[bankId].lenders[msg.sender].balance -= amount;
        bankList[bankId].totalBalance -= amount;
    }

    function setInterestRate(uint256 bankId, uint256 rate) public {
        require(msg.sender == bankList[bankId].admin, "Only admins can change interest rate");
        require(rate > 0, "Interest rate cannot be 0");
        require(bankId != 0, "No Bank found");
        require(bankList[bankId].functional, "Bank is not functional");
        bankList[bankId].interestRate = rate;
    }

    function setTenure(uint256 bankId, uint256 tenure) public {
        require(msg.sender == bankList[bankId].admin, "Only admins can change interest rate");
        require(tenure > 0, "Tenure cannot be 0");
        require(bankId != 0, "No Bank found");
        require(bankList[bankId].functional, "Bank is not functional");
        bankList[bankId].tenure = tenure;
    }

    function LeaveBank() public {
        uint256 bankId = lenderAddressToBankMapping[msg.sender];
        require(bankId != 0, "Sender not part of any bank");
        require(bankList[bankId].functional, "Bank is non functional");
        require(bankList[bankId].lenders[msg.sender].status, "Inactive lender");
        require(bankList[bankId].lenders[msg.sender].balance <= 2 ether, "Bank balance must be less than 2 ether to leave");

        // pay any remaining dues to the person.
        bankList[bankId].totalBalance -= bankList[bankId].lenders[msg.sender].balance;
        bankList[bankId].lenders[msg.sender].balance = 0;
        bankList[bankId].lenders[msg.sender].status = false;
        delete lenderAddressToBankMapping[msg.sender];
        delete bankList[bankId].lenders[msg.sender];
        removeFromKeysBank(bankId, msg.sender);
    }

    function removeFromKeysBank(uint256 bankId, address addr) private {
        uint256 lastIndex = bankList[bankId].keys.length-1;
        uint256 index = bankList[bankId].keys.length-1;
        bool found = false;
        for(uint256 i = 0; i < bankList[bankId].keys.length; i++) {
            if(bankList[bankId].keys[i] == addr) {
                index = i;
                found = true;
                break;
            }
        }

        if (found) {
            if(index != lastIndex) {
                bankList[bankId].keys[index] = bankList[bankId].keys[lastIndex];
            }

            bankList[bankId].keys.pop();
        }
    }

    function closeBank(uint256 bankId) public {
        require(bankId != 0, "Invalid bank id");
        require(lenderAddressToBankMapping[msg.sender] == bankId, "Msg sender not lender or admin of bank");
        require(bankList[bankId].admin == msg.sender, "Msg sender not bank admin");
        // TODO: CHECK IF BANK HAS ANY ACTIVE LOAN OR A LOAN RUNNING IN AUCTION
        bool anyActiveLoanExists = checkIfActiveLoanExists(bankId);
        require(anyActiveLoanExists == false, "Bank has an active loan");
        require(bankList[bankId].totalBalance <= 2 ether, "Amount should be less than <= 2 ether");
        require(bankList[bankId].keys.length == 1, "Multiple members are still present");
        require(bankList[bankId].keys[0] == msg.sender);
        // deduct total balance
        bankList[bankId].totalBalance = 0;
        // deduct lender balance
        delete bankList[bankId].lenders[msg.sender];
        delete lenderAddressToBankMapping[msg.sender];
        bankList[bankId].keys.pop();
        bankList[bankId].functional = false;
    }

    function checkIfActiveLoanExists(uint256 bankId) public view returns (bool) {
        for(uint256 i = 0; i < bankList[bankId].customers.length; i++) {
            address custAddress = bankList[bankId].customers[i];
            for (uint256 j = 0; j < nftDeposits[bankId][custAddress].length; j++) {
                if (nftDeposits[bankId][custAddress][j].isActiveLoan || nftDeposits[bankId][custAddress][j].auctionState == AuctionState.RUNNING) {
                    return true;
                }
            }
        }

        return false;
    }

    function getAllCustomersOfBank(uint256 bankId) public view returns (address[] memory) {
        return bankList[bankId].customers;
    }

    function getAllLendersOfBank(uint256 bankId) public view returns (address[] memory) {
        return bankList[bankId].keys;
    }

    function getAllNftDepositsOfBank(uint256 bankId, address user) public view returns (NFT[] memory) {
        return nftDeposits[bankId][user];
    }

    function getMyBankBalance(uint256 bankId, address myAddress) public view returns (uint256) {
        return bankList[bankId].lenders[myAddress].balance;
    }

    function getStatusOfLender(uint256 bankId, address myAddress) public view returns (bool) {
        return bankList[bankId].lenders[myAddress].status;
    }

    function getTotalNumberOfBanks() public view returns (uint256) {
        return bankCounter-1;
    }

    function getRepaymentSchedules(uint256 loanId) external view returns (RepaymentSchedule[] memory) {
        return repaymentSchedules[loanId];
    }


    function calculateDues(uint256 loanId) public view returns (uint256) {
        uint256 count = 0;
        for(uint256 i = 0; i < repaymentSchedules[loanId].length; i++) {
            if( (block.timestamp > repaymentSchedules[loanId][i].dueDate ) && repaymentSchedules[loanId][i].amountRepaid < repaymentSchedules[loanId][i].amount) {
                count++;
            }
        }

        return count;
    }

    function getAllBanksOfCustomer() public view returns (uint256[] memory) {
        return allMyBankMapping[msg.sender];
    }

    // ------------------------------------------ MOSTLY USER FUNCTIONS EXCEPT APPORTIONMENT, WHICH IS A BANK FUNCTIONALITY--------------------------------------------
    function divideAndRoundUp(uint a, uint b) internal pure returns (uint) {
        require(b > 0, "Division by zero");
        return (a + b - 1) / b;
    }

    // deposit nft
    function depositNFT(uint256 bankId, address nftContractAddress, uint256 tokenId, uint256 price) public {
        require(bankId != 0, "Bank id invalid");
        require(bankList[bankId].functional == true, "Bank id is not functional");
        require(lenderAddressToBankMapping[msg.sender] == 0, "Depositor is a lender ! Not allowed");
        uint256 loanAmount = bankList[bankId].loanPercentage*price/100;
        require(bankList[bankId].totalBalance > loanAmount, "Insufficient fund in bank");
        require(IERC721(nftContractAddress).ownerOf(tokenId) == msg.sender, "You are not the owner of the token");
        require(price > 0, "Invalid price, please pledge a nft with non-zero price");
        // PRT/100
        uint256 interest = divideAndRoundUp(loanAmount * bankList[bankId].interestRate * bankList[bankId].tenure, 100);
        uint256 repaymentAmount = loanAmount + interest;

        // transfer NFT to myself
        IERC721(nftContractAddress).transferFrom(msg.sender, address(this), tokenId);

        // append the log in memory
        NFT memory nft;
        nft.loanId = loanCounter;
        loanCounter++;
        nft.pledger = msg.sender;
        nft.tokenId = tokenId;
        nft.canWithdraw = false;
        nft.loanAmount = loanAmount;
        nft.interestRate = bankList[bankId].interestRate;
        nft.tenure = bankList[bankId].tenure;
        nft.repaymentAmount = repaymentAmount;
        nft.interest = interest;
        nft.amountRepaid = 0;
        nft.bankToWhichItisPledged = bankId;
        nft.nftContractAddress = nftContractAddress;
        nft.isActiveLoan = true;
        nft.lastKnownCurrentOwner = address(this);
        nft.amountReturnedToUser = 0;
        nft.auctionState = AuctionState.NOT_STARTED;
        nft.auctionId = 0;
        nft.amountEarnedFromAuction = 0;
        nft.lateChargesIncurred = 0;
        nft.hasClaimed = false;
        // #########################################################
        // #########################################################
        // This line should be nft.canBeAuctioned = false initially
        // Just setting this here so we can test the auction
           nft.canBeAuctioned = true;
        // ##########################################################
        // ##########################################################
        nft.nftSymbol = IERC721Metadata(nftContractAddress).symbol();
        nftDeposits[bankId][msg.sender].push(nft);

        // create repayment schedule 
        createRepaymentSchedule(nft.loanId, repaymentAmount, nft.tenure);

        // pay the amount to the user
        payable(msg.sender).transfer(loanAmount);

        // apportion profit
        apportionProfitOrLoss(bankId, loanAmount, true);

        bool customerExisting = checkIfCustomerExists(bankId, msg.sender);
        if (!customerExisting) {
            bankList[bankId].customers.push(msg.sender);
        }
        loanIdToAddressMapping[nft.loanId] = msg.sender;
        loanIdToBankIdMapping[nft.loanId] = bankId;
        allMyLoanMapping[msg.sender].push(nft.loanId);
        bool bankAlreadyAddedToList = checkBankIdExists(bankId, msg.sender);
        if(!bankAlreadyAddedToList) {
            allMyBankMapping[msg.sender].push(bankId);
        }
    }

    function checkBankIdExists(uint256 bankId, address customerAddress) private view returns (bool) {
        for(uint256 i = 0; i < allMyBankMapping[customerAddress].length; i++) {
            if(allMyBankMapping[customerAddress][i] == bankId) {
                return true;
            }
        }

        return false;
    }

    function createRepaymentSchedule(uint256 loanId, uint256 repaymentAmount, uint256 tenure) private  {
        uint256 emi = repaymentAmount / (tenure*12);
        uint256 currentTime = block.timestamp;
        for(uint256 i = 0; i < tenure*12; i++) {
            RepaymentSchedule memory rps = RepaymentSchedule({
                rpsIndex: i, 
                amount: emi,
                dueDate: currentTime + (oneMonth * (i+1)),
                amountRepaid: 0 
            });

            repaymentSchedules[loanId].push(rps);
        }

        // adjustment: Pay the extraa due in 1st emi cycle.
        uint256 adjustment = repaymentAmount - (emi*tenure*12);
        repaymentSchedules[loanId][0].amount += adjustment;
    }

    function apportionProfitOrLoss(uint256 bankId, uint256 lossOrProfit, bool isDeducted) private {
        uint256 totalSum = 0;
        for(uint256 i = 0; i < bankList[bankId].keys.length; i++) {
            address currentLendersAddress = bankList[bankId].keys[i];
            uint256 currentLendersShare = (bankList[bankId].lenders[currentLendersAddress].balance * lossOrProfit)/bankList[bankId].totalBalance;
            totalSum += currentLendersShare;
            uint256 deductedBalance = bankList[bankId].lenders[currentLendersAddress].balance - currentLendersShare;
            uint256 addedBalance =  bankList[bankId].lenders[currentLendersAddress].balance + currentLendersShare;
            bankList[bankId].lenders[currentLendersAddress].balance = isDeducted ?  deductedBalance : addedBalance;
        }

        // adjust the balance
        uint256 adjustment = lossOrProfit - totalSum;
        uint256 deductedBalance1 = bankList[bankId].lenders[bankList[bankId].admin].balance - adjustment;
        uint256 addedBalance1 = bankList[bankId].lenders[bankList[bankId].admin].balance + adjustment;
        bankList[bankId].lenders[bankList[bankId].admin].balance = isDeducted ? deductedBalance1 : addedBalance1;

        bankList[bankId].totalBalance = isDeducted ? bankList[bankId].totalBalance - lossOrProfit : bankList[bankId].totalBalance + lossOrProfit;
    }

    function checkIfCustomerExists(uint256 bankId, address customerAddress) private view returns (bool) {
        for(uint256 i = 0; i < bankList[bankId].customers.length; i++) {
            if(bankList[bankId].customers[i] == customerAddress) {
                return true;
            }
        }

        return false;
    }

    function withdrawNFT(uint256 bankId, address nftContract, uint256 loanId) public   {
        require(bankId != 0, "Bank Id is zero");
        require(loanId != 0, "Loan id is 0");
        address nftPledger = loanIdToAddressMapping[loanId];
        require(nftPledger == msg.sender,  "Invalid nft pledger");

        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.NOT_STARTED, "NFT In auction");
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan == false, "Loan is active");

        IERC721(nftContract).transferFrom(address(this), msg.sender, nftDeposits[bankId][nftPledger][index].tokenId);

        nftDeposits[bankId][nftPledger][index].lastKnownCurrentOwner = msg.sender;
    }

    function repayEmi(uint256 bankId, uint256 loanId, uint256 rpsIndex) public payable {
        address nftPledger = loanIdToAddressMapping[loanId];
        require(nftPledger != address(0), "Invalid uadr");
        require(nftPledger == msg.sender, "you cannot repay for someone else's loan");
        require(bankList[bankId].functional == true, "Bank is not functional");
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        isLoanClosed(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan == true, "Loan is closed");
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.NOT_STARTED, "In auction or finalized");
        require(repaymentSchedules[loanId][rpsIndex].amountRepaid < repaymentSchedules[loanId][rpsIndex].amount, "already paid");

        if(block.timestamp > repaymentSchedules[loanId][rpsIndex].dueDate) {
            // this is a due
            require(msg.value >= repaymentSchedules[loanId][rpsIndex].amount + latePaymentCharge, "amount not correct");
            repaymentSchedules[loanId][rpsIndex].amountRepaid += msg.value;
            nftDeposits[bankId][nftPledger][index].amountRepaid += msg.value;
            nftDeposits[bankId][nftPledger][index].lateChargesIncurred += latePaymentCharge;
        } else {
            require(msg.value >= repaymentSchedules[loanId][rpsIndex].amount);
            repaymentSchedules[loanId][rpsIndex].amountRepaid += msg.value;
            nftDeposits[bankId][nftPledger][index].amountRepaid += msg.value;
        }

        // apportion the received amount
        apportionProfitOrLoss(bankId, msg.value, false);

        // check if loan is closed
        isLoanClosed(bankId, loanId, nftPledger);

    }


    function isLoanClosed(uint256 bankId, uint256 loanId, address userAddress) public returns (bool) {
        require(bankId != 0, "Bank Id invalid");
        require(loanId != 0, "Loan Id invalid");
        require(userAddress != address(0), "User address invalid");

        uint256 index = findIndexOfLoan(bankId, loanId, userAddress);
        require(nftDeposits[bankId][userAddress][index].auctionState != AuctionState.RUNNING, "Auction is running");
        if (nftDeposits[bankId][userAddress][index].isActiveLoan == false) {
            return false;
        }


        uint256 expectedRepayment =  nftDeposits[bankId][userAddress][index].repaymentAmount + calculateDues(loanId)*latePaymentCharge;

        if( nftDeposits[bankId][userAddress][index].amountRepaid >= expectedRepayment) {
            nftDeposits[bankId][userAddress][index].isActiveLoan = false;
        }

        return nftDeposits[bankId][userAddress][index].isActiveLoan;
    }

    function claimBalanceFromLoan(uint256 bankId, uint256 loanId) public payable {
        address nftPledger = loanIdToAddressMapping[loanId];
        require(nftPledger == msg.sender);
        require(bankList[bankId].functional == true, "Bank is not functional ! They fled with your funds");
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan == false, "Loan should be inactive");
        require(nftDeposits[bankId][nftPledger][index].hasClaimed == false, "Already claimed");

        
        uint256 claimAmount = 0;
        if(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.NOT_STARTED) {
            claimAmount = nftDeposits[bankId][nftPledger][index].amountRepaid - nftDeposits[bankId][nftPledger][index].lateChargesIncurred - nftDeposits[bankId][nftPledger][index].repaymentAmount;
        } else if (nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.FINALIZED) {
            claimAmount = nftDeposits[bankId][nftPledger][index].amountRepaid;
        }

        require(claimAmount >= 1 ether, "Claim amount is too low, it should be greater than 1 ether");

        payable(msg.sender).transfer(claimAmount);

        nftDeposits[bankId][nftPledger][index].amountReturnedToUser = claimAmount;
        nftDeposits[bankId][nftPledger][index].hasClaimed = true;
        // apportion the profit
        apportionProfitOrLoss(bankId, claimAmount, true);

    }


    function findIndexOfLoan(uint256 bankId, uint256 loanId, address uadr) private view returns (uint256) {
        uint256 index = 9999999999999999999;
        for(uint256 i = 0; i < nftDeposits[bankId][uadr].length; i++) {
            if(nftDeposits[bankId][uadr][i].loanId == loanId) {
                index = i;
                break;
            }
        }

        require(index != 9999999999999999999, "Invalid loan id");
        return index;
    }

    // -------------------------------------------------------auction functions--------------------------------------------------------------------------------------------
    
    // check if NFT is auctionable
    // For now we are not going to use this because we want to show auction in demo
    // this part of require statement will be commented out
    function checkIfAuctionable(uint256 bankId, uint256 loanId) public returns (bool) {
        require(bankId != 0, "Bank Id is 0");
        require(loanId != 0, "Loan Id is 0");
        // do for an active loan
        address nftPledger = loanIdToAddressMapping[loanId];
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.NOT_STARTED);
        isLoanClosed(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan, "Loan is not active");
        // calculate number of times he has defaulted.
        // here no. of default doesn't includes the payment that is already done even though it is late.
        // It is like he hasnt paid it on the run day
        uint256 noOfDefaulted = calculateDues(loanId);
        if (noOfDefaulted >= bankList[bankId].defaultingLimit) {
            nftDeposits[bankId][nftPledger][index].canBeAuctioned = true;
            return true;
        }

        // check defaulting limit and defaulting time limit if both are crossed go for auction
        uint256 defaultPeriod = repaymentSchedules[loanId][repaymentSchedules[loanId].length-1].dueDate + bankList[bankId].defaultingTimeLimit;
        if(block.timestamp > defaultPeriod) {
            nftDeposits[bankId][nftPledger][index].canBeAuctioned = true;
            return true;
        }

        nftDeposits[bankId][nftPledger][index].canBeAuctioned = false;
        return false;
    }

    // lenders can issue an option to start for auction.
    function startAuction(uint256 bankId, uint256 loanId) public  {
        require(bankId != 0, "Bank Id invalid");
        require(loanId != 0, "Loan Id is invalid");
        // only lenders can start the auction.
        require(lenderAddressToBankMapping[msg.sender] == bankId, "Lender is not part of the bank");
        address nftPledger = loanIdToAddressMapping[loanId]; 
        // #####################################################################################
        // #####################################################################################
        // This part will be commented out as of now for the purpose of demonstration
        // This needs to be uncommented after test while deploying.
        // require(checkIfAuctionable(bankId, loanId), "Not auctionable according to the check");
        // ######################################################################################
        // ######################################################################################

        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].tokenId != 0, "Invalid Loan Id");
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.NOT_STARTED, "Auction is either running or finalized");
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan, "Loan is not active");

        // approve the auction contract
        IERC721(nftDeposits[bankId][nftPledger][index].nftContractAddress).approve(address(auctionSmartContract), nftDeposits[bankId][nftPledger][index].tokenId);

        // start the auction
        uint256 auctionBindTokenId = nftDeposits[bankId][nftPledger][index].tokenId;
        uint256 auctionBindRepaymentAmount = nftDeposits[bankId][nftPledger][index].repaymentAmount;
        address auctionBindNFTContractAddress = nftDeposits[bankId][nftPledger][index].nftContractAddress;
        address auctionBindBankAdmin = bankList[bankId].admin;
        uint256 auctionId = auctionSmartContract.startAuction(bankId, loanId, nftPledger, auctionBindTokenId, auctionBindRepaymentAmount, auctionBindNFTContractAddress, address(this), auctionBindBankAdmin);

        // black list the lenders to participate in auction
        auctionSmartContract.blackListUsersFromAuction(loanId, auctionId, bankList[bankId].keys);

        // update my NFT's status
        nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.RUNNING;
        nftDeposits[bankId][nftPledger][index].auctionId = auctionId;
    }

    function finalizeAuctionBeforeTime(uint256 bankId, uint256 loanId) public {
        require(bankId != 0, "Bank Id invalid");
        require(loanId != 0, "Loan Id is invalid");
        // only lenders can finalize the auction.
        require(lenderAddressToBankMapping[msg.sender] == bankId, "Lender is not part of the bank");
        address nftPledger = loanIdToAddressMapping[loanId]; 
        require(nftPledger != address(0), "Invalid nft pledger");
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].tokenId != 0, "Invalid Loan Id");
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.RUNNING, "Auction is either not running");
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan, "Loan is not active");
        
        // fialize the auction
        (uint256 amountEarned, address sent) = auctionSmartContract.finalizeAuctionBeforeTime(loanId);
        if(sent == address(this)) {
            // no one placed a bid and bank is the owner
            // revoke approval
            IERC721(nftDeposits[bankId][nftPledger][index].nftContractAddress).approve(address(0), nftDeposits[bankId][nftPledger][index].tokenId);
            // update our nft deposit 
            nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.NOT_STARTED;
            nftDeposits[bankId][nftPledger][index].auctionId = 0;
        } else {
            // nft has been transferred to someone
            // we received some profit we need to apportion it
            apportionProfitOrLoss(bankId, amountEarned, false);
            // update our nft details
            nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.FINALIZED;
            nftDeposits[bankId][nftPledger][index].isActiveLoan = false;
            nftDeposits[bankId][nftPledger][index].canBeAuctioned = false;
            nftDeposits[bankId][nftPledger][index].lastKnownCurrentOwner = sent;
            nftDeposits[bankId][nftPledger][index].amountEarnedFromAuction = amountEarned;
        }
    }

    function checkIfAuctionHasExpired(uint256 bankId, uint256 loanId) public {
        require(bankId != 0, "Bank Id invalid");
        require(loanId != 0, "Loan Id is invalid");
        // anyone can check it
        address nftPledger = loanIdToAddressMapping[loanId]; 
        require(nftPledger != address(0), "Invalid nft pledger");
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].tokenId != 0, "Invalid Loan Id");
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.RUNNING, "Auction is either not running");
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan, "Loan is not active");

        (uint256 amountEarned, address sent) = auctionSmartContract.checkIfAuctionHasExpired(loanId);

        if(sent == address(0)) {
            // auction hasn't expired and hence dont do anything
            return;
        } else if (sent == address(this)) {
            // no decision was taken as no one placed a bid and our contract is the highest bidder
            IERC721(nftDeposits[bankId][nftPledger][index].nftContractAddress).approve(address(0), nftDeposits[bankId][nftPledger][index].tokenId);

            // just revoke the approval and update our nft details
            nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.NOT_STARTED;
            nftDeposits[bankId][nftPledger][index].auctionId = 0;
        } else {
            // auction has ended we have earned some money and hence we need to apportion it
            apportionProfitOrLoss(bankId, amountEarned, false);

            // update our nft details and loan is closed
            nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.FINALIZED;
            nftDeposits[bankId][nftPledger][index].isActiveLoan = false;
            nftDeposits[bankId][nftPledger][index].canBeAuctioned = false;
            nftDeposits[bankId][nftPledger][index].lastKnownCurrentOwner = sent;
            nftDeposits[bankId][nftPledger][index].amountEarnedFromAuction = amountEarned;
        }
    }

    function cancelAuctionBeforeTime(uint256 bankId, uint256 loanId) public  {
        require(bankId != 0, "Bank Id invalid");
        require(loanId != 0, "Loan Id is invalid");
        // only lenders can finalize the auction.
        require(lenderAddressToBankMapping[msg.sender] == bankId, "Lender is not part of the bank");
        address nftPledger = loanIdToAddressMapping[loanId]; 
        require(nftPledger != address(0), "Invalid nft pledger");
        uint256 index = findIndexOfLoan(bankId, loanId, nftPledger);
        require(nftDeposits[bankId][nftPledger][index].tokenId != 0, "Invalid Loan Id");
        require(nftDeposits[bankId][nftPledger][index].auctionState == AuctionState.RUNNING, "Auction is either not running");
        require(nftDeposits[bankId][nftPledger][index].isActiveLoan, "Loan is not active");

        auctionSmartContract.cancelAuctionBeforeTime(loanId);
        // people placed a bid but bank is the owner as auction was cancelled
        // revoke approval
        IERC721(nftDeposits[bankId][nftPledger][index].nftContractAddress).approve(address(0), nftDeposits[bankId][nftPledger][index].tokenId);
        // update our nft deposit 
        nftDeposits[bankId][nftPledger][index].auctionState = AuctionState.NOT_STARTED;
        nftDeposits[bankId][nftPledger][index].auctionId = 0;

        // users can take their money from withdraw amount, thats upto the auction contract and we don't need to handle it here 😉
    }

    // Fallback function
    receive() external payable {}
}