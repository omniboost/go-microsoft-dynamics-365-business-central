package tripletex

type VoucherType struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	Name    string `json:"name"`
}

type Postings []Posting

type Posting struct {
	ID                  int         `json:"id,omitempty"`
	Version             int         `json:"version,omitempty"`
	URL                 string      `json:"url"`
	Date                string      `json:"date"`
	Description         string      `json:"description"`
	Account             *Account    `json:"account"`
	Customer            *Customer   `json:"customer,omitempty"`
	Supplier            *Supplier   `json:"supplier,omitempty"`
	Employee            *Employee   `json:"employee,omitempty"`
	Project             *Project    `json:"project,omitempty"`
	Product             *Product    `json:"project,omitempty"`
	Department          *Department `json:"department,omitempty"`
	VATType             *VATType    `json:"vatType,omitempty"`
	Amount              int         `json:"amount,omitempty"`
	AmountCurrency      int         `json:"amountCurrency,omitempty"`
	AmountGross         float64     `json:"amountGross"`
	AmountGrossCurrency float64     `json:"amountGrossCurrency"`
	Currency            *Currency   `json:"currency,omitempty"`
	CloseGroup          *CloseGroup `json:"closeGroup,omitempty"`
	InvoiceNumber       string      `json:"invoiceNumber,omitempty"`
	TermOfPayment       string      `json:"termOfPayment,omitempty"`
	Row                 int         `json:"row,omitempty"`
}

type Accounts []Account

