package devices

import (
	"database/sql"
	"github.com/ryu/hxs/config"
	"time"
)

type Device struct {
	ID                          int
	Name                        sql.NullString
	UpdateAt                    *time.Time
	Address                     sql.NullString
	Netmask                     sql.NullString
	MacAddress                  sql.NullString
	Utilities                   sql.NullString
	SoftwareVersion             sql.NullString
	InstalledVersion            sql.NullString
	DeviceType                  sql.NullString
	Groups                      sql.NullString
	PlatformVersion             sql.NullString
	PlatformArchitecture        sql.NullString
	PlatformAvrVersion          sql.NullString
	PlatformGps1FirmwareVersion sql.NullString
	PlatformGps2FirmwareVersion sql.NullString
}

func GetDevice(id sql.NullInt64) (Device, error) {
	row := config.DB.QueryRow(
		"SELECT "+
			"* "+
			"FROM "+
			"devices "+
			"WHERE "+
			"id = $1",
		id)

	d := Device{}
	err := row.Scan(
		&d.ID,
		&d.Name,
		&d.UpdateAt,
		&d.Address,
		&d.Netmask,
		&d.MacAddress,
		&d.Utilities,
		&d.SoftwareVersion,
		&d.InstalledVersion,
		&d.DeviceType,
		&d.Groups,
		&d.PlatformVersion,
		&d.PlatformArchitecture,
		&d.PlatformAvrVersion,
		&d.PlatformGps1FirmwareVersion,
		&d.PlatformGps2FirmwareVersion,
	)
	if err != nil {
		return d, nil
	}

	return d, nil
}
