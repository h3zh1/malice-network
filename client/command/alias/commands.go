package alias

import (
	"github.com/chainreactors/malice-network/client/command/help"
	"github.com/chainreactors/malice-network/client/console"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func Commands(con *console.Console) []*cobra.Command {
	aliasCmd := &cobra.Command{
		Use:   consts.CommandAlias,
		Short: "List current aliases",
		Long:  help.GetHelpFor(consts.CommandAlias),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			return
		},
	}

	aliasLoadCmd := &cobra.Command{
		Use:   consts.CommandAliasLoad,
		Short: "Load a command alias",
		Long:  help.GetHelpFor(consts.CommandAlias + " " + consts.CommandAliasLoad),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			AliasesLoadCmd(cmd, con)
			return
		},
	}
	carapace.Gen(aliasLoadCmd).PositionalCompletion(
		carapace.ActionFiles().Usage("local path where the downloaded file will be saved (optional)"),
	)

	aliasInstallCmd := &cobra.Command{
		Use:   consts.CommandAliasInstall,
		Short: "Install a command alias",
		Long:  help.GetHelpFor(consts.CommandAlias + " " + consts.CommandAliasInstall),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			AliasesInstallCmd(cmd, con)
			return
		},
	}
	carapace.Gen(aliasInstallCmd).PositionalCompletion(
		carapace.ActionFiles().Usage("local path where the downloaded file will be saved (optional)"),
	)

	aliasRemoveCmd := &cobra.Command{
		Use:   consts.CommandAliasRemove,
		Short: "Remove an alias",
		Long:  help.GetHelpFor(consts.CommandAlias + " " + consts.CommandAliasRemove),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			AliasesRemoveCmd(cmd, con)
			return
		},
	}
	carapace.Gen(aliasRemoveCmd).PositionalCompletion(
		AliasCompleter(),
	)

	aliasCmd.AddCommand(aliasLoadCmd, aliasInstallCmd, aliasRemoveCmd)
	return []*cobra.Command{aliasCmd}

}