type Account struct {
	ID int `json:"id,omitempty"`
	// Version                        int       `json:"version,omitempty"`
	URL    string `json:"url"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	// Description                    string    `json:"description"`
	// VatType                        VATType   `json:"vatType"`
	// VatLocked                      bool      `json:"vatLocked"`
	// Currency                       *Currency `json:"currency"`
	// IsCloseable                    bool      `json:"isCloseable"`
	// IsApplicableForSupplierInvoice bool      `json:"isApplicableForSupplierInvoice"`
	// RequireReconciliation          bool      `json:"requireReconciliation"`
	// IsInactive                     bool      `json:"isInactive"`
	// IsBankAccount                  bool      `json:"isBankAccount"`
	// IsInvoiceAccount               bool      `json:"isInvoiceAccount"`
	// BankAccountNumber              string    `json:"bankAccountNumber"`
	// BankAccountCountry             struct {
	// 	ID      int `json:"id"`
	// 	Version int `json:"version"`
	// } `json:"bankAccountCountry"`
	// BankName         string `json:"bankName"`
	// BankAccountIBAN  string `json:"bankAccountIBAN"`
	// BankAccountSWIFT string `json:"bankAccountSWIFT"`
}

type Customers []Customer

type Customer struct {
	ID                    int             `json:"id,omitempty"`
	Version               int             `json:"version,omitempty"`
	URL                   string          `json:"url"`
	Name                  string          `json:"name"`
	OrganizationNumber    string          `json:"organizationNumber,omitempty"`
	SupplierNumber        int             `json:"supplierNumber,omitempty"`
	CustomerNumber        int             `json:"customerNumber,omitempty"`
	IsSupplier            bool            `json:"isSupplier,omitempty"`
	IsCustomer            bool            `json:"isCustomer,omitempty"`
	IsInactive            bool            `json:"isInactive,omitempty"`
	AccountManager        *AccountManager `json:"accountManager,omitempty"`
	Email                 string          `json:"email,omitempty"`
	InvoiceEmail          string          `json:"invoiceEmail,omitempty"`
	OverdueNoticeEmail    string          `json:"overdueNoticeEmail,omitempty"`
	BankAccounts          []string        `json:"bankAccounts,omitempty"`
	PhoneNumber           string          `json:"phoneNumber,omitempty"`
	PhoneNumberMobile     string          `json:"phoneNumberMobile,omitempty"`
	Description           string          `json:"description,omitempty"`
	IsPrivateIndividual   bool            `json:"isPrivateIndividual,omitempty"`
	SingleCustomerInvoice bool            `json:"singleCustomerInvoice,omitempty"`
	InvoiceSendMethod     string          `json:"invoiceSendMethod,omitempty"`
	EmailAttachmentType   string          `json:"emailAttachmentType,omitempty"`
	PostalAddress         *Address        `json:"postalAddress,omitempty"`
	PhysicalAddress       *Address        `json:"physicalAddress,omitempty"`
	DeliveryAddress       *Address        `json:"deliveryAddress,omitempty"`
	Category1             *Category       `json:"category1,omitempty"`
	Category2             *Category       `json:"category2,omitempty"`
	Category3             *Category       `json:"category3,omitempty"`
	InvoicesDueIn         int             `json:"invoicesDueIn,omitempty"`
	InvoicesDueInType     string          `json:"invoicesDueInType,omitempty"`
}

type Supplier struct {
	// 		ID                  int      `json:"id"`
	// 		Version             int      `json:"version"`
	// 		Name                string   `json:"name"`
	// 		OrganizationNumber  string   `json:"organizationNumber"`
	// 		SupplierNumber      int      `json:"supplierNumber"`
	// 		CustomerNumber      int      `json:"customerNumber"`
	// 		IsCustomer          bool     `json:"isCustomer"`
	// 		Email               string   `json:"email"`
	// 		BankAccounts        []string `json:"bankAccounts"`
	// 		InvoiceEmail        string   `json:"invoiceEmail"`
	// 		OverdueNoticeEmail  string   `json:"overdueNoticeEmail"`
	// 		PhoneNumber         string   `json:"phoneNumber"`
	// 		PhoneNumberMobile   string   `json:"phoneNumberMobile"`
	// 		Description         string   `json:"description"`
	// 		IsPrivateIndividual bool     `json:"isPrivateIndividual"`
	// 		ShowProducts        bool     `json:"showProducts"`
	// 		AccountManager      struct {
	// 			ID                       int    `json:"id"`
	// 			Version                  int    `json:"version"`
	// 			FirstName                string `json:"firstName"`
	// 			LastName                 string `json:"lastName"`
	// 			EmployeeNumber           string `json:"employeeNumber"`
	// 			DateOfBirth              string `json:"dateOfBirth"`
	// 			Email                    string `json:"email"`
	// 			PhoneNumberMobileCountry struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"phoneNumberMobileCountry"`
	// 			PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 			PhoneNumberHome        string `json:"phoneNumberHome"`
	// 			PhoneNumberWork        string `json:"phoneNumberWork"`
	// 			NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 			Dnumber                string `json:"dnumber"`
	// 			InternationalID        struct {
	// 				IntAmeldingType string `json:"intAmeldingType"`
	// 				Country         struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 				Number string `json:"number"`
	// 			} `json:"internationalId"`
	// 			BankAccountNumber     string `json:"bankAccountNumber"`
	// 			Iban                  string `json:"iban"`
	// 			Bic                   string `json:"bic"`
	// 			CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 			UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 			UserType              string `json:"userType"`
	// 			Comments              string `json:"comments"`
	// 			Address               struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				AddressLine1 string `json:"addressLine1"`
	// 				AddressLine2 string `json:"addressLine2"`
	// 				PostalCode   string `json:"postalCode"`
	// 				City         string `json:"city"`
	// 				Country      struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 			} `json:"address"`
	// 			Department struct {
	// 				ID               int    `json:"id"`
	// 				Version          int    `json:"version"`
	// 				Name             string `json:"name"`
	// 				DepartmentNumber string `json:"departmentNumber"`
	// 			} `json:"department"`
	// 			Employments []struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				EmploymentID string `json:"employmentId"`
	// 				StartDate    string `json:"startDate"`
	// 				EndDate      string `json:"endDate"`
	// 				Division     struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Name               string `json:"name"`
	// 					StartDate          string `json:"startDate"`
	// 					EndDate            string `json:"endDate"`
	// 					OrganizationNumber string `json:"organizationNumber"`
	// 					Municipality       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"municipality"`
	// 				} `json:"division"`
	// 				LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 				NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 				IsMainEmployer           bool   `json:"isMainEmployer"`
	// 				TaxDeductionCode         string `json:"taxDeductionCode"`
	// 				EmploymentDetails        []struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Date               string `json:"date"`
	// 					EmploymentType     string `json:"employmentType"`
	// 					MaritimeEmployment struct {
	// 						ShipRegister string `json:"shipRegister"`
	// 						ShipType     string `json:"shipType"`
	// 						TradeArea    string `json:"tradeArea"`
	// 					} `json:"maritimeEmployment"`
	// 					RemunerationType   string `json:"remunerationType"`
	// 					WorkingHoursScheme string `json:"workingHoursScheme"`
	// 					ShiftDurationHours int    `json:"shiftDurationHours"`
	// 					OccupationCode     struct {
	// 						ID      int    `json:"id"`
	// 						Version int    `json:"version"`
	// 						NameNO  string `json:"nameNO"`
	// 						Code    string `json:"code"`
	// 					} `json:"occupationCode"`
	// 					PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 					AnnualSalary                   int `json:"annualSalary"`
	// 					HourlyWage                     int `json:"hourlyWage"`
	// 					PayrollTaxMunicipalityID       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"payrollTaxMunicipalityId"`
	// 				} `json:"employmentDetails"`
	// 			} `json:"employments"`
	// 			HolidayAllowanceEarned struct {
	// 				Year                   int `json:"year"`
	// 				Amount                 int `json:"amount"`
	// 				Basis                  int `json:"basis"`
	// 				AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 			} `json:"holidayAllowanceEarned"`
	// 		} `json:"accountManager"`
	// 		PostalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"postalAddress"`
	// 		PhysicalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"physicalAddress"`
	// 		DeliveryAddress struct {
	// 			ID       int `json:"id"`
	// 			Version  int `json:"version"`
	// 			Employee struct {
	// 				ID                       int    `json:"id"`
	// 				Version                  int    `json:"version"`
	// 				FirstName                string `json:"firstName"`
	// 				LastName                 string `json:"lastName"`
	// 				EmployeeNumber           string `json:"employeeNumber"`
	// 				DateOfBirth              string `json:"dateOfBirth"`
	// 				Email                    string `json:"email"`
	// 				PhoneNumberMobileCountry struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"phoneNumberMobileCountry"`
	// 				PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 				PhoneNumberHome        string `json:"phoneNumberHome"`
	// 				PhoneNumberWork        string `json:"phoneNumberWork"`
	// 				NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 				Dnumber                string `json:"dnumber"`
	// 				InternationalID        struct {
	// 					IntAmeldingType string `json:"intAmeldingType"`
	// 					Country         struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 					Number string `json:"number"`
	// 				} `json:"internationalId"`
	// 				BankAccountNumber     string `json:"bankAccountNumber"`
	// 				Iban                  string `json:"iban"`
	// 				Bic                   string `json:"bic"`
	// 				CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 				UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 				UserType              string `json:"userType"`
	// 				Comments              string `json:"comments"`
	// 				Address               struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					AddressLine1 string `json:"addressLine1"`
	// 					AddressLine2 string `json:"addressLine2"`
	// 					PostalCode   string `json:"postalCode"`
	// 					City         string `json:"city"`
	// 					Country      struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 				} `json:"address"`
	// 				Department struct {
	// 					ID               int    `json:"id"`
	// 					Version          int    `json:"version"`
	// 					Name             string `json:"name"`
	// 					DepartmentNumber string `json:"departmentNumber"`
	// 				} `json:"department"`
	// 				Employments []struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					EmploymentID string `json:"employmentId"`
	// 					StartDate    string `json:"startDate"`
	// 					EndDate      string `json:"endDate"`
	// 					Division     struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Name               string `json:"name"`
	// 						StartDate          string `json:"startDate"`
	// 						EndDate            string `json:"endDate"`
	// 						OrganizationNumber string `json:"organizationNumber"`
	// 						Municipality       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"municipality"`
	// 					} `json:"division"`
	// 					LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 					NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 					IsMainEmployer           bool   `json:"isMainEmployer"`
	// 					TaxDeductionCode         string `json:"taxDeductionCode"`
	// 					EmploymentDetails        []struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Date               string `json:"date"`
	// 						EmploymentType     string `json:"employmentType"`
	// 						MaritimeEmployment struct {
	// 							ShipRegister string `json:"shipRegister"`
	// 							ShipType     string `json:"shipType"`
	// 							TradeArea    string `json:"tradeArea"`
	// 						} `json:"maritimeEmployment"`
	// 						RemunerationType   string `json:"remunerationType"`
	// 						WorkingHoursScheme string `json:"workingHoursScheme"`
	// 						ShiftDurationHours int    `json:"shiftDurationHours"`
	// 						OccupationCode     struct {
	// 							ID      int    `json:"id"`
	// 							Version int    `json:"version"`
	// 							NameNO  string `json:"nameNO"`
	// 							Code    string `json:"code"`
	// 						} `json:"occupationCode"`
	// 						PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 						AnnualSalary                   int `json:"annualSalary"`
	// 						HourlyWage                     int `json:"hourlyWage"`
	// 						PayrollTaxMunicipalityID       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"payrollTaxMunicipalityId"`
	// 					} `json:"employmentDetails"`
	// 				} `json:"employments"`
	// 				HolidayAllowanceEarned struct {
	// 					Year                   int `json:"year"`
	// 					Amount                 int `json:"amount"`
	// 					Basis                  int `json:"basis"`
	// 					AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 				} `json:"holidayAllowanceEarned"`
	// 			} `json:"employee"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 			Name string `json:"name"`
	// 		} `json:"deliveryAddress"`
	// 		Category1 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category1"`
	// 		Category2 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category2"`
	// 		Category3 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category3"`
}

