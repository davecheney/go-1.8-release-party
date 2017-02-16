package main

import (
	"sort"
	"strings"
)

func main() {
	switch strings.ToLower(field) {
	case "updated":
		sort.Slice(subs, func(i, j int) bool { return subs[i].Updated.After(subs[j].Updated) })
	case "trust":
		sort.Slice(subs, func(i, j int) bool { return subs[i].Trust > subs[j].Trust })
	case "stddev":
		sort.Slice(subs, func(i, j int) bool {
			var s1, s2 stats.Sample
			for _, r := range subs[i].Ratings {
				s1.Xs = append(s1.Xs, float64(r.Value))
			}
			for _, r := range subs[j].Ratings {
				s2.Xs = append(s2.Xs, float64(r.Value))
			}
			return s1.StdDev() < s2.StdDev()
		})
	case "rating":
		fallthrough
	default:
		sort.Slice(subs, func(i, j int) bool { return subs[i].Rating > subs[j].Rating })

	}
}
