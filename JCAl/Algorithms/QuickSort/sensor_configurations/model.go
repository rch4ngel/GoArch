package sensorconfigs

import (
	"database/sql"
	"fmt"
	"github.com/ryu/hxs/config"
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

type SensorConfigs []SensorConfig

func GetAllSensorConfigs() (SensorConfigs, error) {
	rows, err := config.DB.Query(
		"SELECT * " +
			"FROM " +
			"sensor_configs " +
			"ORDER BY id ASC")

	if err != nil {
		return nil, err
	}

	xsc := SensorConfigs{}
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

		if err != nil {
			return nil, err
		}
		xsc = append(xsc, sc)
	}

	return xsc, nil
}

func OneSensorConfig(id sql.NullInt64) (SensorConfig, error) {
	row := config.DB.QueryRow(
		"SELECT "+
			"* "+
			"FROM "+
			"sensor_configs "+
			"WHERE "+
			"id = $1", id)

	sc := SensorConfig{}
	err := row.Scan(
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

	switch {
	case err == sql.ErrNoRows:
		fmt.Println(err)
		return sc, err
	case err != nil:
		panic(err)
		return sc, err
	}

	return sc, nil
}
