package example

import (
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
