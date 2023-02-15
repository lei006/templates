package webrtc_camera

import (
	"context"
	"errors"
	"math/rand"
	"server/utils/packet"

	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pion/webrtc/v2"
	"github.com/pion/webrtc/v2/pkg/media"
	//"github.com/pion/webrtc/v2/examples/internal/signal"
)

const (
	rtcpPLIInterval = time.Second * 3
)

type WebRtcPion struct {
	PeerConn *webrtc.PeerConnection

	//localTrackChan chan *webrtc.TrackLocalStaticRTP

	videoTrack *webrtc.Track

	Answer_offer SessionDescriptionEx

	firstFrameTimestamp uint64 //第一帧时间戳
	lastFrameTimestamp  uint64 //上一帧时间戳

}

func CreateWebRtcPeerConn(request_offer SessionDescriptionEx) (*WebRtcPion, error) {
	capture := WebRtcPion{}

	//
	answer_offer, err := capture.InitWebrtcPeerConn(request_offer.SessionDescription)
	if err != nil {
		return nil, errors.New("初始化WebRTC的Peer连接失败")
	}

	//err = server_rtsp.SubPacket("/capture.sdp", capture.SendPacket)

	capture.Answer_offer.SessionDescription = *answer_offer

	return &capture, nil
}

func (cap *WebRtcPion) SendPacket(packet *packet.H264Packet) {

	// 计算采样
	var time_len uint32

	curTimestamp := uint64(time.Now().UnixNano() / 1e6)

	if cap.firstFrameTimestamp == 0 {
		// 如果是第一帧
		time_len = 10
		cap.firstFrameTimestamp = curTimestamp
		cap.lastFrameTimestamp = curTimestamp
	} else {
		time_len = uint32(curTimestamp - cap.lastFrameTimestamp)
		cap.lastFrameTimestamp = curTimestamp
	}

	cap.videoTrack.WriteSample(media.Sample{Data: packet.GetData(), Samples: time_len})

}

func (cap *WebRtcPion) InitWebrtcPeerConn(request_offer webrtc.SessionDescription) (*webrtc.SessionDescription, error) {

	mediaEngine := webrtc.MediaEngine{}
	err := mediaEngine.PopulateFromSDP(request_offer)
	if err != nil {
		return nil, errors.New("create MediaEngine.PopulateFromSDP fail:" + err.Error())
	}

	// 新的连接
	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return nil, errors.New("新的Peer连接失败!")
	}

	_, iceConnectedCtxCancel := context.WithCancel(context.Background())

	// 创建视频通道

	videoID, err := getPayloadType(mediaEngine, webrtc.RTPCodecTypeVideo, webrtc.H264)
	if err != nil {
		iceConnectedCtxCancel()
		return nil, errors.New("new Video Track fail!" + err.Error())
	}

	videoTrack, err := peerConnection.NewTrack(videoID, rand.Uint32(), "video", "LiveRTC")
	if err != nil {
		iceConnectedCtxCancel()
		return nil, errors.New("getPayloadType Video Type fail!" + err.Error())
	}

	if _, err = peerConnection.AddTrack(videoTrack); err != nil {
		iceConnectedCtxCancel()
		return nil, errors.New("加入 Video Track失败!")
	}

	cap.videoTrack = videoTrack
	// 状态改变通知
	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		logs.Debug("Connection State has changed :", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			iceConnectedCtxCancel()
		}
	})

	// Set the remote SessionDescription
	if err = peerConnection.SetRemoteDescription(request_offer); err != nil {
		logs.Debug("SetRemoteDescription error:", request_offer, err.Error())
		return nil, err
	}

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		logs.Debug("CreateAnswer error:", err.Error())
		return nil, err
	}

	// Sets the LocalDescription, and starts our UDP listeners
	if err = peerConnection.SetLocalDescription(answer); err != nil {
		logs.Debug("SetLocalDescription error:", err.Error())
		return nil, err
	}

	cap.PeerConn = peerConnection

	return &answer, nil
}

func getPayloadType(m webrtc.MediaEngine, codecType webrtc.RTPCodecType, codecName string) (uint8, error) {
	for _, codec := range m.GetCodecsByKind(codecType) {
		if codec.Name == codecName {
			return codec.PayloadType, nil
		}
	}
	return 0, errors.New("Remote peer does not support " + codecName)
}
