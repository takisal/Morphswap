// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract GovernanceContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function voteOnProposal(
        uint ballotIndex,
        uint voteAmount
    ) public returns (bool) {
        require(_morphswapToken.balanceOf(msg.sender) >= voteAmount);
        require(
            _morphswapToken.transferFrom(msg.sender, address(this), voteAmount)
        );
        require(_ballot[ballotIndex].validUntil < block.number);
        addressToProposalsVotedOn[msg.sender].push(ballotIndex);
        delegatedTokensToAddress[msg.sender] += voteAmount;
        _ballot[ballotIndex].votes += voteAmount;
        addressToBallotToVotes[msg.sender][ballotIndex] += voteAmount;
        if (_ballot[ballotIndex].proposalType == 9) {} else {
            if (
                _ballot[ballotIndex].votes >
                _morphswapToken.totalSupply() / 3 &&
                _ballot[ballotIndex + 1].votes < _ballot[ballotIndex].votes
            ) {
                if (_ballot[ballotIndex].proposalType == 1) {
                    _fee = _ballot[ballotIndex].newValue;
                } else if (_ballot[ballotIndex].proposalType == 2) {
                    _referralBonusMultiplier = _ballot[ballotIndex].newValue;
                } else if (_ballot[ballotIndex].proposalType == 3) {
                    chainlinkFee = _ballot[ballotIndex].newValue;
                } else if (_ballot[ballotIndex].proposalType == 4) {
                    if (_ballot[ballotIndex].newValue % 2 == 1) {
                        _alternatePriceFeed = true;
                        _alternateFeeActive = true;
                    } else {
                        _alternatePriceFeed = false;
                        _alternateFeeActive = false;
                    }
                } else if (_ballot[ballotIndex].proposalType == 5) {
                    if (_ballot[ballotIndex].newValue % 2 == 1) {
                        _alternateJobID = true;
                        _alternateFeeActive = true;
                    } else {
                        _alternateJobID = false;
                        _alternateFeeActive = false;
                    }
                } else if (_ballot[ballotIndex].proposalType == 6) {
                    defaultTipMultiplier = uint128(
                        _ballot[ballotIndex].newValue
                    );
                } else if (_ballot[ballotIndex].proposalType == 7) {
                    alternateTipMultiplier = _ballot[ballotIndex].newValue;
                } else if (
                    _ballot[ballotIndex].proposalType == 8 &&
                    _ballot[ballotIndex].votes >
                    _morphswapToken.totalSupply() / 2
                ) {
                    governanceContract = address(
                        uint160(_ballot[ballotIndex].newValue)
                    );
                }
            }
        }
        return true;
    }

    function withdrawAllVotes() public returns (bool) {
        require(delegatedTokensToAddress[msg.sender] > 0);
        uint totalvotes = delegatedTokensToAddress[msg.sender];
        delegatedTokensToAddress[msg.sender] = 0;
        for (uint i; i < addressToProposalsVotedOn[msg.sender].length; i++) {
            uint current_ballot_index = addressToProposalsVotedOn[msg.sender][
                i
            ];
            uint proposal_votes = addressToBallotToVotes[msg.sender][
                current_ballot_index
            ];
            addressToBallotToVotes[msg.sender][current_ballot_index] = 0;
            _ballot[current_ballot_index].votes -= proposal_votes;
        }
        delete addressToProposalsVotedOn[msg.sender];
        require(_morphswapToken.transfer(msg.sender, totalvotes));
        return true;
    }

    function withdrawVotesSpecificProposal(
        uint ballotIndex
    ) public returns (bool) {
        require(delegatedTokensToAddress[msg.sender] > 0);
        uint caller_votes = addressToBallotToVotes[msg.sender][ballotIndex];
        addressToBallotToVotes[msg.sender][ballotIndex] = 0;
        delegatedTokensToAddress[msg.sender] -= caller_votes;
        _ballot[ballotIndex].votes -= caller_votes;
        require(_morphswapToken.transfer(msg.sender, caller_votes));
        return true;
    }

    function addProposal(
        uint proposalType,
        uint newRate,
        uint startingWeight
    ) public returns (uint) {
        require(
            _morphswapToken.balanceOf(msg.sender) >
                _morphswapToken.totalSupply() / 50
        );
        require(startingWeight > _morphswapToken.totalSupply() / 50);
        require(
            _morphswapToken.transferFrom(
                msg.sender,
                address(this),
                startingWeight
            )
        );
        require(proposalType < 9);
        addressToProposalsVotedOn[msg.sender].push(_ballot.length);
        delegatedTokensToAddress[msg.sender] = startingWeight;
        addressToBallotToVotes[msg.sender][_ballot.length] = startingWeight;
        _ballot.push(
            Proposal(
                proposalType,
                newRate,
                startingWeight,
                block.number + _proposalLifespan
            )
        );
        _ballot.push(Proposal(9, 0, 0, block.number + _proposalLifespan));
        return _ballot.length - 2;
    }
}
