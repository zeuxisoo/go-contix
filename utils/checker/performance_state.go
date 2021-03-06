package checker

import (
    "fmt"
    "errors"
    "time"
    "encoding/json"
    "math/rand"
    "crypto/tls"

    "github.com/parnurzeal/gorequest"

    "github.com/zeuxisoo/go-contix/models"
)

type PerformanceStateChecker struct {
    Agent               *gorequest.SuperAgent
    PerformanceList     []models.PerformanceList

    PerformanceId       string
    PerPage             int
}

func NewPerformanceStateChecker() *PerformanceStateChecker {
    request := gorequest.New().
        TLSClientConfig(&tls.Config{ InsecureSkipVerify: true }).
        Timeout(3000 * time.Millisecond).
        Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8").
        Set("Accept-Language", "en-US,en;q=0.8").
        Set("Connection", "keep-alive")

    checker := &PerformanceStateChecker{
        Agent          : request,
        PerformanceList: []models.PerformanceList{},
        PerformanceId  : "",
        PerPage        : 5,
    }

    // Set the default user agent to prevent the user agent is empty
    return checker.SetUserAgents([]models.CronTaskUserAgent{
        models.CronTaskUserAgent{
            Name : "Default Agent A",
            Agent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36",
        },
        models.CronTaskUserAgent{
            Name : "Default Agent B",
            Agent: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; fr) AppleWebKit/416.12 (KHTML, like Gecko) Safari/412.5",
        },
        models.CronTaskUserAgent{
            Name : "Default Agent C",
            Agent: "Mozilla/5.0 (Windows NT 6.1; rv:15.0) Gecko/20120819 Firefox/15.0 PaleMoon/15.0",
        },
        models.CronTaskUserAgent{
            Name : "Default Agent D",
            Agent: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB6; Acoo Browser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
        },
        models.CronTaskUserAgent{
            Name : "Default Agent E",
            Agent: "Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-BR) AppleWebKit/534.12 (KHTML, like Gecko) NavscapeNavigator/Pre-0.1 Safari/534.12",
        },
        models.CronTaskUserAgent{
            Name : "Default Agent F",
            Agent: "Mozilla/5.0 (Windows; U; WinNT4.0; de-AT; rv:1.7.11) Gecko/20050728",
        },
    })
}

func (this *PerformanceStateChecker) SetPerformanceId(id string) *PerformanceStateChecker {
    this.PerformanceId = id
    return this
}

func (this *PerformanceStateChecker) SetProxy(proxy string) *PerformanceStateChecker {
    if proxy != "" {
        this.Agent = this.Agent.Proxy(proxy)
    }

    return this
}

func (this *PerformanceStateChecker) SetUserAgents(userAgents []models.CronTaskUserAgent) *PerformanceStateChecker {
    if len(userAgents) > 0 {
        this.Agent = this.Agent.Set("User-Agent", randomUserAgent(userAgents))
    }

    return this
}

func (this *PerformanceStateChecker) SetTimeout(timeout int) *PerformanceStateChecker {
    if (timeout > 1000) {
        this.Agent = this.Agent.Timeout(time.Duration(timeout) * time.Millisecond)
    }

    return this
}

func (this *PerformanceStateChecker) GetPerformanceList() ([]models.PerformanceList, error) {
    if _, err := this.makeAuth(); err != nil {
        return nil, err
    }

    if events, err := this.fetchEvent(1); err != nil {
        return nil, err
    }else{
        return events, nil
    }
}


func (this PerformanceStateChecker) makeAuth() (string, error) {
    response, body, errs := this.Agent.Get("http://www.urbtix.hk/").End()
    if errs != nil {
        return "", errs[0]
    }

    if response.Status != "200 OK" {
        return "", errors.New("Target page return status code: " + response.Status)
    }

    if response.Request.URL.String() != "https://ticket.urbtix.hk/internet/" {
        return "", errors.New("Redirect error: " + response.Request.URL.String())
    }

    return body, nil
}

func (this *PerformanceStateChecker) fetchEvent(pageNo int) ([]models.PerformanceList, error) {
    timestamp := time.Now().Unix()
    targetUrl := fmt.Sprintf("https://ticket.urbtix.hk/internet/json/event/%s/performance/%d/%d/perf.json?locale=zh_TW&%d", this.PerformanceId, this.PerPage, pageNo, timestamp)

    _, body, errs := this.Agent.Get(targetUrl).End()
    if errs != nil {
        return this.PerformanceList, errs[0]
    }

    var performanceData models.PerformanceData
    if err := json.Unmarshal([]byte(body), &performanceData); err != nil {
        return this.PerformanceList, err
    }else{
        for k, v := range performanceData.PerformanceList {
            timeString := time.Unix(v.PerformanceDateTime/1000, 0).Format(time.RFC3339)

            this.PerformanceList = append(this.PerformanceList, models.PerformanceList{
                Name  : v.PerformanceName,
                Time  : timeString,
                Status: performanceData.StatusList[k],
            })
        }

        if len(performanceData.PerformanceList) > 0 {
            return this.fetchEvent(pageNo + 1)
        }else{
            return this.PerformanceList, nil
        }
    }
}

func randomUserAgent(userAgents []models.CronTaskUserAgent) string {
    rand.Seed(time.Now().UTC().UnixNano())

    return userAgents[rand.Intn(len(userAgents))].Agent
}
