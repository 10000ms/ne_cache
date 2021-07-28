package common

const (
	NodeStatusUnKnow     NodeStatus = iota
	NodeStatusServing    NodeStatus = iota
	NodeStatusNotServing NodeStatus = iota
)

const (
	RedisCommandCommand RedisCommand = "COMMAND"
	RedisCommandSet     RedisCommand = "SET"
	RedisCommandGet     RedisCommand = "GET"
)

const (
	StringKeyLiveNodeList = "live_node_list"
	StringKeyLastUpdateTime = "last_update_time"
)
