package config

import "os"

var SECRET_JWT = os.Getenv("SECRET_JWT")
