package process

import (
	"io"
	"io/ioutil"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

// Excel processes an excel row into a schema
func Excel(schema *models.Schema, file io.Reader) (*Result, error) {
	logrus.Info("processing excel")
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	xlFile, err := xlsx.OpenBinaryWithRowLimit(data, constants.InitialRowLimit)
	if err != nil {
		return nil, errors.Wrap(err, "xlsx.OpenBinary")
	}

	if len(xlFile.Sheets) < 1 {
		return nil, errors.New("no sheets")
	}
	return excelSheet(schema, xlFile.Sheets[0])
}

func excelSheet(schema *models.Schema, sheet *xlsx.Sheet) (*Result, error) {
	if !sheetHasData(sheet) {
		return nil, errors.New("no data")
	}
	startRow, err := sheetStartRow(sheet)
	if err != nil {
		return nil, errors.Wrap(err, "sheetStartRow")
	}
	startColumn, err := sheetStartColumn(sheet, startRow)
	if err != nil {
		return nil, errors.Wrap(err, "sheetStartColumn")
	}
	headers, err := sheetHeaders(sheet, startRow, startColumn)
	if err != nil {
		return nil, errors.Wrap(err, "sheetHeader")
	}
	if err := sheetDataTypes(sheet, startRow, headers); err != nil {
		return nil, errors.Wrap(err, "sheetDataTypes")
	}
	return &Result{
		StartRow:    startRow,
		StartColumn: startColumn,
		Headers:     headers,
	}, nil
}

func sheetHasData(sheet *xlsx.Sheet) bool {
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			if cell.Value != "" {
				return true
			}
		}
	}
	return false
}

func sheetStartRow(sheet *xlsx.Sheet) (int, error) {
	var count, length, startRow int
	var remainingRowCount = len(sheet.Rows)
	for rowNum, row := range sheet.Rows {
		l := rowLength(row)
		if l == length {
			count++
		} else {
			length = l
			remainingRowCount -= count
			startRow = rowNum
			count = 0
		}
		if count == constants.AssumedRowCount || count == remainingRowCount {
			return startRow, nil
		}
	}
	return -1, errors.New("count not determine start row")
}

func rowLength(row *xlsx.Row) int {
	return len(row.Cells)
}

func sheetStartColumn(sheet *xlsx.Sheet, startRow int) (startColumn int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%v", r)
		}
	}()
	row := sheet.Rows[startRow]
	for i, cell := range row.Cells {
		if cell.Value != "" {
			return i, nil
		}
	}
	return -1, errors.New("could not determing start column")
}

func sheetHeaders(sheet *xlsx.Sheet, startRow, startColumn int) (headers []*models.Header, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered err: %v", r)
		}
	}()
	row := sheet.Rows[startRow]
	unique := map[string]struct{}{}
	for i, cell := range row.Cells[startColumn:] {
		headerName := helpers.CamelCase(cell.Value)
		if _, ok := unique[headerName]; ok {
			return nil, errors.Errorf("duplicate headers: %s", headerName)
		}
		headers = append(headers, &models.Header{
			Name:  headerName,
			Index: startColumn + i,
		})
	}
	return headers, nil
}

func sheetDataTypes(sheet *xlsx.Sheet, startRow int, headers []*models.Header) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered err: %v", r)
		}
	}()
	for _, header := range headers {
		dataType, err := columnDataTypeExcel(sheet.Rows[startRow+1:], int(header.Index))
		if err != nil {
			return errors.Wrap(err, "columnDataType")
		}
		header.DataType = dataType
	}
	return
}

func columnDataTypeExcel(rows []*xlsx.Row, headerIdx int) (dataType string, err error) {
	if err := tryDataTypeExcel(rows, headerIdx, validateTimeExcel); err == nil {
		return models.DataTypeTime, err
	}
	if err := tryDataTypeExcel(rows, headerIdx, validateIntExcel); err == nil {
		return models.DataTypeFloat, err
	}
	if err := tryDataTypeExcel(rows, headerIdx, validateFloatExcel); err == nil {
		return models.DataTypeInt, err
	}
	if err := tryDataTypeExcel(rows, headerIdx, validateUUIDExcel); err == nil {
		return models.DataTypeUUID, err
	}
	return models.DataTypeString, nil
}

func tryDataTypeExcel(rows []*xlsx.Row, headerIdx int, validator func(*xlsx.Cell) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered err: %v", r)
		}
	}()
	var found bool
	for _, row := range rows {
		cell := row.Cells[headerIdx]
		if cell.Value == "" {
			continue
		}
		if err := validator(cell); err != nil {
			return err
		}
		found = true
	}
	if !found {
		return errors.New("not found")
	}
	return nil
}

func validateTimeExcel(cell *xlsx.Cell) error {
	if cell.IsTime() {
		return nil
	}
	return errors.New("isTime")
}

func validateIntExcel(cell *xlsx.Cell) error {
	_, err := cell.Int()
	return err
}

func validateFloatExcel(cell *xlsx.Cell) error {
	_, err := cell.Float()
	return err
}

func validateUUIDExcel(cell *xlsx.Cell) error {
	_, err := uuid.FromString(cell.Value)
	return err
}
