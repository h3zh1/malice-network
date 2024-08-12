package rpc

import (
	"fmt"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/proto/services/listenerrpc"
	"github.com/chainreactors/malice-network/server/internal/core"
)

func (rpc *Server) JobStream(stream listenerrpc.ListenerRPC_JobStreamServer) error {
	go func() {
		for {
			select {
			case msg := <-core.Jobs.Ctrl:
				err := stream.Send(msg)
				if err != nil {
					return
				}
			}
		}
	}()

	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		if msg.Status == consts.CtrlStatusSuccess {
			core.EventBroker.Publish(core.Event{
				Job:       core.Jobs.Get(msg.Job.Id),
				EventType: consts.EventPipeline,
				Message:   "pipeline start",
			})
		} else {
			core.EventBroker.Publish(core.Event{
				EventType: consts.EventPipeline,
				Err:       fmt.Sprintf("%d, %s", msg.Status, msg.Error),
			})
		}
		// todo stop pipeline
	}
}
