package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{partnerBaseURLs: partnerBaseURLs}
}

func (factory *DefaultPartnerFactory) CreatePartner(partnerID int) (Partner, error) {
	baseURL, ok := factory.partnerBaseURLs[partnerID]

	if !ok {
		return nil, fmt.Errorf("unsupported partner id: %d", partnerID)
	}

	switch partnerID {
	case 1:
		return &LocalPartner{BaseURL: baseURL}, nil
	case 2:
		return &ExternalPartner{BaseURL: baseURL}, nil
	default:
		return nil, fmt.Errorf("unsupported partner id: %d", partnerID)
	}
}
