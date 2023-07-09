package processor

import (
	"regexp"
)

type Transaction struct {
	TransactionDate string
	Amount          string
	Merchant        string
}

func ExtractTransactionDetails(message string) *Transaction {
	var (
		transactionDate string
		amount          string
		merchant        string
	)

	// BOI UPI -VPA message format
	boiUPIRegex := regexp.MustCompile(`BOI UPI -VPA`)
	if boiUPIRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`debited Rs.(\d+\.\d+)`)
		merchantRegex := regexp.MustCompile(`credited to (\S+) -Ref`)
		dateRegex := regexp.MustCompile(`on (\d+[A-Za-z]{3}\d+)`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// SBI UPI message format
	sbiUPIRegex := regexp.MustCompile(`Dear SBI User, your A/c`)
	if sbiUPIRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`by Rs(\d+\.\d+)`)
		merchantRegex := regexp.MustCompile(`transfer to (\S+) Ref`)
		dateRegex := regexp.MustCompile(`on (\d+[A-Za-z]{3}\d+)`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// SBIDC message format
	sbiDCRegex := regexp.MustCompile(`Dear SBI Customer, Rs\.\d+ withdrawn at SBI ATM`)
	if sbiDCRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`Rs\.(\d+)`)
		merchantRegex := regexp.MustCompile(`at (\S+) ATM`)
		dateRegex := regexp.MustCompile(`on (\d+[A-Za-z]{3}\d+)`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// ICICICC message format
	iciciCCRegex := regexp.MustCompile(`INR \d+\.\d+ spent on ICICI Bank Card`)
	if iciciCCRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`INR (\d+\.\d+)`)
		merchantRegex := regexp.MustCompile(`at (\S+)`)
		dateRegex := regexp.MustCompile(`on (\d+-[A-Za-z]{3}-\d+)`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// IDFCCC message format
	idfcCCRegex := regexp.MustCompile(`Transaction Successful! INR \d+\.\d+ spent on your IDFC FIRST Bank Credit Card`)
	if idfcCCRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`INR (\d+\.\d+)`)
		merchantRegex := regexp.MustCompile(`ending XX\d+ at (\S+)`)
		dateRegex := regexp.MustCompile(`on (\d+-[A-Za-z]{3}-\d+)`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// STANCCC message format
	stanCCRegex := regexp.MustCompile(`Thank you for using StanChart Credit Card No`)
	if stanCCRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`for INR (\d+\.\d+) at (\S+)`)
		dateRegex := regexp.MustCompile(`on (\d{2}/\d{2}/\d{2})`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
			merchant = amountMatch[2]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	// AMEXCC message format
	amexCCRegex := regexp.MustCompile(`Alert: You've spent INR (\d+,\d+\.\d+) on your AMEX card`)
	if amexCCRegex.MatchString(message) {
		amountRegex := regexp.MustCompile(`INR (\d+,\d+\.\d+)`)
		merchantRegex := regexp.MustCompile(`at (\S+)`)
		dateRegex := regexp.MustCompile(`on (\d{1,2} [A-Za-z]+, \d{4})`)

		amountMatch := amountRegex.FindStringSubmatch(message)
		merchantMatch := merchantRegex.FindStringSubmatch(message)
		dateMatch := dateRegex.FindStringSubmatch(message)

		if len(amountMatch) >= 2 {
			amount = amountMatch[1]
		}
		if len(merchantMatch) >= 2 {
			merchant = merchantMatch[1]
		}
		if len(dateMatch) >= 2 {
			transactionDate = dateMatch[1]
		}
	}

	return &Transaction{
		TransactionDate: transactionDate,
		Amount:          amount,
		Merchant:        merchant,
	}
}

