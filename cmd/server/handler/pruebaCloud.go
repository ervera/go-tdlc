package handler

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

const (
	cloudname        = "dangvuvyq"
	cloudapikey      = "754218821349648"
	cloudinarysecret = "rjiWDoS5G0yNdiY4NZkEXtvit8k"
)

func (c *userHandler) ABC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, _, _ := ctx.Request.FormFile("image")
		cld, _ := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
		options := uploader.UploadParams{}
		resp, _ := cld.Upload.Upload(ctx, file, options)
		fmt.Println(resp)
	}
}

// &{
// 	AssetID:aac8fd89108ff834dd27f979fd9ce77e
// 	PublicID:hl22acprlomnycgiudor
// 	Version:1591095352
// 	VersionID:909700634231dbaaf8b06d7a5940299e
// 	Signature:86922996d63e596464ea3d7a5e86e8de8123f23f
// 	Width:1200
// 	Height:1200
// 	Format:jpg
// 	ResourceType:image
// 	CreatedAt:2020-06-02 10:55:52 +0000 UTC
// 	Tags:[]
// 	Pages:0
// 	Bytes:460268
// 	Type:upload
// 	Etag:2c7e88604ba3f340a0c5bc8cd418b4d9
// 	Placeholder:false
// 	URL:  http://res.cloudinary.com/demo/image/upload/v1591095352/hl22acprlomnycgiudor.jpg
// 	SecureURL:  https://res.cloudinary.com/demo/image/upload/v1591095352/hl22acprlomnycgiudor.jpg
// 	AccessMode:public
// 	Context:map[]
// 	Metadata:map[]
// 	Overwritten:true
// 	OriginalFilename:my_image
// 	Error:{Message:}
//   }
