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

func playSound(soundName string) {
	/*var sound string
	switch soundNo {
		case 0: // Start 
			sound = "sound/START.mp3"
			break 
		case 1: // End 
			sound = "sound/STOP.mp3"
		case 2: // Error 
			sound = "sound/ERROR.mp3"
		case 3: // GCM (Go Compare) 
			sound = "sound/GCM.mp3" 
		default: 
			sound = "sound/WAITING.mp3"
	}*/
	
	f, err := os.Open("sound/" + soundName + ".mp3")
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

func Sound_main(soundName string) {
	playSound(soundName)
}



// sound mapping

// interface for sounds

// Part of the program calls playSound for a different parameter, perhaps 0 if error, 1 if start, 2 if GCP...
	// Can introduce a list in a const later.
// Then check which option (switch case)
// Directly play the corresponding sound to the speaker and then end the function as nothing else needs to happen. 
