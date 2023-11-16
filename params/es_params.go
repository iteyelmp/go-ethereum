package params

const (
	EsGetBlobPerByteGas uint64 = 1     // Per-byte price for an esGetBlob operation.
	EsGetBlobBaseGas    uint64 = 50000 // Base price for an esGetBlob operation.
	ESMemoryGas         uint64 = 1
	ESRPCGasCap         uint64 = 100000000
)

var (
	EsNodeURL = "http://localhost:9545"
)
