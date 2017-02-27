package site

import (
    "strings"

    "github.com/parnurzeal/gorequest"
    "github.com/PuerkitoBio/goquery"

    "github.com/zeuxisoo/go-contix/models"
)

type XiCiDaiLiProxySite struct {
}

func (this *XiCiDaiLiProxySite) Name() (string) {
    return "Xi Ci Dai Li"
}

func (this *XiCiDaiLiProxySite) Fetch() ([]models.ProxyInfo, error) {
    request := gorequest.New()
    response, _, errs := request.
        Get("http://www.xicidaili.com/nn/").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36").
        End()
    if errs != nil {
        return nil, errs[0]
    }

    document, err := goquery.NewDocumentFromResponse(response)
    if err != nil {
        return nil, err
    }

    var proxyList []models.ProxyInfo

    proxyTable := document.Find("table#ip_list")
    proxyRows  := proxyTable.Find("tr:nth-child(n+2)")

    proxyRows.Each(func(i int, s *goquery.Selection) {
        country  := s.Find("td:nth-child(1) img").AttrOr("alt", "n/a")
        ip       := s.Find("td:nth-child(2)").Text()
        port     := s.Find("td:nth-child(3)").Text()
        protocol := s.Find("td:nth-child(6)").Text()

        if ip != "" && port != "" {
            proxyList = append(proxyList, models.ProxyInfo{
                IP      : ip,
                Port    : port,
                Protocol: strings.ToLower(protocol),
                Country : strings.ToLower(country),
            })
        }
    })

    return proxyList, nil
}