type Employee struct {
	// 	Employee struct {
	// 		ID                       int    `json:"id"`
	// 		Version                  int    `json:"version"`
	// 		FirstName                string `json:"firstName"`
	// 		LastName                 string `json:"lastName"`
	// 		EmployeeNumber           string `json:"employeeNumber"`
	// 		DateOfBirth              string `json:"dateOfBirth"`
	// 		Email                    string `json:"email"`
	// 		PhoneNumberMobileCountry struct {
	// 			ID      int `json:"id"`
	// 			Version int `json:"version"`
	// 		} `json:"phoneNumberMobileCountry"`
	// 		PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 		PhoneNumberHome        string `json:"phoneNumberHome"`
	// 		PhoneNumberWork        string `json:"phoneNumberWork"`
	// 		NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 		Dnumber                string `json:"dnumber"`
	// 		InternationalID        struct {
	// 			IntAmeldingType string `json:"intAmeldingType"`
	// 			Country         struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 			Number string `json:"number"`
	// 		} `json:"internationalId"`
	// 		BankAccountNumber     string `json:"bankAccountNumber"`
	// 		Iban                  string `json:"iban"`
	// 		Bic                   string `json:"bic"`
	// 		CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 		UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 		UserType              string `json:"userType"`
	// 		Comments              string `json:"comments"`
	// 		Address               struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"address"`
	// 		Department struct {
	// 			ID               int    `json:"id"`
	// 			Version          int    `json:"version"`
	// 			Name             string `json:"name"`
	// 			DepartmentNumber string `json:"departmentNumber"`
	// 		} `json:"department"`
	// 		Employments []struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			EmploymentID string `json:"employmentId"`
	// 			StartDate    string `json:"startDate"`
	// 			EndDate      string `json:"endDate"`
	// 			Division     struct {
	// 				ID                 int    `json:"id"`
	// 				Version            int    `json:"version"`
	// 				Name               string `json:"name"`
	// 				StartDate          string `json:"startDate"`
	// 				EndDate            string `json:"endDate"`
	// 				OrganizationNumber string `json:"organizationNumber"`
	// 				Municipality       struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"municipality"`
	// 			} `json:"division"`
	// 			LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 			NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 			IsMainEmployer           bool   `json:"isMainEmployer"`
	// 			TaxDeductionCode         string `json:"taxDeductionCode"`
	// 			EmploymentDetails        []struct {
	// 				ID                 int    `json:"id"`
	// 				Version            int    `json:"version"`
	// 				Date               string `json:"date"`
	// 				EmploymentType     string `json:"employmentType"`
	// 				MaritimeEmployment struct {
	// 					ShipRegister string `json:"shipRegister"`
	// 					ShipType     string `json:"shipType"`
	// 					TradeArea    string `json:"tradeArea"`
	// 				} `json:"maritimeEmployment"`
	// 				RemunerationType   string `json:"remunerationType"`
	// 				WorkingHoursScheme string `json:"workingHoursScheme"`
	// 				ShiftDurationHours int    `json:"shiftDurationHours"`
	// 				OccupationCode     struct {
	// 					ID      int    `json:"id"`
	// 					Version int    `json:"version"`
	// 					NameNO  string `json:"nameNO"`
	// 					Code    string `json:"code"`
	// 				} `json:"occupationCode"`
	// 				PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 				AnnualSalary                   int `json:"annualSalary"`
	// 				HourlyWage                     int `json:"hourlyWage"`
	// 				PayrollTaxMunicipalityID       struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"payrollTaxMunicipalityId"`
	// 			} `json:"employmentDetails"`
	// 		} `json:"employments"`
	// 		HolidayAllowanceEarned struct {
	// 			Year                   int `json:"year"`
	// 			Amount                 int `json:"amount"`
	// 			Basis                  int `json:"basis"`
	// 			AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 		} `json:"holidayAllowanceEarned"`
	// 	} `json:"employee"`
}

