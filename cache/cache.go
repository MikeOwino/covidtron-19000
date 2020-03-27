/*
 * Covidtron-19000 - a bot for monitoring data about COVID-19.
 * Copyright (C) 2020 Nicolò Santamaria.
 *
 * Covidtron-19000 is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Covidtron-19000 is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package cache

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Cache struct {
	botName string
	Sessions []int64 `json:"Sessions"`
}

func NewCache(bname string) *Cache {
	var cache = &Cache{botName: bname}

	fpath := fmt.Sprintf("%s/.cache/%s.json", os.Getenv("HOME"), bname)
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Println(err)
		goto exit
	}

	err = json.Unmarshal(data, cache)
	if err != nil {
		log.Println(err)
	}


exit:
	return cache
}

func (c Cache) isin(s int64) bool {
	for _, v := range c.Sessions {
		if s == v {
			return true
		}
	}

	return false
}

func (c *Cache) SaveSession(s int64) {
	if c.isin(s) {
		c.Sessions = append(c.Sessions, s)

		b, err := json.Marshal(c)
		if err != nil {
			log.Println(err)
			return
		}

		fpath := fmt.Sprintf("%s/.cache/%s.json", os.Getenv("HOME"), c.botName)
		err = ioutil.WriteFile(fpath, b, 0755)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *Cache) DelSession(s int64) {
	for k, v := range c.Sessions {
		if v == s {
			c.Sessions = append(c.Sessions[:k], c.Sessions[k+1:]...)
			break
		}
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return
	}

	fpath := fmt.Sprintf("%s/.cache/%s.json", os.Getenv("HOME"), c.botName)
	err = ioutil.WriteFile(fpath, b, 0755)
	if err != nil {
		log.Println(err)
	}
}

func (c Cache) GetSessions() []int64 {
	return c.Sessions
}
