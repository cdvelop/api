module github.com/cdvelop/api

go 1.20

require github.com/cdvelop/model v0.0.68

require (
	github.com/cdvelop/input v0.0.51
	github.com/cdvelop/testools v0.0.35
)

require (
	github.com/cdvelop/cutkey v0.6.0 // indirect
	github.com/cdvelop/maps v0.0.1 // indirect
	github.com/cdvelop/object v0.0.26 // indirect
	github.com/cdvelop/strings v0.0.6 // indirect
	github.com/cdvelop/timeserver v0.0.12 // indirect
	github.com/cdvelop/timetools v0.0.17 // indirect
)

require (
	github.com/cdvelop/fetchserver v0.0.2 // indirect
	github.com/cdvelop/filehandler v0.0.3
	github.com/cdvelop/gotools v0.0.56
	github.com/cdvelop/logserver v0.0.3 // indirect
	github.com/cdvelop/output v0.0.15
	github.com/cdvelop/unixid v0.0.17
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/logserver => ../logserver

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/testools => ../testools

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/output => ../output
