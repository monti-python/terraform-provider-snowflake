// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type WarehouseShowOutputAssert struct {
	*assert.ResourceAssert
}

func WarehouseShowOutput(t *testing.T, name string) *WarehouseShowOutputAssert {
	t.Helper()

	warehouseAssert := WarehouseShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	warehouseAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &warehouseAssert
}

func ImportedWarehouseShowOutput(t *testing.T, id string) *WarehouseShowOutputAssert {
	t.Helper()

	warehouseAssert := WarehouseShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	warehouseAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &warehouseAssert
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (w *WarehouseShowOutputAssert) HasName(expected string) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasState(expected sdk.WarehouseState) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueSet("state", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasType(expected sdk.WarehouseType) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueSet("type", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasSize(expected sdk.WarehouseSize) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueSet("size", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasMinClusterCount(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("min_cluster_count", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasMaxClusterCount(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("max_cluster_count", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasStartedClusters(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("started_clusters", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasRunning(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("running", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasQueued(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("queued", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasIsDefault(expected bool) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_default", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasIsCurrent(expected bool) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_current", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasAutoSuspend(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("auto_suspend", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasAutoResume(expected bool) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueSet("auto_resume", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasAvailable(expected float64) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueSet("available", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasProvisioning(expected float64) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueSet("provisioning", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasQuiescing(expected float64) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueSet("quiescing", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasOther(expected float64) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueSet("other", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasCreatedOn(expected time.Time) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return w
}

func (w *WarehouseShowOutputAssert) HasResumedOn(expected time.Time) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("resumed_on", expected.String()))
	return w
}

func (w *WarehouseShowOutputAssert) HasUpdatedOn(expected time.Time) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("updated_on", expected.String()))
	return w
}

func (w *WarehouseShowOutputAssert) HasOwner(expected string) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasComment(expected string) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasEnableQueryAcceleration(expected bool) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueSet("enable_query_acceleration", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasQueryAccelerationMaxScaleFactor(expected int) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueSet("query_acceleration_max_scale_factor", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasResourceMonitor(expected sdk.AccountObjectIdentifier) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueSet("resource_monitor", expected.Name()))
	return w
}

func (w *WarehouseShowOutputAssert) HasScalingPolicy(expected sdk.ScalingPolicy) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueSet("scaling_policy", expected))
	return w
}

func (w *WarehouseShowOutputAssert) HasOwnerRoleType(expected string) *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return w
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (w *WarehouseShowOutputAssert) HasNoName() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("name"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoState() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueNotSet("state"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoType() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueNotSet("type"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoSize() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueNotSet("size"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoMinClusterCount() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("min_cluster_count"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoMaxClusterCount() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("max_cluster_count"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoStartedClusters() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("started_clusters"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoRunning() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("running"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoQueued() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("queued"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoIsDefault() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueNotSet("is_default"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoIsCurrent() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueNotSet("is_current"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoAutoSuspend() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("auto_suspend"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoAutoResume() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueNotSet("auto_resume"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoAvailable() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueNotSet("available"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoProvisioning() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueNotSet("provisioning"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoQuiescing() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueNotSet("quiescing"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoOther() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputFloatValueNotSet("other"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoCreatedOn() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("created_on"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoResumedOn() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("resumed_on"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoUpdatedOn() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("updated_on"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoOwner() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("owner"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoComment() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("comment"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoEnableQueryAcceleration() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputBoolValueNotSet("enable_query_acceleration"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoQueryAccelerationMaxScaleFactor() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputIntValueNotSet("query_acceleration_max_scale_factor"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoResourceMonitor() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueNotSet("resource_monitor"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoScalingPolicy() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputStringUnderlyingValueNotSet("scaling_policy"))
	return w
}

func (w *WarehouseShowOutputAssert) HasNoOwnerRoleType() *WarehouseShowOutputAssert {
	w.AddAssertion(assert.ResourceShowOutputValueNotSet("owner_role_type"))
	return w
}