type Project struct {
}

type Product struct {
}

type Department struct {
	ID               int    `json:"id"`
	Version          int    `json:"version"`
	URL              string `json:"url"`
	Name             string `json:"name"`
	DepartmentNumber string `json:"departmentNumber"`
}

type VATTypes []VATType

type VATType struct {
	ID         int     `json:"id"`
	Version    int     `json:"version"`
	URL        string  `json:"url"`
	Name       string  `json:"name"`
	Number     string  `json:"number"`
	Percentage float64 `json:"percentage"`
}

type Currency struct {
	// 		ID          int    `json:"id"`
	// 		Version     int    `json:"version"`
	// 		Code        string `json:"code"`
	// 		Description string `json:"description"`
	// 		Factor      int    `json:"factor"`
}

type CloseGroup struct {
	// 		ID      int    `json:"id"`
	// 		Version int    `json:"version"`
	// 		Date    string `json:"date"`
}

type Document struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Attachment struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type EDIDocument struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Invoice struct {
	ID             int      `json:"id"`
	Version        int      `json:"version"`
	URL            string   `json:"url"`
	InvoiceNumber  int      `json:"invoiceNumber"`
	InvoiceDate    string   `json:"invoiceDate"`
	Customer       Customer `json:"customer"`
	InvoiceDueDate string   `json:"invoiceDueDate"`
	Kid            string   `json:"kid"`
	Comment        string   `json:"comment"`
	Orders         Orders   `json:"orders"`
	Voucher        Voucher  `json:"voucher"`
	Currency       Currency `json:"currency"`
	InvoiceRemarks string   `json:"invoiceRemarks"`
	PaymentTypeID  int      `json:"paymentTypeId"`
	PaidAmount     int      `json:"paidAmount"`
	EhfSendStatus  string   `json:"ehfSendStatus"`
}

