package tui

type Model struct {
	step         int
	ProjectDir   string
	AiBackend    string
	AiModel      string
	AiEndpoint   string
	AiAPIKey     string
	OutputFolder string
	verbose      bool

	inputValue string   // current text input buffer
	choices    []string // for selection steps
	cursor     int      // for selecting choices
}

func initialModel() Model {
	return Model{
		step:         0,
		ProjectDir:   ".",
		OutputFolder: "docs",
		verbose:      true,
		inputValue:   "",
		cursor:       0,
	}
}
