package process

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
	"github.com/sirupsen/logrus"
)

func CSV(schema *models.Schema, data []byte) error {
	logrus.Info("processing csv")
	r := csv.NewReader(bytes.NewReader(data))
	headers := []*models.Header{}
	rows := [][]string{}
	var count int
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "reader.Read")
		}
		if len(headers) == 0 {
			hs, err := csvHeaders(row)
			if err != nil {
				return errors.Wrap(err, "csvHeaders")
			}
			headers = hs
			continue
		}
		rows = append(rows, row)
		count++
		if count > constants.InitialRowLimit {
			break
		}
	}
	if err := csvDataTypes(rows, headers); err != nil {
		return errors.Wrap(err, "csvDataTypes")
	}

	schema.StartRow = 1
	schema.StartColumn = 1
	schema.Headers = headers
	schema.SourceType = constants.SourceTypeCSV

	return nil
}

func csvHeaders(row []string) ([]*models.Header, error) {
	unique := map[string]struct{}{}
	headers := []*models.Header{}
	for i, cell := range row {
		headerName := helpers.CamelCase(cell)
		if _, ok := unique[headerName]; ok {
			return nil, errors.Errorf("duplicate headers: %s", headerName)
		}
		headers = append(headers, &models.Header{
			Name:  headerName,
			Index: uint(i),
		})
	}
	return headers, nil
}

func csvDataTypes(rows [][]string, headers []*models.Header) error {
	for _, header := range headers {
		dataType, err := columnDataTypeCSV(rows, int(header.Index))
		if err != nil {
			return errors.Wrap(err, "columnDataTypeCSV")
		}
		header.DataType = dataType
	}
	return nil
}

func columnDataTypeCSV(rows [][]string, headerIdx int) (dataType string, err error) {
	if err := tryDataTypeCSV(rows, headerIdx, validateTimeCSV); err == nil {
		return models.DataTypeTime, err
	}
	if err := tryDataTypeCSV(rows, headerIdx, validateIntCSV); err == nil {
		return models.DataTypeFloat, err
	}
	if err := tryDataTypeCSV(rows, headerIdx, validateFloatCSV); err == nil {
		return models.DataTypeInt, err
	}
	if err := tryDataTypeCSV(rows, headerIdx, validateUUIDCSV); err == nil {
		return models.DataTypeUUID, err
	}
	return models.DataTypeString, nil
}

func tryDataTypeCSV(rows [][]string, headerIdx int, validator func(string) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%v", r)
		}
	}()
	var found bool
	for _, row := range rows {
		cell := row[headerIdx]
		if cell == "" {
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

func validateTimeCSV(cell string) error {
	_, err := time.Parse("2006-01-02 15:04", cell)
	return err
}

func validateFloatCSV(cell string) error {
	_, err := strconv.ParseFloat(cell, 10)
	return err
}

func validateIntCSV(cell string) error {
	_, err := strconv.Atoi(cell)
	return err
}

func validateUUIDCSV(cell string) error {
	_, err := uuid.FromString(cell)
	return err
}
