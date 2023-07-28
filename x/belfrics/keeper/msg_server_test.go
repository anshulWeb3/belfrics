package keeper_test

import (
	"context"
	"testing"

	keepertest "belfrics/testutil/keeper"
	"belfrics/x/belfrics/keeper"
	"belfrics/x/belfrics/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BelfricsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
