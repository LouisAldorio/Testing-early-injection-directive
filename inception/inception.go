package inception

import (
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
)

type inception struct {
}

// New returns a federation plugin that injects
// federated directives and types into the schema
func New() plugin.Plugin {
	return &inception{}
}

// Name returns the plugin name
func (f *inception) Name() string {
	return "inception"
}

// MutateConfig mutates the configuration
func (f *inception) MutateConfig(cfg *config.Config) error {
	cfg.Directives["hasRole"] = config.DirectiveConfig{SkipRuntime: false}

	return nil
}

func (f *inception) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name: "inception/directives.graphql",
		Input: `

			directive @hasRole(role: String!) on FIELD_DEFINITION

			directive @goField(
				forceResolver: Boolean
				name: String
			  ) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION
		`,
		BuiltIn: false,
	}
}
