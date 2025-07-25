// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type GitRepositoryResourceAssert struct {
	*assert.ResourceAssert
}

func GitRepositoryResource(t *testing.T, name string) *GitRepositoryResourceAssert {
	t.Helper()

	return &GitRepositoryResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedGitRepositoryResource(t *testing.T, id string) *GitRepositoryResourceAssert {
	t.Helper()

	return &GitRepositoryResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (g *GitRepositoryResourceAssert) HasDatabaseString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("database", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasSchemaString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("schema", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasNameString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("name", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasApiIntegrationString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("api_integration", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasCommentString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("comment", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasFullyQualifiedNameString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasGitCredentialsString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("git_credentials", expected))
	return g
}

func (g *GitRepositoryResourceAssert) HasOriginString(expected string) *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("origin", expected))
	return g
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (g *GitRepositoryResourceAssert) HasNoDatabase() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("database"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoSchema() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("schema"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoName() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("name"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoApiIntegration() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("api_integration"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoComment() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("comment"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoFullyQualifiedName() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoGitCredentials() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("git_credentials"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNoOrigin() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueNotSet("origin"))
	return g
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (g *GitRepositoryResourceAssert) HasCommentEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("comment", ""))
	return g
}

func (g *GitRepositoryResourceAssert) HasFullyQualifiedNameEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return g
}

func (g *GitRepositoryResourceAssert) HasGitCredentialsEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValueSet("git_credentials", ""))
	return g
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (g *GitRepositoryResourceAssert) HasDatabaseNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("database"))
	return g
}

func (g *GitRepositoryResourceAssert) HasSchemaNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("schema"))
	return g
}

func (g *GitRepositoryResourceAssert) HasNameNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("name"))
	return g
}

func (g *GitRepositoryResourceAssert) HasApiIntegrationNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("api_integration"))
	return g
}

func (g *GitRepositoryResourceAssert) HasCommentNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("comment"))
	return g
}

func (g *GitRepositoryResourceAssert) HasFullyQualifiedNameNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return g
}

func (g *GitRepositoryResourceAssert) HasGitCredentialsNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("git_credentials"))
	return g
}

func (g *GitRepositoryResourceAssert) HasOriginNotEmpty() *GitRepositoryResourceAssert {
	g.AddAssertion(assert.ValuePresent("origin"))
	return g
}
