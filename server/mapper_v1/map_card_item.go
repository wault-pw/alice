package mapper_v1

import (
	"github.com/oka-is/alice/desc/alice_v1"
	"github.com/oka-is/alice/pkg/domain"
)

func MapCardItem(input domain.CardItem) *alice_v1.CardItem {
	return &alice_v1.CardItem{
		Id:       input.ID.String,
		CardId:   input.CardID.String,
		TitleEnc: input.TitleEnc.Bytea,
		BodyEnc:  input.BodyEnc.Bytea,
	}
}

func MapCardItems(input []domain.CardItem) []*alice_v1.CardItem {
	out := make([]*alice_v1.CardItem, len(input))

	for ix := range input {
		out[ix] = MapCardItem(input[ix])
	}

	return out
}
