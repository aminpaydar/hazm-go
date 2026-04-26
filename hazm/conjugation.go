package hazm

import "strings"

// Conjugation mirrors hazm.lemmatizer.Conjugation (Python).
type Conjugation struct{}

func zipSpace(a, b []string) []string {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = a[i] + " " + b[i]
	}
	return out
}

func prefixEach(prefix string, xs []string) []string {
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = prefix + x
	}
	return out
}

func joinPrefix(prefix string, xs []string) []string {
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = prefix + x
	}
	return out
}

var (
	pastSuffixes           = []string{"م", "ی", "", "یم", "ید", "ند"}
	presentPerfectSuffixes = []string{"ه‌ام", "ه‌ای", "ه است", "ه", "ه‌ایم", "ه‌اید", "ه‌اند"}
	presentSuffixes        = []string{"م", "ی", "د", "یم", "ید", "ند"}
)

func (*Conjugation) PerfectivePast(ri string) []string {
	out := make([]string, len(pastSuffixes))
	for i, s := range pastSuffixes {
		out[i] = ri + s
	}
	return out
}

func (c *Conjugation) NegativePerfectivePast(ri string) []string {
	return prefixEach("ن", c.PerfectivePast(ri))
}

func (c *Conjugation) PassivePerfectivePast(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectivePast("شد"))
}

func (c *Conjugation) NegativePassivePerfectivePast(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePerfectivePast("شد"))
}

func (c *Conjugation) ImperfectivePast(ri string) []string {
	return prefixEach("می‌", c.PerfectivePast(ri))
}

func (c *Conjugation) NegativeImperfectivePast(ri string) []string {
	return prefixEach("ن", c.ImperfectivePast(ri))
}

func (c *Conjugation) PassiveImperfectivePast(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectivePast("شد"))
}

func (c *Conjugation) NegativePassiveImperfectivePast(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectivePast("شد"))
}

func (c *Conjugation) PastProgresive(ri string) []string {
	return zipSpace(c.PerfectivePast("داشت"), c.ImperfectivePast(ri))
}

func (c *Conjugation) PassivePastProgresive(ri string) []string {
	return zipSpace(c.PerfectivePast("داشت"), c.PassiveImperfectivePast(ri))
}

func (*Conjugation) PresentPerfect(ri string) []string {
	out := make([]string, len(presentPerfectSuffixes))
	for i, s := range presentPerfectSuffixes {
		out[i] = ri + s
	}
	return out
}

func (c *Conjugation) NegativePresentPerfect(ri string) []string {
	return prefixEach("ن", c.PresentPerfect(ri))
}

func (c *Conjugation) SubjunctivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectivePresent("باش"))
}

func (c *Conjugation) NegativeSubjunctivePresentPerfect(ri string) []string {
	return prefixEach("ن", c.SubjunctivePresentPerfect(ri))
}

func (c *Conjugation) GrammaticalPresentPerfect(ri string) []string {
	xs := c.PerfectivePresent("باش")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "باشی" {
			out[i] = ri + "ه باش"
		} else {
			out[i] = ri + "ه " + x
		}
	}
	return out
}

func (c *Conjugation) NegativeGrammaticalPresentPerfect(ri string) []string {
	xs := c.PerfectivePresent("باش")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "باشی" {
			out[i] = "ن" + ri + "ه باش"
		} else {
			out[i] = "ن" + ri + "ه " + x
		}
	}
	return out
}

func (c *Conjugation) PassivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.PresentPerfect("شد"))
}

func (c *Conjugation) NegativePassivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePresentPerfect("شد"))
}

func (c *Conjugation) PassiveSubjunctivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.SubjunctivePresentPerfect("شد"))
}

func (c *Conjugation) NegativePassiveSubjunctivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeSubjunctivePresentPerfect("شد"))
}

