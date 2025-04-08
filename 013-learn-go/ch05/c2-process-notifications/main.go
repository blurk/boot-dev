package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch notification := n.(type) {
	case directMessage:
		return notification.senderUsername, notification.importance()
	case groupMessage:
		return notification.groupName, notification.importance()
	case systemAlert:
		return notification.alertCode, notification.importance()
	default:
		return "", 0
	}
}
