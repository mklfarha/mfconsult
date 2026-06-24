package nda_document

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/nda_document/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
)

func (m *module) Insert(
	ctx context.Context,
	req types.UpsertRequest,
	opts ...Option,
) (types.UpsertResponse, error) {

	optConfig := applyAllOptions(opts)

	tx := optConfig.SQLTx
	createdTx := false
	if tx == nil {
		ntx, err := m.repository.DB.Begin()
		if err != nil {
			return types.UpsertResponse{}, err
		}
		tx = ntx
		defer tx.Rollback()
		createdTx = true
	}

	qtx := m.repository.Queries.WithTx(tx)
	params := mapUpsertRequestToInsertParams(req)

	_, err := qtx.InsertNdaDocument(
		ctx,
		params,
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {

			return types.UpsertResponse{}, err
		}
	}

	return buildInsertResponse(req), nil
}

func buildInsertResponse(req types.UpsertRequest) types.UpsertResponse {
	return types.UpsertResponse{

		ID: req.NdaDocument.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertNdaDocumentParams {
	return mfconsultdb.InsertNdaDocumentParams{
		ID: req.NdaDocument.ID.String(),

		ClientId: req.NdaDocument.ClientId.String(),

		URL: req.NdaDocument.URL,

		Status: req.NdaDocument.Status.ToNullInt(),

		SignedAt: req.NdaDocument.SignedAt,

		CreatedAt: req.NdaDocument.CreatedAt,

		EnvelopeId: req.NdaDocument.EnvelopeId,

		CertificateURL: req.NdaDocument.CertificateURL,
	}
}
