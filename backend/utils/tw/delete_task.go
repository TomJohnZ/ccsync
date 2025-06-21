package tw

import (
	"ccsync_backend/utils"
	"fmt"
	"os"
)

func DeleteTaskInTaskwarrior(email, encryptionSecret, uuid, taskuuid string) error {
	tempDir, err := os.MkdirTemp("", "taskwarrior-"+email)
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	origin := os.Getenv("CONTAINER_ORIGIN")
	if err := SetTaskwarriorConfig(tempDir, encryptionSecret, origin, uuid); err != nil {
		return err
	}

	if err := SyncTaskwarrior(tempDir); err != nil {
		return err
	}

	if err := utils.ExecCommandInDir(tempDir, "task", taskuuid, "delete", "rc.confirmation=off"); err != nil {
		return fmt.Errorf("failed to mark task as deleted: %v", err)
	}

	// Sync Taskwarrior again
	if err := SyncTaskwarrior(tempDir); err != nil {
		return err
	}

	return nil
}
