package server_rtsp

type Receiver struct {
	//engine.PublishContext

	mid string
	url string
	//client *rtsp.Client
}

func MakeNewReceiver(mid string, url string) *Receiver {

	recevier := &Receiver{}
	recevier.mid = mid
	recevier.url = url
	//recevier.url = "rtsp://127.0.0.1/test.sdp";
	//recevier.pubsuber = core.NewPublisher("aaaaaaaa",100*time.Millisecond, 10);
	//recevier.pubsuber = core.CreateNewPublisher("", "rtsp", 100, 10)

	go recevier.Run()

	return recevier
}

func (this *Receiver) Run() (err error) {

	return nil
}
