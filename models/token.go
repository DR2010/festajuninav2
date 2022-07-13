// Package models is a dish for packages
// -------------------------------------
// .../restauranteapi/models/dishes.go
// -------------------------------------
package models

//  TokenStruct is to be exported

type TokenStruct struct {
	AccessToken           string `json:"AccessToken"`
	RefreshToken          string `json:"RefreshToken"`
	AccessTokenExpiresAt  string `json:"AccessTokenExpiresAt"`
	RefreshTokenExpiresAt string `json:"RefreshTokenExpiresAt"`
	AccessTokenExpired    bool   `json:"AccessTokenExpired"`
	RefreshTokenExpired   bool   `json:"RefreshTokenExpired"`
}

// type TokenStruct struct {
// 	AccessToken          string
// 	RefreshToken         string
// 	AccessTokenExpiresAt string
// 	AccessTokenExpired   string
// 	RefreshTokenExpired  string
// }
