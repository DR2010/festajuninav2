// Package btcmarketshandler API calls for dishes web
// --------------------------------------------------------------
// .../src/restauranteweb/areas/btcmarkets/btcmarketscalls.go
// --------------------------------------------------------------
package admshandler

import (
	"festajuninav2/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
)

// months probation 3.0
// training hours 32.0

type TrainingContractSummary struct {
	ID                            int    `json:"id"`
	StateTrainingAuthority        string `json:"stateTrainingAuthority"`
	ApprenticeID                  int    `json:"apprenticeId"`
	EmployerID                    int    `json:"employerId"`
	EmployerWorkplaceID           int    `json:"employerWorkplaceId"`
	ApprenticeshipID              int    `json:"apprenticeshipId"`
	ApprenticeFirstName           string `json:"apprenticeFirstName"`
	ApprenticeSurname             string `json:"apprenticeSurname"`
	EmployerBusinessName          string `json:"employerBusinessName"`
	Status                        string `json:"status"`
	StatusDate                    string `json:"statusDate"`
	ContractAwardName             string `json:"contractAwardName"`
	NetworkProviderContractSiteID int    `json:"networkProviderContractSiteId"`
	CreatedOn                     string `json:"createdOn"`
	UpdatedOn                     string `json:"updatedOn"`
}

