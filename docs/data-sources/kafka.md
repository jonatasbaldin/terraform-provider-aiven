# Kafka Data Source

The Kafka data source provides information about the existing Aiven Kafka services.

## Example Usage

```hcl
data "aiven_kafka" "kafka1" {
    project = data.aiven_project.pr1.project
    service_name = "my-kafka1"
}
```

## Argument Reference

* `project` - (Required) identifies the project the service belongs to. To set up proper dependency
between the project and the service, refer to the project as shown in the above example.
Project cannot be changed later without destroying and re-creating the service.

* `service_name` - (Required) specifies the actual name of the service. The name cannot be changed
later without destroying and re-creating the service so name should be picked based on
intended service usage rather than current attributes.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_name` - defines where the cloud provider and region where the service is hosted
in. This can be changed freely after service is created. Changing the value will trigger
a potentially lengthy migration process for the service. Format is cloud provider name
(`aws`, `azure`, `do` `google`, `upcloud`, etc.), dash, and the cloud provider
specific region name. These are documented on each Cloud provider's own support articles,
like [here for Google](https://cloud.google.com/compute/docs/regions-zones/) and
[here for AWS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).

* `plan` - defines what kind of computing resources are allocated for the service. It can
be changed after creation, though there are some restrictions when going to a smaller
plan such as the new plan must have sufficient amount of disk space to store all current
data and switching to a plan with fewer nodes might not be supported. The basic plan
names are `hobbyist`, `startup-x`, `business-x` and `premium-x` where `x` is
(roughly) the amount of memory on each node (also other attributes like number of CPUs
and amount of disk space varies but naming is based on memory). The exact options can be
seen from the Aiven web console's Create Service dialog.

* `project_vpc_id` - optionally specifies the VPC the service should run in. If the value
is not set the service is not run inside a VPC. When set, the value should be given as a
reference as shown above to set up dependencies correctly and the VPC must be in the same
cloud and region as the service itself. Project can be freely moved to and from VPC after
creation but doing so triggers migration to new servers so the operation can take
significant amount of time to complete if the service has a lot of data.

* `termination_protection` - prevents the service from being deleted. It is recommended to
set this to `true` for all production services to prevent unintentional service
deletion. This does not shield against deleting databases or topics but for services
with backups much of the content can at least be restored from backup in case accidental
deletion is done.

* `maintenance_window_dow` - day of week when maintenance operations should be performed. 
On monday, tuesday, wednesday, etc.

* `maintenance_window_time` - time of day when maintenance operations should be performed. 
UTC time in HH:mm:ss format.

* `kafka_user_config` - defines Kafka specific additional configuration options. The following 
configuration options available:
    * `custom_domain` - Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.
    * `ip_filter` - Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
    
    * `kafka` - Kafka broker configuration values.
        * `auto_create_topics_enable` - Enable auto creation of topics
        * `compression_type` - Specify the final compression type for a given topic. This 
        configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). 
        It additionally accepts 'uncompressed' which is equivalent to no compression; and 'producer' 
        which means retain the original compression codec set by the producer.
        * `connections_max_idle_ms` - Idle connections timeout: the server socket processor 
        threads close the connections that idle for longer than this.
        * `default_replication_factor` - Replication factor for autocreated topics
        * `group_max_session_timeout_ms` - The maximum allowed session timeout for registered 
        consumers. Longer timeouts give consumers more time to process messages in between heartbeats 
        at the cost of a longer time to detect failures.
        * `group_min_session_timeout_ms` - The minimum allowed session timeout for registered 
        consumers. Longer timeouts give consumers more time to process messages in between heartbeats 
        at the cost of a longer time to detect failures.
        * `log_cleaner_max_compaction_lag_ms` - The maximum amount of time message will 
        remain uncompacted. Only applicable for logs that are being compacted
        * `log_cleaner_min_cleanable_ratio` - Controls log compactor frequency. Larger 
        value means more frequent compactions but also more space wasted for logs. Consider setting 
        log.cleaner.max.compaction.lag.ms to enforce compactions sooner, instead of setting a very 
        high value for this option.
        * `log_cleaner_min_compaction_lag_ms` - The minimum time a message will remain 
        uncompacted in the log. Only applicable for logs that are being compacted.
        * `log_cleanup_policy` - The default cleanup policy for segments beyond the retention window.
        * `log_flush_interval_messages` - The number of messages accumulated on a log partition 
        before messages are flushed to disk.
        * `log_flush_interval_ms` - The maximum time in ms that a message in any topic is kept 
        in memory before flushed to disk. If not set, the value in log.flush.scheduler.interval.ms is used.
        * `log_index_interval_bytes` - The interval with which Kafka adds an entry to the offset index.
        * `log_index_size_max_bytes` - The maximum size in bytes of the offset index.
        * `log_message_downconversion_enable` - This configuration controls whether down-conversion 
        of message formats is enabled to satisfy consume requests.
        * `log_message_timestamp_difference_max_ms` - The maximum difference allowed between 
        the timestamp when a broker receives a message and the timestamp specified in the message
        * `log_message_timestamp_type` - Define whether the timestamp in the message is 
        message create time or log append time.
        * `log_preallocate` - Should pre allocate file when create new segment?
        * `log_retention_bytes` - The maximum size of the log before deleting messages
        * `log_retention_hours` - The number of hours to keep a log file before deleting it.
        * `log_retention_ms` - The number of milliseconds to keep a log file before deleting it 
        (in milliseconds), If not set, the value in log.retention.minutes is used. If set to -1, no 
        time limit is applied.
        * `log_roll_jitter_ms` - The maximum jitter to subtract from logRollTimeMillis 
        (in milliseconds). If not set, the value in log.roll.jitter.hours is used.
        * `log_roll_ms` - The maximum time before a new log segment is rolled out (in milliseconds).
        * `log_segment_bytes` - The maximum size of a single log file
        * `max_connections_per_ip` - The maximum number of connections allowed from each ip 
        address (defaults to 2147483647).
        * `max_incremental_fetch_session_cache_slots` - The maximum number of incremental fetch 
        sessions that the broker will maintain.
        * `log_segment_delete_delay_ms` - The amount of time to wait before deleting a file 
        from the filesystem.
        * `message_max_bytes` - The maximum size of message that the server can receive.
        * `num_partitions` - Number of partitions for autocreated topics
        * `offsets_retention_minutes` - Log retention window in minutes for offsets topic.
        * `min_insync_replicas` - When a producer sets acks to 'all' (or '-1'), 
        min.insync.replicas specifies the minimum number of replicas that must acknowledge a write for 
        the write to be considered successful.
        * `producer_purgatory_purge_interval_requests` - The purge interval (in number of 
        requests) of the producer request purgatory(defaults to 1000).
        * `replica_fetch_max_bytes` - The number of bytes of messages to attempt to fetch 
        for each partition (defaults to 1048576). This is not an absolute maximum, if the first record 
        batch in the first non-empty partition of the fetch is larger than this value, the record batch 
        will still be returned to ensure that progress can be made.
        * `replica_fetch_response_max_bytes` - Maximum bytes expected for the entire fetch 
        response (defaults to 10485760). Records are fetched in batches, and if the first record batch 
        in the first non-empty partition of the fetch is larger than this value, the record batch will 
        still be returned to ensure that progress can be made. As such, this is not an absolute maximum.
        * `socket_request_max_bytes` - The maximum number of bytes in a socket request 
        (defaults to 104857600).
         
    * `kafka_authentication_methods` - Kafka authentication methods
        * `certificate` - Enable certificate/SSL authentication
        * `sasl` - Enable SASL authentication
    
    * `kafka_connect` - Enable Kafka Connect service
    * `kafka_connect_config` - Kafka Connect configuration values
        * `connector_client_config_override_policy` - Defines what client configurations can 
        be overridden by the connector. Default is None
        * `consumer_auto_offset_reset` - What to do when there is no initial offset in Kafka or 
        if the current offset does not exist any more on the server. Default is earliest.
        * `consumer_fetch_max_bytes` - Records are fetched in batches by the consumer, and 
        if the first record batch in the first non-empty partition of the fetch is larger than this value, 
        the record batch will still be returned to ensure that the consumer can make progress. As such, 
        this is not a absolute maximum.
        * `consumer_isolation_level` - Transaction read isolation level. read_uncommitted is 
        the default, but read_committed can be used if consume-exactly-once behavior is desired.
        * `consumer_max_partition_fetch_bytes` - Records are fetched in batches by the consumer.If 
        the first record batch in the first non-empty partition of the fetch is larger than this limit, 
        the batch will still be returned to ensure that the consumer can make progress.
        * `consumer_max_poll_interval_ms` - The maximum delay in milliseconds between invocations 
        of poll() when using consumer group management (defaults to 300000).
        * `consumer_max_poll_records` - The maximum number of records returned in a single call 
        to poll() (defaults to 500).
        * `offset_flush_interval_ms` - The interval at which to try committing offsets for 
        tasks (defaults to 60000).
        * `offset_flush_timeout_ms` - Maximum number of milliseconds to wait for records to 
        flush and partition offset data to be committed to offset storage before cancelling the process 
        and restoring the offset data to be committed in a future attempt (defaults to 5000).
        * `producer_max_request_size` - This setting will limit the number of record batches 
        the producer will send in a single request to avoid sending huge requests.
        * `session_timeout_ms` - The timeout in milliseconds used to detect failures when 
        using Kafka’s group management facilities (defaults to 10000). 
        * `transaction_remove_expired_transaction_cleanup_interval_ms` - The interval at which 
        to remove transactions that have expired due to transactional.id.expiration.ms passing (defaults 
        to 3600000 (1 hour)).
        * `transaction_state_log_segment_bytes` - The transaction topic segment bytes should 
        be kept relatively small in order to facilitate faster log compaction and cache loads (defaults 
        to 104857600 (100 mebibytes)).
    
    * `kafka_rest` - Enable Kafka-REST service
    * `kafka_rest_config` - Kafka-REST configuration
        * `consumer_enable_auto_commit` - If true the consumer's offset will be periodically 
        committed to Kafka in the background
        * `consumer_request_max_bytes` - Maximum number of bytes in unencoded message keys and 
        values by a single request
        * `consumer_request_timeout_ms` - The maximum total time to wait for messages for a 
        request if the maximum number of messages has not yet been reached
        * `producer_acks` - The number of acknowledgments the producer requires the leader to 
        have received before considering a request complete. If set to 'all' or '-1', the leader will wait 
        for the full set of in-sync replicas to acknowledge the record.
        * `producer_linger_ms` - Wait for up to the given delay to allow batching records together
        * `simpleconsumer_pool_size_max` - Maximum number of SimpleConsumers that can be 
        instantiated per broker.
        
    * `schema_registry_config` - Schema Registry configuration
        * `leader_eligibility` - If true, Karapace / Schema Registry on the service nodes can 
        participate in leader election. It might be needed to disable this when the schemas topic is replicated 
        to a secondary cluster and Karapace / Schema Registry there must not participate in leader election. 
        Defaults to 'true'.
        * `topic_name` - The durable single partition topic that acts as the durable log for the 
        data. This topic must be compacted to avoid losing data due to retention policy. Please note that 
        changing this configuration in an existing Schema Registry / Karapace setup leads to previous 
        schemas being inaccessible, data encoded with them potentially unreadable and schema ID sequence 
        put out of order. It's only possible to do the switch while Schema Registry / Karapace is disabled. 
        Defaults to '_schemas'.
    
    * `kafka_version` - Kafka major version
    
    * `private_access` - Allow access to selected service ports from private networks
        * `prometheus` - Allow clients to connect to prometheus with a DNS name that always resolves 
        to the service's private IP addresses. Only available in certain network locations
        
    * `public_access` - Allow access to selected service ports from the public Internet
        * `kafka` - Allow clients to connect to kafka from the public internet for service 
        nodes that are in a project VPC or another type of private network
        * `kafka_connect` - Allow clients to connect to kafka_connect from the public internet 
        for service nodes that are in a project VPC or another type of private network
        * `kafka_rest` - Allow clients to connect to kafka_rest from the public internet for 
        service nodes that are in a project VPC or another type of private network
        * `prometheus` - Allow clients to connect to prometheus from the public internet for 
        service nodes that are in a project VPC or another type of private network
        * `schema_registry` - Allow clients to connect to schema_registry from the public 
        internet for service nodes that are in a project VPC or another type of private network
        
    * `schema_registry` - Enable Schema-Registry service
    
    * `privatelink_access` - Allow access to selected service components through Privatelink
        * `kafka` - Enable kafka
        * `kafka_connect` - Enable kafka_connect
        * `kafka_rest` - Enable kafka_rest
        * `schema_registry` - Enable schema_registry
    
* `service_uri` - URI for connecting to the Kafka service.

* `service_host` - Kafka hostname.

* `service_port` - Kafka port.

* `service_password` - Password used for connecting to the Kafka service, if applicable.

* `service_username` - Username used for connecting to the Kafka service, if applicable.

* `state` - Service state.

* `kafka` - Kafka server provided values:
    * `access_cert` - The Kafka client certificate
    * `access_key` - The Kafka client certificate key
    * `connect_uri` - The Kafka Connect URI, if any
    * `rest_uri` - The Kafka REST URI, if any
    * `schema_registry_uri` - The Schema Registry URI, if any
