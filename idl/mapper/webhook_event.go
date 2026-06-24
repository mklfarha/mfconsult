package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/webhook_event"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/mklfarha/mfconsult/enum"

	"github.com/guregu/null/v6"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func WebhookEventToProto(e main_entity.WebhookEvent) *pb.WebhookEvent {
	return &pb.WebhookEvent{
		Id:          e.ID,
		Source:      pb.WebhookSource(e.Source),
		EventType:   e.EventType.ValueOrZero(),
		Payload:     e.Payload.ValueOrZero(),
		ProcessedAt: timestamppb.New(e.ProcessedAt.ValueOrZero()),
		CreatedAt:   timestamppb.New(e.CreatedAt.ValueOrZero()),
	}
}

func WebhookEventSliceToProto(es []main_entity.WebhookEvent) []*pb.WebhookEvent {
	res := []*pb.WebhookEvent{}
	for _, e := range es {
		res = append(res, WebhookEventToProto(e))
	}
	return res
}

func WebhookEventFromProto(m *pb.WebhookEvent) main_entity.WebhookEvent {
	if m == nil {
		return main_entity.WebhookEvent{}
	}
	return main_entity.WebhookEvent{
		ID:          m.GetId(),
		Source:      enum.WebhookSource(m.GetSource()),
		EventType:   null.StringFrom(m.EventType),
		Payload:     null.StringFrom(m.Payload),
		ProcessedAt: null.TimeFrom(m.GetProcessedAt().AsTime()),
		CreatedAt:   null.TimeFrom(m.GetCreatedAt().AsTime()),
	}
}

func WebhookEventSliceFromProto(es []*pb.WebhookEvent) []main_entity.WebhookEvent {
	if es == nil {
		return []main_entity.WebhookEvent{}
	}
	res := []main_entity.WebhookEvent{}
	for _, e := range es {
		res = append(res, WebhookEventFromProto(e))
	}
	return res
}
