package bank

import "bitbucket.org/mundipagg/boletoapi/models"

func baseValidateRecipientDocumentNumber(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Recipient.Document.IsCPF() {
			return t.Recipient.Document.ValidateCPF()
		}
		if t.Recipient.Document.IsCNPJ() {
			return t.Recipient.Document.ValidateCNPJ()
		}
		return models.NewErrorResponse("MPRecipientDocumentType", "Tipo de Documento inválido")
	default:
		return invalidType(t)
	}
}

func baseValidateBuyerDocumentNumber(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Buyer.Document.IsCPF() {
			return t.Buyer.Document.ValidateCPF()
		}
		if t.Buyer.Document.IsCNPJ() {
			return t.Buyer.Document.ValidateCNPJ()
		}
		return models.NewErrorResponse("MPBuyerDocumentType", "Tipo de Documento inválido")
	default:
		return invalidType(t)
	}
}

func baseValidateExpireDate(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		return t.Title.IsExpireDateValid()
	default:
		return invalidType(t)
	}
}

func baseValidateAmountInCents(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		return t.Title.IsAmountInCentsValid()
	default:
		return invalidType(t)
	}
}
