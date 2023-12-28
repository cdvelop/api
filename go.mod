module github.com/cdvelop/api

go 1.20

require (
	github.com/cdvelop/filehandler v0.0.35
	github.com/cdvelop/fileserver v0.0.54
	github.com/cdvelop/input v0.0.81
	github.com/cdvelop/model v0.0.107
	github.com/cdvelop/output v0.0.16
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/structs v0.0.1
	github.com/cdvelop/testools v0.0.79
	github.com/cdvelop/unixid v0.0.48
)

require (
	github.com/cdvelop/cutkey v1.0.14 // indirect
	github.com/cdvelop/fetchserver v0.0.26 // indirect
	github.com/cdvelop/logserver v0.0.33 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/object v0.0.67 // indirect
	github.com/cdvelop/timetools v0.0.33 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/testools => ../testools
