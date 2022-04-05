package Models

// Pkg exist to be imported by other pkgs, never imports pkgs itself.

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // use interface when data type is unknown
	CSRFToken string                 // Security token
	Flash     string                 // flash message to end user
	Warning   string
	Error     string
}
