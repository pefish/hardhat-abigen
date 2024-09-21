package type_

type ArtifactData struct {
	ContractName string      `json:"contractName"`
	ABI          interface{} `json:"abi"`
	ByteCode     string      `json:"bytecode"`
}
