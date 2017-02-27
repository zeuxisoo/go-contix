package site

import (
    "encoding/json"

    "github.com/parnurzeal/gorequest"

    "github.com/zeuxisoo/go-contix/models"
)

type NyLonerProxyList struct {
    Time    string  `json:"time"`
    IP      string  `json:"ip"`
    Port    string  `json:"port"`
}

type NyLonerProxy struct {
    List    []NyLonerProxyList  `json:"list"`
    Status  string              `json:"status"`
}

type NyLonerProxySite struct {
}

func (this *NyLonerProxySite) Name() (string) {
    return "Hide My Ass"
}

func (this *NyLonerProxySite) Fetch() ([]models.ProxyInfo, error) {
    request := gorequest.New()
    _, body, errs := request.Get("http://nyloner.cn/proxy?page=1&num=15").End()
    if errs != nil {
        return nil, errs[0]
    }

    var nylonerProxy NyLonerProxy
    if err := json.Unmarshal([]byte(body), &nylonerProxy); err != nil {
        return nil, err
    }

    var proxyList []models.ProxyInfo
    for _, proxy := range nylonerProxy.List {
        proxyList = append(proxyList, models.ProxyInfo{
            IP      : proxy.IP,
            Port    : proxy.Port,
            Protocol: "http",
            Country : "",
        })
    }

    return proxyList, nil
}
