package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Info(args []string) []byte {
	return utils.ToRESP(` 
 # Server 
 redis_version:7.4.2 
 redis_git_sha1:00000000
 redis_git_dirty:0
 redis_build_id:d7b80adc76717456 
 redis_mode:standalone 
 os:Linux 6.10.14-linuxkit aarch64 
 arch_bits:64 
 monotonic_clock:POSIX clock_gettime 
 multiplexing_api:epoll 
 atomicvar_api:c11-builtin 
 gcc_version:12.2.0 
 process_id:1 
 process_supervised:no 
 run_id:2cd67f7eee234fce49c33788fe16f8548eb949d6 
 tcp_port:6379 
 server_time_usec:1745411481497518 
 uptime_in_seconds:69 
 uptime_in_days:0 
 hz:10 
 configured_hz:10 
 lru_clock:581017 
 executable:/data/redis-server 
 config_file: 
 io_threads_active:0 
 listener0:name=tcp,bind=*,bind=-::*,port=6379 
 
 # Clients 
 connected_clients:3 
 cluster_connections:0 
 maxclients:10000 
 client_recent_max_input_buffer:20480 
 client_recent_max_output_buffer:0 
 blocked_clients:0 
 tracking_clients:0 
 pubsub_clients:0 
 watching_clients:0 
 clients_in_timeout_table:0 
 total_watched_keys:0 
 total_blocking_keys:0 
 total_blocking_keys_on_nokey:0 
 
 # Memory 
 used_memory:1207064 
 used_memory_human:1.15M 
 used_memory_rss:15822848 
 used_memory_rss_human:15.09M 
 used_memory_peak:1229496 
 used_memory_peak_human:1.17M 
 used_memory_peak_perc:98.18% 
 used_memory_overhead:1005536 
 used_memory_startup:979088 
 used_memory_dataset:201528 
 used_memory_dataset_perc:88.40% 
 allocator_allocated:6265120 
 allocator_active:15335424 
 allocator_resident:17760256 
 allocator_muzzy:0 
 total_system_memory:8218251264 
 total_system_memory_human:7.65G 
 used_memory_lua:31744 
 used_memory_vm_eval:31744 
 used_memory_lua_human:31.00K 
 used_memory_scripts_eval:0 
 number_of_cached_scripts:0 
 number_of_functions:0 
 number_of_libraries:0 
 used_memory_vm_functions:32768 
 used_memory_vm_total:64512 
 used_memory_vm_total_human:63.00K 
 used_memory_functions:192 
 used_memory_scripts:192 
 used_memory_scripts_human:192B 
 maxmemory:0 
 maxmemory_human:0B 
 maxmemory_policy:noeviction 
 allocator_frag_ratio:2.41 
 allocator_frag_bytes:5980640 
 allocator_rss_ratio:1.16 
 allocator_rss_bytes:2424832 
 rss_overhead_ratio:0.89 
 rss_overhead_bytes:-1937408 
 mem_fragmentation_ratio:13.13 
 mem_fragmentation_bytes:14617608 
 mem_not_counted_for_evict:0 
 mem_replication_backlog:0 
 mem_total_replication_buffers:0 
 mem_clients_slaves:0 
 mem_clients_normal:26256 
 mem_cluster_links:0 
 mem_aof_buffer:0 
 mem_allocator:jemalloc-5.3.0 
 mem_overhead_db_hashtable_rehashing:0 
 active_defrag_running:0 
 lazyfree_pending_objects:0 
 lazyfreed_objects:0 
 
 # Persistence 
 loading:0 
 async_loading:0 
 current_cow_peak:0 
 current_cow_size:0 
 current_cow_size_age:0 
 current_fork_perc:0.00 
 current_save_keys_processed:0 
 current_save_keys_total:0 
 rdb_changes_since_last_save:0 
 rdb_bgsave_in_progress:0 
 rdb_last_save_time:1745411412 
 rdb_last_bgsave_status:ok 
 rdb_last_bgsave_time_sec:-1 
 rdb_current_bgsave_time_sec:-1 
 rdb_saves:0 
 rdb_last_cow_size:0 
 rdb_last_load_keys_expired:0 
 rdb_last_load_keys_loaded:0 
 aof_enabled:0 
 aof_rewrite_in_progress:0 
 aof_rewrite_scheduled:0 
 aof_last_rewrite_time_sec:-1 
 aof_current_rewrite_time_sec:-1 
 aof_last_bgrewrite_status:ok 
 aof_rewrites:0 
 aof_rewrites_consecutive_failures:0 
 aof_last_write_status:ok 
 aof_last_cow_size:0 
 module_fork_in_progress:0 
 module_fork_last_cow_size:0 
 
 # Stats 
 total_connections_received:3 
 total_commands_processed:23 
 instantaneous_ops_per_sec:0 
 total_net_input_bytes:713 
 total_net_output_bytes:57022 
 total_net_repl_input_bytes:0 
 total_net_repl_output_bytes:0 
 instantaneous_input_kbps:0.00 
 instantaneous_output_kbps:0.00 
 instantaneous_input_repl_kbps:0.00 
 instantaneous_output_repl_kbps:0.00 
 rejected_connections:0 
 sync_full:0 
 sync_partial_ok:0 
 sync_partial_err:0 
 expired_subkeys:0 
 expired_keys:0 
 expired_stale_perc:0.00 
 expired_time_cap_reached_count:0 
 expire_cycle_cpu_milliseconds:2 
 evicted_keys:0 
 evicted_clients:0 
 evicted_scripts:0 
 total_eviction_exceeded_time:0 
 current_eviction_exceeded_time:0 
 keyspace_hits:0 
 keyspace_misses:0 
 pubsub_channels:0 
 pubsub_patterns:0 
 pubsubshard_channels:0 
 latest_fork_usec:0 
 total_forks:0 
 migrate_cached_sockets:0 
 slave_expires_tracked_keys:0 
 active_defrag_hits:0 
 active_defrag_misses:0 
 active_defrag_key_hits:0 
 active_defrag_key_misses:0 
 total_active_defrag_time:0 
 current_active_defrag_time:0 
 tracking_total_keys:0 
 tracking_total_items:0 
 tracking_total_prefixes:0 
 unexpected_error_replies:0 
 total_error_replies:4 
 dump_payload_sanitizations:0 
 total_reads_processed:30 
 total_writes_processed:26 
 io_threaded_reads_processed:0 
 io_threaded_writes_processed:0 
 client_query_buffer_limit_disconnections:0 
 client_output_buffer_limit_disconnections:0 
 reply_buffer_shrinks:5 
 reply_buffer_expands:2 
 eventloop_cycles:698 
 eventloop_duration_sum:466522 
 eventloop_duration_cmd_sum:1241 
 instantaneous_eventloop_cycles_per_sec:9 
 instantaneous_eventloop_duration_usec:682 
 acl_access_denied_auth:0 
 acl_access_denied_cmd:0 
 acl_access_denied_key:0 
 acl_access_denied_channel:0 
 
 # Replication 
 role:master 
 connected_slaves:0 
 master_failover_state:no-failover 
 master_replid:2e46b7c375773930bb85913d597f25dcdabf584f 
 master_replid2:0000000000000000000000000000000000000000 
 master_repl_offset:0 
 second_repl_offset:-1 
 repl_backlog_active:0 
 repl_backlog_size:1048576 
 repl_backlog_first_byte_offset:0 
 repl_backlog_histlen:0 
 
 # CPU 
 used_cpu_sys:0.124799 
 used_cpu_user:0.412718 
 used_cpu_sys_children:0.001039 
 used_cpu_user_children:0.002030 
 used_cpu_sys_main_thread:0.123637 
 used_cpu_user_main_thread:0.412925 
 
 # Modules 
 
 # Errorstats 
 errorstat_ERR:count=4 
 
 # Cluster 
 cluster_enabled:0 
 
 # Keyspace 
 
 write this as INFO`)
}
