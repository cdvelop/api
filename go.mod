module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.63

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fileserver v0.0.13
	github.com/cdvelop/input v0.0.47
	github.com/cdvelop/testools v0.0.31
)

require (
	github.com/cdvelop/timetools v0.0.13 // indirect
	golang.org/x/net v0.17.0 // indirect
)

require (
	github.com/cdvelop/gotools v0.0.52
	github.com/cdvelop/output v0.0.12
	github.com/cdvelop/unixid v0.0.13
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
