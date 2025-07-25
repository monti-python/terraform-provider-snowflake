---
page_title: "snowflake_database Resource - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  Represents a standard database. If replication configuration is specified, the database is promoted to serve as a primary database for replication.
---

!> **Note** The provider does not detect external changes on database type. In this case, remove the database of wrong type manually with `terraform destroy` and recreate the resource. It will be addressed in the future.

!> **Note** A database cannot be dropped successfully if it contains network rule-network policy associations. The error looks like `098507 (2BP01): Cannot drop database DATABASE as it includes network rule - policy associations.`. Currently, the provider does not unassign such objects automatically. Before dropping the resource, first unassign the network rule from the relevant objects. See [guide](../guides/unassigning_policies) for more details.

# snowflake_database (Resource)

Represents a standard database. If replication configuration is specified, the database is promoted to serve as a primary database for replication.

## Example Usage

```terraform
## Minimal
resource "snowflake_database" "primary" {
  name = "database_name"
}

## Complete (with every optional set)
resource "snowflake_database" "primary" {
  name         = "database_name"
  is_transient = false
  comment      = "my standard database"

  data_retention_time_in_days                   = 10
  max_data_extension_time_in_days               = 20
  external_volume                               = snowflake_external_volume.example.fully_qualified_name
  catalog                                       = snowflake_catalog.example.fully_qualified_name
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

  replication {
    enable_to_account {
      account_identifier = "<secondary_account_organization_name>.<secondary_account_name>"
      with_failover      = true
    }
    ignore_edition_check = true
  }
}

## Replication with dynamic block
locals {
  replication_configs = [
    {
      account_identifier = "\"<secondary_account_organization_1_name>\".\"<secondary_account_1_name>\""
      with_failover      = true
    },
    {
      account_identifier = "\"<secondary_account_organization_2_name>\".\"<secondary_account_2_name>\""
      with_failover      = true
    },
  ]
}

resource "snowflake_database" "primary" {
  name = "database_name"

  replication {
    dynamic "enable_to_account" {
      for_each = { for rc in local.replication_configs : rc.account_identifier => rc }
      content {
        account_identifier = enable_to_account.value.account_identifier
        with_failover      = enable_to_account.value.with_failover
      }
    }
    ignore_edition_check = true
  }
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specifies the identifier for the database; must be unique for your account. As a best practice for [Database Replication and Failover](https://docs.snowflake.com/en/user-guide/db-replication-intro), it is recommended to give each secondary database the same name as its primary database. This practice supports referencing fully-qualified objects (i.e. '<db>.<schema>.<object>') by other objects in the same database, such as querying a fully-qualified table name in a view. If a secondary database has a different name from the primary database, then these object references would break in the secondary database. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.

### Optional

- `catalog` (String) The database parameter that specifies the default catalog to use for Iceberg tables. For more information, see [CATALOG](https://docs.snowflake.com/en/sql-reference/parameters#catalog).
- `comment` (String) Specifies a comment for the database.
- `data_retention_time_in_days` (Number) Specifies the number of days for which Time Travel actions (CLONE and UNDROP) can be performed on the database, as well as specifying the default Time Travel retention time for all schemas created in the database. For more details, see [Understanding & Using Time Travel](https://docs.snowflake.com/en/user-guide/data-time-travel).
- `default_ddl_collation` (String) Specifies a default collation specification for all schemas and tables added to the database. It can be overridden on schema or table level. For more information, see [collation specification](https://docs.snowflake.com/en/sql-reference/collation#label-collation-specification).
- `drop_public_schema_on_creation` (Boolean) Specifies whether to drop public schema on creation or not. Modifying the parameter after database is already created won't have any effect.
- `enable_console_output` (Boolean) If true, enables stdout/stderr fast path logging for anonymous stored procedures.
- `external_volume` (String) The database parameter that specifies the default external volume to use for Iceberg tables. For more information, see [EXTERNAL_VOLUME](https://docs.snowflake.com/en/sql-reference/parameters#external-volume).
- `is_transient` (Boolean) Specifies the database as transient. Transient databases do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss.
- `log_level` (String) Specifies the severity level of messages that should be ingested and made available in the active event table. Valid options are: [TRACE DEBUG INFO WARN ERROR FATAL OFF]. Messages at the specified level (and at more severe levels) are ingested. For more information, see [LOG_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters.html#label-log-level).
- `max_data_extension_time_in_days` (Number) Object parameter that specifies the maximum number of days for which Snowflake can extend the data retention period for tables in the database to prevent streams on the tables from becoming stale. For a detailed description of this parameter, see [MAX_DATA_EXTENSION_TIME_IN_DAYS](https://docs.snowflake.com/en/sql-reference/parameters.html#label-max-data-extension-time-in-days).
- `quoted_identifiers_ignore_case` (Boolean) If true, the case of quoted identifiers is ignored. For more information, see [QUOTED_IDENTIFIERS_IGNORE_CASE](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
- `replace_invalid_characters` (Boolean) Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�) in query results for an Iceberg table. You can only set this parameter for tables that use an external Iceberg catalog. For more information, see [REPLACE_INVALID_CHARACTERS](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
- `replication` (Block List, Max: 1) Configures replication for a given database. When specified, this database will be promoted to serve as a primary database for replication. A primary database can be replicated in one or more accounts, allowing users in those accounts to query objects in each secondary (i.e. replica) database. (see [below for nested schema](#nestedblock--replication))
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

<a id="nestedblock--replication"></a>
### Nested Schema for `replication`

Required:

- `enable_to_account` (Block List, Min: 1) Entry to enable replication and optionally failover for a given account identifier. (see [below for nested schema](#nestedblock--replication--enable_to_account))

Optional:

- `ignore_edition_check` (Boolean) Allows replicating data to accounts on lower editions in either of the following scenarios: 1. The primary database is in a Business Critical (or higher) account but one or more of the accounts approved for replication are on lower editions. Business Critical Edition is intended for Snowflake accounts with extremely sensitive data. 2. The primary database is in a Business Critical (or higher) account and a signed business associate agreement is in place to store PHI data in the account per HIPAA and HITRUST regulations, but no such agreement is in place for one or more of the accounts approved for replication, regardless if they are Business Critical (or higher) accounts. Both scenarios are prohibited by default in an effort to help prevent account administrators for Business Critical (or higher) accounts from inadvertently replicating sensitive data to accounts on lower editions.

<a id="nestedblock--replication--enable_to_account"></a>
### Nested Schema for `replication.enable_to_account`

Required:

- `account_identifier` (String) Specifies account identifier for which replication should be enabled. The account identifiers should be in the form of `"<organization_name>"."<account_name>"`. For more information about this resource, see [docs](./account).

Optional:

- `with_failover` (Boolean) Specifies if failover should be enabled for the specified account identifier



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
terraform import snowflake_database.example '"<database_name>"'
```
