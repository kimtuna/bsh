// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CompanyServerAccess {
    struct Company {
        string name;
        string ceoName;
        uint256 subscriptionEnd;  // 구독 만료일
        bool isActive;            // 서비스 활성화 상태
    }
    
    // 회사 주소 => 회사 정보
    mapping(address => Company) public companies;
    
    // 구독 가격 (wei 단위)
    uint256 public monthlyPrice = 0.01 ether;    // 1개월
    uint256 public quarterlyPrice = 0.025 ether; // 3개월
    uint256 public yearlyPrice = 0.08 ether;     // 1년
    
    // 이벤트
    event CompanyRegistered(address indexed company, string name, uint256 subscriptionEnd);
    event SubscriptionExtended(address indexed company, uint256 newEndDate);
    event ServiceExpired(address indexed company);
    
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    // 회사 등록 및 구독
    function registerCompany(
        string memory _name,
        string memory _ceoName,
        uint256 _subscriptionType // 1: 1개월, 2: 3개월, 3: 1년
    ) external payable {
        require(msg.value >= getSubscriptionPrice(_subscriptionType), "Insufficient payment");
        uint256 subscriptionDuration;
        if (_subscriptionType == 1) {
            subscriptionDuration = 30 days;
        } else if (_subscriptionType == 2) {
            subscriptionDuration = 90 days;
        } else if (_subscriptionType == 3) {
            subscriptionDuration = 365 days;
        } else {
            revert("Invalid subscription type");
        }
        companies[msg.sender] = Company({
            name: _name,
            ceoName: _ceoName,
            subscriptionEnd: block.timestamp + subscriptionDuration,
            isActive: true
        });
        emit CompanyRegistered(msg.sender, _name, block.timestamp + subscriptionDuration);
    }
    
    // 구독 연장
    function extendSubscription(uint256 _subscriptionType) external payable {
        require(companies[msg.sender].isActive, "Company not registered");
        require(msg.value >= getSubscriptionPrice(_subscriptionType), "Insufficient payment");
        
        uint256 subscriptionDuration;
        if (_subscriptionType == 1) {
            subscriptionDuration = 30 days;
        } else if (_subscriptionType == 2) {
            subscriptionDuration = 90 days;
        } else if (_subscriptionType == 3) {
            subscriptionDuration = 365 days;
        } else {
            revert("Invalid subscription type");
        }
        
        companies[msg.sender].subscriptionEnd = block.timestamp + subscriptionDuration; // 구독 끝날 경우 서비스 중지
        emit SubscriptionExtended(msg.sender, block.timestamp + subscriptionDuration);
    }
    
    // 구독 상태 확인
    function checkSubscription(address _company) public view returns (bool) {
        return companies[_company].isActive && 
               companies[_company].subscriptionEnd > block.timestamp;
    }
    
    // 구독 가격 조회
    function getSubscriptionPrice(uint256 _type) public view returns (uint256) {
        if (_type == 1) return monthlyPrice;
        if (_type == 2) return quarterlyPrice;
        if (_type == 3) return yearlyPrice;
        revert("Invalid subscription type");
    }
    
    // 회사 정보 조회
    function getCompanyInfo(address _company) external view returns (
        string memory name,
        string memory ceoName,
        uint256 subscriptionEnd,
        bool isActive
    ) {
        Company storage company = companies[_company];
        return (
            company.name,
            company.ceoName,
            company.subscriptionEnd,
            company.isActive
        );
    }
    
    // 관리자 함수: 가격 설정
    function setPrices(
        uint256 _monthly,
        uint256 _quarterly,
        uint256 _yearly
    ) external onlyOwner {
        monthlyPrice = _monthly;
        quarterlyPrice = _quarterly;
        yearlyPrice = _yearly;
    }
    
    // 관리자 함수: 구독 만료 처리
    function expireSubscription(address _company) external onlyOwner {
        companies[_company].isActive = false;
        emit ServiceExpired(_company);
    }
}
