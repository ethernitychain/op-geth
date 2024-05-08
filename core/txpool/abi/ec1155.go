// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ec1155

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

// Ec1155MetaData contains all meta data concerning the Ec1155 contract.
var Ec1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516125c03803806125c0833981810160405281019061003191906100c7565b60405180602001604052805f81525061004f8161005660201b60201c565b50506103fb565b8060029081610065919061032c565b5050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100968261006d565b9050919050565b6100a68161008c565b81146100b0575f80fd5b50565b5f815190506100c18161009d565b92915050565b5f602082840312156100dc576100db610069565b5b5f6100e9848285016100b3565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061016d57607f821691505b6020821081036101805761017f610129565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101e27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826101a7565b6101ec86836101a7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61023061022b61022684610204565b61020d565b610204565b9050919050565b5f819050919050565b61024983610216565b61025d61025582610237565b8484546101b3565b825550505050565b5f90565b610271610265565b61027c818484610240565b505050565b5b8181101561029f576102945f82610269565b600181019050610282565b5050565b601f8211156102e4576102b581610186565b6102be84610198565b810160208510156102cd578190505b6102e16102d985610198565b830182610281565b50505b505050565b5f82821c905092915050565b5f6103045f19846008026102e9565b1980831691505092915050565b5f61031c83836102f5565b9150826002028217905092915050565b610335826100f2565b67ffffffffffffffff81111561034e5761034d6100fc565b5b6103588254610156565b6103638282856102a3565b5f60209050601f831160018114610394575f8415610382578287015190505b61038c8582610311565b8655506103f3565b601f1984166103a286610186565b5f5b828110156103c9578489015182556001820191506020850194506020810190506103a4565b868310156103e657848901516103e2601f8916826102f5565b8355505b6001600288020188555050505b505050505050565b6121b8806104085f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80632eb2c2d6116100645780632eb2c2d6146101405780634e1273f41461015c578063a22cb4651461018c578063e985e9c5146101a8578063f242432a146101d857610090565b8062fdd58e1461009457806301ffc9a7146100c457806302fe5305146100f45780630e89341c14610110575b5f80fd5b6100ae60048036038101906100a99190611335565b6101f4565b6040516100bb9190611382565b60405180910390f35b6100de60048036038101906100d991906113f0565b610249565b6040516100eb9190611435565b60405180910390f35b61010e6004803603810190610109919061158a565b61032a565b005b61012a600480360381019061012591906115d1565b610336565b604051610137919061165c565b60405180910390f35b61015a600480360381019061015591906117de565b6103c8565b005b61017660048036038101906101719190611969565b61046f565b6040516101839190611a96565b60405180910390f35b6101a660048036038101906101a19190611ae0565b610576565b005b6101c260048036038101906101bd9190611b1e565b61058c565b6040516101cf9190611435565b60405180910390f35b6101f260048036038101906101ed9190611b5c565b61061a565b005b5f805f8381526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7fd9b67a26000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061031357507f0e89341c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103235750610322826106c1565b5b9050919050565b6103338161072a565b50565b60606002805461034590611c1c565b80601f016020809104026020016040519081016040528092919081815260200182805461037190611c1c565b80156103bc5780601f10610393576101008083540402835291602001916103bc565b820191905f5260205f20905b81548152906001019060200180831161039f57829003601f168201915b50505050509050919050565b5f6103d161073d565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156104165750610414868261058c565b155b1561045a5780866040517fe237d922000000000000000000000000000000000000000000000000000000008152600401610451929190611c5b565b60405180910390fd5b6104678686868686610744565b505050505050565b606081518351146104bb57815183516040517f5b0599910000000000000000000000000000000000000000000000000000000081526004016104b2929190611c82565b60405180910390fd5b5f835167ffffffffffffffff8111156104d7576104d6611466565b5b6040519080825280602002602001820160405280156105055781602001602082028036833780820191505090505b5090505f5b845181101561056b57610541610529828761083890919063ffffffff16565b61053c838761084b90919063ffffffff16565b6101f4565b82828151811061055457610553611ca9565b5b60200260200101818152505080600101905061050a565b508091505092915050565b61058861058161073d565b838361085e565b5050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f61062361073d565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156106685750610666868261058c565b155b156106ac5780866040517fe237d9220000000000000000000000000000000000000000000000000000000081526004016106a3929190611c5b565b60405180910390fd5b6106b986868686866109c7565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b80600290816107399190611e73565b5050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036107b4575f6040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016107ab9190611f42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610824575f6040517f01a8351400000000000000000000000000000000000000000000000000000000815260040161081b9190611f42565b60405180910390fd5b6108318585858585610acd565b5050505050565b5f60208202602084010151905092915050565b5f60208202602084010151905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108ce575f6040517fced3e1000000000000000000000000000000000000000000000000000000000081526004016108c59190611f42565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109ba9190611435565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a37575f6040517f57f447ce000000000000000000000000000000000000000000000000000000008152600401610a2e9190611f42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610aa7575f6040517f01a83514000000000000000000000000000000000000000000000000000000008152600401610a9e9190611f42565b60405180910390fd5b5f80610ab38585610b79565b91509150610ac48787848487610acd565b50505050505050565b610ad985858585610ba9565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610b72575f610b1561073d565b90506001845103610b61575f610b345f8661084b90919063ffffffff16565b90505f610b4a5f8661084b90919063ffffffff16565b9050610b5a838989858589610f39565b5050610b70565b610b6f8187878787876110e8565b5b505b5050505050565b60608060405191506001825283602083015260408201905060018152826020820152604081016040529250929050565b8051825114610bf357815181516040517f5b059991000000000000000000000000000000000000000000000000000000008152600401610bea929190611c82565b60405180910390fd5b5f610bfc61073d565b90505f5b8351811015610df8575f610c1d828661084b90919063ffffffff16565b90505f610c33838661084b90919063ffffffff16565b90505f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610d56575f805f8481526020019081526020015f205f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610d0257888183856040517f03dee4c5000000000000000000000000000000000000000000000000000000008152600401610cf99493929190611f5b565b60405180910390fd5b8181035f808581526020019081526020015f205f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610deb57805f808481526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610de39190611fcb565b925050819055505b5050806001019050610c00565b506001835103610eb3575f610e165f8561084b90919063ffffffff16565b90505f610e2c5f8561084b90919063ffffffff16565b90508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051610ea4929190611c82565b60405180910390a45050610f32565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8686604051610f29929190611ffe565b60405180910390a45b5050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b11156110e0578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b8152600401610f99959493929190612085565b6020604051808303815f875af1925050508015610fd457506040513d601f19601f82011682018060405250810190610fd191906120f1565b60015b611055573d805f8114611002576040519150601f19603f3d011682016040523d82523d5f602084013e611007565b606091505b505f81510361104d57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016110449190611f42565b60405180910390fd5b805181602001fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146110de57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016110d59190611f42565b60405180910390fd5b505b505050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b111561128f578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b815260040161114895949392919061211c565b6020604051808303815f875af192505050801561118357506040513d601f19601f8201168201806040525081019061118091906120f1565b60015b611204573d805f81146111b1576040519150601f19603f3d011682016040523d82523d5f602084013e6111b6565b606091505b505f8151036111fc57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016111f39190611f42565b60405180910390fd5b805181602001fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461128d57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016112849190611f42565b60405180910390fd5b505b505050505050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112d1826112a8565b9050919050565b6112e1816112c7565b81146112eb575f80fd5b50565b5f813590506112fc816112d8565b92915050565b5f819050919050565b61131481611302565b811461131e575f80fd5b50565b5f8135905061132f8161130b565b92915050565b5f806040838503121561134b5761134a6112a0565b5b5f611358858286016112ee565b925050602061136985828601611321565b9150509250929050565b61137c81611302565b82525050565b5f6020820190506113955f830184611373565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113cf8161139b565b81146113d9575f80fd5b50565b5f813590506113ea816113c6565b92915050565b5f60208284031215611405576114046112a0565b5b5f611412848285016113dc565b91505092915050565b5f8115159050919050565b61142f8161141b565b82525050565b5f6020820190506114485f830184611426565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61149c82611456565b810181811067ffffffffffffffff821117156114bb576114ba611466565b5b80604052505050565b5f6114cd611297565b90506114d98282611493565b919050565b5f67ffffffffffffffff8211156114f8576114f7611466565b5b61150182611456565b9050602081019050919050565b828183375f83830152505050565b5f61152e611529846114de565b6114c4565b90508281526020810184848401111561154a57611549611452565b5b61155584828561150e565b509392505050565b5f82601f8301126115715761157061144e565b5b813561158184826020860161151c565b91505092915050565b5f6020828403121561159f5761159e6112a0565b5b5f82013567ffffffffffffffff8111156115bc576115bb6112a4565b5b6115c88482850161155d565b91505092915050565b5f602082840312156115e6576115e56112a0565b5b5f6115f384828501611321565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61162e826115fc565b6116388185611606565b9350611648818560208601611616565b61165181611456565b840191505092915050565b5f6020820190508181035f8301526116748184611624565b905092915050565b5f67ffffffffffffffff82111561169657611695611466565b5b602082029050602081019050919050565b5f80fd5b5f6116bd6116b88461167c565b6114c4565b905080838252602082019050602084028301858111156116e0576116df6116a7565b5b835b8181101561170957806116f58882611321565b8452602084019350506020810190506116e2565b5050509392505050565b5f82601f8301126117275761172661144e565b5b81356117378482602086016116ab565b91505092915050565b5f67ffffffffffffffff82111561175a57611759611466565b5b61176382611456565b9050602081019050919050565b5f61178261177d84611740565b6114c4565b90508281526020810184848401111561179e5761179d611452565b5b6117a984828561150e565b509392505050565b5f82601f8301126117c5576117c461144e565b5b81356117d5848260208601611770565b91505092915050565b5f805f805f60a086880312156117f7576117f66112a0565b5b5f611804888289016112ee565b9550506020611815888289016112ee565b945050604086013567ffffffffffffffff811115611836576118356112a4565b5b61184288828901611713565b935050606086013567ffffffffffffffff811115611863576118626112a4565b5b61186f88828901611713565b925050608086013567ffffffffffffffff8111156118905761188f6112a4565b5b61189c888289016117b1565b9150509295509295909350565b5f67ffffffffffffffff8211156118c3576118c2611466565b5b602082029050602081019050919050565b5f6118e66118e1846118a9565b6114c4565b90508083825260208201905060208402830185811115611909576119086116a7565b5b835b81811015611932578061191e88826112ee565b84526020840193505060208101905061190b565b5050509392505050565b5f82601f8301126119505761194f61144e565b5b81356119608482602086016118d4565b91505092915050565b5f806040838503121561197f5761197e6112a0565b5b5f83013567ffffffffffffffff81111561199c5761199b6112a4565b5b6119a88582860161193c565b925050602083013567ffffffffffffffff8111156119c9576119c86112a4565b5b6119d585828601611713565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611a1181611302565b82525050565b5f611a228383611a08565b60208301905092915050565b5f602082019050919050565b5f611a44826119df565b611a4e81856119e9565b9350611a59836119f9565b805f5b83811015611a89578151611a708882611a17565b9750611a7b83611a2e565b925050600181019050611a5c565b5085935050505092915050565b5f6020820190508181035f830152611aae8184611a3a565b905092915050565b611abf8161141b565b8114611ac9575f80fd5b50565b5f81359050611ada81611ab6565b92915050565b5f8060408385031215611af657611af56112a0565b5b5f611b03858286016112ee565b9250506020611b1485828601611acc565b9150509250929050565b5f8060408385031215611b3457611b336112a0565b5b5f611b41858286016112ee565b9250506020611b52858286016112ee565b9150509250929050565b5f805f805f60a08688031215611b7557611b746112a0565b5b5f611b82888289016112ee565b9550506020611b93888289016112ee565b9450506040611ba488828901611321565b9350506060611bb588828901611321565b925050608086013567ffffffffffffffff811115611bd657611bd56112a4565b5b611be2888289016117b1565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611c3357607f821691505b602082108103611c4657611c45611bef565b5b50919050565b611c55816112c7565b82525050565b5f604082019050611c6e5f830185611c4c565b611c7b6020830184611c4c565b9392505050565b5f604082019050611c955f830185611373565b611ca26020830184611373565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611d327fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611cf7565b611d3c8683611cf7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611d77611d72611d6d84611302565b611d54565b611302565b9050919050565b5f819050919050565b611d9083611d5d565b611da4611d9c82611d7e565b848454611d03565b825550505050565b5f90565b611db8611dac565b611dc3818484611d87565b505050565b5b81811015611de657611ddb5f82611db0565b600181019050611dc9565b5050565b601f821115611e2b57611dfc81611cd6565b611e0584611ce8565b81016020851015611e14578190505b611e28611e2085611ce8565b830182611dc8565b50505b505050565b5f82821c905092915050565b5f611e4b5f1984600802611e30565b1980831691505092915050565b5f611e638383611e3c565b9150826002028217905092915050565b611e7c826115fc565b67ffffffffffffffff811115611e9557611e94611466565b5b611e9f8254611c1c565b611eaa828285611dea565b5f60209050601f831160018114611edb575f8415611ec9578287015190505b611ed38582611e58565b865550611f3a565b601f198416611ee986611cd6565b5f5b82811015611f1057848901518255600182019150602085019450602081019050611eeb565b86831015611f2d5784890151611f29601f891682611e3c565b8355505b6001600288020188555050505b505050505050565b5f602082019050611f555f830184611c4c565b92915050565b5f608082019050611f6e5f830187611c4c565b611f7b6020830186611373565b611f886040830185611373565b611f956060830184611373565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fd582611302565b9150611fe083611302565b9250828201905080821115611ff857611ff7611f9e565b5b92915050565b5f6040820190508181035f8301526120168185611a3a565b9050818103602083015261202a8184611a3a565b90509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f61205782612033565b612061818561203d565b9350612071818560208601611616565b61207a81611456565b840191505092915050565b5f60a0820190506120985f830188611c4c565b6120a56020830187611c4c565b6120b26040830186611373565b6120bf6060830185611373565b81810360808301526120d1818461204d565b90509695505050505050565b5f815190506120eb816113c6565b92915050565b5f60208284031215612106576121056112a0565b5b5f612113848285016120dd565b91505092915050565b5f60a08201905061212f5f830188611c4c565b61213c6020830187611c4c565b818103604083015261214e8186611a3a565b905081810360608301526121628185611a3a565b90508181036080830152612176818461204d565b9050969550505050505056fea2646970667358221220d556faaf605d8719cd332cf99979b8a37fbee22ce0d290b7e4c328266371f72e64736f6c63430008190033",
}

