package consts

var ContentTypeRe = map[string]string{
	"":                                  "",
	"application/json":                  ".*?application\\/json.*?",
	"application/xml":                   ".*?application\\/xml.*?",
	"application/x-www-form-urlencoded": ".*?application/x-www-form-urlencoded.*?",
	"multipart/form-data":               ".*?multipart/form-data.*?",
}
