package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1/"

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	/*
		   Private API の呼出には認証が必要です。
		   ログイン後、開発者ページ において発行した API key と API secret を使用します （API key をご利用いただけるのは、bitFlyer Lightning をご利用可能なお客様のみとなります）。

		   以下の情報を HTTP リクエストヘッダに含めます。

		   ACCESS-KEY: 開発者ページで発行した API key
		   ACCESS-TIMESTAMP: リクエスト時の Unix Timestamp
		   ACCESS-SIGN: 以下の方法でリクエストごとに生成した署名
		   ACCESS-SIGN は、ACCESS-TIMESTAMP, HTTP メソッド, リクエストのパス, リクエストボディ を文字列として連結したものを、 API secret で HMAC-SHA256 署名を行った結果です。

			 // Node.js のサンプル
			var request = require('request');
			var crypto = require('crypto');

			var key = '{{ YOUR API KEY }}';
			var secret = '{{ YOUR API SECRET }}';

			var timestamp = Date.now().toString();
			var method = 'POST';
			var path = '/v1/me/sendchildorder';
			var body = JSON.stringify({
					product_code: 'BTC_JPY',
					child_order_type: 'LIMIT',
					side: 'BUY',
					price: 30000,
					size: 0.1
			});

			var text = timestamp + method + path + body;
			var sign = crypto.createHmac('sha256', secret).update(text).digest('hex');

			var options = {
					url: 'https://api.bitflyer.com' + path,
					method: method,
					body: body,
					headers: {
							'ACCESS-KEY': key,
							'ACCESS-TIMESTAMP': timestamp,
							'ACCESS-SIGN': sign,
							'Content-Type': 'application/json'
					}
			};
			request(options, function (err, response, payload) {
					console.log(payload);
			});
	*/
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) //10進数
	log.Panicln(timestamp)
	message := timestamp + endpoint + string(body)
	h := hmac.New(sha256.New, []byte(api.secret))
	h.Write([]byte(message))
	sign := hex.EncodeToString(h.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL) //正しいURLか確認
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Panicf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
