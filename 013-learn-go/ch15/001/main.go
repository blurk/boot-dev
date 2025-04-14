package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	errUpdateStatus := em.recipient.updateStatus(em.status)

	if errUpdateStatus != nil {
		return fmt.Errorf("error updating user status: %w", errUpdateStatus)
	}

	errTrack := a.track(em.status)

	if errTrack != nil {
		return fmt.Errorf("error tracking user bounce: %w", errTrack)
	}

	return nil
}
