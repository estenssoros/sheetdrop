package process

import (
	"io"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/common"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// Excel processes an excel row into a schema
func Excel(r io.ReaderAt, size int64) (interface{}, error) {
	xlFile, err := xlsx.OpenReaderAtWithRowLimit(r, size, constants.InitialRowLimit)
	if err != nil {
		return nil, errors.Wrap(err, "xlsx.OpenBinary")
	}
	schemas := []*models.Schema{}
	for _, sheet := range xlFile.Sheets {
		schema, err := excelSheet(sheet)
		if err != nil {
			return nil, errors.Wrap(err, "excelSheet")
		}
		schemas = append(schemas, schema)
	}
	return schemas, nil
}

func excelSheet(sheet *xlsx.Sheet) (*models.Schema, error) {
	if !sheetHasData(sheet) {
		return nil, common.ErrNoData
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
	dataTypes, err := sheetDataTypes(sheet, startRow, headers)
	if err != nil {
		return nil, errors.Wrap(err, "sheetDataTypes")
	}
	return &models.Schema{
		StartRow:    startRow,
		StartColumn: startColumn,
		Headers:     headers,
		DataTypes:   dataTypes,
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

func sheetHeaders(sheet *xlsx.Sheet, startRow, startColumn int) (headers map[string]int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%v", r)
		}
	}()
	row := sheet.Rows[startRow]
	headers = map[string]int{}
	for i, cell := range row.Cells[startColumn:] {
		headerName := helpers.CamelCase(cell.Value)
		if _, ok := headers[headerName]; ok {
			return nil, errors.Errorf("duplicate headers: %s", headerName)
		}
		headers[headerName] = startColumn + i
	}
	return headers, nil
}

func sheetDataTypes(sheet *xlsx.Sheet, startRow int, headers map[string]int) (dataTypes []*models.DataType, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%v", r)
		}
	}()
	for _, headerIndex := range headers {
		dataType, err := columnDataType(sheet.Rows[startRow+1:], headerIndex)
		if err != nil {
			return nil, errors.Wrap(err, "columnDataType")
		}
		dataTypes = append(dataTypes, dataType)
	}
	return
}

func columnDataType(rows []*xlsx.Row, headerIdx int) (dataType *models.DataType, err error) {
	dataType = &models.DataType{Idx: headerIdx}
	if err := tryDataType(rows, headerIdx, validateTime); err == nil {
		dataType.Type = models.DataTypeTime
		return dataType, err
	}
	if err := tryDataType(rows, headerIdx, validateFloat); err == nil {
		dataType.Type = models.DataTypeInt
		return dataType, err
	}
	if err := tryDataType(rows, headerIdx, validateFloat); err == nil {
		dataType.Type = models.DataTypeFloat
		return dataType, err
	}
	dataType.Type = models.DataTypeString
	return
}

func tryDataType(rows []*xlsx.Row, headerIdx int, validator func(*xlsx.Cell) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%v", r)
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

func validateTime(cell *xlsx.Cell) error {
	if cell.IsTime() {
		return nil
	}
	return errors.New("isTime")
}

func validateInt(cell *xlsx.Cell) error {
	_, err := cell.Int()
	return err
}

func validateFloat(cell *xlsx.Cell) error {
	_, err := cell.Float()
	return err
}
