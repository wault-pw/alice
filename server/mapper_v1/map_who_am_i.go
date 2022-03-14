package mapper_v1

import (
	"github.com/wault-pw/alice/desc/alice_v1"
	"github.com/wault-pw/alice/pkg/domain"
)

func MapWhoAmI(input domain.User) *alice_v1.WhoAmIResponse {
	return &alice_v1.WhoAmIResponse{
		Id:         input.ID.String,
		PasswdSalt: input.PasswdSalt.Bytea,
		PrivKeyEnc: input.PrivKeyEnc.Bytea,
		PubKey:     input.PubKey.Bytea,
	}
}
