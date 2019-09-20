package cmd

import "github.com/chaosblade-io/chaosblade/version"

func CmdInit(ver, env, buildTime string) *baseCommand {
	//version init
	version.InitVersion(ver, env, buildTime)

	cli := NewCli()
	baseCmd := &baseCommand{
		command: cli.rootCmd,
	}
	// add version command
	baseCmd.AddCommand(&VersionCommand{})
	// add prepare command
	prepareCommand := &PrepareCommand{}
	baseCmd.AddCommand(prepareCommand)
	prepareCommand.AddCommand(&PrepareJvmCommand{})
	prepareCommand.AddCommand(&PrepareCPlusCommand{})

	// add revoke command
	baseCmd.AddCommand(&RevokeCommand{})

	// add create command
	createCommand := &CreateCommand{}
	baseCmd.AddCommand(createCommand)

	// add exp command
	expCommand := NewExpCommand()
	expCommand.AddCommandTo(createCommand)

	// add destroy command
	baseCmd.AddCommand(&DestroyCommand{exp: expCommand})

	// add status command
	baseCmd.AddCommand(&StatusCommand{exp: expCommand})

	// add query command
	queryCommand := &QueryCommand{}
	baseCmd.AddCommand(queryCommand)
	queryCommand.AddCommand(&QueryDiskCommand{})
	queryCommand.AddCommand(&QueryNetworkCommand{})
	queryCommand.AddCommand(&QueryJvmCommand{})

	return baseCmd
}
