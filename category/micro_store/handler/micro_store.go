package handler
import (
	"context"
    "github.com/coder279/micro_store/domain/service"
	log "github.com/micro/go-micro/v2/logger"
	micro_store "github.com/coder279/micro_store/proto/micro_store"
)
type Micro_store struct{
     Micro_storeDataService service.IMicro_storeDataService
}
// Call is a single request handler called via client.Call or the generated client code
func (e *Micro_store) Call(ctx context.Context, req *micro_store.Request, rsp *micro_store.Response) error {
	log.Info("Received Micro_store.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Micro_store) Stream(ctx context.Context, req *micro_store.StreamingRequest, stream micro_store.Micro_store_StreamStream) error {
	log.Infof("Received Micro_store.Stream request with count: %d", req.Count)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&micro_store.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}
	return nil
}
// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Micro_store) PingPong(ctx context.Context, stream micro_store.Micro_store_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&micro_store.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
