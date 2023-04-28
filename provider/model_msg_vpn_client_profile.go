package provider

import (
	"telusag/terraform-provider-solace/sempv2"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// MsgVpnClientProfile struct for MsgVpnClientProfile
type MsgVpnClientProfile struct {
	AllowBridgeConnectionsEnabled                            *bool                    `tfsdk:"allow_bridge_connections_enabled"`
	AllowCutThroughForwardingEnabled                         *bool                    `tfsdk:"allow_cut_through_forwarding_enabled"`
	AllowGuaranteedEndpointCreateDurability                  *string                  `tfsdk:"allow_guaranteed_endpoint_create_durability"`
	AllowGuaranteedEndpointCreateEnabled                     *bool                    `tfsdk:"allow_guaranteed_endpoint_create_enabled"`
	AllowGuaranteedMsgReceiveEnabled                         *bool                    `tfsdk:"allow_guaranteed_msg_receive_enabled"`
	AllowGuaranteedMsgSendEnabled                            *bool                    `tfsdk:"allow_guaranteed_msg_send_enabled"`
	AllowSharedSubscriptionsEnabled                          *bool                    `tfsdk:"allow_shared_subscriptions_enabled"`
	AllowTransactedSessionsEnabled                           *bool                    `tfsdk:"allow_transacted_sessions_enabled"`
	ApiQueueManagementCopyFromOnCreateName                   *string                  `tfsdk:"api_queue_management_copy_from_on_create_name"`
	ApiQueueManagementCopyFromOnCreateTemplateName           *string                  `tfsdk:"api_queue_management_copy_from_on_create_template_name"`
	ApiTopicEndpointManagementCopyFromOnCreateName           *string                  `tfsdk:"api_topic_endpoint_management_copy_from_on_create_name"`
	ApiTopicEndpointManagementCopyFromOnCreateTemplateName   *string                  `tfsdk:"api_topic_endpoint_management_copy_from_on_create_template_name"`
	ClientProfileName                                        *string                  `tfsdk:"client_profile_name"`
	CompressionEnabled                                       *bool                    `tfsdk:"compression_enabled"`
	ElidingDelay                                             *int64                   `tfsdk:"eliding_delay"`
	ElidingEnabled                                           *bool                    `tfsdk:"eliding_enabled"`
	ElidingMaxTopicCount                                     *int64                   `tfsdk:"eliding_max_topic_count"`
	EventClientProvisionedEndpointSpoolUsageThreshold        *EventThresholdByPercent `tfsdk:"event_client_provisioned_endpoint_spool_usage_threshold"`
	EventConnectionCountPerClientUsernameThreshold           *EventThreshold          `tfsdk:"event_connection_count_per_client_username_threshold"`
	EventEgressFlowCountThreshold                            *EventThreshold          `tfsdk:"event_egress_flow_count_threshold"`
	EventEndpointCountPerClientUsernameThreshold             *EventThreshold          `tfsdk:"event_endpoint_count_per_client_username_threshold"`
	EventIngressFlowCountThreshold                           *EventThreshold          `tfsdk:"event_ingress_flow_count_threshold"`
	EventServiceSmfConnectionCountPerClientUsernameThreshold *EventThreshold          `tfsdk:"event_service_smf_connection_count_per_client_username_threshold"`
	EventServiceWebConnectionCountPerClientUsernameThreshold *EventThreshold          `tfsdk:"event_service_web_connection_count_per_client_username_threshold"`
	EventSubscriptionCountThreshold                          *EventThreshold          `tfsdk:"event_subscription_count_threshold"`
	EventTransactedSessionCountThreshold                     *EventThreshold          `tfsdk:"event_transacted_session_count_threshold"`
	EventTransactionCountThreshold                           *EventThreshold          `tfsdk:"event_transaction_count_threshold"`
	MaxConnectionCountPerClientUsername                      *int64                   `tfsdk:"max_connection_count_per_client_username"`
	MaxEgressFlowCount                                       *int64                   `tfsdk:"max_egress_flow_count"`
	MaxEndpointCountPerClientUsername                        *int64                   `tfsdk:"max_endpoint_count_per_client_username"`
	MaxIngressFlowCount                                      *int64                   `tfsdk:"max_ingress_flow_count"`
	MaxMsgsPerTransaction                                    *int32                   `tfsdk:"max_msgs_per_transaction"`
	MaxSubscriptionCount                                     *int64                   `tfsdk:"max_subscription_count"`
	MaxTransactedSessionCount                                *int64                   `tfsdk:"max_transacted_session_count"`
	MaxTransactionCount                                      *int64                   `tfsdk:"max_transaction_count"`
	MsgVpnName                                               *string                  `tfsdk:"msg_vpn_name"`
	QueueControl1MaxDepth                                    *int32                   `tfsdk:"queue_control1_max_depth"`
	QueueControl1MinMsgBurst                                 *int32                   `tfsdk:"queue_control1_min_msg_burst"`
	QueueDirect1MaxDepth                                     *int32                   `tfsdk:"queue_direct1_max_depth"`
	QueueDirect1MinMsgBurst                                  *int32                   `tfsdk:"queue_direct1_min_msg_burst"`
	QueueDirect2MaxDepth                                     *int32                   `tfsdk:"queue_direct2_max_depth"`
	QueueDirect2MinMsgBurst                                  *int32                   `tfsdk:"queue_direct2_min_msg_burst"`
	QueueDirect3MaxDepth                                     *int32                   `tfsdk:"queue_direct3_max_depth"`
	QueueDirect3MinMsgBurst                                  *int32                   `tfsdk:"queue_direct3_min_msg_burst"`
	QueueGuaranteed1MaxDepth                                 *int32                   `tfsdk:"queue_guaranteed1_max_depth"`
	QueueGuaranteed1MinMsgBurst                              *int32                   `tfsdk:"queue_guaranteed1_min_msg_burst"`
	RejectMsgToSenderOnNoSubscriptionMatchEnabled            *bool                    `tfsdk:"reject_msg_to_sender_on_no_subscription_match_enabled"`
	ReplicationAllowClientConnectWhenStandbyEnabled          *bool                    `tfsdk:"replication_allow_client_connect_when_standby_enabled"`
	ServiceMinKeepaliveTimeout                               *int32                   `tfsdk:"service_min_keepalive_timeout"`
	ServiceSmfMaxConnectionCountPerClientUsername            *int64                   `tfsdk:"service_smf_max_connection_count_per_client_username"`
	ServiceSmfMinKeepaliveEnabled                            *bool                    `tfsdk:"service_smf_min_keepalive_enabled"`
	ServiceWebInactiveTimeout                                *int64                   `tfsdk:"service_web_inactive_timeout"`
	ServiceWebMaxConnectionCountPerClientUsername            *int64                   `tfsdk:"service_web_max_connection_count_per_client_username"`
	ServiceWebMaxPayload                                     *int64                   `tfsdk:"service_web_max_payload"`
	TcpCongestionWindowSize                                  *int64                   `tfsdk:"tcp_congestion_window_size"`
	TcpKeepaliveCount                                        *int64                   `tfsdk:"tcp_keepalive_count"`
	TcpKeepaliveIdleTime                                     *int64                   `tfsdk:"tcp_keepalive_idle_time"`
	TcpKeepaliveInterval                                     *int64                   `tfsdk:"tcp_keepalive_interval"`
	TcpMaxSegmentSize                                        *int64                   `tfsdk:"tcp_max_segment_size"`
	TcpMaxWindowSize                                         *int64                   `tfsdk:"tcp_max_window_size"`
	TlsAllowDowngradeToPlainTextEnabled                      *bool                    `tfsdk:"tls_allow_downgrade_to_plain_text_enabled"`
}

func (tfData *MsgVpnClientProfile) ToTF(apiData *sempv2.MsgVpnClientProfile) {
	AssignIfDstNotNil(&tfData.AllowBridgeConnectionsEnabled, apiData.AllowBridgeConnectionsEnabled)
	AssignIfDstNotNil(&tfData.AllowCutThroughForwardingEnabled, apiData.AllowCutThroughForwardingEnabled)
	AssignIfDstNotNil(&tfData.AllowGuaranteedEndpointCreateDurability, apiData.AllowGuaranteedEndpointCreateDurability)
	AssignIfDstNotNil(&tfData.AllowGuaranteedEndpointCreateEnabled, apiData.AllowGuaranteedEndpointCreateEnabled)
	AssignIfDstNotNil(&tfData.AllowGuaranteedMsgReceiveEnabled, apiData.AllowGuaranteedMsgReceiveEnabled)
	AssignIfDstNotNil(&tfData.AllowGuaranteedMsgSendEnabled, apiData.AllowGuaranteedMsgSendEnabled)
	AssignIfDstNotNil(&tfData.AllowSharedSubscriptionsEnabled, apiData.AllowSharedSubscriptionsEnabled)
	AssignIfDstNotNil(&tfData.AllowTransactedSessionsEnabled, apiData.AllowTransactedSessionsEnabled)
	AssignIfDstNotNil(&tfData.ApiQueueManagementCopyFromOnCreateName, apiData.ApiQueueManagementCopyFromOnCreateName)
	AssignIfDstNotNil(&tfData.ApiQueueManagementCopyFromOnCreateTemplateName, apiData.ApiQueueManagementCopyFromOnCreateTemplateName)
	AssignIfDstNotNil(&tfData.ApiTopicEndpointManagementCopyFromOnCreateName, apiData.ApiTopicEndpointManagementCopyFromOnCreateName)
	AssignIfDstNotNil(&tfData.ApiTopicEndpointManagementCopyFromOnCreateTemplateName, apiData.ApiTopicEndpointManagementCopyFromOnCreateTemplateName)
	AssignIfDstNotNil(&tfData.ClientProfileName, apiData.ClientProfileName)
	AssignIfDstNotNil(&tfData.CompressionEnabled, apiData.CompressionEnabled)
	AssignIfDstNotNil(&tfData.ElidingDelay, apiData.ElidingDelay)
	AssignIfDstNotNil(&tfData.ElidingEnabled, apiData.ElidingEnabled)
	AssignIfDstNotNil(&tfData.ElidingMaxTopicCount, apiData.ElidingMaxTopicCount)
	AssignIfDstNotNil(&tfData.EventClientProvisionedEndpointSpoolUsageThreshold, EventThresholdByPercentToTF(apiData.EventClientProvisionedEndpointSpoolUsageThreshold))
	AssignIfDstNotNil(&tfData.EventConnectionCountPerClientUsernameThreshold, EventThresholdToTF(apiData.EventConnectionCountPerClientUsernameThreshold))
	AssignIfDstNotNil(&tfData.EventEgressFlowCountThreshold, EventThresholdToTF(apiData.EventEgressFlowCountThreshold))
	AssignIfDstNotNil(&tfData.EventEndpointCountPerClientUsernameThreshold, EventThresholdToTF(apiData.EventEndpointCountPerClientUsernameThreshold))
	AssignIfDstNotNil(&tfData.EventIngressFlowCountThreshold, EventThresholdToTF(apiData.EventIngressFlowCountThreshold))
	AssignIfDstNotNil(&tfData.EventServiceSmfConnectionCountPerClientUsernameThreshold, EventThresholdToTF(apiData.EventServiceSmfConnectionCountPerClientUsernameThreshold))
	AssignIfDstNotNil(&tfData.EventServiceWebConnectionCountPerClientUsernameThreshold, EventThresholdToTF(apiData.EventServiceWebConnectionCountPerClientUsernameThreshold))
	AssignIfDstNotNil(&tfData.EventSubscriptionCountThreshold, EventThresholdToTF(apiData.EventSubscriptionCountThreshold))
	AssignIfDstNotNil(&tfData.EventTransactedSessionCountThreshold, EventThresholdToTF(apiData.EventTransactedSessionCountThreshold))
	AssignIfDstNotNil(&tfData.EventTransactionCountThreshold, EventThresholdToTF(apiData.EventTransactionCountThreshold))
	AssignIfDstNotNil(&tfData.MaxConnectionCountPerClientUsername, apiData.MaxConnectionCountPerClientUsername)
	AssignIfDstNotNil(&tfData.MaxEgressFlowCount, apiData.MaxEgressFlowCount)
	AssignIfDstNotNil(&tfData.MaxEndpointCountPerClientUsername, apiData.MaxEndpointCountPerClientUsername)
	AssignIfDstNotNil(&tfData.MaxIngressFlowCount, apiData.MaxIngressFlowCount)
	AssignIfDstNotNil(&tfData.MaxMsgsPerTransaction, apiData.MaxMsgsPerTransaction)
	AssignIfDstNotNil(&tfData.MaxSubscriptionCount, apiData.MaxSubscriptionCount)
	AssignIfDstNotNil(&tfData.MaxTransactedSessionCount, apiData.MaxTransactedSessionCount)
	AssignIfDstNotNil(&tfData.MaxTransactionCount, apiData.MaxTransactionCount)
	AssignIfDstNotNil(&tfData.MsgVpnName, apiData.MsgVpnName)
	AssignIfDstNotNil(&tfData.QueueControl1MaxDepth, apiData.QueueControl1MaxDepth)
	AssignIfDstNotNil(&tfData.QueueControl1MinMsgBurst, apiData.QueueControl1MinMsgBurst)
	AssignIfDstNotNil(&tfData.QueueDirect1MaxDepth, apiData.QueueDirect1MaxDepth)
	AssignIfDstNotNil(&tfData.QueueDirect1MinMsgBurst, apiData.QueueDirect1MinMsgBurst)
	AssignIfDstNotNil(&tfData.QueueDirect2MaxDepth, apiData.QueueDirect2MaxDepth)
	AssignIfDstNotNil(&tfData.QueueDirect2MinMsgBurst, apiData.QueueDirect2MinMsgBurst)
	AssignIfDstNotNil(&tfData.QueueDirect3MaxDepth, apiData.QueueDirect3MaxDepth)
	AssignIfDstNotNil(&tfData.QueueDirect3MinMsgBurst, apiData.QueueDirect3MinMsgBurst)
	AssignIfDstNotNil(&tfData.QueueGuaranteed1MaxDepth, apiData.QueueGuaranteed1MaxDepth)
	AssignIfDstNotNil(&tfData.QueueGuaranteed1MinMsgBurst, apiData.QueueGuaranteed1MinMsgBurst)
	AssignIfDstNotNil(&tfData.RejectMsgToSenderOnNoSubscriptionMatchEnabled, apiData.RejectMsgToSenderOnNoSubscriptionMatchEnabled)
	AssignIfDstNotNil(&tfData.ReplicationAllowClientConnectWhenStandbyEnabled, apiData.ReplicationAllowClientConnectWhenStandbyEnabled)
	AssignIfDstNotNil(&tfData.ServiceMinKeepaliveTimeout, apiData.ServiceMinKeepaliveTimeout)
	AssignIfDstNotNil(&tfData.ServiceSmfMaxConnectionCountPerClientUsername, apiData.ServiceSmfMaxConnectionCountPerClientUsername)
	AssignIfDstNotNil(&tfData.ServiceSmfMinKeepaliveEnabled, apiData.ServiceSmfMinKeepaliveEnabled)
	AssignIfDstNotNil(&tfData.ServiceWebInactiveTimeout, apiData.ServiceWebInactiveTimeout)
	AssignIfDstNotNil(&tfData.ServiceWebMaxConnectionCountPerClientUsername, apiData.ServiceWebMaxConnectionCountPerClientUsername)
	AssignIfDstNotNil(&tfData.ServiceWebMaxPayload, apiData.ServiceWebMaxPayload)
	AssignIfDstNotNil(&tfData.TcpCongestionWindowSize, apiData.TcpCongestionWindowSize)
	AssignIfDstNotNil(&tfData.TcpKeepaliveCount, apiData.TcpKeepaliveCount)
	AssignIfDstNotNil(&tfData.TcpKeepaliveIdleTime, apiData.TcpKeepaliveIdleTime)
	AssignIfDstNotNil(&tfData.TcpKeepaliveInterval, apiData.TcpKeepaliveInterval)
	AssignIfDstNotNil(&tfData.TcpMaxSegmentSize, apiData.TcpMaxSegmentSize)
	AssignIfDstNotNil(&tfData.TcpMaxWindowSize, apiData.TcpMaxWindowSize)
	AssignIfDstNotNil(&tfData.TlsAllowDowngradeToPlainTextEnabled, apiData.TlsAllowDowngradeToPlainTextEnabled)
}

func (tfData *MsgVpnClientProfile) ToApi() *sempv2.MsgVpnClientProfile {
	return &sempv2.MsgVpnClientProfile{
		AllowBridgeConnectionsEnabled:                            tfData.AllowBridgeConnectionsEnabled,
		AllowCutThroughForwardingEnabled:                         tfData.AllowCutThroughForwardingEnabled,
		AllowGuaranteedEndpointCreateDurability:                  tfData.AllowGuaranteedEndpointCreateDurability,
		AllowGuaranteedEndpointCreateEnabled:                     tfData.AllowGuaranteedEndpointCreateEnabled,
		AllowGuaranteedMsgReceiveEnabled:                         tfData.AllowGuaranteedMsgReceiveEnabled,
		AllowGuaranteedMsgSendEnabled:                            tfData.AllowGuaranteedMsgSendEnabled,
		AllowSharedSubscriptionsEnabled:                          tfData.AllowSharedSubscriptionsEnabled,
		AllowTransactedSessionsEnabled:                           tfData.AllowTransactedSessionsEnabled,
		ApiQueueManagementCopyFromOnCreateName:                   tfData.ApiQueueManagementCopyFromOnCreateName,
		ApiQueueManagementCopyFromOnCreateTemplateName:           tfData.ApiQueueManagementCopyFromOnCreateTemplateName,
		ApiTopicEndpointManagementCopyFromOnCreateName:           tfData.ApiTopicEndpointManagementCopyFromOnCreateName,
		ApiTopicEndpointManagementCopyFromOnCreateTemplateName:   tfData.ApiTopicEndpointManagementCopyFromOnCreateTemplateName,
		ClientProfileName:                                        tfData.ClientProfileName,
		CompressionEnabled:                                       tfData.CompressionEnabled,
		ElidingDelay:                                             tfData.ElidingDelay,
		ElidingEnabled:                                           tfData.ElidingEnabled,
		ElidingMaxTopicCount:                                     tfData.ElidingMaxTopicCount,
		EventClientProvisionedEndpointSpoolUsageThreshold:        tfData.EventClientProvisionedEndpointSpoolUsageThreshold.ToApi(),
		EventConnectionCountPerClientUsernameThreshold:           tfData.EventConnectionCountPerClientUsernameThreshold.ToApi(),
		EventEgressFlowCountThreshold:                            tfData.EventEgressFlowCountThreshold.ToApi(),
		EventEndpointCountPerClientUsernameThreshold:             tfData.EventEndpointCountPerClientUsernameThreshold.ToApi(),
		EventIngressFlowCountThreshold:                           tfData.EventIngressFlowCountThreshold.ToApi(),
		EventServiceSmfConnectionCountPerClientUsernameThreshold: tfData.EventServiceSmfConnectionCountPerClientUsernameThreshold.ToApi(),
		EventServiceWebConnectionCountPerClientUsernameThreshold: tfData.EventServiceWebConnectionCountPerClientUsernameThreshold.ToApi(),
		EventSubscriptionCountThreshold:                          tfData.EventSubscriptionCountThreshold.ToApi(),
		EventTransactedSessionCountThreshold:                     tfData.EventTransactedSessionCountThreshold.ToApi(),
		EventTransactionCountThreshold:                           tfData.EventTransactionCountThreshold.ToApi(),
		MaxConnectionCountPerClientUsername:                      tfData.MaxConnectionCountPerClientUsername,
		MaxEgressFlowCount:                                       tfData.MaxEgressFlowCount,
		MaxEndpointCountPerClientUsername:                        tfData.MaxEndpointCountPerClientUsername,
		MaxIngressFlowCount:                                      tfData.MaxIngressFlowCount,
		MaxMsgsPerTransaction:                                    tfData.MaxMsgsPerTransaction,
		MaxSubscriptionCount:                                     tfData.MaxSubscriptionCount,
		MaxTransactedSessionCount:                                tfData.MaxTransactedSessionCount,
		MaxTransactionCount:                                      tfData.MaxTransactionCount,
		MsgVpnName:                                               tfData.MsgVpnName,
		QueueControl1MaxDepth:                                    tfData.QueueControl1MaxDepth,
		QueueControl1MinMsgBurst:                                 tfData.QueueControl1MinMsgBurst,
		QueueDirect1MaxDepth:                                     tfData.QueueDirect1MaxDepth,
		QueueDirect1MinMsgBurst:                                  tfData.QueueDirect1MinMsgBurst,
		QueueDirect2MaxDepth:                                     tfData.QueueDirect2MaxDepth,
		QueueDirect2MinMsgBurst:                                  tfData.QueueDirect2MinMsgBurst,
		QueueDirect3MaxDepth:                                     tfData.QueueDirect3MaxDepth,
		QueueDirect3MinMsgBurst:                                  tfData.QueueDirect3MinMsgBurst,
		QueueGuaranteed1MaxDepth:                                 tfData.QueueGuaranteed1MaxDepth,
		QueueGuaranteed1MinMsgBurst:                              tfData.QueueGuaranteed1MinMsgBurst,
		RejectMsgToSenderOnNoSubscriptionMatchEnabled:            tfData.RejectMsgToSenderOnNoSubscriptionMatchEnabled,
		ReplicationAllowClientConnectWhenStandbyEnabled:          tfData.ReplicationAllowClientConnectWhenStandbyEnabled,
		ServiceMinKeepaliveTimeout:                               tfData.ServiceMinKeepaliveTimeout,
		ServiceSmfMaxConnectionCountPerClientUsername:            tfData.ServiceSmfMaxConnectionCountPerClientUsername,
		ServiceSmfMinKeepaliveEnabled:                            tfData.ServiceSmfMinKeepaliveEnabled,
		ServiceWebInactiveTimeout:                                tfData.ServiceWebInactiveTimeout,
		ServiceWebMaxConnectionCountPerClientUsername:            tfData.ServiceWebMaxConnectionCountPerClientUsername,
		ServiceWebMaxPayload:                                     tfData.ServiceWebMaxPayload,
		TcpCongestionWindowSize:                                  tfData.TcpCongestionWindowSize,
		TcpKeepaliveCount:                                        tfData.TcpKeepaliveCount,
		TcpKeepaliveIdleTime:                                     tfData.TcpKeepaliveIdleTime,
		TcpKeepaliveInterval:                                     tfData.TcpKeepaliveInterval,
		TcpMaxSegmentSize:                                        tfData.TcpMaxSegmentSize,
		TcpMaxWindowSize:                                         tfData.TcpMaxWindowSize,
		TlsAllowDowngradeToPlainTextEnabled:                      tfData.TlsAllowDowngradeToPlainTextEnabled,
	}
}

// Terraform Resource schema for MsgVpnClientProfile
func MsgVpnClientProfileResourceSchema(requiredAttributes ...string) schema.Schema {
	schema := schema.Schema{
		Description: "MsgVpnClientProfile",
		Attributes: map[string]schema.Attribute{
			"allow_bridge_connections_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing Bridge clients using the Client Profile to connect. Changing this setting does not affect existing Bridge client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "allow_bridge_connections_enabled"),
				Optional:    !contains(requiredAttributes, "allow_bridge_connections_enabled"),
			},
			"allow_cut_through_forwarding_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to bind to endpoints with the cut-through forwarding delivery mode. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.22. This attribute has been deprecated. Please visit the Solace Product Lifecycle Policy web page for details on deprecated features.",
				Required:    contains(requiredAttributes, "allow_cut_through_forwarding_enabled"),
				Optional:    !contains(requiredAttributes, "allow_cut_through_forwarding_enabled"),
			},
			"allow_guaranteed_endpoint_create_durability": schema.StringAttribute{
				Description: "The types of Queues and Topic Endpoints that clients using the client-profile can create. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"all\"`. The allowed values and their meaning are:  <pre> \"all\" - Client can create any type of endpoint. \"durable\" - Client can create only durable endpoints. \"non-durable\" - Client can create only non-durable endpoints. </pre>  Available since 2.14.",
				Required:    contains(requiredAttributes, "allow_guaranteed_endpoint_create_durability"),
				Optional:    !contains(requiredAttributes, "allow_guaranteed_endpoint_create_durability"),

				Validators: []validator.String{
					stringvalidator.OneOf("all", "durable", "non-durable"),
				},
				PlanModifiers: StringPlanModifiersFor("allow_guaranteed_endpoint_create_durability", requiredAttributes),
			},
			"allow_guaranteed_endpoint_create_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to create topic endponts or queues. Changing this value does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "allow_guaranteed_endpoint_create_enabled"),
				Optional:    !contains(requiredAttributes, "allow_guaranteed_endpoint_create_enabled"),
			},
			"allow_guaranteed_msg_receive_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to receive guaranteed messages. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "allow_guaranteed_msg_receive_enabled"),
				Optional:    !contains(requiredAttributes, "allow_guaranteed_msg_receive_enabled"),
			},
			"allow_guaranteed_msg_send_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to send guaranteed messages. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "allow_guaranteed_msg_send_enabled"),
				Optional:    !contains(requiredAttributes, "allow_guaranteed_msg_send_enabled"),
			},
			"allow_shared_subscriptions_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing shared subscriptions. Changing this setting does not affect existing subscriptions. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.11.",
				Required:    contains(requiredAttributes, "allow_shared_subscriptions_enabled"),
				Optional:    !contains(requiredAttributes, "allow_shared_subscriptions_enabled"),
			},
			"allow_transacted_sessions_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to establish transacted sessions. Changing this setting does not affect existing client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "allow_transacted_sessions_enabled"),
				Optional:    !contains(requiredAttributes, "allow_transacted_sessions_enabled"),
			},
			"api_queue_management_copy_from_on_create_name": schema.StringAttribute{
				Description: "The name of a queue to copy settings from when a new queue is created by a client using the Client Profile. The referenced queue must exist in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.14. This attribute has been replaced with `apiQueueManagementCopyFromOnCreateTemplateName`.",
				Required:    contains(requiredAttributes, "api_queue_management_copy_from_on_create_name"),
				Optional:    !contains(requiredAttributes, "api_queue_management_copy_from_on_create_name"),

				PlanModifiers: StringPlanModifiersFor("api_queue_management_copy_from_on_create_name", requiredAttributes),
			},
			"api_queue_management_copy_from_on_create_template_name": schema.StringAttribute{
				Description: "The name of a queue template to copy settings from when a new queue is created by a client using the Client Profile. If the referenced queue template does not exist, queue creation will fail when it tries to resolve this template. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since 2.14.",
				Required:    contains(requiredAttributes, "api_queue_management_copy_from_on_create_template_name"),
				Optional:    !contains(requiredAttributes, "api_queue_management_copy_from_on_create_template_name"),

				PlanModifiers: StringPlanModifiersFor("api_queue_management_copy_from_on_create_template_name", requiredAttributes),
			},
			"api_topic_endpoint_management_copy_from_on_create_name": schema.StringAttribute{
				Description: "The name of a topic endpoint to copy settings from when a new topic endpoint is created by a client using the Client Profile. The referenced topic endpoint must exist in the Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Deprecated since 2.14. This attribute has been replaced with `apiTopicEndpointManagementCopyFromOnCreateTemplateName`.",
				Required:    contains(requiredAttributes, "api_topic_endpoint_management_copy_from_on_create_name"),
				Optional:    !contains(requiredAttributes, "api_topic_endpoint_management_copy_from_on_create_name"),

				PlanModifiers: StringPlanModifiersFor("api_topic_endpoint_management_copy_from_on_create_name", requiredAttributes),
			},
			"api_topic_endpoint_management_copy_from_on_create_template_name": schema.StringAttribute{
				Description: "The name of a topic endpoint template to copy settings from when a new topic endpoint is created by a client using the Client Profile. If the referenced topic endpoint template does not exist, topic endpoint creation will fail when it tries to resolve this template. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`. Available since 2.14.",
				Required:    contains(requiredAttributes, "api_topic_endpoint_management_copy_from_on_create_template_name"),
				Optional:    !contains(requiredAttributes, "api_topic_endpoint_management_copy_from_on_create_template_name"),

				PlanModifiers: StringPlanModifiersFor("api_topic_endpoint_management_copy_from_on_create_template_name", requiredAttributes),
			},
			"client_profile_name": schema.StringAttribute{
				Description: "The name of the Client Profile.",
				Required:    contains(requiredAttributes, "client_profile_name"),
				Optional:    !contains(requiredAttributes, "client_profile_name"),

				PlanModifiers: StringPlanModifiersFor("client_profile_name", requiredAttributes),
			},
			"compression_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to use compression. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since 2.10.",
				Required:    contains(requiredAttributes, "compression_enabled"),
				Optional:    !contains(requiredAttributes, "compression_enabled"),
			},
			"eliding_delay": schema.Int64Attribute{
				Description: "The amount of time to delay the delivery of messages to clients using the Client Profile after the initial message has been delivered (the eliding delay interval), in milliseconds. A value of 0 means there is no delay in delivering messages to clients. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Required:    contains(requiredAttributes, "eliding_delay"),
				Optional:    !contains(requiredAttributes, "eliding_delay"),
			},
			"eliding_enabled": schema.BoolAttribute{
				Description: "Enable or disable message eliding for clients using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "eliding_enabled"),
				Optional:    !contains(requiredAttributes, "eliding_enabled"),
			},
			"eliding_max_topic_count": schema.Int64Attribute{
				Description: "The maximum number of topics tracked for message eliding per client connection using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`.",
				Required:    contains(requiredAttributes, "eliding_max_topic_count"),
				Optional:    !contains(requiredAttributes, "eliding_max_topic_count"),
			},
			"event_client_provisioned_endpoint_spool_usage_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_client_provisioned_endpoint_spool_usage_threshold"),
				Optional:    !contains(requiredAttributes, "event_client_provisioned_endpoint_spool_usage_threshold"),

				AttributeTypes: EventThresholdByPercentAttributeTypes,
			},
			"event_connection_count_per_client_username_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_connection_count_per_client_username_threshold"),
				Optional:    !contains(requiredAttributes, "event_connection_count_per_client_username_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_egress_flow_count_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_egress_flow_count_threshold"),
				Optional:    !contains(requiredAttributes, "event_egress_flow_count_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_endpoint_count_per_client_username_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_endpoint_count_per_client_username_threshold"),
				Optional:    !contains(requiredAttributes, "event_endpoint_count_per_client_username_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_ingress_flow_count_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_ingress_flow_count_threshold"),
				Optional:    !contains(requiredAttributes, "event_ingress_flow_count_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_service_smf_connection_count_per_client_username_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_service_smf_connection_count_per_client_username_threshold"),
				Optional:    !contains(requiredAttributes, "event_service_smf_connection_count_per_client_username_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_service_web_connection_count_per_client_username_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_service_web_connection_count_per_client_username_threshold"),
				Optional:    !contains(requiredAttributes, "event_service_web_connection_count_per_client_username_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_subscription_count_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_subscription_count_threshold"),
				Optional:    !contains(requiredAttributes, "event_subscription_count_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_transacted_session_count_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_transacted_session_count_threshold"),
				Optional:    !contains(requiredAttributes, "event_transacted_session_count_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"event_transaction_count_threshold": schema.ObjectAttribute{
				Description: "",
				Required:    contains(requiredAttributes, "event_transaction_count_threshold"),
				Optional:    !contains(requiredAttributes, "event_transaction_count_threshold"),

				AttributeTypes: EventThresholdAttributeTypes,
			},
			"max_connection_count_per_client_username": schema.Int64Attribute{
				Description: "The maximum number of client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Required:    contains(requiredAttributes, "max_connection_count_per_client_username"),
				Optional:    !contains(requiredAttributes, "max_connection_count_per_client_username"),
			},
			"max_egress_flow_count": schema.Int64Attribute{
				Description: "The maximum number of transmit flows that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Required:    contains(requiredAttributes, "max_egress_flow_count"),
				Optional:    !contains(requiredAttributes, "max_egress_flow_count"),
			},
			"max_endpoint_count_per_client_username": schema.Int64Attribute{
				Description: "The maximum number of queues and topic endpoints that can be created by clients with the same Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Required:    contains(requiredAttributes, "max_endpoint_count_per_client_username"),
				Optional:    !contains(requiredAttributes, "max_endpoint_count_per_client_username"),
			},
			"max_ingress_flow_count": schema.Int64Attribute{
				Description: "The maximum number of receive flows that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Required:    contains(requiredAttributes, "max_ingress_flow_count"),
				Optional:    !contains(requiredAttributes, "max_ingress_flow_count"),
			},
			"max_msgs_per_transaction": schema.Int64Attribute{
				Description: "The maximum number of publisher and consumer messages combined that is allowed within a transaction for each client associated with this client-profile. Exceeding this limit will result in a transaction prepare or commit failure. Changing this value during operation will not affect existing sessions. It is only validated at transaction creation time. Large transactions consume more resources and are more likely to require retrieving messages from the ADB or from disk to process the transaction prepare or commit requests. The transaction processing rate may diminish if a large number of messages must be retrieved from the ADB or from disk. Care should be taken to not use excessively large transactions needlessly to avoid exceeding resource limits and to avoid reducing the overall broker performance. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`. Available since 2.20.",
				Required:    contains(requiredAttributes, "max_msgs_per_transaction"),
				Optional:    !contains(requiredAttributes, "max_msgs_per_transaction"),
			},
			"max_subscription_count": schema.Int64Attribute{
				Description: "The maximum number of subscriptions per client using the Client Profile. This limit is not enforced when a client adds a subscription to an endpoint, except for MQTT QoS 1 subscriptions. In addition, this limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default varies by platform.",
				Required:    contains(requiredAttributes, "max_subscription_count"),
				Optional:    !contains(requiredAttributes, "max_subscription_count"),
			},
			"max_transacted_session_count": schema.Int64Attribute{
				Description: "The maximum number of transacted sessions that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10`.",
				Required:    contains(requiredAttributes, "max_transacted_session_count"),
				Optional:    !contains(requiredAttributes, "max_transacted_session_count"),
			},
			"max_transaction_count": schema.Int64Attribute{
				Description: "The maximum number of transactions that can be created by one client using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default varies by platform.",
				Required:    contains(requiredAttributes, "max_transaction_count"),
				Optional:    !contains(requiredAttributes, "max_transaction_count"),
			},
			"msg_vpn_name": schema.StringAttribute{
				Description: "The name of the Message VPN.",
				Required:    contains(requiredAttributes, "msg_vpn_name"),
				Optional:    !contains(requiredAttributes, "msg_vpn_name"),

				PlanModifiers: StringPlanModifiersFor("msg_vpn_name", requiredAttributes),
			},
			"queue_control1_max_depth": schema.Int64Attribute{
				Description: "The maximum depth of the \"Control 1\" (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Required:    contains(requiredAttributes, "queue_control1_max_depth"),
				Optional:    !contains(requiredAttributes, "queue_control1_max_depth"),
			},
			"queue_control1_min_msg_burst": schema.Int64Attribute{
				Description: "The number of messages that are always allowed entry into the \"Control 1\" (C-1) priority queue, regardless of the `queueControl1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Required:    contains(requiredAttributes, "queue_control1_min_msg_burst"),
				Optional:    !contains(requiredAttributes, "queue_control1_min_msg_burst"),
			},
			"queue_direct1_max_depth": schema.Int64Attribute{
				Description: "The maximum depth of the \"Direct 1\" (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Required:    contains(requiredAttributes, "queue_direct1_max_depth"),
				Optional:    !contains(requiredAttributes, "queue_direct1_max_depth"),
			},
			"queue_direct1_min_msg_burst": schema.Int64Attribute{
				Description: "The number of messages that are always allowed entry into the \"Direct 1\" (D-1) priority queue, regardless of the `queueDirect1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Required:    contains(requiredAttributes, "queue_direct1_min_msg_burst"),
				Optional:    !contains(requiredAttributes, "queue_direct1_min_msg_burst"),
			},
			"queue_direct2_max_depth": schema.Int64Attribute{
				Description: "The maximum depth of the \"Direct 2\" (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Required:    contains(requiredAttributes, "queue_direct2_max_depth"),
				Optional:    !contains(requiredAttributes, "queue_direct2_max_depth"),
			},
			"queue_direct2_min_msg_burst": schema.Int64Attribute{
				Description: "The number of messages that are always allowed entry into the \"Direct 2\" (D-2) priority queue, regardless of the `queueDirect2MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Required:    contains(requiredAttributes, "queue_direct2_min_msg_burst"),
				Optional:    !contains(requiredAttributes, "queue_direct2_min_msg_burst"),
			},
			"queue_direct3_max_depth": schema.Int64Attribute{
				Description: "The maximum depth of the \"Direct 3\" (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Required:    contains(requiredAttributes, "queue_direct3_max_depth"),
				Optional:    !contains(requiredAttributes, "queue_direct3_max_depth"),
			},
			"queue_direct3_min_msg_burst": schema.Int64Attribute{
				Description: "The number of messages that are always allowed entry into the \"Direct 3\" (D-3) priority queue, regardless of the `queueDirect3MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `4`.",
				Required:    contains(requiredAttributes, "queue_direct3_min_msg_burst"),
				Optional:    !contains(requiredAttributes, "queue_direct3_min_msg_burst"),
			},
			"queue_guaranteed1_max_depth": schema.Int64Attribute{
				Description: "The maximum depth of the \"Guaranteed 1\" (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `20000`.",
				Required:    contains(requiredAttributes, "queue_guaranteed1_max_depth"),
				Optional:    !contains(requiredAttributes, "queue_guaranteed1_max_depth"),
			},
			"queue_guaranteed1_min_msg_burst": schema.Int64Attribute{
				Description: "The number of messages that are always allowed entry into the \"Guaranteed 1\" (G-3) priority queue, regardless of the `queueGuaranteed1MaxDepth` value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `255`.",
				Required:    contains(requiredAttributes, "queue_guaranteed1_min_msg_burst"),
				Optional:    !contains(requiredAttributes, "queue_guaranteed1_min_msg_burst"),
			},
			"reject_msg_to_sender_on_no_subscription_match_enabled": schema.BoolAttribute{
				Description: "Enable or disable the sending of a negative acknowledgement (NACK) to a client using the Client Profile when discarding a guaranteed message due to no matching subscription found. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.2.",
				Required:    contains(requiredAttributes, "reject_msg_to_sender_on_no_subscription_match_enabled"),
				Optional:    !contains(requiredAttributes, "reject_msg_to_sender_on_no_subscription_match_enabled"),
			},
			"replication_allow_client_connect_when_standby_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing clients using the Client Profile to connect to the Message VPN when its replication state is standby. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Required:    contains(requiredAttributes, "replication_allow_client_connect_when_standby_enabled"),
				Optional:    !contains(requiredAttributes, "replication_allow_client_connect_when_standby_enabled"),
			},
			"service_min_keepalive_timeout": schema.Int64Attribute{
				Description: "The minimum client keepalive timeout which will be enforced for client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30`. Available since 2.19.",
				Required:    contains(requiredAttributes, "service_min_keepalive_timeout"),
				Optional:    !contains(requiredAttributes, "service_min_keepalive_timeout"),
			},
			"service_smf_max_connection_count_per_client_username": schema.Int64Attribute{
				Description: "The maximum number of SMF client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Required:    contains(requiredAttributes, "service_smf_max_connection_count_per_client_username"),
				Optional:    !contains(requiredAttributes, "service_smf_max_connection_count_per_client_username"),
			},
			"service_smf_min_keepalive_enabled": schema.BoolAttribute{
				Description: "Enable or disable the enforcement of a minimum keepalive timeout for SMF clients. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.19.",
				Required:    contains(requiredAttributes, "service_smf_min_keepalive_enabled"),
				Optional:    !contains(requiredAttributes, "service_smf_min_keepalive_enabled"),
			},
			"service_web_inactive_timeout": schema.Int64Attribute{
				Description: "The timeout for inactive Web Transport client sessions using the Client Profile, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30`.",
				Required:    contains(requiredAttributes, "service_web_inactive_timeout"),
				Optional:    !contains(requiredAttributes, "service_web_inactive_timeout"),
			},
			"service_web_max_connection_count_per_client_username": schema.Int64Attribute{
				Description: "The maximum number of Web Transport client connections per Client Username using the Client Profile. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default is the maximum value supported by the platform.",
				Required:    contains(requiredAttributes, "service_web_max_connection_count_per_client_username"),
				Optional:    !contains(requiredAttributes, "service_web_max_connection_count_per_client_username"),
			},
			"service_web_max_payload": schema.Int64Attribute{
				Description: "The maximum Web Transport payload size before fragmentation occurs for clients using the Client Profile, in bytes. The size of the header is not included. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000000`.",
				Required:    contains(requiredAttributes, "service_web_max_payload"),
				Optional:    !contains(requiredAttributes, "service_web_max_payload"),
			},
			"tcp_congestion_window_size": schema.Int64Attribute{
				Description: "The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `2`.",
				Required:    contains(requiredAttributes, "tcp_congestion_window_size"),
				Optional:    !contains(requiredAttributes, "tcp_congestion_window_size"),
			},
			"tcp_keepalive_count": schema.Int64Attribute{
				Description: "The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5`.",
				Required:    contains(requiredAttributes, "tcp_keepalive_count"),
				Optional:    !contains(requiredAttributes, "tcp_keepalive_count"),
			},
			"tcp_keepalive_idle_time": schema.Int64Attribute{
				Description: "The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Required:    contains(requiredAttributes, "tcp_keepalive_idle_time"),
				Optional:    !contains(requiredAttributes, "tcp_keepalive_idle_time"),
			},
			"tcp_keepalive_interval": schema.Int64Attribute{
				Description: "The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Required:    contains(requiredAttributes, "tcp_keepalive_interval"),
				Optional:    !contains(requiredAttributes, "tcp_keepalive_interval"),
			},
			"tcp_max_segment_size": schema.Int64Attribute{
				Description: "The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1460`.",
				Required:    contains(requiredAttributes, "tcp_max_segment_size"),
				Optional:    !contains(requiredAttributes, "tcp_max_segment_size"),
			},
			"tcp_max_window_size": schema.Int64Attribute{
				Description: "The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `256`.",
				Required:    contains(requiredAttributes, "tcp_max_window_size"),
				Optional:    !contains(requiredAttributes, "tcp_max_window_size"),
			},
			"tls_allow_downgrade_to_plain_text_enabled": schema.BoolAttribute{
				Description: "Enable or disable allowing a client using the Client Profile to downgrade an encrypted connection to plain text. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since 2.8.",
				Required:    contains(requiredAttributes, "tls_allow_downgrade_to_plain_text_enabled"),
				Optional:    !contains(requiredAttributes, "tls_allow_downgrade_to_plain_text_enabled"),
			},
		},
	}

	return schema
}
