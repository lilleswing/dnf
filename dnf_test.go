package dnf

import "testing"

func TestDNFIdentity(t *testing.T) {
    cases := []struct {
        in, want string
    } {
        {"A+B", "A+B"},
        //{"B+A", "A+B"}, Fails for now
        {"A", "A"},
    }
    for _, c := range cases {
        got := toDNF(c.in)
        if got != c.want {
            t.Errorf("toDNF(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
