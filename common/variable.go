package common

var NodeList = MangerNodeManage{
	List: make(map[string]*MangeSingleNode),
}

var NodeManager = ServerNodeManger{
	RawNodeList:  make(map[string]*ServerSingleNode),
	HashMap:      make(map[int]*ServerSingleNode),
	NodeMultiple: 4, // 4倍节点数
	NodeHash:     make([]int, 0),
}
