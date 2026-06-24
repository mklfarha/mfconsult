package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=DocumentKind -json
type DocumentKind int64

const (
	DOCUMENT_KIND_INVALID = iota
	DOCUMENT_KIND_FILE
	DOCUMENT_KIND_REPO_URL
	DOCUMENT_KIND_LINK
)

func (e DocumentKind) ToInt64() int64 {
	return int64(e)
}

func (e DocumentKind) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func DocumentKindFromString(in string) DocumentKind {
	switch in {
	case "invalid":
		return DOCUMENT_KIND_INVALID
	case "file":
		return DOCUMENT_KIND_FILE
	case "repo_url":
		return DOCUMENT_KIND_REPO_URL
	case "link":
		return DOCUMENT_KIND_LINK
	}
	return DOCUMENT_KIND_INVALID
}

func DocumentKindFromPointerString(in *string) DocumentKind {
	if in == nil {
		return DOCUMENT_KIND_INVALID
	}
	return DocumentKindFromString(*in)
}

func (e DocumentKind) String() string {
	switch e {
	case DOCUMENT_KIND_INVALID:
		return "invalid"
	case DOCUMENT_KIND_FILE:
		return "file"
	case DOCUMENT_KIND_REPO_URL:
		return "repo_url"
	case DOCUMENT_KIND_LINK:
		return "link"
	}

	return "invalid"
}

func (e DocumentKind) StringPtr() *string {
	val := e.String()
	return &val
}

func DocumentKindSliceToJSON(in []DocumentKind) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling DocumentKind slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToDocumentKindSlice(in json.RawMessage) []DocumentKind {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling DocumentKind slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []DocumentKind{}
	for _, r := range res {
		finalRes = append(finalRes, DocumentKind(r))
	}
	return finalRes
}
