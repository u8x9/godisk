package token

import (
	"fmt"
	"time"

	"github.com/u8x9/godisk/util"
)

const tokenSalt = "#1(22]]^DFzd"

// GenToken 生成token
func GenToken(username string) string {
	// md5(username + timestamp + tokenSalt) + timestamp[:8]
	// len(token) = 40
	ts := fmt.Sprintf("%x", time.Now().Unix())
	key := fmt.Sprintf("%s%s%s", username, ts, tokenSalt)
	token := util.MD5([]byte(key)) + ts[:8]
	return token
}
