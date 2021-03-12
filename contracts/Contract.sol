pragma solidity 0.6.2;

abstract contract BaseContract {
    string public content;
    address payable public A;
    address payable public B;
    address payable public judge;
    //Server contract address
    address payable public server;
    uint public feePercentLimit;

    uint256 public balance;

    event Deposit(uint256 time, address from, uint256 amount);
    event Withdraw(uint256 time, uint256 amountToA, uint256 amountToB, uint feePercent);
    event Intervene(uint256 time, uint256 amountToA, uint256 amountToB, uint feePercent);

    function _processDeposit(uint256 amount) internal virtual;
    function _processWithdraw(uint256 amountToA, uint256 amountToB, uint feePercent) internal virtual;
    function _processIntervene(uint256 amountToA, uint256 amountToB, uint feePercent) internal virtual;

    constructor(address payable _A, address payable _B, address payable _judge, address payable _server, string memory _content, uint _feePercentLimit) public {
        require(_feePercentLimit <= 100, "Fee percentage limit should be less than 100%");
        content = _content;
        A = _A;
        B = _B;
        judge = _judge;
        server = _server;
        feePercentLimit = _feePercentLimit;
    }

    modifier onlyServer(){
        require(msg.sender == server, "msg.sender must be server.");
        _;
    }

    modifier onlyJudge(){
        require(msg.sender == judge, "msg.sender must be judge.");
        _;
    }

    modifier limitFeePercent(uint feePercent){
        require(feePercent <= feePercentLimit, "Fee percent should be less than fee percent limit.");
        _;
    }

    modifier checkBalance(uint256 amount){
        require(amount <= balance, "Withdraw amount should be less than balance.");
        _;
    }

    function deposit(uint256 amount) external {
        _processDeposit(amount);

        balance += amount;
        emit Deposit(now, msg.sender, amount);
    }

    function withdraw(
        uint256 amountToA,
        uint256 amountToB,
        uint feePercent
    ) external onlyJudge limitFeePercent(feePercent) checkBalance(amountToA + amountToB) {
        _processWithdraw(amountToA, amountToB, feePercent);

        balance -= amountToA + amountToB;
        emit Withdraw(now, amountToA, amountToB, feePercent);
    }

    function intervene(
        uint256 amountToA,
        uint256 amountToB,
        uint feePercent
    ) external onlyServer checkBalance(amountToA + amountToB) {
        _processIntervene(amountToA,amountToA,feePercent);

        balance -= amountToA + amountToB;
        emit Intervene(now, amountToA, amountToB, feePercent);
    }

}
