package rpi

const ADDR = ":8080"

//TODO: Add Cobra and start as Daemon in Background
//TODO: Add Signals for grace shutdown
//TODO: make the Port configurable also the logging files
//TODO: Startup Text where the app is available
//TODO: Implement also abstraction for Commands in order to have to interfaces API HTTP/CLI

type Color struct {
	R            int `json:"r"`
	G            int `json:"g"`
	B            int `json:"b"`
	Duration_sec int `json:"duration_sec"`
}

//func info(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte("With the POST Methode you are able to controll the lights"))
//	w.Write([]byte("you can pass the following paylog"))
//	w.Write([]byte("{\"r\": 125,\"g\": 125,\"b\": 125,\"duration_sec\": 5, }"))
//	return
//}
//
//func ligthON(w http.ResponseWriter, r *http.Request) {
//
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var c Color
//	log.Infof("Body, %v", string(reqBody))
//	err := json.Unmarshal(reqBody, &c)
//	if err != nil {
//		log.Error(err)
//	}
//	log.Infof("Struct, %v", c)
//
//	w.WriteHeader(http.StatusOK)
//
//	fmt.Fprintf(w, "your desired color r: %d, g: %d b: %d and it should shine for %d seconds", c.R, c.G, c.B, c.Duration_sec)
//	turnOn(&c)
//	return
//}
//
//func turnOn(color *Color) {
//	brightness := 0.5
//	log.Infof("Set brightness to %d", brightness)
//	blinkt := NewBlinkt(brightness)
//	log.Info("Create new Instance of Blink")
//	blinkt.SetClearOnExit(true)
//	log.Info("Setup SetClearOnExit: %b", false)
//	blinkt.Setup()
//	log.Info("Setup Blinkt")
//	blinkt.SetAll(color.R, color.G, color.B)
//	log.Info("Setup Color for RGB: %d %d %d", color.R, color.G, color.B)
//	blinkt.Show()
//	log.Info("Show Blinkt")
//	Delay((1000 * color.Duration_sec))
//	log.Info("Deployed after %d ", (1000 * color.Duration_sec))
//	blinkt.Clear()
//	log.Info("Clear Blinkt")
//	blinkt.Show()
//	return
//}
//
//func demo() {
//	brightness := 0.5
//	blinkt := NewBlinkt(brightness)
//
//	blinkt.SetClearOnExit(true)
//
//	blinkt.Setup()
//	Delay(100)
//
//	step := 0
//
//	for {
//
//		step = step % 3
//		switch step {
//		case 0:
//			blinkt.SetAll(128, 0, 0)
//		case 1:
//			blinkt.SetAll(0, 128, 0)
//		case 2:
//			blinkt.SetAll(0, 0, 128)
//		}
//
//		step++
//		blinkt.Show()
//		Delay(500)
//		fmt.Printf("Step %d\n", step)
//	}
//}
