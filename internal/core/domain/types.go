package domain

type ContentType string

const (
	ContentTypeText ContentType = "text"
)

type MessageRole string

const (
	MessageRoleSystem    MessageRole = "system"
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
)

type ContentPart struct {
	Type ContentType `json:"type"`
	Text string      `json:"text,omitempty"`
	URI  string      `json:"uri,omitempty"`
}

type Message struct {
	Role  MessageRole   `json:"role"`
	Parts []ContentPart `json:"parts"`
}

type ProviderRef struct {
	Name string `json:"name"`
}

type ModelRef struct {
	Provider string `json:"provider"`
	Name     string `json:"name"`
	Alias    string `json:"alias,omitempty"`
}

type CapabilitySet struct {
	Chat        bool `json:"chat"`
	Streaming   bool `json:"streaming"`
	Local       bool `json:"local"`
	Remote      bool `json:"remote"`
	ToolCalling bool `json:"tool_calling"`
}

type ExecutionParameters struct {
	Temperature float64 `json:"temperature" toml:"temperature"`
	MaxTokens   int     `json:"max_tokens" toml:"max_tokens"`
	Thinking    bool    `json:"thinking" toml:"thinking"`
}

type ExecutionRequest struct {
	Model        ModelRef            `json:"model"`
	Messages     []Message           `json:"messages"`
	SystemPrompt string              `json:"system_prompt,omitempty"`
	Parameters   ExecutionParameters `json:"parameters"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type ExecutionResult struct {
	Provider ProviderRef `json:"provider"`
	Model    ModelRef    `json:"model"`
	Message  Message     `json:"message"`
	Usage    Usage       `json:"usage"`
}

type Session struct {
	ID       string    `json:"id"`
	Messages []Message `json:"messages"`
}

type ResourceSnapshot struct {
	CPUPercent  float64  `json:"cpu_percent"`
	MemoryBytes uint64   `json:"memory_bytes"`
	GPUPercent  *float64 `json:"gpu_percent,omitempty"`
	VRAMBytes   *uint64  `json:"vram_bytes,omitempty"`
}

type PathsConfig struct {
	ConfigDir string `toml:"config_dir"`
	DataDir   string `toml:"data_dir"`
	CacheDir  string `toml:"cache_dir"`
	LogDir    string `toml:"log_dir"`
}

type DefaultsConfig struct {
	Provider     string  `toml:"provider"`
	Model        string  `toml:"model"`
	Alias        string  `toml:"alias"`
	SystemPrompt string  `toml:"system_prompt"`
	Temperature  float64 `toml:"temperature"`
	MaxTokens    int     `toml:"max_tokens"`
	Thinking     bool    `toml:"thinking"`
}

type ProfileConfig struct {
	Provider     *string  `toml:"provider"`
	Model        *string  `toml:"model"`
	Alias        *string  `toml:"alias"`
	SystemPrompt *string  `toml:"system_prompt"`
	Temperature  *float64 `toml:"temperature"`
	MaxTokens    *int     `toml:"max_tokens"`
	Thinking     *bool    `toml:"thinking"`
}

type AliasConfig struct {
	Provider string `toml:"provider"`
	Model    string `toml:"model"`
}

type AppConfig struct {
	DefaultProfile string                   `toml:"default_profile"`
	Paths          PathsConfig              `toml:"paths"`
	Defaults       DefaultsConfig           `toml:"defaults"`
	Profiles       map[string]ProfileConfig `toml:"profiles"`
	Aliases        map[string]AliasConfig   `toml:"aliases"`
}

type CommandOverrides struct {
	Profile      string
	Provider     *string
	Model        *string
	Alias        *string
	SystemPrompt *string
	Temperature  *float64
	MaxTokens    *int
	Thinking     *bool
}

type ResolvedSettings struct {
	Profile      string              `json:"profile"`
	Provider     string              `json:"provider"`
	Model        string              `json:"model"`
	Alias        string              `json:"alias,omitempty"`
	SystemPrompt string              `json:"system_prompt,omitempty"`
	Parameters   ExecutionParameters `json:"parameters"`
}

func UserMessage(text string) Message {
	return Message{
		Role: MessageRoleUser,
		Parts: []ContentPart{{
			Type: ContentTypeText,
			Text: text,
		}},
	}
}

func AssistantMessage(text string) Message {
	return Message{
		Role: MessageRoleAssistant,
		Parts: []ContentPart{{
			Type: ContentTypeText,
			Text: text,
		}},
	}
}

func MessageText(message Message) string {
	for _, part := range message.Parts {
		if part.Type == ContentTypeText {
			return part.Text
		}
	}

	return ""
}

func DefaultConfig() AppConfig {
	return AppConfig{
		DefaultProfile: "default",
		Defaults: DefaultsConfig{
			Provider:    "local-stub",
			Model:       "echo-1",
			Temperature: 0.2,
			MaxTokens:   512,
		},
		Profiles: map[string]ProfileConfig{},
		Aliases:  map[string]AliasConfig{},
	}
}
