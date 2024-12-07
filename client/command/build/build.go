package build

import (
	"errors"
	"fmt"
	"github.com/chainreactors/malice-network/client/command/common"
	"github.com/chainreactors/malice-network/client/repl"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	"github.com/chainreactors/malice-network/helper/types"
	"github.com/spf13/cobra"
	"os"
)

func BeaconCmd(cmd *cobra.Command, con *repl.Console) error {
	name, address, buildTarget, modules, ca, interval, jitter, srdi := common.ParseGenerateFlags(cmd)
	if buildTarget == "" {
		return errors.New("require build target")
	}
	go func() {
		params := &types.ProfileParams{
			Interval: interval,
			Jitter:   jitter,
		}
		_, err := con.Rpc.Build(con.Context(), &clientpb.Generate{
			ProfileName: name,
			Address:     address,
			Type:        consts.CommandBuildBeacon,
			Target:      buildTarget,
			Modules:     modules,
			Ca:          ca,
			Params:      params.String(),
			Srdi:        srdi,
		})
		if err != nil {
			con.Log.Errorf("Build beacon failed: %v", err)
			return
		}
	}()
	return nil
}

func BindCmd(cmd *cobra.Command, con *repl.Console) error {
	name, address, buildTarget, modules, ca, interval, jitter, srdi := common.ParseGenerateFlags(cmd)
	if buildTarget == "" {
		return errors.New("require build target")
	}
	go func() {
		params := &types.ProfileParams{
			Interval: interval,
			Jitter:   jitter,
		}
		_, err := con.Rpc.Build(con.Context(), &clientpb.Generate{
			ProfileName: name,
			Address:     address,
			Type:        consts.CommandBuildBind,
			Target:      buildTarget,
			Modules:     modules,
			Ca:          ca,
			Params:      params.String(),
			Srdi:        srdi,
		})
		if err != nil {
			con.Log.Errorf("Build bind failed: %v", err)
			return
		}
	}()
	return nil
}

func PreludeCmd(cmd *cobra.Command, con *repl.Console) error {
	name, address, buildTarget, modules, ca, _, _, srdi := common.ParseGenerateFlags(cmd)
	if buildTarget == "" {
		return errors.New("require build target")
	}
	autorunPath, _ := cmd.Flags().GetString("autorun")
	if autorunPath == "" {
		return errors.New("require autorun.yaml path")
	}
	file, err := os.ReadFile(autorunPath)
	if err != nil {
		return err
	}
	go func() {
		_, err := con.Rpc.Build(con.Context(), &clientpb.Generate{
			ProfileName: name,
			Address:     address,
			Type:        consts.CommandBuildPrelude,
			Target:      buildTarget,
			Modules:     modules,
			Ca:          ca,
			Srdi:        srdi,
			Bin:         file,
		})
		if err != nil {
			con.Log.Errorf("Build prelude failed: %v\n", err)
			return
		}
	}()
	return nil
}

func ModulesCmd(cmd *cobra.Command, con *repl.Console) error {
	name, address, buildTarget, modules, _, _, _, srdi := common.ParseGenerateFlags(cmd)
	if len(modules) == 0 {
		modules = []string{"full"}
	}
	if buildTarget == "" {
		return errors.New("require build target")
	}
	go func() {
		_, err := BuildModules(con, name, address, buildTarget, modules, srdi)
		if err != nil {
			con.Log.Errorf("Build modules failed: %v", err)
			return
		}
	}()
	return nil
}

func PulseCmd(cmd *cobra.Command, con *repl.Console) error {
	profile, _ := cmd.Flags().GetString("profile")
	address, _ := cmd.Flags().GetString("address")
	buildTarget, _ := cmd.Flags().GetString("target")
	srdi, _ := cmd.Flags().GetBool("srdi")
	artifactId, _ := cmd.Flags().GetUint32("artifact-id")
	if buildTarget != consts.TargetX64Windows && buildTarget != consts.TargetX86Windows {
		return errors.New("pulse build target must be x86_64-pc-windows-msvc or i686-pc-windows-msvc")
	}
	go func() {
		_, err := con.Rpc.Build(con.Context(), &clientpb.Generate{
			ProfileName: profile,
			Address:     address,
			Target:      buildTarget,
			Type:        consts.CommandBuildPulse,
			Srdi:        srdi,
			ArtifactId:  artifactId,
		})
		if err != nil {
			con.Log.Errorf("Build loader failed: %v", err)
			return
		}
	}()
	return nil
}

func BuildLogCmd(cmd *cobra.Command, con *repl.Console) error {
	name := cmd.Flags().Arg(0)
	num, _ := cmd.Flags().GetInt("limit")
	builder, err := con.Rpc.BuildLog(con.Context(), &clientpb.Builder{
		Name: name,
		Num:  uint32(num),
	})
	if err != nil {
		return err
	}
	if len(builder.Log) == 0 {
		con.Log.Infof("No log for %s", name)
		return nil
	}
	fmt.Println(string(builder.Log))
	return nil
}

func BuildModules(con *repl.Console, name, address, buildTarget string, modules []string, srdi bool) (bool, error) {
	_, err := con.Rpc.Build(con.Context(), &clientpb.Generate{
		ProfileName: name,
		Address:     address,
		Target:      buildTarget,
		Type:        consts.CommandBuildModules,
		Modules:     modules,
		Srdi:        srdi,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
