/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotas{}).Type1()):             EnvironmentSpecQuotasCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDdus{}).Type1()):         EnvironmentSpecQuotasDdusCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDemUnits{}).Type1()):     EnvironmentSpecQuotasDemUnitsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasLogs{}).Type1()):         EnvironmentSpecQuotasLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasSynthetic{}).Type1()):    EnvironmentSpecQuotasSyntheticCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasUserSessions{}).Type1()): EnvironmentSpecQuotasUserSessionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorage{}).Type1()):            EnvironmentSpecStorageCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageLimits{}).Type1()):      EnvironmentSpecStorageLimitsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageRetention{}).Type1()):   EnvironmentSpecStorageRetentionCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotas{}).Type1()):             EnvironmentSpecQuotasCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDdus{}).Type1()):         EnvironmentSpecQuotasDdusCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDemUnits{}).Type1()):     EnvironmentSpecQuotasDemUnitsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasLogs{}).Type1()):         EnvironmentSpecQuotasLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasSynthetic{}).Type1()):    EnvironmentSpecQuotasSyntheticCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasUserSessions{}).Type1()): EnvironmentSpecQuotasUserSessionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorage{}).Type1()):            EnvironmentSpecStorageCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageLimits{}).Type1()):      EnvironmentSpecStorageLimitsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageRetention{}).Type1()):   EnvironmentSpecStorageRetentionCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasCodec struct {
}

func (EnvironmentSpecQuotasCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotas)(ptr) == nil
}

func (EnvironmentSpecQuotasCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotas)(ptr)
	var objs []EnvironmentSpecQuotas
	if obj != nil {
		objs = []EnvironmentSpecQuotas{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotas{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotas)(ptr) = EnvironmentSpecQuotas{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotas

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotas{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotas)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotas)(ptr) = EnvironmentSpecQuotas{}
			}
		} else {
			*(*EnvironmentSpecQuotas)(ptr) = EnvironmentSpecQuotas{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotas

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotas{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotas)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotas)(ptr) = EnvironmentSpecQuotas{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotas", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasDdusCodec struct {
}

func (EnvironmentSpecQuotasDdusCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotasDdus)(ptr) == nil
}

func (EnvironmentSpecQuotasDdusCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotasDdus)(ptr)
	var objs []EnvironmentSpecQuotasDdus
	if obj != nil {
		objs = []EnvironmentSpecQuotasDdus{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDdus{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasDdusCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotasDdus)(ptr) = EnvironmentSpecQuotasDdus{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotasDdus

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDdus{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotasDdus)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotasDdus)(ptr) = EnvironmentSpecQuotasDdus{}
			}
		} else {
			*(*EnvironmentSpecQuotasDdus)(ptr) = EnvironmentSpecQuotasDdus{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotasDdus

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDdus{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotasDdus)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotasDdus)(ptr) = EnvironmentSpecQuotasDdus{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotasDdus", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasDemUnitsCodec struct {
}

func (EnvironmentSpecQuotasDemUnitsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotasDemUnits)(ptr) == nil
}

func (EnvironmentSpecQuotasDemUnitsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotasDemUnits)(ptr)
	var objs []EnvironmentSpecQuotasDemUnits
	if obj != nil {
		objs = []EnvironmentSpecQuotasDemUnits{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDemUnits{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasDemUnitsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotasDemUnits)(ptr) = EnvironmentSpecQuotasDemUnits{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotasDemUnits

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDemUnits{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotasDemUnits)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotasDemUnits)(ptr) = EnvironmentSpecQuotasDemUnits{}
			}
		} else {
			*(*EnvironmentSpecQuotasDemUnits)(ptr) = EnvironmentSpecQuotasDemUnits{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotasDemUnits

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasDemUnits{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotasDemUnits)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotasDemUnits)(ptr) = EnvironmentSpecQuotasDemUnits{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotasDemUnits", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasLogsCodec struct {
}

func (EnvironmentSpecQuotasLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotasLogs)(ptr) == nil
}

func (EnvironmentSpecQuotasLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotasLogs)(ptr)
	var objs []EnvironmentSpecQuotasLogs
	if obj != nil {
		objs = []EnvironmentSpecQuotasLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotasLogs)(ptr) = EnvironmentSpecQuotasLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotasLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotasLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotasLogs)(ptr) = EnvironmentSpecQuotasLogs{}
			}
		} else {
			*(*EnvironmentSpecQuotasLogs)(ptr) = EnvironmentSpecQuotasLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotasLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotasLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotasLogs)(ptr) = EnvironmentSpecQuotasLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotasLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasSyntheticCodec struct {
}

func (EnvironmentSpecQuotasSyntheticCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotasSynthetic)(ptr) == nil
}

func (EnvironmentSpecQuotasSyntheticCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotasSynthetic)(ptr)
	var objs []EnvironmentSpecQuotasSynthetic
	if obj != nil {
		objs = []EnvironmentSpecQuotasSynthetic{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasSynthetic{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasSyntheticCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotasSynthetic)(ptr) = EnvironmentSpecQuotasSynthetic{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotasSynthetic

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasSynthetic{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotasSynthetic)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotasSynthetic)(ptr) = EnvironmentSpecQuotasSynthetic{}
			}
		} else {
			*(*EnvironmentSpecQuotasSynthetic)(ptr) = EnvironmentSpecQuotasSynthetic{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotasSynthetic

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasSynthetic{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotasSynthetic)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotasSynthetic)(ptr) = EnvironmentSpecQuotasSynthetic{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotasSynthetic", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecQuotasUserSessionsCodec struct {
}

func (EnvironmentSpecQuotasUserSessionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecQuotasUserSessions)(ptr) == nil
}

func (EnvironmentSpecQuotasUserSessionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecQuotasUserSessions)(ptr)
	var objs []EnvironmentSpecQuotasUserSessions
	if obj != nil {
		objs = []EnvironmentSpecQuotasUserSessions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasUserSessions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecQuotasUserSessionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecQuotasUserSessions)(ptr) = EnvironmentSpecQuotasUserSessions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecQuotasUserSessions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasUserSessions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecQuotasUserSessions)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecQuotasUserSessions)(ptr) = EnvironmentSpecQuotasUserSessions{}
			}
		} else {
			*(*EnvironmentSpecQuotasUserSessions)(ptr) = EnvironmentSpecQuotasUserSessions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecQuotasUserSessions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecQuotasUserSessions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecQuotasUserSessions)(ptr) = obj
		} else {
			*(*EnvironmentSpecQuotasUserSessions)(ptr) = EnvironmentSpecQuotasUserSessions{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecQuotasUserSessions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecStorageCodec struct {
}

func (EnvironmentSpecStorageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecStorage)(ptr) == nil
}

func (EnvironmentSpecStorageCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecStorage)(ptr)
	var objs []EnvironmentSpecStorage
	if obj != nil {
		objs = []EnvironmentSpecStorage{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorage{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecStorageCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecStorage)(ptr) = EnvironmentSpecStorage{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecStorage

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorage{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecStorage)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecStorage)(ptr) = EnvironmentSpecStorage{}
			}
		} else {
			*(*EnvironmentSpecStorage)(ptr) = EnvironmentSpecStorage{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecStorage

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorage{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecStorage)(ptr) = obj
		} else {
			*(*EnvironmentSpecStorage)(ptr) = EnvironmentSpecStorage{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecStorage", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecStorageLimitsCodec struct {
}

func (EnvironmentSpecStorageLimitsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecStorageLimits)(ptr) == nil
}

func (EnvironmentSpecStorageLimitsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecStorageLimits)(ptr)
	var objs []EnvironmentSpecStorageLimits
	if obj != nil {
		objs = []EnvironmentSpecStorageLimits{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageLimits{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecStorageLimitsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecStorageLimits)(ptr) = EnvironmentSpecStorageLimits{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecStorageLimits

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageLimits{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecStorageLimits)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecStorageLimits)(ptr) = EnvironmentSpecStorageLimits{}
			}
		} else {
			*(*EnvironmentSpecStorageLimits)(ptr) = EnvironmentSpecStorageLimits{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecStorageLimits

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageLimits{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecStorageLimits)(ptr) = obj
		} else {
			*(*EnvironmentSpecStorageLimits)(ptr) = EnvironmentSpecStorageLimits{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecStorageLimits", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecStorageRetentionCodec struct {
}

func (EnvironmentSpecStorageRetentionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecStorageRetention)(ptr) == nil
}

func (EnvironmentSpecStorageRetentionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecStorageRetention)(ptr)
	var objs []EnvironmentSpecStorageRetention
	if obj != nil {
		objs = []EnvironmentSpecStorageRetention{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageRetention{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecStorageRetentionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecStorageRetention)(ptr) = EnvironmentSpecStorageRetention{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecStorageRetention

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageRetention{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecStorageRetention)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecStorageRetention)(ptr) = EnvironmentSpecStorageRetention{}
			}
		} else {
			*(*EnvironmentSpecStorageRetention)(ptr) = EnvironmentSpecStorageRetention{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecStorageRetention

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecStorageRetention{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecStorageRetention)(ptr) = obj
		} else {
			*(*EnvironmentSpecStorageRetention)(ptr) = EnvironmentSpecStorageRetention{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecStorageRetention", "unexpected JSON type")
	}
}
