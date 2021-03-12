pragma solidity 0.6.2;

import { BaseGuarantee } from "./Guarantee.sol";

contract USDTGuarantee is BaseGuarantee{
    address public token;
    constructor(address payable _server) public BaseGuarantee(_server) {
        // Mainnet Tether contract address: 0xdAC17F958D2ee523a2206206994597C13D831ec7
        token = 0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660;
    }

    function _processDeposit(address _guarantor, uint256 amount) internal override{
        _safeErc20TransferFrom(_guarantor, address(this), amount);
    }
    function _processWithdraw(address to, uint256 amount) internal override{
        _safeErc20Transfer(to, amount);
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
