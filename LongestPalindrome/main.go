package main

import ( 
	"strings"
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	chars := strings.Split(s, "")
	pa := []string{}

	for s := range subStrings(chars, len(chars)) {
		if s == reverse(s) {
			pa = append(pa, s)
		}
	}

	var l string
	var ll int
	for _, p := range pa {
		if len(p) > ll {
			l = p
			ll = len(p)
		}
	}

	return l
}

func subStrings(chars []string, n int) map[string]int {
	m := make(map[string]int)

	for len := 1; len <= n; len++ {
		for i := 0; i <= n - len; i++ {
			j := i + len - 1;
			s := ""
			for k := i; k <= j; k++ {
				s+=chars[k]
			}

			if _, ok := m[s]; !ok {
				m[s] = 1; 
			}
		}
	}

	return m
}

func reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}