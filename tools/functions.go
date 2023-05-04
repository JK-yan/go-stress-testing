package tools

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var Functions = map[string]interface{}{
	"get_timestamp": getTimestamp, // call without arguments
	"get_uuid":      getUUID,
	"get_oid":       getOid,
	"get_oid_one":   getOidOneByOne,
}
var IndexId = 0
var oids []string
var mutex sync.Mutex

func init() {
	rand.Seed(time.Now().UnixNano())

}
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func getUUID() uuid.UUID {
	return uuid.NewV4()
}

// MergeMap 以master 为主合并second。
func MergeMap(master map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	for k, v := range second {
		if _, ok := master[k]; !ok {
			master[k] = v
		}
	}
	return master
}

func ReadCsv(filePath string) []string {
	f := readFile(filePath)
	return strings.Split(string(f), ",")
}
func SetOid(oid []string) {
	oids = oid
}

func readFile(filePath string) []byte {
	r, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return r
}

func getOid() string {
	oid := []string{"630f0dff5d2b2a0001235e69", "630f0e015d2b2a0001236062", "630f0e035d2b2a000123625b", "630f0e055d2b2a0001236454", "630f0e075d2b2a000123664d", "630f0e095d2b2a0001236846", "630f0e0b5d2b2a0001236a3f", "630f0e0d5d2b2a0001236c38", "630f0e0f5d2b2a0001236e31", "630f0e105d2b2a000123702a", "630f0e125d2b2a0001237223", "630f0e145d2b2a000123741c", "630f0e165d2b2a0001237615", "630f0e185d2b2a000123780e", "630f0e195d2b2a0001237a07", "630f0e1b5d2b2a0001237c00", "630f0e1d5d2b2a0001237df9", "630f0e1f5d2b2a0001237ff2", "630f0e205d2b2a00012381eb", "630f0e225d2b2a00012383e4", "630f0e245d2b2a00012385dd", "630f0e255d2b2a00012387d6", "630f0e275d2b2a00012389cf", "630f0e295d2b2a0001238bc8", "630f0e2a5d2b2a0001238dc1", "630f0e2c5d2b2a0001238fba", "630f0e2e5d2b2a00012391b3", "630f0e2f5d2b2a00012393ac", "630f0e315d2b2a00012395a5", "630f0e335d2b2a000123979e", "630f0e345d2b2a0001239997", "630f0e365d2b2a0001239b90", "630f0e385d2b2a0001239d89", "630f0e3a5d2b2a0001239f82", "630f0e3b5d2b2a000123a17b", "630f0e3d5d2b2a000123a374", "630f0e3e5d2b2a000123a56d", "630f0e405d2b2a000123a766", "630f0e425d2b2a000123a95f", "630f0e435d2b2a000123ab58", "630f0e455d2b2a000123ad51", "630f0e465d2b2a000123af4a", "630f0e485d2b2a000123b143", "630f0e495d2b2a000123b33c", "630f0e4b5d2b2a000123b535", "630f0e4c5d2b2a000123b72e", "630f0e4e5d2b2a000123b927", "630f0e505d2b2a000123bb20", "630f0e515d2b2a000123bd19", "630f0e535d2b2a000123bf12", "630f0e545d2b2a000123c10b", "630f0e565d2b2a000123c304", "630f0e575d2b2a000123c4fd", "630f0e595d2b2a000123c6f6", "630f0e5b5d2b2a000123c8ef", "630f0e5c5d2b2a000123cae8", "630f0e5e5d2b2a000123cce1", "630f0e5f5d2b2a000123ceda", "630f0e615d2b2a000123d0d3", "630f0e625d2b2a000123d2cc", "630f0e645d2b2a000123d4c5", "630f0e655d2b2a000123d6be", "630f0e675d2b2a000123d8b7", "630f0e685d2b2a000123dab0", "630f0e6a5d2b2a000123dca9", "630f0e6b5d2b2a000123dea2", "630f0e6d5d2b2a000123e09b", "630f0e6e5d2b2a000123e294", "630f0e705d2b2a000123e48d", "630f0e725d2b2a000123e686", "630f0e735d2b2a000123e87f", "630f0e755d2b2a000123ea78", "630f0e765d2b2a000123ec71", "630f0e785d2b2a000123ee6a", "630f0e795d2b2a000123f063", "630f0e7b5d2b2a000123f25c", "630f0e7c5d2b2a000123f455", "630f0e7e5d2b2a000123f64e", "630f0e805d2b2a000123f847", "630f0e815d2b2a000123fa40", "630f0e835d2b2a000123fc39", "630f0e845d2b2a000123fe32", "630f0e865d2b2a000124002b", "630f0e875d2b2a0001240224", "630f0e895d2b2a000124041d", "630f0e8a5d2b2a0001240616", "630f0e8c5d2b2a000124080f", "630f0e8d5d2b2a0001240a08", "630f0e8f5d2b2a0001240c01", "630f0e905d2b2a0001240dfa", "630f0e925d2b2a0001240ff3", "630f0e945d2b2a00012411ec", "630f0e955d2b2a00012413e5", "630f0e975d2b2a00012415de", "630f0e985d2b2a00012417d7", "630f0e9a5d2b2a00012419d0", "630f0e9b5d2b2a0001241bc9", "630f0e9d5d2b2a0001241dc2", "630f0e9e5d2b2a0001241fbb", "630f0ea05d2b2a00012421b4"}
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(len(oid))
	return oid[l]
}

func getOidOneByOne() string {
	defer mutex.Unlock()
	mutex.Lock()
	if IndexId < len(oids) {
		id := oids[IndexId]
		IndexId = IndexId + 1
		return strings.TrimSpace(id)
	} else {
		id := oids[0]
		IndexId = 1
		return strings.TrimSpace(id)
	}

}
