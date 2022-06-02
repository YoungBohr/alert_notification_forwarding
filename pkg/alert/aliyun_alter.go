package alert

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AliyunAlertMessage struct {
	Form    map[string][]string
	Product string
	Project string
	Unit    string
}

func (a *AliyunAlertMessage) Get(key string) (s string) {
	if val, ok := a.Form[key]; ok {
		if len(val) > 0 {
			s = val[0]
		}
	}
	return
}

func (a *AliyunAlertMessage) ToMarkdown() (markdown string, err error) {
	if a.Form == nil {
		err = errors.New("post form is null\n")
		return
	}

	ns := a.Get("namespace")
	switch {
	case strings.Contains(ns, "ecs"):
		a.Product = "ECS"
	case strings.Contains(ns, "kvstore"):
		a.Product = "Redis"
	case strings.Contains(ns, "polardb"):
		a.Product = "PolarDB"
	case strings.Contains(ns, "rds"):
		a.Product = "RDS"
	case strings.Contains(ns, "rocketmq"):
		a.Product = "RocketMQ"
	case strings.Contains(ns, "oss"):
		a.Product = "OSS"
	default:
		a.Product = "云服务"
	}

	an := a.Get("alertName")
	switch {
	case strings.Contains(an, "率"):
		a.Unit = "%"
	case strings.Contains(an, "流量"):
		a.Unit = "Mbytes"
	case strings.Contains(an, "数量"):
		a.Unit = "Counts/s"
	case strings.Contains(an, "次数"):
		a.Unit = "Counts/s"
	case strings.Contains(an, "响应时间"):
		a.Unit = "us"
	default:
		a.Unit = ""
	}

	var project string
	switch a.Project {
	case "cem":
		project = "CEM"
	case "prw":
		project = "拼任务"
	case "kbt":
		project = "刊播通"
	case "wjyb":
		project = "问卷100"
	case "sls":
		project = "手拉手"
	default:
		project = a.Project
	}

	var title, corlor string
	switch a.Get("alertState") {
	case "ALERT":
		title = a.Product + "发生告警"
		corlor = "#FF0000"
	case "OK":
		title = a.Product + "恢复正常"
		corlor = "#32CD32"
	}

	timestamp, _ := strconv.ParseInt(a.Get("timestamp"), 10, 64)
	datetime := time.Unix(timestamp/1000, 0)

	markdown = fmt.Sprintf(`
			# [云监控] <font color=%s size=18>%s</font>
			> 时间: <font color=\"comment\">%s</font>
			> 项目: <font color=\"comment\">%v</font>
			> 实例: <font color=\"comment\">%v</font>
			> 监控指标: <font color=\"comment\">%v</font>
			> 报警条件: <font color=\"comment\">%v</font>
			> 现在情况: <font color=\"warning\">%v%v</font>
			> 持续时间: <font color=\"comment\">%v</font>`,
		corlor, title, datetime.String(), project,
		a.Get("instanceName"), a.Get("alertName"),
		a.Get("expression"), a.Get("curValue"),
		a.Unit, a.Get("lastTime"))
	return
}
