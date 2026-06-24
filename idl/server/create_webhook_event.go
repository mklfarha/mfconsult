package server

import (
	"context"
	"errors"
	webhook_eventmodule "github.com/mklfarha/mfconsult/core/module/webhook_event"
	"github.com/mklfarha/mfconsult/core/module/webhook_event/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateWebhookEvent(ctx context.Context, req *pb.CreateWebhookEventRequest) (*pb.WebhookEvent, error) {
	res, err := s.core.WebhookEvent().Insert(ctx, types.UpsertRequest{
		WebhookEvent: pbmapper.WebhookEventFromProto(req.GetWebhookEvent()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.WebhookEvent().FetchWebhookEventById(ctx, types.FetchWebhookEventByIdRequest(res), webhook_eventmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.WebhookEventToProto(fetchRes.Results[0]), nil
}
