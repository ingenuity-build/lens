package cmd

import (
	"github.com/spf13/cobra"
)

// TxCommand registers a new tx command.
func txCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "broadcast transactions to a chain",
	}

	cmd.AddCommand(
		authTxCmd(),
		authzTxCmd(),
		bankTxCmd(),
		distributionTxCmd(),
		feegrantTxCmd(),
		govTxCmd(),
		stakingTxCmd(),
		slashingTxCmd(),
	)

	return cmd
}

// authCmd returns the transaction commands for this module
func authTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auth",
		Aliases: []string{"a"},
		Short:   "auth transaction commands",
	}

	// TODO: do we need this? I don't know if we do.
	// cmd.AddCommand(authSignTxCmd())

	return cmd
}

// authzTxCmd returns the authz tx commands for this module
func authzTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "authz",
		Aliases: []string{"az"},
		Short:   "authz transaction commands",
	}

	cmd.AddCommand(
		authzGrantAuthorizationCmd(),
		authzRevokeAuthorizationCmd(),
		authzExecAuthorizationCmd(),
	)

	return cmd
}

// bankTxCmd returns the bank tx commands for this module
func bankTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bank",
		Aliases: []string{"b", "bnk"},
		Short:   "bank transaction commands",
	}

	cmd.AddCommand(bankSendCmd())

	return cmd
}

// distributionTxCmd returns the distribution tx commands for this module
func distributionTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "distribution",
		Aliases: []string{"dist", "distr", "d"},
		Short:   "distribution transaction commands",
	}

	cmd.AddCommand(
		distributionWithdrawRewardsCmd(),
		// distributionSetWithdrawAddressCmd(),
		// distributionFundCommunityPoolCmd(),
	)

	return cmd
}

// feegrantTxCmd returns the fee grant tx commands for this module
func feegrantTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "feegrant",
		Aliases: []string{"f", "fee"},
		Short:   "fee grant transaction commands",
	}

	cmd.AddCommand(
	// feegrantFeeGrantCmd(),
	// feegrantRevoteFeeGrantCmd(),
	)

	return cmd
}

// stakingTxCmd returns the staking tx commands for this module
func stakingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "staking",
		Aliases: []string{"stake", "stk"},
		Short:   "staking transaction commands",
	}

	cmd.AddCommand(
		stakingDelegateCmd(),
		stakingRedelegateCmd(),
		// stakingCreateValidatorCmd(),
		// stakingEditValidatorCmd(),
		// stakingUnbondCmd(),
	)

	return cmd
}

// govTxCmd returns the gov tx commands for this module
func govTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "governance",
		Aliases: []string{"gov", "g"},
		Short:   "governance transaction commands",
	}

	cmd.AddCommand(
	// govSubmitProposalCmd(),
	// govDepositCmd(),
	// govVoteCmd(),
	// govWeightedVoteCmd(),
	)

	return cmd
}

// slashingTxCmd returns the slashing tx commands for this module
func slashingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "slashing",
		Aliases: []string{"sl", "slash"},
		Short:   "slashing transaction commands",
	}

	cmd.AddCommand(
	// slashingUnjailCmd(),
	)

	return cmd
}
