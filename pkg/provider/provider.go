package provider

import "time"

// DefaultTimeout - for main purposes
const DefaultTimeout = 4 * time.Second

// UserAgent - user agent for http client
const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:84.0) Gecko/20100101 Firefox/84.0"

// URIPatternFZ44Purchase - pattern for parse purchases
const URIPatternFZ44Purchase = "https://zakupki.gov.ru/epz/order/notice/ea44/view/common-info.html?regNumber=%d&formatInJson=true"

type EntryDto struct { //TODO add filters for api
}

// Purchase - represent table columns
type Purchase struct {
	Id                          string
	Fz                          string
	Customer                    string
	CustomerLink                string
	CustomerInn                 string
	CustomerRegion              string
	BiddingRegion               string
	CustomerActivityField       string
	BiddingVolume               string
	BiddingCount                string
	PurchaseTarget              string
	RegistryBiddingNumber       string
	ContractPrice               string
	ParticipationSecurityAmount string
	ExecutionSecurityAmount     string
	PublishedAt                 string
	RequisitionDeadlineAt       string
	ContractStartAt             string
	ContractEndAt               string
	Playground                  string
	PurchaseLink                string
}

// Provider - represent table columns
type Provider struct {
	Id               string
	Fz               string
	PublishedAt      string
	Nmck             string
	SecurityAmount   string
	Winner           string
	ProviderRegion   string
	Address          string
	LeaderFio        string
	INN              string
	Email            string
	Email2           string
	WinCount         string
	ParticipateCount string
	LastWinDate      string
	Purpose          string
	ProcedureType    string
	Playground       string
	Customer         string
	WinCountYear     string
	Bad              string
	ActivityField    string
	PhoneNumber      string
}

