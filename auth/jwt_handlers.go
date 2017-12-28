// Package auth relation
// and more
package auth

import (
	"encoding/json"
	"net/http"

	"stagezero.com/leandro/marketbin/api"
)

var maxKeys = 3

var privateKeys = []string{`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7hdoAWaR712lcIkaB60oUwIl6CalMg6nVxcYWZVK17LfEVac
hjnZvnfCveCpvMGTWH1ThUKmjdLTQKWZuMWnHjGCcKeHkt89yty1Eo5zC1A7ukVL
u3aix2v/lVv8ALo4dPjNng+13BxgF2pX/z7xhecmASqgmhwnoDnjoqU08ys9j7Wo
8/V2KVUyBXrn4ingpB6vGCCrhwgxbYT/9oqgtlLeNfnHulbsL599JRV5f/kITa/q
tv0VnUN8G7Q61zzZg5xE6u7tClULaPoT1j1gbZZSKAqVOXwkn2wW5f6Dl5QTGqqT
VkuxHlkPlq8UffkrqfT9ZdsObESnsiY7CammrQIDAQABAoIBABdAMefxHE9D1eQn
f2NAha+Vhh4lp/w7chwPJVGaQrTNwvruelqhS6JOD7Z7Ohg0zy7VDlL1L06qR/cI
NPrWUnugWhymP5cYNfCZnRUy1AlGzI5kNgEYlMzkvxDW6sUnalwB8BJ/dkMIglnH
CNDkVLG+4Pc8MnLaTQRYouI/P+x/HERJqj/l1O3JI2Exd+MDS9ZEqLv7dCg7nep2
6Coy4O2G+r2sDjJ03ehHBZaeLOlF+9QWFu2gnv3h09oCub4MOIXPEabLJJ9D7V2E
xvrNOU3B520zMb/6IBX/yQcIr1GHGQw11cRjXXAEvozJSBP4mYIyy6oQ/8GnV5z5
OoUIvAECgYEA/574a/7TetLzhO2YadbA1WB3C/6c0vB+MkNRov6Ybufmg2W7yox0
W9AXL2XT9JGhYFo0B9PEZusA+fu4wHgbWcMxL3Tm5q1w04GMxeYIxc4iCs97FTZz
kkquujWCtsd/n0XVV6PeWcY0uhbeOSUiNuD1agslV552alm5isupG48CgYEA7nHI
LTcMqzFWv+FKcvw+NcL+t43T1ykbq7i6TEjIyhwv3S/4ZOsz1FPQtYWBYMsGz95J
hpQzDsGWR2lRulc2WzDBL3dwhzn9UCqKkIPgmYsqeKlqbAuEuqVihWBqRBWVTzgT
38SfMCtBMPVDOOWc0Sc55q4LFMcfJiH9qzESbAMCgYEAwkGA8DE7bX+aaE5XITd+
W6lvTsIzU2pHvNLD22Y3WTEKUJijWY3bb1p4BCESLi3twVdLaxdXjg7RMyhEgp/D
yTc4zaO9RVhRAarV3B6wVAIOhMCf/MLgmTAAEKpRp618IwHi2zNA6mBh+Xkfb7X0
hlf2qJvHyQo3WyLMidmzrakCgYBEKBfs/LhNUtwWGuK5/WoW2fcPJqYv8VssebAe
0As84lO4KNcambSF87NLv66cqUv4LPTdWA1EIYfAP9WRqw4pgMUnuT9cF/JVcOOM
rXWMzh/Ev8Bgw+Ybp9yCfW3Cqly0eTYNF1ndXe/Te7fMUq3BhzHgw0z9knFP6BVn
Uq6OWwKBgQDHvcyRKruGrIGQ6FMBMEWsqjtUHjLWDCE63ujlkufzIzWyxyFnU476
TPvC0HIZm17vHHYjLcjqprw3NupdxYCeK9p+0jBO7ygU0wHaDJBBuypwO/wXiKIf
jvegQdUcY9KK6rZibE32YbfpuaqcN873mnzsBGgOsk0TLBJNpmDOtA==
-----END RSA PRIVATE KEY-----
`}

var publicKeys = []string{`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7hdoAWaR712lcIkaB60o
UwIl6CalMg6nVxcYWZVK17LfEVachjnZvnfCveCpvMGTWH1ThUKmjdLTQKWZuMWn
HjGCcKeHkt89yty1Eo5zC1A7ukVLu3aix2v/lVv8ALo4dPjNng+13BxgF2pX/z7x
hecmASqgmhwnoDnjoqU08ys9j7Wo8/V2KVUyBXrn4ingpB6vGCCrhwgxbYT/9oqg
tlLeNfnHulbsL599JRV5f/kITa/qtv0VnUN8G7Q61zzZg5xE6u7tClULaPoT1j1g
bZZSKAqVOXwkn2wW5f6Dl5QTGqqTVkuxHlkPlq8UffkrqfT9ZdsObESnsiY7Camm
rQIDAQAB
-----END PUBLIC KEY-----
`}

type keyPair struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// JWTKeys handler.
func JWTKeys(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	var data keyPair
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		api.ResponseError(w, api.JSONValidationError)
		return
	}

	// logger.Info("PrivateKey: %s", data.PrivateKey)
	// logger.Info("PublicKey: %s", data.PublicKey)

	AddKeys(data.PrivateKey, data.PublicKey)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":      "added",
			"private_key": data.PrivateKey,
			"public_key":  data.PublicKey,
		},
	)
}

// AddKeys adds a keys pair to JWT
func AddKeys(private string, public string) {
	AddPrivateKey(private)
	AddPublicKey(public)
}

// AddPrivateKey adds a key to the beginign of the array
func AddPrivateKey(key string) {
	found := false
	for _, v := range privateKeys {
		if v == key {
			found = true
			break
		}
	}

	if !found {
		if len(privateKeys) == 0 {
			privateKeys = append(privateKeys, key)
		} else if len(privateKeys) > 0 && len(privateKeys) < maxKeys {
			tmp := privateKeys
			privateKeys = []string{key}
			privateKeys = append(privateKeys, tmp...)

		} else {
			tmp := privateKeys[:maxKeys-1]
			privateKeys = []string{key}
			privateKeys = append(privateKeys, tmp...)
		}
	}

	// logger.Info("SignKeys: %v", SignKeys)
}

// AddPublicKey adds a key to the beginign of the array
func AddPublicKey(key string) {
	found := false
	for _, v := range publicKeys {
		if v == key {
			found = true
			break
		}
	}

	if !found {
		if len(publicKeys) == 0 {
			publicKeys = append(publicKeys, key)
		} else if len(publicKeys) > 0 && len(publicKeys) < maxKeys {
			tmp := publicKeys
			publicKeys = []string{key}
			publicKeys = append(publicKeys, tmp...)

		} else {
			tmp := publicKeys[:maxKeys-1]
			publicKeys = []string{key}
			publicKeys = append(publicKeys, tmp...)
		}
	}

	// logger.Info("VerifyKeys: %v", VerifyKeys)
}
