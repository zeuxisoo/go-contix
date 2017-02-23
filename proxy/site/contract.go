package site

type Contract interface {
    Name() (string)
    Fetch() ([]string, error)
}
