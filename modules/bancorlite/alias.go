package bancorlite

import (
	"github.com/coinexchain/dex/modules/bancorlite/internal/keepers"
	"github.com/coinexchain/dex/modules/bancorlite/internal/types"
)

const (
	StoreKey   = types.StoreKey
	ModuleName = types.ModuleName
)

var (
	NewBaseKeeper       = keepers.NewKeeper
	NewBancorInfoKeeper = keepers.NewBancorInfoKeeper
	DefaultParams       = types.DefaultParams
)

type (
	Keeper                     = keepers.Keeper
	MsgBancorTradeInfoForKafka = types.MsgBancorTradeInfoForKafka
	MsgBancorInfoForKafka      = types.MsgBancorInfoForKafka
)
