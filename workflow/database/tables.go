package database

import "time"

//流程定义表
type ProcDef struct {
	ID        int       `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;comment:流程ID"`
	Name      string    `gorm:"column:name;type:VARCHAR(250) NOT NULL;comment:流程名字;uniqueIndex:uix_name_source"`
	Version   int       `gorm:"column:version;type:INT UNSIGNED NOT NULL DEFAULT 1;comment:版本号"`
	Resource  string    `gorm:"column:resource;type:TEXT NOT NULL;comment:流程定义模板"`
	UserID    int       `gorm:"column:user_id;type:VARCHAR(250) NOT NULL;comment:创建者ID"`
	Source    string    `gorm:"column:source;type:VARCHAR(250) NOT NULL;uniqueIndex:uix_name_source;comment:来源(引擎可能被多个系统、组件等使用，这里记下从哪个来源创建的流程);"`
	CreatTime time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
}

type CommonID struct {
	ID int `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
}

//流程定义历史表
type HistProcDef struct {
	CommonID
	ProcID    int       `gorm:"column:proc_id;type:INT UNSIGNED NOT NULL;comment:流程ID"`
	Name      string    `gorm:"column:name;type:VARCHAR(250) NOT NULL;comment:流程名字;uniqueIndex:uix_name_source"`
	Version   int       `gorm:"column:version;type:INT UNSIGNED NOT NULL DEFAULT 1;comment:版本号"`
	Resource  string    `gorm:"column:resource;type:TEXT NOT NULL;comment:流程定义模板"`
	UserID    int       `gorm:"column:user_id;type:VARCHAR(250) NOT NULL;comment:创建者ID"`
	Source    string    `gorm:"column:source;type:VARCHAR(250) NOT NULL;uniqueIndex:uix_name_source;comment:来源(引擎可能被多个系统、组件等使用，这里记下从哪个来源创建的流程);"`
	CreatTime time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
}

//流程实例表
type ProcInst struct {
	ID            int       `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;comment:流程实例ID"`
	ProcID        int       `gorm:"column:proc_id;type:INT NOT NULL;index:ix_proc_id;comment:流程ID"`
	ProcVersion   int       `gorm:"column:proc_version;type:INT UNSIGNED NOT NULL;comment:流程版本号"`
	BusinessID    string    `gorm:"column:business_id;type:VARCHAR(250) DEFAULT NULL;comment:业务ID"`
	CurrentNodeID string    `gorm:"column:current_node_id;type:VARCHAR(250) NOT NULL;comment:当前进行节点ID"`
	CreateTime    time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
	Status        int       `gorm:"column:status;type:TINYINT DEFAULT 0 ;comment:0:未完成 1:已完成 2:撤销"`
}

//流程实例历史表
type HistProcInst struct {
	CommonID
	ProcInstID    int       `gorm:"column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	ProcID        int       `gorm:"column:proc_id;type:INT NOT NULL;index:ix_proc_id;comment:流程ID"`
	ProcVersion   int       `gorm:"column:proc_version;type:INT UNSIGNED NOT NULL;comment:流程版本号"`
	BusinessID    string    `gorm:"column:business_id;type:VARCHAR(250) DEFAULT NULL;comment:业务ID"`
	CurrentNodeID string    `gorm:"column:current_node_id;type:VARCHAR(250) NOT NULL;comment:当前进行节点ID"`
	CreateTime    time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
	Status        int       `gorm:"column:status;type:TINYINT DEFAULT 0 ;comment:0:未完成 1:已完成 2:撤销"`
}

//任务表
type Task struct {
	ID           int       `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;comment:任务ID"`
	ProcID       int       `gorm:"index:ix_proc_id;column:proc_id;type:INT UNSIGNED NOT NULL;comment:流程ID,冗余字段，偷懒用"`
	ProcInstID   int       `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	NodeID       string    `gorm:"column:node_id;type:VARCHAR(250) NOT NULL;comment:节点ID"`
	PrevNodeID   string    `gorm:"column:prev_node_id;type:VARCHAR(250) DEFAULT NULL;comment:上个处理节点ID,注意这里和execution中的上一个节点不一样，这里是实际审批处理时上个已处理节点的ID"`
	IsCosigned   int       `gorm:"column:is_cosigned;type:TINYINT DEFAULT 0;comment:0:任意一人通过即可 1:会签"`
	BatchCode    string    `gorm:"index:ix_batch_code;column:batch_code;type:VARCHAR(50) DEFAULT NULL;comment:批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别"`
	UserID       string    `gorm:"column:user_id;type:VARCHAR(250) NOT NULL;comment:分配用户ID"`
	IsPassed     int       `gorm:"column:is_passed;type:TINYINT DEFAULT NULL;comment:任务是否通过 0:驳回 1:通过"`
	IsFinished   int       `gorm:"column:is_finished;type:TINYINT DEFAULT 0;comment:0:任务未处理 1:处理完成.任务未必都是用户处理的，比如会签时一人驳回，其他任务系统自动设为已处理"`
	CreateTime   time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:系统创建任务时间"`
	FinishedTime time.Time `gorm:"column:finished_time;type:DATETIME DEFAULT NULL;comment:处理任务时间"`
}