type Voucher struct {
	ID          int          `json:"id,omitempty"`
	Version     int          `json:"version,omitempty"`
	Date        string       `json:"date"`
	Description string       `json:"description"`
	VoucherType *VoucherType `json:"voucherType,omitempty"`
	Postings    Postings     `json:"postings"`
	Document    *Document    `json:"document,omitempty"`
	Attachment  *Attachment  `json:"attachment,omitempty"`
	EDIDocument *EDIDocument `json:"ediDocument,omitempty"`
}

type Orders []Order

type Order struct {
	ID                 int      `json:"id"`
	Version            int      `json:"version"`
	URL                string   `json:"url"`
	Customer           Customer `json:"customer"`
	Contact            Contact  `json:"contact"`
	Attn               Attn     `json:"attn"`
	ReceiverEmail      string   `json:"receiverEmail"`
	OverdueNoticeEmail string   `json:"overdueNoticeEmail"`
	Number             string   `json:"number"`
	Reference          string   `json:"reference"`
	OurContact         Contact  `json:"ourContact"`
	OurContactEmployee Contact  `json:"ourContactEmployee"`
	Department         struct {
		ID               int    `json:"id"`
		Version          int    `json:"version"`
		Name             string `json:"name"`
		DepartmentNumber string `json:"departmentNumber"`
	} `json:"department"`
	OrderDate                                   string     `json:"orderDate"`
	Project                                     Project    `json:"project"`
	InvoiceComment                              string     `json:"invoiceComment"`
	Currency                                    Currency   `json:"currency"`
	InvoicesDueIn                               int        `json:"invoicesDueIn"`
	InvoicesDueInType                           string     `json:"invoicesDueInType"`
	IsShowOpenPostsOnInvoices                   bool       `json:"isShowOpenPostsOnInvoices"`
	IsClosed                                    bool       `json:"isClosed"`
	DeliveryDate                                string     `json:"deliveryDate"`
	DeliveryAddress                             Address    `json:"deliveryAddress"`
	DeliveryComment                             string     `json:"deliveryComment"`
	IsPrioritizeAmountsIncludingVat             bool       `json:"isPrioritizeAmountsIncludingVat"`
	OrderLineSorting                            string     `json:"orderLineSorting"`
	OrderLines                                  OrderLines `json:"orderLines"`
	IsSubscription                              bool       `json:"isSubscription"`
	SubscriptionDuration                        int        `json:"subscriptionDuration"`
	SubscriptionDurationType                    string     `json:"subscriptionDurationType"`
	SubscriptionPeriodsOnInvoice                int        `json:"subscriptionPeriodsOnInvoice"`
	SubscriptionInvoicingTimeInAdvanceOrArrears string     `json:"subscriptionInvoicingTimeInAdvanceOrArrears"`
	SubscriptionInvoicingTime                   int        `json:"subscriptionInvoicingTime"`
	SubscriptionInvoicingTimeType               string     `json:"subscriptionInvoicingTimeType"`
	IsSubscriptionAutoInvoicing                 bool       `json:"isSubscriptionAutoInvoicing"`
}

