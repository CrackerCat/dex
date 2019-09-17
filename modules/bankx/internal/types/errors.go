package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	CodeSpaceBankx sdk.CodespaceType = ModuleName

	// 301 ～ 399
	CodeMemoMissing                     sdk.CodeType = 301
	CodeInsufficientCETForActivationFee sdk.CodeType = 302
	CodeInvalidActivationFee            sdk.CodeType = 303
	CodeInvalidUnlockTime               sdk.CodeType = 304
	CodeTokenForbiddenByOwner           sdk.CodeType = 305
	CodeInvalidLockCoinsFee             sdk.CodeType = 306
	CodeNoInputs                        sdk.CodeType = 307
	CodeNoOutputs                       sdk.CodeType = 308
	CodeInputOutputMismatch             sdk.CodeType = 309
)

func ErrMemoMissing() sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeMemoMissing, "memo is empty")
}

func ErrorInsufficientCETForActivatingFee() sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeInsufficientCETForActivationFee, "Insufficient CET for Activating fees")
}

func ErrUnlockTime(msg string) sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeInvalidUnlockTime, msg)
}

func ErrTokenForbiddenByOwner(msg string) sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeTokenForbiddenByOwner, msg)
}
func ErrNoInputs(msg string) sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeNoInputs, msg)
}
func ErrNoOutputs(msg string) sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeNoOutputs, msg)
}
func ErrInputOutputMismatch(msg string) sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeInputOutputMismatch, msg)
}
func ErrorInvalidActivatingFee() sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeInvalidActivationFee, "invalid activated fees")
}

func ErrorInvalidLockCoinsFee() sdk.Error {
	return sdk.NewError(CodeSpaceBankx, CodeInvalidLockCoinsFee, "invalid lock coins fee")
}
