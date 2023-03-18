package packet

import "github.com/pion/rtp"

type H264Packet struct {
	Buffer    []byte
	FrameData []byte
	Samples   uint32
	RtpPacket rtp.Packet

	PPS      []byte
	SPS      []byte
	FatherId string //父id: 就是生成这个数据的人
}

func MakeH264Packet(data []byte) *H264Packet {

	packet := &H264Packet{
		Buffer: data,
	}
	return packet
}

func (this *H264Packet) DealFrame() {

	if this.IsKeyFrame() {
		this.FrameData = append([]byte("\000\000\000\001"+string(this.SPS)+"\000\000\000\001"+string(this.PPS)+"\000\000\000\001"), this.Buffer[:]...)
	} else {
		this.FrameData = append([]byte("\000\000\000\001"), this.Buffer[:]...)
	}
}

func (this *H264Packet) IsKeyFrame() bool {
	nal_unit_type := this.Buffer[0] & 0b00011111
	if nal_unit_type == 5 {
		return true
	}
	return false
}

func (this *H264Packet) GetData() []byte {
	return this.FrameData
}

func (this *H264Packet) GetSamples() uint32 {
	return this.Samples
}
