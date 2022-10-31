package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRegistry{}, "registry/CreateRegistry", nil)
	cdc.RegisterConcrete(&MsgJoinRegistryCoOperator{}, "registry/JoinRegistryCoOperator", nil)
	cdc.RegisterConcrete(&MsgJoinRegistryMember{}, "registry/JoinRegistryMember", nil)
	cdc.RegisterConcrete(&MsgUnbondRegistry{}, "registry/UnbondRegistry", nil)
	cdc.RegisterConcrete(&MsgModifyRegistry{}, "registry/ModifyRegistry", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRegistry{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinRegistryCoOperator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinRegistryMember{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnbondRegistry{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgModifyRegistry{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
