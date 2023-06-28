package main

import (
	"log"
	"strings"

	"github.com/packethost/packngo"
)

func main() {
	c, err := packngo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ps, _, err := c.Projects.List(nil)
	if err != nil {
		log.Fatal(err)
	}
	pids := []string{}
	for _, p := range ps {
		if strings.HasPrefix(p.Name, "PACKNGO_TEST_DELME_2d768716_") ||
			strings.HasPrefix(p.Name, "foobar-") ||
			strings.HasPrefix(p.Name, "My project") ||
			strings.HasPrefix(p.Name, "ansible-inte") ||
			strings.HasPrefix(p.Name, "ansible-test") ||
			strings.HasPrefix(p.Name, "ff") ||
			strings.HasPrefix(p.Name, "tftest") ||
			strings.HasPrefix(p.Name, "tfacc") ||
			strings.HasPrefix(p.Name, "jrpq6f7n") ||
			strings.HasPrefix(p.Name, "TerraformTestProject-") {
			log.Println(p.Name)
			pids = append(pids, p.ID)
		}
	}
	for _, pid := range pids {
		ds, _, err := c.Devices.List(pid, nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range ds {
			log.Printf("removing dev %s", d.ID)
			_, err = c.Devices.Delete(d.ID, false)
			if err != nil {
				log.Println("ERR, proj", pid, err)
			}
		}

	}
	for _, pid := range pids {
		log.Printf("removing project %s", pid)
		_, err = c.Projects.Delete(pid)
		if err != nil {
			log.Println("ERR, proj", pid, err.Error())
		}

	}

	log.Println(pids)
}
