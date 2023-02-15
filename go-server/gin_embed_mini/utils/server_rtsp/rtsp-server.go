package server_rtsp

import (
	"errors"
	"fmt"
	"net"
	"runtime/debug"
	"server/utils/packet"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

type PacketCallback func(*packet.H264Packet)

type RtspServer struct {
	SessionLogger
	TCPListener *net.TCPListener
	TCPPort     int
	Stoped      bool
	pushers     map[string]*Pusher // Path <-> Pusher
	pushersLock sync.RWMutex
	Name        string

	handlePlayFilter func(url string) (string, error)
	handleRecvFilter func(url string) (string, error)
}

func MakeServer() *RtspServer {

	//port := appContext.Config.GetConfigInt32("RtspPort", Default_RtspPort)

	srv := &RtspServer{
		//SessionLogger:  SessionLogger{log.New(os.Stdout, "[LiveRTC-RTSP]", log.LstdFlags|log.Lshortfile)},
		SessionLogger: SessionLogger{logger: logs.NewLogger()},
		Stoped:        true,
		pushers:       make(map[string]*Pusher),
		Name:          "rtsp-server",
	}

	//log := logs.NewLogger()
	//log.SetLogger(logs.AdapterConsole)

	return srv
}

func (server *RtspServer) Start(port int) (err error) {
	server.TCPPort = port
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", server.TCPPort))
	if err != nil {
		return
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}

	go server.ListenStart(listener)

	time.Sleep(2 * time.Second)

	return err
}

func (this *RtspServer) SubPacket(path string, callback PacketCallback) error {

	pusher := this.GetPusher(path)
	if pusher == nil {
		return errors.New("no find rtsp path:" + path)
	}

	pusher.SetPacketCallback(callback)
	return nil
}

func (server *RtspServer) ListenStart(listener *net.TCPListener) (err error) {

	server.Stoped = false
	server.TCPListener = listener

	for !server.Stoped {
		conn, err := server.TCPListener.Accept()
		if err != nil {
			server.logger.Debug(err.Error())
			continue
		}

		handle := MakeRtspHandle(conn)

		go server.handleConn(handle)

	}
	return
}

func (this *RtspServer) handleConn(handle *RTSPHandle) error {
	defer func() {
		if r := recover(); r != nil {
			this.logger.Error("rtmp server handleConn panic: ", r, string(debug.Stack()))
		}
	}()

	networkBuffer := 1048576

	if tcpConn, ok := handle.conn.(*net.TCPConn); ok {
		if err := tcpConn.SetReadBuffer(networkBuffer); err != nil {
			this.logger.Debug("rtsp server conn set read buffer error, %v", err)
		}
		if err := tcpConn.SetWriteBuffer(networkBuffer); err != nil {
			this.logger.Debug("rtsp server conn set write buffer error, %v", err)
		}
	}

	session := NewSession(this, handle.conn)
	session.Start()

	return nil
}

func (server *RtspServer) Stop() {

	server.logger.Debug("rtsp server stop on", server.TCPPort)
	server.Stoped = true
	if server.TCPListener != nil {
		server.TCPListener.Close()
		server.TCPListener = nil
	}
	server.pushersLock.Lock()
	server.pushers = make(map[string]*Pusher)
	server.pushersLock.Unlock()

}

func (server *RtspServer) AddPusher(pusher *Pusher) bool {

	added := false
	server.pushersLock.Lock()
	_, ok := server.pushers[pusher.Path()]
	if !ok {
		server.pushers[pusher.Path()] = pusher
		//server.logger.Debug("pusher start:", pusher, "     now pusher size:", len(server.pushers))
		added = true
	} else {
		added = false
	}
	server.pushersLock.Unlock()
	if added {
		go pusher.Start()

	}
	return added
}

func (server *RtspServer) TryAttachToPusher(session *Session) (int, *Pusher) {
	server.pushersLock.Lock()
	attached := 0
	var pusher *Pusher = nil
	if _pusher, ok := server.pushers[session.Path]; ok {
		if _pusher.RebindSession(session) {
			server.logger.Debug("Attached to a pusher")
			attached = 1
			pusher = _pusher
		} else {
			attached = -1
		}
	}
	server.pushersLock.Unlock()
	return attached, pusher
}

func (server *RtspServer) RemovePusher(pusher *Pusher) {

	//pusher.publisher.Close("rtsp pusher close 2")
	server.pushersLock.Lock()
	if _pusher, ok := server.pushers[pusher.Path()]; ok && pusher.ID() == _pusher.ID() {
		delete(server.pushers, pusher.Path())
		server.logger.Debug("pusher end:", pusher, "     now pusher size:", len(server.pushers))
	}
	server.pushersLock.Unlock()
}

func (server *RtspServer) GetPusher(path string) (pusher *Pusher) {
	server.pushersLock.RLock()
	pusher = server.pushers[path]
	server.pushersLock.RUnlock()
	return
}

func (server *RtspServer) GetPushers() (pushers map[string]*Pusher) {
	pushers = make(map[string]*Pusher)
	server.pushersLock.RLock()
	for k, v := range server.pushers {
		pushers[k] = v
	}
	server.pushersLock.RUnlock()
	return
}

func (server *RtspServer) GetPusherSize() (size int) {
	server.pushersLock.RLock()
	size = len(server.pushers)
	server.pushersLock.RUnlock()
	return
}

func (server *RtspServer) SetRecvFilter(cb_handle func(url string) (string, error)) {

	server.handleRecvFilter = cb_handle

}

func (server *RtspServer) SetPlayFilter(cb_handle func(url string) (string, error)) {

	server.handlePlayFilter = cb_handle

}
