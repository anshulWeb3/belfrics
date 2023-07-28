package keeper

import (
	"context"

	"belfrics/x/belfrics/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKyc2(goCtx context.Context, msg *types.MsgCreateKyc2) (*types.MsgCreateKyc2Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKyc2(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kyc2 = types.Kyc2{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetKyc2(
		ctx,
		kyc2,
	)
	return &types.MsgCreateKyc2Response{}, nil
}

func (k msgServer) UpdateKyc2(goCtx context.Context, msg *types.MsgUpdateKyc2) (*types.MsgUpdateKyc2Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc2(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kyc2 = types.Kyc2{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetKyc2(ctx, kyc2)

	return &types.MsgUpdateKyc2Response{}, nil
}

func (k msgServer) DeleteKyc2(goCtx context.Context, msg *types.MsgDeleteKyc2) (*types.MsgDeleteKyc2Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc2(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKyc2(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteKyc2Response{}, nil
}