// Ec1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ec1155MetaData.ABI instead.
var Ec1155ABI = Ec1155MetaData.ABI

// Ec1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Ec1155MetaData.Bin instead.
var Ec1155Bin = Ec1155MetaData.Bin

// DeployEc1155 deploys a new Ethereum contract, binding an instance of Ec1155 to it.
func DeployEc1155(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *Ec1155, error) {
	parsed, err := Ec1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Ec1155Bin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ec1155{Ec1155Caller: Ec1155Caller{contract: contract}, Ec1155Transactor: Ec1155Transactor{contract: contract}, Ec1155Filterer: Ec1155Filterer{contract: contract}}, nil
}

// Ec1155 is an auto generated Go binding around an Ethereum contract.
type Ec1155 struct {
	Ec1155Caller     // Read-only binding to the contract
	Ec1155Transactor // Write-only binding to the contract
	Ec1155Filterer   // Log filterer for contract events
}

// Ec1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ec1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ec1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ec1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ec1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ec1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ec1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ec1155Session struct {
	Contract     *Ec1155           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ec1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ec1155CallerSession struct {
	Contract *Ec1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Ec1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ec1155TransactorSession struct {
	Contract     *Ec1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ec1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ec1155Raw struct {
	Contract *Ec1155 // Generic contract binding to access the raw methods on
}

// Ec1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ec1155CallerRaw struct {
	Contract *Ec1155Caller // Generic read-only contract binding to access the raw methods on
}

// Ec1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ec1155TransactorRaw struct {
	Contract *Ec1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEc1155 creates a new instance of Ec1155, bound to a specific deployed contract.
func NewEc1155(address common.Address, backend bind.ContractBackend) (*Ec1155, error) {
	contract, err := bindEc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ec1155{Ec1155Caller: Ec1155Caller{contract: contract}, Ec1155Transactor: Ec1155Transactor{contract: contract}, Ec1155Filterer: Ec1155Filterer{contract: contract}}, nil
}

// NewEc1155Caller creates a new read-only instance of Ec1155, bound to a specific deployed contract.
func NewEc1155Caller(address common.Address, caller bind.ContractCaller) (*Ec1155Caller, error) {
	contract, err := bindEc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ec1155Caller{contract: contract}, nil
}

// NewEc1155Transactor creates a new write-only instance of Ec1155, bound to a specific deployed contract.
func NewEc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Ec1155Transactor, error) {
	contract, err := bindEc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ec1155Transactor{contract: contract}, nil
}

