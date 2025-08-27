# mmdb-exporter
.mmdb file exporter
## Usage
```
$ ./mmdb-exporter -h
Usage of ./mmdb-exporter:
  -c string
        Contry code to print (default "any")
  -p string
        path to databse (default "GeoLite2-City.mmdb")
```
### Example
Note: with -c no country code in output, only entries
```
$ ./mmdb-exporter -p ../GeoLite2-Country.mmdb -c US | head
1.178.4.0/23
1.178.6.0/24
1.178.8.0/23
1.178.65.0/24
1.178.68.0/22

$ ./mmdb-exporter -p ../GeoLite2-Country.mmdb | head
 1.0.0.0/24
CN 1.0.1.0/24
CN 1.0.2.0/23
AU 1.0.4.0/22
CN 1.0.8.0/21
JP 1.0.16.0/20
```
Country code is a first field, so to speedupt filtering you cpoud use regexp in grep with `^` like `./mmdb-exporter -p ../GeoLite2-Country.mmdb | grep -E '^(US|CA) '`
