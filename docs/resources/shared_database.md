---
page_title: "snowflake_shared_database Resource - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  A shared database creates a database from a share provided by another Snowflake account. For more information about shares, see Introduction to Secure Data Sharing https://docs.snowflake.com/en/user-guide/data-sharing-intro.
---

!> **Note** The provider does not detect external changes on database type. In this case, remove the database of wrong type manually with `terraform destroy` and recreate the resource. It will be addressed in the future.

!> **Note** A database cannot be dropped successfully if it contains network rule-network policy associations. The error looks like `098507 (2BP01): Cannot drop database DATABASE as it includes network rule - policy associations.`. Currently, the provider does not unassign such objects automatically. Before dropping the resource, first unassign the network rule from the relevant objects. See [guide](../guides/unassigning_policies) for more details.

# snowflake_shared_database (Resource)

A shared database creates a database from a share provided by another Snowflake account. For more information about shares, see [Introduction to Secure Data Sharing](https://docs.snowflake.com/en/user-guide/data-sharing-intro).

## Example Usage

```terraform
# 1. Preparing database to share
resource "snowflake_share" "test" {
  provider = primary_account # notice the provider fields
  name     = "share_name"
  accounts = ["<secondary_account_organization_name>.<secondary_account_name>"]
}

resource "snowflake_database" "test" {
  provider = primary_account
  name     = "shared_database"
}

resource "snowflake_grant_privileges_to_share" "test" {
  provider    = primary_account
  to_share    = snowflake_share.test.name
  privileges  = ["USAGE"]
  on_database = snowflake_database.test.name
}

# 2. Creating shared database
## 2.1. Minimal version
resource "snowflake_shared_database" "test" {
  provider   = secondary_account
  depends_on = [snowflake_grant_privileges_to_share.test]
  name       = snowflake_database.test.name # shared database should have the same as the "imported" one
  from_share = "<primary_account_organization_name>.<primary_account_name>.${snowflake_share.test.name}"
}

## 2.2. Complete version (with every optional set)
resource "snowflake_shared_database" "test" {
  provider   = secondary_account
  depends_on = [snowflake_grant_privileges_to_share.test]
  name       = snowflake_database.test.name # shared database should have the same as the "imported" one
  from_share = "<primary_account_organization_name>.<primary_account_name>.${snowflake_share.test.name}"
  comment    = "A shared database"

  external_volume                               = "<external_volume_name>"
  catalog                                       = "<catalog_name>"
  replace_invalid_characters                    = false
  default_ddl_collation                         = "en_US"
  storage_serialization_policy                  = "COMPATIBLE"
  log_level                                     = "INFO"
  trace_level                                   = "ALWAYS"
  suspend_task_after_num_failures               = 10
  task_auto_retry_attempts                      = 10
  user_task_managed_initial_warehouse_size      = "LARGE"
  user_task_timeout_ms                          = 3600000
  user_task_minimum_trigger_interval_in_seconds = 120
  quoted_identifiers_ignore_case                = false
  enable_console_output                         = false
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `from_share` (String) A fully qualified path to a share from which the database will be created. A fully qualified path follows the format of `"<organization_name>"."<account_name>"."<share_name>"`. For more information about this resource, see [docs](./share).
- `name` (String) Specifies the identifier for the database; must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.

### Optional

- `catalog` (String) The database parameter that specifies the default catalog to use for Iceberg tables. For more information, see [CATALOG](https://docs.snowflake.com/en/sql-reference/parameters#catalog).
- `comment` (String) Specifies a comment for the database.
- `default_ddl_collation` (String) Specifies a default collation specification for all schemas and tables added to the database. It can be overridden on schema or table level. For more information, see [collation specification](https://docs.snowflake.com/en/sql-reference/collation#label-collation-specification).
- `enable_console_output` (Boolean) If true, enables stdout/stderr fast path logging for anonymous stored procedures.
- `external_volume` (String) The database parameter that specifies the default external volume to use for Iceberg tables. For more information, see [EXTERNAL_VOLUME](https://docs.snowflake.com/en/sql-reference/parameters#external-volume).
- `log_level` (String) Specifies the severity level of messages that should be ingested and made available in the active event table. Valid options are: [TRACE DEBUG INFO WARN ERROR FATAL OFF]. Messages at the specified level (and at more severe levels) are ingested. For more information, see [LOG_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters.html#label-log-level).
- `quoted_identifiers_ignore_case` (Boolean) If true, the case of quoted identifiers is ignored. For more information, see [QUOTED_IDENTIFIERS_IGNORE_CASE](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
- `replace_invalid_characters` (Boolean) Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�) in query results for an Iceberg table. You can only set this parameter for tables that use an external Iceberg catalog. For more information, see [REPLACE_INVALID_CHARACTERS](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
- `storage_serialization_policy` (String) The storage serialization policy for Iceberg tables that use Snowflake as the catalog. Valid options are: [COMPATIBLE OPTIMIZED]. COMPATIBLE: Snowflake performs encoding and compression of data files that ensures interoperability with third-party compute engines. OPTIMIZED: Snowflake performs encoding and compression of data files that ensures the best table performance within Snowflake. For more information, see [STORAGE_SERIALIZATION_POLICY](https://docs.snowflake.com/en/sql-reference/parameters#storage-serialization-policy).
- `suspend_task_after_num_failures` (Number) How many times a task must fail in a row before it is automatically suspended. 0 disables auto-suspending. For more information, see [SUSPEND_TASK_AFTER_NUM_FAILURES](https://docs.snowflake.com/en/sql-reference/parameters#suspend-task-after-num-failures).
- `task_auto_retry_attempts` (Number) Maximum automatic retries allowed for a user task. For more information, see [TASK_AUTO_RETRY_ATTEMPTS](https://docs.snowflake.com/en/sql-reference/parameters#task-auto-retry-attempts).
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `trace_level` (String) Controls how trace events are ingested into the event table. Valid options are: `ALWAYS` | `ON_EVENT` | `PROPAGATE` | `OFF`. For information about levels, see [TRACE_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters.html#label-trace-level).
- `user_task_managed_initial_warehouse_size` (String) The initial size of warehouse to use for managed warehouses in the absence of history. For more information, see [USER_TASK_MANAGED_INITIAL_WAREHOUSE_SIZE](https://docs.snowflake.com/en/sql-reference/parameters#user-task-managed-initial-warehouse-size).
- `user_task_minimum_trigger_interval_in_seconds` (Number) Minimum amount of time between Triggered Task executions in seconds.
- `user_task_timeout_ms` (Number) User task execution timeout in milliseconds. For more information, see [USER_TASK_TIMEOUT_MS](https://docs.snowflake.com/en/sql-reference/parameters#user-task-timeout-ms).

### Read-Only

- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_shared_database.example '"<shared_database_name>"'
```
