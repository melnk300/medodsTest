package jwt

import (
	"errors"
	"github.com/melnk300/medodsTest/pkg/database"
	"github.com/melnk300/medodsTest/pkg/mail"
	"github.com/melnk300/medodsTest/pkg/tokens"

	"log"
)

func CreateTokens(ip string, guid string) (tokens.JwtToken, tokens.JwtToken) {
	db := database.MakeConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users VALUES ($1, $2) ON CONFLICT (guid) DO NOTHING", guid, "melnk300@gmail.com")
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	access, refresh, jti := tokens.GenerateTokens(ip, guid)

	_, err = db.Exec("INSERT INTO tokens VALUES ($1, $2)", jti, guid)
	if err != nil {
		log.Fatalf("Error inserting token: %v", err)
	}

	return access, refresh
}

func ProcessTokens(acToken string, rfToken string, ip string) (*tokens.JwtToken, *tokens.JwtToken, error) {
	db := database.MakeConnection()
	defer db.Close()

	accessClaims, err := tokens.ParseToken(acToken)
	if err != nil {
		switch err.Error() {
		case "failed to cast claims to struct":
			return nil, nil, errors.New("invalid payload")
		case "token is invalid":
			return nil, nil, errors.New("invalid token")
		}
	}

	refreshClaims, err := tokens.ParseToken(rfToken)
	if err != nil {
		switch err.Error() {
		case "failed to cast claims to struct":
			return nil, nil, errors.New("invalid payload")
		case "token is invalid":
			return nil, nil, errors.New("invalid token")
		}
	}

	if refreshClaims.JTI != accessClaims.JTI {
		return nil, nil, errors.New("different tokens")
	}

	type UserData struct {
		ID    string
		Email string
	}

	userData := UserData{}

	err = db.QueryRow("SELECT tokens.user_id, users.email FROM tokens JOIN users ON users.guid = tokens.user_id WHERE jti = $1", refreshClaims.JTI).Scan(&userData.ID, &userData.Email)

	if err != nil {
		return nil, nil, errors.New("token used")
	}

	_, err = db.Exec("DELETE FROM tokens WHERE jti = $1", refreshClaims.JTI)
	if err != nil {
		log.Fatalf("Query error: %v", err.Error())
	}

	if refreshClaims.IP != ip {
		mail.SendLetter(userData.Email, "email warning")
		return nil, nil, errors.New("different ip")
	}

	newAcToken, newRfToken := CreateTokens(refreshClaims.IP, refreshClaims.GUID)

	return &newAcToken, &newRfToken, nil
}
