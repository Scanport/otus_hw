package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Top10(str string) []string {
	str = strings.ToLower(str)
	re := regexp.MustCompile(`(?i)[\p{L}\p{Mn}]+(?:-\p{L}+)?`)
	parts := re.FindAllString(str, -1)
	top := sortMapByValue(mapGenerator(parts))
	if len(top) > 10 {
		return top[:10]
	}
	return top
}

func mapGenerator(sl []string) map[string]int {
	partMap := make(map[string]int)
	for _, el := range sl {
		partMap[el]++
	}
	return partMap
}

func sortMapByValue(m map[string]int) []string {
	pairs := make([][2]string, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]string{k, strconv.Itoa(v)})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] > pairs[j][1]
	})
	result := make([]string, len(pairs))
	for i, pair := range pairs {
		result[i] = pair[0]
	}
	return result
}