func (c *Conjugation) PassiveGrammaticalPresentPerfect(ri string) []string {
	xs := c.PerfectivePresent("باش")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "باشی" {
			out[i] = ri + "ه شده باش"
		} else {
			out[i] = ri + "ه شده " + x
		}
	}
	return out
}

func (c *Conjugation) NegativePassiveGrammaticalPresentPerfect(ri string) []string {
	xs := c.PerfectivePresent("باش")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "باشی" {
			out[i] = ri + "ه نشده باش"
		} else {
			out[i] = ri + "ه نشده " + x
		}
	}
	return out
}

func (c *Conjugation) ImperfectivePresentPerfect(ri string) []string {
	return prefixEach("می‌", c.PresentPerfect(ri))
}

func (c *Conjugation) NegativeImperfectivePresentPerfect(ri string) []string {
	return prefixEach("ن", c.ImperfectivePresentPerfect(ri))
}

func (c *Conjugation) SubjunctiveImperfectivePresentPerfect(ri string) []string {
	return prefixEach("می‌", c.SubjunctivePresentPerfect(ri))
}

func (c *Conjugation) NegativeSubjunctiveImperfectivePresentPerfect(ri string) []string {
	return prefixEach("ن", c.SubjunctiveImperfectivePresentPerfect(ri))
}

func (c *Conjugation) PassiveImperfectivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectivePresentPerfect("شد"))
}

func (c *Conjugation) NegativePassiveImperfectivePresentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectivePresentPerfect("شد"))
}

func (c *Conjugation) PassiveSubjunctiveImperfectivePresentPerfect(ri string) []string {
	xs := c.SubjunctiveImperfectivePresentPerfect("شد")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = ri + "ه " + x
	}
	return out
}

func (c *Conjugation) NegativePassiveSubjunctiveImperfectivePresentPerfect(ri string) []string {
	xs := c.NegativeSubjunctiveImperfectivePresentPerfect("شد")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = ri + "ه " + x
	}
	return out
}

func (c *Conjugation) PresentPerfectProgressive(ri string) []string {
	return zipSpace(c.PresentPerfect("داشت"), c.ImperfectivePresentPerfect(ri))
}

func (c *Conjugation) PassivePresentPerfectProgressive(ri string) []string {
	return zipSpace(c.PresentPerfect("داشت"), c.PassiveImperfectivePresentPerfect(ri))
}

func (c *Conjugation) PastPrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectivePast("بود"))
}

func (c *Conjugation) NegativePastPrecedent(ri string) []string {
	return prefixEach("ن", c.PastPrecedent(ri))
}

func (c *Conjugation) PassivePastPrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.PastPrecedent("شد"))
}

func (c *Conjugation) NegativePassivePastPrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePastPrecedent("شد"))
}

func (c *Conjugation) ImperfectivePastPrecedent(ri string) []string {
	return prefixEach("می‌", c.PastPrecedent(ri))
}

func (c *Conjugation) NegativeImperfectivePastPrecedent(ri string) []string {
	return prefixEach("ن", c.ImperfectivePastPrecedent(ri))
}

func (c *Conjugation) PassiveImperfectivePastPrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectivePastPrecedent("شد"))
}

func (c *Conjugation) NegativePassiveImperfectivePastPrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectivePastPrecedent("شد"))
}

func (c *Conjugation) PastPrecedentProgressive(ri string) []string {
	return zipSpace(c.PerfectivePast("داشت"), c.ImperfectivePastPrecedent(ri))
}

func (c *Conjugation) PassivePastPrecedentProgressive(ri string) []string {
	return zipSpace(c.PerfectivePast("داشت"), c.PassiveImperfectivePastPrecedent(ri))
}

func (c *Conjugation) PastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.PresentPerfect("بود"))
}

func (c *Conjugation) NegativePastPrecedentPerfect(ri string) []string {
	return prefixEach("ن", c.PastPrecedentPerfect(ri))
}

func (c *Conjugation) SubjunctivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.SubjunctivePresentPerfect("بود"))
}

