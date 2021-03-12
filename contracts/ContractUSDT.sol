// SPDX-License-Identifier: MIT
pragma solidity 0.6.2;

import { BaseContract } from "./Contract.sol";

contract USDTContract is BaseContract {
    address public token;

    constructor(
        address payable _A,
        address payable _B,
        address payable _judge,
        address payable _Server,
        string memory _Content,
        uint _feePercentLimit
    ) public BaseContract(_A, _B, _judge, _Server, _Content, _feePercentLimit) {
        // Mainnet Tether contract address: 0xdAC17F958D2ee523a2206206994597C13D831ec7
        token = 0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660;
    }

    function _processDeposit(uint256 amount) internal override {
        require(
            msg.value == 0,
            "ETH value is supposed to be 0 for ERC20 instance"
        );
        _safeErc20TransferFrom(msg.sender, address(this), amount);
    }

    function _processWithdraw(uint256 amountToA, uint256 amountToB, uint feePercent) internal override {
        uint256 feeA;
        uint256 feeB;

        if(amountToA > 0){
            feeA = amountToA * feePercent / 100;
            _safeErc20Transfer(A, amountToA - feeA);
        }
        if(amountToB > 0){
            feeB = amountToB * feePercent / 100;
            _safeErc20Transfer(B, amountToB - feeB);
        }

        _safeErc20Transfer(judge, feeA + feeB);
    }

    function _processIntervene(uint256 amountToA, uint256 amountToB, uint feePercent) internal override {
        uint256 feeA;
        uint256 feeB;

        if(amountToA > 0){
            feeA = amountToA * feePercent / 100;
            _safeErc20Transfer(A, amountToA - feeA);
        }
        if(amountToB > 0){
            feeB = amountToB * feePercent / 100;
            _safeErc20Transfer(B, amountToB - feeB);
        }

        _safeErc20Transfer(server, feeA + feeB);
    }

    function _safeErc20TransferFrom(
        address _from,
        address _to,
        uint256 _amount
    ) internal {
        (bool success, bytes memory data) = token.call(
            abi.encodeWithSelector(
                0x23b872dd, /* transferFrom */
                _from,
                _to,
                _amount
            )
        );
        require(success, "not enough allowed tokens");

        // if contract returns some data lets make sure that is `true` according to standard
        if (data.length > 0) {
            require(
                data.length == 32,
                "data length should be either 0 or 32 bytes"
            );
            success = abi.decode(data, (bool));
            require(success, "not enough allowed tokens. Token returns false.");
        }
    }

    function _safeErc20Transfer(address _to, uint256 _amount) internal {
        (bool success, bytes memory data) = token.call(
            abi.encodeWithSelector(
                0xa9059cbb, /* transfer */
                _to,
                _amount
            )
        );
        require(success, "not enough tokens");

        // if contract returns some data lets make sure that is `true` according to standard
        if (data.length > 0) {
            require(
                data.length == 32,
                "data length should be either 0 or 32 bytes"
            );
            success = abi.decode(data, (bool));
            require(success, "not enough tokens. Token returns false.");
        }
    }
}
