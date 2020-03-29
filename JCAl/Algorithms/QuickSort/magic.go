package main

type Sensor struct {
	SensorName  string
	SensorValue float64
}

type Key struct {
	Equipment string
	Sensor
}

type SensorMap map[Key]float64

func LoadSensors() SensorMap {
	sm := SensorMap{}
	s1 := Sensor{SensorName: "Hour Meter", SensorValue: 23.43}
	s2 := Sensor{SensorName: "Hour Meter", SensorValue: 43}
	s3 := Sensor{SensorName: "Hour Meter", SensorValue: 554}
	s4 := Sensor{SensorName: "Hour Meter", SensorValue: 9923}

	s5 := Sensor{SensorName: "Speedometer", SensorValue: 532}
	s6 := Sensor{SensorName: "Speedometer", SensorValue: 9482}
	s7 := Sensor{SensorName: "Speedometer", SensorValue: 4443}
	s8 := Sensor{SensorName: "Speedometer", SensorValue: 2111}

	sm["DR02"] = []Sensor{s1, s5}
	sm["TR006"] = []Sensor{s2, s6}
	sm["SH034"] = []Sensor{s3, s7}
	sm["LDR93"] = []Sensor{s4, s8}

	return sm
}
