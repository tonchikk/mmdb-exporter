# mmdb-exporter
.mmdb file exporter
## Build
Just `go build`
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

## Some links
* https://github.com/recall704/geoip2ipset/blob/main/main.go
* https://github.com/chr0mag/geoipsets/blob/main/bash/build-country-sets.sh
* https://habr.com/ru/articles/721070/ - nginx + geodb
* https://stat.ripe.net/data/country-resource-list/data.json?v4_format=prefix&resource=ru - ripe db query for reginal networks
* https://github.com/mivk/ip-country/blob/master/get-ripe-ips
* https://unix.stackexchange.com/questions/502907/how-to-block-countries-iptables-or-firewalld-by-geolite2-mmdb
* https://github.com/knoguchi/mmdb2csv
* https://github.com/P3TERX/GeoLite.mmdb
* https://www.ipdeny.com/ipblocks/
* https://www.cloudflare.com/ips-v4/#
* http://nginx.org/en/docs/http/ngx_http_geoip_module.html
* 
