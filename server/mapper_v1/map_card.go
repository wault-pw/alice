package mapper_v1

import (
	"github.com/wault-pw/alice/desc/alice_v1"
	"github.com/wault-pw/alice/pkg/domain"
)

func MapListCardsResponse(input []domain.Card) *alice_v1.ListCardsResponse {
	return &alice_v1.ListCardsResponse{
		Items: MapCards(input),
	}
}

func MapCard(input domain.Card) *alice_v1.Card {
	return &alice_v1.Card{
		Id:          input.ID.String,
		Archived:    input.Archived.Bool,
		WorkspaceId: input.WorkspaceID.String,
		TitleEnc:    input.TitleEnc.Bytea,
		TagsEnc:     input.TagsEnc.Slice,
	}
}

func MapCards(input []domain.Card) []*alice_v1.Card {
	out := make([]*alice_v1.Card, len(input))

	for ix := range input {
		out[ix] = MapCard(input[ix])
	}

	return out
}