type Contact struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer"`
}

type Attn struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer"`
}

type Address struct {
	ID           int      `json:"id"`
	Version      int      `json:"version"`
	URL          string   `json:"url"`
	Employee     Employee `json:"employee"`
	AddressLine1 string   `json:"addressLine1"`
	AddressLine2 string   `json:"addressLine2"`
	PostalCode   string   `json:"postalCode"`
	City         string   `json:"city"`
	Country      struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"country"`
	Name string `json:"name"`
}

type OrderLines []OrderLine

type OrderLine struct {
	ID        int     `json:"id"`
	Version   int     `json:"version"`
	URL       string  `json:"url"`
	Product   Product `json:"product"`
	Inventory struct {
		ID              int    `json:"id"`
		Version         int    `json:"version"`
		Name            string `json:"name"`
		Number          string `json:"number"`
		IsMainInventory bool   `json:"isMainInventory"`
		IsInactive      bool   `json:"isInactive"`
	} `json:"inventory"`
	Description                   string   `json:"description"`
	Count                         int      `json:"count"`
	UnitCostCurrency              int      `json:"unitCostCurrency"`
	UnitPriceExcludingVatCurrency int      `json:"unitPriceExcludingVatCurrency"`
	Currency                      Currency `json:"currency"`
	Markup                        int      `json:"markup"`
	Discount                      int      `json:"discount"`
	VATType                       VATType  `json:"vatType"`
	UnitPriceIncludingVatCurrency int      `json:"unitPriceIncludingVatCurrency"`
	IsSubscription                bool     `json:"isSubscription"`
	SubscriptionPeriodStart       string   `json:"subscriptionPeriodStart"`
	SubscriptionPeriodEnd         string   `json:"subscriptionPeriodEnd"`
	OrderGroup                    struct {
		ID        int    `json:"id"`
		Version   int    `json:"version"`
		Title     string `json:"title"`
		Comment   string `json:"comment"`
		SortIndex int    `json:"sortIndex"`
	} `json:"orderGroup"`
}

type AccountManager struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	EmployeeNumber           string `json:"employeeNumber"`
	DateOfBirth              string `json:"dateOfBirth"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile      string `json:"phoneNumberMobile"`
	PhoneNumberHome        string `json:"phoneNumberHome"`
	PhoneNumberWork        string `json:"phoneNumberWork"`
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	Dnumber                string `json:"dnumber"`
	InternationalID        struct {
		IntAmeldingType string `json:"intAmeldingType"`
		Country         struct {
			ID      int `json:"id"`
			Version int `json:"version"`
		} `json:"country"`
		Number string `json:"number"`
	} `json:"internationalId"`
	BankAccountNumber     string     `json:"bankAccountNumber"`
	Iban                  string     `json:"iban"`
	Bic                   string     `json:"bic"`
	CreditorBankCountryID int        `json:"creditorBankCountryId"`
	UsesAbroadPayment     bool       `json:"usesAbroadPayment"`
	UserType              string     `json:"userType"`
	Comments              string     `json:"comments"`
	Address               Address    `json:"address"`
	Department            Department `json:"department"`
	Employments           []struct {
		ID           int    `json:"id"`
		Version      int    `json:"version"`
		EmploymentID string `json:"employmentId"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Division     struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Name               string `json:"name"`
			StartDate          string `json:"startDate"`
			EndDate            string `json:"endDate"`
			OrganizationNumber string `json:"organizationNumber"`
			Municipality       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"municipality"`
		} `json:"division"`
		LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
		NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
		IsMainEmployer           bool   `json:"isMainEmployer"`
		TaxDeductionCode         string `json:"taxDeductionCode"`
		EmploymentDetails        []struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Date               string `json:"date"`
			EmploymentType     string `json:"employmentType"`
			MaritimeEmployment struct {
				ShipRegister string `json:"shipRegister"`
				ShipType     string `json:"shipType"`
				TradeArea    string `json:"tradeArea"`
			} `json:"maritimeEmployment"`
			RemunerationType   string `json:"remunerationType"`
			WorkingHoursScheme string `json:"workingHoursScheme"`
			ShiftDurationHours int    `json:"shiftDurationHours"`
			OccupationCode     struct {
				ID      int    `json:"id"`
				Version int    `json:"version"`
				NameNO  string `json:"nameNO"`
				Code    string `json:"code"`
			} `json:"occupationCode"`
			PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
			AnnualSalary                   int `json:"annualSalary"`
			HourlyWage                     int `json:"hourlyWage"`
			PayrollTaxMunicipalityID       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"payrollTaxMunicipalityId"`
		} `json:"employmentDetails"`
	} `json:"employments"`
	HolidayAllowanceEarned struct {
		Year                   int `json:"year"`
		Amount                 int `json:"amount"`
		Basis                  int `json:"basis"`
		AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	} `json:"holidayAllowanceEarned"`
}

type Category struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}
