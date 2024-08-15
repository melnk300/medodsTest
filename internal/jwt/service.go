package jwt

import (
	"github.com/melnk300/medodsTest/pkg/database"
	"github.com/melnk300/medodsTest/pkg/tokens"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateTokens(ip string, guid string) (tokens.Token, tokens.Token) {
	db := database.MakeConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users VALUES ($1, $2) ON CONFLICT (guid) DO NOTHING", guid, "melnk300@gmail.com")
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	access, refresh, jti := tokens.GenerateTokens(ip, guid)

	hashedJti, _ := bcrypt.GenerateFromPassword([]byte(jti), bcrypt.DefaultCost)

	_, err = db.Exec("INSERT INTO tokens VALUES ($1, $2)", string(hashedJti), guid)
	if err != nil {
		log.Fatalf("Error inserting token: %v", err)
	}

	return access, refresh
}
