package sound

import (
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"	
	"github.com/gopxl/beep/v2/speaker"	
	"os"
	"time"
	"log"
)

//sound Types
var ()

func playSound(soundNo int) {
	var sound string
	switch soundNo {
		case 0: // Start 
			sound = "START.mp3"
			break 
		case 1: // End 
			sound = "STOP.mp3"
		case 2: // Error 
			sound = "ERROR.mp3"
		case 3: // GCM (Go Compare) 
			sound = "GCM.mp3" 
		default: 
			sound = "WAITING.mp3"
	}
	
	f, err := os.Open(sound)
	if err != nil {
		log.Fatal(err)
	}
	
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	
	<-done		
}

func sound_main(soundNo int) {
	playSound(soundNo)
}



// sound mapping

// interface for sounds

// Part of the program calls playSound for a different parameter, perhaps 0 if error, 1 if start, 2 if GCP...
	// Can introduce a list in a const later.
// Then check which option (switch case)
// Directly play the corresponding sound to the speaker and then end the function as nothing else needs to happen. 
