// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.9;


contract OverallContractMask {

     struct txobj {
        uint8 method_id;
        uint8 internal_start_chainid;
        uint8 internal_end_chainid;
        uint64 pair_id;
        address finalchain_wallet;
        uint64 secondpair_id;
        address firstchain_asset;
        address finalchain_asset;
        uint64 quadrillionratio;
        uint64 quadrilliontipratio;
        uint128 rtxnum;
        bool alt_fee;
    }
    
  function getTxByRTxNumber(uint8 reqchainid, uint128 rtx_number) public view returns (txobj memory) {
      return txobj(reqchainid, 1, 2, 3, address(0), 5, address(0),address(0), 0, 0, rtx_number, true);
  }
  function oraclePing(uint8 _icid) public returns (uint128, bool){
      return (uint128(568797), true);
  }
   fallback() external payable {  }

    receive() external payable { }
     
}