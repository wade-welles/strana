package tracker

import (
	"testing"
	"time"

	"github.com/blushft/strana/event/contexts"
	"github.com/stretchr/testify/suite"
)

type TrackerSuite struct {
	suite.Suite
	tr *Tracker
}

func TestRunTrackerSuite(t *testing.T) {
	suite.Run(t, new(TrackerSuite))
}

func (s *TrackerSuite) SetupSuite() {
	tr, err := New(TrackingID("1234"), SetAppInfo(&contexts.App{
		Name:    "tracker_test",
		Version: "v0.1.0",
	}))
	s.Require().NoError(err)
	s.tr = tr
}

func (s *TrackerSuite) TearDownSuite() {
	s.tr.Close()
}

func (s *TrackerSuite) TestTrackAction() {
	a := &contexts.Action{
		Category: "tests",
		Action:   "Test Action",
		Label:    "test_track",
		Property: "test_val",
		Value:    99,
	}

	s.Require().NoError(s.tr.Action(a))

	time.Sleep(time.Second * 5)
}
