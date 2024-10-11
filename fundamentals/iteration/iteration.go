package iteration


func ScreamGenerator(screamLetter string, repetitions int) string{
	var fullScream string
	for i := 0; i<repetitions; i++{
		fullScream += screamLetter
	}
	return fullScream	
}

func main() {

}