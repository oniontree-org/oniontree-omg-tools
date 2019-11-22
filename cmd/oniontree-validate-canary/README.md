# oniontree-validate-canary

Validate host's `/canary.txt`.

## Usage

```
Usage of oniontree-validate-canary:
  -date-only
    	Skip message validation, check date only
  -id string
    	Onion service ID
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

$ torsocks oniontree-validate-canary -id darkfail
2019/11/20 15:21:59 http://darkfailllnkf4vf.onion: canary is valid!
```