// Training Contract
type TrainingContract struct {
	ID                                          int      `json:"id"`
	StateTrainingAuthority                      string   `json:"stateTrainingAuthority"`
	ApprenticeID                                int      `json:"apprenticeId"`
	EmployerID                                  int      `json:"employerId"`
	ApprenticeshipID                            int      `json:"apprenticeshipId"`
	ApprenticeshipIDStatusCode                  string   `json:"apprenticeshipIdStatusCode"`
	IsRecommencement                            bool     `json:"isRecommencement"`
	RecommencementApprenticeshipID              int      `json:"recommencementApprenticeshipId"`
	Status                                      string   `json:"status"`
	PriorQualificationsWithNoEffectOnIncentives []string `json:"priorQualificationsWithNoEffectOnIncentives"`
	StatusDate                                  string   `json:"statusDate"`
	SignedDate                                  string   `json:"signedDate"`
	Sta                                         struct {
		Reference string `json:"reference"`
	} `json:"sta"`
	Employer struct {
		Abn                      string `json:"abn"`
		EmployerAddressID        int    `json:"employerAddressId"`
		ContactID                int    `json:"contactId"`
		BusinessEntityID         int    `json:"businessEntityId"`
		BusinessName             string `json:"businessName"`
		WorkplaceID              int    `json:"workplaceId"`
		WorkplaceContactID       int    `json:"workplaceContactId"`
		EmployerIsGto            bool   `json:"employerIsGto"`
		HostOrganisationTypeCode string `json:"hostOrganisationTypeCode"`
		EmployerComments         string `json:"employerComments"`
		WorkplaceComments        string `json:"workplaceComments"`
	} `json:"employer"`
	Apprentice struct {
		Surname                           string `json:"surname"`
		FirstName                         string `json:"firstName"`
		OtherNames                        string `json:"otherNames"`
		BirthDate                         string `json:"birthDate"`
		Usi                               string `json:"usi"`
		UsiExempt                         bool   `json:"usiExempt"`
		UsiExemptReasonCode               string `json:"usiExemptReasonCode"`
		PhoneNumber                       string `json:"phoneNumber"`
		PhoneNumberInternationalPrefix    string `json:"phoneNumberInternationalPrefix"`
		EmailAddress                      string `json:"emailAddress"`
		Comments                          string `json:"comments"`
		CompletedSupersededQualification  bool   `json:"completedSupersededQualification"`
		IntendsToClaimDisabledWageSupport bool   `json:"intendsToClaimDisabledWageSupport"`
		GuardianRequired                  bool   `json:"guardianRequired"`
		GuardianNotAvailable              bool   `json:"guardianNotAvailable"`
		GuardianComments                  string `json:"guardianComments"`
	} `json:"apprentice"`
	ContractDetails struct {
		AwardID                      int     `json:"awardId"`
		AwardName                    string  `json:"awardName"`
		TrainingHours                float32 `json:"trainingHours"`
		AttendanceTypeCode           string  `json:"attendanceTypeCode"`
		SignedDate                   string  `json:"signedDate"`
		ApprenticeIsExistingWorker   bool    `json:"apprenticeIsExistingWorker"`
		ApprenticeIsPreviousEmployee bool    `json:"apprenticeIsPreviousEmployee"`
	} `json:"contractDetails"`
	Rto struct {
		Code                                   string `json:"code"`
		Title                                  string `json:"title"`
		ContactOfficerName                     string `json:"contactOfficerName"`
		ContactOfficerPhoneInternationalPrefix string `json:"contactOfficerPhoneInternationalPrefix"`
		ContactOfficerPhone                    string `json:"contactOfficerPhone"`
	} `json:"rto"`
	Apprenticeship struct {
		CommencementDate          string  `json:"commencementDate"`
		StateQualificationCode    string  `json:"stateQualificationCode"`
		NationalQualificationCode string  `json:"nationalQualificationCode"`
		QualificationTitle        string  `json:"qualificationTitle"`
		QualificationID           int     `json:"qualificationId"`
		Name                      string  `json:"name"`
		TypeCode                  string  `json:"typeCode"`
		LevelCode                 string  `json:"levelCode"`
		AnzscoCode                string  `json:"anzscoCode"`
		MonthsDuration            int     `json:"monthsDuration"`
		MonthsProbation           float32 `json:"monthsProbation"`
		NominalCompletionDate     string  `json:"nominalCompletionDate"`
		IsCustodial               bool    `json:"isCustodial"`
		CustodyReleaseDate        string  `json:"custodyReleaseDate"`
	} `json:"apprenticeship"`
	NetworkProvider struct {
		OrganisationCode                  string `json:"organisationCode"`
		SubmittedForReviewDate            string `json:"submittedForReviewDate"`
		ContractID                        int    `json:"contractId"`
		ContractStartDate                 string `json:"contractStartDate"`
		ContractEndDate                   string `json:"contractEndDate"`
		ContractSiteID                    int    `json:"contractSiteId"`
		ContractOutOfRegion               bool   `json:"contractOutOfRegion"`
		ContactName                       string `json:"contactName"`
		ContactEmail                      string `json:"contactEmail"`
		ContactPhone                      string `json:"contactPhone"`
		ContactPhoneInternationalPrefix   string `json:"contactPhoneInternationalPrefix"`
		ExclusionForIncentives            bool   `json:"exclusionForIncentives"`
		Comments                          string `json:"comments"`
		CommentsForStateTrainingAuthority string `json:"commentsForStateTrainingAuthority"`
	} `json:"networkProvider"`
	PriorQualifications struct {
		UnusableDueToInjury bool `json:"unusableDueToInjury"`
		JobactiveStreamBOrC bool `json:"jobactiveStreamBOrC"`
		CurrentlyUnemployed bool `json:"currentlyUnemployed"`
	} `json:"priorQualifications"`
	ProofOfIdentity struct {
		DocumentCategory                 string `json:"documentCategory"`
		OtherDocumentType                string `json:"otherDocumentType"`
		PhotographicDocumentTypeCode     string `json:"photographicDocumentTypeCode"`
		NonPhotographicDocumentTypeCode1 string `json:"nonPhotographicDocumentTypeCode1"`
		NonPhotographicDocumentTypeCode2 string `json:"nonPhotographicDocumentTypeCode2"`
		Sighted                          bool   `json:"sighted"`
		SightedComments                  string `json:"sightedComments"`
	} `json:"proofOfIdentity"`
	Education struct {
		AttendingSecondarySchool          bool   `json:"attendingSecondarySchool"`
		ApprovedSchoolBasedApprenticeship bool   `json:"approvedSchoolBasedApprenticeship"`
		TrainingPlanSignedByPrincipal     bool   `json:"trainingPlanSignedByPrincipal"`
		CreditSought                      bool   `json:"creditSought"`
		CreditEvidenceSighted             bool   `json:"creditEvidenceSighted"`
		MonthsCreditSought                int    `json:"monthsCreditSought"`
		SecondarySchoolLevelCode          string `json:"secondarySchoolLevelCode"`
		SecondarySchoolCode               string `json:"secondarySchoolCode"`
		SecondarySchoolName               string `json:"secondarySchoolName"`
	} `json:"education"`
	Employment struct {
		ApprenticesEmployed                         int    `json:"apprenticesEmployed"`
		SupervisorsEmployed                         int    `json:"supervisorsEmployed"`
		ContactPerson                               string `json:"contactPerson"`
		ContactPhone                                string `json:"contactPhone"`
		ContactPhoneInternationalPrefix             string `json:"contactPhoneInternationalPrefix"`
		ContactFax                                  string `json:"contactFax"`
		ContactEmail                                string `json:"contactEmail"`
		ArrangementCode                             string `json:"arrangementCode"`
		ApprenticeInBusinessRelationship            bool   `json:"apprenticeInBusinessRelationship"`
		BusinessRelationshipCode                    string `json:"businessRelationshipCode"`
		PreviousIncentives                          bool   `json:"previousIncentives"`
		PreviousIncentiveComments                   string `json:"previousIncentiveComments"`
		ConsideredExistingWorkerByDepartment        bool   `json:"consideredExistingWorkerByDepartment"`
		ExistingWorkerCompletedRecentApprenticeship bool   `json:"existingWorkerCompletedRecentApprenticeship"`
		ExistingWorkerMonths                        int    `json:"existingWorkerMonths"`
	} `json:"employment"`
	NewSouthWales struct {
		ApprenticeshipAlternateDuration      bool   `json:"apprenticeshipAlternateDuration"`
		RecommencementIsTransfer             bool   `json:"recommencementIsTransfer"`
		RecommencementDueToChangeOfOwnership bool   `json:"recommencementDueToChangeOfOwnership"`
		EmployerIsLabourHire                 bool   `json:"employerIsLabourHire"`
		EmployerHostOrganisationAbn          string `json:"employerHostOrganisationAbn"`
	} `json:"newSouthWales"`
	NorthernTerritory struct {
		ApprenticeWorkPhoneInternationalPrefix string `json:"apprenticeWorkPhoneInternationalPrefix"`
		ApprenticeWorkPhone                    string `json:"apprenticeWorkPhone"`
	} `json:"northernTerritory"`
	Queensland struct {
		ApprenticeHasResidency               bool `json:"apprenticeHasResidency"`
		EducationDispensationLetterReceived  bool `json:"educationDispensationLetterReceived"`
		RecommencementIsAvetmissCommencement bool `json:"recommencementIsAvetmissCommencement"`
	} `json:"queensland"`
	Tasmania struct {
		ApprenticeshipProbationDocumentationSupplied bool   `json:"apprenticeshipProbationDocumentationSupplied"`
		ApprenticeHasResidency                       bool   `json:"apprenticeHasResidency"`
		ContractFundingSourceCode                    string `json:"contractFundingSourceCode"`
		EducationDispensationLetterReceived          bool   `json:"educationDispensationLetterReceived"`
	} `json:"tasmania"`
	Victoria struct {
		RecommencementIsOverride bool `json:"recommencementIsOverride"`
	} `json:"victoria"`
	WesternAustralia struct {
		EmployerHostOrganisationAbn                    string `json:"employerHostOrganisationAbn"`
		SecondarySchoolContactPerson                   string `json:"secondarySchoolContactPerson"`
		SecondarySchoolContactPhone                    string `json:"secondarySchoolContactPhone"`
		SecondarySchoolContactPhoneInternationalPrefix string `json:"secondarySchoolContactPhoneInternationalPrefix"`
		ApprenticeshipOccupation                       string `json:"apprenticeshipOccupation"`
		RtoSuggestedCampus                             string `json:"rtoSuggestedCampus"`
		EducationReleaseScheduleSighted                bool   `json:"educationReleaseScheduleSighted"`
		EducationCurriculumCouncilNumber               int    `json:"educationCurriculumCouncilNumber"`
		EmployerContactPerson                          string `json:"employerContactPerson"`
		WorkplaceContactPhoneInternationalPrefix       string `json:"workplaceContactPhoneInternationalPrefix"`
		WorkplaceContactPhone                          string `json:"workplaceContactPhone"`
	} `json:"westernAustralia"`
	Bac struct {
		Eligible            bool     `json:"eligible"`
		DecisionReasons     []string `json:"decisionReasons"`
		EligibleOverride    bool     `json:"eligibleOverride"`
		ClaimEventPublished string   `json:"claimEventPublished"`
	} `json:"bac"`
	Version string `json:"version"`
}

