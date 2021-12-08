package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateAdvertisement{}, "halborn/CreateAdvertisement", nil)
	cdc.RegisterConcrete(MsgDeleteAdvertisement{}, "halborn/DeleteAdvertisement", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "halborn/deposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "halborn/withdraw", nil)
	cdc.RegisterConcrete(MsgCaptureTheFlag{}, "halborn/CaptureTheFlag", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