func (c *Conjugation) NegativeSubjunctivePastPrecedentPerfect(ri string) []string {
	return prefixEach("ن", c.SubjunctivePastPrecedentPerfect(ri))
}

func (c *Conjugation) GrammaticalPastPrecedentPerfect(ri string) []string {
	xs := c.PerfectivePresent("باش")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "باشی" {
			out[i] = ri + "ه بوده باش"
		} else {
			out[i] = ri + "ه بوده " + x
		}
	}
	return out
}

func (c *Conjugation) NegativeGrammaticalPastPrecedentPerfect(ri string) []string {
	return prefixEach("ن", c.GrammaticalPastPrecedentPerfect(ri))
}

func (c *Conjugation) PassivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.PastPrecedentPerfect("شد"))
}

func (c *Conjugation) NegativePassivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePastPrecedentPerfect("شد"))
}

func (c *Conjugation) PassiveSubjunctivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.SubjunctivePastPrecedentPerfect("شد"))
}

func (c *Conjugation) NegativePassiveSubjunctivePastPrecedentPerfect(ri string) []string {
	xs := c.SubjunctivePastPrecedentPerfect("شد")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = ri + "ه ن" + x
	}
	return out
}

func (c *Conjugation) PassiveGrammaticalPastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.GrammaticalPastPrecedentPerfect("شد"))
}

func (c *Conjugation) NegativePassiveGrammaticalPastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeGrammaticalPastPrecedentPerfect("شد"))
}

func (c *Conjugation) ImperfectivePastPrecedentPerfect(ri string) []string {
	return prefixEach("می‌", c.PastPrecedentPerfect(ri))
}

func (c *Conjugation) NegativeImperfectivePastPrecedentPerfect(ri string) []string {
	return prefixEach("ن", c.ImperfectivePastPrecedentPerfect(ri))
}

func (c *Conjugation) SubjunctiveImperfectivePastPrecedentPerfect(ri string) []string {
	return prefixEach("می‌", c.SubjunctivePastPrecedentPerfect(ri))
}

func (c *Conjugation) NegativeSubjunctiveImperfectivePastPrecedentPerfect(ri string) []string {
	return prefixEach("ن", c.SubjunctiveImperfectivePastPrecedentPerfect(ri))
}

func (c *Conjugation) PassiveImperfectivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectivePastPrecedentPerfect("شد"))
}

func (c *Conjugation) NegativePassiveImperfectivePastPrecedentPerfect(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectivePastPrecedentPerfect("شد"))
}

func (c *Conjugation) PassiveSubjunctiveImperfectivePastPrecedentPerfect(ri string) []string {
	xs := c.SubjunctiveImperfectivePastPrecedentPerfect("شد")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = ri + "ه " + x
	}
	return out
}

func (c *Conjugation) NegativePassiveSubjunctiveImperfectivePastPrecedentPerfect(ri string) []string {
	xs := c.SubjunctiveImperfectivePastPrecedentPerfect("شد")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = ri + "ه ن" + x
	}
	return out
}

func (c *Conjugation) PastPrecedentPerfectProgressive(ri string) []string {
	return zipSpace(c.PresentPerfect("داشت"), c.ImperfectivePastPrecedentPerfect(ri))
}

func (c *Conjugation) PassivePastPrecedentPerfectProgressive(ri string) []string {
	return zipSpace(c.PresentPerfect("داشت"), c.PassiveImperfectivePastPrecedentPerfect(ri))
}

func (*Conjugation) PerfectivePresent(rii string) []string {
	out := make([]string, len(presentSuffixes))
	for i, s := range presentSuffixes {
		out[i] = rii + s
	}
	return out
}

func (c *Conjugation) NegativePerfectivePresent(rii string) []string {
	return prefixEach("ن", c.PerfectivePresent(rii))
}

func (c *Conjugation) SubjunctivePerfectivePresent(rii string) []string {
	return prefixEach("ب", c.PerfectivePresent(rii))
}

