package DesighPrinciple

import "testing"

func TestAlertx(t *testing.T) {
	alter := NewAlertx()
	rule := new(AlertRule)
	apiInfo := ApiStatInfo{
		rstcnt:    1000,
		errcnt:    10,
		durOfSecs: 50,
	}
	notification := new(Notification)
	alter.addAlerHandler(NewTpsAlertHandler(rule, notification))
	alter.addAlerHandler(NewErrAlertHandler(rule, notification))
	alter.addAlerHandler(NewTimeAlertHandler(rule, notification))
	alter.check(apiInfo)
}
