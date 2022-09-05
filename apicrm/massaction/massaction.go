package massaction

import "encoding/json"

func (svc *APIMassActionService) Recalculate(model DomainRecalc) (ResponseRecalc, error) {
	payload, err := json.Marshal(model)

	if err != nil {
		return ResponseRecalc{}, err
	}

	response, err := svc.client.CRMHandlerPostService("MassAction", payload)
	if err != nil {
		return ResponseRecalc{}, err
	}

	return convertMarchal(response)
}


// Todos os serviços deverão ter o seu próprio conversor de json para o domain

func convertMarchal(response []byte) (ResponseRecalc, error) {
	var result ResponseRecalc

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseRecalc{}, err
	}

	return result, nil
}
