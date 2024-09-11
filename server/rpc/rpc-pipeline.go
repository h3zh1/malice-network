package rpc

import (
	"context"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/proto/client/clientpb"
	"github.com/chainreactors/malice-network/proto/implant/implantpb"
	"github.com/chainreactors/malice-network/server/internal/certs"
	"github.com/chainreactors/malice-network/server/internal/configs"
	"github.com/chainreactors/malice-network/server/internal/db"
	"github.com/chainreactors/malice-network/server/internal/db/models"

	"github.com/chainreactors/malice-network/proto/listener/lispb"
	"github.com/chainreactors/malice-network/server/internal/core"
)

func (rpc *Server) RegisterPipeline(ctx context.Context, req *lispb.Pipeline) (*implantpb.Empty, error) {
	var name string
	if req.GetTls().Key == "" || req.GetTls().Cert == "" {
		switch req.Body.(type) {
		case *lispb.Pipeline_Tcp:
			name = req.GetTcp().Name
		}
		tlsConfig := configs.GenerateTlsConfig(name)
		cert, key, err := certs.GeneratePipelineCert(&tlsConfig)
		if err != nil {
			return &implantpb.Empty{}, err
		}
		req.Tls.Cert = string(cert)
		req.Tls.Key = string(key)
	}
	err := db.CreatePipeline(req)
	if err != nil {
		return &implantpb.Empty{}, err
	}
	return &implantpb.Empty{}, nil
}

func (rpc *Server) NewPipeline(ctx context.Context, req *lispb.Pipeline) (*clientpb.Empty, error) {
	err := db.CreatePipeline(req)
	if err != nil {
		return &clientpb.Empty{}, err
	}
	return &clientpb.Empty{}, nil
}

func (rpc *Server) ListTcpPipelines(ctx context.Context, req *lispb.ListenerName) (*lispb.Pipelines, error) {
	var result []*lispb.Pipeline
	pipelines, err := db.ListPipelines(req.Name, "tcp")
	if err != nil {
		return &lispb.Pipelines{}, err
	}
	for _, pipeline := range pipelines {
		tcp := lispb.TCPPipeline{
			Name:   pipeline.Name,
			Host:   pipeline.Host,
			Port:   uint32(pipeline.Port),
			Enable: pipeline.Enable,
		}
		result = append(result, &lispb.Pipeline{
			Body: &lispb.Pipeline_Tcp{
				Tcp: &tcp,
			},
		})
	}
	return &lispb.Pipelines{Pipelines: result}, nil
}

func (rpc *Server) StartTcpPipeline(ctx context.Context, req *lispb.CtrlPipeline) (*clientpb.Empty, error) {
	pipelineDB, err := db.FindPipeline(req.Name, req.ListenerId)
	if err != nil {
		return &clientpb.Empty{}, err
	}
	pipeline := models.ToProtobuf(&pipelineDB)
	ch := make(chan bool)
	job := &core.Job{
		ID:      core.CurrentJobID(),
		Message: pipeline,
		JobCtrl: ch,
		Name:    pipeline.GetTcp().Name,
	}
	core.Jobs.Add(job)
	ctrl := clientpb.JobCtrl{
		Id:   core.NextCtrlID(),
		Ctrl: consts.CtrlPipelineStart,
		Job: &clientpb.Job{
			Id:       core.NextJobID(),
			Pipeline: pipeline,
		},
	}
	core.Jobs.Ctrl <- &ctrl
	return &clientpb.Empty{}, nil
}

func (rpc *Server) StopTcpPipeline(ctx context.Context, req *lispb.CtrlPipeline) (*clientpb.Empty, error) {
	pipelineDB, err := db.FindPipeline(req.Name, req.ListenerId)
	if err != nil {
		return &clientpb.Empty{}, err
	}
	pipeline := models.ToProtobuf(&pipelineDB)
	ctrl := clientpb.JobCtrl{
		Id:   core.NextCtrlID(),
		Ctrl: consts.CtrlPipelineStop,
		Job: &clientpb.Job{
			Id:       core.NextJobID(),
			Pipeline: pipeline,
		},
	}
	core.Jobs.Ctrl <- &ctrl
	return &clientpb.Empty{}, nil
}

func (rpc *Server) ListJobs(ctx context.Context, req *clientpb.Empty) (*lispb.Pipelines, error) {
	var pipelines []*lispb.Pipeline
	for _, job := range core.Jobs.All() {
		pipeline, ok := job.Message.(*lispb.Pipeline)
		if !ok {
			continue
		}
		if pipeline.GetTcp() != nil {
			pipelines = append(pipelines, job.Message.(*lispb.Pipeline))
		}
	}
	return &lispb.Pipelines{Pipelines: pipelines}, nil
}
