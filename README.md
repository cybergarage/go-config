# go-json-pointer

go-json-pointer is a utility library to access JSON based configuration easily. Go language supports JSON format in the standard library. However, I feel that the standard library is not useful to parse any JSON formats [\[1\]][json-go] [\[2\]][encoding-json].

For that reason, I have developed the open source package to access JSON based configuration easily. Using go-json-pointer, you can get configuration values in the specified JSON file or string by the given path like Path. The example is bellow.

```
import (
	"json/pointer"
)

config, err := xjson.NewConfig()
if err != nil {
	t.Error(err)
}

err = config.ParseFromFile("/etc/profile.conf")
if err != nil {
	t.Error(err)
}

name, err := config.GetKeyStringByPath("/organizer/name")
if err != nil {
	t.Error(err)
}

age, err := config.GetKeyStringByPath("/organizer/age")
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

- [GitHub](https://github.com/cybergarage/go-json-pointer)

## Documents

Please check the godoc of `config/jpath` as the following.

```
godoc -http=:6060
```

## References

- \[1\] [JSON and Go][json-go]
- \[2\] [encoding/json][encoding-json]
- \[3\] [RFC 6901 - JavaScript Object Notation (JSON) Pointer][json-pointer]

[json-go]: http://blog.golang.org/json-and-go
[encoding-json]: http://golang.org/pkg/encoding/json/
[json-pointer]: https://tools.ietf.org/html/rfc6901