package supervisor

import (
	"encoding/json"

	"github.com/kolo/xmlrpc"
)

type Supervisor struct {
	rpcURL string
	client *xmlrpc.Client
}

func New(url string) *Supervisor {

	client, err := xmlrpc.NewClient("http://localhost:9001/RPC2", nil)
	if err != nil {
		panic(err)
	}
	return &Supervisor{
		rpcURL: url,
		client: client,
	}
}

// ListMethods 列出所有的方法
func (s *Supervisor) ListMethods() (methods []string, err error) {
	err = s.client.Call("system.listMethods", nil, &methods)
	return
}

// GetAPIVersion 返回 supervisord 使用的 RPC API 的版本
func (s *Supervisor) GetAPIVersion() (version string, err error) {
	err = s.client.Call("supervisor.getAPIVersion", nil, &version)
	return
}

// GetAPIVersion 返回supervisord使用的supervisor包的版本
func (s *Supervisor) GetSupervisorVersion() (version string, err error) {
	err = s.client.Call("supervisor.getSupervisorVersion", nil, &version)
	return
}

type state struct {
	StateCode int    `json:"statecode"` // -1, 0, 1, 2  分别对应如下
	StateName string `json:"statename"` // SHUTDOWN, RESTARTING, RUNNING, FATAL
}

// GetAPIVersion 返回supervisord使用的supervisor包的版本
func (s *Supervisor) GetState() (st state, err error) {
	var ret interface{}
	err = s.client.Call("supervisor.getState", nil, &ret)
	b, _ := json.Marshal(ret)
	json.Unmarshal(b, &st)
	return
}

// GetPID 返回supervisord的PID
func (s *Supervisor) GetPID() (pid int, err error) {
	err = s.client.Call("supervisor.getPID", nil, &pid)
	return
}

// ReadLog 读取supervisord的日志
func (s *Supervisor) ReadLog(offset int, length int) (log string, err error) {
	err = s.client.Call("supervisor.readLog", []interface{}{offset, length}, &log)
	return
}

// ClearLog 清空supervisord的日志
func (s *Supervisor) ClearLog() (cleard bool, err error) {
	err = s.client.Call("supervisor.clearLog", nil, &cleard)
	return
}

// Shutdown 关闭 supervisord 进程
func (s *Supervisor) Shutdown() (restarted bool, err error) {
	err = s.client.Call("supervisor.shutdown", nil, &restarted)
	return
}

// Restart 重启supervisord管理下的子进程
func (s *Supervisor) Restart() (restarted bool, err error) {
	err = s.client.Call("supervisor.restart", nil, &restarted)
	return
}

type processInfo struct {
	Name          string `json:"name"`
	Group         string `json:"group"`
	Description   string `json:"description"`
	Start         int    `json:"start"`
	Stop          int    `json:"stop"`
	Now           int    `json:"now"`
	State         int    `json:"state"`
	StateName     string `json:"statename"`
	SpawnErr      string `json:"spawnerr"`
	ExitStatus    int    `json:"exitstatus"`
	StdOutLogFile string `json:"stdout_logfile"`
	StdErrLogFile string `json:"stderr_logfile"`
	PID           int    `json:"pid"`
}

// GetProcessInfo 获取一个指定名字的supervisor进程的信息
func (s *Supervisor) GetProcessInfo(name string) (info processInfo, err error) {
	var ret any
	err = s.client.Call("supervisor.getProcessInfo", []any{name}, &ret)
	b, _ := json.Marshal(ret)
	json.Unmarshal(b, &info)
	return
}

// GetAllProcessInfo 获取所有的的supervisor进程的信息
func (s *Supervisor) GetAllProcessInfo() (info []processInfo, err error) {
	var ret any
	err = s.client.Call("supervisor.getAllProcessInfo", nil, &ret)
	b, _ := json.Marshal(ret)
	json.Unmarshal(b, &info)
	return
}

// StartProcess 通过名称启动进程, 启动成功返回pid, 失败pid为0
func (s *Supervisor) StartProcess(name string) (pid int, err error) {
	var started bool
	err = s.client.Call("supervisor.startProcess", []any{name, true}, &started)
	if err != nil {
		return
	}
	info, err := s.GetProcessInfo(name)
	pid = info.PID
	return
}

// StopProcess 通过名称启动进程, 启动成功返回pid, 失败pid为0
func (s *Supervisor) StopProcess(name string) (stoped bool, err error) {
	err = s.client.Call("supervisor.stopProcess", []any{name, true}, &stoped)
	return
}

/**
 * processStatusInfo
 * status UNKNOWN_METHOD = 1, INCORRECT_PARAMETERS = 2, BAD_ARGUMENTS = 3, SIGNATURE_UNSUPPORTED = 4,
 *        SHUTDOWN_STATE = 6, BAD_NAME = 10, BAD_SIGNAL = 11, NO_FILE = 20, NOT_EXECUTABLE = 21, FAILED = 30,
 *        ABNORMAL_TERMINATION = 40, SPAWN_ERROR = 50, ALREADY_STARTED = 60, NOT_RUNNING = 70, SUCCESS = 80,
 *        ALREADY_ADDED = 90, STILL_RUNNING = 91, CANT_REREAD = 92
 */
type processStatusInfo struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	Description string `json:"description"`
	Status      int    `json:"status"` // 80代表成功
}

// StartAllProcesses 启动所有进程
func (s *Supervisor) StartAllProcesses() (info []processStatusInfo, err error) {
	var ret any
	err = s.client.Call("supervisor.startAllProcesses", []any{true}, &ret)
	b, _ := json.Marshal(ret)
	json.Unmarshal(b, &info)
	return
}

// StopAllProcesses 停止所有进程
func (s *Supervisor) StopAllProcesses() (info []processStatusInfo, err error) {
	var ret any
	err = s.client.Call("supervisor.stopAllProcesses", []any{true}, &ret)
	b, _ := json.Marshal(ret)
	json.Unmarshal(b, &info)
	return
}
