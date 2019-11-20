# oniontree-get-mirrors

Download `/mirrors.txt` from a remote host and extract URLs.

## Usage

```
Usage of oniontree-get-mirrors:
  -id string
    	Onion service ID
  -replace
    	Replace existing URLs
  -timeout duration
    	HTTP request timeout (default 30s)
  -verify-signature
    	Enable signature verification (default true)
```

## Example

```
# Set HTTP proxy settings
$ export http_proxy=socks5://localhost:9050
$ export https_proxy=socks5://localhost:9050

$ torsocks oniontree-get-mirrors -id darkfail
2019/11/20 15:31:58 http://darkfailllnkf4vf.onion: ok
```
