/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package doctl

const (
	// ArgAccessToken is the access token to be used for the operations
	ArgAccessToken = "access-token"
	// ArgContext is the name of the auth context
	ArgContext = "context"
	// ArgDefaultContext is the default auth context
	ArgDefaultContext = "default"
	// ArgActionID is an action id argument.
	ArgActionID = "action-id"
	// ArgActionAfter is an action after argument.
	ArgActionAfter = "after"
	// ArgActionBefore is an action before argument.
	ArgActionBefore = "before"
	// ArgActionResourceType is an action resource type argument.
	ArgActionResourceType = "resource-type"
	// ArgActionRegion is an action region argument.
	ArgActionRegion = "region"
	// ArgActionStatus is an action status argument.
	ArgActionStatus = "status"
	// ArgActionType is an action type argument.
	ArgActionType = "action-type"
	// ArgApp is the app ID.
	ArgApp = "app"
	// ArgAppWithProjects will determine whether project ids should be fetched along with listed apps.
	ArgAppWithProjects = "with-projects"
	// ArgAppSpec is a path to an app spec.
	ArgAppSpec = "spec"
	// ArgAppLogType the type of log.
	ArgAppLogType = "type"
	// ArgAppDeployment is the deployment ID.
	ArgAppDeployment = "deployment"
	// ArgAppInstanceName is the instance name of currently running instances (optional).
	ArgAppInstanceName = "instance-name"
	// ArgAppDevConfig is the path to the app dev link config.
	ArgAppDevConfig = "dev-config"
	// ArgBuildCommand is an optional build command to set for local development.
	ArgBuildCommand = "build-command"
	// ArgBuildpack is a buildpack id.
	ArgBuildpack = "buildpack"
	// ArgAppLogFollow follow logs.
	ArgAppLogFollow = "follow"
	// ArgAppLogTail tail logs.
	ArgAppLogTail = "tail"
	// ArgNoPrefix no prefix to json logs
	ArgNoPrefix = "no-prefix"
	// ArgAppForceRebuild forces a deployment rebuild
	ArgAppForceRebuild = "force-rebuild"
	// ArgAppComponents is a list of components to restart.
	ArgAppComponents = "components"
	// ArgAppAlertDestinations is a path to an app alert destination file.
	ArgAppAlertDestinations = "app-alert-destinations"
	// ArgClusterName is a cluster name argument.
	ArgClusterName = "cluster-name"
	// ArgClusterVersionSlug is a cluster version argument.
	ArgClusterVersionSlug = "version"
	// ArgVPCUUID is a VPC UUID argument.
	ArgVPCUUID = "vpc-uuid"
	// ArgClusterVPCUUID is a cluster vpc-uuid argument.
	ArgClusterVPCUUID = "vpc-uuid"
	// ArgClusterSubnet is a cluster pod CIDR argument.
	ArgClusterSubnet = "cluster-subnet"
	// ArgServiceSubnet is a cluster service CIDR argument.
	ArgServiceSubnet = "service-subnet"
	// ArgClusterNodePool are a cluster's node pools arguments.
	ArgClusterNodePool = "node-pool"
	// ArgClusterUpdateKubeconfig updates the local kubeconfig.
	ArgClusterUpdateKubeconfig = "update-kubeconfig"
	// ArgNoCache represents whether or not to omit the cache on the next command.
	ArgNoCache = "no-cache"
	// ArgNodePoolName is a cluster's node pool name argument.
	ArgNodePoolName = "name"
	// ArgNodePoolCount is a cluster's node pool count argument.
	ArgNodePoolCount = "count"
	// ArgNodePoolAutoScale is a cluster's node pool auto_scale argument.
	ArgNodePoolAutoScale = "auto-scale"
	// ArgNodePoolMinNodes is a cluster's node pool min_nodes argument.
	ArgNodePoolMinNodes = "min-nodes"
	// ArgNodePoolMaxNodes is a cluster's node pool max_nodes argument.
	ArgNodePoolMaxNodes = "max-nodes"
	// ArgNodePoolNodeIDs is a cluster's node pool nodes argument.
	ArgNodePoolNodeIDs = "node-ids"
	// ArgMaintenanceWindow is a cluster's maintenance window argument
	ArgMaintenanceWindow = "maintenance-window"
	// ArgMajorVersion is a major version number.
	ArgMajorVersion = "major-version"
	// ArgAutoUpgrade is a cluster's auto-upgrade argument.
	ArgAutoUpgrade = "auto-upgrade"

	// ArgAutoscaleID is an autoscale id argument.
	ArgAutoscaleID = "id"
	// ArgAutoscaleName is an autoscale name argument.
	ArgAutoscaleName = "name"
	// ArgAutoscaleMinInstances is an autoscale min instance argument.
	ArgAutoscaleMinInstances = "min-instances"
	// ArgAutoscaleMaxInstances is an autoscale max instance argument.
	ArgAutoscaleMaxInstances = "max-instances"
	// ArgAutoscaleCpuTarget is an autoscale target cpu utilization argument.
	ArgAutoscaleCpuTarget = "cpu-target"
	// ArgAutoscaleMemTarget is an autoscale target memory utilization argument.
	ArgAutoscaleMemTarget = "mem-target"
	// ArgAutoscaleCooldownMinutes is an autoscale cooldown duration (minutes) argument.
	ArgAutoscaleCooldownMinutes = "cooldown-minutes"
	// ArgAutoscaleTargetInstances is an autoscale target instance argument.
	ArgAutoscaleTargetInstances = "target-instances"

	// ArgVPCNATGatewayName is a vpc nat gateway name argument.
	ArgVPCNATGatewayName = "name"
	// ArgVPCNATGatewayType is a vpc nat gateway type argument.
	ArgVPCNATGatewayType = "type"
	// ArgVPCNATGatewayRegion is a vpc nat gateway region argument.
	ArgVPCNATGatewayRegion = "region"
	// ArgVPCNATGatewaySize is a vpc nat gateway region argument.
	ArgVPCNATGatewaySize = "size"
	// ArgVPCNATGatewayVPCs is a vpc nat gateway vpcs argument.
	ArgVPCNATGatewayVPCs = "vpcs"
	// ArgVPCNATGatewayUDPTimeout is a vpc nat gateway udp-timeout argument.
	ArgVPCNATGatewayUDPTimeout = "udp-timeout"
	// ArgVPCNATGatewayICMPTimeout is a vpc nat gateway icmp-timeout argument.
	ArgVPCNATGatewayICMPTimeout = "icmp-timeout"
	// ArgVPCNATGatewayTCPTimeout is a vpc nat gateway tcp-timeout argument.
	ArgVPCNATGatewayTCPTimeout = "tcp-timeout"

	// ArgHA is a cluster's highly available control plane argument.
	ArgHA = "ha"
	// ArgEnableControlPlaneFirewall enable control plane firewall.
	ArgEnableControlPlaneFirewall = "enable-control-plane-firewall"
	// ArgControlPlaneFirewallAllowedAddresses list of allowed addresses that can access the control plane.
	ArgControlPlaneFirewallAllowedAddresses = "control-plane-firewall-allowed-addresses"
	// ArgClusterAutoscalerScaleDownUtilizationThreshold is the cluster autoscaler scale down utilization threshold
	ArgClusterAutoscalerScaleDownUtilizationThreshold = "scale-down-utilization-threshold"
	// ArgClusterAutoscalerScaleDownUnneededTime is the cluster autoscaler scale down unneeded time
	ArgClusterAutoscalerScaleDownUnneededTime = "scale-down-unneeded-time"
	// ArgClusterAutoscalerExpanders customizes the expanders used by the cluster autoscaler to scale up the cluster.
	ArgClusterAutoscalerExpanders = "expanders"
	// ArgEnableRoutingAgent enables the routing-agent cluster plugin.
	ArgEnableRoutingAgent = "enable-routing-agent"
	// ArgSurgeUpgrade is a cluster's surge-upgrade argument.
	ArgSurgeUpgrade = "surge-upgrade"
	// ArgCommandUpsert is an upsert for a resource to be created or updated argument.
	ArgCommandUpsert = "upsert"
	// ArgCommandUpdateSources tells the respective operation to also update the underlying sources.
	ArgCommandUpdateSources = "update-sources"
	// ArgCommandWait is a wait for a resource to be created argument.
	ArgCommandWait = "wait"
	// ArgSetCurrentContext is a flag to set the new kubeconfig context as current.
	ArgSetCurrentContext = "set-current-context"
	// ArgDropletID is a droplet id argument.
	ArgDropletID = "droplet-id"
	// ArgDropletIDs is a list of droplet IDs.
	ArgDropletIDs = "droplet-ids"
	// ArgKernelID is a kernel id argument.
	ArgKernelID = "kernel-id"
	// ArgKubernetesLabel is a Kubernetes label argument.
	ArgKubernetesLabel = "label"
	// ArgKubernetesTaint is a Kubernetes taint argument.
	ArgKubernetesTaint = "taint"
	// ArgKubernetesAlias is a Kubernetes alias argument that saves authentication information under the specified context.
	ArgKubernetesAlias = "alias"
	// ArgKubeConfigExpirySeconds indicates the length of time the token in a kubeconfig will be valid in seconds.
	ArgKubeConfigExpirySeconds = "expiry-seconds"
	// ArgImage is an image argument.
	ArgImage = "image"
	// ArgImageID is an image id argument.
	ArgImageID = "image-id"
	// ArgImagePublic is a public image argument.
	ArgImagePublic = "public"
	// ArgImageSlug is an image slug argument.
	ArgImageSlug = "image-slug"
	// ArgInteractive is the argument to enable an interactive CLI.
	ArgInteractive = "interactive"
	// ArgIPAddress is an IP address argument.
	ArgIPAddress = "ip-address"
	// ArgDropletName is a droplet name argument.
	ArgDropletName = "droplet-name"
	// ArgEnvFile is an environment file to load variables from.
	ArgEnvFile = "env-file"
	// ArgResizeDisk is a resize disk argument.
	ArgResizeDisk = "resize-disk"
	// ArgSnapshotName is a snapshot name argument.
	ArgSnapshotName = "snapshot-name"
	// ArgSnapshotDesc is the description for volume snapshot.
	ArgSnapshotDesc = "snapshot-desc"
	// ArgResourceType is the resource type for snapshot.
	ArgResourceType = "resource"
	// ArgBackups is an enable backups argument.
	ArgBackups = "enable-backups"
	// ArgDropletBackupPolicyPlan sets a frequency plan for backups.
	ArgDropletBackupPolicyPlan = "backup-policy-plan"
	// ArgDropletBackupPolicyWeekday sets backup policy day of the week.
	ArgDropletBackupPolicyWeekday = "backup-policy-weekday"
	// ArgDropletBackupPolicyHour sets backup policy hour.
	ArgDropletBackupPolicyHour = "backup-policy-hour"
	// ArgIPv6 is an enable IPv6 argument.
	ArgIPv6 = "enable-ipv6"
	// ArgPrivateNetworking is an enable private networking argument.
	ArgPrivateNetworking = "enable-private-networking"
	// ArgMonitoring is an enable monitoring argument.
	ArgMonitoring = "enable-monitoring"
	// ArgDropletAgent is an argument for enabling/disabling the Droplet agent.
	ArgDropletAgent = "droplet-agent"
	// ArgRecordData is a record data argument.
	ArgRecordData = "record-data"
	// ArgRecordID is a record id argument.
	ArgRecordID = "record-id"
	// ArgRecordName is a record name argument.
	ArgRecordName = "record-name"
	// ArgRecordPort is a record port argument.
	ArgRecordPort = "record-port"
	// ArgRecordPriority is a record priority argument.
	ArgRecordPriority = "record-priority"
	// ArgRecordType is a record type argument.
	ArgRecordType = "record-type"
	// ArgRecordTTL is a record ttl argument.
	ArgRecordTTL = "record-ttl"
	// ArgRecordWeight is a record weight argument.
	ArgRecordWeight = "record-weight"
	// ArgRecordFlags is a record flags argument.
	ArgRecordFlags = "record-flags"
	// ArgRecordTag is a record tag argument.
	ArgRecordTag = "record-tag"
	// ArgRegionSlug is a region slug argument.
	ArgRegionSlug = "region"
	// ArgSchemaOnly is a schema only argument.
	ArgSchemaOnly = "schema-only"
	// ArgSizeSlug is a size slug argument.
	ArgSizeSlug = "size"
	// ArgSizeUnit is a size unit argument.
	ArgSizeUnit = "size-unit"
	// ArgsSSHKeyPath is a ssh argument.
	ArgsSSHKeyPath = "ssh-key-path"
	// ArgSSHKeys is a ssh key argument.
	ArgSSHKeys = "ssh-keys"
	// ArgsSSHPort is a ssh argument.
	ArgsSSHPort = "ssh-port"
	// ArgsSSHAgentForwarding is a ssh argument.
	ArgsSSHAgentForwarding = "ssh-agent-forwarding"
	// ArgsSSHPrivateIP is a ssh argument.
	ArgsSSHPrivateIP = "ssh-private-ip"
	// ArgSSHCommand is a ssh argument.
	ArgSSHCommand = "ssh-command"
	// ArgSSHRetryMax is a ssh argument.
	ArgSSHRetryMax = "ssh-retry-max"
	// ArgUserData is a user data argument.
	ArgUserData = "user-data"
	// ArgUserDataFile is a user data file location argument.
	ArgUserDataFile = "user-data-file"
	// ArgImageName name is an image name argument.
	ArgImageName = "image-name"
	// ArgImageExternalURL is a URL that returns an image file.
	ArgImageExternalURL = "image-url"
	// ArgImageDistro is the name of a custom image's distribution
	ArgImageDistro = "image-distribution"
	// ArgImageDescription is free text that describes the image.
	ArgImageDescription = "image-description"
	// ArgKey is a key argument.
	ArgKey = "key"
	// ArgKeyName is a key name argument.
	ArgKeyName = "key-name"
	// ArgKeyPublicKey is a public key argument.
	ArgKeyPublicKey = "public-key"
	// ArgKeyPublicKeyFile is a public key file argument.
	ArgKeyPublicKeyFile = "public-key-file"
	// ArgSSHUser is a SSH user argument.
	ArgSSHUser = "ssh-user"
	// ArgFormat is columns to include in output argument.
	ArgFormat = "format"
	// ArgNoHeader hides the output header.
	ArgNoHeader = "no-header"
	// ArgPollTime is how long before the next poll argument.
	ArgPollTime = "poll-timeout"
	// ArgTagName is a tag name
	// NOTE: ArgTagName will be deprecated once existing uses have been migrated
	// to use `--tag` (ArgTag). ArgTagName should not be used on new calls.
	ArgTagName = "tag-name"
	// ArgTagNames is a slice of possible tag names
	// NOTE: ArgTagNames will be deprecated once existing uses have been migrated
	// to use `--tag` (ArgTag). ArgTagNames should not be used on new calls.
	ArgTagNames = "tag-names"
	// ArgTag specifies tag.  --tag can be repeated or multiple tags can be , separated.
	ArgTag = "tag"
	//ArgTemplate is template format
	ArgTemplate = "template"
	// ArgTimeout is a timeout duration
	ArgTimeout = "timeout"
	// ArgTriggerDeployment indicates whether to trigger a deployment
	ArgTriggerDeployment = "trigger-deployment"
	// ArgVersion is the version of the command to use
	ArgVersion = "version"
	// ArgVerbose enables verbose output
	ArgVerbose = "verbose"

	// ArgOutput is an output type argument.
	ArgOutput = "output"

	// ArgUptimeCheckName is the name of an uptime check.
	ArgUptimeCheckName = "name"
	// ArgUptimeCheckType is the type of an uptime check.
	ArgUptimeCheckType = "type"
	// ArgUptimeCheckTarget is the target of an uptime check.
	ArgUptimeCheckTarget = "target"
	// ArgUptimeCheckRegions are the regions of an uptime check.
	ArgUptimeCheckRegions = "regions"
	// ArgUptimeCheckEnabled is whether or not an uptime check is enabled.
	ArgUptimeCheckEnabled = "enabled"

	// ArgUptimeAlertName is the name of an uptime alert.
	ArgUptimeAlertName = "name"
	// ArgUptimeAlertType is the type of an uptime alert.
	ArgUptimeAlertType = "type"
	// ArgUptimeAlertThreshold the threshold at which an uptime alert will trigger.
	ArgUptimeAlertThreshold = "threshold"
	// ArgUptimeAlertComparison is the uptime alert comparator.
	ArgUptimeAlertComparison = "comparison"
	// ArgUptimeAlertEmails are the emails to send uptime alerts to.
	ArgUptimeAlertEmails = "emails"
	// ArgUptimeAlertSlackChannels are the Slack channels to send uptime alerts to.
	ArgUptimeAlertSlackChannels = "slack-channels"
	// ArgUptimeAlertSlackURLs are the Slack URLs to send uptime alerts to.
	ArgUptimeAlertSlackURLs = "slack-urls"
	// ArgUptimeAlertPeriod is the time threshold at which an uptime alert will trigger.
	ArgUptimeAlertPeriod = "period"

	// ArgVolumeSize is the size of a volume.
	ArgVolumeSize = "size"
	// ArgVolumeDesc is the description of a volume.
	ArgVolumeDesc = "desc"
	// ArgVolumeRegion is the region of a volume.
	ArgVolumeRegion = "region"
	// ArgVolumeSnapshot is the snapshot from which to create a volume.
	ArgVolumeSnapshot = "snapshot"
	// ArgVolumeFilesystemType is the filesystem type for a volume.
	ArgVolumeFilesystemType = "fs-type"
	// ArgVolumeFilesystemLabel is the filesystem label for a volume.
	ArgVolumeFilesystemLabel = "fs-label"
	// ArgVolumeList is the IDs of many volumes.
	ArgVolumeList = "volumes"
	// ArgVolumeSnapshotList is the IDs of many volume snapshots.
	ArgVolumeSnapshotList = "snapshots"
	// ArgLoadBalancerList is the IDs of many load balancers.
	ArgLoadBalancerList = "load-balancers"

	// ArgCDNTTL is a cdn ttl argument
	ArgCDNTTL = "ttl"
	// ArgCDNDomain is a cdn custom domain argument
	ArgCDNDomain = "domain"
	// ArgCDNCertificateID is a certificate id to use with a custom domain
	ArgCDNCertificateID = "certificate-id"
	// ArgCDNFiles is a cdn files argument
	ArgCDNFiles = "files"

	// ArgCertificateName is a name of the certificate.
	ArgCertificateName = "name"
	// ArgCertificateDNSNames is a list of DNS names.
	ArgCertificateDNSNames = "dns-names"
	// ArgPrivateKeyPath is a path to a private key for the certificate.
	ArgPrivateKeyPath = "private-key-path"
	// ArgLeafCertificatePath is a path to a certificate leaf.
	ArgLeafCertificatePath = "leaf-certificate-path"
	// ArgCertificateChainPath is a path to a certificate chain.
	ArgCertificateChainPath = "certificate-chain-path"
	// ArgCertificateType is a certificate type.
	ArgCertificateType = "type"

	// ArgLoadBalancerName is a name of the load balancer.
	ArgLoadBalancerName = "name"
	// ArgLoadBalancerAlgorithm is a load balancing algorithm.
	ArgLoadBalancerAlgorithm = "algorithm"
	// ArgRedirectHTTPToHTTPS is a flag that indicates whether HTTP requests to the load balancer on port 80 should be redirected to HTTPS on port 443.
	ArgRedirectHTTPToHTTPS = "redirect-http-to-https"
	// ArgEnableProxyProtocol is a flag that indicates whether PROXY protocol should be enabled on the load balancer.
	ArgEnableProxyProtocol = "enable-proxy-protocol"
	// ArgDisableLetsEncryptDNSRecords is a flag that when set will disable the creation of DNS records pointing to the load balancer IP from the apex domain in the cert.
	ArgDisableLetsEncryptDNSRecords = "disable-lets-encrypt-dns-records"
	// ArgEnableBackendKeepalive is a flag that indicates whether keepalive connections should be enabled to target droplets from the load balancer.
	ArgEnableBackendKeepalive = "enable-backend-keepalive"
	// ArgStickySessions is a list of sticky sessions settings for the load balancer.
	ArgStickySessions = "sticky-sessions"
	// ArgHealthCheck is a list of health check settings for the load balancer.
	ArgHealthCheck = "health-check"
	// ArgForwardingRules is a list of forwarding rules for the load balancer.
	ArgForwardingRules = "forwarding-rules"
	// ArgHTTPIdleTimeoutSeconds is the http idle time out configuration for the load balancer
	ArgHTTPIdleTimeoutSeconds = "http-idle-timeout-seconds"
	// ArgAllowList is the list of firewall rules for ALLOWING traffic to the loadbalancer
	ArgAllowList = "allow-list"
	// ArgDenyList is a list of firewall rules for DENYING traffic to the loadbalancer
	ArgDenyList = "deny-list"
	// ArgLoadBalancerType is the type of the load balancer.
	ArgLoadBalancerType = "type"
	// ArgLoadBalancerDomains is list of domains supported for global load balancer.
	ArgLoadBalancerDomains = "domains"
	// ArgGlobalLoadBalancerSettings is global load balancer settings.
	ArgGlobalLoadBalancerSettings = "glb-settings"
	// ArgGlobalLoadBalancerCDNSettings is global load balancer CDN settings.
	ArgGlobalLoadBalancerCDNSettings = "glb-cdn-settings"
	// ArgTargetLoadBalancerIDs is a list of target load balancer IDs.
	ArgTargetLoadBalancerIDs = "target-lb-ids"
	// ArgLoadBalancerNetwork is the type of network the load balancer is accessible from.
	ArgLoadBalancerNetwork = "network"
	// ArgLoadBalancerNetworkStack is the network stack type the load balancer will be configured with (e.g IPv4, Dual Stack: IPv4 and IPv6).
	ArgLoadBalancerNetworkStack = "network-stack"
	// ArgLoadBalancerTLSCipherPolicy is the tls cipher policy to be used for the load balancer
	ArgLoadBalancerTLSCipherPolicy = "tls-cipher-policy"

	// ArgFirewallName is a name of the firewall.
	ArgFirewallName = "name"
	// ArgInboundRules is a list of inbound rules for the firewall.
	ArgInboundRules = "inbound-rules"
	// ArgOutboundRules is a list of outbound rules for the firewall.
	ArgOutboundRules = "outbound-rules"

	// ArgProjectID is the ID of a project.
	ArgProjectID = "project-id"
	// ArgProjectName is the name of a project.
	ArgProjectName = "name"
	// ArgProjectDescription is the description of a project.
	ArgProjectDescription = "description"
	// ArgProjectPurpose is the purpose of a project.
	ArgProjectPurpose = "purpose"
	// ArgProjectEnvironment is the environment of a project. Should be one of 'Development', 'Staging', 'Production'.
	ArgProjectEnvironment = "environment"
	// ArgProjectIsDefault is used to change the default project.
	ArgProjectIsDefault = "is_default"
	// ArgProjectResource is a flag for your resource URNs
	ArgProjectResource = "resource"

	// ArgDatabaseRestoreFromClusterName is a flag for specifying the name of an existing database cluster from which the backup will be restored.
	ArgDatabaseRestoreFromClusterName = "restore-from-cluster-name"
	// ArgDatabaseRestoreFromClusterID is a flag for specifying the id of an existing database cluster from which the new database will be forked from.
	ArgDatabaseRestoreFromClusterID = "restore-from-cluster-id"
	// ArgDatabaseRestoreFromTimestamp is a flag for specifying the timestamp of an existing database cluster backup in ISO8601 combined date and time format. The most recent backup will be used if excluded.
	ArgDatabaseRestoreFromTimestamp = "restore-from-timestamp"
	// ArgDatabaseEngine is a flag for specifying which database engine to use
	ArgDatabaseEngine = "engine"
	// ArgDatabaseConfigJson is a flag for specifying the database configuration in JSON format for an update
	ArgDatabaseConfigJson = "config-json"
	// ArgDatabaseNumNodes is the number of nodes in the database cluster
	ArgDatabaseNumNodes = "num-nodes"
	// ArgDatabaseStorageSizeMib is the amount of disk space, in MiB, that should be allocated to the database cluster
	ArgDatabaseStorageSizeMib = "storage-size-mib"
	// ArgDatabaseMaintenanceDay is the new day for the maintenance window
	ArgDatabaseMaintenanceDay = "day"
	// ArgDatabaseMaintenanceHour is the new hour for the maintenance window
	ArgDatabaseMaintenanceHour = "hour"
	// ArgDatabasePoolUserName is the name of user for use with connection pool
	ArgDatabasePoolUserName = "user"
	// ArgDatabasePoolDBName is the database for use with connection pool
	ArgDatabasePoolDBName = "db"
	// ArgDatabasePoolSize is the flag for connection pool size
	ArgDatabasePoolSize = "size"
	// ArgDatabasePoolMode is the flag for connection pool mode
	ArgDatabasePoolMode = "mode"
	// ArgDatabaseUserMySQLAuthPlugin is a flag for setting the MySQL user auth plugin
	ArgDatabaseUserMySQLAuthPlugin = "mysql-auth-plugin"
	// ArgDatabasePrivateConnectionBool determine if the private connection details should be shown
	ArgDatabasePrivateConnectionBool = "private"
	// ArgDatabaseUserKafkaACLs will specify permissions on topics in kafka clsuter
	ArgDatabaseUserKafkaACLs = "acl"
	// ArgDatabaseUserOpenSearchACLs will specify permissions on indexes in opensearch clsuter
	ArgDatabaseUserOpenSearchACLs = "opensearch-acl"

	// ArgDatabaseTopicReplicationFactor is the replication factor of a kafka topic
	ArgDatabaseTopicReplicationFactor = "replication-factor"
	// ArgDatabaseTopicPartitionCount is the number of partitions that are associated with a kafka topic
	ArgDatabaseTopicPartitionCount = "partition-count"
	// ArgDatabaseTopicCleanupPolicy is the cleanup policy associated with a kafka topic
	ArgDatabaseTopicCleanupPolicy = "cleanup-policy"
	// ArgDatabaseTopicCompressionType is the compression algorithm used for a kafka topic
	ArgDatabaseTopicCompressionType = "compression-type"
	// ArgDatabaseTopicDeleteRetentionMS is the amount of time, in ms, to retain delete tombstone markers for a kafka topic
	ArgDatabaseTopicDeleteRetentionMS = "delete-retention-ms"
	// ArgDatabaseTopicFileDeleteDelayMS is the amount of time, in ms, to wait before deleting a file from the filesystem
	ArgDatabaseTopicFileDeleteDelayMS = "file-delete-delay-ms"
	// ArgDatabaseTopicFlushMessages is the size, in bytes, of all messages to accumulate on a partition before flushing them to disk
	ArgDatabaseTopicFlushMessages = "flush-messages"
	// ArgDatabaseTopicFlushMS is the amount of time, in ms, a message is kept in memory before it is flushed to disk
	ArgDatabaseTopicFlushMS = "flush-ms"
	// ArgDatabaseTopicIntervalIndexBytes is the number of bytes between entries being added into the offset index
	ArgDatabaseTopicIntervalIndexBytes = "interval-index-bytes"
	// ArgDatabaseTopicMaxCompactionLagMS is the maximum amount of time, in ms, that a message will remain uncompacted (if compaction is enabled)
	ArgDatabaseTopicMaxCompactionLagMS = "max-compaction-lag-ms"
	// ArgDatabaseTopicMaxMessageBytes is the maximum size, in bytes, of the largest record batch that can be sent to the server
	ArgDatabaseTopicMaxMessageBytes = "max-message-bytes"
	// ArgDatabaseTopicMessageDownConversionEnable determines whether brokers should convert messages for consumers expecting older message formats
	ArgDatabaseTopicMessageDownConversionEnable = "message-down-conversion-enable"
	// ArgDatabaseTopicMessageFormatVersion is the version used by the broker to append messages to the kafka topic logs
	ArgDatabaseTopicMessageFormatVersion = "message-format-version"
	// ArgDatabaseTopicMessageTimestampType is the timestamp used for messages
	ArgDatabaseTopicMessageTimestampType = "message-timestamp-type"
	// ArgDatabaseTopicMinCleanableDirtyRatio is ratio, between 0 and 1, specifying the frequenty of log compaction
	ArgDatabaseTopicMinCleanableDirtyRatio = "min-cleanable-dirty-ratio"
	// ArgDatabaseTopicMinCompactionLagMS is the minimum time, in ms, that a message will remain uncompacted
	ArgDatabaseTopicMinCompactionLagMS = "min-compaction-lag-ms"
	// ArgDatabaseTopicMinInsyncReplicas is the minimum number of replicas that must ACK a write for the write to be considered successful
	ArgDatabaseTopicMinInsyncReplicas = "min-insync-replicas"
	// ArgDatabaseTopicPreallocate determines whether a file should be preallocated on disk when creating a new log segment
	ArgDatabaseTopicPreallocate = "preallocate"
	// ArgDatabaseTopicRetentionBytes is the maximum size, in bytes, of a topic log before messages are deleted
	ArgDatabaseTopicRetentionBytes = "retention-bytes"
	// ArgDatabaseTopicRetentionMS is the maximum time, in ms, that a message is retained before deleting it
	ArgDatabaseTopicRetentionMS = "retention-ms"
	// ArgDatabaseTopicSegmentBytes is the maximum size, in bytes, of a single log file
	ArgDatabaseTopicSegmentBytes = "segment-bytes"
	// ArgDatabaseTopicSegmentJitterMS is the maximum random jitter, in ms, subtracted from the scheduled segment roll time to avoid thundering herds of segment rolling
	ArgDatabaseTopicSegmentJitterMS = "segment-jitter-ms"
	// ArgDatabaseTopicSegmentMS is the period of time, in ms, after which the log will be forced to roll if the segment file isn't full
	ArgDatabaseTopicSegmentMS = "segment-ms"

	// ArgPrivateNetworkUUID is the flag for VPC UUID
	ArgPrivateNetworkUUID = "private-network-uuid"

	// ArgForce forces confirmation on actions
	ArgForce = "force"

	// ArgObjectName is the Kubernetes object name
	ArgObjectName = "name"
	// ArgObjectNamespace is the Kubernetes object namespace
	ArgObjectNamespace = "namespace"

	// ArgVPCName is a name of the VPC.
	ArgVPCName = "name"
	// ArgVPCDescription is a VPC description.
	ArgVPCDescription = "description"
	// ArgVPCDefault is the VPC default argument, to update a specific VPC to the default VPC.
	ArgVPCDefault = "default"
	// ArgVPCIPRange is a VPC range of IP addresses in CIDR notation.
	ArgVPCIPRange = "ip-range"

	// ArgVPCPeeringName is a name of the VPC Peering.
	ArgVPCPeeringName = "name"
	// ArgVPCPeeringVPCIDs is the vpc ids of the peering
	ArgVPCPeeringVPCIDs = "vpc-ids"
	// ArgVPCPeeringVPCID is id of the VPC.
	ArgVPCPeeringVPCID = "vpc-id"

	// ArgPartnerAttachmentType is the type of the Partner Attachment e.g. "partner".
	ArgPartnerAttachmentType = "type"
	// ArgPartnerAttachmentName is a name of the Partner Attachment.
	ArgPartnerAttachmentName = "name"
	// ArgPartnerAttachmentBandwidthInMbps is the connection bandwidth in megabits per second.
	ArgPartnerAttachmentBandwidthInMbps = "connection-bandwidth-in-mbps"
	// ArgPartnerAttachmentRegion is the region slug.
	ArgPartnerAttachmentRegion = "region"
	// ArgPartnerAttachmentNaaSProvider is the name of the Network as a Service provider
	ArgPartnerAttachmentNaaSProvider = "naas-provider"
	// ArgPartnerAttachmentVPCIDs are the IDs of the VPCs which the Partner Attachment is connected
	ArgPartnerAttachmentVPCIDs = "vpc-ids"
	// ArgPartnerAttachmentBGPLocalASN is the BGP Autonomous System Number (ASN) of the local device
	ArgPartnerAttachmentBGPLocalASN = "bgp-local-asn"
	// ArgPartnerAttachmentBGPLocalRouterIP is the BGP IP address of the local device
	ArgPartnerAttachmentBGPLocalRouterIP = "bgp-local-router-ip"
	// ArgPartnerAttachmentBGPPeerASN is the BGP Autonomous System Number (ASN) of the peer device
	ArgPartnerAttachmentBGPPeerASN = "bgp-peer-asn"
	// ArgPartnerAttachmentBGPPeerRouterIP is the BGP IP address of the peer device
	ArgPartnerAttachmentBGPPeerRouterIP = "bgp-peer-router-ip"
	// ArgPartnerAttachmentBGPAuthKey is the BGP MD5 authentication key
	ArgPartnerAttachmentBGPAuthKey = "bgp-auth-key"
	// ArgPartnerAttachmentRedundancyZone is the redundancy zone of the Partner Attachment
	ArgPartnerAttachmentRedundancyZone = "redundancy-zone"
	// ArgPartnerAttachmentParentUUID is the ha parent uuid of the Partner Attachment
	ArgPartnerAttachmentParentUUID = "parent-uuid"

	// ArgReadWrite indicates a generated token should be read/write.
	ArgReadWrite = "read-write"
	// ArgRegistry indicates the name of the registry.
	ArgRegistry = "registry"
	// ArgRegistryExpirySeconds indicates the length of time the token will be valid in seconds.
	ArgRegistryExpirySeconds = "expiry-seconds"
	// ArgRegistryReadOnly indicates that a generated registry API token should be read-only.
	ArgRegistryReadOnly = "read-only"
	// ArgRegistryNeverExpire indicates that a generated registry API token should never expire.
	ArgRegistryNeverExpire = "never-expire"
	// ArgSubscriptionTier is a subscription tier slug.
	ArgSubscriptionTier = "subscription-tier"
	// ArgGCIncludeUntaggedManifests indicates that a garbage collection should delete
	// all untagged manifests.
	ArgGCIncludeUntaggedManifests = "include-untagged-manifests"
	// ArgGCExcludeUnreferencedBlobs indicates that a garbage collection should
	// not delete unreferenced blobs.
	ArgGCExcludeUnreferencedBlobs = "exclude-unreferenced-blobs"
	// ArgRegistryAuthorizationServerEndpoint is the endpoint of the OAuth authorization server
	// used to revoke credentials on logout.
	ArgRegistryAuthorizationServerEndpoint = "authorization-server-endpoint"

	// 1-Click Args

	// ArgOneClicks is the flag to pass in 1-click application slugs
	ArgOneClicks = "1-clicks"

	// ArgOneClickType is the type of 1-Click
	ArgOneClickType = "type"

	//ArgDangerous indicates whether to delete the cluster and all it's associated resources
	ArgDangerous = "dangerous"

	// ArgDatabaseFirewallRule the firewall rules.
	ArgDatabaseFirewallRule = "rule"

	// ArgDatabaseFirewallRuleUUID is the UUID for the firewall rules.
	ArgDatabaseFirewallRuleUUID = "uuid"

	// Monitoring Args

	// ArgAlertPolicyDescription is the flag to pass in the alert policy description.
	ArgAlertPolicyDescription = "description"

	// ArgAlertPolicyType is the alert policy type.
	ArgAlertPolicyType = "type"

	// ArgAlertPolicyValue is the alert policy value.
	ArgAlertPolicyValue = "value"

	// ArgAlertPolicyWindow is the alert policy window.
	ArgAlertPolicyWindow = "window"

	// ArgAlertPolicyTags is the alert policy tags.
	ArgAlertPolicyTags = "tags"

	// ArgAlertPolicyEntities is the alert policy entities.
	ArgAlertPolicyEntities = "entities"

	// ArgAlertPolicyEnabled is whether the alert policy is enabled.
	ArgAlertPolicyEnabled = "enabled"

	// ArgAlertPolicyCompare is the alert policy comparator.
	ArgAlertPolicyCompare = "compare"

	// ArgAlertPolicyEmails are the emails to send alerts to.
	ArgAlertPolicyEmails = "emails"

	// ArgAlertPolicySlackChannels are the Slack channels to send alerts to.
	ArgAlertPolicySlackChannels = "slack-channels"

	// ArgAlertPolicySlackURLs are the Slack URLs to send alerts to.
	ArgAlertPolicySlackURLs = "slack-urls"

	// ArgTokenValidationServer is the server used to validate an OAuth token
	ArgTokenValidationServer = "token-validation-server"

	// ArgGPUs specifies to list GPU Droplets
	ArgGPUs = "gpus"

	// Agent Args

	// ArgAgentId is the ID of the agent.
	ArgAgentId = "agent-id"

	// ArgAgentName is the name of the agent.
	ArgAgentName = "name"

	// ArgAgentInstruction is the instruction for the agent.
	ArgAgentInstruction = "instruction"

	// ArgAgentModelId is the model ID for the agent.
	ArgModelId = "model-id"

	// ArgAgentProjectId is the project ID for the agent.
	ArgAgentProjectId = "project-id"

	// ArgAgentRegion is the region where the agent is deployed.
	ArgAgentRegion = "region"

	// ArgAnthropicKeyId is the Anthropic key ID for the agent.
	ArgAnthropicKeyId = "Anthropic-key-id"

	// ArgAgentDescription is the description for the agent.
	ArgAgentDescription = "description"

	// ArgKnowledgeBaseId is the knowledge base ID(s) attached to the agent.
	ArgKnowledgeBaseId = "knowledge-base-id"

	// ArgOpenAIKeyId is the OpenAI API key ID for the agent.
	ArgOpenAIKeyId = "openai-key-id"

	// ArgTags are the tags applied to the agent.
	ArgTags = "tags"

	// ArgAgentForce forces agent deletion without confirmation.
	ArgAgentForce = "force"

	// ArgK specifies how many results should be considered from an attached knowledge base.
	ArgK = "k"

	// ArgMaxTokens is the maximum number of tokens to generate in the response.
	ArgMaxTokens = "max-tokens"

	// ArgRetrievalMethod is the method used to retrieve information from the knowledge base.
	ArgRetrievalMethod = "retrieval-method"

	// ArgTemperature is the temperature for the response generation.
	ArgTemperature = "temperature"

	// ArgTopProbability is the top probability for the response generation.
	ArgTopProbability = "top-p"

	// ArgKnowledgeBaseDatabaseID is the ID of the database
	ArgKnowledgeBaseDatabaseID = "database-id"

	// ArgKnowledgeBaseName is the name of the knowledge base
	ArgKnowledgeBaseName = "name"

	// ArgKnowledgeBaseRegion is the region of the knowledge base
	ArgKnowledgeBaseRegion = "region"

	// ArgKnowledgeBaseProjectID is the project ID of the knowledge base
	ArgKnowledgeBaseProjectID = "project-id"

	// ArgKnowledgeBaseEmbeddingModelUUID is the UUID of the embedding model
	ArgKnowledgeBaseEmbeddingModelUUID = "embedding-model-uuid"

	// ArgKnowledgeBaseTags is the tags of the knowledge base
	ArgKnowledgeBaseTags = "tags"

	// ArgKnowledgeBaseVPCUUID is the UUID of the VPC
	ArgKnowledgeBaseVPCUUID = "vpc-uuid"

	// ArgKnowledgeBaseDataSource is the data source of the knowledge base
	ArgKnowledgeBaseDataSource = "data-sources"

	// ArgKnowledgeBaseBaseURL is the base URL of the knowledge base
	ArgKnowledgeBaseBaseURL = "base-url"

	// ArgKnowledgeBaseCrawlingOption is the crawling option of the knowledge base
	ArgKnowledgeBaseCrawlingOption = "crawling-option"

	// ArgKnowledgeBaseEmbedMedia is the embed media option of the knowledge base
	ArgKnowledgeBaseEmbedMedia = "embed-media"

	// ArgKnowledgeBaseUUID is the UUID of the knowledge base
	ArgKnowledgeBaseUUID = "uuid"

	// ArgKnowledgeBaseDataSourceID is the ID of the data source
	ArgKnowledgeBaseDataSourceID = "data-source-id"

	// ArgKnowledgeBaseBucketName is name of data source from Spaces
	ArgKnowledgeBaseBucketName = "bucket-name"

	// ArgKnowledgeBaseItemPath is the item path of the data source
	ArgKnowledgeBaseItemPath = "item-path"

	// ArgPrefix is a byoip prefix argument.
	ArgPrefix = "prefix"

	// ArgSignature is a byoip prefix argument.
	ArgSignature = "signature"

	// ArgAgentParentID is the ID for the parent agent.
	ArgParentAgentId = "parent-agent-id"

	// ArgChildAgentID is the ID for the child agent.
	ArgChildAgentId = "child-agent-id"

	// ArgAgentRouteId is the UUID for the agent linkage.
	ArgAgentRouteId = "route-id"

	// ArgAgentRouteName is the name of the route.
	ArgAgentRouteName = "route-name"

	// ArgAgentRouteIfCase is the if-case condition for the route.
	ArgAgentRouteIfCase = "if-case"

	// ArgFunctionIDs is the name of the function.
	ArgFunctionID = "function-id"

	//ArgFunctionName is the name of the function.
	ArgFunctionName = "name"

	// ArgFunctionDescription is the description of the function.
	ArgFunctionRouteDescription = "description"

	// ArgFunctionRouteFaasName is the name of the function route in the DigitalOcean functions platform
	ArgFunctionRouteFaasName = "faas-name"

	// ArgFunctionRouteFaasNamespace is the namespace of the function route in the DigitalOcean functions platform
	ArgFunctionRouteFaasNamespace = "faas-namespace"

	// ArgFunctionRouteInputSchema is the input schema of the function route
	ArgFunctionRouteInputSchema = "input-schema"

	// ArgFunctionRouteOutputSchema is the output schema of the function route
	ArgFunctionRouteOutputSchema = "output-schema"

	// ArgAgentAPIKeyName is the name of API Key of the agent.
	ArgAgentAPIKeyName = "name"

	// ArgAgentUUID is the uuid of the apikey.
	ArgAPIkeyUUID = "api-key-uuid"

	// ArgAPIKeyForce forces API Key deletion without confirmation.
	ArgAPIKeyForce = "force"
)
