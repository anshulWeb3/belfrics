package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKyc2 = "create_kyc_2"
	TypeMsgUpdateKyc2 = "update_kyc_2"
	TypeMsgDeleteKyc2 = "delete_kyc_2"
)

var _ sdk.Msg = &MsgCreateKyc2{}

func NewMsgCreateKyc2(
	creator string,
	address string,

) *MsgCreateKyc2 {
	return &MsgCreateKyc2{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateKyc2) Route() string {
	return RouterKey
}

func (msg *MsgCreateKyc2) Type() string {
	return TypeMsgCreateKyc2
}

func (msg *MsgCreateKyc2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKyc2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKyc2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKyc2{}

func NewMsgUpdateKyc2(
	creator string,
	address string,

) *MsgUpdateKyc2 {
	return &MsgUpdateKyc2{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateKyc2) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKyc2) Type() string {
	return TypeMsgUpdateKyc2
}

func (msg *MsgUpdateKyc2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKyc2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKyc2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKyc2{}

func NewMsgDeleteKyc2(
	creator string,
	address string,

) *MsgDeleteKyc2 {
	return &MsgDeleteKyc2{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteKyc2) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKyc2) Type() string {
	return TypeMsgDeleteKyc2
}

func (msg *MsgDeleteKyc2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKyc2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKyc2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
