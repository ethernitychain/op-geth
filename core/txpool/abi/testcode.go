// example of transaction input data
txInput := "0xa5142faa00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003"

// load contract ABI
abi, err := abi.JSON(strings.NewReader(myContractAbi))
if err != nil {
	log.Fatal(err)
}

// decode txInput method signature
decodedSig, err := hex.DecodeString(txInput[2:10])
if err != nil {
	log.Fatal(err)
}

// recover Method from signature and ABI
method, err := abi.MethodById(decodedSig)
if err != nil {
	log.Fatal(err)
}

// decode txInput Payload
decodedData, err := hex.DecodeString(txInput[10:])
if err != nil {
	log.Fatal(err)
}

// create strut that matches input names to unpack
// for example my function takes 2 inputs, with names "Name1" and "Name2" and of type uint256 (solidity)
type FunctionInputs struct {
	Name1 *big.Int // *big.Int for uint256 for example
	Name2 *big.Int
}

var data FunctionInputs

// unpack method inputs
err = method.Inputs.Unpack(&data, decodedData)
if err != nil {
	log.Fatal(err)
}

fmt.Println(data)