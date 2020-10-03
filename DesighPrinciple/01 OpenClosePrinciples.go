package DesighPrinciple

import "fmt"

// 开闭原则：对扩展开放，对修改关闭

// 模块功能：当接口请求个数超过某个阈值发出告警
type AlertRule struct{}

func (a *AlertRule) getMathedRule(api string) int64 {
	return 0
}

type Notification struct{}

func (n *Notification) notify(level string) {

}

type Alert struct {
	rule         *AlertRule
	notification *Notification
}

func NewAlert(rule *AlertRule, notification *Notification) *Alert {
	return &Alert{
		rule:         rule,
		notification: notification,
	}
}

func (a *Alert) check(api string, requestCnt, errCnt, durationOfSeconds int64) {
	tps := requestCnt / durationOfSeconds
	if tps > a.rule.getMathedRule(api) {
		a.notification.notify("URGENCY")
	}
	if errCnt > a.rule.getMathedRule(api) {
		a.notification.notify("SEVERE")
	}
}

/*
	抛出问题：假如新增加一个功能：当每秒钟超时请求个数，超过某个预先配置最大阈值
	那么需要在check方法中：1，引入timeout超时参数，2：并在内部添加超时处理逻辑
	这样实现不友好，如果基于开闭原则，那么代码扩展性则很好。*/

// 开闭原则实现
type Alertx struct {
	alertHandlers []AlertHandler
}
type ApiStatInfo struct {
	api       string
	rstcnt    int64
	errcnt    int64
	durOfSecs int64
}
type AlertHandler interface {
	check(info ApiStatInfo)
}

func NewAlertx() *Alertx {
	return &Alertx{
		alertHandlers: make([]AlertHandler, 0),
	}
}

func (a *Alertx) addAlerHandler(handler AlertHandler) {
	a.alertHandlers = append(a.alertHandlers, handler)
}

func (a *Alertx) check(info ApiStatInfo) {
	for _, handle := range a.alertHandlers {
		handle.check(info)
	}
}

type TpsAlertHandler struct {
	rule         *AlertRule
	notification *Notification
}

func NewTpsAlertHandler(rule *AlertRule, notification *Notification) *TpsAlertHandler {
	return &TpsAlertHandler{
		rule:         rule,
		notification: notification,
	}
}
func (t *TpsAlertHandler) check(info ApiStatInfo) {
	tps := info.rstcnt / info.durOfSecs
	if tps > t.rule.getMathedRule(info.api) {
		t.notification.notify("URGENCY")
	}
	fmt.Println("TPS")
}

type ErrAlertHandler struct {
	rule         *AlertRule
	notification *Notification
}

func NewErrAlertHandler(rule *AlertRule, notification *Notification) *ErrAlertHandler {
	return &ErrAlertHandler{
		rule:         rule,
		notification: notification,
	}
}
func (e *ErrAlertHandler) check(info ApiStatInfo) {
	if info.errcnt > e.rule.getMathedRule(info.api) {
		e.notification.notify("SEVER")
	}
	fmt.Println("Err")
}

// 那么新增timeOut规则，只需要实现AlertHandler接口，并注册到Alertx
type TimeOutAlertHandler struct {
	rule         *AlertRule
	notification *Notification
}

func NewTimeAlertHandler(rule *AlertRule, notification *Notification) *TimeOutAlertHandler {
	return &TimeOutAlertHandler{
		rule:         rule,
		notification: notification,
	}
}
func (t *TimeOutAlertHandler) check(info ApiStatInfo) {
	fmt.Println("TimeOut!")
}
