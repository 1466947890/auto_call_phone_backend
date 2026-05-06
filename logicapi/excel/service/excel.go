package service

import (
	"auto_call_phone/common/base"
	"auto_call_phone/logicapi/excel/viewmodels"
	"context"
	"io"

	"github.com/xuri/excelize/v2"
)

type ExcelService struct {
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (p ExcelService) UploadPhoneExcel(c context.Context, commonParams *base.CommonParams, req *viewmodels.UploadPhoneExcelReq) (*viewmodels.UploadPhoneExcelRsp, error) {
	file, err := req.File.Open()
	if err != nil {
		return nil, err
	}
	records, err := paseExcel(file)
	if err != nil {
		return nil, err
	}
	return &viewmodels.UploadPhoneExcelRsp{
		PhoneList: records,
		Total:     len(records),
	}, nil
}

func paseExcel(file io.Reader) ([]viewmodels.PhoneInfo, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rows, err := f.GetRows("Sheet1")
	var records []viewmodels.PhoneInfo
	for _, row := range rows {
		records = append(records, viewmodels.PhoneInfo{
			PhoneNumber: row[0],
			Remark:      row[1],
		})
	}
	return records, nil
}
