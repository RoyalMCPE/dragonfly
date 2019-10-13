package session

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
)

// handleText ...
func (s *Session) handleText(pk *packet.Text) error {
	if pk.TextType != packet.TextTypeChat {
		return fmt.Errorf("text packet can only contain text type of type chat (%v) but got %v", packet.TextTypeChat, pk.TextType)
	}
	if pk.SourceName != s.conn.IdentityData().DisplayName {
		return fmt.Errorf("text packet source name must be equal to display name")
	}
	s.c.Chat(pk.Message)
	return nil
}

// SendMessage ...
func (s *Session) SendMessage(message string) {
	s.writePacket(&packet.Text{
		TextType: packet.TextTypeRaw,
		Message:  message,
	})
}

// SendTip ...
func (s *Session) SendTip(message string) {
	s.writePacket(&packet.Text{
		TextType: packet.TextTypePopup,
		Message:  message,
	})
}

// SendAnnouncement ...
func (s *Session) SendAnnouncement(message string) {
	s.writePacket(&packet.Text{
		TextType: packet.TextTypeAnnouncement,
		Message:  message,
	})
}

// SendPopup ...
func (s *Session) SendPopup(message string) {
	s.writePacket(&packet.Text{
		TextType: packet.TextTypePopup,
		Message:  message,
	})
}

// SendJukeBoxPopup ...
func (s *Session) SendJukeBoxPopup(message string) {
	s.writePacket(&packet.Text{
		TextType: packet.TextTypeJukeboxPopup,
		Message:  message,
	})
}

// SendScoreboard ...
func (s *Session) SendScoreboard(displayName string, objName string) {
	s.writePacket(&packet.SetDisplayObjective{
		DisplaySlot:   "sidebar",
		ObjectiveName: objName,
		DisplayName:   displayName,
		CriteriaName:  "dummy",
	})
}

// RemoveScoreboard ...
func (s *Session) RemoveScoreboard(objName string) {
	s.writePacket(&packet.RemoveObjective{
		ObjectiveName: objName,
	})
}

const tickLength = time.Second / 20

// SetTitleDurations ...
func (s *Session) SetTitleDurations(fadeInDuration, remainDuration, fadeOutDuration time.Duration) {
	s.writePacket(&packet.SetTitle{
		ActionType:      packet.TitleActionSetDurations,
		FadeInDuration:  int32(fadeInDuration / tickLength),
		RemainDuration:  int32(remainDuration / tickLength),
		FadeOutDuration: int32(fadeOutDuration / tickLength),
	})
}

// SendTitle ...
func (s *Session) SendTitle(text string) {
	s.writePacket(&packet.SetTitle{ActionType: packet.TitleActionSetTitle, Text: text})
}

// SendSubtitle ...
func (s *Session) SendSubtitle(text string) {
	s.writePacket(&packet.SetTitle{ActionType: packet.TitleActionSetSubtitle, Text: text})
}

// SendActionbarMessage ...
func (s *Session) SendActionBarMessage(text string) {
	s.writePacket(&packet.SetTitle{ActionType: packet.TitleActionSetActionBar, Text: text})
}
