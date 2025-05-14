// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CompanyServerAccessMetaData contains all meta data concerning the CompanyServerAccess contract.
var CompanyServerAccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"company\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subscriptionEnd\",\"type\":\"uint256\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"company\",\"type\":\"address\"}],\"name\":\"ServiceExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"company\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEndDate\",\"type\":\"uint256\"}],\"name\":\"SubscriptionExtended\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_company\",\"type\":\"address\"}],\"name\":\"checkSubscription\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"companies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ceoName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"subscriptionEnd\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_company\",\"type\":\"address\"}],\"name\":\"expireSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subscriptionType\",\"type\":\"uint256\"}],\"name\":\"extendSubscription\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_company\",\"type\":\"address\"}],\"name\":\"getCompanyInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ceoName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"subscriptionEnd\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"getSubscriptionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"monthlyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quarterlyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ceoName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_subscriptionType\",\"type\":\"uint256\"}],\"name\":\"registerCompany\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_monthly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quarterly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_yearly\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yearlyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052662386f26fc100006001556658d15e1762800060025567011c37937e0800006003553480156030575f5ffd5b503360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506116288061007e5f395ff3fe6080604052600436106100a6575f3560e01c8063971dc02811610063578063971dc028146101e0578063a06c5a24146101fc578063a88fe42d14610226578063ad644ccb1461024e578063fc59282d1461028a578063ff82286b146102a6576100a6565b806318a36b09146100aa578063355e6ce8146100d257806342e12c83146101115780634f62647e146101505780636d2db8841461018c5780638da5cb5b146101b6575b5f5ffd5b3480156100b5575f5ffd5b506100d060048036038101906100cb9190610d01565b6102d0565b005b3480156100dd575f5ffd5b506100f860048036038101906100f39190610d01565b6103fb565b6040516101089493929190610dce565b60405180910390f35b34801561011c575f5ffd5b5061013760048036038101906101329190610d01565b61053e565b6040516101479493929190610dce565b60405180910390f35b34801561015b575f5ffd5b5061017660048036038101906101719190610d01565b6106c3565b6040516101839190610e1f565b60405180910390f35b348015610197575f5ffd5b506101a0610761565b6040516101ad9190610e38565b60405180910390f35b3480156101c1575f5ffd5b506101ca610767565b6040516101d79190610e60565b60405180910390f35b6101fa60048036038101906101f59190610fcf565b61078c565b005b348015610207575f5ffd5b50610210610970565b60405161021d9190610e38565b60405180910390f35b348015610231575f5ffd5b5061024c60048036038101906102479190611057565b610976565b005b348015610259575f5ffd5b50610274600480360381019061026f91906110a7565b610a1f565b6040516102819190610e38565b60405180910390f35b6102a4600480360381019061029f91906110a7565b610a96565b005b3480156102b1575f5ffd5b506102ba610c90565b6040516102c79190610e38565b60405180910390f35b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461035f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035690611142565b60405180910390fd5b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f6101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167ff62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe360405160405180910390a250565b5f602052805f5260405f205f91509050805f0180546104199061118d565b80601f01602080910402602001604051908101604052809291908181526020018280546104459061118d565b80156104905780601f1061046757610100808354040283529160200191610490565b820191905f5260205f20905b81548152906001019060200180831161047357829003601f168201915b5050505050908060010180546104a59061118d565b80601f01602080910402602001604051908101604052809291908181526020018280546104d19061118d565b801561051c5780601f106104f35761010080835404028352916020019161051c565b820191905f5260205f20905b8154815290600101906020018083116104ff57829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b6060805f5f5f5f5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050805f01816001018260020154836003015f9054906101000a900460ff168380546105ab9061118d565b80601f01602080910402602001604051908101604052809291908181526020018280546105d79061118d565b80156106225780601f106105f957610100808354040283529160200191610622565b820191905f5260205f20905b81548152906001019060200180831161060557829003601f168201915b505050505093508280546106359061118d565b80601f01602080910402602001604051908101604052809291908181526020018280546106619061118d565b80156106ac5780601f10610683576101008083540402835291602001916106ac565b820191905f5260205f20905b81548152906001019060200180831161068f57829003601f168201915b505050505092509450945094509450509193509193565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f9054906101000a900460ff16801561075a5750425f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060020154115b9050919050565b60025481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61079581610a1f565b3410156107d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ce90611207565b60405180910390fd5b5f600182036107eb5762278d00905061084f565b600282036107fe576276a700905061084e565b60038203610812576301e13380905061084d565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108449061126f565b60405180910390fd5b5b5b6040518060800160405280858152602001848152602001824261087291906112ba565b8152602001600115158152505f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0190816108cd919061148d565b5060208201518160010190816108e3919061148d565b50604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167e78dab93f1aeefc09535aa58402d91dde97c29a340cd7a0a57e12529c94e21485834261095491906112ba565b60405161096292919061155c565b60405180910390a250505050565b60015481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a05576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fc90611142565b60405180910390fd5b826001819055508160028190555080600381905550505050565b5f60018203610a32576001549050610a91565b60028203610a44576002549050610a91565b60038203610a56576003549050610a91565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a889061126f565b60405180910390fd5b919050565b5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f9054906101000a900460ff16610b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b18906115d4565b60405180910390fd5b610b2a81610a1f565b341015610b6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6390611207565b60405180910390fd5b5f60018203610b805762278d009050610be4565b60028203610b93576276a7009050610be3565b60038203610ba7576301e133809050610be2565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd99061126f565b60405180910390fd5b5b5b8042610bf091906112ba565b5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600201819055503373ffffffffffffffffffffffffffffffffffffffff167f6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb25548242610c7791906112ba565b604051610c849190610e38565b60405180910390a25050565b60035481565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cd082610ca7565b9050919050565b610ce081610cc6565b8114610cea575f5ffd5b50565b5f81359050610cfb81610cd7565b92915050565b5f60208284031215610d1657610d15610c9f565b5b5f610d2384828501610ced565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610d6e82610d2c565b610d788185610d36565b9350610d88818560208601610d46565b610d9181610d54565b840191505092915050565b5f819050919050565b610dae81610d9c565b82525050565b5f8115159050919050565b610dc881610db4565b82525050565b5f6080820190508181035f830152610de68187610d64565b90508181036020830152610dfa8186610d64565b9050610e096040830185610da5565b610e166060830184610dbf565b95945050505050565b5f602082019050610e325f830184610dbf565b92915050565b5f602082019050610e4b5f830184610da5565b92915050565b610e5a81610cc6565b82525050565b5f602082019050610e735f830184610e51565b92915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610eb782610d54565b810181811067ffffffffffffffff82111715610ed657610ed5610e81565b5b80604052505050565b5f610ee8610c96565b9050610ef48282610eae565b919050565b5f67ffffffffffffffff821115610f1357610f12610e81565b5b610f1c82610d54565b9050602081019050919050565b828183375f83830152505050565b5f610f49610f4484610ef9565b610edf565b905082815260208101848484011115610f6557610f64610e7d565b5b610f70848285610f29565b509392505050565b5f82601f830112610f8c57610f8b610e79565b5b8135610f9c848260208601610f37565b91505092915050565b610fae81610d9c565b8114610fb8575f5ffd5b50565b5f81359050610fc981610fa5565b92915050565b5f5f5f60608486031215610fe657610fe5610c9f565b5b5f84013567ffffffffffffffff81111561100357611002610ca3565b5b61100f86828701610f78565b935050602084013567ffffffffffffffff8111156110305761102f610ca3565b5b61103c86828701610f78565b925050604061104d86828701610fbb565b9150509250925092565b5f5f5f6060848603121561106e5761106d610c9f565b5b5f61107b86828701610fbb565b935050602061108c86828701610fbb565b925050604061109d86828701610fbb565b9150509250925092565b5f602082840312156110bc576110bb610c9f565b5b5f6110c984828501610fbb565b91505092915050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f61112c602183610d36565b9150611137826110d2565b604082019050919050565b5f6020820190508181035f83015261115981611120565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111a457607f821691505b6020821081036111b7576111b6611160565b5b50919050565b7f496e73756666696369656e74207061796d656e740000000000000000000000005f82015250565b5f6111f1601483610d36565b91506111fc826111bd565b602082019050919050565b5f6020820190508181035f83015261121e816111e5565b9050919050565b7f496e76616c696420737562736372697074696f6e2074797065000000000000005f82015250565b5f611259601983610d36565b915061126482611225565b602082019050919050565b5f6020820190508181035f8301526112868161124d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6112c482610d9c565b91506112cf83610d9c565b92508282019050808211156112e7576112e661128d565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261130e565b611353868361130e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61138e61138961138484610d9c565b61136b565b610d9c565b9050919050565b5f819050919050565b6113a783611374565b6113bb6113b382611395565b84845461131a565b825550505050565b5f5f905090565b6113d26113c3565b6113dd81848461139e565b505050565b5b81811015611400576113f55f826113ca565b6001810190506113e3565b5050565b601f82111561144557611416816112ed565b61141f846112ff565b8101602085101561142e578190505b61144261143a856112ff565b8301826113e2565b50505b505050565b5f82821c905092915050565b5f6114655f198460080261144a565b1980831691505092915050565b5f61147d8383611456565b9150826002028217905092915050565b61149682610d2c565b67ffffffffffffffff8111156114af576114ae610e81565b5b6114b9825461118d565b6114c4828285611404565b5f60209050601f8311600181146114f5575f84156114e3578287015190505b6114ed8582611472565b865550611554565b601f198416611503866112ed565b5f5b8281101561152a57848901518255600182019150602085019450602081019050611505565b868310156115475784890151611543601f891682611456565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f8301526115748185610d64565b90506115836020830184610da5565b9392505050565b7f436f6d70616e79206e6f742072656769737465726564000000000000000000005f82015250565b5f6115be601683610d36565b91506115c98261158a565b602082019050919050565b5f6020820190508181035f8301526115eb816115b2565b905091905056fea2646970667358221220d41c59921246a5351ae9d17d19a011af61b54eac5b002f86eec74227e8f0c14664736f6c634300081d0033",
}

// CompanyServerAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use CompanyServerAccessMetaData.ABI instead.
var CompanyServerAccessABI = CompanyServerAccessMetaData.ABI

// CompanyServerAccessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompanyServerAccessMetaData.Bin instead.
var CompanyServerAccessBin = CompanyServerAccessMetaData.Bin

// DeployCompanyServerAccess deploys a new Ethereum contract, binding an instance of CompanyServerAccess to it.
func DeployCompanyServerAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompanyServerAccess, error) {
	parsed, err := CompanyServerAccessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompanyServerAccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompanyServerAccess{CompanyServerAccessCaller: CompanyServerAccessCaller{contract: contract}, CompanyServerAccessTransactor: CompanyServerAccessTransactor{contract: contract}, CompanyServerAccessFilterer: CompanyServerAccessFilterer{contract: contract}}, nil
}

// CompanyServerAccess is an auto generated Go binding around an Ethereum contract.
type CompanyServerAccess struct {
	CompanyServerAccessCaller     // Read-only binding to the contract
	CompanyServerAccessTransactor // Write-only binding to the contract
	CompanyServerAccessFilterer   // Log filterer for contract events
}

// CompanyServerAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompanyServerAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompanyServerAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompanyServerAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompanyServerAccessSession struct {
	Contract     *CompanyServerAccess // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CompanyServerAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompanyServerAccessCallerSession struct {
	Contract *CompanyServerAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CompanyServerAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompanyServerAccessTransactorSession struct {
	Contract     *CompanyServerAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CompanyServerAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompanyServerAccessRaw struct {
	Contract *CompanyServerAccess // Generic contract binding to access the raw methods on
}

// CompanyServerAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompanyServerAccessCallerRaw struct {
	Contract *CompanyServerAccessCaller // Generic read-only contract binding to access the raw methods on
}

// CompanyServerAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompanyServerAccessTransactorRaw struct {
	Contract *CompanyServerAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompanyServerAccess creates a new instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccess(address common.Address, backend bind.ContractBackend) (*CompanyServerAccess, error) {
	contract, err := bindCompanyServerAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccess{CompanyServerAccessCaller: CompanyServerAccessCaller{contract: contract}, CompanyServerAccessTransactor: CompanyServerAccessTransactor{contract: contract}, CompanyServerAccessFilterer: CompanyServerAccessFilterer{contract: contract}}, nil
}

// NewCompanyServerAccessCaller creates a new read-only instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessCaller(address common.Address, caller bind.ContractCaller) (*CompanyServerAccessCaller, error) {
	contract, err := bindCompanyServerAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessCaller{contract: contract}, nil
}

// NewCompanyServerAccessTransactor creates a new write-only instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*CompanyServerAccessTransactor, error) {
	contract, err := bindCompanyServerAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessTransactor{contract: contract}, nil
}

// NewCompanyServerAccessFilterer creates a new log filterer instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*CompanyServerAccessFilterer, error) {
	contract, err := bindCompanyServerAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessFilterer{contract: contract}, nil
}

// bindCompanyServerAccess binds a generic wrapper to an already deployed contract.
func bindCompanyServerAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompanyServerAccessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompanyServerAccess *CompanyServerAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompanyServerAccess.Contract.CompanyServerAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompanyServerAccess *CompanyServerAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.CompanyServerAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompanyServerAccess *CompanyServerAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.CompanyServerAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompanyServerAccess *CompanyServerAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompanyServerAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompanyServerAccess *CompanyServerAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompanyServerAccess *CompanyServerAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.contract.Transact(opts, method, params...)
}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessCaller) CheckSubscription(opts *bind.CallOpts, _company common.Address) (bool, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "checkSubscription", _company)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessSession) CheckSubscription(_company common.Address) (bool, error) {
	return _CompanyServerAccess.Contract.CheckSubscription(&_CompanyServerAccess.CallOpts, _company)
}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) CheckSubscription(_company common.Address) (bool, error) {
	return _CompanyServerAccess.Contract.CheckSubscription(&_CompanyServerAccess.CallOpts, _company)
}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCaller) Companies(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "companies", arg0)

	outstruct := new(struct {
		Name            string
		CeoName         string
		SubscriptionEnd *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CeoName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SubscriptionEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessSession) Companies(arg0 common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.Companies(&_CompanyServerAccess.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) Companies(arg0 common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.Companies(&_CompanyServerAccess.CallOpts, arg0)
}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCaller) GetCompanyInfo(opts *bind.CallOpts, _company common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "getCompanyInfo", _company)

	outstruct := new(struct {
		Name            string
		CeoName         string
		SubscriptionEnd *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CeoName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SubscriptionEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessSession) GetCompanyInfo(_company common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.GetCompanyInfo(&_CompanyServerAccess.CallOpts, _company)
}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) GetCompanyInfo(_company common.Address) (struct {
	Name            string
	CeoName         string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.GetCompanyInfo(&_CompanyServerAccess.CallOpts, _company)
}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) GetSubscriptionPrice(opts *bind.CallOpts, _type *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "getSubscriptionPrice", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) GetSubscriptionPrice(_type *big.Int) (*big.Int, error) {
	return _CompanyServerAccess.Contract.GetSubscriptionPrice(&_CompanyServerAccess.CallOpts, _type)
}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) GetSubscriptionPrice(_type *big.Int) (*big.Int, error) {
	return _CompanyServerAccess.Contract.GetSubscriptionPrice(&_CompanyServerAccess.CallOpts, _type)
}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) MonthlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "monthlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) MonthlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.MonthlyPrice(&_CompanyServerAccess.CallOpts)
}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) MonthlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.MonthlyPrice(&_CompanyServerAccess.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessSession) Owner() (common.Address, error) {
	return _CompanyServerAccess.Contract.Owner(&_CompanyServerAccess.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) Owner() (common.Address, error) {
	return _CompanyServerAccess.Contract.Owner(&_CompanyServerAccess.CallOpts)
}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) QuarterlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "quarterlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) QuarterlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.QuarterlyPrice(&_CompanyServerAccess.CallOpts)
}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) QuarterlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.QuarterlyPrice(&_CompanyServerAccess.CallOpts)
}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) YearlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "yearlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) YearlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.YearlyPrice(&_CompanyServerAccess.CallOpts)
}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) YearlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.YearlyPrice(&_CompanyServerAccess.CallOpts)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) ExpireSubscription(opts *bind.TransactOpts, _company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "expireSubscription", _company)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessSession) ExpireSubscription(_company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExpireSubscription(&_CompanyServerAccess.TransactOpts, _company)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) ExpireSubscription(_company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExpireSubscription(&_CompanyServerAccess.TransactOpts, _company)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) ExtendSubscription(opts *bind.TransactOpts, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "extendSubscription", _subscriptionType)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessSession) ExtendSubscription(_subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExtendSubscription(&_CompanyServerAccess.TransactOpts, _subscriptionType)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) ExtendSubscription(_subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExtendSubscription(&_CompanyServerAccess.TransactOpts, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x971dc028.
//
// Solidity: function registerCompany(string _name, string _ceoName, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) RegisterCompany(opts *bind.TransactOpts, _name string, _ceoName string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "registerCompany", _name, _ceoName, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x971dc028.
//
// Solidity: function registerCompany(string _name, string _ceoName, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessSession) RegisterCompany(_name string, _ceoName string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.RegisterCompany(&_CompanyServerAccess.TransactOpts, _name, _ceoName, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x971dc028.
//
// Solidity: function registerCompany(string _name, string _ceoName, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) RegisterCompany(_name string, _ceoName string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.RegisterCompany(&_CompanyServerAccess.TransactOpts, _name, _ceoName, _subscriptionType)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) SetPrices(opts *bind.TransactOpts, _monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "setPrices", _monthly, _quarterly, _yearly)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessSession) SetPrices(_monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.SetPrices(&_CompanyServerAccess.TransactOpts, _monthly, _quarterly, _yearly)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) SetPrices(_monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.SetPrices(&_CompanyServerAccess.TransactOpts, _monthly, _quarterly, _yearly)
}

// CompanyServerAccessCompanyRegisteredIterator is returned from FilterCompanyRegistered and is used to iterate over the raw logs and unpacked data for CompanyRegistered events raised by the CompanyServerAccess contract.
type CompanyServerAccessCompanyRegisteredIterator struct {
	Event *CompanyServerAccessCompanyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompanyServerAccessCompanyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessCompanyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompanyServerAccessCompanyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompanyServerAccessCompanyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessCompanyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessCompanyRegistered represents a CompanyRegistered event raised by the CompanyServerAccess contract.
type CompanyServerAccessCompanyRegistered struct {
	Company         common.Address
	Name            string
	SubscriptionEnd *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompanyRegistered is a free log retrieval operation binding the contract event 0x0078dab93f1aeefc09535aa58402d91dde97c29a340cd7a0a57e12529c94e214.
//
// Solidity: event CompanyRegistered(address indexed company, string name, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterCompanyRegistered(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessCompanyRegisteredIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "CompanyRegistered", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessCompanyRegisteredIterator{contract: _CompanyServerAccess.contract, event: "CompanyRegistered", logs: logs, sub: sub}, nil
}

// WatchCompanyRegistered is a free log subscription operation binding the contract event 0x0078dab93f1aeefc09535aa58402d91dde97c29a340cd7a0a57e12529c94e214.
//
// Solidity: event CompanyRegistered(address indexed company, string name, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchCompanyRegistered(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessCompanyRegistered, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "CompanyRegistered", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessCompanyRegistered)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompanyRegistered is a log parse operation binding the contract event 0x0078dab93f1aeefc09535aa58402d91dde97c29a340cd7a0a57e12529c94e214.
//
// Solidity: event CompanyRegistered(address indexed company, string name, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseCompanyRegistered(log types.Log) (*CompanyServerAccessCompanyRegistered, error) {
	event := new(CompanyServerAccessCompanyRegistered)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompanyServerAccessServiceExpiredIterator is returned from FilterServiceExpired and is used to iterate over the raw logs and unpacked data for ServiceExpired events raised by the CompanyServerAccess contract.
type CompanyServerAccessServiceExpiredIterator struct {
	Event *CompanyServerAccessServiceExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompanyServerAccessServiceExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessServiceExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompanyServerAccessServiceExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompanyServerAccessServiceExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessServiceExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessServiceExpired represents a ServiceExpired event raised by the CompanyServerAccess contract.
type CompanyServerAccessServiceExpired struct {
	Company common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterServiceExpired is a free log retrieval operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterServiceExpired(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessServiceExpiredIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "ServiceExpired", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessServiceExpiredIterator{contract: _CompanyServerAccess.contract, event: "ServiceExpired", logs: logs, sub: sub}, nil
}

// WatchServiceExpired is a free log subscription operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchServiceExpired(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessServiceExpired, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "ServiceExpired", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessServiceExpired)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "ServiceExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceExpired is a log parse operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseServiceExpired(log types.Log) (*CompanyServerAccessServiceExpired, error) {
	event := new(CompanyServerAccessServiceExpired)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "ServiceExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompanyServerAccessSubscriptionExtendedIterator is returned from FilterSubscriptionExtended and is used to iterate over the raw logs and unpacked data for SubscriptionExtended events raised by the CompanyServerAccess contract.
type CompanyServerAccessSubscriptionExtendedIterator struct {
	Event *CompanyServerAccessSubscriptionExtended // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompanyServerAccessSubscriptionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessSubscriptionExtended)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompanyServerAccessSubscriptionExtended)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompanyServerAccessSubscriptionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessSubscriptionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessSubscriptionExtended represents a SubscriptionExtended event raised by the CompanyServerAccess contract.
type CompanyServerAccessSubscriptionExtended struct {
	Company    common.Address
	NewEndDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionExtended is a free log retrieval operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterSubscriptionExtended(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessSubscriptionExtendedIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "SubscriptionExtended", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessSubscriptionExtendedIterator{contract: _CompanyServerAccess.contract, event: "SubscriptionExtended", logs: logs, sub: sub}, nil
}

// WatchSubscriptionExtended is a free log subscription operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchSubscriptionExtended(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessSubscriptionExtended, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "SubscriptionExtended", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessSubscriptionExtended)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "SubscriptionExtended", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionExtended is a log parse operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseSubscriptionExtended(log types.Log) (*CompanyServerAccessSubscriptionExtended, error) {
	event := new(CompanyServerAccessSubscriptionExtended)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "SubscriptionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
