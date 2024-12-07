package helper

import (
	"aidanwoods.dev/go-paseto"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func GeneratePaseto(key paseto.V4SymmetricKey, claims map[string]interface{}) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetAudience(viper.GetString("PASETO_AUDIENCE"))
	token.SetIssuer(viper.GetString("PASETO_ISSUER"))
	token.SetSubject(viper.GetString("PASETO_SUBJECT"))
	token.SetExpiration(time.Now().Add(24 * time.Hour))

	token.SetString("sec", viper.GetString("PASETO_SECRET"))
	for k, v := range claims {
		err := token.Set(k, v)
		if err != nil {
			return "", err
		}
	}

	return token.V4Encrypt(key, nil), nil
}

func ValidatePaseto(key paseto.V4SymmetricKey, token string) (map[string]interface{}, error) {
	decrypted := paseto.NewParser()
	v4Local, err := decrypted.ParseV4Local(key, token, nil)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}, err
	}

	issuer, err := v4Local.GetIssuer()
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}, err
	}

	audience, err := v4Local.GetAudience()
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}, err
	}

	subject, err := v4Local.GetSubject()
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}, err
	}

	if issuer != viper.GetString("PASETO_ISSUER") {
		fmt.Println("Invalid Issuer")
		return map[string]interface{}{}, errors.New("error validate")
	}

	if audience != viper.GetString("PASETO_AUDIENCE") {
		fmt.Println("Invalid Audience")
		return map[string]interface{}{}, errors.New("error validate")
	}

	if subject != viper.GetString("PASETO_SUBJECT") {
		fmt.Println("Invalid Subject")
		return map[string]interface{}{}, errors.New("error validate")
	}

	secret, err := v4Local.GetString("sec")
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}, errors.New("error validate")
	}

	if secret != viper.GetString("PASETO_SECRET") {
		fmt.Println("Invalid Secret")
		return map[string]interface{}{}, errors.New("error validate")
	}

	return v4Local.Claims(), nil
}
