// Code generated by ogen, DO NOT EDIT.

package rippling

import (
	"time"

	"github.com/go-faster/jx"
)

// SetFake set fake values.
func (s *Address) SetFake() {
	{
		{
			s.City.SetFake()
		}
	}
	{
		{
			s.Country.SetFake()
		}
	}
	{
		{
			s.State.SetFake()
		}
	}
	{
		{
			s.StreetLine1.SetFake()
		}
	}
	{
		{
			s.StreetLine2.SetFake()
		}
	}
	{
		{
			s.Zip.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *AuthenticatedUserMe) SetFake() {
	{
		{
			s.Company.SetFake()
		}
	}
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.WorkEmail.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Candidate) SetFake() {
	{
		{
			s.Attachments = nil
			for i := 0; i < 0; i++ {
				var elem CandidateAttachmentsItem
				{
					elem.SetFake()
				}
				s.Attachments = append(s.Attachments, elem)
			}
		}
	}
	{
		{
			s.CandidateId.SetFake()
		}
	}
	{
		{
			s.Currency.SetFake()
		}
	}
	{
		{
			s.Department.SetFake()
		}
	}
	{
		{
			s.Email.SetFake()
		}
	}
	{
		{
			s.EmploymentType.SetFake()
		}
	}
	{
		{
			s.EquityShares.SetFake()
		}
	}
	{
		{
			s.JobTitle.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.PhoneNumber.SetFake()
		}
	}
	{
		{
			s.SalaryPerUnit.SetFake()
		}
	}
	{
		{
			s.SalaryUnit.SetFake()
		}
	}
	{
		{
			s.SigningBonus.SetFake()
		}
	}
	{
		{
			s.StartDate.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *CandidateAttachmentsItem) SetFake() {
	{
		{
			s.FileName.SetFake()
		}
	}
	{
		{
			s.FileURL.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *CandidateEmploymentType) SetFake() {
	*s = CandidateEmploymentTypeCONTRACTOR
}

// SetFake set fake values.
func (s *CandidateSalaryUnit) SetFake() {
	*s = CandidateSalaryUnitHOUR
}

// SetFake set fake values.
func (s *Company) SetFake() {
	{
		{
			s.Address.SetFake()
		}
	}
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Phone.SetFake()
		}
	}
	{
		{
			s.PrimaryEmail.SetFake()
		}
	}
	{
		{
			s.WorkLocations = nil
			for i := 0; i < 0; i++ {
				var elem WorkLocation
				{
					elem.SetFake()
				}
				s.WorkLocations = append(s.WorkLocations, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *CustomFields) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Required.SetFake()
		}
	}
	{
		{
			s.Title.SetFake()
		}
	}
	{
		{
			s.Type.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *CustomFieldsType) SetFake() {
	*s = CustomFieldsTypeTEXT
}

// SetFake set fake values.
func (s *Department) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Parent.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Employee) SetFake() {
	{
		{
			s.CustomFields.SetFake()
		}
	}
	{
		{
			s.Department.SetFake()
		}
	}
	{
		{
			s.EmploymentType.SetFake()
		}
	}
	{
		{
			s.EndDate.SetFake()
		}
	}
	{
		{
			s.FirstName.SetFake()
		}
	}
	{
		{
			s.Gender.SetFake()
		}
	}
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.IdentifiedGender.SetFake()
		}
	}
	{
		{
			s.LastName.SetFake()
		}
	}
	{
		{
			s.Manager.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.RoleState.SetFake()
		}
	}
	{
		{
			s.SpokeId.SetFake()
		}
	}
	{
		{
			s.Title.SetFake()
		}
	}
	{
		{
			s.WorkEmail.SetFake()
		}
	}
	{
		{
			s.WorkLocation.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EmployeeEmploymentType) SetFake() {
	*s = EmployeeEmploymentTypeCONTRACTOR
}

// SetFake set fake values.
func (s *EmployeeGender) SetFake() {
	*s = EmployeeGenderMALE
}

// SetFake set fake values.
func (s *EmployeeIdentifiedGender) SetFake() {
	*s = EmployeeIdentifiedGenderMALE
}

// SetFake set fake values.
func (s *EmployeeRoleState) SetFake() {
	*s = EmployeeRoleStateINIT
}

// SetFake set fake values.
func (s *Event) SetFake() {
	{
		{
			s.Company.SetFake()
		}
	}
	{
		{
			s.EventReason.SetFake()
		}
	}
	{
		{
			s.EventType.SetFake()
		}
	}
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Initiator.SetFake()
		}
	}
	{
		{
			s.LinkedEvents = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.LinkedEvents = append(s.LinkedEvents, elem)
			}
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.RequestData.SetFake()
		}
	}
	{
		{
			s.Spoke.SetFake()
		}
	}
	{
		{
			s.Subjects = nil
			for i := 0; i < 0; i++ {
				var elem EventSubjectsItem
				{
					elem.SetFake()
				}
				s.Subjects = append(s.Subjects, elem)
			}
		}
	}
	{
		{
			s.Timestamp.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EventEventReason) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Reason.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EventEventType) SetFake() {
	*s = EventEventTypeEXTERNALACCOUNTCREATE
}

// SetFake set fake values.
func (s *EventInitiator) SetFake() {
	{
		{
			s.DisplayName.SetFake()
		}
	}
	{
		{
			s.Icon.SetFake()
		}
	}
	{
		{
			s.Role.SetFake()
		}
	}
	{
		{
			s.Type.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EventInitiatorType) SetFake() {
	*s = EventInitiatorTypeROLE
}

// SetFake set fake values.
func (s *EventRequestData) SetFake() {
	{
		{
			s.City.SetFake()
		}
	}
	{
		{
			s.Country.SetFake()
		}
	}
	{
		{
			s.IP.SetFake()
		}
	}
	{
		{
			s.Latitude.SetFake()
		}
	}
	{
		{
			s.Longitude.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EventSubjectsItem) SetFake() {
	{
		{
			s.DisplayName.SetFake()
		}
	}
	{
		{
			s.Icon.SetFake()
		}
	}
	{
		{
			s.Instance.SetFake()
		}
	}
	{
		{
			s.Type.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *EventSubjectsItemType) SetFake() {
	*s = EventSubjectsItemTypeROLE
}

// SetFake set fake values.
func (s *GetCompanyActivityOK) SetFake() {
	{
		{
			s.Data.SetFake()
		}
	}
	{
		{
			s.Error.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *GetCompanyActivityOKData) SetFake() {
	{
		{
			s.Events.SetFake()
		}
	}
	{
		{
			s.Next.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Group) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.SpokeId.SetFake()
		}
	}
	{
		{
			s.Users = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.Users = append(s.Users, elem)
			}
		}
	}
	{
		{
			s.Version.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *GroupUpdatePayload) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.SpokeId.SetFake()
		}
	}
	{
		{
			s.Users = nil
			for i := 0; i < 0; i++ {
				var elem jx.Raw
				{
					elem = []byte("null")
				}
				s.Users = append(s.Users, elem)
			}
		}
	}
	{
		{
			s.Version.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LeaveRequest) SetFake() {
	{
		{
			s.Comments.SetFake()
		}
	}
	{
		{
			s.CreatedAt.SetFake()
		}
	}
	{
		{
			s.Dates = nil
			for i := 0; i < 0; i++ {
				var elem LeaveRequestDatesItem
				{
					elem.SetFake()
				}
				s.Dates = append(s.Dates, elem)
			}
		}
	}
	{
		{
			s.EndDate.SetFake()
		}
	}
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.LeavePolicy.SetFake()
		}
	}
	{
		{
			s.LeaveType.SetFake()
		}
	}
	{
		{
			s.NumHours.SetFake()
		}
	}
	{
		{
			s.NumMinutes.SetFake()
		}
	}
	{
		{
			s.PolicyDisplayName.SetFake()
		}
	}
	{
		{
			s.ProcessedAt.SetFake()
		}
	}
	{
		{
			s.ProcessedBy.SetFake()
		}
	}
	{
		{
			s.ProcessedByName.SetFake()
		}
	}
	{
		{
			s.ReasonForLeave.SetFake()
		}
	}
	{
		{
			s.RequestedBy.SetFake()
		}
	}
	{
		{
			s.RequestedByName.SetFake()
		}
	}
	{
		{
			s.Role.SetFake()
		}
	}
	{
		{
			s.RoleName.SetFake()
		}
	}
	{
		{
			s.StartDate.SetFake()
		}
	}
	{
		{
			s.Status.SetFake()
		}
	}
	{
		{
			s.UpdatedAt.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LeaveRequestDatesItem) SetFake() {
	{
		{
			s.Date.SetFake()
		}
	}
	{
		{
			s.NumMinutes.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LeaveRequestLeaveType) SetFake() {
	*s = LeaveRequestLeaveTypeVACATION
}

// SetFake set fake values.
func (s *LeaveRequestStatus) SetFake() {
	*s = LeaveRequestStatusPENDING
}

// SetFake set fake values.
func (s *Level) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Parent.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OptBool) SetFake() {
	var elem bool
	{
		elem = true
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCandidate) SetFake() {
	var elem Candidate
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCandidateEmploymentType) SetFake() {
	var elem CandidateEmploymentType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCandidateSalaryUnit) SetFake() {
	var elem CandidateSalaryUnit
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCustomFields) SetFake() {
	var elem CustomFields
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCustomFieldsType) SetFake() {
	var elem CustomFieldsType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDate) SetFake() {
	var elem time.Time
	{
		elem = time.Now()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEmployeeEmploymentType) SetFake() {
	var elem EmployeeEmploymentType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEmployeeRoleState) SetFake() {
	var elem EmployeeRoleState
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEventEventType) SetFake() {
	var elem EventEventType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEventInitiatorType) SetFake() {
	var elem EventInitiatorType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEventSubjectsItemType) SetFake() {
	var elem EventSubjectsItemType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptFloat64) SetFake() {
	var elem float64
	{
		elem = float64(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptGroupUpdatePayload) SetFake() {
	var elem GroupUpdatePayload
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt) SetFake() {
	var elem int
	{
		elem = int(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLeaveRequestLeaveType) SetFake() {
	var elem LeaveRequestLeaveType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLeaveRequestStatus) SetFake() {
	var elem LeaveRequestStatus
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptNilAddress) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEmployeeGender) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEmployeeIdentifiedGender) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEventArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEventEventReason) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEventInitiator) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilEventRequestData) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilGetCompanyActivityOKData) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilInt) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilString) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptPostGroupsReq) SetFake() {
	var elem PostGroupsReq
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptString) SetFake() {
	var elem string
	{
		elem = "string"
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *PostGroupsReq) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.SpokeId.SetFake()
		}
	}
	{
		{
			s.Users = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.Users = append(s.Users, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *PostMarkAppInstalledOK) SetFake() {
	{
		{
			s.Ok = true
		}
	}
}

// SetFake set fake values.
func (s *Team) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Parent.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *WorkLocation) SetFake() {
	{
		{
			s.Address.SetFake()
		}
	}
	{
		{
			s.Nickname.SetFake()
		}
	}
}