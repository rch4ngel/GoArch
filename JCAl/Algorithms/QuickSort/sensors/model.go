package sensors

import (
	"database/sql"
	"github.com/ryu/hxs/config"
	"net/http"
	"time"
)

type Sensor struct {
	ID           sql.NullInt64
	Equipment    sql.NullString
	Enum         sql.NullString
	Interface    sql.NullString
	Device       sql.NullString
	SensorValues sql.NullString
}

type EquipmentSensors map[sql.NullString][]Sensor

func AllSensors() (map[sql.NullString][]Sensor, error) {
	es := EquipmentSensors{}
	rows, err := config.DB.Query(
		"SELECT " +
			"sensors.id," +
			"equipment.name," +
			"enum_tables.name," +
			"sensor_configs.interface," +
			"devices.name," +
			"sensors.sensor_values " +
			"FROM " +
			"sensors " +
			"inner join equipment  on sensors.equipment_id = equipment.id " +
			"inner join sensor_configs on sensors.config_id = sensor_configs.id " +
			"inner join enum_tables on sensor_configs.sensor_type_id = enum_tables.id  " +
			"inner join devices on sensor_configs.device_id = devices.id " +
			"ORDER BY equipment.name ASC")

	if err != nil {
		return es, err
	}

	for rows.Next() {
		s := Sensor{}
		err := rows.Scan(
			&s.ID,
			&s.Equipment,
			&s.Enum,
			&s.Interface,
			&s.Device,
			&s.SensorValues,
		)
		if err != nil {
			return nil, err
		}
		_, ok := es[s.Equipment]
		if !ok {
			xs := []Sensor{}
			xs = append(xs, s)
			es[s.Equipment] = xs
		} else {
			es[s.Equipment] = append(es[s.Equipment], s)
		}
	}

	//for k, v := range es {
	//	fmt.Printf("%s:\n", k.String)
	//	for _, v := range v {
	//		fmt.Printf("\t\t\t%s:\n", v.Equipment.String)
	//		fmt.Printf("\t\t\t%s:\n", v.Enum.String)
	//		fmt.Printf("\t\t\t%s:\n", v.SensorValues.String)
	//	}
	//}

	return es, nil
}

func OneSensor(r *http.Request) (map[sql.NullString][]Sensor, error) {
	es := EquipmentSensors{}
	e := r.FormValue("equipment")
	rows, err := config.DB.Query(
		"SELECT "+
			"sensors.id,"+
			"equipment.name,"+
			"enum_tables.name,"+
			"sensor_configs.interface,"+
			"devices.name,"+
			"sensors.sensor_values "+
			"FROM "+
			"sensors "+
			"inner join equipment  on sensors.equipment_id = equipment.id "+
			"inner join sensor_configs on sensors.config_id = sensor_configs.id "+
			"inner join enum_tables on sensor_configs.sensor_type_id = enum_tables.id  "+
			"inner join devices on sensor_configs.device_id = devices.id "+
			"WHERE equipment.name = $1", e)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		s := Sensor{}
		err := rows.Scan(
			&s.ID,
			&s.Equipment,
			&s.Enum,
			&s.Interface,
			&s.Device,
			&s.SensorValues,
		)

		if err != nil {
			return nil, err
		}
		_, ok := es[s.Equipment]
		if !ok {
			xs := []Sensor{}
			xs = append(xs, s)
			es[s.Equipment] = xs
		} else {
			es[s.Equipment] = append(es[s.Equipment], s)
		}
	}

	//for k, v := range es {
	//	fmt.Printf("%s:\n", k.String)
	//	for _, v := range v {
	//		fmt.Printf("\t\t\t%s:\n", v.Equipment.String)
	//		fmt.Printf("\t\t\t%s:\n", v.Enum.String)
	//		fmt.Printf("\t\t\t%s:\n", v.SensorValues.String)
	//	}
	//}

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
