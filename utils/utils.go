package utils

import (
	"encoding/xml"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os"
)

type FileJson struct {
	Status           string         `json:"status" db:"status"`
	Uid              string         `json:"Uid" db:"entity_id"`
	PurchaseNumber   string         `json:"PurchaseNumber" db:"purchase_number"`
	Description      string         `json:"Description" db:"description"`
	DateStart        string         `json:"DateStart" db:"date_start"`
	DateEnd          string         `json:"DateEnd" db:"date_end"`
	DateStartAuction string         `json:"DateStartAuction" db:"date_start_auction"`
	DateEndAuction   string         `json:"DateEndAuction" db:"date_end_auction"`
	OriginalUrl      string         `json:"OriginalUrl" db:"original_url"`
	NameEtp          string         `json:"NameEtp" db:"name_etp"`
	UrlNameEtp       string         `json:"UrlNameEtp" db:"url_name_etp"`
	Fz               string         `json:"Fz" db:"fz"`
	Type             string         `json:"Type" db:"type"`
	PurchaseStatus   string         `json:"PurchaseStatus" db:"purchase_status"`
	Customer         customer       `json:"Customer"`
	Delivery         deliveryPlace  `json:"Delivery"`
	DeliveryTime     string         `json:"DeliveryTime" db:"delivery_time"`
	Price            float64        `json:"Price" db:"price"`
	Region           string         `json:"Region" db:"region"`
	Currency         string         `json:"Currency" db:"currency"`
	AmountCollateral string         `json:"AmountCollateral" db:"amount_collateral"`
	Documents        []Documents    `json:"Documents"`
	LotItems         []LotItems     `json:"LotItems"`
	DrugLotItems     []DrugLotItems `json:"DrugLotItems"`
	//db only fields
	KladrCode     string `db:"delivery_kladr_code"`
	FullName      string `db:"delivery_full_name"`
	Kladr         string `db:"delivery_kladr"`
	DeliveryPlace string `db:"delivery_place"`

	CustomerPhone        string `db:"customer_phone"`
	CustomerEmail        string `db:"customer_email"`
	CustomerLocation     string `db:"customer_location"`
	CustomerOrgFullName  string `db:"customer_org_full_name"`
	CustomerOrgShortName string `db:"customer_org_short_name"`
	CustomerRegNum       string `db:"customer_reg_num"`
	Lastname             string `db:"customer_contact_last_name"`
	FirstName            string `db:"customer_contact_first_name"`
	MiddleName           string `db:"customer_contact_middle_name"`
}
type deliveryPlace struct {
	KladrCode     string `json:"KladrCode"`
	FullName      string `json:"FullName"`
	Kladr         string `json:"Kladr"`
	DeliveryPlace string `json:"DeliveryPlace"`
}
type customer struct {
	CustomerPhone        string `json:"CustomerPhone" db:"customer_phone"`
	CustomerEmail        string `json:"CustomerEmail" db:"customer_email"`
	CustomerLocation     string `json:"CustomerLocation"`
	CustomerOrgFullName  string `json:"CustomerOrgFullName"`
	CustomerOrgShortName string `json:"CustomerOrgShortName"`
	CustomerRegNum       string `json:"CustomerRegNum"`
	CustomerContact      customerContact
}
type customerContact struct {
	Lastname   string `json:"Lastname"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
}
type Documents struct {
	Did          string `db:"document_id"`
	DocumentName string `json:"DocumentName" db:"document_name"`
	DocumentUrl  string `json:"DocumentUrl" db:"document_url"`
}
type LotItems struct {
	Sid             string  `json:"Sid" db:"sid"`
	Name            string  `json:"Name" db:"name"`
	UnitPrice       float64 `json:"UnitPrice" db:"unit_price"`
	Sum             string  `json:"Sum" db:"sum"`
	Quantity        string  `json:"Quantity" db:"quantity"`
	Okpd2Code       string  `json:"Okpd2Code" db:"okpd_2_code"`
	OkeiTitle       string  `json:"OkeiTitle" db:"okei_title"`
	CodeMes         string  `json:"CodeMes" db:"code_mes"`
	NationalCodeMes string  `json:"NationalCodeMes" db:"national_code_mes"`
	NameMes         string  `json:"NameMes" db:"name_mes"`
}
type DrugLotItems struct {
	Sid             string `json:"Sid" db:"sid"`
	MNNinfo         string `json:"MNNinfo" db:"MNNinfo"`
	MedFormName     string `json:"MedFormName" db:"med_form_name"`
	DosageGRLSValue string `json:"DosageGRLSValue" db:"dosage_grlsv_values"`
	DosageCode      string `json:"DosageCode" db:"dosage_code"`
	DosageOKEIName  string `json:"DosageOKEIName" db:"dosage_okei_name"`
	DosageUserName  string `json:"DosageUserName" db:"dosage_user_name"`
	DrugQuantity    string `json:"DrugQuantity" db:"drug_quantity"`
}
type DocEntity struct {
	DocumentId string `db:"document_id"`
	EntityId   string `db:"entity_id"`
}
type LotEntity struct {
	LotId    string `db:"lot_id"`
	EntityId string `db:"entity_id"`
}
type DrugEntity struct {
	DrugId   string `db:"drug_id"`
	EntityId string `db:"entity_id"`
}

type FileXml struct {
	XMLName              xml.Name `xml:"export" json:"XmlName"`
	Text                 string   `xml:",chardata"`
	Xmlns                string   `xml:"xmlns,attr"`
	Ns2                  string   `xml:"ns2,attr"`
	Ns4                  string   `xml:"ns4,attr"`
	Ns3                  string   `xml:"ns3,attr"`
	Ns6                  string   `xml:"ns6,attr"`
	Ns5                  string   `xml:"ns5,attr"`
	Ns8                  string   `xml:"ns8,attr"`
	Ns7                  string   `xml:"ns7,attr"`
	Ns13                 string   `xml:"ns13,attr"`
	Ns9                  string   `xml:"ns9,attr"`
	Ns12                 string   `xml:"ns12,attr"`
	Ns11                 string   `xml:"ns11,attr"`
	Ns10                 string   `xml:"ns10,attr"`
	EpNotificationEF2020 struct {
		Text          string `xml:",chardata"`
		SchemeVersion string `xml:"schemeVersion,attr"`
		ID            string `xml:"id" json:"purchase_id"`
		VersionNumber string `xml:"versionNumber"`
		CommonInfo    struct {
			Text               string `xml:",chardata"`
			PurchaseNumber     string `xml:"purchaseNumber"`
			DocNumber          string `xml:"docNumber"`
			PlannedPublishDate string `xml:"plannedPublishDate"`
			PublishDTInEIS     string `xml:"publishDTInEIS"`
			Href               string `xml:"href"`
			NotPublishedOnEIS  string `xml:"notPublishedOnEIS"`
			PlacingWay         struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"placingWay"`
			ETP struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
				URL  string `xml:"url"`
			} `xml:"ETP"`
			ContractConclusionOnSt83Ch2 string `xml:"contractConclusionOnSt83Ch2"`
			PurchaseObjectInfo          string `xml:"purchaseObjectInfo"`
		} `xml:"commonInfo"`
		PurchaseResponsibleInfo struct {
			Text               string `xml:",chardata"`
			ResponsibleOrgInfo struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				ShortName       string `xml:"shortName"`
				PostAddress     string `xml:"postAddress"`
				FactAddress     string `xml:"factAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
			} `xml:"responsibleOrgInfo"`
			ResponsibleRole string `xml:"responsibleRole"`
			ResponsibleInfo struct {
				Text              string `xml:",chardata"`
				OrgPostAddress    string `xml:"orgPostAddress"`
				OrgFactAddress    string `xml:"orgFactAddress"`
				ContactPersonInfo struct {
					Text       string `xml:",chardata"`
					LastName   string `xml:"lastName"`
					FirstName  string `xml:"firstName"`
					MiddleName string `xml:"middleName"`
				} `xml:"contactPersonInfo"`
				ContactEMail string `xml:"contactEMail"`
				ContactPhone string `xml:"contactPhone"`
			} `xml:"responsibleInfo"`
		} `xml:"purchaseResponsibleInfo"`
		PrintFormInfo struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url"`
		} `xml:"printFormInfo"`
		AttachmentsInfo struct {
			Text           string `xml:",chardata"`
			AttachmentInfo []struct {
				Text               string `xml:",chardata"`
				PublishedContentId string `xml:"publishedContentId"`
				FileName           string `xml:"fileName"`
				FileSize           string `xml:"fileSize"`
				DocDescription     string `xml:"docDescription"`
				DocDate            string `xml:"docDate"`
				URL                string `xml:"url"`
				DocKindInfo        struct {
					Text string `xml:",chardata"`
					Code string `xml:"code"`
					Name string `xml:"name"`
				} `xml:"docKindInfo"`
				CryptoSigns struct {
					Text      string `xml:",chardata"`
					Signature struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"signature"`
				} `xml:"cryptoSigns"`
			} `xml:"attachmentInfo"`
		} `xml:"attachmentsInfo"`
		NotificationInfo struct {
			Text          string `xml:",chardata"`
			ProcedureInfo struct {
				Text           string `xml:",chardata"`
				CollectingInfo struct {
					Text    string `xml:",chardata"`
					StartDT string `xml:"startDT"`
					EndDT   string `xml:"endDT"`
				} `xml:"collectingInfo"`
				BiddingDate     string `xml:"biddingDate"`
				SummarizingDate string `xml:"summarizingDate"`
			} `xml:"procedureInfo"`
			ContractConditionsInfo struct {
				Text         string `xml:",chardata"`
				MaxPriceInfo struct {
					Text     string  `xml:",chardata"`
					MaxPrice float64 `xml:"maxPrice"`
					Currency struct {
						Text string `xml:",chardata"`
						Code string `xml:"code"`
						Name string `xml:"name"`
					} `xml:"currency"`
				} `xml:"maxPriceInfo"`
			} `xml:"contractConditionsInfo"`
			CustomerRequirementsInfo struct {
				Text                    string `xml:",chardata"`
				CustomerRequirementInfo struct {
					Text     string `xml:",chardata"`
					Customer struct {
						Text            string `xml:",chardata"`
						RegNum          string `xml:"regNum"`
						ConsRegistryNum string `xml:"consRegistryNum"`
						FullName        string `xml:"fullName"`
					} `xml:"customer"`
					ContractGuarantee struct {
						Text    string `xml:",chardata"`
						Account struct {
							Text              string `xml:",chardata"`
							Bik               string `xml:"bik"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
							CreditOrgName     string `xml:"creditOrgName"`
						} `xml:"account"`
						ProcedureInfo string `xml:"procedureInfo"`
						Part          string `xml:"part"`
					} `xml:"contractGuarantee"`
					ContractConditionsInfo struct {
						Text         string `xml:",chardata"`
						MaxPriceInfo struct {
							Text     string `xml:",chardata"`
							MaxPrice string `xml:"maxPrice"`
						} `xml:"maxPriceInfo"`
						MustPublicDiscussion string `xml:"mustPublicDiscussion"`
						IKZInfo              struct {
							Text         string `xml:",chardata"`
							PurchaseCode string `xml:"purchaseCode"`
							PublishYear  string `xml:"publishYear"`
							OKPD2Info    struct {
								Text      string `xml:",chardata"`
								Undefined string `xml:"undefined"`
							} `xml:"OKPD2Info"`
							KVRInfo struct {
								Text string `xml:",chardata"`
								KVR  struct {
									Text string `xml:",chardata"`
									Code string `xml:"code"`
									Name string `xml:"name"`
								} `xml:"KVR"`
							} `xml:"KVRInfo"`
							CustomerCode        string `xml:"customerCode"`
							PurchaseNumber      string `xml:"purchaseNumber"`
							PurchaseOrderNumber string `xml:"purchaseOrderNumber"`
						} `xml:"IKZInfo"`
						TenderPlan2020Info struct {
							Text               string `xml:",chardata"`
							Plan2020Number     string `xml:"plan2020Number"`
							Position2020Number string `xml:"position2020Number"`
						} `xml:"tenderPlan2020Info"`
						ContractExecutionPaymentPlan struct {
							Text                 string `xml:",chardata"`
							FinancingSourcesInfo struct {
								Text            string `xml:",chardata"`
								FinancingSource string `xml:"financingSource"`
								CurrentYear     string `xml:"currentYear"`
								FinanceInfo     struct {
									Text        string `xml:",chardata"`
									Total       string `xml:"total"`
									CurrentYear string `xml:"currentYear"`
									FirstYear   string `xml:"firstYear"`
									SecondYear  string `xml:"secondYear"`
									SubsecYears string `xml:"subsecYears"`
								} `xml:"financeInfo"`
							} `xml:"financingSourcesInfo"`
							KVRsInfo struct {
								Text        string `xml:",chardata"`
								CurrentYear string `xml:"currentYear"`
								KVRInfo     struct {
									Text string `xml:",chardata"`
									KVR  struct {
										Text string `xml:",chardata"`
										Code string `xml:"code"`
										Name string `xml:"name"`
									} `xml:"KVR"`
									KVRYearsInfo struct {
										Text        string `xml:",chardata"`
										Total       string `xml:"total"`
										CurrentYear string `xml:"currentYear"`
										FirstYear   string `xml:"firstYear"`
										SecondYear  string `xml:"secondYear"`
										SubsecYears string `xml:"subsecYears"`
									} `xml:"KVRYearsInfo"`
								} `xml:"KVRInfo"`
								TotalSum string `xml:"totalSum"`
							} `xml:"KVRsInfo"`
						} `xml:"contractExecutionPaymentPlan"`
						DeliveryPlacesInfo struct {
							Text              string `xml:",chardata"`
							DeliveryPlaceInfo struct {
								Text  string `xml:",chardata"`
								Kladr struct {
									Text      string `xml:",chardata"`
									KladrCode string `xml:"kladrCode"`
									FullName  string `xml:"fullName"`
								} `xml:"kladr"`
								DeliveryPlace string `xml:"deliveryPlace"`
							} `xml:"deliveryPlaceInfo"`
						} `xml:"deliveryPlacesInfo"`
						DeliveryTerm           string `xml:"deliveryTerm"`
						IsOneSideRejectionSt95 string `xml:"isOneSideRejectionSt95"`
					} `xml:"contractConditionsInfo"`
				} `xml:"customerRequirementInfo"`
			} `xml:"customerRequirementsInfo"`
			PurchaseObjectsInfo struct {
				Text                       string `xml:",chardata"`
				NotDrugPurchaseObjectsInfo struct {
					Text           string `xml:",chardata"`
					PurchaseObject []struct {
						Text string `xml:",chardata"`
						Sid  string `xml:"sid"`
						KTRU struct {
							Text            string `xml:",chardata"`
							Code            string `xml:"code"`
							Name            string `xml:"name"`
							VersionId       string `xml:"versionId"`
							VersionNumber   string `xml:"versionNumber"`
							Characteristics struct {
								Text                              string `xml:",chardata"`
								CharacteristicsUsingReferenceInfo []struct {
									Text   string `xml:",chardata"`
									Code   string `xml:"code"`
									Name   string `xml:"name"`
									Type   string `xml:"type"`
									Kind   string `xml:"kind"`
									Values struct {
										Text  string `xml:",chardata"`
										Value struct {
											Text               string `xml:",chardata"`
											QualityDescription string `xml:"qualityDescription"`
											OKEI               struct {
												Text         string `xml:",chardata"`
												Code         string `xml:"code"`
												NationalCode string `xml:"nationalCode"`
												Name         string `xml:"name"`
											} `xml:"OKEI"`
											ValueFormat string `xml:"valueFormat"`
											RangeSet    struct {
												Text       string `xml:",chardata"`
												ValueRange struct {
													Text            string `xml:",chardata"`
													MinMathNotation string `xml:"minMathNotation"`
													Min             string `xml:"min"`
												} `xml:"valueRange"`
											} `xml:"rangeSet"`
										} `xml:"value"`
									} `xml:"values"`
								} `xml:"characteristicsUsingReferenceInfo"`
							} `xml:"characteristics"`
						} `xml:"KTRU"`
						Name string `xml:"name"`
						OKEI struct {
							Text         string `xml:",chardata"`
							Code         string `xml:"code"`
							NationalCode string `xml:"nationalCode"`
							Name         string `xml:"name"`
						} `xml:"OKEI"`
						Price    float64 `xml:"price"`
						Quantity struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"quantity"`
						Sum              string `xml:"sum"`
						IsMedicalProduct string `xml:"isMedicalProduct"`
						Type             string `xml:"type"`
						HierarchyType    string `xml:"hierarchyType"`
						OKPD2            struct {
							Text     string `xml:",chardata"`
							OKPDCode string `xml:"OKPDCode"`
							OKPDName string `xml:"OKPDName"`
						} `xml:"OKPD2"`
					} `xml:"purchaseObject"`
					TotalSum          string `xml:"totalSum"`
					QuantityUndefined string `xml:"quantityUndefined"`
				} `xml:"notDrugPurchaseObjectsInfo"`
				DrugPurchaseObjectsInfo struct {
					Text                   string `xml:",chardata"`
					DrugPurchaseObjectInfo []struct {
						Text                         string `xml:",chardata"`
						Sid                          string `xml:"sid"`
						ObjectInfoUsingReferenceInfo struct {
							Text      string `xml:",chardata"`
							DrugsInfo struct {
								Text     string `xml:",chardata"`
								DrugInfo struct {
									Text    string `xml:",chardata"`
									MNNInfo struct {
										Text            string `xml:",chardata"`
										MNNExternalCode string `xml:"MNNExternalCode"`
										MNNName         string `xml:"MNNName"`
									} `xml:"MNNInfo"`
									MedicamentalFormInfo struct {
										Text                 string `xml:",chardata"`
										MedicamentalFormName string `xml:"medicamentalFormName"`
									} `xml:"medicamentalFormInfo"`
									DosageInfo struct {
										Text            string `xml:",chardata"`
										DosageGRLSValue string `xml:"dosageGRLSValue"`
										DosageUserOKEI  struct {
											Text         string `xml:",chardata"`
											Code         string `xml:"code"`
											NationalCode string `xml:"nationalCode"`
											Name         string `xml:"name"`
										} `xml:"dosageUserOKEI"`
										DosageUserName string `xml:"dosageUserName"`
									} `xml:"dosageInfo"`
									BasicUnit            string `xml:"basicUnit"`
									DrugQuantity         string `xml:"drugQuantity"`
									LimPriceValuePerUnit string `xml:"limPriceValuePerUnit"`
								} `xml:"drugInfo"`
							} `xml:"drugsInfo"`
						} `xml:"objectInfoUsingReferenceInfo"`
						IsZNVLP                   string `xml:"isZNVLP"`
						DrugQuantityCustomersInfo struct {
							Text                     string `xml:",chardata"`
							DrugQuantityCustomerInfo struct {
								Text     string `xml:",chardata"`
								Customer struct {
									Text            string `xml:",chardata"`
									RegNum          string `xml:"regNum"`
									ConsRegistryNum string `xml:"consRegistryNum"`
									FullName        string `xml:"fullName"`
								} `xml:"customer"`
								Quantity string `xml:"quantity"`
							} `xml:"drugQuantityCustomerInfo"`
							Total string `xml:"total"`
						} `xml:"drugQuantityCustomersInfo"`
						PricePerUnit  string `xml:"pricePerUnit"`
						PositionPrice string `xml:"positionPrice"`
					} `xml:"drugPurchaseObjectInfo"`
					Total             string `xml:"total"`
					QuantityUndefined string `xml:"quantityUndefined"`
				} `xml:"drugPurchaseObjectsInfo"`
			} `xml:"purchaseObjectsInfo"`
			PreferensesInfo struct {
				Text           string `xml:",chardata"`
				PreferenseInfo []struct {
					Text                      string `xml:",chardata"`
					PreferenseRequirementInfo struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
					} `xml:"preferenseRequirementInfo"`
					PrefValue string `xml:"prefValue"`
				} `xml:"preferenseInfo"`
			} `xml:"preferensesInfo"`
			RequirementsInfo struct {
				Text            string `xml:",chardata"`
				RequirementInfo struct {
					Text                      string `xml:",chardata"`
					PreferenseRequirementInfo struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
					} `xml:"preferenseRequirementInfo"`
				} `xml:"requirementInfo"`
			} `xml:"requirementsInfo"`
			RestrictionsInfo struct {
				Text            string `xml:",chardata"`
				RestrictionInfo struct {
					Text                      string `xml:",chardata"`
					PreferenseRequirementInfo struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
					} `xml:"preferenseRequirementInfo"`
					RestrictionsSt14 struct {
						Text            string `xml:",chardata"`
						RestrictionSt14 struct {
							Text             string `xml:",chardata"`
							RequirementsType struct {
								Text            string `xml:",chardata"`
								RequirementType struct {
									Text string `xml:",chardata"`
									Type string `xml:"type"`
								} `xml:"requirementType"`
							} `xml:"requirementsType"`
							NPAInfo struct {
								Text      string `xml:",chardata"`
								Code      string `xml:"code"`
								Name      string `xml:"name"`
								ShortName string `xml:"shortName"`
							} `xml:"NPAInfo"`
						} `xml:"restrictionSt14"`
					} `xml:"restrictionsSt14"`
				} `xml:"restrictionInfo"`
			} `xml:"restrictionsInfo"`
			Flags struct {
				Text                   string `xml:",chardata"`
				PurchaseObjectsCh9St37 string `xml:"purchaseObjectsCh9St37"`
			} `xml:"flags"`
		} `xml:"notificationInfo"`
	} `xml:"epNotificationEF2020"`
}
type FilesParsed struct {
	File        string
	Dir         string
	FullNameDir string
}

func GetDir(root string) ([]FilesParsed, error) {
	var filesParsed []FilesParsed
	var fileParsed FilesParsed
	allDirs, _ := ioutil.ReadDir(root)
	var dirs []string
	for _, dir := range allDirs {
		dirs = append(dirs, dir.Name()[13:len(dir.Name())-27])
		files, _ := ioutil.ReadDir(root + dir.Name())
		for _, file := range files {
			if file.Name()[len(file.Name())-3:] == "xml" {
				fileParsed.File = file.Name()
				fileParsed.Dir = dir.Name()[13 : len(dir.Name())-27]
				fileParsed.FullNameDir = dir.Name()
				filesParsed = append(filesParsed, fileParsed)
			}
		}
	}
	return filesParsed, nil
}

func InsertInto(fn string, dirs FilesParsed) (error, FileJson) {
	commits := map[string]string{
		"Adygeja_Resp":                   "Республика Адыгея",
		"Altaj_Resp":                     "Республика алтай",
		"Altajskij_kraj":                 "Алтайский край",
		"Amurskaja_obl":                  "Амурская область",
		"Arkhangelskaja_obl":             "Архангельская область",
		"Astrakhanskaja_obl":             "Астраханская область",
		"Bajkonur_g":                     "Байконур",
		"Bashkortostan_Resp":             "Республика Башкортостан",
		"Belgorodskaja_obl":              "Белгородская область",
		"Brjanskaja_obl":                 "Брянская область",
		"Burjatija_Resp":                 "Республика Бурятия",
		"Chechenskaja_Resp":              "Чеченская Республика",
		"Cheljabinskaja_obl":             "Челябинская область",
		"Chukotskij_AO":                  "Чукотский АО",
		"Chuvashskaja_Resp":              "Чувашская Республика",
		"Dagestan_Resp":                  "Республика Дагестан",
		"Evrejskaja_Aobl":                "Еврейская АО",
		"Ingushetija_Resp":               "Республика Ингушетия",
		"Irkutskaja_obl":                 "Иркутская область",
		"Ivanovskaja_obl":                "Ивановская область",
		"Jamalo-Neneckij_AO":             "Ямало-Ненецкий АО",
		"Jaroslavskaja_obl":              "Ярославская область",
		"Kabardino-Balkarskaja_Resp":     "Кабардино-Балкарская Республика",
		"Kaliningradskaja_obl":           "Калининградская область",
		"Kalmykija_Resp":                 "Республика Калмыкия",
		"Kaluzhskaja_obl":                "Калужская область",
		"Kamchatskij_kraj":               "Камчатский край",
		"Karachaevo-Cherkesskaja_Resp":   "Карачаево-Черкесская Республика",
		"Karelija_Resp":                  "Республика Карелия",
		"Kemerovskaja_obl":               "Кемеровская область",
		"Khabarovskij_kraj":              "Хабаровский край",
		"Khakasija_Resp":                 "Республика Хакасия",
		"Khanty-Mansijskij_AO-Jugra_AO":  "Ханты-Мансийский АО",
		"Kirovskaja_obl":                 "Кировская_область",
		"Komi_Resp":                      "Республика Коми",
		"Kostromskaja_obl":               "Костромская область",
		"Krasnodarskij_kraj":             "Краснодарский край",
		"Krasnojarskij_kraj":             "Красноярский край",
		"Krim_Resp":                      "Республика Крым",
		"Kurganskaja_obl":                "Курганская_область",
		"Kurskaja_obl":                   "Курская область",
		"Leningradskaja_obl":             "Ленинградская область",
		"Lipeckaja_obl":                  "Липецкая область",
		"Magadanskaja_obl":               "Магаданская область",
		"Marij_El_Resp":                  "Республика Марий Эл",
		"Mordovija_Resp":                 "Республика Мордовия",
		"Moskovskaja_obl":                "Московская область",
		"Moskva":                         "Москва",
		"Murmanskaja_obl":                "Мурманская область",
		"Neneckij_AO":                    "Ненецкий АО",
		"Nizhegorodskaja_obl":            "Нижегородская область",
		"Novgorodskaja_obl":              "Новгородская область",
		"Novosibirskaja_obl":             "Новосибирская область",
		"Omskaja_obl":                    "Омская область",
		"Orenburgskaja_obl":              "Оренбургская область",
		"Orlovskaja_obl":                 "Орловская область",
		"Penzenskaja_obl":                "Пензенская область",
		"Permskij_kraj":                  "Пермский край",
		"Primorskij_kraj":                "Приморский край",
		"Pskovskaja_obl":                 "Псковская область",
		"Rjazanskaja_obl":                "Рязанская область",
		"Rostovskaja_obl":                "Ростовская область",
		"Sakha_Jakutija_Resp":            "Республика Саха",
		"Sakhalinskaja_obl":              "Сахалинская область",
		"Samarskaja_obl":                 "Самарская область",
		"Sankt-Peterburg":                "Санкт-Петербург",
		"Saratovskaja_obl":               "Саратовская область",
		"Sevastopol_g":                   "Севастополь",
		"Severnaja_Osetija-Alanija_Resp": "Северная Осетия",
		"Smolenskaja_obl":                "Смоленская область",
		"Stavropolskij_kraj":             "Ставропольский край",
		"Sverdlovskaja_obl":              "Свердловская область",
		"Tambovskaja_obl":                "Тамбовская область",
		"Tatarstan_Resp":                 "Республика Татарстан",
		"Tjumenskaja_obl":                "Тюменская область",
		"Tomskaja_obl":                   "Томская область",
		"Tulskaja_obl":                   "Тульская область",
		"Tverskaja_obl":                  "Тверская область",
		"Tyva_Resp":                      "Республика Тыва",
		"Udmurtskaja_Resp":               "Удмуртская Республика",
		"Uljanovskaja_obl":               "Ульяновская область",
		"Vladimirskaja_obl":              "Владимирская область",
		"Volgogradskaja_obl":             "Волгоградская область",
		"Vologodskaja_obl":               "Вологодская область",
		"Voronezhskaja_obl":              "Воронежская область",
		"Zabajkalskij_kraj":              "Забайкальский край",
	}
	xmlFile, err := os.Open(fn)

	//fmt.Printf("Successfully Opened %s\n", xmlFile.Name())
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var file FileXml

	_ = xml.Unmarshal(byteValue, &file)

	var docs []Documents
	for _, item := range file.EpNotificationEF2020.AttachmentsInfo.AttachmentInfo {
		var doc Documents
		doc.DocumentName = item.DocDescription
		doc.DocumentUrl = item.URL
		doc.Did = hash(item.URL + item.DocDescription)
		//fmt.Println(hash(item.URL+item.DocDescription) + "\n\n")
		docs = append(docs, doc)
	}
	var lots []LotItems
	for _, item := range file.EpNotificationEF2020.NotificationInfo.PurchaseObjectsInfo.NotDrugPurchaseObjectsInfo.PurchaseObject {
		var lot LotItems
		lot.Name = item.Name
		lot.UnitPrice = item.Price
		lot.Okpd2Code = item.KTRU.Code
		lot.Quantity = item.Quantity.Value
		lot.NameMes = item.Name
		lot.CodeMes = item.OKEI.Code
		lot.Sum = item.Sum
		lot.Sid = item.Sid
		lot.NationalCodeMes = item.OKEI.NationalCode
		lots = append(lots, lot)
	}
	var drugs []DrugLotItems
	for _, item := range file.EpNotificationEF2020.NotificationInfo.PurchaseObjectsInfo.DrugPurchaseObjectsInfo.DrugPurchaseObjectInfo {
		var lot DrugLotItems
		lot.Sid = item.Sid
		lot.MNNinfo = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.MNNInfo.MNNName
		lot.MedFormName = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.MedicamentalFormInfo.MedicamentalFormName
		lot.DosageGRLSValue = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.DosageInfo.DosageGRLSValue
		lot.DosageCode = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.DosageInfo.DosageUserOKEI.Code
		lot.DosageOKEIName = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.DosageInfo.DosageUserOKEI.Name
		lot.DosageUserName = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.DosageInfo.DosageUserName
		lot.DrugQuantity = item.ObjectInfoUsingReferenceInfo.DrugsInfo.DrugInfo.DrugQuantity
		drugs = append(drugs, lot)
	}
	fileJson := FileJson{
		Status:           "Подача заявок",
		Uid:              file.EpNotificationEF2020.ID,
		PurchaseNumber:   file.EpNotificationEF2020.CommonInfo.PurchaseNumber,
		Description:      file.EpNotificationEF2020.CommonInfo.PurchaseObjectInfo,
		DateStart:        file.EpNotificationEF2020.CommonInfo.PublishDTInEIS,
		DateEnd:          file.EpNotificationEF2020.NotificationInfo.ProcedureInfo.CollectingInfo.EndDT,
		DateStartAuction: file.EpNotificationEF2020.NotificationInfo.ProcedureInfo.BiddingDate,
		DateEndAuction:   file.EpNotificationEF2020.NotificationInfo.ProcedureInfo.SummarizingDate,
		OriginalUrl:      file.EpNotificationEF2020.CommonInfo.Href,
		NameEtp:          file.EpNotificationEF2020.CommonInfo.ETP.Name,
		UrlNameEtp:       file.EpNotificationEF2020.CommonInfo.ETP.URL,
		Region:           commits[dirs.Dir],
		Fz:               "44-ФЗ",
		Type:             file.EpNotificationEF2020.CommonInfo.PlacingWay.Name,

		CustomerOrgShortName: file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.ShortName,
		CustomerOrgFullName:  file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.FullName,
		CustomerRegNum:       file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.RegNum,
		CustomerPhone:        file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPhone,
		Lastname:             file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.LastName,
		FirstName:            file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.FirstName,
		MiddleName:           file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.MiddleName,
		Customer: customer{
			CustomerOrgShortName: file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.ShortName,
			CustomerOrgFullName:  file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.FullName,
			CustomerRegNum:       file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.RegNum,
			CustomerPhone:        file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPhone,
			CustomerContact: customerContact{
				Lastname:   file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.LastName,
				FirstName:  file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.FirstName,
				MiddleName: file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactPersonInfo.MiddleName,
			},
			CustomerEmail:    file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleInfo.ContactEMail,
			CustomerLocation: file.EpNotificationEF2020.PurchaseResponsibleInfo.ResponsibleOrgInfo.FactAddress,
		},
		Delivery: deliveryPlace{
			Kladr:     file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.Kladr.KladrCode,
			KladrCode: file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.Kladr.FullName,
			FullName:  file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.DeliveryPlace,
		},
		Kladr:            file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.Kladr.KladrCode,
		KladrCode:        file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.Kladr.FullName,
		FullName:         file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryPlacesInfo.DeliveryPlaceInfo.DeliveryPlace,
		DeliveryTime:     file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractConditionsInfo.DeliveryTerm,
		Price:            file.EpNotificationEF2020.NotificationInfo.ContractConditionsInfo.MaxPriceInfo.MaxPrice,
		Currency:         file.EpNotificationEF2020.NotificationInfo.ContractConditionsInfo.MaxPriceInfo.Currency.Code,
		AmountCollateral: file.EpNotificationEF2020.NotificationInfo.CustomerRequirementsInfo.CustomerRequirementInfo.ContractGuarantee.Part,
		Documents:        docs,
		LotItems:         lots,
		DrugLotItems:     drugs,
	}

	if err != nil {
		return err, fileJson
	}
	return err, fileJson
}
func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	str := fmt.Sprint(h.Sum32())
	return str
}
func UniqueLot(s []LotItems) []LotItems {
	inResult := make(map[string]bool)
	var result []LotItems
	for _, str := range s {
		if _, ok := inResult[str.Sid]; !ok {
			inResult[str.Sid] = true
			result = append(result, str)
		}
	}
	return result
}
func UniqueDoc(s []Documents) []Documents {
	inResult := make(map[string]bool)
	var result []Documents
	for _, str := range s {
		if _, ok := inResult[str.Did]; !ok {
			inResult[str.Did] = true
			result = append(result, str)
		}
	}
	return result
}
func UniqueDrug(s []DrugLotItems) []DrugLotItems {
	inResult := make(map[string]bool)
	var result []DrugLotItems
	for _, str := range s {
		if _, ok := inResult[str.Sid]; !ok {
			inResult[str.Sid] = true
			result = append(result, str)
		}
	}
	return result
}
