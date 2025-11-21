package codelocationmatch

import (
	"reflect"
	"regexp"
)

func toMap(s []*Location) map[string]map[int]*Location {
	ret := make(map[string]map[int]*Location)
	for _, loc := range s {
		if ret[loc.filePath] == nil {
			ret[loc.filePath] = make(map[int]*Location)
		}
		ret[loc.filePath][loc.lineNo] = loc
	}
	return ret
}

func allKeys[T string | int](s ...any) []T {
	keyMap := make(map[T]bool)
	keys := make([]T, 0)
	for _, m := range s {
		values := reflect.ValueOf(m).MapKeys()
		for _, value := range values {
			v := value.Interface().(T)
			if !keyMap[v] {
				keyMap[v] = true
				keys = append(keys, v)
			}
		}
	}
	return keys
}

func MatchLocations(loc1, loc2 []*Location, ignoreLineNo bool) bool {
	m1 := toMap(loc1)
	m2 := toMap(loc2)
	keys := allKeys[string](m1, m2)
	for _, k := range keys {
		mm1, mm2 := m1[k], m2[k]
		if ignoreLineNo {
			if mm1 != nil && mm2 != nil {
				return true
			}
		} else {
			keys2 := allKeys[int](mm1, mm2)
			for _, kk := range keys2 {
				l1, l2 := mm1[kk], mm2[kk]
				if l1 != nil && l2 != nil {
					return true
				}
			}
		}
	}
	return false
}

func Match(s1, s2 string, ignoreLineNo bool, filePathPattern *regexp.Regexp) bool {
	loc1 := ParseLocations(s1, filePathPattern)
	loc2 := ParseLocations(s2, filePathPattern)
	return MatchLocations(loc1, loc2, ignoreLineNo)
}
