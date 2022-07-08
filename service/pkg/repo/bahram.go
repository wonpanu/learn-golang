package repo

import (
	"strconv"
	"strings"
)

type BahRamRepo interface {
	BahRam(n string) ([]string, error)
}

type BahRamAdapter []string

func (b BahRamAdapter) BahRam(n string) ([]string, error) {
	result := []string{}
	if num, err := strconv.Atoi(n); err == nil {
		for count := 1; count <= num; count++ {
			if count%3 == 0 || count%5 == 0 {
				result = append(result, "Zzz")
			} else {
				s := ""
				s += strings.Repeat("Bah", count%3)
				s += strings.Repeat("Ram", count%5)
				result = append(result, s)
			}
		}
	}
	return result, nil
}

func BahRam(n string) ([]string, error) {
	result := []string{}
	if num, err := strconv.Atoi(n); err == nil {
		for count := 1; count <= num; count++ {
			if count%3 == 0 || count%5 == 0 {
				result = append(result, "Zzz")
			} else {
				s := ""
				s += strings.Repeat("Bah", count%3)
				s += strings.Repeat("Ram", count%5)
				result = append(result, s)
			}
		}
	}
	return result, nil
}

func NewBahRamRepo() BahRamRepo {
	return BahRamAdapter{}
}
