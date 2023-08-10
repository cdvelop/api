module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.36

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/input v0.0.13
	github.com/cdvelop/testools v0.0.0-00010101000000-000000000000
)

require (
	github.com/cdvelop/gotools v0.0.22
	golang.org/x/text v0.9.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/testools => ../testools

replace github.com/cdvelop/gotools => ../gotools
