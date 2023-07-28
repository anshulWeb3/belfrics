package keeper

import (
	"context"

	"belfrics/x/belfrics/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Kyc2All(goCtx context.Context, req *types.QueryAllKyc2Request) (*types.QueryAllKyc2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kyc2s []types.Kyc2
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	kyc2Store := prefix.NewStore(store, types.KeyPrefix(types.Kyc2KeyPrefix))

	pageRes, err := query.Paginate(kyc2Store, req.Pagination, func(key []byte, value []byte) error {
		var kyc2 types.Kyc2
		if err := k.cdc.Unmarshal(value, &kyc2); err != nil {
			return err
		}

		kyc2s = append(kyc2s, kyc2)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKyc2Response{Kyc2: kyc2s, Pagination: pageRes}, nil
}

func (k Keeper) Kyc2(goCtx context.Context, req *types.QueryGetKyc2Request) (*types.QueryGetKyc2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKyc2(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKyc2Response{Kyc2: val}, nil
}
