package simulation

import (
	"math/rand"

	"github.com/PriceChain/cprc/x/registry/keeper"
	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUnbondRegistry(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUnbondRegistry{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UnbondRegistry simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UnbondRegistry simulation not implemented"), nil, nil
	}
}
