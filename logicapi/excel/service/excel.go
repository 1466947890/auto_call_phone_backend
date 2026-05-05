package service

import (
	"auto_call_phone/common/base"
	"auto_call_phone/logicapi/excel/viewmodels"
	"context"
)

type ExcelService struct {
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (p ExcelService) UploadPhoneExcel(c context.Context, commonParams *base.CommonParams, req *viewmodels.UploadPhoneExcelReq) (*viewmodels.UploadPhoneExcelRsp, error) {
	return nil, nil
}
