/**
 * файл с ошибками
 */
package store

import "errors"

var (
	// если запись не найдена
	ErrRecordNotFound = errors.New("record not found")
)