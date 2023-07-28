package keeper

import (
	"context"

	"belfrics/x/belfrics/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKyc3(goCtx context.Context, msg *types.MsgCreateKyc3) (*types.MsgCreateKyc3Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKyc3(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kyc3 = types.Kyc3{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetKyc3(
		ctx,
		kyc3,
	)
	return &types.MsgCreateKyc3Response{}, nil
}

func (k msgServer) UpdateKyc3(goCtx context.Context, msg *types.MsgUpdateKyc3) (*types.MsgUpdateKyc3Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc3(
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

	var kyc3 = types.Kyc3{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetKyc3(ctx, kyc3)

	return &types.MsgUpdateKyc3Response{}, nil
}

func (k msgServer) DeleteKyc3(goCtx context.Context, msg *types.MsgDeleteKyc3) (*types.MsgDeleteKyc3Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc3(
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

	k.RemoveKyc3(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteKyc3Response{}, nil
}
