module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.40

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/input v0.0.21
	github.com/cdvelop/testools v0.0.3
)

require (
	github.com/cdvelop/gotools v0.0.30
	github.com/cdvelop/output v0.0.2
	golang.org/x/text v0.12.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/testools => ../testools

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/output => ../output
