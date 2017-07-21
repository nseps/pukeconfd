package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/openstack", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requesting openstack metadata service", r.RemoteAddr)
		fmt.Fprintln(w, "latest")
	})

	http.HandleFunc("/openstack/latest/meta_data.json", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requesting meta_data.json", r.RemoteAddr)

		meta := MetaData{
			UUID:     "dc6302fc-6db0-11e7-882e-e30d38b305d3",
			Hostname: "mahhostnm",
		}

		metaBytes, err := json.Marshal(meta)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintln(w, string(metaBytes))
	})

	http.HandleFunc("/openstack/latest/network_data.json", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requesting network_data.json", r.RemoteAddr)
		fmt.Fprintf(w, "{}")
	})

	http.HandleFunc("/openstack/latest/vendor_data.json", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requesting vendor_data.json", r.RemoteAddr)
		fmt.Fprintf(w, "{}")
	})

	http.HandleFunc("/openstack/latest/user_data", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requesting user_data", r.RemoteAddr)

		userdata, err := ioutil.ReadFile("user-data.yaml")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintln(w, string(userdata))
	})

	log.Println(http.ListenAndServe("169.254.169.254:80", nil))
}
