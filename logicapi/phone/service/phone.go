package service

import (
	"auto_call_phone/common/base"
	"auto_call_phone/logicapi/excel/viewmodels"
	"context"
)

type PhoneService struct {
}

func NewPhoneService() *PhoneService {
	return &PhoneService{}
}

func (p PhoneService) UploadExcel(c context.Context, commonParams *base.CommonParams, req *viewmodels.UploadPhoneExcelReq) (*viewmodels.UploadPhoneExcelRsp, error) {
	return nil, nil
}
