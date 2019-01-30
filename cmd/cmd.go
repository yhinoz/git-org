package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhinoz/git-org/cmd/grep"
	"github.com/yhinoz/git-org/cmd/repos"
	cmdutil "github.com/yhinoz/git-org/cmd/util"
	"github.com/yhinoz/git-org/cmd/version"
)

func NewDefaultCmd() *cobra.Command {
	f := cmdutil.NewFactory()

	cmds := &cobra.Command{
		Use:   "git-org",
		Short: "operation to the github orgranization, git-org is git subcommand",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmds.PersistentFlags().StringP("org", "o", "", "organization name")

	cmds.AddCommand(repos.NewReposCmd(f))
	cmds.AddCommand(grep.NewGrepCmd(f))
	cmds.AddCommand(version.NewVersionCmd())

	return cmds
}