// Dto44fz - represent json response from purchase by 44fz
type Dto44fz struct {
	Dto struct {
		HeaderBlock struct {
			CapitalRepairsLink string `json:"capitalRepairsLink"`
			ChangedDate        int64  `json:"changedDate"`
			ComplaintsDto      struct {
				ComplaintID     interface{} `json:"complaintId"`
				ComplaintNumber interface{} `json:"complaintNumber"`
				HasComplaint    bool        `json:"hasComplaint"`
			} `json:"complaintsDto"`
			CreatedDate             int64       `json:"createdDate"`
			CryptoSignLink          string      `json:"cryptoSignLink"`
			CurrencySymbol          string      `json:"currencySymbol"`
			CustomerOrgType         bool        `json:"customerOrgType"`
			ExpirationDate          int64       `json:"expirationDate"`
			IsCentralized           bool        `json:"isCentralized"`
			Nmck                    float64     `json:"nmck"`
			OrderBlockingDTO        interface{} `json:"orderBlockingDTO"`
			OrderShared             bool        `json:"orderShared"`
			OrderType               string      `json:"orderType"`
			OrganizationPublishLink string      `json:"organizationPublishLink"`
			OrganizationPublishName string      `json:"organizationPublishName"`
			PdiscLink               interface{} `json:"pdiscLink"`
			PlacingWayFZ            string      `json:"placingWayFZ"`
			PlacingWayName          string      `json:"placingWayName"`
			PrintFormLink           string      `json:"printFormLink"`
			ProtocolName            interface{} `json:"protocolName"`
			PublishedDate           int64       `json:"publishedDate"`
			PublishedEtpDate        interface{} `json:"publishedEtpDate"`
			PurchaseNumber          string      `json:"purchaseNumber"`
			PurchaseObjectName      string      `json:"purchaseObjectName"`
			StageNumber             int         `json:"stageNumber"`
			StageOfPurchase         string      `json:"stageOfPurchase"`
			SubsystemType           interface{} `json:"subsystemType"`
			Suspend                 bool        `json:"suspend"`
			TargetEA                interface{} `json:"targetEA"`
			TimeZoneAbbrev          string      `json:"timeZoneAbbrev"`
			MultiLots               bool        `json:"multiLots"`
		} `json:"headerBlock"`
		OrderType string `json:"orderType"`
		TabsBlock struct {
			ActualNotices []struct {
				ChangeLotNum  interface{} `json:"changeLotNum"`
				DocNum        string      `json:"docNum"`
				DocTypeCode   string      `json:"docTypeCode"`
				OrderType     string      `json:"orderType"`
				PrintFormLink struct {
					Link  string      `json:"link"`
					Show  bool        `json:"show"`
					Title interface{} `json:"title"`
				} `json:"printFormLink"`
				PublishDate    int64  `json:"publishDate"`
				TimeZoneAbbrev string `json:"timeZoneAbbrev"`
			} `json:"actualNotices"`
			ChangedStProtocol      bool        `json:"changedStProtocol"`
			CommonLink             string      `json:"commonLink"`
			DecisionInfoLink       interface{} `json:"decisionInfoLink"`
			DocLink                string      `json:"docLink"`
			DocTypeCode            string      `json:"docTypeCode"`
			EventLink              string      `json:"eventLink"`
			LotsLink               string      `json:"lotsLink"`
			ManyLots               bool        `json:"manyLots"`
			PlacingWay             string      `json:"placingWay"`
			PlacingWayCode         string      `json:"placingWayCode"`
			ProtocolBidListLink    interface{} `json:"protocolBidListLink"`
			ProtocolBidReviewLink  interface{} `json:"protocolBidReviewLink"`
			ProtocolCancelInfoLink interface{} `json:"protocolCancelInfoLink"`
			ProtocolChangeLink     interface{} `json:"protocolChangeLink"`
			ProtocolDocsLink       interface{} `json:"protocolDocsLink"`
			ProtocolMainInfoLink   interface{} `json:"protocolMainInfoLink"`
			ProtocolName           interface{} `json:"protocolName"`
			RegNumber              string      `json:"regNumber"`
			ShowSupplierResultsTab bool        `json:"showSupplierResultsTab"`
			StructuredProtocolTabs interface{} `json:"structuredProtocolTabs"`
			SupplierResultsLink    string      `json:"supplierResultsLink"`
			TabParams              interface{} `json:"tabParams"`
			TimeZoneAbbrev         string      `json:"timeZoneAbbrev"`
		} `json:"tabsBlock"`
		FeatureFCSNF9370 bool `json:"featureFCSNF9370"`
		ConditionsBlock  struct {
			AdditionalInformation             interface{} `json:"additionalInformation"`
			Advantages                        []string    `json:"advantages"`
			AllRequirementsForParticipantsMap struct {
			} `json:"allRequirementsForParticipantsMap"`
			HasComment                bool `json:"hasComment"`
			HasConditionsForExclusion bool `json:"hasConditionsForExclusion"`
			LotID                     int  `json:"lotId"`
			LotRestrictionNpaDtoList  []struct {
				AddInfo   interface{} `json:"addInfo"`
				Exclusion bool        `json:"exclusion"`
				NpaID     int         `json:"npaId"`
				NpaText   string      `json:"npaText"`
				Reason    interface{} `json:"reason"`
				TypeID    int         `json:"typeId"`
				TypeText  string      `json:"typeText"`
			} `json:"lotRestrictionNpaDtoList"`
			LotRestrictionNpaPagingDto struct {
				CriteriaPageNumber int    `json:"criteriaPageNumber"`
				PageCount          int    `json:"pageCount"`
				RecordsPerPage     string `json:"recordsPerPage"`
				Rounded            bool   `json:"rounded"`
				Total              int    `json:"total"`
				TotalRounded       int    `json:"totalRounded"`
				RssLink            string `json:"rssLink"`
				MoreThanTen        bool   `json:"moreThanTen"`
				ExceedCount        bool   `json:"exceedCount"`
			} `json:"lotRestrictionNpaPagingDto"`
			LotRestrictionNpaURL              string `json:"lotRestrictionNpaUrl"`
			NotAbleToAbideJustification       bool   `json:"notAbleToAbideJustification"`
			ParentRequirementsForParticipants []struct {
				AdditionalInformation interface{} `json:"additionalInformation"`
				FeatureID             int         `json:"featureId"`
				Lvl                   interface{} `json:"lvl"`
				Name                  string      `json:"name"`
				ParentID              interface{} `json:"parentId"`
			} `json:"parentRequirementsForParticipants"`
			RestrictionsAndProhibitions []struct {
				AdditionalInformation interface{} `json:"additionalInformation"`
				Name                  string      `json:"name"`
			} `json:"restrictionsAndProhibitions"`
			RestrictsVendorSelect interface{} `json:"restrictsVendorSelect"`
			ShowConditions        bool        `json:"showConditions"`
		} `json:"conditionsBlock"`
		CustomerRequirementsBlock []struct {
			BankSupportBlock struct {
				BankSupportType     string `json:"bankSupportType"`
				BankSupportTypeText string `json:"bankSupportTypeText"`
			} `json:"bankSupportBlock"`
			ConditionsOfContract struct {
				BtkNumber              interface{} `json:"btkNumber"`
				Currency               interface{} `json:"currency"`
				DeliveryPlace          []string    `json:"deliveryPlace"`
				DeliveryTime           string      `json:"deliveryTime"`
				InitialContractPrice   interface{} `json:"initialContractPrice"`
				InterBudgetaryTransfer interface{} `json:"interBudgetaryTransfer"`
				Show                   bool        `json:"show"`
				UnilateralRefusal      interface{} `json:"unilateralRefusal"`
			} `json:"conditionsOfContract"`
			CustomerID                  int         `json:"customerId"`
			CustomerName                string      `json:"customerName"`
			DiscussionNeed              interface{} `json:"discussionNeed"`
			EnsuringPerformanceContract struct {
				AdditionalInformation       interface{} `json:"additionalInformation"`
				AmountContractEnforcement   interface{} `json:"amountContractEnforcement"`
				ContractGrntShare           int         `json:"contractGrntShare"`
				ContractualSecurityRequired string      `json:"contractualSecurityRequired"`
				Currency                    string      `json:"currency"`
				EnforcementProcedure        string      `json:"enforcementProcedure"`
				OfferGrnt                   bool        `json:"offerGrnt"`
				PaymentRequisites           struct {
					Bik string `json:"bik"`
					Ls  string `json:"ls"`
					Rs  string `json:"rs"`
				} `json:"paymentRequisites"`
				SmpSono               bool `json:"smpSono"`
				EnergyServiceContract bool `json:"energyServiceContract"`
			} `json:"ensuringPerformanceContract"`
			EnsuringPurchase struct {
				AmountEnforcement    interface{} `json:"amountEnforcement"`
				Currency             interface{} `json:"currency"`
				EnforcementProcedure interface{} `json:"enforcementProcedure"`
				OfferGrnt            bool        `json:"offerGrnt"`
				PaymentRequisites    interface{} `json:"paymentRequisites"`
				SecurityRequired     string      `json:"securityRequired"`
			} `json:"ensuringPurchase"`
			FinanceSource        interface{} `json:"financeSource"`
			InitialContractPrice struct {
				AdvancePercent      interface{} `json:"advancePercent"`
				ContractCurrency    interface{} `json:"contractCurrency"`
				ContractPaymentPlan struct {
					AnnualPaymentPlanTable struct {
						BudgetRowList               []interface{} `json:"budgetRowList"`
						BudgetTotalRowTitle         string        `json:"budgetTotalRowTitle"`
						ColumnAmount                int           `json:"columnAmount"`
						CommonTotalRowTitle         string        `json:"commonTotalRowTitle"`
						ExtrabudgetRowList          []interface{} `json:"extrabudgetRowList"`
						ExtrabudgetTotalRowTitle    string        `json:"extrabudgetTotalRowTitle"`
						FirstYearAmountColumnHeader string        `json:"firstYearAmountColumnHeader"`
						HasFinancialSourceColumn    bool          `json:"hasFinancialSourceColumn"`
						HasFirstYearAmountColumn    bool          `json:"hasFirstYearAmountColumn"`
						HasKbkKosguKvrColumn        bool          `json:"hasKbkKosguKvrColumn"`
						HasOtherYearAmountColumn    bool          `json:"hasOtherYearAmountColumn"`
						HasSecondYearAmountColumn   bool          `json:"hasSecondYearAmountColumn"`
						HasTable                    bool          `json:"hasTable"`
						HasThirdYearAmountColumn    bool          `json:"hasThirdYearAmountColumn"`
						HasTotalAmountColumn        bool          `json:"hasTotalAmountColumn"`
						KbkArticleRowList           []interface{} `json:"kbkArticleRowList"`
						KvrCodeRowList              []struct {
							FinancialSource  string `json:"financialSource"`
							FirstYearAmount  string `json:"firstYearAmount"`
							KbkKosguKvr      string `json:"kbkKosguKvr"`
							OtherYearAmount  string `json:"otherYearAmount"`
							SecondYearAmount string `json:"secondYearAmount"`
							ThirdYearAmount  string `json:"thirdYearAmount"`
							TotalYearAmount  string `json:"totalYearAmount"`
						} `json:"kvrCodeRowList"`
						OtherYearAmountColumnHeader  interface{} `json:"otherYearAmountColumnHeader"`
						SecondYearAmountColumnHeader string      `json:"secondYearAmountColumnHeader"`
						ThirdYearAmountColumnHeader  string      `json:"thirdYearAmountColumnHeader"`
						TotalAmountColumnHeader      string      `json:"totalAmountColumnHeader"`
						TotalBudgetRow               interface{} `json:"totalBudgetRow"`
						TotalCommonRow               struct {
							FinancialSource  string `json:"financialSource"`
							FirstYearAmount  string `json:"firstYearAmount"`
							KbkKosguKvr      string `json:"kbkKosguKvr"`
							OtherYearAmount  string `json:"otherYearAmount"`
							SecondYearAmount string `json:"secondYearAmount"`
							ThirdYearAmount  string `json:"thirdYearAmount"`
							TotalYearAmount  string `json:"totalYearAmount"`
						} `json:"totalCommonRow"`
						KvrCodeSumTitle       string   `json:"kvrCodeSumTitle"`
						BudgetRowListSum      []string `json:"budgetRowListSum"`
						KbkArticleRowListSum  []string `json:"kbkArticleRowListSum"`
						KbkArticleSumTitle    string   `json:"kbkArticleSumTitle"`
						ExtrabudgetRowListSum []string `json:"extrabudgetRowListSum"`
						ExtraBudgetSumTitle   string   `json:"extraBudgetSumTitle"`
						KvrCodeRowListSum     []string `json:"kvrCodeRowListSum"`
						BudgetSumTitle        string   `json:"budgetSumTitle"`
					} `json:"annualPaymentPlanTable"`
					ExtraBudgetPayment                interface{}   `json:"extraBudgetPayment"`
					KbkAfter2016List                  []interface{} `json:"kbkAfter2016List"`
					KbkAfter2016TotalAmountByYearMaps struct {
					} `json:"kbkAfter2016TotalAmountByYearMaps"`
					KbkAfter2016TotalAmountByYearMapsFormated struct {
					} `json:"kbkAfter2016TotalAmountByYearMapsFormated"`
					KbkArticleList                  []interface{} `json:"kbkArticleList"`
					KbkArticleTotalAmountByYearMaps struct {
					} `json:"kbkArticleTotalAmountByYearMaps"`
					KbkArticleTotalAmountByYearMapsFormated struct {
					} `json:"kbkArticleTotalAmountByYearMapsFormated"`
					KbkList                  []interface{} `json:"kbkList"`
					KbkTotalAmountByYearMaps struct {
					} `json:"kbkTotalAmountByYearMaps"`
					KbkTotalAmountByYearMapsFormated struct {
					} `json:"kbkTotalAmountByYearMapsFormated"`
					KosguList                          interface{} `json:"kosguList"`
					KosguTotalAmountByYearMaps         interface{} `json:"kosguTotalAmountByYearMaps"`
					KosguTotalAmountByYearMapsFormated interface{} `json:"kosguTotalAmountByYearMapsFormated"`
					KvrCodeList                        []struct {
						AmountByYearMaps struct {
							Num2020 []int     `json:"2020"`
							Num2021 []float64 `json:"2021"`
							Num2022 []int     `json:"2022"`
							Num2023 []int     `json:"2023"`
						} `json:"amountByYearMaps"`
						AmountByYearMapsFormated struct {
							Num2020 []string `json:"2020"`
							Num2021 []string `json:"2021"`
							Num2022 []string `json:"2022"`
							Num2023 []string `json:"2023"`
						} `json:"amountByYearMapsFormated"`
						ContractPaymentPlanID interface{} `json:"contractPaymentPlanId"`
						Number                string      `json:"number"`
					} `json:"kvrCodeList"`
					KvrCodeTotalAmountByYearMaps struct {
						Num2020 int     `json:"2020"`
						Num2021 float64 `json:"2021"`
						Num2022 int     `json:"2022"`
						Num2023 int     `json:"2023"`
					} `json:"kvrCodeTotalAmountByYearMaps"`
					KvrCodeTotalAmountByYearMapsFormated struct {
						Num2020 string `json:"2020"`
						Num2021 string `json:"2021"`
						Num2022 string `json:"2022"`
						Num2023 string `json:"2023"`
					} `json:"kvrCodeTotalAmountByYearMapsFormated"`
					KvrList                  []interface{} `json:"kvrList"`
					KvrTotalAmountByYearMaps struct {
					} `json:"kvrTotalAmountByYearMaps"`
					KvrTotalAmountByYearMapsFormated struct {
					} `json:"kvrTotalAmountByYearMapsFormated"`
					Total                     int   `json:"total"`
					Years                     []int `json:"years"`
					HasAnnualPaymentPlanTable bool  `json:"hasAnnualPaymentPlanTable"`
				} `json:"contractPaymentPlan"`
				ContractPriceFormula                interface{} `json:"contractPriceFormula"`
				Currency                            string      `json:"currency"`
				CurrencyCbRate                      interface{} `json:"currencyCbRate"`
				CurrencyRate                        interface{} `json:"currencyRate"`
				CurrencyRateDate                    interface{} `json:"currencyRateDate"`
				DiscussionNeed                      bool        `json:"discussionNeed"`
				FinanceSource                       interface{} `json:"financeSource"`
				Ikz                                 []string    `json:"ikz"`
				ImposDefWorkSize                    bool        `json:"imposDefWorkSize"`
				InitialContractPrice                float64     `json:"initialContractPrice"`
				InterBudgetaryTransfer              interface{} `json:"interBudgetaryTransfer"`
				NmckContractCurrency                interface{} `json:"nmckContractCurrency"`
				NmckContractCurrencyForeignCurrency bool        `json:"nmckContractCurrencyForeignCurrency"`
				NominalRate                         interface{} `json:"nominalRate"`
				NotNeedObozDescription              interface{} `json:"notNeedObozDescription"`
				Oboz                                interface{} `json:"oboz"`
				ObozStatus                          interface{} `json:"obozStatus"`
				RestrictsVendorSelect               interface{} `json:"restrictsVendorSelect"`
				SpecifyContractPriceFormula         bool        `json:"specifyContractPriceFormula"`
				CurrencyNotEquals                   bool        `json:"currencyNotEquals"`
				HasContractPaymentPlan              bool        `json:"hasContractPaymentPlan"`
			} `json:"initialContractPrice"`
			OrganizationID      int `json:"organizationId"`
			RelationshipWithRPG struct {
				Link  string `json:"link"`
				Show  bool   `json:"show"`
				Title string `json:"title"`
			} `json:"relationshipWithRPG"`
			ShowContractPaymentPlanInfo bool `json:"showContractPaymentPlanInfo"`
			SingleCustomer              bool `json:"singleCustomer"`
			WarrantyObligations         struct {
				WarrantyObligations            bool        `json:"warrantyObligations"`
				WarrantyObligationsAccount     interface{} `json:"warrantyObligationsAccount"`
				WarrantyObligationsBik         interface{} `json:"warrantyObligationsBik"`
				WarrantyObligationsMethod      interface{} `json:"warrantyObligationsMethod"`
				WarrantyObligationsPersAccount interface{} `json:"warrantyObligationsPersAccount"`
				WarrantyObligationsRequisites  string      `json:"warrantyObligationsRequisites"`
				WarrantyObligationsSize        interface{} `json:"warrantyObligationsSize"`
			} `json:"warrantyObligations"`
			ImposDefWorkSize bool `json:"imposDefWorkSize"`
		} `json:"customerRequirementsBlock"`
		CustomerRequirementsPagingDto struct {
			CriteriaPageNumber int    `json:"criteriaPageNumber"`
			PageCount          int    `json:"pageCount"`
			RecordsPerPage     string `json:"recordsPerPage"`
			Rounded            bool   `json:"rounded"`
			Total              int    `json:"total"`
			TotalRounded       int    `json:"totalRounded"`
			RssLink            string `json:"rssLink"`
			MoreThanTen        bool   `json:"moreThanTen"`
			ExceedCount        bool   `json:"exceedCount"`
		} `json:"customerRequirementsPagingDto"`
		GeneralInformationOnPurchaseBlock struct {
			AdditionalInformationOnChange         interface{} `json:"additionalInformationOnChange"`
			Article15PartType                     interface{} `json:"article15PartType"`
			BudgetUnionState                      bool        `json:"budgetUnionState"`
			InterBudgetaryTransfer                bool        `json:"interBudgetaryTransfer"`
			LifeCycleCases                        interface{} `json:"lifeCycleCases"`
			MedicinePurchase                      interface{} `json:"medicinePurchase"`
			NumberStandardContract                interface{} `json:"numberStandardContract"`
			OrderFormat                           string      `json:"orderFormat"`
			OrderName                             string      `json:"orderName"`
			OrganizationRole                      string      `json:"organizationRole"`
			PlacementGeneralInformationOfPurchase struct {
				Link                string `json:"link"`
				OrderShared         bool   `json:"orderShared"`
				OrganizationID      int    `json:"organizationId"`
				PowerOfOrganization string `json:"powerOfOrganization"`
				Title               string `json:"title"`
				Show                bool   `json:"show"`
			} `json:"placementGeneralInformationOfPurchase"`
			PublicDiscussion    interface{} `json:"publicDiscussion"`
			PurchasesAbroad     bool        `json:"purchasesAbroad"`
			RelationshipWithRPG []struct {
				Link  string `json:"link"`
				Show  bool   `json:"show"`
				Title string `json:"title"`
			} `json:"relationshipWithRPG"`
			RequisitesOfLegalAct           interface{} `json:"requisitesOfLegalAct"`
			StageOfPurchase                string      `json:"stageOfPurchase"`
			WayDefinitionOfSupplier        string      `json:"wayDefinitionOfSupplier"`
			NameOfSubjectPurchase          interface{} `json:"nameOfSubjectPurchase"`
			NameOfSubjectPurchaseForOther  string      `json:"nameOfSubjectPurchaseForOther"`
			NameOfElectronicPlatform       string      `json:"nameOfElectronicPlatform"`
			WebAddressOfElectronicPlatform struct {
				Link  string `json:"link"`
				Show  bool   `json:"show"`
				Title string `json:"title"`
			} `json:"webAddressOfElectronicPlatform"`
			Goz                   bool `json:"goz"`
			WillContractLifeCycle bool `json:"willContractLifeCycle"`
			Part2Article84        bool `json:"part2Article84"`
			OrganizationRoleTKO   bool `json:"organizationRoleTKO"`
		} `json:"generalInformationOnPurchaseBlock"`
		InitialContractPriceBlock struct {
			AdvancePercent                      interface{} `json:"advancePercent"`
			ContractCurrency                    interface{} `json:"contractCurrency"`
			ContractPaymentPlan                 interface{} `json:"contractPaymentPlan"`
			ContractPriceFormula                interface{} `json:"contractPriceFormula"`
			Currency                            string      `json:"currency"`
			CurrencyCbRate                      interface{} `json:"currencyCbRate"`
			CurrencyRate                        interface{} `json:"currencyRate"`
			CurrencyRateDate                    interface{} `json:"currencyRateDate"`
			DiscussionNeed                      bool        `json:"discussionNeed"`
			FinanceSource                       interface{} `json:"financeSource"`
			Ikz                                 []string    `json:"ikz"`
			ImposDefWorkSize                    bool        `json:"imposDefWorkSize"`
			InitialContractPrice                float64     `json:"initialContractPrice"`
			InterBudgetaryTransfer              bool        `json:"interBudgetaryTransfer"`
			NmckContractCurrency                interface{} `json:"nmckContractCurrency"`
			NmckContractCurrencyForeignCurrency bool        `json:"nmckContractCurrencyForeignCurrency"`
			NominalRate                         interface{} `json:"nominalRate"`
			NotNeedObozDescription              interface{} `json:"notNeedObozDescription"`
			Oboz                                interface{} `json:"oboz"`
			ObozStatus                          interface{} `json:"obozStatus"`
			RestrictsVendorSelect               interface{} `json:"restrictsVendorSelect"`
			SpecifyContractPriceFormula         bool        `json:"specifyContractPriceFormula"`
			CurrencyNotEquals                   bool        `json:"currencyNotEquals"`
			HasContractPaymentPlan              bool        `json:"hasContractPaymentPlan"`
		} `json:"initialContractPriceBlock"`
		OrderGrantingCompetitiveDocCommonBlock interface{} `json:"orderGrantingCompetitiveDocCommonBlock"`
		OrganizationDefinitionSupplierBlock    struct {
			AdditionalInformation              interface{} `json:"additionalInformation"`
			ContactPhoneNumber                 string      `json:"contactPhoneNumber"`
			EmailAddress                       string      `json:"emailAddress"`
			Fax                                interface{} `json:"fax"`
			InformationAboutContractService    interface{} `json:"informationAboutContractService"`
			Location                           string      `json:"location"`
			OrganizationProvidingAccommodation string      `json:"organizationProvidingAccommodation"`
			OrganizationRoleType               string      `json:"organizationRoleType"`
			PostAddress                        string      `json:"postAddress"`
			ResponsibleOfficer                 string      `json:"responsibleOfficer"`
			SpecOrgAddress                     interface{} `json:"specOrgAddress"`
			SpecOrgLocation                    interface{} `json:"specOrgLocation"`
			SpecOrgName                        interface{} `json:"specOrgName"`
			SpecOrgURL                         interface{} `json:"specOrgUrl"`
		} `json:"organizationDefinitionSupplierBlock"`
		ProcedurePurchaseBlock struct {
			ActualOfferEndDate        int64       `json:"actualOfferEndDate"`
			AdditionalInformation     interface{} `json:"additionalInformation"`
			ConditionsParticAviod     interface{} `json:"conditionsParticAviod"`
			ContractPeriod            interface{} `json:"contractPeriod"`
			EndDateTime               int64       `json:"endDateTime"`
			FinalOfferOpenDate        interface{} `json:"finalOfferOpenDate"`
			FinalOfferOpenInfo        interface{} `json:"finalOfferOpenInfo"`
			FinalOfferOpenPlace       interface{} `json:"finalOfferOpenPlace"`
			OfferDeliverPlace         string      `json:"offerDeliverPlace"`
			OfferDeliverProcedure     string      `json:"offerDeliverProcedure"`
			OfferDiscussDate          interface{} `json:"offerDiscussDate"`
			OfferDiscussPlace         interface{} `json:"offerDiscussPlace"`
			OfferOpenDate             interface{} `json:"offerOpenDate"`
			OfferOpenInfo             interface{} `json:"offerOpenInfo"`
			OfferOpenPlace            interface{} `json:"offerOpenPlace"`
			OfferReviewDate           interface{} `json:"offerReviewDate"`
			OfferReviewFirstPartDate  int64       `json:"offerReviewFirstPartDate"`
			OfferReviewFirstPartFlag  bool        `json:"offerReviewFirstPartFlag"`
			OfferReviewInfo           interface{} `json:"offerReviewInfo"`
			OfferReviewPlace          interface{} `json:"offerReviewPlace"`
			OfferReviewSecondPartDate interface{} `json:"offerReviewSecondPartDate"`
			Preselect                 bool        `json:"preselect"`
			PreselectDate             interface{} `json:"preselectDate"`
			PreselectPlace            interface{} `json:"preselectPlace"`
			ProlongDate               interface{} `json:"prolongDate"`
			QuotBidForm               interface{} `json:"quotBidForm"`
			StartDateTime             int64       `json:"startDateTime"`
			TransferOfferOpenDate     interface{} `json:"transferOfferOpenDate"`
			AuctionEtpDate            int64       `json:"auctionEtpDate"`
			ClosedAuctionDate         interface{} `json:"closedAuctionDate"`
			DateTime                  int64       `json:"dateTime"`
			EndReviewDate             interface{} `json:"endReviewDate"`
			ExistsPublishedUIA        bool        `json:"existsPublishedUIA"`
			NewActualOfferEndDate     interface{} `json:"newActualOfferEndDate"`
			NewDateTime               interface{} `json:"newDateTime"`
			NewOfferEndDate           interface{} `json:"newOfferEndDate"`
			NewOfferReviewFirstPDate  interface{} `json:"newOfferReviewFirstPDate"`
			NewOfferReviewSecondPDate interface{} `json:"newOfferReviewSecondPDate"`
			NewTime                   interface{} `json:"newTime"`
		} `json:"procedurePurchaseBlock"`
		ProcedurePurchaseBlockSt2 interface{} `json:"procedurePurchaseBlockSt2"`
		PublicDiscussionBlock     interface{} `json:"publicDiscussionBlock"`
		PurchaseObjectInfoBlock   struct {
			CurrencyName                        interface{} `json:"currencyName"`
			CurrencyRate                        interface{} `json:"currencyRate"`
			CurrencySymbol                      string      `json:"currencySymbol"`
			LotID                               int         `json:"lotId"`
			MedicinePositionsInfo               interface{} `json:"medicinePositionsInfo"`
			NeedPackInfo                        bool        `json:"needPackInfo"`
			NmckContractCurrencyForeignCurrency bool        `json:"nmckContractCurrencyForeignCurrency"`
			Paging                              struct {
				CriteriaPageNumber int    `json:"criteriaPageNumber"`
				PageCount          int    `json:"pageCount"`
				RecordsPerPage     string `json:"recordsPerPage"`
				Rounded            bool   `json:"rounded"`
				Total              int    `json:"total"`
				TotalRounded       int    `json:"totalRounded"`
				RssLink            string `json:"rssLink"`
				MoreThanTen        bool   `json:"moreThanTen"`
				ExceedCount        bool   `json:"exceedCount"`
			} `json:"paging"`
			PlacingWay    interface{} `json:"placingWay"`
			PositionsInfo []struct {
				AdditionalCharacteristicsTRU []struct {
					CharacteristicsTRU []struct {
						ShortUnit string `json:"shortUnit"`
						Unit      string `json:"unit"`
						Value     string `json:"value"`
					} `json:"characteristicsTRU"`
					Name string `json:"name"`
				} `json:"additionalCharacteristicsTRU"`
				Cost                   interface{} `json:"cost"`
				Count                  int         `json:"count"`
				CustomerCountTru       interface{} `json:"customerCountTru"`
				JustificationInclusion interface{} `json:"justificationInclusion"`
				PositionCode           struct {
					Code string `json:"code"`
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"positionCode"`
				PositionID            int         `json:"positionId"`
				PositionName          string      `json:"positionName"`
				PriceForUnit          interface{} `json:"priceForUnit"`
				ReasonCharacteristics string      `json:"reasonCharacteristics"`
				Unit                  string      `json:"unit"`
				MedicalProduct        bool        `json:"medicalProduct"`
				Ktru                  bool        `json:"ktru"`
			} `json:"positionsInfo"`
			PurchaseDescription          interface{} `json:"purchaseDescription"`
			RestrictsForeignServices     interface{} `json:"restrictsForeignServices"`
			ShowCost                     bool        `json:"showCost"`
			ShowCountOrganization        bool        `json:"showCountOrganization"`
			ShowPrice                    bool        `json:"showPrice"`
			ShowQuantity                 bool        `json:"showQuantity"`
			ShowRestrictsForeignServices bool        `json:"showRestrictsForeignServices"`
			ShowUnit                     bool        `json:"showUnit"`
			TotalCost                    interface{} `json:"totalCost"`
			TotalCostByCurrency          interface{} `json:"totalCostByCurrency"`
			ShowMedicalProduct           bool        `json:"showMedicalProduct"`
			ImposDefWorkSize             bool        `json:"imposDefWorkSize"`
			Medicine                     bool        `json:"medicine"`
		} `json:"purchaseObjectInfoBlock"`
		SingleCustomer                  bool `json:"singleCustomer"`
		StructuredDocumentationBlockDto struct {
			ClarProcedureDelivery         string `json:"clarProcedureDelivery"`
			ClarProcedureEndDate          int64  `json:"clarProcedureEndDate"`
			ClarProcedureStartDate        int64  `json:"clarProcedureStartDate"`
			ContractNecessaryGoodsCh9St37 bool   `json:"contractNecessaryGoodsCh9St37"`
			ModifiableContactCountGoods   bool   `json:"modifiableContactCountGoods"`
			OnesideRejectionCh9St37       bool   `json:"onesideRejectionCh9St37"`
		} `json:"structuredDocumentationBlockDto"`
		TenderDocumentationBlockDto interface{} `json:"tenderDocumentationBlockDto"`
	} `json:"dto"`
}
