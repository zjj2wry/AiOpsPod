package llmservice

import (
	"fmt"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

func NewAgent(llm llms.Model, tools []tools.Tool, document string) *agents.Executor {
	prefix := `Today is {{.today}}.
	Answer the following questions as best you can. You have access to the following tools:
	
	{{.tool_descriptions}}`

	if document != "" {
		prefix = fmt.Sprintf(`The documents found here may be related, if not, you can ignore them:
		"%s"\n%s`, document, prefix)
	}
	agent := agents.NewOneShotAgent(
		llm,
		tools,
		agents.WithMaxIterations(3),
		agents.WithReturnIntermediateSteps(),
		agents.WithPromptPrefix(prefix),
	)

	return agents.NewExecutor(agent)
}
