package equipment

import (
	"database/sql"
	"github.com/archangel/JCAl/Algorithms/QuickSort/activity"
	"github.com/archangel/JCAl/Algorithms/QuickSort/contractor"
	"github.com/archangel/JCAl/Algorithms/QuickSort/department"
	"github.com/archangel/JCAl/Algorithms/QuickSort/devices"
	"github.com/archangel/JCAl/Algorithms/QuickSort/enum_tables"
	"github.com/archangel/JCAl/Algorithms/QuickSort/project"
	"github.com/archangel/JCAl/Algorithms/QuickSort/reason"
	"github.com/archangel/JCAl/Algorithms/QuickSort/status"
	"github.com/archangel/JCAl/Algorithms/QuickSort/tiedown"
	"github.com/archangel/JCAl/Algorithms/QuickSort/unit"
	"time"
)

type Equipment struct {
	ID                 int64
	Type               sql.NullString
	Name               string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
	Revision           sql.NullInt64
	Device             devices.Device
	Status             status.Status
	Reason             reason.Reason
	LineupStatus       string
	LineUpReason       enumtables.Enum
	Activity           activity.Activity
	ActivityStart      *time.Time
	Size               sql.NullFloat64
	Tiedown            tiedown.Tiedown
	Unit               unit.Unit
	Length             sql.NullFloat64
	Height             sql.NullFloat64
	Project            project.Project
	TimeLate           *time.Time
	PrestartCheck      sql.NullBool
	Warnings           string
	HpEquipmentType    enumtables.Enum
	Department         department.Department
	LineupComment      sql.NullString
	NoAutoAccept       sql.NullBool
	Contractor         contractor.Contractor
	Weight             sql.NullInt64
	RepairStatusId     sql.NullInt64
	LineupPriorityId   sql.NullInt64
	GeometryNodeSetId  sql.NullInt64
	GeometryParamSetId sql.NullInt64
	IconSetId          sql.NullInt64
	LastActivity       activity.Activity
}