func (c *Conjugation) NegativeSubjunctivePerfectivePresent(rii string) []string {
	return prefixEach("ن", c.PerfectivePresent(rii))
}

func (c *Conjugation) GrammaticalPerfectivePresent(rii string) []string {
	xs := c.SubjunctivePerfectivePresent(rii)
	out := make([]string, len(xs))
	for i, x := range xs {
		// Matches Python: only substitutes the docstring example forms.
		if x == "ببینی" {
			out[i] = "ببین"
		} else {
			out[i] = x
		}
	}
	return out
}

func (c *Conjugation) NegativeGrammaticalPerfectivePresent(rii string) []string {
	xs := c.PerfectivePresent(rii)
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "بینی" {
			out[i] = "نبین"
		} else {
			out[i] = "ن" + x
		}
	}
	return out
}

func (c *Conjugation) PassivePerfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectivePresent("شو"))
}

func (c *Conjugation) NegativePassivePerfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePerfectivePresent("شو"))
}

func (c *Conjugation) PassiveSubjunctivePerfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.SubjunctivePerfectivePresent("شو"))
}

func (c *Conjugation) NegativePassiveSubjunctivePerfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeSubjunctivePerfectivePresent("شو"))
}

func (c *Conjugation) PassiveGrammaticalPerfectivePresent(ri string) []string {
	xs := c.GrammaticalPerfectivePresent("شو")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "بشوی" {
			out[i] = ri + "ه بشو"
		} else {
			out[i] = ri + "ه " + x
		}
	}
	return out
}

func (c *Conjugation) NegativePassiveGrammaticalPerfectivePresent(ri string) []string {
	xs := c.NegativeGrammaticalPerfectivePresent("شو")
	out := make([]string, len(xs))
	for i, x := range xs {
		if x == "نشوی" {
			out[i] = ri + "ه نشو"
		} else {
			out[i] = ri + "ه " + x
		}
	}
	return out
}

func (c *Conjugation) ImperfectivePresent(rii string) []string {
	return prefixEach("می‌", c.PerfectivePresent(rii))
}

func (c *Conjugation) NegativeImperfectivePresent(rii string) []string {
	return prefixEach("ن", c.ImperfectivePresent(rii))
}

func (c *Conjugation) PassiveImperfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectivePresent("شو"))
}

func (c *Conjugation) NegativePassiveImperfectivePresent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectivePresent("شو"))
}

func (c *Conjugation) PresentProgressive(rii string) []string {
	return zipSpace(c.PerfectivePresent("دار"), c.ImperfectivePresent(rii))
}

func (c *Conjugation) PassivePresentProgressive(ri string) []string {
	return zipSpace(c.PerfectivePresent("دار"), c.PassiveImperfectivePresent(ri))
}

func (c *Conjugation) PerfectiveFuture(ri string) []string {
	xs := c.PerfectivePresent("خواه")
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = x + " " + ri
	}
	return out
}

func (c *Conjugation) NegativePerfectiveFuture(ri string) []string {
	return prefixEach("ن", c.PerfectiveFuture(ri))
}

func (c *Conjugation) PassivePerfectiveFuture(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectiveFuture("شد"))
}

func (c *Conjugation) NegativePassivePerfectiveFuture(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativePerfectiveFuture("شد"))
}

func (c *Conjugation) ImperfectiveFuture(ri string) []string {
	return prefixEach("می‌", c.PerfectiveFuture(ri))
}

func (c *Conjugation) NegativeImperfectiveFuture(ri string) []string {
	return prefixEach("ن", c.ImperfectiveFuture(ri))
}

func (c *Conjugation) PassiveImperfectiveFuture(ri string) []string {
	return joinPrefix(ri+"ه ", c.ImperfectiveFuture("شد"))
}

func (c *Conjugation) NegativePassiveImperfectiveFuture(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeImperfectiveFuture("شد"))
}

func (c *Conjugation) FuturePrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.PerfectiveFuture("بود"))
}

