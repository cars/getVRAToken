package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type logonMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain,omitempty"`
}

type logonResponse struct {
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func main() {
	userPtr := flag.String("username", "administrator", "User Name")
	domainPtr := flag.CommandLine.String("domain", "", "Domain to log into")
	serverPtr := flag.String("server", "api.mgmt.cloud.vmware.com", "FQDN/Hostname of vRA server")
	passwordPtr := flag.String("password", "", "Password")
	flag.Parse()

	logonJSON, err := json.Marshal(logonMessage{*userPtr, *passwordPtr, *domainPtr})

	if err != nil {
		fmt.Print("Error occurred")
	}

	postURL := fmt.Sprintf("https://%s/csp/gateway/am/api/login?access_token", *serverPtr)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(logonJSON))
	if err != nil {
		log.Fatal(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	respBody, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Non-200 Status Code received", resp.StatusCode, string(respBody[:]))
		log.Fatal("Non-200 Status Code received")
	}

	var logonResp logonResponse
	json.Unmarshal(respBody, &logonResp)
	fmt.Print(logonResp.RefreshToken)
}
