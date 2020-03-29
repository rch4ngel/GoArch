package sensors

import (
	"database/sql"
	"github.com/archangel/JCAl/Algorithms/QuickSort/config"
	"time"
)

type Sensor struct {
	ID                      int64
	EquipmentId             sql.NullInt64
	SensorStatusId          sql.NullInt64
	UpdatedAt               *time.Time
	SensorValues            sql.NullString
	TemporaryId             sql.NullInt64
	Offset0                 sql.NullFloat64
	Scale0                  sql.NullFloat64
	AlarmStatusId           sql.NullInt64
	SensorConfigId          sql.NullInt64
	ReferencePositionTypeId sql.NullFloat64
	ReferencePosition       sql.NullInt64
	AverageValues           sql.NullFloat64
}

type EquipmentSensors []Sensor

func AllSensors() ([]Sensor, error) {
	es := EquipmentSensors{}
	rows, err := config.DB.Query(
		"SELECT " +
			"* " +
			"FROM " +
			"sensors " +
			"ORDER BY " +
			"id ASC")

	if err != nil {
		return es, err
	}

	for rows.Next() {
		s := Sensor{}
		err := rows.Scan(
			&s.ID,
			&s.EquipmentId,
			&s.SensorStatusId,
			&s.UpdatedAt,
			&s.SensorValues,
			&s.TemporaryId,
			&s.Offset0,
			&s.Scale0,
			&s.AlarmStatusId,
			&s.SensorConfigId,
			&s.ReferencePositionTypeId,
			&s.ReferencePosition,
			&s.AverageValues,
		)

		if err != nil {
			return nil, err
		}

		es = append(es, s)
	}

	return es, nil
}

// Config structure and functionality is defined below.
type Config struct {
	ID          int64
	UpdatedAt   *time.Time
	DeviceName  sql.NullString
	Module      sql.NullString
	Attribute   sql.NullString
	Value       sql.NullString
	Application sql.NullString
}

type SensorTypeConfigurations map[sql.NullString][]Config

func AllSensorDeviceConfigs() (map[sql.NullString][]Config, error) {
	cm := SensorTypeConfigurations{}
	rows, err := config.DB.Query("" +
		"SELECT * " +
		"FROM " +
		"configurations " +
		"WHERE " +
		"device_id " +
		"IN " +
		"(" +
		"SELECT DISTINCT " +
		"devices.id " +
		"FROM sensors " +
		"INNER JOIN sensor_configs on sensors.config_id = sensor_configs.id " +
		"INNER JOIN devices on sensor_configs.device_id = devices.id order by devices.id asc" +
		") " +
		"AND application " +
		"LIKE 'JAMS';")
	if err != nil {
		return nil, err
	}

	c := Config{}
	for rows.Next() {
		rows.Scan(
			&c.ID,
			&c.UpdatedAt,
			&c.DeviceName,
			&c.Module,
			&c.Attribute,
			&c.Value,
			&c.Application,
		)
		_, ok := cm[c.DeviceName]
		if !ok {
			xc := []Config{}
			xc = append(xc, c)
			cm[c.DeviceName] = xc
		} else {
			cm[c.DeviceName] = append(cm[c.DeviceName], c)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cm, nil
}

func quickSort(sm map[string][]Sensor) {

}
