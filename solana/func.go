package solana

import (
    "github.com/agrison/go-commons-lang/stringUtils"
    _ "github.com/agrison/go-commons-lang/stringUtils"
    "regexp"
    "strings"
)

func FindSolanaAddresses(text string) []string {
    var addresses []string
    re := regexp.MustCompile("[1-9A-HJ-NP-Za-km-z]{44}")
    matches := re.FindAllString(text, -1)
    m := make(map[string]bool)
    for _, v := range matches {
        if _, ok := m[v]; ok {
            continue
        }
        i := strings.Index(text, v)
        if i > 0 {
            front := text[i-1 : i]
            if stringUtils.IsAlphanumeric(front) {
                continue
            }
        }

        j := i + len(v)
        if j < len(text) {
            back := text[j : j+1]
            if stringUtils.IsAlphanumeric(back) {
                continue
            }
        }

        m[v] = true
        addresses = append(addresses, v)
    }

    return addresses
}
