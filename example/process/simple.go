package process

import (
	. "github.com/Bunny3th/easy-workflow/workflow/engine"
	. "github.com/Bunny3th/easy-workflow/workflow/model"
	"log"
)

func CreateSimpleProcessJson() (string, error) {
	// 开始节点
	startNode := Node{NodeID: "Start", NodeName: "申请会议室",
		NodeType: 0, UserIDs: []string{"$starter"},
	}

	// 判断节点
	GWConfig_Conditional := HybridGateway{[]Condition{{Expression: "$hours>=3", NodeID: "Manager"},
		{Expression: "$hours<3", NodeID: "END"}}, []string{}, 0}
	Node1 := Node{NodeID: "GW-Hours", NodeName: "校验会议时间",
		NodeType: 2, GWConfig: GWConfig_Conditional,
		PrevNodeIDs: []string{"Start"},
	}

	// 经理审批
	Node2 := Node{NodeID: "Manager", NodeName: "经理审批",
		NodeType: 1, Roles: []string{"主管"},
		UserIDs:     []string{"2"},
		PrevNodeIDs: []string{"GW-Hours"},
	}

	//结束节点
	endNode := Node{NodeID: "END", NodeName: "END",
		NodeType: 3, PrevNodeIDs: []string{"GW-Hours", "Manager"}}

	//流程是节点的集合，所以要把上面所有的节点放在一个切片中
	var Nodelist []Node
	Nodelist = append(Nodelist, Node1)
	Nodelist = append(Nodelist, Node2)
	Nodelist = append(Nodelist, endNode)
	Nodelist = append(Nodelist, startNode)

	process := Process{ProcessName: "会议申请", Source: "办公系统", Nodes: Nodelist}

	//转化为json
	j, err := JSONMarshal(process, false)
	return string(j), err
}

func CreateSimpleExampleProcess() {
	//获得示例流程json
	j, err := CreateSimpleProcessJson()
	if err != nil {
		log.Fatal(err)
	}

	//保存流程
	id, err := ProcessSave(j, "system")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("流程保存成功，ID：", id)
}