func (c *Conjugation) NegativeFuturePrecedent(ri string) []string {
	return prefixEach("ن", c.FuturePrecedent(ri))
}

func (c *Conjugation) PassiveFuturePrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.FuturePrecedent("شد"))
}

func (c *Conjugation) NegativePassiveFuturePrecedent(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeFuturePrecedent("شد"))
}

func (c *Conjugation) FuturePrecedentImperfective(ri string) []string {
	return prefixEach("می‌", c.FuturePrecedent(ri))
}

func (c *Conjugation) NegativeFuturePrecedentImperfective(ri string) []string {
	return prefixEach("ن", c.FuturePrecedentImperfective(ri))
}

func (c *Conjugation) PassiveFuturePrecedentImperfective(ri string) []string {
	return joinPrefix(ri+"ه ", c.FuturePrecedentImperfective("شد"))
}

func (c *Conjugation) NegativePassiveFuturePrecedentImperfective(ri string) []string {
	return joinPrefix(ri+"ه ", c.NegativeFuturePrecedentImperfective("شد"))
}

// GetAll returns the same flattened list as Conjugation.get_all in Python.
func (c *Conjugation) GetAll(verb string) []string {
	parts := strings.SplitN(verb, "#", 2)
	if len(parts) != 2 {
		return nil
	}
	ri, rii := parts[0], parts[1]

	var blocks [][]string
	blocks = append(blocks, []string{ri + "ن"})
	blocks = append(blocks, c.PerfectivePast(ri))
	blocks = append(blocks, c.NegativePerfectivePast(ri))
	blocks = append(blocks, c.PassivePerfectivePast(ri))
	blocks = append(blocks, c.NegativePassivePerfectivePast(ri))
	blocks = append(blocks, c.ImperfectivePast(ri))
	blocks = append(blocks, c.NegativeImperfectivePast(ri))
	blocks = append(blocks, c.PassiveImperfectivePast(ri))
	blocks = append(blocks, c.NegativePassiveImperfectivePast(ri))
	blocks = append(blocks, c.PastProgresive(ri))
	blocks = append(blocks, c.PassivePastProgresive(ri))
	blocks = append(blocks, c.PresentPerfect(ri))
	blocks = append(blocks, c.NegativePresentPerfect(ri))
	blocks = append(blocks, c.SubjunctivePresentPerfect(ri))
	blocks = append(blocks, c.NegativeSubjunctivePresentPerfect(ri))
	blocks = append(blocks, c.GrammaticalPresentPerfect(ri))
	blocks = append(blocks, c.NegativeGrammaticalPresentPerfect(ri))
	blocks = append(blocks, c.PassivePresentPerfect(ri))
	blocks = append(blocks, c.NegativePassivePresentPerfect(ri))
	blocks = append(blocks, c.PassiveSubjunctivePresentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveSubjunctivePresentPerfect(ri))
	blocks = append(blocks, c.PassiveGrammaticalPresentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveGrammaticalPresentPerfect(ri))
	blocks = append(blocks, c.ImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.NegativeImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.SubjunctiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.NegativeSubjunctiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.PassiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.PassiveSubjunctiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveSubjunctiveImperfectivePresentPerfect(ri))
	blocks = append(blocks, c.PresentPerfectProgressive(ri))
	blocks = append(blocks, c.PassivePresentPerfectProgressive(ri))
	blocks = append(blocks, c.PastPrecedent(ri))
	blocks = append(blocks, c.NegativePastPrecedent(ri))
	blocks = append(blocks, c.PassivePastPrecedent(ri))
	blocks = append(blocks, c.NegativePassivePastPrecedent(ri))
	blocks = append(blocks, c.ImperfectivePastPrecedent(ri))
	blocks = append(blocks, c.NegativeImperfectivePastPrecedent(ri))
	blocks = append(blocks, c.PassiveImperfectivePastPrecedent(ri))
	blocks = append(blocks, c.NegativePassiveImperfectivePastPrecedent(ri))
	blocks = append(blocks, c.PastPrecedentProgressive(ri))
	blocks = append(blocks, c.PassivePastPrecedentProgressive(ri))
	blocks = append(blocks, c.PastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePastPrecedentPerfect(ri))
	blocks = append(blocks, c.SubjunctivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativeSubjunctivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.GrammaticalPastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativeGrammaticalPastPrecedentPerfect(ri))
	blocks = append(blocks, c.PassivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePassivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.PassiveSubjunctivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveSubjunctivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.PassiveGrammaticalPastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveGrammaticalPastPrecedentPerfect(ri))
	blocks = append(blocks, c.ImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativeImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.SubjunctiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativeSubjunctiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.PassiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.PassiveSubjunctiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.NegativePassiveSubjunctiveImperfectivePastPrecedentPerfect(ri))
	blocks = append(blocks, c.PastPrecedentPerfectProgressive(ri))
	blocks = append(blocks, c.PassivePastPrecedentPerfectProgressive(ri))
	blocks = append(blocks, c.PerfectivePresent(rii))
	blocks = append(blocks, c.NegativePerfectivePresent(rii))
	blocks = append(blocks, c.SubjunctivePerfectivePresent(rii))
	blocks = append(blocks, c.NegativeSubjunctivePerfectivePresent(rii))
	blocks = append(blocks, c.GrammaticalPerfectivePresent(rii))
	blocks = append(blocks, c.NegativeGrammaticalPerfectivePresent(rii))
	blocks = append(blocks, c.PassivePerfectivePresent(ri))
	blocks = append(blocks, c.NegativePassivePerfectivePresent(ri))
	blocks = append(blocks, c.PassiveSubjunctivePerfectivePresent(ri))
	blocks = append(blocks, c.NegativePassiveSubjunctivePerfectivePresent(ri))
	blocks = append(blocks, c.PassiveGrammaticalPerfectivePresent(ri))
	blocks = append(blocks, c.NegativePassiveGrammaticalPerfectivePresent(ri))
	blocks = append(blocks, c.ImperfectivePresent(rii))
	blocks = append(blocks, c.NegativeImperfectivePresent(rii))
	blocks = append(blocks, c.PassiveImperfectivePresent(ri))
	blocks = append(blocks, c.NegativePassiveImperfectivePresent(ri))
	blocks = append(blocks, c.PresentProgressive(rii))
	blocks = append(blocks, c.PassivePresentProgressive(ri))
	blocks = append(blocks, c.PerfectiveFuture(ri))
	blocks = append(blocks, c.NegativePerfectiveFuture(ri))
	blocks = append(blocks, c.PassivePerfectiveFuture(ri))
	blocks = append(blocks, c.NegativePassivePerfectiveFuture(ri))
	blocks = append(blocks, c.ImperfectiveFuture(ri))
	blocks = append(blocks, c.NegativeImperfectiveFuture(ri))
	blocks = append(blocks, c.PassiveImperfectiveFuture(ri))
	blocks = append(blocks, c.NegativePassiveImperfectiveFuture(ri))
	blocks = append(blocks, c.FuturePrecedent(ri))
	blocks = append(blocks, c.NegativeFuturePrecedent(ri))
	blocks = append(blocks, c.PassiveFuturePrecedent(ri))
	blocks = append(blocks, c.NegativePassiveFuturePrecedent(ri))
	blocks = append(blocks, c.FuturePrecedentImperfective(ri))
	blocks = append(blocks, c.NegativeFuturePrecedentImperfective(ri))
	blocks = append(blocks, c.PassiveFuturePrecedentImperfective(ri))
	blocks = append(blocks, c.NegativePassiveFuturePrecedentImperfective(ri))

	n := 0
	for _, b := range blocks {
		n += len(b)
	}
	out := make([]string, 0, n)
	for _, b := range blocks {
		out = append(out, b...)
	}
	return out
}
