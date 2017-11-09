# yestest

## ????????????

```bash
$ go build -o yestest
$ ./yestest | pv -r > /dev/null
[1.97MiB/s]
```

## too fast...

```bash
$ yes | pv -r > /dev/null
[31.5MiB/s]
```
