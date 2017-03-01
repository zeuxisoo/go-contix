package models

type Performance struct {
    IsFirstDayPerformance                bool    `json:"isFirstDayPerformance"`
    PerformanceAcsId                     int     `json:"performanceAcsId"`
    PerformanceDisplayFormat             string  `json:"performanceDisplayFormat"`
    EventId                              int     `json:"eventId"`
    PerformanceId                        int     `json:"performanceId"`
    BookmarkCreateTime                   int     `json:"bookmarkCreateTime"`
    BookmarkStatus                       int     `json:"bookmarkStatus"`
    PerformanceCategoryClass             string  `json:"performanceCategoryClass"`
    TransactionMaxQuota                  int     `json:"transactionMaxQuota"`
    PerformanceDateTime                  int64   `json:"performanceDateTime"`
    IsPurchasable                        bool    `json:"isPurchasable"`
    CounterSalesStartDate                *string `json:"counterSalesStartDate"`
    CounterSalesEndDate                  *string `json:"counterSalesEndDate"`
    DisplayDate                          bool    `json:"displayDate"`
    DisplayTime                          bool    `json:"displayTime"`
    ExternalReferenceKey                 string  `json:"externalReferenceKey"`
    PerformanceDisplayFormatValue        int     `json:"performanceDisplayFormatValue"`
    IsNotAllowedToPurchaseBeforeShowTime bool    `json:"isNotAllowedToPurchaseBeforeShowTime"`
    Note                                 *string `json:"note"`
    PerformanceName                      string  `json:"performanceName"`
}

type PerformanceData struct {
    PerformanceList []Performance `json:"performanceList"`
    StatusList      []string      `json:"performanceQuotaStatusList"`
}

type PerformanceList struct {
    Name    string
    Time    string
    Status  string
}
