package plugin

import (
	"github.com/ChengjinWu/gojson"
	"github.com/gin-gonic/gin"
)

func Res401(c *gin.Context) {
	c.JSON(401, gin.H{"err": "permission denied"})
}
func Res401str(c *gin.Context, str string) {
	c.JSON(401, gin.H{"err": str})
}
func Res403(c *gin.Context) {
	c.JSON(403, gin.H{"err": "forbidden"})
}
func Res403str(c *gin.Context, str string) {
	c.JSON(403, gin.H{"err": str})
}
func Res404(c *gin.Context) {
	c.JSON(404, gin.H{"err": "not found"})
}
func Res404str(c *gin.Context, str string) {
	c.JSON(404, gin.H{"err": str})
}

func Res406(c *gin.Context) {
	c.JSON(406, gin.H{"err": "request type error"})
}
func Res406str(c *gin.Context, str string) {
	c.JSON(406, gin.H{"err": str})
}

func Res500(c *gin.Context) {
	c.JSON(500, gin.H{"err": "server error"})
}

func Res500str(c *gin.Context, str string) {
	c.JSON(500, gin.H{"err": str})
}

func RequestJsonCheck(c *gin.Context) (data []byte, msg string, err error) {
	data, err = c.GetRawData()
	if err != nil {
		msg = "get raw data faild"
		return
	}
	err = gojson.CheckValid([]byte(data))
	if err != nil {
		msg = "get raw data json check faild"
		return
	}
	if len(data) == 0 {
		msg = "get raw data len check faild"
		return
	}
	return
}

func RequestData(c *gin.Context) (data []byte, msg string, err error) {
	data, err = c.GetRawData()
	if err != nil {
		msg = "get raw data faild"
		return
	}
	return
}

// func RecordMake(resource string, action string, args string, state_before string, state_after string, project_id int64) error {
// 	record := model.Record{
// 		Action:       action,
// 		Resource:     resource,
// 		Args:         args,
// 		State_before: state_before,
// 		State_after:  state_after,
// 		Project_id:   project_id,
// 	}
// 	err := record.Add()
// 	return err
// }

//map get keys type string
func mapGetKeysString(m map[string]interface{}) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func mapListStringCompare(a []string, b []string) (flag int64) {
	flag = 1
	a_len := len(a)
	b_len := len(b)
	if a_len == b_len {
		for i, v := range a {
			if v != b[i] {
				flag = 0
			}
		}
	} else {
		flag = 0
	}
	return flag
}
