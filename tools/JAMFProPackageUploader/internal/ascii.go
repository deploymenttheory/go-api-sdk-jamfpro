package jamfpropackageuploader

import "fmt"

// Helper function to print ASCII art
func PrintASCIILogo() {
	asciiArt := `
   ___  ___  ___  ___ _____   _____            _____          _                      _   _       _                 _            
  |_  |/ _ \ |  \/  ||  ___| | ___ \          | ___ \        | |                    | | | |     | |               | |           
    | / /_\ \| .  . || |_    | |_/ / __ ___   | |_/ /_ _  ___| | __ ___  __ _  ___  | | | |_ __ | | ___   __ _  __| | ___ _ __  
    | |  _  || |\/| ||  _|   |  __/ '__/ _ \  |  __/ _  |/ __| |/ / _  |/ _  |/ _ \ | | | | '_ \| |/ _ \ / _  |/ _` + "`" + ` |/ _ \ '__| 
/\__/ / | | || |  | || |     | |  | | | (_) | | | | (_| | (__|   < (_| | (_| |  __/ | |_| | |_)|| | (_) | (_| | (_| |  __/ |    
\____/\_| |_/\_|  |_/\_|     \_|  |_|  \___/  \_|  \__,_|\___|_|\_\__,_|\__, |\___|  \___/| .__/|_|\___/ \__,_|\__,_|\___|_|    
                                                                         __/ |            | |                                   
                                                                        |___/             |_|                                   
`
	fmt.Println(asciiArt)
}
