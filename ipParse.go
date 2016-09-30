// ipParse project ipParse.go
package ipParse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseIP(IP_RANGE string) []string {
	var ipList []string
	//Проверка IP регулярному выражению
	match, _ := regexp.MatchString("([0-9]{1,3}-?[0-9]{1,3}|[0-9]{1,3}).([0-9]{1,3}-?[0-9]{1,3}|[0-9]{1,3}).([0-9]{1,3}-?[0-9]{1,3}|[0-9]{1,3}).([0-9]{1,3}-?[0-9]{1,3}|[0-9]{1,3})", IP_RANGE)
	fmt.Println(match)
	rng := strings.Split(IP_RANGE, ".")
	if len(rng) != 4 || !match {
		return ipList
	}
	d := make(map[string]map[int]int)
	mask := ""
	for i, item := range rng {
		d[fmt.Sprintf("o%d", i+1)] = map[int]int{0: 0, 1: 0}
		if strings.Contains(item, "-") {
			tmp := strings.Split(item, "-")
			//
			for y, el := range tmp {
				el, _ := strconv.Atoi(el)
				//TODO Добавить обработку ошибок, если нужна будет
				d[fmt.Sprintf("o%d", i+1)][y] = el
			}
		} else {
			el, _ := strconv.Atoi(item)
			d[fmt.Sprintf("o%d", i+1)][0] = el
			d[fmt.Sprintf("o%d", i+1)][1] = el
		}
		//Генерим маску
		if len(mask) > 0 {
			if mask[len(mask)-1] == '.' {
				mask += "%d"
			} else {
				mask += ".%d"
			}
		} else {
			mask += "%d."
		}
	}
	//Генерим список адресов
	for o1 := d["o1"][0]; o1 <= d["o1"][1]; o1++ {
		for o2 := d["o2"][0]; o2 <= d["o2"][1]; o2++ {
			for o3 := d["o3"][0]; o3 <= d["o3"][1]; o3++ {
				for o4 := d["o4"][0]; o4 <= d["o4"][1]; o4++ {
					ipList = append(ipList, fmt.Sprintf(mask, o1, o2, o3, o4))
				}
			}
		}
	}

	return ipList
}
