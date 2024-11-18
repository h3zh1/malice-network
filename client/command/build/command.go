package build

import (
	"fmt"
	"github.com/chainreactors/malice-network/client/command/common"
	"github.com/chainreactors/malice-network/client/repl"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	"github.com/kballard/go-shellquote"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

func Commands(con *repl.Console) []*cobra.Command {
	profileCmd := &cobra.Command{
		Use:   consts.CommandProfile,
		Short: "Show all compile profile ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ProfileShowCmd(cmd, con)
		},
	}

	newCmd := &cobra.Command{
		Use:   "new",
		Short: "New compile profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ProfileNewCmd(cmd, con)
		},
		Example: `~~~`,
	}
	common.BindFlag(newCmd, common.ProfileSet)
	common.BindFlagCompletions(newCmd, func(comp carapace.ActionMap) {
		comp["name"] = carapace.ActionValues("profile name")
		comp["target"] = common.BuildTargetCompleter(con)
		comp["pipeline_id"] = common.AllPipelineCompleter(con)
		comp["proxy"] = carapace.ActionValues("http", "socks5")
		//comp["obfuscate"] = carapace.ActionValues("true", "false")
		comp["modules"] = carapace.ActionValues("e.g.: execute_exe,execute_dll")
		comp["ca"] = carapace.ActionValues("true", "false")

		comp["interval"] = carapace.ActionValues("5")
		comp["jitter"] = carapace.ActionValues("0.2")
	})

	profileCmd.AddCommand(newCmd)

	buildCmd := &cobra.Command{
		Use:   consts.CommandBuild,
		Short: "build",
	}
	// build beacon --format/-f exe,dll,shellcode -i 1.1.1 -m load_pe
	beaconCmd := &cobra.Command{
		Use:   consts.CommandBuildBeacon,
		Short: "build beacon",
		RunE: func(cmd *cobra.Command, args []string) error {
			return BeaconCmd(cmd, con)
		},
	}
	common.BindFlag(beaconCmd, common.GenerateFlagSet)
	common.BindFlagCompletions(beaconCmd, func(comp carapace.ActionMap) {
		comp["profile_name"] = common.ProfileCompleter(con)
		comp["target"] = common.BuildTargetCompleter(con)
	})
	common.BindArgCompletions(beaconCmd, nil, common.ProfileCompleter(con))

	bindCmd := &cobra.Command{
		Use:   consts.CommandBuildBind,
		Short: "build bind",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return BindCmd(cmd, con)
		},
	}

	common.BindFlag(bindCmd, common.GenerateFlagSet)
	common.BindFlagCompletions(bindCmd, func(comp carapace.ActionMap) {
		comp["profile_name"] = common.ProfileCompleter(con)
		comp["target"] = common.BuildTargetCompleter(con)
	})
	common.BindArgCompletions(bindCmd, nil, common.ProfileCompleter(con))

	shellCodeCmd := &cobra.Command{
		Use:   consts.CommandBuildShellCode,
		Short: "build ShellCode",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ShellCodeCmd(cmd, con)
		},
	}

	common.BindFlag(shellCodeCmd, common.GenerateFlagSet)
	common.BindFlagCompletions(shellCodeCmd, func(comp carapace.ActionMap) {
		comp["profile_name"] = common.ProfileCompleter(con)
		comp["target"] = common.BuildTargetCompleter(con)
	})
	common.BindArgCompletions(shellCodeCmd, nil, common.ProfileCompleter(con))

	preludeCmd := &cobra.Command{
		Use:   consts.CommandBuildPrelude,
		Short: "build prelude",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return PreludeCmd(cmd, con)
		},
	}

	common.BindFlag(preludeCmd, common.GenerateFlagSet)
	common.BindFlagCompletions(preludeCmd, func(comp carapace.ActionMap) {
		comp["profile_name"] = common.ProfileCompleter(con)
		comp["target"] = common.BuildTargetCompleter(con)
	})
	common.BindArgCompletions(preludeCmd, nil, common.ProfileCompleter(con))

	modulesCmd := &cobra.Command{
		Use:   consts.CommandBuildModules,
		Short: "show all modules",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ModulesCmd(cmd, con)
		},
	}

	common.BindFlag(modulesCmd, common.GenerateFlagSet, func(f *pflag.FlagSet) {
		f.StringSlice("feature", []string{}, "Set modules e.g.: nano, full")
	})

	common.BindFlagCompletions(modulesCmd, func(comp carapace.ActionMap) {
		comp["profile_name"] = common.ProfileCompleter(con)
		comp["target"] = common.BuildTargetCompleter(con)
		//comp["type"] = common.BuildFormatCompleter(con)
	})
	common.BindArgCompletions(modulesCmd, nil, common.ProfileCompleter(con))

	buildCmd.AddCommand(beaconCmd, bindCmd, shellCodeCmd, preludeCmd, modulesCmd)

	srdiCmd := &cobra.Command{
		Use:   consts.CommandSRDI,
		Short: "build srdi",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SRDICmd(cmd, con)
		},
	}
	common.BindFlag(srdiCmd, common.SRDIFlagSet)
	common.BindFlagCompletions(srdiCmd, func(comp carapace.ActionMap) {
		comp["path"] = carapace.ActionFiles().Usage("file path")
		comp["id"] = common.ArtifactCompleter(con)
	})

	artifactCmd := &cobra.Command{
		Use:   consts.CommandArtifact,
		Short: "artifact manage",
	}

	listArtifactCmd := &cobra.Command{
		Use:   consts.CommandArtifactList,
		Short: "list build output file in server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListArtifactCmd(cmd, con)
		},
	}

	downloadCmd := &cobra.Command{
		Use:   consts.CommandArtifactDownload,
		Short: "download build output file in server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return DownloadArtifactCmd(cmd, con)
		},
	}
	common.BindFlag(downloadCmd, func(f *pflag.FlagSet) {
		f.StringP("output", "o", "", "output path")
	})
	common.BindArgCompletions(downloadCmd, nil, common.ArtifactCompleter(con))

	uploadCmd := &cobra.Command{
		Use:   consts.CommandArtifactUpload,
		Short: "upload build output file in server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return UploadArtifactCmd(cmd, con)
		},
	}
	common.BindArgCompletions(uploadCmd, nil, carapace.ActionFiles().Usage("custom artifact"))
	common.BindFlag(uploadCmd, func(f *pflag.FlagSet) {
		f.StringP("stage", "s", "", "Set stage")
		f.StringP("type", "t", "", "Set type")
		f.StringP("name", "n", "", "alias name")
	})

	artifactCmd.AddCommand(listArtifactCmd, downloadCmd, uploadCmd)

	return []*cobra.Command{profileCmd, buildCmd, artifactCmd, srdiCmd}
}

