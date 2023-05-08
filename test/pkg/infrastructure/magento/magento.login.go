package magento

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (m *magnetoRepo) Login() (string, error) {

	var buf bytes.Buffer
	auth := struct {
		User     string `json:"username"`
		Password string `json:"password"`
	}{}
	auth.User = m.user
	auth.Password = m.password

	if err := json.NewEncoder(&buf).Encode(auth); err != nil {
		return "", err
	}
	resp, err := http.Post(m.url+"/rest/V1/integration/admin/token", "application/json", &buf)
	if err != nil {
		return "", nil
	}

	var res string
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}
