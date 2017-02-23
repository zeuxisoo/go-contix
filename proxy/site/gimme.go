package site

import (
    "fmt"
    "time"
    "encoding/json"

    "github.com/parnurzeal/gorequest"
)

type GimmeProxyWebsites struct {
    Example bool `json:"example"`
    Google  bool `json:"google"`
    Amazon  bool `json:"amazon"`
}

type GimmeProxy struct {
    Get             bool                `json:"get"`
    Post            bool                `json:"post"`
    cookies         bool                `json:"cookies"`
    Referer         bool                `json:"referer"`
    UserAgent       bool                `json:"user-agent"`
    AnonymityLevel  int                 `json:"anonymityLevel"`
    SupportsHTTPS   bool                `json:"supportsHttps"`
    Protocol        string              `json:"protocol"`
    IP              string              `json:"ip"`
    Port            string              `json:"port"`
    Websites        GimmeProxyWebsites  `json:"websites"`
    Country         string              `json:"country"`
    TsChecked       int                 `json:"tsChecked"`
    Curl            string              `json:"curl"`
    IPPort          string              `json:"ipPort"`
    Type            string              `json:"type"`
    Speed           float64             `json:"speed"`
    OtherProtocols  struct{}            `json:"otherProtocols"`
}

type GimmeProxySite struct {
}

func (this *GimmeProxySite) Name() (string) {
    return "Gimme"
}

func (this *GimmeProxySite) Fetch() ([]string, error) {
    request    := gorequest.New()
    gimmeProxy := GimmeProxy{}

    var proxyList []string
    for i := 0; i < 5; i++ {
        _, body, _ := request.Get("https://gimmeproxy.com/api/getProxy").End()

        if err := json.Unmarshal([]byte(body), &gimmeProxy); err != nil {
            continue
        }

        proxyList = append(proxyList, fmt.Sprintf("%s:%s", gimmeProxy.IP, gimmeProxy.Port))

        time.Sleep(time.Second * 2)
    }

    return proxyList, nil
}
