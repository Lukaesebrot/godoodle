package godoodle

// Response represents an execute response
type Response struct {
	Output  string `json:"output"`
	Memory  string `json:"memory"`
	CPUTime string `json:"cpuTime"`
}
