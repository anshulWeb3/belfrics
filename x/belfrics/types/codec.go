package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateKyc{}, "belfrics/CreateKyc", nil)
	cdc.RegisterConcrete(&MsgUpdateKyc{}, "belfrics/UpdateKyc", nil)
	cdc.RegisterConcrete(&MsgDeleteKyc{}, "belfrics/DeleteKyc", nil)
	cdc.RegisterConcrete(&MsgCreateKyc2{}, "belfrics/CreateKyc2", nil)
	cdc.RegisterConcrete(&MsgUpdateKyc2{}, "belfrics/UpdateKyc2", nil)
	cdc.RegisterConcrete(&MsgDeleteKyc2{}, "belfrics/DeleteKyc2", nil)
	cdc.RegisterConcrete(&MsgCreateKyc3{}, "belfrics/CreateKyc3", nil)
	cdc.RegisterConcrete(&MsgUpdateKyc3{}, "belfrics/UpdateKyc3", nil)
	cdc.RegisterConcrete(&MsgDeleteKyc3{}, "belfrics/DeleteKyc3", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKyc{},
		&MsgUpdateKyc{},
		&MsgDeleteKyc{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKyc2{},
		&MsgUpdateKyc2{},
		&MsgDeleteKyc2{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKyc3{},
		&MsgUpdateKyc3{},
		&MsgDeleteKyc3{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
