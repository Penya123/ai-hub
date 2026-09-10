package models

// OllamaRequest define la estructura para enviar prompts al celular
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaResponse captura la respuesta generada por Aigis
type OllamaResponse struct {
	Response string `json:"response"`
}
