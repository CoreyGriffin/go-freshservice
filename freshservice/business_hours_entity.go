package freshservice

import "time"

//BusinessHours holds the Business Hours Configurations in Freshservice
type BusinessHours struct {
	List []BusinessHoursDetails `json:"business_hours"`
}

// BusinessHoursConfig holds a configuration for business hours in Freshservice
type BusinessHoursConfig struct {
	Details BusinessHoursDetails `json:"business_hours"`
}

// BusinessHoursDetails holds a configuration for business hours in Freshservice
type BusinessHoursDetails struct {
	ID               int              `json:"id"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	IsDefault        bool             `json:"is_default"`
	TimeZone         string           `json:"time_zone"`
	ServiceDeskHours ServiceDeskHours `json:"service_desk_hours"`
	ListOfHolidays   []WorkdayHoliday `json:"list_of_holidays"`
}

// ServiceDeskHours contains the time at which the workday begins and ends for the seven days of the week.
type ServiceDeskHours struct {
	Monday    WorkdayHours `json:"monday"`
	Tuesday   WorkdayHours `json:"tuesday"`
	Wednesday WorkdayHours `json:"wednesday"`
	Thursday  WorkdayHours `json:"thursday"`
	Friday    WorkdayHours `json:"friday"`
}

// WorkdayHours contains the time at which the workday begins and ends
type WorkdayHours struct {
	BeginningOfWorkday string `json:"beginning_of_workday"`
	EndOfWorkday       string `json:"end_of_workday"`
}

// WorkdayHoliday holds a configured holiday for the year. Dates are in ISO --MM-DD format.
type WorkdayHoliday struct {
	HolidayDate string `json:"holiday_date"`
	HolidayName string `json:"holiday_name"`
}
