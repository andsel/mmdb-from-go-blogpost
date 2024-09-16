package main

import (
	"log"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

func main() {
	// Load the database we wish to enrich.
	writer, err := mmdbwriter.Load("/path/to/logstash_plugins/logstash-filter-geoip/src/test/resources/maxmind-test-data/GeoIP2-Country-Test.mmdb", mmdbwriter.Options{})
	if err != nil {
		log.Fatal(err)
	}

    // Modify existing IP
	_, singleNet, err := net.ParseCIDR("216.160.83.58/32")
	if err != nil {
		log.Fatal(err)
	}
	singleHostData := mmdbtype.Map{
		//"is_in_european_union": mmdbtype.Uint16(1),
		"country": mmdbtype.Map{
			"geoname_id": mmdbtype.Uint64(2750405),
			"is_in_european_union": mmdbtype.Uint16(1),
			"iso_code": mmdbtype.String("NL"),
			"names": mmdbtype.Map{
				"en": mmdbtype.String("Netherlands"),
			},
		},
	}
	if err := writer.InsertFunc(singleNet, inserter.TopLevelMergeWith(singleHostData)); err != nil {
		log.Fatal(err)
	}


	// Write the newly enriched DB to the filesystem.
	fh, err := os.Create("/path/to/logstash_plugins/logstash-filter-geoip/src/test/resources/maxmind-test-data/GeoIP2-Country-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}
}
