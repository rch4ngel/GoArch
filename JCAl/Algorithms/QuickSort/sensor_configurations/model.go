package sensorconfigs

import (
	"database/sql"
	"fmt"
	"github.com/ryu/hxs/config"
	"github.com/ryu/hxs/devices"
	"github.com/ryu/hxs/enum_tables"
	"time"
)

type SensorConfig struct {
	ID            int
	UpdatedAt     *time.Time
	DeviceId      sql.NullInt64
	GroupId       sql.NullInt64
	SensorTypeId  sql.NullInt64
	SensorUnitsId sql.NullInt64
	Enabled       sql.NullBool
	PollMillis    sql.NullFloat64
	SaveSeconds   sql.NullFloat64
	Interface     sql.NullString
	Channel       sql.NullInt64
	Index0        sql.NullInt64
	AverageCount  sql.NullFloat64
	MinValue      sql.NullFloat64
	MaxValue      sql.NullFloat64
	ChannelType   sql.NullInt64
	Index1        sql.NullInt64
	Ordinal       sql.NullInt64
	DeletedAt     *time.Time
	PublishMqtt   sql.NullBool
}

type Key struct {
	devices.Device
	enumtables.Enum
}

type SensorConfigMap map[Key]SensorConfig

func GetAllSensorConfigs() (scm SensorConfigMap, err error) {
	rows, err := config.DB.Query(
		"SELECT * " +
			"FROM " +
			"sensor_configs " +
			"ORDER BY id ASC")

	if err != nil {
		return nil, err
	}

	scm = SensorConfigMap{}
	for rows.Next() {
		sc := SensorConfig{}
		err := rows.Scan(
			&sc.ID,
			&sc.UpdatedAt,
			&sc.DeviceId,
			&sc.GroupId,
			&sc.SensorTypeId,
			&sc.SensorUnitsId,
			&sc.Enabled,
			&sc.PollMillis,
			&sc.SaveSeconds,
			&sc.Interface,
			&sc.Channel,
			&sc.Index0,
			&sc.AverageCount,
			&sc.MinValue,
			&sc.MaxValue,
			&sc.ChannelType,
			&sc.Index1,
			&sc.Ordinal,
			&sc.DeletedAt,
			&sc.PublishMqtt,
		)
		if err != nil {
			fmt.Println("Error creating sensor config:", err)
			return nil, err
		}

		d, err := devices.GetDevice(sc.DeviceId)
		if err != nil {
			fmt.Println("Cannot find Device:", d)
			return nil, err
		}

		e, err := enumtables.GetEnum(sc.SensorTypeId)

		if err != nil {
			return nil, err
		}
		scm[Key{d, e}] = sc
		fmt.Println("Adding Object:", sc)
	}

	return scm, nil
}
