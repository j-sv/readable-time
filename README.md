# readable-time

Prints the current time in a slightly more readable way.

```bash
$ readable-time
ten past six, monday, august tenth
```

## Formatting

The command takes an optional format, in the form of a template string, which
can be used to customize the output.

The functions exported by [`time.Time`](./time/time.go#17-19) can be used like this:

```bash
$ readable-time --format '{{ .Clock }}'
ten past six
````

```bash
$ readable-time --format '{{ .Month }} {{ .Day }}'
november first
````

## When

The `--when` flag can be used to set which time is used when formatting.

It supports either a datetime given in [`time.RFC3999`](https://pkg.go.dev/time#pkg-constants)
format or a relative timestamp, e.g. "half an hour ago".

```bash
$ readable-time --format "{{ .Month }} {{ .Day }}"
november first
$ readable-time --format "{{ .Month }} {{ .Day }}" --when "a week ago"
october twenty-fifth
```
