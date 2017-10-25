package auth

var private = []string{`
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAwLWa6VkfyIjZPS22C3NxzSqP9lZNrUq0HYKU2skrPd+nKGne
W2l9G2zP4nqY25PesuoF+nhczl3uFiSxnYWE5M6kbPMoOAeDV2JLMfRAC/lE5mE9
f4yFGGeF7usaOa2+DbzwYCw2KL9SlG4+yszviXkoOlzaAamjpl9YAKsWVT+YypHl
UddkwamMsnggPCwpC04op8gcgIKGnOfxpldhfH2LzyPA0hrUJfKDB9QZi18Cterq
/ySkBsPAu8dddaoJRLrVlpRF9r6DyYBXV4pf7vhDH1jZjAkyc9BmLoEdgJbdbjxo
DsoMTOHfg9U5DZZocSw0UeiSJOlwnE/88VuJwwIDAQABAoIBAQCPgO0Z5Z9wcs/q
6fZNcquFD7PiygPQDvgvnYUBo5qK2didlLDbu6NJX+8yhP79TJEjBHhtO2JI3tOk
M7D8T2hWBreU1kHV72pUEUNTsMJc6EPxlun36IlYUgb/kp2V5BbKHi+WPnYaQ9bX
53zlAlVCNzHIvEovwMa35voejUj1MzfeOOcsRU9reeXzPvnlOTUNZMWqa5ob68ZF
8kv48MCtN57qoUZyb9l4cfXuTRF+2THNeK73VuoMmc+bhO7tmZOsLc6rLQ0aEk9q
wBMWqq0Q6WgCzp95dA+ZhFWSSap8497h5v15VBXdrJLYrXn/qtzRPa2ICPGo5E2W
PHjmFwoZAoGBAP485i9vEt4O7U7s/xC64xxR2M7QreDcQOPO9ErzjkRXDY2XAFP1
WCyrsRZm1JfO1PsgfMtUT0f5HNKDMAFHjZIWxGWxzgzgyUnYtuZCQW+4C5BFjBT7
dwj3oIogFMG+vfZUpNHJKPcAdphiBQC3ElzoetfPPLe05LPs7oqiIVj9AoGBAMIL
iMyRGMLkb32xeM2DvAgGyM7Lb1WgH5yzOlCV8QzqpLnk8UbwNZyPWdyMFX34eLc+
giTO0EDrpU2m8z1Ch5LXPCbRTIeDBD9CX8EwktruHmVJ3PRHMr6j3r8qaZWuZRDY
cVWSQ6MupJE9hNEKj3+1Dq0Zisq/uFPSuEVIokm/AoGBAPwrk3WBConO7HrYf7Ys
aI/ybsXUHmNmk8Zhw9WD9py6a+sA14ZvV+IW+jNqE3vv3zinZKCZI3oUEQ6MqNTc
EAPTKUJlNid33q+skN2a4iTZvD6BfQxi0BLI6yeV4oC5nNnz4vdiO4ujnf5PWv72
lvQoc5ATMfpVJnOAkqpXXhRRAoGAOigu1fDe1PqWF7vrEt1aq6Us5h2+vpEBKHvn
DGQEHPTubfCaB8LSrpugOSObBWhE9da8Nr/tVqfJoV5aJJAeBfqQQqoUH8E6sqL4
A/TE7uzTG1Rp7qSwJscCaZUSlBPyonvca+MsdmnyVL11YxmhLItdXK/9Ewsm+ah8
JffA/A0CgYEAy66bD8/8ZoB9brFyWYxmE3G3Sj6oNJM/1MDLbg3/Hr+9QgHihniJ
FWZp/5Uiur1Xl8pRPy/8pIh1dvkUHcOjAnhpMWWKKyPLOFEAQV9w5WIL2XPCZU5m
Kns6ZaJjZJNCwcskNVZh1zyqtwKOnDqfJgPR/yYorODlRbpYm/DuHeI=
-----END RSA PRIVATE KEY-----
`}

var public = []string{`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwLWa6VkfyIjZPS22C3Nx
zSqP9lZNrUq0HYKU2skrPd+nKGneW2l9G2zP4nqY25PesuoF+nhczl3uFiSxnYWE
5M6kbPMoOAeDV2JLMfRAC/lE5mE9f4yFGGeF7usaOa2+DbzwYCw2KL9SlG4+yszv
iXkoOlzaAamjpl9YAKsWVT+YypHlUddkwamMsnggPCwpC04op8gcgIKGnOfxpldh
fH2LzyPA0hrUJfKDB9QZi18Cterq/ySkBsPAu8dddaoJRLrVlpRF9r6DyYBXV4pf
7vhDH1jZjAkyc9BmLoEdgJbdbjxoDsoMTOHfg9U5DZZocSw0UeiSJOlwnE/88VuJ
wwIDAQAB
-----END PUBLIC KEY-----
`}

// // ParseJWTCookieToken returns a token if found
// func ParseJWTCookieToken(r *http.Request) *jwt.Token {
// 	c, _ := r.Cookie("token")
// 	if c != nil {
// 		token, terr := jwt.Parse(c.Value, keyLookupFunc)
// 		if terr == nil {
// 			return token
// 		}
// 	}
// 	return nil
// }

// func keyLookupFunc(token *jwt.Token) (interface{}, error) {
// 	// Don't forget to validate the alg is what you expect:
// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}
// 	// Unpack key from PEM encoded PKCS8
// 	return []byte(VerifyKeys[0]), nil
// }

// CreateCustomJWTToken returns a token with map
// func CreateCustomJWTToken(m map[string]interface{}, timeout time.Duration) string {
// 	var id, issuer, subject, audience string
// 	var notbefore int64
// 	logger.Info("JWT CustomMap: %v", m)
// 	if i := m["issuer"]; i != nil {
// 		delete(m, "issuer")
// 		issuer = i.(string)
// 	}
// 	if s := m["subject"]; s != nil {
// 		delete(m, "subject")
// 		subject = s.(string)
// 	}
// 	if a := m["audience"]; a != nil {
// 		delete(m, "audience")
// 		audience = a.(string)
// 	}
// 	if n := m["notbefore"]; n != nil {
// 		delete(m, "notbefore")
// 		notbefore = n.(int64)
// 	}
// 	if d := m["id"]; d != nil {
// 		delete(m, "id")
// 		id = d.(string)
// 	}
// 	// j := jwt.New(jwt.SigningMethodHS256)
// 	// j.Claims = make(map[string]interface{})
// 	// j.Claims["exp"] = time.Now().Add(time.Minute * timeout).UTC().Unix()
// 	// j.Claims["iat"] = time.Now().UTC().Unix()

// 	claims := jwtCustomClaims{
// 		m,
// 		jwt.StandardClaims{
// 			Id:        id,
// 			ExpiresAt: time.Now().Add(time.Second * timeout).UTC().Unix(),
// 			IssuedAt:  time.Now().UTC().Unix(),
// 			Issuer:    issuer,
// 			Subject:   subject,
// 			NotBefore: notbefore,
// 			Audience:  audience,
// 		},
// 	}

// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	logger.Info("Create Token: %v", t)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return t
// }
