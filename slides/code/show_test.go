package code

import (
	"testing"
	"time"

	protoc_gen_go "github.com/lcmaguire/protoc-gen-go/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
)

func TestXxx(t *testing.T) {
	message := &protoc_gen_go.User{UserId: uuid.NewString(), DisplayName: "Tom o Noku", Verified: true, CreateTime: timestamppb.New(time.Now().Add(-time.Hour)), PostCount: 10, Interests: []string{"surfing"},
		Membership: protoc_gen_go.MembershipType_MEMBERSHIP_FREE,
	}
	bites, _ := proto.Marshal(message)
	t.Log(bites)
	t.Log(proto.Size(message))
}
