package asset

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	CodeSpaceAsset = "asset"

	CodeInvalidTokenName   = 0
	CodeInvalidTokenSymbol = 1
	CodeInvalidTokenSupply = 2
	CodeInvalidTokenOwner  = 3
	CodeNoTokenPersist     = 4
)

func ErrorInvalidTokenName(codespace sdk.CodespaceType, fmt string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTokenName, fmt)
}
func ErrorInvalidTokenSymbol(codespace sdk.CodespaceType, fmt string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTokenSymbol, fmt)
}
func ErrorInvalidTokenSupply(codespace sdk.CodespaceType, fmt string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTokenSupply, fmt)
}
func ErrorInvalidTokenOwner(codespace sdk.CodespaceType, fmt string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTokenOwner, fmt)
}
func ErrorNoTokenPersist(codespace sdk.CodespaceType, fmt string) sdk.Error {
	return sdk.NewError(codespace, CodeNoTokenPersist, fmt)
}
