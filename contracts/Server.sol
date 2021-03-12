pragma solidity 0.6.2;

import { USDTContract } from "./ContractUSDT.sol";
import { USDTGuarantee } from "./GuaranteeUSDT.sol";

contract Server is USDTGuarantee {
    address[] public contracts;
    uint public contractsNum;

    mapping(address => uint[]) public relations;
    mapping(address => uint) public relationsNum;

    constructor(address payable _server) USDTGuarantee(_server) public {
    }

    modifier requireDifferentAddress(
        address payable _A,
        address payable _B,
        address payable _judge
    ) {
        require(_A != _B, "Address A and address B should not be equal.");
        require(_A != _judge, "Address A and address judge should not be equal.");
        require(_B != _judge, "Address B and address judge should not be equal.");
        _;
    }

    function createContract(
        address payable _A,
        address payable _B,
        address payable _judge,
        string memory _Content,
        uint _feePercentLimit) public requireDifferentAddress(_A, _B, _judge) returns (uint)
    {
        USDTContract newContract = new USDTContract(_A, _B, _judge, payable(address(this)), _Content, _feePercentLimit);

        address contractAdr = address(newContract);

        contracts.push(contractAdr);

        relations[_A].push(contractsNum);
        relations[_B].push(contractsNum);
        relations[_judge].push(contractsNum);

        relationsNum[_A]++;
        relationsNum[_B]++;
        relationsNum[_judge]++;

        return contractsNum++;
    }

    function interveneContract(
        address contractAdr,
        uint256 amountToA,
        uint256 amountToB,
        uint feePercent
    ) public onlyServer {
        USDTContract c = USDTContract(contractAdr);
        c.intervene(amountToA, amountToB, feePercent);
    }


}