// Row is
type Row struct {
	Description []string
}

// ControllerInfo is
type ControllerInfo struct {
	Name          string
	Message       string
	UserID        string
	UserName      string
	ApplicationID string //
	IsAdmin       string //
	Token         string //
}

type DisplayTemplate struct {
	Info              ControllerInfo
	FieldNames        []string
	Rows              []Row
	TrainingContracts []TrainingContractSummary
}

func AdmsIndex(httpwriter http.ResponseWriter, redisclient *redis.Client, credentials models.Credentials, sysid string) {

	// create new template
	t, _ := template.ParseFiles("html/index.html", "templates/adms/listtemplate.html")

	// Get list of users (api call)
	//
	// actlist, error := securityhandler.UserListAPI()
	// tcget := TrainingContractGet(httpwriter, redisclient, credentials)

	list := TrainingContractList(httpwriter, redisclient, credentials)

	token := GetToken()

	// Assemble the display structure for html template
	//
	items := DisplayTemplate{}
	items.Info.Name = "ADMS Training Contract List"
	items.Info.UserID = credentials.UserID
	items.Info.UserName = credentials.Name
	items.Info.ApplicationID = credentials.ApplicationID
	items.Info.IsAdmin = credentials.IsAdmin
	items.Info.Token = token

	var numberoffields = 5

	// Set colum names
	items.FieldNames = make([]string, numberoffields)
	items.FieldNames[0] = "TC ID"
	items.FieldNames[1] = "Apprentice Name"
	items.FieldNames[2] = "Apprentice DOB"
	items.FieldNames[3] = "TC Status"
	items.FieldNames[4] = "TC Status Date"

	// Set rows to be displayed
	items.Rows = make([]Row, len(list))
	items.TrainingContracts = make([]TrainingContractSummary, len(list))
	// items.RowID = make([]int, len(dishlist))

	for i := 0; i < len(list); i++ {
		items.Rows[i] = Row{}
		items.Rows[i].Description = make([]string, numberoffields)
		items.Rows[i].Description[0] = strconv.Itoa(list[i].ID)
		items.Rows[i].Description[1] = list[i].ApprenticeFirstName
		items.Rows[i].Description[2] = list[i].ApprenticeSurname
		items.Rows[i].Description[3] = list[i].Status
		items.Rows[i].Description[4] = list[i].StatusDate

		items.TrainingContracts[i] = list[i]
	}

	t.Execute(httpwriter, items)
}
