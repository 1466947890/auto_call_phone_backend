package service

import (
	"auto_call_phone/common/base"
	"auto_call_phone/data/datamodels"
	"auto_call_phone/data/repo"
	"auto_call_phone/logicapi/excel/viewmodels"
	"context"
	"errors"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm/clause"
)

type ExcelService struct {
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (p ExcelService) UploadPhoneExcel(c context.Context, commonParams *base.CommonParams, req *viewmodels.UploadPhoneExcelReq) (*viewmodels.UploadPhoneExcelRsp, error) {
	if strings.TrimSpace(req.DeviceID) == "" {
		return nil, errors.New("device_id is required")
	}
	file, err := req.File.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := parseExcel(file)
	if err != nil {
		return nil, err
	}

	// 去除文件内重复
	seen := make(map[string]struct{}, len(records))
	uniq := make([]viewmodels.PhoneInfo, 0, len(records))
	for _, r := range records {
		if r.PhoneNumber == "" {
			continue
		}
		if _, ok := seen[r.PhoneNumber]; ok {
			continue
		}
		seen[r.PhoneNumber] = struct{}{}
		uniq = append(uniq, r)
	}

	total := len(records)
	if len(uniq) == 0 {
		return &viewmodels.UploadPhoneExcelRsp{
			PhoneList:      []viewmodels.PhoneInfo{},
			EffectiveCount: 0,
			RepeatCount:    total,
			Total:          total,
		}, nil
	}

	// 查询数据库已存在的号码（含软删除，软删的也视为已存在，不重复插入）
	phones := make([]string, 0, len(uniq))
	for _, r := range uniq {
		phones = append(phones, r.PhoneNumber)
	}
	var existed []string
	if err := repo.DB.WithContext(c).Unscoped().
		Model(&datamodels.DevicePhone{}).
		Where("device_id = ? AND phone_number IN ?", req.DeviceID, phones).
		Pluck("phone_number", &existed).Error; err != nil {
		return nil, err
	}
	existedSet := make(map[string]struct{}, len(existed))
	for _, p := range existed {
		existedSet[p] = struct{}{}
	}

	toInsert := make([]datamodels.DevicePhone, 0, len(uniq))
	effective := make([]viewmodels.PhoneInfo, 0, len(uniq))
	for _, r := range uniq {
		if _, ok := existedSet[r.PhoneNumber]; ok {
			continue
		}
		toInsert = append(toInsert, datamodels.DevicePhone{
			PhoneNumber: r.PhoneNumber,
			DeviceID:    req.DeviceID,
			Status:      int8(r.Status),
			Remarks:     r.Remark,
		})
		effective = append(effective, r)
	}

	if len(toInsert) > 0 {
		if err := repo.DB.WithContext(c).
			Clauses(clause.OnConflict{DoNothing: true}).
			CreateInBatches(&toInsert, 500).Error; err != nil {
			return nil, err
		}
	}

	return &viewmodels.UploadPhoneExcelRsp{
		PhoneList:      effective,
		EffectiveCount: len(effective),
		RepeatCount:    total - len(effective),
		Total:          total,
	}, nil
}

func (p ExcelService) GetPhones(c context.Context, commonParams *base.CommonParams, req *viewmodels.GetPhonesReq) (*viewmodels.GetPhonesRsp, error) {
	if strings.TrimSpace(req.DeviceID) == "" {
		return nil, errors.New("device_id is required")
	}
	var rows []datamodels.DevicePhone
	if err := repo.DB.WithContext(c).
		Where("device_id = ?", req.DeviceID).
		Order("id ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	list := make([]viewmodels.PhoneInfo, 0, len(rows))
	for _, r := range rows {
		list = append(list, viewmodels.PhoneInfo{
			ID:          r.ID,
			PhoneNumber: r.PhoneNumber,
			Remark:      r.Remarks,
			Status:      int(r.Status),
		})
	}
	return &viewmodels.GetPhonesRsp{
		PhoneList: list,
		Total:     len(list),
	}, nil
}

func (p ExcelService) DeletePhone(c context.Context, commonParams *base.CommonParams, req *viewmodels.DeletePhoneReq) (*viewmodels.DeletePhoneRsp, error) {
	if req.ID <= 0 {
		return nil, errors.New("id is required")
	}
	res := repo.DB.WithContext(c).Where("id = ?", req.ID).Delete(&datamodels.DevicePhone{})
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("phone not found")
	}
	return &viewmodels.DeletePhoneRsp{ID: req.ID}, nil
}

func (p ExcelService) UpdatePhone(c context.Context, commonParams *base.CommonParams, req *viewmodels.UpdatePhoneReq) (*viewmodels.UpdatePhoneRsp, error) {
	if req.ID <= 0 {
		return nil, errors.New("id is required")
	}
	if req.Remark == nil && req.Status == nil {
		return nil, errors.New("nothing to update")
	}

	updates := make(map[string]interface{}, 2)
	if req.Remark != nil {
		updates["remarks"] = *req.Remark
	}
	if req.Status != nil {
		updates["status"] = int8(*req.Status)
	}

	res := repo.DB.WithContext(c).
		Model(&datamodels.DevicePhone{}).
		Where("id = ?", req.ID).
		Updates(updates)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("phone not found")
	}

	var row datamodels.DevicePhone
	if err := repo.DB.WithContext(c).Where("id = ?", req.ID).First(&row).Error; err != nil {
		return nil, err
	}
	return &viewmodels.UpdatePhoneRsp{
		PhoneInfo: viewmodels.PhoneInfo{
			ID:          row.ID,
			PhoneNumber: row.PhoneNumber,
			Remark:      row.Remarks,
			Status:      int(row.Status),
		},
	}, nil
}

func parseExcel(file io.Reader) ([]viewmodels.PhoneInfo, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	if sheet == "" {
		return nil, errors.New("excel has no sheet")
	}
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	records := make([]viewmodels.PhoneInfo, 0, len(rows))
	for i, row := range rows {
		if len(row) == 0 {
			continue
		}
		phone := strings.TrimSpace(row[0])
		// 跳过表头：首行如果不像号码，直接忽略
		if i == 0 && !looksLikePhone(phone) {
			continue
		}
		if phone == "" {
			continue
		}
		remark := ""
		if len(row) > 1 {
			remark = strings.TrimSpace(row[1])
		}
		records = append(records, viewmodels.PhoneInfo{
			PhoneNumber: phone,
			Remark:      remark,
		})
	}
	return records, nil
}

func looksLikePhone(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '+' || r == '-' || r == ' ' {
			continue
		}
		return false
	}
	return true
}
