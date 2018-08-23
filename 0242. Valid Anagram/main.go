package main

import (
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func Runes(a []rune) { sort.Sort(RuneSlice(a)) }

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	runesS := []rune(s)
	runesT := []rune(t)
	Runes(runesS)
	Runes(runesT)
	i := 0
	for i < len(runesS){
		if runesT[i] != runesS[i]{
			return false
		}
		i++
	}
	return true
}
