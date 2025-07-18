// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type AuthenticationPolicyShowOutputAssert struct {
	*assert.ResourceAssert
}

func AuthenticationPolicyShowOutput(t *testing.T, name string) *AuthenticationPolicyShowOutputAssert {
	t.Helper()

	authenticationPolicyAssert := AuthenticationPolicyShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	authenticationPolicyAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &authenticationPolicyAssert
}

func ImportedAuthenticationPolicyShowOutput(t *testing.T, id string) *AuthenticationPolicyShowOutputAssert {
	t.Helper()

	authenticationPolicyAssert := AuthenticationPolicyShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	authenticationPolicyAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &authenticationPolicyAssert
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (a *AuthenticationPolicyShowOutputAssert) HasCreatedOn(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasName(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasComment(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasDatabaseName(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasSchemaName(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasOwner(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasOwnerRoleType(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasOptions(expected string) *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueSet("options", expected))
	return a
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (a *AuthenticationPolicyShowOutputAssert) HasNoCreatedOn() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("created_on"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoName() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("name"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoComment() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("comment"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoDatabaseName() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("database_name"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoSchemaName() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("schema_name"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoOwner() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("owner"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoOwnerRoleType() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("owner_role_type"))
	return a
}

func (a *AuthenticationPolicyShowOutputAssert) HasNoOptions() *AuthenticationPolicyShowOutputAssert {
	a.AddAssertion(assert.ResourceShowOutputValueNotSet("options"))
	return a
}
