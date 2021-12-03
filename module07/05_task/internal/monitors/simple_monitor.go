package monitors

// import (
// 	"fmt"
// )

// type Monitor interface {
// 	Type() string
// 	Run() error
// }

// const EnvironmentMonitorType = "monitor_type_simple"

// type simpleMonitor struct {
// 	monitorType string
// }

// func NewSimpleMonitor() Monitor {
// 	return &simpleMonitor{
// 		monitorType: EnvironmentMonitorType,
// 	}
// }

// func (m *simpleMonitor) Type() string {
// 	return m.monitorType
// }

// func (m *simpleMonitor) Run() error {
// 	fmt.Printf("run monitor: %s\n", m.monitorType)
// 	return nil
// }
