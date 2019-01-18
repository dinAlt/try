package try

import "github.com/pkg/errors"

var (
	ErrEmptyValue = errors.New("Value is empty")
	ErrWrongType  = errors.New("Value has wrong type")
)

func GetMap(value interface{}) (map[string]interface{}, error) {
	if value == nil {
		return nil, ErrEmptyValue
	}

	switch v := value.(type) {
	case map[string]interface{}:
		if v == nil {
			return nil, ErrEmptyValue
		}
		return v, nil
	case *map[string]interface{}:
		if v == nil {
			return nil, ErrEmptyValue
		}
		if *v == nil {
			return nil, ErrEmptyValue
		}
		return *v, nil
	}

	return nil, ErrWrongType
}

func GetStringsMap(value map[string]interface{}) (map[string]string, error) {
	res := make(map[string]string)

	for key := range value {
		v, err := GetString(value[key])
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res[key] = v
	}
	return res, nil
}

func GetInt64Map(value map[string]interface{}) (map[string]int64, error) {
	res := make(map[string]int64)

	for key := range value {
		v, err := GetInt64(value[key])
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res[key] = v
	}
	return res, nil
}

func GetSlice(value interface{}) ([]interface{}, error) {
	if value == nil {
		return nil, ErrEmptyValue
	}

	switch v := value.(type) {
	case []interface{}:
		if v == nil {
			return nil, ErrEmptyValue
		}
		return v, nil
	case *[]interface{}:
		if v == nil {
			return nil, ErrEmptyValue
		}
		if *v == nil {
			if v == nil {
				return nil, ErrEmptyValue
			}
		}
		return *v, nil
	}
	return nil, ErrWrongType
}

func GetBool(value interface{}) (bool, error) {
	if value == nil {
		return false, ErrEmptyValue
	}

	switch v := value.(type) {
	case *bool:
		if v == nil {
			return false, ErrEmptyValue
		}
		return *v, nil
	case bool:
		return v, nil
	}

	return false, ErrWrongType
}

func GetInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, ErrEmptyValue
	}

	switch v := value.(type) {
	case *int64:
		if v == nil {
			return 0, ErrEmptyValue
		}
		return *v, nil
	case int64:
		return v, nil
	case *int32:
		if v == nil {
			return 0, ErrEmptyValue
		}
		return int64(*v), nil
	case int32:
		return int64(v), nil
	case *float64:
		if v == nil {
			return 0, nil
		}
		return int64(*v), nil
	case float64:
		return int64(v), nil
	}

	return 0, ErrWrongType
}

func GetString(value interface{}) (string, error) {
	if value == nil {
		return "", ErrEmptyValue
	}

	switch v := value.(type) {
	case *string:
		{
			if v == nil {
				return "", ErrEmptyValue
			}
			return *v, nil
		}
	case string:
		return v, nil
	}

	return "", ErrWrongType
}
