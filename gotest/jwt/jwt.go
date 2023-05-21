/**
 * @Author: xueyanghan
 * @File: jwt.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/4/1 13:37
 */

package main

import "golang.org/x/oauth2/jwt"

func main() {
	jwtCfg := jwt.Config{}
	jwtCfg.TokenURL = "https://www.googleapis.com/oauth2/v4/token"
}
