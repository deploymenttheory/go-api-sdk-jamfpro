package jamfpro

// Type refers to string representation of target object type. I.e buildings, policies, computergroups

const (
	// Pagination - type: string, error: any
	errMsgFailedPaginatedGet = "failed to get paginated %s, error: %v"

	// CRUD - format always type: string, id/name: any, error: any
	errMsgFailedGetByID      = "failed to get %s by id: %v, error: %v"
	errMsgFailedGetByName    = "failed to get %s by name: %s, error: %v"
	errMsgFailedCreate       = "failed to create %s, error: %v"
	errMsgFailedUpdateByID   = "failed to update %s by id: %v, error: %v"
	errMsgFailedUpdateByName = "failed to update %s by name: %s, error: %v"
	errMsgFailedDeleteByID   = "failed to delete %s by id: %v, error: %v"
	errMsgFailedDeleteByName = "failed to delete %s by name: %s, error: %v"

	// Mapstructure - type: string, error: any
	errMsgFailedMapstruct = "failed to map interfaced %s to structs, error: %v"
)
