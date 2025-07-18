---
page_title: "snowflake_cortex_search_services Data Source - terraform-provider-snowflake"
subcategory: "Preview"
description: |-
  
---

!> **Caution: Preview Feature** This feature is considered a preview feature in the provider, regardless of the state of the resource in Snowflake. We do not guarantee its stability. It will be reworked and marked as a stable feature in future releases. Breaking changes are expected, even without bumping the major version. To use this feature, add the relevant feature name to `preview_features_enabled` field in the [provider configuration](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#schema). Please always refer to the [Getting Help](https://github.com/snowflakedb/terraform-provider-snowflake?tab=readme-ov-file#getting-help) section in our Github repo to best determine how to get help for your questions.

# snowflake_cortex_search_services (Data Source)



## Example Usage

```terraform
# Simple usage
data "snowflake_cortex_search_services" "simple" {
}

output "simple_output" {
  value = data.snowflake_cortex_search_services.simple.cortex_search_services
}

# Filtering (like)
data "snowflake_cortex_search_services" "like" {
  like = "some-name"
}

output "like_output" {
  value = data.snowflake_cortex_search_services.like.cortex_search_services
}

# Filtering (starts_with)
data "snowflake_cortex_search_services" "starts_with" {
  starts_with = "prefix-"
}

output "starts_with_output" {
  value = data.snowflake_cortex_search_services.starts_with.cortex_search_services
}

# Filtering (limit)
data "snowflake_cortex_search_services" "limit" {
  limit {
    rows = 10
    from = "prefix-"
  }
}

output "limit_output" {
  value = data.snowflake_cortex_search_services.limit.cortex_search_services
}
```

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `in` (Block List, Max: 1) IN clause to filter the list of cortex search services. (see [below for nested schema](#nestedblock--in))
- `like` (String) Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
- `limit` (Block List, Max: 1) Limits the number of rows returned. If the `limit.from` is set, then the limit will start from the first element matched by the expression. The expression is only used to match with the first element, later on the elements are not matched by the prefix, but you can enforce a certain pattern with `starts_with` or `like`. (see [below for nested schema](#nestedblock--limit))
- `starts_with` (String) Filters the output with **case-sensitive** characters indicating the beginning of the object name.

### Read-Only

- `cortex_search_services` (List of Object) Holds the output of SHOW CORTEX SEARCH SERVICES. (see [below for nested schema](#nestedatt--cortex_search_services))
- `id` (String) The ID of this resource.

<a id="nestedblock--in"></a>
### Nested Schema for `in`

Optional:

- `account` (Boolean) Returns records for the entire account.
- `database` (String) Returns records for the current database in use or for a specified database (db_name).
- `schema` (String) Returns records for the current schema in use or a specified schema (schema_name).


<a id="nestedblock--limit"></a>
### Nested Schema for `limit`

Required:

- `rows` (Number) The maximum number of rows to return.

Optional:

- `from` (String) Specifies a **case-sensitive** pattern that is used to match object name. After the first match, the limit on the number of rows will be applied.


<a id="nestedatt--cortex_search_services"></a>
### Nested Schema for `cortex_search_services`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `database_name` (String)
- `name` (String)
- `schema_name` (String)
