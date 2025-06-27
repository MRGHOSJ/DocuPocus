package ai

// Documentation represents the comprehensive documentation structure
type Documentation struct {
	Summary         string   `json:"summary"`
	Parameters      []Param  `json:"parameters"`
	Returns         string   `json:"returns"`
	TimeComplexity  string   `json:"time_complexity"`
	SpaceComplexity string   `json:"space_complexity"`
	UsageExample    string   `json:"usage_example"`
	EdgeCases       []string `json:"edge_cases"`
}

// Param describes a function parameter
type Param struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type YAMLField struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // scalar, map, array
	Description string `json:"description"`
}

type YAMLDocumentation struct {
	Summary       string                 `json:"summary"`
	Fields        []YAMLField            `json:"fields"`
	Examples      map[string]interface{} `json:"examples,omitempty"`
	Defaults      map[string]interface{} `json:"defaults,omitempty"`
	Usage         string                 `json:"usage,omitempty"`
	BestPractices []string               `json:"best_practices,omitempty"`
}
