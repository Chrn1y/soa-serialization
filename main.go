package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack"

	"github.com/Chrn1y/soa-serialization/protos/models"
)

func writeToFile(name string, bytes []byte) error {
	f, _ := os.Create(name)
	defer f.Close()

	if _, err := f.Write(bytes); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func readFromFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type AdditionalStruct struct {
	Str  string
	Strs []string
}

type BasicStruct struct {
	Name       string
	Id         int32
	ServiceIds []uint32
	Adds       []*AdditionalStruct
	Dict       map[string]float32
}

const num = 1000

func main() {
	basic := getBasicStruct()
	println("Serialization and Deserialization in Go:\n")

	// gob
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			if err := enc.Encode(basic); err != nil {
				println(err.Error())
				return
			}
		}

		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(basic); err != nil {
			println(err.Error())
			return
		}

		fmt.Printf("encoding/gob serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(buf.Bytes()),
			(time.Now().UnixNano()-start)/(num+1))

		err := writeToFile("out/gob.txt", buf.Bytes())
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &BasicStruct{}
			var buf bytes.Buffer
			enc := gob.NewDecoder(&buf)
			data, err := readFromFile("out/gob.txt")
			if err != nil {
				println(err.Error())
				return
			}
			_, err = buf.Write(data)
			if err != nil {
				println(err.Error())
				return
			}
			if err = enc.Decode(out); err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("encoding/gob deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)

	}

	// json
	{
		out := make([]byte, 0)
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var err error
			out, err = json.Marshal(basic)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("json serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(out),
			(time.Now().UnixNano()-start)/num)

		err := writeToFile("out/json.txt", out)
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &BasicStruct{}
			data, err := readFromFile("out/json.txt")
			if err != nil {
				println(err.Error())
				return
			}
			err = json.Unmarshal(data, out)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("json deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)
	}

	// xml
	{
		basic := getBasicStruct().toXML()
		out := make([]byte, 0)
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var err error
			out, err = xml.Marshal(basic)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("xml serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(out),
			(time.Now().UnixNano()-start)/num)

		err := writeToFile("out/xml.txt", out)
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &BasicXMLStruct{}
			data, err := readFromFile("out/xml.txt")
			if err != nil {
				println(err.Error())
				return
			}
			err = xml.Unmarshal(data, out)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("xml deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)
	}

	//proto
	{
		basic := getBasicStruct().toProto()
		out := make([]byte, 0)
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var err error
			out, err = proto.Marshal(basic)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("proto serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(out),
			(time.Now().UnixNano()-start)/num)
		err := writeToFile("out/proto.txt", out)
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &models.Basic{}
			data, err := readFromFile("out/proto.txt")
			if err != nil {
				println(err.Error())
				return
			}
			err = proto.Unmarshal(data, out)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("proto deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)
	}

	//yaml
	{
		out := make([]byte, 0)
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var err error
			out, err = yaml.Marshal(basic)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("yaml serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(out),
			(time.Now().UnixNano()-start)/num)

		err := writeToFile("out/yaml.txt", out)
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &BasicStruct{}
			data, err := readFromFile("out/yaml.txt")
			if err != nil {
				println(err.Error())
				return
			}
			err = yaml.Unmarshal(data, out)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("yaml deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)
	}

	//messagepack
	{
		out := make([]byte, 0)
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			var err error
			out, err = msgpack.Marshal(basic)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("messagepack serialization\nsize: %d bytes\ntime: %d nanoseconds\n\n",
			len(out),
			(time.Now().UnixNano()-start)/num)

		err := writeToFile("out/msgpack.txt", out)
		if err != nil {
			println(err.Error())
			return
		}
	}
	{
		start := time.Now().UnixNano()
		for i := 0; i < num; i++ {
			out := &BasicStruct{}
			data, err := readFromFile("out/msgpack.txt")
			if err != nil {
				println(err.Error())
				return
			}
			err = msgpack.Unmarshal(data, out)
			if err != nil {
				println(err.Error())
				return
			}
		}
		fmt.Printf("messagepack deserialization\ntime: %d nanoseconds\n\n",
			(time.Now().UnixNano()-start)/num)
	}
}

func getBasicStruct() *BasicStruct {
	servIds := make([]uint32, 0)
	for i := range make([]struct{}, 1000) {
		servIds = append(servIds, uint32(i))
	}
	floats := make(map[string]float32)
	for i := range make([]struct{}, 1000) {
		floats[fmt.Sprintf("float%d", i)] = float32(i) + float32(i)/100
	}
	adds := make([]*AdditionalStruct, 0)
	for i := range make([]struct{}, 100) {
		temp := &AdditionalStruct{
			Str:  fmt.Sprintf("some_string%d", i),
			Strs: make([]string, 0),
		}
		for j := range make([]struct{}, i) {
			temp.Strs = append(temp.Strs, fmt.Sprintf("str%d", j))
		}
		adds = append(adds, temp)
	}
	return &BasicStruct{
		Name:       "some_name",
		Id:         12345,
		ServiceIds: servIds,
		Adds:       adds,
		Dict:       floats,
	}
}

func (b *BasicStruct) toProto() *models.Basic {
	temp := make([]*models.Basic_Additional, 0)
	for _, a := range b.Adds {
		temp = append(temp, &models.Basic_Additional{
			Str:  a.Str,
			Strs: a.Strs,
		})
	}
	return &models.Basic{
		Name:       b.Name,
		Id:         b.Id,
		ServiceIds: b.ServiceIds,
		Additional: temp,
		Dict:       b.Dict,
	}
}

func (b *BasicStruct) toXML() *BasicXMLStruct {
	values := make([]float32, 0)
	keys := make([]string, 0)
	for k, v := range b.Dict {
		values = append(values, v)
		keys = append(keys, k)
	}
	return &BasicXMLStruct{
		Name:       b.Name,
		Id:         b.Id,
		ServiceIds: b.ServiceIds,
		Adds:       b.Adds,
		DictValues: values,
		DictKeys:   keys,
	}
}

type BasicXMLStruct struct {
	Name       string
	Id         int32
	ServiceIds []uint32
	Adds       []*AdditionalStruct
	DictKeys   []string
	DictValues []float32
}
