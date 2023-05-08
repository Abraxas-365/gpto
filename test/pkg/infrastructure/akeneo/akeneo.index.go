package akeneo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/ports"
	"io/ioutil"
	"net/http"
	"time"
)

type akeneo struct {
	db       *sql.DB
	secret   string
	user     string
	password string
	url      string
}

func AkeneoFactory(db *sql.DB, secret string, user string, password string, url string) ports.Akeneo {
	return &akeneo{
		db:       db,
		secret:   secret,
		user:     user,
		password: password,
		url:      url,
	}
}

func (a *akeneo) Login() (string, error) {
	var buf bytes.Buffer
	var bearer = "Basic " + a.secret
	auth := struct {
		GrantType string `json:"grant_type"`
		User      string `json:"username"`
		Password  string `json:"password"`
	}{}
	auth.User = a.user
	auth.Password = a.password
	auth.GrantType = "password"

	akenenoResponse := struct {
		Token string `json:"access_token"`
	}{}

	if err := json.NewEncoder(&buf).Encode(auth); err != nil {
		return "", err
	}

	var client http.Client
	req, err := http.NewRequest("post", a.url+"/api/oauth/v1/token", &buf)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &akenenoResponse); err != nil {
		return "", err
	}
	return akenenoResponse.Token, nil
}

func uploadData[T any](url string, token string, data []T) error {
	var bearer = "Bearer " + token
	fmt.Println("akeneo", bearer)
	var client http.Client
	var buf bytes.Buffer

	for index := range data {
		if err := json.NewEncoder(&buf).Encode(data[index]); err != nil {
			return err
		}
		if index%100 == 0 {
			req, err := http.NewRequest("PATCH", url, &buf)
			if err != nil {
				return err
			}
			req.Header.Add("Authorization", bearer)
			req.Header.Add("Content-Type", "application/vnd.akeneo.collection+json")
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			buf.Reset()
			time.Sleep(15000 * time.Millisecond)

		}
		fmt.Println(index)
	}
	//enviar restantes
	req, err := http.NewRequest("PATCH", url, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/vnd.akeneo.collection+json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	buf.Reset()
	return nil
}

func getData[T any](url string, token string, data *T) error {

	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/vnd.akeneo.collection+json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	return nil
}
