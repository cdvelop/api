module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.62

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fileserver v0.0.0-00010101000000-000000000000
	github.com/cdvelop/input v0.0.46
	github.com/cdvelop/testools v0.0.30
)

require (
	github.com/cdvelop/timetools v0.0.12 // indirect
	golang.org/x/net v0.17.0 // indirect
)

require (
	github.com/cdvelop/gotools v0.0.51
	github.com/cdvelop/output v0.0.11
	github.com/cdvelop/unixid v0.0.12
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/testools => ../testools

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/output => ../output
