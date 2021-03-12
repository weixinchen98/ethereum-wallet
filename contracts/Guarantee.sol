pragma solidity 0.6.2;

abstract contract BaseGuarantee {
    address payable public server;
    address payable[] public guarantor;
    mapping(address => uint256) public bail;
    mapping(address => string) public name;

    event Register(address guarantor, uint256 amount, string name);
    event Withdraw(address guarantor, address payable _to, uint256 amount);
    event Intervene(address guarantor, address payable _to, uint256 amount);

    function _processDeposit(address _guarantor, uint256 amount) internal virtual;
    function _processWithdraw(address to, uint256 amount) internal virtual;

    constructor(address payable _server) public {
        server = _server;
    }

    modifier onlyServer(){
        require(msg.sender == server, "msg.sender should be server.");
        _;
    }

    modifier onlyNotRegistered(){
        require(bail[msg.sender] == 0, "This address is already registered.");
        _;
    }

    function register(uint256 amount, string calldata _name) external onlyNotRegistered {
        //TODO: setting lower bound of registering amount.
        require(amount > 0, "Registering amount should not be zero.");

        _processDeposit(msg.sender, amount);

        guarantor.push(msg.sender);
        bail[msg.sender] = amount;
        name[msg.sender] = _name;
        emit Register(msg.sender, amount, _name);
    }

    function intervene(address payable _guarantor, address payable _to, uint256 amount) external onlyServer {
        require(bail[_guarantor] >= amount, "Withdraw amount should be less than guarantor's bail.");

        _processWithdraw(_to, amount);

        bail[_guarantor] -= amount;
        emit Intervene(_guarantor, _to, amount);
    }
}
