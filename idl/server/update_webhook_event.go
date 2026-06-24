package server

import (
	"context"
	"errors"
	webhook_eventmodule "github.com/mklfarha/mfconsult/core/module/webhook_event"
	"github.com/mklfarha/mfconsult/core/module/webhook_event/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateWebhookEvent(ctx context.Context, req *pb.UpdateWebhookEventRequest) (*pb.WebhookEvent, error) {

	if req.WebhookEvent.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetWebhookEvent())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetWebhookEvent(), "id")
		}

		pkEntity := pbmapper.WebhookEventFromProto(req.GetWebhookEvent())
		existingRes, err := s.core.WebhookEvent().FetchWebhookEventById(ctx,
			types.FetchWebhookEventByIdRequest{
				ID: pkEntity.ID,
			},
			webhook_eventmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.WebhookEventToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetWebhookEvent())
		req = &pb.UpdateWebhookEventRequest{WebhookEvent: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.WebhookEvent().Update(ctx, types.UpsertRequest{
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
