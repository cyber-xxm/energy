package main

import (
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/gson"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/launcher"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func getSchema() gson.JSON {
	//var flagDevToolWsURL = "ws://127.0.0.1:9222/devtools/browser/41c8af06-f898-4e32-bb5c-a65506066a3f"
	//browse := rod.New().ControlURL(flagDevToolWsURL).MustConnect()
	//l := launcher.New().Bin(launcher.NewBrowser().MustGet())
	l := launcher.New().Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe")
	defer l.Kill()
	u := l.MustLaunch()
	parsed, err := url.Parse(u)
	utils.E(err)
	parsed.Scheme = "http"
	parsed.Path = "/json/protocol"
	fmt.Println("parsed.String():", parsed.String())
	res, err := http.Get(parsed.String()) //nolint: noctx
	utils.E(err)
	defer func() { _ = res.Body.Close() }()

	data, err := ioutil.ReadAll(res.Body)
	utils.E(err)

	obj := gson.New(data)

	utils.E(utils.OutputFile("tmp/proto.json", obj.JSON("", "  ")))

	return obj
}

func mapType(n string) string {
	return map[string]string{
		"boolean": "bool",
		"number":  "float64",
		"integer": "int",
		"string":  "string",
		"binary":  "[]byte",
		"object":  "map[string]gson.JSON",
		"any":     "gson.JSON",
	}[n]
}

func typeName(domain *domain, schema gson.JSON) string {
	typeName := ""
	if schema.Has("type") {
		typeName = schema.Get("type").Str()
	}

	if typeName == "array" { //nolint: nestif
		item := schema.Get("items")

		if item.Has("type") {
			typeName = "[]" + mapType(item.Get("type").Str())
		} else {
			ref := item.Get("$ref").Str()
			if domain.ref(ref) {
				typeName = "[]*" + refName(domain.name, ref)
			} else {
				typeName = "[]" + refName(domain.name, ref)
			}
		}
	} else if schema.Has("$ref") {
		ref := schema.Get("$ref").Str()
		if domain.ref(ref) {
			typeName += "*"
		}
		typeName += refName(domain.name, ref)
	} else {
		typeName = mapType(typeName)
	}

	switch typeName {
	case "NetworkTimeSinceEpoch", "InputTimeSinceEpoch":
		typeName = "TimeSinceEpoch"
	case "NetworkMonotonicTime":
		typeName = "MonotonicTime"
	}

	return typeName
}

func enumList(schema gson.JSON) []string {
	var enum []string
	if schema.Has("enum") {
		enum = []string{}
		for _, v := range schema.Get("enum").Arr() {
			if _, ok := v.Val().(string); !ok {
				panic("enum type error")
			}
			enum = append(enum, v.Str())
		}
	}

	return enum
}

func jsonTag(name string, optional bool) string {
	jsonTagValue := name
	if optional {
		jsonTagValue += ",omitempty"
	}
	return fmt.Sprintf("`json:\"%s\"`", jsonTagValue)
}

func refName(domain, id string) string {
	if strings.Contains(id, ".") {
		return symbol(id)
	}
	return domain + symbol(id)
}

// make sure golint works fine.
func symbol(n string) string {
	if n == "" {
		return ""
	}

	n = strings.ReplaceAll(n, ".", "")

	dashed := regexp.MustCompile(`[-_]`).Split(n, -1)
	if len(dashed) > 1 {
		converted := []string{}
		for _, part := range dashed {
			converted = append(converted, strings.ToUpper(part[:1])+part[1:])
		}
		n = strings.Join(converted, "")
	}

	n = strings.ToUpper(n[:1]) + n[1:]

	n = replaceLower(n, "Id")
	n = replaceLower(n, "Css")
	n = replaceLower(n, "Url")
	n = replaceLower(n, "Uuid")
	n = replaceLower(n, "Xml")
	n = replaceLower(n, "Http")
	n = replaceLower(n, "Dns")
	n = replaceLower(n, "Cpu")
	n = replaceLower(n, "Mime")
	n = replaceLower(n, "Json")
	n = replaceLower(n, "Html")
	n = replaceLower(n, "Guid")
	n = replaceLower(n, "Sql")
	n = replaceLower(n, "Eof")
	n = replaceLower(n, "Api")
	n = replaceLower(n, "Ui")
	n = replaceLower(n, "Https")

	n = strings.Replace(n, "Ids", "IDs", -1)

	return n
}

func replaceLower(n, word string) string {
	return regexp.MustCompile(word+`([A-Z-_]|$)`).ReplaceAllStringFunc(n, strings.ToUpper)
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}