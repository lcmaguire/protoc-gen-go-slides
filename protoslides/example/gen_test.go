package example

import (
	"errors"
	"testing"

	"github.com/lcmaguire/protoc-gen-go-slides/example"
)

func TestXxx(t *testing.T) {
	message := &example.Post{}

	message.PostContent = &example.Post_Audio{
		Audio: &example.Audio{
			/* now populate entire audio*/
		},
	}

	// with setters
	message.SetAudio(&example.Audio{ /* now populate entire audio*/ })

	message.SetText("post")

	message.PostContent = &example.Post_Text{
		Text: "text",
	}

}

type PageSizer interface {
	GetPageSize() int32
}

func HandlePageSize(in PageSizer) (int32, error) {
	if in.GetPageSize() < 0 {
		return -1, errors.New("invalid cannot be less than 0")
	}

	if in.GetPageSize() > 1000 {
		return 1000, nil
	}

	if in.GetPageSize() == 0 {
		return 25, nil
	}

	return in.GetPageSize(), nil
}
