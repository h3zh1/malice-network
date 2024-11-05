package consts

const (
	CalleeCMD      = "cmd"
	CalleeMal      = "mal"
	CalleeSDK      = "sdk"
	CalleeExplorer = "explorer"
)

const (
	CtrlTaskCallback      = "task_callback"
	CtrlTaskFinish        = "task_finish"
	CtrlTaskCancel        = "task_cancel"
	CtrlTaskError         = "task_error"
	CtrlClientJoin        = "client_join"
	CtrlClientLeft        = "client_left"
	CtrlWebUpload         = "web_upload"
	CtrlListenerStart     = "listener_start"
	CtrlListenerStop      = "listener_stop"
	CtrlPipelineStart     = "pipeline_start"
	CtrlPipelineStop      = "pipeline_stop"
	CtrlWebsiteStart      = "website_start"
	CtrlWebsiteStop       = "website_stop"
	CtrlWebsiteRegister   = "website_register"
	CtrlJobStart          = "job_start"
	CtrlJobStop           = "job_stop"
	CtrlSessionRegister   = "session_register"
	CtrlSessionReRegister = "session_re-register"
	CtrlSessionDead       = "session_dead"
	CtrlSessionInit       = "session_init"
	CtrlSessionReborn     = "session_reborn"
	CtrlSessionLog        = "session_log"
	CtrlSessionTask       = "session_task"
	CtrlSessionError      = "session_error"
	CtrlSessionStop       = "session_stop"
	CtrlSessionCheckin    = "session_checkin"
)

// ctrl status
const (
	CtrlStatusSuccess = 0 + iota
	CtrlStatusFailed
)

// event
const (
	EventJoin        = "join"
	EventLeft        = "left"
	EventClient      = "client"
	EventBroadcast   = "broadcast"
	EventNotify      = "notify"
	EventSession     = "session"
	EventListener    = "listener"
	EventTask        = "task"
	EventWebsite     = "website"
	EventTcpPipeline = "tcp"
	EventJob         = "job"
)
