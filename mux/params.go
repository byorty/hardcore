package mux

import "strconv"

type RequestScopeParamsImpl map[string]string

func (r RequestScopeParamsImpl) GetString(key string) string {
	if value, ok := r[key]; ok {
		return value
	} else {
		return ""
	}
}

func (r RequestScopeParamsImpl) GetInt(key string) int {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 0)
	return int(i)
}

func (r RequestScopeParamsImpl) GetInt8(key string) int8 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 8)
	return int8(i)
}

func (r RequestScopeParamsImpl) GetInt16(key string) int16 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 16)
	return int16(i)
}

func (r RequestScopeParamsImpl) GetInt32(key string) int32 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 32)
	return int32(i)
}

func (r RequestScopeParamsImpl) GetInt64(key string) int64 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 64)
	return i
}

func (r RequestScopeParamsImpl) GetUint(key string) uint {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 0)
	return uint(i)
}

func (r RequestScopeParamsImpl) GetUint8(key string) uint8 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 8)
	return uint8(i)
}

func (r RequestScopeParamsImpl) GetUint16(key string) uint16 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 16)
	return uint16(i)
}

func (r RequestScopeParamsImpl) GetUint32(key string) uint32 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 32)
	return uint32(i)
}

func (r RequestScopeParamsImpl) GetUint64(key string) uint64 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 64)
	return i
}

func (r RequestScopeParamsImpl) GetBool(key string) bool {
	b, _ := strconv.ParseBool(r.GetString(key))
	return b
}

func (r RequestScopeParamsImpl) GetFloat32(key string) float32 {
	f, _ := strconv.ParseFloat(r.GetString(key), 32)
	return float32(f)
}

func (r RequestScopeParamsImpl) GetFloat64(key string) float64 {
	f, _ := strconv.ParseFloat(r.GetString(key), 64)
	return f
}

func (r RequestScopeParamsImpl) Set(key, value string) {
	r[key] = value
}
