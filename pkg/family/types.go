package family

type TreeComponentConfig struct {
	FamilyName string `json:"familyName"`
}

type FamilyConfig struct {
	Parents  []string `json:"parents,omitempty"`
	Children []string `json:"children,omitempty"`
}
