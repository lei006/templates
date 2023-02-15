package server_rtsp

var rtsp_server *RtspServer

func Start(port int) error {
	tmp_server := MakeServer()
	err := tmp_server.Start(port)
	if err != nil {
		return err
	}

	rtsp_server = tmp_server
	return nil
}

func Stop() {
	if rtsp_server != nil {
		rtsp_server.Stop()
		rtsp_server = nil
	}
}

func SubPacket(path string, callback PacketCallback) error {
	return rtsp_server.SubPacket(path, callback)
}
