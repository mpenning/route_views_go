package main

import (
	"fmt"

	"github.com/gleich/logoru"
	"github.com/melbahja/goph"
	"github.com/sirikothe/gotextfsm"
)

func main() {

	// route-views username 'rviews' with no password
	client01, err := goph.New("rviews", "route-views.routeviews.org", goph.Password(""))
	if err != nil {
		logoru.Critical(err.Error())
	}
	defer client01.Close()

	// get the best BGP route info for 4.2.2.2
	cmdout01, err := client01.Run(fmt.Sprintf("show ip bgp %s bestpath", "4.2.2.2"))
	if err != nil {
		logoru.Critical(err.Error())
	}

	// Begin gotextfsm template
	template := `Value bgpAsPath (\d+\s*.*?)
Value bgpIpv4Prefix (\d+\.\d+\.\d+\.\d+)
Value bgpIpv4NextHop (\d+\.\d+\.\d+\.\d+)
Value bgpPrefixLength (\d+)
Value bgpTableVersion (\d+)


Start
	^BGP routing table entry for ${bgpIpv4Prefix}/${bgpPrefixLength}, version ${bgpTableVersion}
	^Paths:
	^\s{2}Not advertised
	^\s{2}Refresh
	^\s{2}${bgpAsPath}
	^\s+${bgpIpv4NextHop} from
	^\s+Origin
`
	// End gotextfsm template

	// Use textfsm to read the bgp table...
	fsm := gotextfsm.TextFSM{}
	fsm.ParseString(template)
	if err != nil {
		logoru.Critical(err)
	}
	parser := gotextfsm.ParserOutput{}

	/////////////////////////////////////////////////////////////
	// parse cmdout01
	err = parser.ParseTextString(string(cmdout01), fsm, true)
	if err != nil {
		logoru.Critical(err)
	}

	// populate a map called 'output'...
	output01 := make(map[string]string)
	for _, record := range parser.Dict {
		for key, value := range record {
			switch value.(type) {
			case string:
				output01[key] = value.(string)
			}
		}
	}
	logoru.Info(output01)

}
