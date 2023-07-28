package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKyc3 = "create_kyc_3"
	TypeMsgUpdateKyc3 = "update_kyc_3"
	TypeMsgDeleteKyc3 = "delete_kyc_3"
)

var _ sdk.Msg = &MsgCreateKyc3{}

func NewMsgCreateKyc3(
	creator string,
	address string,
	value bool,

) *MsgCreateKyc3 {
	return &MsgCreateKyc3{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgCreateKyc3) Route() string {
	return RouterKey
}

func (msg *MsgCreateKyc3) Type() string {
	return TypeMsgCreateKyc3
}

func (msg *MsgCreateKyc3) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKyc3) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKyc3) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKyc3{}

func NewMsgUpdateKyc3(
	creator string,
	address string,
	value bool,

) *MsgUpdateKyc3 {
	return &MsgUpdateKyc3{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgUpdateKyc3) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKyc3) Type() string {
	return TypeMsgUpdateKyc3
}

func (msg *MsgUpdateKyc3) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKyc3) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKyc3) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKyc3{}

func NewMsgDeleteKyc3(
	creator string,
	address string,

) *MsgDeleteKyc3 {
	return &MsgDeleteKyc3{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteKyc3) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKyc3) Type() string {
	return TypeMsgDeleteKyc3
}

func (msg *MsgDeleteKyc3) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKyc3) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKyc3) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
