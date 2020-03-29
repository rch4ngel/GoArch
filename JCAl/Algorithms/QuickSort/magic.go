package main

import (
	"fmt"
	"github.com/archangel/JCAl/Algorithms/QuickSort/devices"
	"github.com/archangel/JCAl/Algorithms/QuickSort/enum_tables"
	"github.com/archangel/JCAl/Algorithms/QuickSort/sensor_configurations"
	"github.com/archangel/JCAl/Algorithms/QuickSort/sensors"
)

type MagicEquipmentKey struct {
	Device devices.Device
	Sensor sensors.Sensor
	Enum   enumtables.Enum
}

type SensorMap map[MagicEquipmentKey]float64

func LoadSensorMagic() (map[MagicEquipmentKey]enumtables.Enum, error) {
	xs, err := sensors.AllSensors()
	if err != nil {
		return nil, err
	}

	xsc := []sensorconfigs.SensorConfig{}
	go func() {
		for _, s := range xs {
			sc, err := sensorconfigs.OneSensorConfig(s.SensorConfigId)
			if err != nil {
				fmt.Println(err)
			}
			xsc = append(xsc, sc)
		}
	}()

	xe := []enumtables.Enum{}
	go func() {
		for _, s := range xsc {
			e, err := enumtables.GetEnum(s.SensorTypeId)
			if err != nil {
				fmt.Println(err)
			}
			xe = append(xe, e)
		}
	}()

}
