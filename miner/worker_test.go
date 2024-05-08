// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package miner

import (
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// testCode is the testing contract binary code which will initialises some
	// variables in constructor
	testCode  = "0x02fe53050000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000e67646a736867646a736764676a73000000000000000000000000000000000000"
	testCode2 = "0x02fe53050000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000e67646a736867646a736764676a73000000000000000000000000000000000000"
	// testCode2 = "0x608060405234801561000f575f80fd5b506040518060400160405280601281526020017f68697468697369736e6f74636f727265637400000000000000000000000000008152506100558161005b60201b60201c565b50610377565b806002908161006a91906102a8565b5050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100e957607f821691505b6020821081036100fc576100fb6100a5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261015e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610123565b6101688683610123565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101ac6101a76101a284610180565b610189565b610180565b9050919050565b5f819050919050565b6101c583610192565b6101d96101d1826101b3565b84845461012f565b825550505050565b5f90565b6101ed6101e1565b6101f88184846101bc565b505050565b5b8181101561021b576102105f826101e5565b6001810190506101fe565b5050565b601f8211156102605761023181610102565b61023a84610114565b81016020851015610249578190505b61025d61025585610114565b8301826101fd565b50505b505050565b5f82821c905092915050565b5f6102805f1984600802610265565b1980831691505092915050565b5f6102988383610271565b9150826002028217905092915050565b6102b18261006e565b67ffffffffffffffff8111156102ca576102c9610078565b5b6102d482546100d2565b6102df82828561021f565b5f60209050601f831160018114610310575f84156102fe578287015190505b610308858261028d565b86555061036f565b601f19841661031e86610102565b5f5b8281101561034557848901518255600182019150602085019450602081019050610320565b86831015610362578489015161035e601f891682610271565b8355505b6001600288020188555050505b505050505050565b6121b8806103845f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80632eb2c2d6116100645780632eb2c2d6146101405780634e1273f41461015c578063a22cb4651461018c578063e985e9c5146101a8578063f242432a146101d857610090565b8062fdd58e1461009457806301ffc9a7146100c457806302fe5305146100f45780630e89341c14610110575b5f80fd5b6100ae60048036038101906100a99190611335565b6101f4565b6040516100bb9190611382565b60405180910390f35b6100de60048036038101906100d991906113f0565b610249565b6040516100eb9190611435565b60405180910390f35b61010e6004803603810190610109919061158a565b61032a565b005b61012a600480360381019061012591906115d1565b610336565b604051610137919061165c565b60405180910390f35b61015a600480360381019061015591906117de565b6103c8565b005b61017660048036038101906101719190611969565b61046f565b6040516101839190611a96565b60405180910390f35b6101a660048036038101906101a19190611ae0565b610576565b005b6101c260048036038101906101bd9190611b1e565b61058c565b6040516101cf9190611435565b60405180910390f35b6101f260048036038101906101ed9190611b5c565b61061a565b005b5f805f8381526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7fd9b67a26000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061031357507f0e89341c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103235750610322826106c1565b5b9050919050565b6103338161072a565b50565b60606002805461034590611c1c565b80601f016020809104026020016040519081016040528092919081815260200182805461037190611c1c565b80156103bc5780601f10610393576101008083540402835291602001916103bc565b820191905f5260205f20905b81548152906001019060200180831161039f57829003601f168201915b50505050509050919050565b5f6103d161073d565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156104165750610414868261058c565b155b1561045a5780866040517fe237d922000000000000000000000000000000000000000000000000000000008152600401610451929190611c5b565b60405180910390fd5b6104678686868686610744565b505050505050565b606081518351146104bb57815183516040517f5b0599910000000000000000000000000000000000000000000000000000000081526004016104b2929190611c82565b60405180910390fd5b5f835167ffffffffffffffff8111156104d7576104d6611466565b5b6040519080825280602002602001820160405280156105055781602001602082028036833780820191505090505b5090505f5b845181101561056b57610541610529828761083890919063ffffffff16565b61053c838761084b90919063ffffffff16565b6101f4565b82828151811061055457610553611ca9565b5b60200260200101818152505080600101905061050a565b508091505092915050565b61058861058161073d565b838361085e565b5050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f61062361073d565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156106685750610666868261058c565b155b156106ac5780866040517fe237d9220000000000000000000000000000000000000000000000000000000081526004016106a3929190611c5b565b60405180910390fd5b6106b986868686866109c7565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b80600290816107399190611e73565b5050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036107b4575f6040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016107ab9190611f42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610824575f6040517f01a8351400000000000000000000000000000000000000000000000000000000815260040161081b9190611f42565b60405180910390fd5b6108318585858585610acd565b5050505050565b5f60208202602084010151905092915050565b5f60208202602084010151905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108ce575f6040517fced3e1000000000000000000000000000000000000000000000000000000000081526004016108c59190611f42565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109ba9190611435565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a37575f6040517f57f447ce000000000000000000000000000000000000000000000000000000008152600401610a2e9190611f42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610aa7575f6040517f01a83514000000000000000000000000000000000000000000000000000000008152600401610a9e9190611f42565b60405180910390fd5b5f80610ab38585610b79565b91509150610ac48787848487610acd565b50505050505050565b610ad985858585610ba9565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610b72575f610b1561073d565b90506001845103610b61575f610b345f8661084b90919063ffffffff16565b90505f610b4a5f8661084b90919063ffffffff16565b9050610b5a838989858589610f39565b5050610b70565b610b6f8187878787876110e8565b5b505b5050505050565b60608060405191506001825283602083015260408201905060018152826020820152604081016040529250929050565b8051825114610bf357815181516040517f5b059991000000000000000000000000000000000000000000000000000000008152600401610bea929190611c82565b60405180910390fd5b5f610bfc61073d565b90505f5b8351811015610df8575f610c1d828661084b90919063ffffffff16565b90505f610c33838661084b90919063ffffffff16565b90505f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610d56575f805f8481526020019081526020015f205f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610d0257888183856040517f03dee4c5000000000000000000000000000000000000000000000000000000008152600401610cf99493929190611f5b565b60405180910390fd5b8181035f808581526020019081526020015f205f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610deb57805f808481526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610de39190611fcb565b925050819055505b5050806001019050610c00565b506001835103610eb3575f610e165f8561084b90919063ffffffff16565b90505f610e2c5f8561084b90919063ffffffff16565b90508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051610ea4929190611c82565b60405180910390a45050610f32565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8686604051610f29929190611ffe565b60405180910390a45b5050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b11156110e0578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b8152600401610f99959493929190612085565b6020604051808303815f875af1925050508015610fd457506040513d601f19601f82011682018060405250810190610fd191906120f1565b60015b611055573d805f8114611002576040519150601f19603f3d011682016040523d82523d5f602084013e611007565b606091505b505f81510361104d57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016110449190611f42565b60405180910390fd5b805181602001fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146110de57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016110d59190611f42565b60405180910390fd5b505b505050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b111561128f578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b815260040161114895949392919061211c565b6020604051808303815f875af192505050801561118357506040513d601f19601f8201168201806040525081019061118091906120f1565b60015b611204573d805f81146111b1576040519150601f19603f3d011682016040523d82523d5f602084013e6111b6565b606091505b505f8151036111fc57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016111f39190611f42565b60405180910390fd5b805181602001fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461128d57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016112849190611f42565b60405180910390fd5b505b505050505050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112d1826112a8565b9050919050565b6112e1816112c7565b81146112eb575f80fd5b50565b5f813590506112fc816112d8565b92915050565b5f819050919050565b61131481611302565b811461131e575f80fd5b50565b5f8135905061132f8161130b565b92915050565b5f806040838503121561134b5761134a6112a0565b5b5f611358858286016112ee565b925050602061136985828601611321565b9150509250929050565b61137c81611302565b82525050565b5f6020820190506113955f830184611373565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113cf8161139b565b81146113d9575f80fd5b50565b5f813590506113ea816113c6565b92915050565b5f60208284031215611405576114046112a0565b5b5f611412848285016113dc565b91505092915050565b5f8115159050919050565b61142f8161141b565b82525050565b5f6020820190506114485f830184611426565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61149c82611456565b810181811067ffffffffffffffff821117156114bb576114ba611466565b5b80604052505050565b5f6114cd611297565b90506114d98282611493565b919050565b5f67ffffffffffffffff8211156114f8576114f7611466565b5b61150182611456565b9050602081019050919050565b828183375f83830152505050565b5f61152e611529846114de565b6114c4565b90508281526020810184848401111561154a57611549611452565b5b61155584828561150e565b509392505050565b5f82601f8301126115715761157061144e565b5b813561158184826020860161151c565b91505092915050565b5f6020828403121561159f5761159e6112a0565b5b5f82013567ffffffffffffffff8111156115bc576115bb6112a4565b5b6115c88482850161155d565b91505092915050565b5f602082840312156115e6576115e56112a0565b5b5f6115f384828501611321565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61162e826115fc565b6116388185611606565b9350611648818560208601611616565b61165181611456565b840191505092915050565b5f6020820190508181035f8301526116748184611624565b905092915050565b5f67ffffffffffffffff82111561169657611695611466565b5b602082029050602081019050919050565b5f80fd5b5f6116bd6116b88461167c565b6114c4565b905080838252602082019050602084028301858111156116e0576116df6116a7565b5b835b8181101561170957806116f58882611321565b8452602084019350506020810190506116e2565b5050509392505050565b5f82601f8301126117275761172661144e565b5b81356117378482602086016116ab565b91505092915050565b5f67ffffffffffffffff82111561175a57611759611466565b5b61176382611456565b9050602081019050919050565b5f61178261177d84611740565b6114c4565b90508281526020810184848401111561179e5761179d611452565b5b6117a984828561150e565b509392505050565b5f82601f8301126117c5576117c461144e565b5b81356117d5848260208601611770565b91505092915050565b5f805f805f60a086880312156117f7576117f66112a0565b5b5f611804888289016112ee565b9550506020611815888289016112ee565b945050604086013567ffffffffffffffff811115611836576118356112a4565b5b61184288828901611713565b935050606086013567ffffffffffffffff811115611863576118626112a4565b5b61186f88828901611713565b925050608086013567ffffffffffffffff8111156118905761188f6112a4565b5b61189c888289016117b1565b9150509295509295909350565b5f67ffffffffffffffff8211156118c3576118c2611466565b5b602082029050602081019050919050565b5f6118e66118e1846118a9565b6114c4565b90508083825260208201905060208402830185811115611909576119086116a7565b5b835b81811015611932578061191e88826112ee565b84526020840193505060208101905061190b565b5050509392505050565b5f82601f8301126119505761194f61144e565b5b81356119608482602086016118d4565b91505092915050565b5f806040838503121561197f5761197e6112a0565b5b5f83013567ffffffffffffffff81111561199c5761199b6112a4565b5b6119a88582860161193c565b925050602083013567ffffffffffffffff8111156119c9576119c86112a4565b5b6119d585828601611713565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611a1181611302565b82525050565b5f611a228383611a08565b60208301905092915050565b5f602082019050919050565b5f611a44826119df565b611a4e81856119e9565b9350611a59836119f9565b805f5b83811015611a89578151611a708882611a17565b9750611a7b83611a2e565b925050600181019050611a5c565b5085935050505092915050565b5f6020820190508181035f830152611aae8184611a3a565b905092915050565b611abf8161141b565b8114611ac9575f80fd5b50565b5f81359050611ada81611ab6565b92915050565b5f8060408385031215611af657611af56112a0565b5b5f611b03858286016112ee565b9250506020611b1485828601611acc565b9150509250929050565b5f8060408385031215611b3457611b336112a0565b5b5f611b41858286016112ee565b9250506020611b52858286016112ee565b9150509250929050565b5f805f805f60a08688031215611b7557611b746112a0565b5b5f611b82888289016112ee565b9550506020611b93888289016112ee565b9450506040611ba488828901611321565b9350506060611bb588828901611321565b925050608086013567ffffffffffffffff811115611bd657611bd56112a4565b5b611be2888289016117b1565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611c3357607f821691505b602082108103611c4657611c45611bef565b5b50919050565b611c55816112c7565b82525050565b5f604082019050611c6e5f830185611c4c565b611c7b6020830184611c4c565b9392505050565b5f604082019050611c955f830185611373565b611ca26020830184611373565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611d327fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611cf7565b611d3c8683611cf7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611d77611d72611d6d84611302565b611d54565b611302565b9050919050565b5f819050919050565b611d9083611d5d565b611da4611d9c82611d7e565b848454611d03565b825550505050565b5f90565b611db8611dac565b611dc3818484611d87565b505050565b5b81811015611de657611ddb5f82611db0565b600181019050611dc9565b5050565b601f821115611e2b57611dfc81611cd6565b611e0584611ce8565b81016020851015611e14578190505b611e28611e2085611ce8565b830182611dc8565b50505b505050565b5f82821c905092915050565b5f611e4b5f1984600802611e30565b1980831691505092915050565b5f611e638383611e3c565b9150826002028217905092915050565b611e7c826115fc565b67ffffffffffffffff811115611e9557611e94611466565b5b611e9f8254611c1c565b611eaa828285611dea565b5f60209050601f831160018114611edb575f8415611ec9578287015190505b611ed38582611e58565b865550611f3a565b601f198416611ee986611cd6565b5f5b82811015611f1057848901518255600182019150602085019450602081019050611eeb565b86831015611f2d5784890151611f29601f891682611e3c565b8355505b6001600288020188555050505b505050505050565b5f602082019050611f555f830184611c4c565b92915050565b5f608082019050611f6e5f830187611c4c565b611f7b6020830186611373565b611f886040830185611373565b611f956060830184611373565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fd582611302565b9150611fe083611302565b9250828201905080821115611ff857611ff7611f9e565b5b92915050565b5f6040820190508181035f8301526120168185611a3a565b9050818103602083015261202a8184611a3a565b90509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f61205782612033565b612061818561203d565b9350612071818560208601611616565b61207a81611456565b840191505092915050565b5f60a0820190506120985f830188611c4c565b6120a56020830187611c4c565b6120b26040830186611373565b6120bf6060830185611373565b81810360808301526120d1818461204d565b90509695505050505050565b5f815190506120eb816113c6565b92915050565b5f60208284031215612106576121056112a0565b5b5f612113848285016120dd565b91505092915050565b5f60a08201905061212f5f830188611c4c565b61213c6020830187611c4c565b818103604083015261214e8186611a3a565b905081810360608301526121628185611a3a565b90508181036080830152612176818461204d565b9050969550505050505056fea2646970667358221220df76bd4aa4fd8e9f4beefd9fb4b46b642e7980fee67a0c29ca41e35e5a4a04c664736f6c63430008190033"
	// testGas is the gas required for contract deployment.
	testGas  = 144109
	testGas2 = 2517211
)