// NewEc1155Filterer creates a new log filterer instance of Ec1155, bound to a specific deployed contract.
func NewEc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Ec1155Filterer, error) {
	contract, err := bindEc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ec1155Filterer{contract: contract}, nil
}

// bindEc1155 binds a generic wrapper to an already deployed contract.
func bindEc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ec1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ec1155 *Ec1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ec1155.Contract.Ec1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ec1155 *Ec1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ec1155.Contract.Ec1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ec1155 *Ec1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ec1155.Contract.Ec1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ec1155 *Ec1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ec1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ec1155 *Ec1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ec1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ec1155 *Ec1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ec1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Ec1155 *Ec1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ec1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Ec1155 *Ec1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Ec1155.Contract.BalanceOf(&_Ec1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Ec1155 *Ec1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Ec1155.Contract.BalanceOf(&_Ec1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Ec1155 *Ec1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Ec1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Ec1155 *Ec1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Ec1155.Contract.BalanceOfBatch(&_Ec1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Ec1155 *Ec1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Ec1155.Contract.BalanceOfBatch(&_Ec1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Ec1155 *Ec1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ec1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Ec1155 *Ec1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Ec1155.Contract.IsApprovedForAll(&_Ec1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Ec1155 *Ec1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Ec1155.Contract.IsApprovedForAll(&_Ec1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ec1155 *Ec1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ec1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ec1155 *Ec1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ec1155.Contract.SupportsInterface(&_Ec1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ec1155 *Ec1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ec1155.Contract.SupportsInterface(&_Ec1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Ec1155 *Ec1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Ec1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Ec1155 *Ec1155Session) Uri(arg0 *big.Int) (string, error) {
	return _Ec1155.Contract.Uri(&_Ec1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Ec1155 *Ec1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Ec1155.Contract.Uri(&_Ec1155.CallOpts, arg0)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Ec1155 *Ec1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Ec1155 *Ec1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.Contract.SafeBatchTransferFrom(&_Ec1155.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Ec1155 *Ec1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.Contract.SafeBatchTransferFrom(&_Ec1155.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Ec1155 *Ec1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Ec1155 *Ec1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.Contract.SafeTransferFrom(&_Ec1155.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Ec1155 *Ec1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Ec1155.Contract.SafeTransferFrom(&_Ec1155.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ec1155 *Ec1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ec1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ec1155 *Ec1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ec1155.Contract.SetApprovalForAll(&_Ec1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ec1155 *Ec1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ec1155.Contract.SetApprovalForAll(&_Ec1155.TransactOpts, operator, approved)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_Ec1155 *Ec1155Transactor) SetURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _Ec1155.contract.Transact(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_Ec1155 *Ec1155Session) SetURI(newuri string) (*types.Transaction, error) {
	return _Ec1155.Contract.SetURI(&_Ec1155.TransactOpts, newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_Ec1155 *Ec1155TransactorSession) SetURI(newuri string) (*types.Transaction, error) {
	return _Ec1155.Contract.SetURI(&_Ec1155.TransactOpts, newuri)
}

// Ec1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ec1155 contract.
type Ec1155ApprovalForAllIterator struct {
	Event *Ec1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Ec1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ec1155ApprovalForAll)
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
		it.Event = new(Ec1155ApprovalForAll)
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
func (it *Ec1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ec1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ec1155ApprovalForAll represents a ApprovalForAll event raised by the Ec1155 contract.
type Ec1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Ec1155 *Ec1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Ec1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ec1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Ec1155ApprovalForAllIterator{contract: _Ec1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Ec1155 *Ec1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Ec1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ec1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ec1155ApprovalForAll)
				if err := _Ec1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Ec1155 *Ec1155Filterer) ParseApprovalForAll(log types.Log) (*Ec1155ApprovalForAll, error) {
	event := new(Ec1155ApprovalForAll)
	if err := _Ec1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ec1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Ec1155 contract.
type Ec1155TransferBatchIterator struct {
	Event *Ec1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *Ec1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ec1155TransferBatch)
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
		it.Event = new(Ec1155TransferBatch)
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
func (it *Ec1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ec1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ec1155TransferBatch represents a TransferBatch event raised by the Ec1155 contract.
type Ec1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Ec1155 *Ec1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Ec1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ec1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Ec1155TransferBatchIterator{contract: _Ec1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Ec1155 *Ec1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Ec1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ec1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ec1155TransferBatch)
				if err := _Ec1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Ec1155 *Ec1155Filterer) ParseTransferBatch(log types.Log) (*Ec1155TransferBatch, error) {
	event := new(Ec1155TransferBatch)
	if err := _Ec1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ec1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Ec1155 contract.
type Ec1155TransferSingleIterator struct {
	Event *Ec1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *Ec1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ec1155TransferSingle)
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
		it.Event = new(Ec1155TransferSingle)
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
func (it *Ec1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ec1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ec1155TransferSingle represents a TransferSingle event raised by the Ec1155 contract.
type Ec1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Ec1155 *Ec1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Ec1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ec1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Ec1155TransferSingleIterator{contract: _Ec1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Ec1155 *Ec1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Ec1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ec1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ec1155TransferSingle)
				if err := _Ec1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Ec1155 *Ec1155Filterer) ParseTransferSingle(log types.Log) (*Ec1155TransferSingle, error) {
	event := new(Ec1155TransferSingle)
	if err := _Ec1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ec1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Ec1155 contract.
type Ec1155URIIterator struct {
	Event *Ec1155URI // Event containing the contract specifics and raw log

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
func (it *Ec1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ec1155URI)
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
		it.Event = new(Ec1155URI)
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
func (it *Ec1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ec1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ec1155URI represents a URI event raised by the Ec1155 contract.
type Ec1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Ec1155 *Ec1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Ec1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ec1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Ec1155URIIterator{contract: _Ec1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Ec1155 *Ec1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Ec1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ec1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ec1155URI)
				if err := _Ec1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Ec1155 *Ec1155Filterer) ParseURI(log types.Log) (*Ec1155URI, error) {
	event := new(Ec1155URI)
	if err := _Ec1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
