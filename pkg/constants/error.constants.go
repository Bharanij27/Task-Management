package constants

const (
	Invalid_Payload     = "Invalid Request payload"
	Internal_Error      = "Server error: Can't process the request please vist after some time"
	Creation_Success    = "Successfully created a task with id %s"
	Update_Success      = "Successfully updated a task %s"
	Invalid_Task_Fetch  = "Requested taskId %s does not exists"
	Invalid_Task_Update = "Update failed: Requested taskId %s does not exists"
	Deletion_Failed     = "Deletion failed: Failed to delete task %s"
	Deletion_Success    = "Successfully deleted a task %s"
	Cron_Failed         = "Cron job is failed..."
)
