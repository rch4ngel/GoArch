package equipment

type Device devices.Device
type Reason reason.Reason
type Status status.Status
type Enum enumtables.Enum
type Activity activity.Activity
type Contractor contractor.Contracter
type Tiedown tiedown.Tiedown
type Unit unit.Unit
type Project project.Project
type Department department.Department

type Equipment struct {
	ID        int64
	Type      string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Revision  sql.NullInt64
	Device
	Status
	Reason
	LineupStatus string
	Enum
	Activity
	ActivityStart *time.Time
	Size          sql.NullFloat64
	Tiedown
	Unit
	Length sql.NullFloat64
	Height sql.NullFloat64
	Project
	TimeLate      *time.Time
	PrestartCheck sql.NullBool
	Warnings      string
	Enum
	Department
	LineupComment sql.NullString
	NoAutoAccept  sql.NullBool
	Contractor
	Weight             sql.NullInt64
	RepairStatusId     sql.NullInt64
	LineupPriorityId   sql.NullInt64
	GeometryNodeSetId  sql.NullInt64
	GeometryParamSetId sql.NullInt64
	IconSetId          sql.NullInt64
	Activity
}
