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
		jsoniter.MustGetKind(reflect2.TypeOf(CaptureRuleSpecMatches{}).Type1()):        CaptureRuleSpecMatchesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ContextPropagationSpecMatches{}).Type1()): ContextPropagationSpecMatchesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EntryPointSpecMatches{}).Type1()):         EntryPointSpecMatchesCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CaptureRuleSpecMatches{}).Type1()):        CaptureRuleSpecMatchesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ContextPropagationSpecMatches{}).Type1()): ContextPropagationSpecMatchesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EntryPointSpecMatches{}).Type1()):         EntryPointSpecMatchesCodec{},
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
type CaptureRuleSpecMatchesCodec struct {
}

func (CaptureRuleSpecMatchesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CaptureRuleSpecMatches)(ptr) == nil
}

func (CaptureRuleSpecMatchesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CaptureRuleSpecMatches)(ptr)
	var objs []CaptureRuleSpecMatches
	if obj != nil {
		objs = []CaptureRuleSpecMatches{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CaptureRuleSpecMatches{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CaptureRuleSpecMatchesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CaptureRuleSpecMatches)(ptr) = CaptureRuleSpecMatches{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CaptureRuleSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CaptureRuleSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CaptureRuleSpecMatches)(ptr) = objs[0]
			} else {
				*(*CaptureRuleSpecMatches)(ptr) = CaptureRuleSpecMatches{}
			}
		} else {
			*(*CaptureRuleSpecMatches)(ptr) = CaptureRuleSpecMatches{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CaptureRuleSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CaptureRuleSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CaptureRuleSpecMatches)(ptr) = obj
		} else {
			*(*CaptureRuleSpecMatches)(ptr) = CaptureRuleSpecMatches{}
		}
	default:
		iter.ReportError("decode CaptureRuleSpecMatches", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ContextPropagationSpecMatchesCodec struct {
}

func (ContextPropagationSpecMatchesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ContextPropagationSpecMatches)(ptr) == nil
}

func (ContextPropagationSpecMatchesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ContextPropagationSpecMatches)(ptr)
	var objs []ContextPropagationSpecMatches
	if obj != nil {
		objs = []ContextPropagationSpecMatches{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ContextPropagationSpecMatches{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ContextPropagationSpecMatchesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ContextPropagationSpecMatches)(ptr) = ContextPropagationSpecMatches{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ContextPropagationSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ContextPropagationSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ContextPropagationSpecMatches)(ptr) = objs[0]
			} else {
				*(*ContextPropagationSpecMatches)(ptr) = ContextPropagationSpecMatches{}
			}
		} else {
			*(*ContextPropagationSpecMatches)(ptr) = ContextPropagationSpecMatches{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ContextPropagationSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ContextPropagationSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ContextPropagationSpecMatches)(ptr) = obj
		} else {
			*(*ContextPropagationSpecMatches)(ptr) = ContextPropagationSpecMatches{}
		}
	default:
		iter.ReportError("decode ContextPropagationSpecMatches", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EntryPointSpecMatchesCodec struct {
}

func (EntryPointSpecMatchesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EntryPointSpecMatches)(ptr) == nil
}

func (EntryPointSpecMatchesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EntryPointSpecMatches)(ptr)
	var objs []EntryPointSpecMatches
	if obj != nil {
		objs = []EntryPointSpecMatches{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EntryPointSpecMatches{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EntryPointSpecMatchesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EntryPointSpecMatches)(ptr) = EntryPointSpecMatches{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EntryPointSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EntryPointSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EntryPointSpecMatches)(ptr) = objs[0]
			} else {
				*(*EntryPointSpecMatches)(ptr) = EntryPointSpecMatches{}
			}
		} else {
			*(*EntryPointSpecMatches)(ptr) = EntryPointSpecMatches{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EntryPointSpecMatches

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EntryPointSpecMatches{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EntryPointSpecMatches)(ptr) = obj
		} else {
			*(*EntryPointSpecMatches)(ptr) = EntryPointSpecMatches{}
		}
	default:
		iter.ReportError("decode EntryPointSpecMatches", "unexpected JSON type")
	}
}
