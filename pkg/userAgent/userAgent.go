package userAgent

import "github.com/mssola/user_agent"

type Ua struct {
	Mozilla        string
	Platform       string
	Os             string
	Localization   string
	Bot            bool
	Mobile         bool
	Engine         string
	EngineVersion  string
	Browser        string
	BrowserVersion string
}

func UaSearch(s string) Ua {
	ua := user_agent.New(s)
	var uaStruct Ua
	uaStruct.Mozilla = ua.Mozilla()
	uaStruct.Platform = ua.Platform()
	uaStruct.Os = ua.OS()
	uaStruct.Localization = ua.Localization()
	uaStruct.Bot = ua.Bot()
	uaStruct.Mobile = ua.Mobile()
	uaStruct.Engine, uaStruct.EngineVersion = ua.Engine()
	uaStruct.Browser, uaStruct.BrowserVersion = ua.Browser()
	return uaStruct
}
