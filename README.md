# cert-check
A simple Go TLS certificate checker

#### Usage:
```
$ ./cert-check 
Usage:
  -d="google.com": domain name to alert on
  -n=30: Number of days to alert on. Only print if this is triggered.
  -v=false: always print number of days until expiration
```

#### Check a domain, with default options:
```$ ./cert-check -d google.com```

#### Check a domain in verbose-mode, with the default number of days set
```
$ ./cert-check -d google.com -v
domain: google.com, issued by: [GeoTrust Inc.] expires in: 89 days
```

#### Check a domain and alert if the certificate will expire in less than 100 days
```
$ ./cert-check -d google.com -v -n 100
domain: google.com, issued by: [GeoTrust Inc.] expires in: 89 days
$ echo $?
1
```
