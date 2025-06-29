package dao

type StudentData struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	SystemAccess       bool   `json:"systemAccess"`
	Phone              string `json:"phone"`
	Gender             string `json:"gender"`
	DOB                string `json:"dob"`
	ClassName          string `json:"class"`
	SectionName        string `json:"section"`
	Roll               int    `json:"roll"`
	FatherName         string `json:"fatherName"`
	FatherPhone        string `json:"fatherPhone"`
	MotherName         string `json:"montherName"`
	MotherPhone        string `json:"motherPhone"`
	GuardianName       string `json:"guardianName"`
	GuardianPhone      string `json:"guardianPhone"`
	RelationOfGuardian string `json:"relationOfGuardian"`
	CurrentAddress     string `json:"currentAddress"`
	PermanentAddress   string `json:"permanentAddress"`
	AdmissionDate      string `json:"admissionDate"`
	ReporterName       string `json:"reporterName"`
}