var (
	// Test chain configurations
	testTxPoolConfig  legacypool.Config
	ethashChainConfig *params.ChainConfig
	cliqueChainConfig *params.ChainConfig

	// Test accounts
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	testBankFunds   = big.NewInt(1000000000000000000)

	testUserKey, _  = crypto.GenerateKey()
	testUserAddress = crypto.PubkeyToAddress(testUserKey.PublicKey)

	// Test transactions
	pendingTxs []*types.Transaction
	newTxs     []*types.Transaction

	testConfig = &Config{
		Recommit: time.Second,
		GasCeil:  params.GenesisGasLimit,
	}
)

func init() {
	testTxPoolConfig = legacypool.DefaultConfig
	testTxPoolConfig.Journal = ""
	ethashChainConfig = new(params.ChainConfig)
	*ethashChainConfig = *params.TestChainConfig
	cliqueChainConfig = new(params.ChainConfig)
	*cliqueChainConfig = *params.TestChainConfig
	cliqueChainConfig.Clique = &params.CliqueConfig{
		Period: 10,
		Epoch:  30000,
	}

	signer := types.LatestSigner(params.TestChainConfig)
	tx1 := types.MustSignNewTx(testBankKey, signer, &types.AccessListTx{
		ChainID:  params.TestChainConfig.ChainID,
		Nonce:    0,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	tx3 := types.MustSignNewTx(testBankKey, signer, &types.AccessListTx{
		ChainID:  params.TestChainConfig.ChainID,
		Nonce:    0,
		To:       &testUserAddress,
		Value:    big.NewInt(0),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
		Data:     []byte(testCode2),
	})
	tx4 := types.MustSignNewTx(testBankKey, signer, &types.AccessListTx{
		ChainID:  params.TestChainConfig.ChainID,
		Nonce:    0,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
		Data:     []byte(testCode2),
	})

	pendingTxs = append(pendingTxs, tx1)
	pendingTxs = append(pendingTxs, tx3)
	pendingTxs = append(pendingTxs, tx4)
	// pendingTxs = append(pendingTxs, tx4)

	tx2 := types.MustSignNewTx(testBankKey, signer, &types.LegacyTx{
		Nonce:    1,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	newTxs = append(newTxs, tx2)
}

// testWorkerBackend implements worker.Backend interfaces and wraps all information needed during the testing.
type testWorkerBackend struct {
	db      ethdb.Database
	txPool  *txpool.TxPool
	chain   *core.BlockChain
	genesis *core.Genesis
}

func newTestWorkerBackend(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, n int) *testWorkerBackend {
	var gspec = &core.Genesis{
		Config: chainConfig,
		Alloc:  core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}},
	}
	switch e := engine.(type) {
	case *clique.Clique:
		gspec.ExtraData = make([]byte, 32+common.AddressLength+crypto.SignatureLength)
		copy(gspec.ExtraData[32:32+common.AddressLength], testBankAddress.Bytes())
		e.Authorize(testBankAddress, func(account accounts.Account, s string, data []byte) ([]byte, error) {
			return crypto.Sign(crypto.Keccak256(data), testBankKey)
		})
	case *ethash.Ethash:
	default:
		t.Fatalf("unexpected consensus engine type: %T", engine)
	}
	chain, err := core.NewBlockChain(db, &core.CacheConfig{TrieDirtyDisabled: true}, gspec, nil, engine, vm.Config{}, nil, nil)
	if err != nil {
		t.Fatalf("core.NewBlockChain failed: %v", err)
	}
	pool := legacypool.New(testTxPoolConfig, chain)
	txpool, _ := txpool.New(new(big.Int).SetUint64(testTxPoolConfig.PriceLimit), chain, []txpool.SubPool{pool})

	return &testWorkerBackend{
		db:      db,
		chain:   chain,
		txPool:  txpool,
		genesis: gspec,
	}
}

func (b *testWorkerBackend) BlockChain() *core.BlockChain { return b.chain }
func (b *testWorkerBackend) TxPool() *txpool.TxPool       { return b.txPool }

func (b *testWorkerBackend) newRandomTx(creation bool) *types.Transaction {
	var tx *types.Transaction
	gasPrice := big.NewInt(10 * params.InitialBaseFee)

	if creation {
		tx, _ = types.SignTx(types.NewContractCreation(b.txPool.Nonce(testBankAddress), big.NewInt(0), testGas, gasPrice, common.FromHex(testCode)), types.HomesteadSigner{}, testBankKey)
	} else {
		tx, _ = types.SignTx(types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, gasPrice, nil), types.HomesteadSigner{}, testBankKey)
	}
	return tx
}

func (b *testWorkerBackend) newRandomTx2(creation bool) *types.Transaction {
	var tx *types.Transaction
	gasPrice := big.NewInt(10 * params.InitialBaseFee)

	if creation {
		// tx, _ = types.SignTx(types.NewContractCreation(b.txPool.Nonce(testBankAddress), big.NewInt(0), testGas2, gasPrice, common.FromHex(testCode2)), types.HomesteadSigner{}, testBankKey)
		tx, _ = types.SignTx(types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(0), params.TxGas, gasPrice, common.FromHex(testCode2)), types.HomesteadSigner{}, testBankKey)
	}
	return tx
}

