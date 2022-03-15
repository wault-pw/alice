package api_v1

import (
	"net/http"

	"github.com/wault-pw/alice/desc/alice_v1"
	"github.com/wault-pw/alice/server/engine"
	"github.com/wault-pw/alice/server/mapper_v1"
)

func CloneCard(ctx *engine.Context) {
	cardID, _ := ctx.Param(paramCardID), ctx.Param(paramWorkspaceID)
	req := new(alice_v1.CloneCardRequest)
	err := ctx.MustBindProto(req)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	user := ctx.MustGetUser()
	err = ctx.NewUserPolicy(user).CanWrite()
	if err != nil {
		ctx.HandleError(err)
		return
	}

	card, err := ctx.GetStore().CloneCard(ctx.Context, cardID, req.GetTitleEnc())
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.ProtoBuf(http.StatusOK, &alice_v1.CloneCardResponse{
		Card: mapper_v1.MapCard(card),
	})
}
