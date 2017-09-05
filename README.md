# go-config

`go-config` is a utility library to access JSON based configuration easily. Go language supports JSON format in the standard library. However, I feel that the standard library is not useful to parse any JSON formats [\[1\]][json-go] [\[2\]][encoding-json].

For that reason, I have developed the open source package to access JSON properties easily. Using `go-config`, you can get the JSON properties in the specified JSON file or string by the given path like XPath for XML specification.

Finally, I will support RFC 6901, JSON Pointer, [\[3\]][json-pointer] specification.

Using `go-config`, you can read JSON easily as the following.

Using `Config`, you can read a configuration file based on JSON easily. The sample configuration is bellow.

```
	config, err := NewConfig()
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

- /etc/profile.conf
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

- [GitHub](https://github.com/cybergarage/go-config)

## Documents

Please check the godoc as the following.

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
