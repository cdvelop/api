module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.34

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/input v0.0.13
)

require (
	github.com/cdvelop/gotools v0.0.15
	golang.org/x/text v0.9.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/cutkey => ../cutkey
