package util 


import (
	"github.com/joho/godotenv"
	"log"

)

// loads the enviroment vars and returns a slice of keys
//0=openai; 
// 1=tmdb; 
// 2=tursokey; 
// 3=tursodb;
//4=HBO URL
// 5=NETFLIX URL
// 6=Disney URL
// 7=AppleTV URL
// 8=PRIME URL
func LoadEnviroment() []string {
	myEnv, err := godotenv.Read()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	myKeys := []string{}
	openaiKey := myEnv["OPENAI_API_KEY"];
	tmdbKey := myEnv["TMDB_API_KEY"];
	tursoKey := myEnv["TURSO_AUTH_KEY"]
	tursoDB := myEnv["TURSO_DATABASE_LINK"]
	HBO := myEnv["HBO"]
	NETFLIX := myEnv["NETFLIX"]
	Disney := myEnv["DISNEY"]
	AppleTV := myEnv["APPLE"]
	PRIME := myEnv["PRIME"]
	
	slice := append(myKeys,openaiKey,tmdbKey,tursoKey,tursoDB, HBO, NETFLIX, Disney, AppleTV, PRIME)
  
	return slice
  }