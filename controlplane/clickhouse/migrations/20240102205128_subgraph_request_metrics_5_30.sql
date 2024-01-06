-- migrate:up

-- This table is used to aggregate subgraph request metrics.

CREATE TABLE IF NOT EXISTS cosmo.subgraph_request_metrics_5_30 (
   Timestamp DateTime('UTC') CODEC (Delta(4), ZSTD(1)),
   FederatedGraphID LowCardinality(String) CODEC(ZSTD(1)),
   OrganizationID LowCardinality(String) CODEC(ZSTD(1)),
   SubgraphID String CODEC (ZSTD(1)),
   TotalRequests UInt64 CODEC(ZSTD(1)),
   TotalErrors UInt64 CODEC(ZSTD(1))
) ENGINE = SummingMergeTree
PARTITION BY toDate(Timestamp)
-- This allows us to fetch request metrics by subgraph ID in the most efficient way
ORDER BY (
    SubgraphID, FederatedGraphID, OrganizationID, toUnixTimestamp(Timestamp)
)
TTL toDateTime(Timestamp) + toIntervalDay(30) SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- migrate:down

DROP TABLE IF EXISTS cosmo.subgraph_request_metrics_5_30