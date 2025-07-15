package appjwt


type JwtPlayload struct{
    username string
    email    string 
}


func (J *JwtPlayload)GenerateToken(username string,email string){
	

}