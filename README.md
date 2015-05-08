# go-json-path

go-json-path is a utility library to access JSON based configuration easily. Go language supports JSON format in the standard library. However, I feel that the standard library is not useful to parse any JSON formats [\[1\]][json-go] [\[2\]][encoding-json].

For that reason, I have developed the open source package to access JSON based configuration easily. Using go-json-path, you can get configuration values in the specified JSON file or string by the given path like XPath. The example is bellow.

```
import (
	"config/jpath"
)

config, err := xjson.NewConfig()
if err != nil {
	t.Error(err)
}

err = config.ParseFromFile("/etc/profile.conf")
if err != nil {
	t.Error(err)
}

name, err := config.GetKeyStringByXPath("/organizer/name")
if err != nil {
	t.Error(err)
}

age, err := config.GetKeyStringByXPath("/organizer/age")
if err != nil {
	t.Error(err)
}
```

The configuration file format is based on JSON as the following.

```
#
#  /etc/profile.conf
#

{
	"organizer": {
		"name": "John Smith",
		"age": 33
	}
}
```

## Repository

- [GitHub](https://github.com/cybergarage/go-json-path)

## Documents

Please check the godoc of `config/jpath` as the following.

```
godoc -http=:6060
```

## References

- \[1\] [JSON and Go][json-go]
- \[2\] [encoding/json][encoding-json]

[json-go]: http://blog.golang.org/json-and-go
[encoding-json]: http://golang.org/pkg/encoding/json/
