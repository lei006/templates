package webrtc_camera

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
)

type SessionDescriptionEx struct {
	webrtc.SessionDescription
}

func (desc *SessionDescriptionEx) ToJson() (string, error) {
	b, err := json.Marshal(desc)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (desc *SessionDescriptionEx) FromJson(str string) error {

	err := json.Unmarshal([]byte(str), desc)
	if err != nil {
		return err
	}

	return nil
}
