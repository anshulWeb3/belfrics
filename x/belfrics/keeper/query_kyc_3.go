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

func (k Keeper) Kyc3All(goCtx context.Context, req *types.QueryAllKyc3Request) (*types.QueryAllKyc3Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kyc3s []types.Kyc3
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	kyc3Store := prefix.NewStore(store, types.KeyPrefix(types.Kyc3KeyPrefix))

	pageRes, err := query.Paginate(kyc3Store, req.Pagination, func(key []byte, value []byte) error {
		var kyc3 types.Kyc3
		if err := k.cdc.Unmarshal(value, &kyc3); err != nil {
			return err
		}

		kyc3s = append(kyc3s, kyc3)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKyc3Response{Kyc3: kyc3s, Pagination: pageRes}, nil
}

func (k Keeper) Kyc3(goCtx context.Context, req *types.QueryGetKyc3Request) (*types.QueryGetKyc3Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKyc3(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKyc3Response{Kyc3: val}, nil
}
