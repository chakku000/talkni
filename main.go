package main

import (
	"os"
	"log"
	"net"
	"net/http"
	"html/template"
)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		hostname := get_hostname()
		ip_addrs := get_ip()
		log.Println("Hostname : ", hostname)

		type TemplateArgs struct{
			Hostname string
			Ips map[string][]string
		}
		targs := TemplateArgs{
			Hostname : hostname,
			Ips : ip_addrs,
		}
		tpl := template.Must(template.ParseFiles("index.html"))
		err := tpl.Execute(w, targs)
		if err != nil{
			log.Println(err)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}

func get_hostname() string{
	name, err := os.Hostname()
	if err != nil{
		log.Println(err)
		return "Error. Failed to get hostname."
	}
	return name
}

func get_ip() map[string][]string{
	ips := make(map[string][]string)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips
	}

	for _, iface := range ifaces {
		iname := iface.Name
		addrs, err := iface.Addrs()
		if err != nil {
			ips[iname] = nil
			continue
		}

		list := []string{}
		for _, addr := range addrs {
			list = append(list, addr.String())
		}
		ips[iname] = list
	}

	return ips
}
