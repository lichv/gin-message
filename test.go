package main

import (
	"fmt"
	"gin-message/app/models"
	"gin-message/utils"
	"gin-message/utils/setting"
)

func init() {
	setting.Setup()
	models.Setup()

}
func main() {
	//var buff bytes.Buffer
	//
	//stu := map[string]interface{}{"Name":"hello","Id":11}
	//
	//tmpl, err := template.New("test").Parse("{{.Name}} ID is {{ .Id }}")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = tmpl.Execute(&buff, stu)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(buff.String())
	result, err := utils.ParseTemplateWithParams("{{Name}} ID is {{ ID}}", map[string]interface{}{"Name": 223, "ID": 34534})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	//reg := regexp.MustCompile(`{{\s*(.*?)\s*}}`)
	//var content = "{{Name}} ID is {{ Id }}"
	//submatchs := reg.FindAllStringSubmatch(content, -1)
	//for _,submatch := range submatchs{
	//	fmt.Println(submatch)
	//	tmp,_ := regexp.Compile(submatch[1])
	//	content = tmp.ReplaceAllString(content,"."+submatch[1])
	//}
	//fmt.Println(content)
	//all := reg.ReplaceAll([]byte(content), []byte(`{{\s*\.(.*?)\s}}`))
	//all := reg.ReplaceAllLiteralString(content,`{{\s*\.(.*?)\s}}`)
	//fmt.Println(string(all))

}
