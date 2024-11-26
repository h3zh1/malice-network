package listener

import (
	"fmt"
	"github.com/chainreactors/malice-network/client/repl"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	"github.com/chainreactors/tui"
	"github.com/evertras/bubble-table/table"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"strconv"
)

func ListPipelineCmd(cmd *cobra.Command, con *repl.Console) error {
	listenerID, _ := cmd.Flags().GetString("listener")
	pipelines, err := con.Rpc.ListPipelines(con.Context(), &clientpb.Listener{
		Id: listenerID,
	})
	if err != nil {
		return err
	}
	if len(pipelines.Pipelines) == 0 {
		con.Log.Warnf("No pipelines found")
		return nil
	}
	var rowEntries []table.Row
	var row table.Row
	tableModel := tui.NewTable([]table.Column{
		table.NewColumn("Name", "Name", 20),
		table.NewColumn("Enable", "Enable", 7),
		table.NewColumn("Type", "Type", 10),
		table.NewColumn("ListenerID", "ListenerID", 15),
		table.NewColumn("Address", "Address", 20),
		table.NewColumn("Parser", "Parser", 10),
		table.NewColumn("Encryption", "Encryption", 10),
		table.NewColumn("TLS", "TLS", 10),
	}, true)
	for _, pipeline := range pipelines.GetPipelines() {
		newRow := table.RowData{}
		if pipeline.Enable {
			newRow["Enable"] = termenv.String(strconv.FormatBool(pipeline.Enable)).Foreground(tui.Green).String()
		} else {
			newRow["Enable"] = termenv.String(strconv.FormatBool(pipeline.Enable)).Foreground(tui.Red).String()
		}
		if pipeline.Tls.Enable {
			newRow["TLS"] = termenv.String(strconv.FormatBool(pipeline.Tls.Enable)).Foreground(tui.Green).String()
		} else {
			newRow["TLS"] = termenv.String(strconv.FormatBool(pipeline.Tls.Enable)).Foreground(tui.Red).String()
		}
		if pipeline.Encryption.Enable {
			newRow["Encryption"] = pipeline.Encryption.Type
		} else {
			newRow["Encryption"] = "raw"
		}
		switch body := pipeline.Body.(type) {
		case *clientpb.Pipeline_Tcp:
			newRow["Name"] = pipeline.Name
			newRow["Type"] = consts.TCPPipeline
			newRow["ListenerID"] = pipeline.ListenerId
			newRow["Address"] = body.Tcp.Host + ":" + strconv.Itoa(int(body.Tcp.Port))
			newRow["Parser"] = pipeline.Parser
			row = table.NewRow(newRow)
		case *clientpb.Pipeline_Bind:
			newRow["Name"] = pipeline.Name
			newRow["Type"] = consts.BindPipeline
			newRow["ListenerID"] = pipeline.ListenerId
			newRow["Parser"] = pipeline.Parser
			row = table.NewRow(newRow)
		}

		rowEntries = append(rowEntries, row)
	}
	tableModel.SetRows(rowEntries)
	fmt.Printf(tableModel.View())
	return nil
}

func StartPipelineCmd(cmd *cobra.Command, con *repl.Console) error {
	name := cmd.Flags().Arg(0)
	_, err := con.Rpc.StartPipeline(con.Context(), &clientpb.CtrlPipeline{
		Name: name,
	})
	if err != nil {
		return err
	}
	return nil
}

func StopPipelineCmd(cmd *cobra.Command, con *repl.Console) error {
	name := cmd.Flags().Arg(0)
	_, err := con.Rpc.StopPipeline(con.Context(), &clientpb.CtrlPipeline{
		Name: name,
	})
	if err != nil {
		return err
	}
	return nil
}
