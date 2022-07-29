package freshservice

import "time"

// Applications holds a list of Freshservice application details
type Applications struct {
	List []ApplicationDetails `json:"applications"`
}

// Application holds the details of a specific Freshservice application
type Application struct {
	Details ApplicationDetails `json:"application"`
}

// ApplicationDetails are the details related to a specific application in Freshservice
type ApplicationDetails struct {
	AdditionalData    AdditionalData `json:"additional_data"`
	UserCount         int            `json:"user_count"`
	InstallationCount int            `json:"installation_count"`
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Description       interface{}    `json:"description"`
	Notes             interface{}    `json:"notes"`
	PublisherID       int64          `json:"publisher_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	ApplicationType   string         `json:"application_type"`
	Status            string         `json:"status"`
	ManagedByID       int64          `json:"managed_by_id"`
	Category          string         `json:"category"`
	Sources           []interface{}  `json:"sources"`
}

type AdditionalData struct {
	Overview     interface{} `json:"overview"`
	GraphData    interface{} `json:"graph_data"`
	LastSyncDate interface{} `json:"last_sync_date"`
}

// ApplicationListOptions holds the available options that can be
// passed when requesting a list of Freshservice Applications
type ApplicationListOptions struct {
	PageQuery string
}

// Licenses holds a list of Freshservice licenses for an application
type Licenses struct {
	List []LicensesDetails `json:"licenses"`
}

// LicenseDetails holds the details of a specific Freshservice application license
type LicensesDetails struct {
	ID          int       `json:"id"`
	ContractID  string    `json:"contract_id"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

// ApplicationUsers holds a list of Freshservice application users
type ApplicationUsers struct {
	List []ApplicationUserDetails `json:"application_users"`
}

// ApplicationUserDetails holds the details of users for a specific Freshservice application
type ApplicationUserDetails struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        int       `json:"user_id"`
	LicenseID     int       `json:"license_id"`
	AllocatedDate time.Time `json:"allocated_date"`
	FirstUsed     time.Time `json:"first_used"`
	LastUsed      time.Time `json:"last_used"`
	Source        string    `json:"source"`
}

// ApplicationInstallations holds a list of Freshservice application installations
type ApplicationInstallations struct {
	List []ApplicationInstallationDetails `json:"installations"`
}

// ApplicationInstallationDetails holds the details of installations for a specific Freshservice application
type ApplicationInstallationDetails struct {
	ID                    int       `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	InstallationPath      string    `json:"installation_path"`
	Version               string    `json:"version"`
	InstallationMachineID int       `json:"installation_machine_id"`
	UserID                int       `json:"user_id"`
	DepartmentID          int       `json:"department_id"`
	InstallationDate      time.Time `json:"installation_date"`
}
