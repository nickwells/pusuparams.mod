package pusuparams_test

import (
	"errors"
	"testing"
	"time"

	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/param.mod/v6/paramset"
	"github.com/nickwells/pusu.mod/pusuclt"
	"github.com/nickwells/pusuparams.mod/pusuparams"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

type cciSetter func(*pusuclt.ConnInfo)

// setSvrAddress returns a cci setter that will set the SvrAddress to the
// given string
func setSvrAddress(addr string) cciSetter {
	return func(cci *pusuclt.ConnInfo) {
		cci.SvrAddress = addr
	}
}

// setConnTimeout returns a cci setter that will set the ConnTimeout to the
// given duration
func setConnTimeout(d time.Duration) cciSetter {
	return func(cci *pusuclt.ConnInfo) {
		cci.ConnTimeout = d
	}
}

// setPingInterval returns a cci setter that will set the PingInterval to the
// given duration
func setPingInterval(d time.Duration) cciSetter {
	return func(cci *pusuclt.ConnInfo) {
		cci.PingInterval = d
	}
}

func TestAddPusuParams(t *testing.T) {
	svrAddr := "localhost:4242"
	connTimeout := "42s"
	connTimeoutDuration := 42 * time.Second
	pingInterval := "42s"
	pingIntervalDuration := 42 * time.Second

	testCases := []struct {
		testhelper.ID
		params        []string
		expCCISetters []cciSetter
		expErrMap     param.ErrMap
	}{
		{
			ID: testhelper.MkID("set address"),
			params: []string{
				"-pubsub-server-address", svrAddr,
			},
			expCCISetters: []cciSetter{setSvrAddress(svrAddr)},
		},
		{
			ID: testhelper.MkID("set timeout and address"),
			params: []string{
				"-pubsub-server-address", svrAddr,
				"-pubsub-conn-timeout", connTimeout,
			},
			expCCISetters: []cciSetter{
				setConnTimeout(connTimeoutDuration),
				setSvrAddress(svrAddr),
			},
		},
		{
			ID: testhelper.MkID("set ping interval and address"),
			params: []string{
				"-pubsub-server-address", svrAddr,
				"-pubsub-ping-interval", pingInterval,
			},
			expCCISetters: []cciSetter{
				setPingInterval(pingIntervalDuration),
				setSvrAddress(svrAddr),
			},
		},
		{
			ID: testhelper.MkID("missing address"),
			expErrMap: param.ErrMap{
				"pubsub-server-address": []error{
					errors.New("this parameter must be set somewhere"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			expectedCCI := pusuclt.NewConnInfo(nil)
			for _, f := range tc.expCCISetters {
				f(expectedCCI)
			}

			targetCCI := pusuclt.NewConnInfo(nil)

			ps := paramset.NewNoHelpNoExitNoErrRptOrPanic(
				pusuparams.AddPusuParams(targetCCI, ""))

			errMap := ps.Parse(tc.params)
			if err := testhelper.DiffVals(errMap, tc.expErrMap); err != nil {
				t.Log(tc.IDStr())
				t.Logf("\t%s", err)
				t.Error("\terror maps differ\n")
			}

			if err := testhelper.DiffVals(targetCCI, expectedCCI); err != nil {
				t.Log(tc.IDStr())
				t.Logf("\t%s", err)
				t.Error("\tpost-parsing values differ\n")
			}
		})
	}
}
