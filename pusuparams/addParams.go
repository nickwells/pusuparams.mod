package pusuparams

import (
	"time"

	"github.com/nickwells/check.mod/v2/check"
	"github.com/nickwells/filecheck.mod/filecheck"
	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/param.mod/v6/psetter"
	"github.com/nickwells/pusu.mod/pusu"
	"github.com/nickwells/pusu.mod/pusuclt"
)

const (
	paramNameServerAddress = "pubsub-server-address"
	paramNameConnTimeout   = "pubsub-conn-timeout"
	paramNamePingInterval  = "pubsub-ping-interval"

	paramNameCACertFilename = "ca-cert-filename"
	paramNameCertFilename   = "cert-filename"
	paramNameKeyFilename    = "key-filename"

	groupNameDesc = "publish/subscribe server parameters"
)

// AddPusuParams returns a function that will add parameters for setting
// information needed to connect to a publish/subscribe server
func AddPusuParams(ci *pusuclt.ConnInfo, groupName string) param.PSetOptFunc {
	prefix := ""

	if groupName != "" {
		if err := param.ParameterNameCheck(groupName); err != nil {
			panic(err)
		}

		prefix = groupName + "-"
	}

	return func(ps *param.PSet) error {
		mustBeSet := []param.OptFunc{param.Attrs(param.MustBeSet)}
		dontShow := []param.OptFunc{param.Attrs(param.DontShowInStdUsage)}

		if groupName != "" {
			ps.AddGroup(groupName, groupNameDesc)

			mustBeSet = append(mustBeSet, param.GroupName(groupName))
			dontShow = append(dontShow, param.GroupName(groupName))
		}

		ps.Add(prefix+paramNameServerAddress,
			psetter.String[string]{
				Value: &ci.SvrAddress,
			},
			"the address of the pub/sub server to connect to",
			mustBeSet...)

		ps.Add(prefix+paramNameConnTimeout,
			psetter.Duration{
				Value: &ci.ConnTimeout,
			},
			"the timeout. This is how long to wait before"+
				" abandoning the attempt to connect to"+
				" the pub/sub server",
			dontShow...)

		ps.Add(prefix+paramNamePingInterval,
			psetter.Duration{
				Value: &ci.PingInterval,
				Checks: []check.Duration{
					check.ValGT(time.Duration(0)),
				},
			},
			"the ping interval. This is how long to wait between"+
				" sending ping messages to the pub/sub server",
			dontShow...)

		return nil
	}
}

// AddCertInfoParams returns a function that will add parameters for setting
// information needed to connect to a publish/subscribe server. If the
// groupName argument is not an empty string then a parameter group will be
// created, the parameters will all be in that group and the groupName will
// be used as a prefix to the parameter names.
func AddCertInfoParams(ci *pusu.CertInfo, groupName string) param.PSetOptFunc {
	prefix := ""

	if groupName != "" {
		if err := param.ParameterNameCheck(groupName); err != nil {
			panic(err)
		}

		prefix = groupName + "-"
	}

	return func(ps *param.PSet) error {
		mustBeSet := []param.OptFunc{param.Attrs(param.MustBeSet)}

		if groupName != "" {
			ps.AddGroup(groupName, groupNameDesc)

			mustBeSet = append(mustBeSet, param.GroupName(groupName))
		}

		ps.Add(prefix+paramNameCACertFilename,
			psetter.Pathname{
				Value:       &ci.CACertFilename,
				Expectation: filecheck.FileExists(),
			},
			"the name of the file holding"+
				" the certificate for the certification authority (CA)",
			mustBeSet...)

		ps.Add(prefix+paramNameCertFilename,
			psetter.Pathname{
				Value:       &ci.CertFilename,
				Expectation: filecheck.FileExists(),
			},
			"the name of the file holding"+
				" the certificate for the program",
			mustBeSet...)

		ps.Add(prefix+paramNameKeyFilename,
			psetter.Pathname{
				Value:       &ci.KeyFilename,
				Expectation: filecheck.FileExists(),
			},
			"the name of the file holding"+
				" the private key for the program's certificate",
			mustBeSet...)

		return nil
	}
}