//任务历史表
type HistTask struct {
	CommonID
	TaskID       int       `gorm:"column:task_id;type:INT UNSIGNED NOT NULL;comment:任务ID"`
	ProcID       int       `gorm:"index:ix_proc_id;column:proc_id;type:INT UNSIGNED NOT NULL;comment:流程ID,冗余字段，偷懒用"`
	ProcInstID   int       `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	NodeID       string    `gorm:"column:node_id;type:VARCHAR(250) NOT NULL;comment:节点ID"`
	PrevNodeID   string    `gorm:"column:prev_node_id;type:VARCHAR(250) DEFAULT NULL;comment:上个处理节点ID,注意这里和execution中的上一个节点不一样，这里是实际审批处理时上个已处理节点的ID"`
	IsCosigned   int       `gorm:"column:is_cosigned;type:TINYINT DEFAULT 0;comment:0:任意一人通过即可 1:会签"`
	BatchCode    string    `gorm:"index:ix_batch_code;column:batch_code;type:VARCHAR(50) DEFAULT NULL;comment:批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别"`
	UserID       string    `gorm:"column:user_id;type:VARCHAR(250) NOT NULL;comment:分配用户ID"`
	IsPassed     int       `gorm:"column:is_passed;type:TINYINT DEFAULT NULL;comment:任务是否通过 0:驳回 1:通过"`
	IsFinished   int       `gorm:"column:is_finished;type:TINYINT DEFAULT 0;comment:0:任务未处理 1:处理完成.任务未必都是用户处理的，比如会签时一人驳回，其他任务系统自动设为已处理"`
	CreateTime   time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:系统创建任务时间"`
	FinishedTime time.Time `gorm:"column:finished_time;type:DATETIME DEFAULT NULL;comment:处理任务时间"`
}

//任务备注表
type TaskComment struct {
	ID         int    `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcInstID int    `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	TaskID     int    `gorm:"index:ix_task_id;column:task_id;type:INT UNSIGNED NOT NULL;comment:任务ID"`
	Comment    string `gorm:"column:comment;type:TEXT;comment:任务备注"`
}

//任务备注历史表
type HistTaskComment struct {
	ID         int    `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcInstID int    `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	TaskID     int    `gorm:"index:ix_task_id;column:task_id;type:INT UNSIGNED NOT NULL;comment:任务ID"`
	Comment    string `gorm:"column:comment;type:TEXT;comment:任务备注"`
}

//流程节点执行关系定义表
type ProcExecution struct {
	ID          int       `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcID      int       `gorm:"index:ix_proc_id;column:proc_id;type:INT NOT NULL;comment:流程ID"`
	ProcVersion int       `gorm:"column:proc_version;type:INT UNSIGNED NOT NULL;comment:流程版本号"`
	NodeID      string    `gorm:"column:node_id;type:VARCHAR(250) NOT NULL;comment:节点ID"`
	NodeName    string    `gorm:"column:node_name;type:VARCHAR(250) NOT NULL;comment:节点名称"`
	PrevNodeID  string    `gorm:"column:prev_node_id;type:VARCHAR(250) DEFAULT NULL;comment:上级节点ID"`
	NodeType    int       `gorm:"column:node_type;type:TINYINT NOT NULL;comment:流程类型 0:开始节点 1:任务节点 2:网关节点 3:结束节点"`
	IsCosigned  int       `gorm:"column:is_cosigned;type:TINYINT NOT NULL;comment:是否会签"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
}

//流程节点执行关系定义历史表
type HistProcExecution struct {
	ID          int       `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcID      int       `gorm:"index:ix_proc_id;column:proc_id;type:INT NOT NULL;comment:流程ID"`
	ProcVersion int       `gorm:"column:proc_version;type:INT UNSIGNED NOT NULL;comment:流程版本号"`
	NodeID      string    `gorm:"column:node_id;type:VARCHAR(250) NOT NULL;comment:节点ID"`
	NodeName    string    `gorm:"column:node_name;type:VARCHAR(250) NOT NULL;comment:节点名称"`
	PrevNodeID  string    `gorm:"column:prev_node_id;type:VARCHAR(250) DEFAULT NULL;comment:上级节点ID"`
	NodeType    int       `gorm:"column:node_type;type:TINYINT NOT NULL;comment:流程类型 0:开始节点 1:任务节点 2:网关节点 3:结束节点"`
	IsCosigned  int       `gorm:"column:is_cosigned;type:TINYINT NOT NULL;comment:是否会签"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME DEFAULT NOW();comment:创建时间"`
}

//流程实例变量表
type ProcInstVariable struct {
	ID         int    `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcInstID int    `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	Key        string `gorm:"column:key;type:VARCHAR(250) NOT NULL;comment:变量key"`
	Value      string `gorm:"column:value;type:VARCHAR(250) NOT NULL;comment:变量value"`
}

//流程实例变量历史表
type HistProcInstVariable struct {
	ID         int    `gorm:"primaryKey;column:id;type:INT UNSIGNED NOT NULL AUTO_INCREMENT;"`
	ProcInstID int    `gorm:"index:ix_proc_inst_id;column:proc_inst_id;type:INT UNSIGNED NOT NULL;comment:流程实例ID"`
	Key        string `gorm:"column:key;type:VARCHAR(250) NOT NULL;comment:变量key"`
	Value      string `gorm:"column:value;type:VARCHAR(250) NOT NULL;comment:变量value"`
}
