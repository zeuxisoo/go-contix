package site

import (
    "github.com/zeuxisoo/go-contix/models"
)

type Contract interface {
    Name() (string)
    Fetch() ([]models.ProxyInfo, error)
}
