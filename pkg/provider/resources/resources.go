package resources

type resource string

const (
	Account                                                resource = "snowflake_account"
	AccountAuthenticationPolicyAttachment                  resource = "snowflake_account_authentication_policy_attachment"
	AccountParameter                                       resource = "snowflake_account_parameter"
	AccountPasswordPolicyAttachment                        resource = "snowflake_account_password_policy_attachment"
	AccountRole                                            resource = "snowflake_account_role"
	Alert                                                  resource = "snowflake_alert"
	ApiAuthenticationIntegrationWithAuthorizationCodeGrant resource = "snowflake_api_authentication_integration_with_authorization_code_grant"
	ApiAuthenticationIntegrationWithClientCredentials      resource = "snowflake_api_authentication_integration_with_client_credentials"
	ApiAuthenticationIntegrationWithJwtBearer              resource = "snowflake_api_authentication_integration_with_jwt_bearer"
	ApiIntegration                                         resource = "snowflake_api_integration"
	AuthenticationPolicy                                   resource = "snowflake_authentication_policy"
	ComputePool                                            resource = "snowflake_compute_pool"
	CortexSearchService                                    resource = "snowflake_cortex_search_service"
	CurrentAccount                                         resource = "snowflake_current_account"
	Database                                               resource = "snowflake_database"
	DatabaseRole                                           resource = "snowflake_database_role"
	DynamicTable                                           resource = "snowflake_dynamic_table"
	EmailNotificationIntegration                           resource = "snowflake_email_notification_integration"
	Execute                                                resource = "snowflake_execute"
	ExternalFunction                                       resource = "snowflake_external_function"
	ExternalTable                                          resource = "snowflake_external_table"
	ExternalOauthSecurityIntegration                       resource = "snowflake_external_oauth_integration"
	ExternalVolume                                         resource = "snowflake_external_volume"
	FailoverGroup                                          resource = "snowflake_failover_group"
	FileFormat                                             resource = "snowflake_file_format"
	FunctionJava                                           resource = "snowflake_function_java"
	FunctionJavascript                                     resource = "snowflake_function_javascript"
	FunctionPython                                         resource = "snowflake_function_python"
	FunctionScala                                          resource = "snowflake_function_scala"
	FunctionSql                                            resource = "snowflake_function_sql"
	GitRepository                                          resource = "snowflake_git_repository"
	GrantAccountRole                                       resource = "snowflake_grant_account_role"
	GrantApplicationRole                                   resource = "snowflake_grant_application_role"
	GrantDatabaseRole                                      resource = "snowflake_grant_database_role"
	GrantOwnership                                         resource = "snowflake_grant_ownership"
	GrantPrivilegesToAccountRole                           resource = "snowflake_grant_privileges_to_account_role"
	GrantPrivilegesToDatabaseRole                          resource = "snowflake_grant_privileges_to_database_role"
	GrantPrivilegesToShare                                 resource = "snowflake_grant_privileges_to_share"
	ImageRepository                                        resource = "snowflake_image_repository"
	JobService                                             resource = "snowflake_job_service"
	LegacyServiceUser                                      resource = "snowflake_legacy_service_user"
	ManagedAccount                                         resource = "snowflake_managed_account"
	MaskingPolicy                                          resource = "snowflake_masking_policy"
	MaterializedView                                       resource = "snowflake_materialized_view"
	NetworkPolicy                                          resource = "snowflake_network_policy"
	NetworkPolicyAttachment                                resource = "snowflake_network_policy_attachment"
	NetworkRule                                            resource = "snowflake_network_rule"
	NotificationIntegration                                resource = "snowflake_notification_integration"
	OauthIntegration                                       resource = "snowflake_oauth_integration"
	OauthIntegrationForCustomClients                       resource = "snowflake_oauth_integration_for_custom_clients"
	OauthIntegrationForPartnerApplications                 resource = "snowflake_oauth_integration_for_partner_applications"
	ObjectParameter                                        resource = "snowflake_object_parameter"
	PasswordPolicy                                         resource = "snowflake_password_policy"
	Pipe                                                   resource = "snowflake_pipe"
	PrimaryConnection                                      resource = "snowflake_primary_connection"
	ProcedureJava                                          resource = "snowflake_procedure_java"
	ProcedureJavascript                                    resource = "snowflake_procedure_javascript"
	ProcedurePython                                        resource = "snowflake_procedure_python"
	ProcedureScala                                         resource = "snowflake_procedure_scala"
	ProcedureSql                                           resource = "snowflake_procedure_sql"
	ResourceMonitor                                        resource = "snowflake_resource_monitor"
	RowAccessPolicy                                        resource = "snowflake_row_access_policy"
	SamlSecurityIntegration                                resource = "snowflake_saml_integration"
	Saml2SecurityIntegration                               resource = "snowflake_saml2_integration"
	Schema                                                 resource = "snowflake_schema"
	ScimSecurityIntegration                                resource = "snowflake_scim_integration"
	SecondaryConnection                                    resource = "snowflake_secondary_connection"
	SecondaryDatabase                                      resource = "snowflake_secondary_database"
	SecretWithAuthorizationCodeGrant                       resource = "snowflake_secret_with_authorization_code_grant"
	SecretWithBasicAuthentication                          resource = "snowflake_secret_with_basic_authentication"
	SecretWithClientCredentials                            resource = "snowflake_secret_with_client_credentials"
	SecretWithGenericString                                resource = "snowflake_secret_with_generic_string"
	SessionParameter                                       resource = "snowflake_session_parameter"
	Sequence                                               resource = "snowflake_sequence"
	Service                                                resource = "snowflake_service"
	ServiceUser                                            resource = "snowflake_service_user"
	Share                                                  resource = "snowflake_share"
	SharedDatabase                                         resource = "snowflake_shared_database"
	Stage                                                  resource = "snowflake_stage"
	StorageIntegration                                     resource = "snowflake_storage_integration"
	StreamOnDirectoryTable                                 resource = "snowflake_stream_on_directory_table"
	StreamOnExternalTable                                  resource = "snowflake_stream_on_external_table"
	StreamOnTable                                          resource = "snowflake_stream_on_table"
	StreamOnView                                           resource = "snowflake_stream_on_view"
	Streamlit                                              resource = "snowflake_streamlit"
	Table                                                  resource = "snowflake_table"
	TableColumnMaskingPolicyApplication                    resource = "snowflake_table_column_masking_policy_application"
	TableConstraint                                        resource = "snowflake_table_constraint"
	Tag                                                    resource = "snowflake_tag"
	TagAssociation                                         resource = "snowflake_tag_association"
	TagMaskingPolicyAssociation                            resource = "snowflake_tag_masking_policy_association"
	Task                                                   resource = "snowflake_task"
	User                                                   resource = "snowflake_user"
	UserAuthenticationPolicyAttachment                     resource = "snowflake_user_authentication_policy_attachment"
	UserPasswordPolicyAttachment                           resource = "snowflake_user_password_policy_attachment"
	UserPublicKeys                                         resource = "snowflake_user_public_keys"
	View                                                   resource = "snowflake_view"
	Warehouse                                              resource = "snowflake_warehouse"
)

type Resource interface {
	xxxProtected()
	String() string
}

func (r resource) xxxProtected() {}

func (r resource) String() string {
	return string(r)
}