func Register(con *repl.Console) {
	con.RegisterServerFunc("payload_local", func(shellcodePath string) (string, error) {
		if shellcodePath != "" {
			shellcode, _ := os.ReadFile(shellcodePath)
			if _, err := os.Stat(shellcodePath); os.IsNotExist(err) {
				return "", fmt.Errorf("shellcode file does not exist: %s", shellcodePath)
			}
			return string(shellcode), nil
		} else {
			return "shellcode123", nil
		}
	}, nil)

	con.RegisterServerFunc("donut_exe2shellcode", func(exe []byte, arch string, param string) (string, error) {
		cmdline, err := shellquote.Split(param)
		if err != nil {
			return "", err
		}

		bin, err := con.Rpc.EXE2Shellcode(con.Context(), &clientpb.EXE2Shellcode{
			Bin:    exe,
			Arch:   arch,
			Type:   "donut",
			Params: strings.Join(cmdline, ","),
		})
		if err != nil {
			return "", err
		}
		return string(bin.Bin), nil
	}, nil)

	con.RegisterServerFunc("donut_dll2shellcode", func(dll []byte, arch string, param string) (string, error) {
		cmdline, err := shellquote.Split(param)
		if err != nil {
			return "", err
		}

		bin, err := con.Rpc.DLL2Shellcode(con.Context(), &clientpb.DLL2Shellcode{
			Bin:    dll,
			Arch:   arch,
			Type:   "donut",
			Params: strings.Join(cmdline, ","),
		})
		if err != nil {
			return "", err
		}
		return string(bin.Bin), nil
	}, nil)

	con.RegisterServerFunc("srdi", func(dll []byte, entry string, arch string, param string) (string, error) {
		bin, err := con.Rpc.DLL2Shellcode(con.Context(), &clientpb.DLL2Shellcode{
			Bin:        dll,
			Arch:       arch,
			Type:       "srdi",
			Entrypoint: entry,
			Params:     param,
		})
		if err != nil {
			return "", err
		}
		return string(bin.Bin), nil
	}, nil)

	con.RegisterServerFunc("sgn_encode", func(shellcode []byte, arch string, iterations int32) (string, error) {
		bin, err := con.Rpc.ShellcodeEncode(con.Context(), &clientpb.ShellcodeEncode{
			Shellcode:  shellcode,
			Arch:       arch,
			Type:       "sgn",
			Iterations: iterations,
		})
		if err != nil {
			return "", err
		}
		return string(bin.Bin), nil
	}, nil)
}
