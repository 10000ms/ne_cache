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