func newTestWorker(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, blocks int) (*worker, *testWorkerBackend) {
	backend := newTestWorkerBackend(t, chainConfig, engine, db, blocks)
	backend.txPool.Add(newTxs, true, false)
	w := newWorker(testConfig, chainConfig, engine, backend, new(event.TypeMux), nil, false)
	w.setEtherbase(testBankAddress)
	return w, backend
}

func TestGenerateAndImportBlock(t *testing.T) {
	t.Parallel()
	var (
		db     = rawdb.NewMemoryDatabase()
		config = *params.AllCliqueProtocolChanges
	)
	config.Clique = &params.CliqueConfig{Period: 1, Epoch: 30000}
	engine := clique.New(config.Clique, db)

	w, b := newTestWorker(t, &config, engine, db, 0)
	defer w.close()

	// This test chain imports the mined blocks.
	chain, _ := core.NewBlockChain(rawdb.NewMemoryDatabase(), nil, b.genesis, nil, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	// Ignore empty commit here for less noise.
	w.skipSealHook = func(task *task) bool {
		return len(task.receipts) == 0
	}

	// Wait for mined blocks.
	sub := w.mux.Subscribe(core.NewMinedBlockEvent{})
	defer sub.Unsubscribe()

	// Start mining!
	w.start()

	for i := 0; i < 5; i++ {

		b.txPool.Add([]*types.Transaction{b.newRandomTx2(true)}, true, false)
		b.txPool.Add([]*types.Transaction{b.newRandomTx(true)}, true, false)
		b.txPool.Add([]*types.Transaction{b.newRandomTx(false)}, true, false)
		b.txPool.Add([]*types.Transaction{b.newRandomTx2(true)}, true, false)

		select {
		case ev := <-sub.Chan():
			block := ev.Data.(core.NewMinedBlockEvent).Block
			if _, err := chain.InsertChain([]*types.Block{block}); err != nil {
				t.Fatalf("failed to insert new mined block %d: %v", block.NumberU64(), err)
			}
		case <-time.After(3 * time.Second): // Worker needs 1s to include new changes.
			t.Fatalf("timeout")
		}
	}

}

func TestEmptyWorkEthash(t *testing.T) {
	t.Parallel()
	testEmptyWork(t, ethashChainConfig, ethash.NewFaker())
}
func TestEmptyWorkClique(t *testing.T) {
	t.Parallel()
	testEmptyWork(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func testEmptyWork(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, _ := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	taskCh := make(chan struct{}, 2)
	checkEqual := func(t *testing.T, task *task) {
		// The work should contain 1 tx
		receiptLen, balance := 1, big.NewInt(1000)
		if len(task.receipts) != receiptLen {
			t.Fatalf("receipt number mismatch: have %d, want %d", len(task.receipts), receiptLen)
		}
		if task.state.GetBalance(testUserAddress).Cmp(balance) != 0 {
			t.Fatalf("account balance mismatch: have %d, want %d", task.state.GetBalance(testUserAddress), balance)
		}
	}
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 1 {
			checkEqual(t, task)
			taskCh <- struct{}{}
		}
	}
	w.skipSealHook = func(task *task) bool { return true }
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	w.start() // Start mining!
	select {
	case <-taskCh:
	case <-time.NewTimer(3 * time.Second).C:
		t.Error("new task timeout")
	}
}

func TestAdjustIntervalEthash(t *testing.T) {
	t.Parallel()
	testAdjustInterval(t, ethashChainConfig, ethash.NewFaker())
}

func TestAdjustIntervalClique(t *testing.T) {
	t.Parallel()
	testAdjustInterval(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func testAdjustInterval(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, _ := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	var (
		progress = make(chan struct{}, 10)
		result   = make([]float64, 0, 10)
		index    = 0
		start    atomic.Bool
	)
	w.resubmitHook = func(minInterval time.Duration, recommitInterval time.Duration) {
		// Short circuit if interval checking hasn't started.
		if !start.Load() {
			return
		}
		var wantMinInterval, wantRecommitInterval time.Duration

		switch index {
		case 0:
			wantMinInterval, wantRecommitInterval = 3*time.Second, 3*time.Second
		case 1:
			origin := float64(3 * time.Second.Nanoseconds())
			estimate := origin*(1-intervalAdjustRatio) + intervalAdjustRatio*(origin/0.8+intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 2:
			estimate := result[index-1]
			min := float64(3 * time.Second.Nanoseconds())
			estimate = estimate*(1-intervalAdjustRatio) + intervalAdjustRatio*(min-intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 3:
			// lower than upstream test, since enforced min recommit interval is lower
			wantMinInterval, wantRecommitInterval = 500*time.Millisecond, 500*time.Millisecond
		}

		// Check interval
		if minInterval != wantMinInterval {
			t.Errorf("resubmit min interval mismatch: have %v, want %v ", minInterval, wantMinInterval)
		}
		if recommitInterval != wantRecommitInterval {
			t.Errorf("resubmit interval mismatch: have %v, want %v", recommitInterval, wantRecommitInterval)
		}
		result = append(result, float64(recommitInterval.Nanoseconds()))
		index += 1
		progress <- struct{}{}
	}
	w.start()

	time.Sleep(time.Second) // Ensure two tasks have been submitted due to start opt
	start.Store(true)

	w.setRecommitInterval(3 * time.Second)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: true, ratio: 0.8}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: false}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.setRecommitInterval(500 * time.Millisecond)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}
}

func TestGetSealingWorkEthash(t *testing.T) {
	t.Parallel()
	testGetSealingWork(t, ethashChainConfig, ethash.NewFaker())
}

func TestGetSealingWorkClique(t *testing.T) {
	t.Parallel()
	testGetSealingWork(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func TestGetSealingWorkPostMerge(t *testing.T) {
	t.Parallel()
	local := new(params.ChainConfig)
	*local = *ethashChainConfig
	local.TerminalTotalDifficulty = big.NewInt(0)
	testGetSealingWork(t, local, ethash.NewFaker())
}

func testGetSealingWork(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	w.setExtra([]byte{0x01, 0x02})

	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	timestamp := uint64(time.Now().Unix())
	assertBlock := func(block *types.Block, number uint64, coinbase common.Address, random common.Hash) {
		if block.Time() != timestamp {
			// Sometime the timestamp will be mutated if the timestamp
			// is even smaller than parent block's. It's OK.
			t.Logf("Invalid timestamp, want %d, get %d", timestamp, block.Time())
		}
		_, isClique := engine.(*clique.Clique)
		if !isClique {
			if len(block.Extra()) != 2 {
				t.Error("Unexpected extra field")
			}
			if block.Coinbase() != coinbase {
				t.Errorf("Unexpected coinbase got %x want %x", block.Coinbase(), coinbase)
			}
		} else {
			if block.Coinbase() != (common.Address{}) {
				t.Error("Unexpected coinbase")
			}
		}
		if !isClique {
			if block.MixDigest() != random {
				t.Error("Unexpected mix digest")
			}
		}
		if block.Nonce() != 0 {
			t.Error("Unexpected block nonce")
		}
		if block.NumberU64() != number {
			t.Errorf("Mismatched block number, want %d got %d", number, block.NumberU64())
		}
	}
	var cases = []struct {
		parent       common.Hash
		coinbase     common.Address
		random       common.Hash
		expectNumber uint64
		expectErr    bool
	}{
		{
			b.chain.Genesis().Hash(),
			common.HexToAddress("0xdeadbeef"),
			common.HexToHash("0xcafebabe"),
			uint64(1),
			false,
		},
		{
			b.chain.CurrentBlock().Hash(),
			common.HexToAddress("0xdeadbeef"),
			common.HexToHash("0xcafebabe"),
			b.chain.CurrentBlock().Number.Uint64() + 1,
			false,
		},
		{
			b.chain.CurrentBlock().Hash(),
			common.Address{},
			common.HexToHash("0xcafebabe"),
			b.chain.CurrentBlock().Number.Uint64() + 1,
			false,
		},
		{
			b.chain.CurrentBlock().Hash(),
			common.Address{},
			common.Hash{},
			b.chain.CurrentBlock().Number.Uint64() + 1,
			false,
		},
		{
			common.HexToHash("0xdeadbeef"),
			common.HexToAddress("0xdeadbeef"),
			common.HexToHash("0xcafebabe"),
			0,
			true,
		},
	}

	// This API should work even when the automatic sealing is not enabled
	for _, c := range cases {
		r := w.getSealingBlock(&generateParams{
			parentHash:  c.parent,
			timestamp:   timestamp,
			coinbase:    c.coinbase,
			random:      c.random,
			withdrawals: nil,
			beaconRoot:  nil,
			noTxs:       false,
			forceTime:   true,
		})
		if c.expectErr {
			if r.err == nil {
				t.Error("Expect error but get nil")
			}
		} else {
			if r.err != nil {
				t.Errorf("Unexpected error %v", r.err)
			}
			assertBlock(r.block, c.expectNumber, c.coinbase, c.random)
		}
	}

	// This API should work even when the automatic sealing is enabled
	w.start()
	for _, c := range cases {
		r := w.getSealingBlock(&generateParams{
			parentHash:  c.parent,
			timestamp:   timestamp,
			coinbase:    c.coinbase,
			random:      c.random,
			withdrawals: nil,
			beaconRoot:  nil,
			noTxs:       false,
			forceTime:   true,
		})
		if c.expectErr {
			if r.err == nil {
				t.Error("Expect error but get nil")
			}
		} else {
			if r.err != nil {
				t.Errorf("Unexpected error %v", r.err)
			}
			assertBlock(r.block, c.expectNumber, c.coinbase, c.random)
		}
	}
}
