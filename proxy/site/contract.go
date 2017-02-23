package site

type Contract interface {
    Name() (string)
    Fetch() ([]ProxyInfo, error)
}

type ProxyInfo struct {
    IP          string
    Port        string
    Protocol    string
    Country     string
}
