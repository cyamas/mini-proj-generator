package main

import "testing"

var testMap1 = map[string][]string{
	"cat1": {"feat1", "feat2"},
	"cat2": {"feat3", "feat4"},
	"cat3": {"feat5", "feat6"},
}

var testMap2 = make(map[string][]string)

func TestGetCategoryNames(t *testing.T) {
	result1 := getCategoryNames(&testMap1)
	ans1 := map[string]bool{"cat1": true, "cat2": true, "cat3": true}
	if len(result1) != len(ans1) {
		t.Errorf("length of result1 was %d; want %d", len(result1), len(ans1))
	}
	for i := range len(result1) {
		if !ans1[result1[i]] {
			t.Errorf("%v not in ans1 set", result1[i])
		}
	}

	result2 := getCategoryNames(&testMap2)
	ans2 := []string{}
	if len(result2) != 0 {
		t.Errorf("length of result2 was %d; want %d", len(result2), len(ans2))
	}
}

func TestSelectRandoFeature(t *testing.T) {
	copyTestMap1 := make(map[string][]string)
	for k, v := range testMap1 {
		copyTestMap1[k] = v
	}
	res1Cat, res1Ft := selectRandoFeature(testMap1)

	isCat, isFeat := isValidKeyValuePair(res1Cat, res1Ft, copyTestMap1)
	if !isCat {
		t.Errorf("%v is not a valid category in original testMap1", res1Cat)
	}
	if !isFeat {
		t.Errorf("%v is not a valid feature in category %v", res1Ft, res1Cat)
	}

	if !categoryIsDeleted(res1Cat, testMap1) {
		t.Errorf("%v should no longer exist in testMap1", res1Cat)
	}
}

func isValidKeyValuePair(resCat string, resFt string, ogMap map[string][]string) (bool, bool) {
	isCat := false
	isFeat := false
	for cat, feats := range ogMap {
		if resCat == cat {
			isCat = true
			for _, feat := range feats {
				if resFt == feat {
					isFeat = true
				}
			}
		}
	}
	return isCat, isFeat
}

func categoryIsDeleted(resCat string, testMap map[string][]string) bool {
	cats := make(map[string]bool)
	for cat := range testMap {
		cats[cat] = true
	}
	return !cats[resCat]
}
