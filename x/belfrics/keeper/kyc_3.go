package keeper

import (
	"belfrics/x/belfrics/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKyc3 set a specific kyc3 in the store from its index
func (k Keeper) SetKyc3(ctx sdk.Context, kyc3 types.Kyc3) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc3KeyPrefix))
	b := k.cdc.MustMarshal(&kyc3)
	store.Set(types.Kyc3Key(
		kyc3.Address,
	), b)
}

// GetKyc3 returns a kyc3 from its index
func (k Keeper) GetKyc3(
	ctx sdk.Context,
	address string,

) (val types.Kyc3, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc3KeyPrefix))

	b := store.Get(types.Kyc3Key(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKyc3 removes a kyc3 from the store
func (k Keeper) RemoveKyc3(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc3KeyPrefix))
	store.Delete(types.Kyc3Key(
		address,
	))
}

// GetAllKyc3 returns all kyc3
func (k Keeper) GetAllKyc3(ctx sdk.Context) (list []types.Kyc3) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc3KeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Kyc3
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
