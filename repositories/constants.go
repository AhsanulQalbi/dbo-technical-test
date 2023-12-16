package repositories

const (
	// WHERE
	// WHERE_USER_ID       = "user_id = ?"
	// WHERE_ASSIGNEE_ID   = "assignee_id = ?"
	// WHERE_MANPOWER_ID   = "manpower_id = ?"
	// WHERE_CLEAR_PIC_ID  = "clear_pic_id = ?"
	// WHERE_ID            = "id = ?"
	// WHERE_TRIP_ID       = "trip_id = ?"
	// WHERE_RESIGN_ID     = "resign_id = ?"
	// WHERE_TRAINING_ID   = "training_id = ?"
	// WHERE_DIRECTLINE_ID = "direct_line_id = ?"
	// WHERE_STATIONERY_ID = "stationery_id = ?"
	// WHERE_STATUS        = "status = ?"
	// WHERE_NOT_STATUS    = "status != ?"

	// WHERE_BETWEEN_LAST_DAY           = "last_day BETWEEN ? AND ?"
	// WHERE_BETWEEN_CREATED_AT         = "created_at BETWEEN ? AND ?"
	// WHERE_RESIGN_ID_AND_CLEAR_PIC_ID = "resignation_id = ? AND clear_pic_id = ?"
	// WHERE_RESIGN_ID_AND_CREATED_BY   = "resign_id = ? AND created_by = ?"

	// // IN
	// WHERE_ID_IN       = "id IN (?)"
	// WHERE_ID_NOT_IN   = "id NOT IN (?)"
	// WHERE_NOTIFIER_IN = "notifier IN ?"

	// // COUNT
	// ASSIGNEE_COUNT = "assignee_id, count(*) as count"

	// // ORDER
	// ID_ASC           = "id ASC"
	// TRIP_ID_ASC      = "trip_id ASC"
	// TRIP_ID_DESC     = "trip_id DESC"
	// MANPOWER_ID_DESC = "manpower_id DESC"
	// CREATED_AT_DESC  = "created_at DESC"
	// CREATED_AT_ASC   = "created_at ASC"
	// CLEAR_ID_ASC     = "clear_id ASC"
	NAME_ASC = "name ASC"
)
