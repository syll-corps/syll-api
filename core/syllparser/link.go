package syllparser 

type SyllabLinker struct {
	// Late: cashed links for the fast uri-getting.
	template map[string]string

	core string 
} 

func (sl *SyllabLinker) MakeGroupUri(g string) string {
	return sl.core + g 
}