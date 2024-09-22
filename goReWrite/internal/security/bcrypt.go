package security


import "golang.org/x/crypto/bcrypt"



// HashPassword hashes the user's password

func HashPassword(password string) (string, error) {

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    if err != nil {

        return "", err

    }

    return string(hashedPassword), nil

}


func ComparePasswords(hashedPassword, password string) error {

    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), 
	[]byte(password))
}