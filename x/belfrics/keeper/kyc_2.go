package keeper

import (
	"belfrics/x/belfrics/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKyc2 set a specific kyc2 in the store from its index
func (k Keeper) SetKyc2(ctx sdk.Context, kyc2 types.Kyc2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc2KeyPrefix))
	b := k.cdc.MustMarshal(&kyc2)
	store.Set(types.Kyc2Key(
		kyc2.Address,
	), b)
}

// GetKyc2 returns a kyc2 from its index
func (k Keeper) GetKyc2(
	ctx sdk.Context,
	address string,

) (val types.Kyc2, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc2KeyPrefix))

	b := store.Get(types.Kyc2Key(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKyc2 removes a kyc2 from the store
func (k Keeper) RemoveKyc2(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc2KeyPrefix))
	store.Delete(types.Kyc2Key(
		address,
	))
}

// GetAllKyc2 returns all kyc2
func (k Keeper) GetAllKyc2(ctx sdk.Context) (list []types.Kyc2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Kyc2KeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Kyc2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
