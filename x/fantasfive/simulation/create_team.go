package simulation

import (
	"math/rand"

	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateTeam(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateTeam{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateTeam simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateTeam simulation not implemented"), nil, nil
	}
}
