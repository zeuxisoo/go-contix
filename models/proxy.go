package models

type ProxyInfo struct {
    IP          string
    Port        string
    Protocol    string
    Country     string
}

type ProxyState struct {
    Usable  bool
    Proxy   string
}
