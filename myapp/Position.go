// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package myapp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Position struct {
	_tab flatbuffers.Struct
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Position) Latitude() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Position) MutateLatitude(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Position) Longitude() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Position) MutateLongitude(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreatePosition(builder *flatbuffers.Builder, latitude float32, longitude float32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependFloat32(longitude)
	builder.PrependFloat32(latitude)
	return builder.Offset()
}
