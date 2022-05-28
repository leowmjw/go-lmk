package main

import (
	"app/api/rippling"
	"fmt"
	"net/http"
)

type FakerRipplingServer struct{}

func (f FakerRipplingServer) PostAtsCandidatesPushCandidate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetCompanies(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetCompanyActivity(w http.ResponseWriter, r *http.Request, params rippling.GetCompanyActivityParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetCustomFields(w http.ResponseWriter, r *http.Request, params rippling.GetCustomFieldsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetDepartments(w http.ResponseWriter, r *http.Request, params rippling.GetDepartmentsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetEmployees(w http.ResponseWriter, r *http.Request, params rippling.GetEmployeesParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetEmployeesIncludeTerminated(w http.ResponseWriter, r *http.Request, params rippling.GetEmployeesIncludeTerminatedParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetEmployeesEmployeeId(w http.ResponseWriter, r *http.Request, employeeId rippling.EmployeeId) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetGroups(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) PostGroups(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) DeleteGroupsGroupId(w http.ResponseWriter, r *http.Request, groupId rippling.GroupId) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) PatchGroupsGroupId(w http.ResponseWriter, r *http.Request, groupId rippling.GroupId) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) PutGroupsGroupId(w http.ResponseWriter, r *http.Request, groupId rippling.GroupId) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetLeaveRequests(w http.ResponseWriter, r *http.Request, params rippling.GetLeaveRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) ProcessLeaveRequests(w http.ResponseWriter, r *http.Request, id string, params rippling.ProcessLeaveRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetLevels(w http.ResponseWriter, r *http.Request, params rippling.GetLevelsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) PostMarkAppInstalled(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetMe(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetSamlIdpMetadata(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetTeams(w http.ResponseWriter, r *http.Request, params rippling.GetTeamsParams) {
	//TODO implement me
	panic("implement me")
}

func (f FakerRipplingServer) GetWorkLocations(w http.ResponseWriter, r *http.Request, params rippling.GetWorkLocationsParams) {
	//TODO implement me
	panic("implement me")
}

func main() {

	fmt.Println("Welcome Lu Mau Kerja? (LMK)")

	frs := FakerRipplingServer{}
	http.ListenAndServe(":8080", rippling.Handler(frs))

}
