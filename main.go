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

	ps, _, err := c.Projects.List()
	if err != nil {
		log.Fatal(err)
	}
	pids := []string{}
	for _, p := range ps {
		if strings.HasPrefix(p.Name, "PACKNGO_TEST_DELME_2d768716_") {
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
			log.Println("removing dev %s", d.ID)
			_, err = c.Devices.Delete(d.ID)
			if err != nil {
				log.Fatal(err)
			}
		}

		vs, _, err := c.Volumes.List(pid, nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range vs {
			log.Println("removing vol %s", v.ID)
			if v.Locked {
				_, err = c.Volumes.Unlock(v.ID)
				if err != nil {
					log.Fatal(err)
				}
			}
			_, err = c.Volumes.Delete(v.ID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	for _, pid := range pids {
		log.Println("removing project %s", pid)
		_, err = c.Projects.Delete(pid)
		if err != nil {
			log.Fatal(err)
		}

	}

	log.Println(pids)
}
