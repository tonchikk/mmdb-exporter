package main

import (
  "fmt"
  "log"
  "flag"
//  "net/netip"

  "github.com/oschwald/maxminddb-golang/v2"
)

func main() {

  var country = flag.String("c", "any", "Contry code to print")
  var dbpath = flag.String("p", "GeoLite2-City.mmdb", "path to databse")

  flag.Parse();

  db, err := maxminddb.Open(*dbpath)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()


  for result := range db.Networks() {
    var record struct {
      Country struct {
      ISOCode string `maxminddb:"iso_code"`
      } `maxminddb:"country"`
    }
    err := result.Decode(&record)
    if err != nil {
    log.Fatal(err)
    }
    if *country == "any" {
      fmt.Printf("%s %s\n", record.Country.ISOCode, result.Prefix())
    } else if record.Country.ISOCode == *country {
      fmt.Printf("%s\n", result.Prefix())
    }
  }
}
